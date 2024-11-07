package spine

import (
	"sync"
	"sync/atomic"
	"time"

	"github.com/lyn0904/spine-go/api"
	"github.com/lyn0904/spine-go/model"
)

type HeartbeatManager struct {
	localEntity  api.EntityLocalInterface
	localFeature api.FeatureLocalInterface

	heartBeatNum   uint64 // see https://github.com/golang/go/issues/11891
	stopHeartbeatC chan struct{}
	stopMux        sync.Mutex

	heartBeatTimeout *model.DurationType

	mux sync.Mutex
}

var _ api.HeartbeatManagerInterface = (*HeartbeatManager)(nil)

// Create a new Heartbeat Manager which handles sending of heartbeats
func NewHeartbeatManager(localEntity api.EntityLocalInterface, timeout time.Duration) *HeartbeatManager {
	h := &HeartbeatManager{
		localEntity:      localEntity,
		heartBeatTimeout: model.NewDurationType(timeout),
	}

	return h
}

func (c *HeartbeatManager) IsHeartbeatRunning() bool {
	c.stopMux.Lock()
	defer c.stopMux.Unlock()

	if c.stopHeartbeatC != nil && !c.isHeartbeatClosed() {
		return true
	}

	return false
}

func (c *HeartbeatManager) SetLocalFeature(entity api.EntityLocalInterface, feature api.FeatureLocalInterface) {
	if entity == nil || feature == nil {
		return
	}

	if feature.Type() != model.FeatureTypeTypeDeviceDiagnosis ||
		feature.Role() != model.RoleTypeServer {
		return
	}

	// check if the local device diagnosis server feature, supports the heartbeat function
	ops, ok := feature.Operations()[model.FunctionTypeDeviceDiagnosisHeartbeatData]
	if !ok || !ops.Read() {
		return
	}

	c.mux.Lock()

	c.localEntity = entity
	c.localFeature = feature

	// initialise heartbeat data
	heartbeatData := c.heartbeatData(time.Now().UTC(), c.heartBeatCounter())

	// updating the data will automatically notify all subscribed remote features
	feature.SetData(model.FunctionTypeDeviceDiagnosisHeartbeatData, heartbeatData)

	c.mux.Unlock()

	// start creating heartbeats
	_ = c.StartHeartbeat()
}

// Start setting heartbeat data
// Make sure the a required FeatureTypeTypeDeviceDiagnosis with the role server is present
// otherwise this will end with an error
// Note: Remote features need to have a subscription to get notifications
func (c *HeartbeatManager) StartHeartbeat() error {
	timeout, err := c.heartBeatTimeout.GetTimeDuration()
	if err != nil {
		return err
	}

	// stop an already running heartbeat
	c.StopHeartbeat()

	c.stopHeartbeatC = make(chan struct{})

	go c.updateHeartbeatData(c.stopHeartbeatC, timeout)

	return nil
}

// Stop updating heartbeat data
// Note: No active subscribers will get any further notifications!
func (c *HeartbeatManager) StopHeartbeat() {
	if c.IsHeartbeatRunning() {
		close(c.stopHeartbeatC)
	}
}

func (c *HeartbeatManager) heartbeatData(t time.Time, counter *uint64) *model.DeviceDiagnosisHeartbeatDataType {
	timestamp := model.NewAbsoluteOrRelativeTimeTypeFromTime(t)

	return &model.DeviceDiagnosisHeartbeatDataType{
		Timestamp:        timestamp,
		HeartbeatCounter: counter,
		HeartbeatTimeout: c.heartBeatTimeout,
	}
}

func (c *HeartbeatManager) updateHeartbeatData(stopC chan struct{}, d time.Duration) {
	// Substract two seconds, because some devices (like Elli Connect/Pro) with OPEV/OSCEV interpret
	// the heartbeat timeout (<= 4s) as the time within which a heartbeat should be received and otherwise
	// will go into fallback mode.
	// But other EVSE devices and in LPC (<= 60s), the heartbeat should be considered missing, if it is not
	// received within twice the heartbeat timeout timeframe.
	if d > 2*time.Second {
		d -= 2 * time.Second
	}
	ticker := time.NewTicker(d)
	for {
		select {
		case <-ticker.C:

			heartbeatData := c.heartbeatData(time.Now().UTC(), c.heartBeatCounter())

			c.mux.Lock()
			// updating the data will automatically notify all subscribed remote features
			c.localFeature.SetData(model.FunctionTypeDeviceDiagnosisHeartbeatData, heartbeatData)
			c.mux.Unlock()

		case <-stopC:
			return
		}
	}
}

func (c *HeartbeatManager) isHeartbeatClosed() bool {
	select {
	case <-c.stopHeartbeatC:
		return true
	default:
	}

	return false
}

// TODO heartBeatCounter should be global on CEM level, not on connection level
func (c *HeartbeatManager) heartBeatCounter() *uint64 {
	i := atomic.AddUint64(&c.heartBeatNum, 1)
	return &i
}
