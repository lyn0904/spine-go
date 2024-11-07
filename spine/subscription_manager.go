package spine

import (
	"errors"
	"fmt"
	"reflect"
	"sync"
	"sync/atomic"

	"github.com/ahmetb/go-linq/v3"
	"github.com/lyn0904/spine-go/api"
	"github.com/lyn0904/spine-go/model"
	"github.com/lyn0904/spine-go/util"
)

type SubscriptionManager struct {
	localDevice api.DeviceLocalInterface

	subscriptionNum     uint64
	subscriptionEntries []*api.SubscriptionEntry

	mux sync.Mutex
	// TODO: add persistence
}

func NewSubscriptionManager(localDevice api.DeviceLocalInterface) *SubscriptionManager {
	c := &SubscriptionManager{
		subscriptionNum: 0,
		localDevice:     localDevice,
	}

	return c
}

// is sent from the client (remote device) to the server (local device)
func (c *SubscriptionManager) AddSubscription(remoteDevice api.DeviceRemoteInterface, data model.SubscriptionManagementRequestCallType) error {
	serverFeature := c.localDevice.FeatureByAddress(data.ServerAddress)
	if serverFeature == nil {
		return fmt.Errorf("server feature '%s' in local device '%s' not found", data.ServerAddress, *c.localDevice.Address())
	}
	if err := c.checkRoleAndType(serverFeature, model.RoleTypeServer, *data.ServerFeatureType); err != nil {
		return err
	}

	clientFeature := remoteDevice.FeatureByAddress(data.ClientAddress)
	if clientFeature == nil {
		return fmt.Errorf("client feature '%s' in remote device '%s' not found", data.ClientAddress, *remoteDevice.Address())
	}
	if err := c.checkRoleAndType(clientFeature, model.RoleTypeClient, *data.ServerFeatureType); err != nil {
		return err
	}

	subscriptionEntry := &api.SubscriptionEntry{
		Id:            c.subscriptionId(),
		ServerFeature: serverFeature,
		ClientFeature: clientFeature,
	}

	c.mux.Lock()
	defer c.mux.Unlock()

	for _, item := range c.subscriptionEntries {
		if reflect.DeepEqual(item.ServerFeature, serverFeature) && reflect.DeepEqual(item.ClientFeature, clientFeature) {
			return fmt.Errorf("requested subscription is already present")
		}
	}

	c.subscriptionEntries = append(c.subscriptionEntries, subscriptionEntry)

	payload := api.EventPayload{
		Ski:          remoteDevice.Ski(),
		EventType:    api.EventTypeSubscriptionChange,
		ChangeType:   api.ElementChangeAdd,
		Data:         data,
		Device:       remoteDevice,
		Entity:       clientFeature.Entity(),
		Feature:      clientFeature,
		LocalFeature: serverFeature,
	}
	Events.Publish(payload)

	return nil
}

// Remove a specific subscription that is provided by a delete message from a remote device
func (c *SubscriptionManager) RemoveSubscription(data model.SubscriptionManagementDeleteCallType, remoteDevice api.DeviceRemoteInterface) error {
	var newSubscriptionEntries []*api.SubscriptionEntry

	// according to the spec 7.4.4
	// a. The absence of "subscriptionDelete. clientAddress. device" SHALL be treated as if it was
	//    present and set to the sender's "device" address part.
	// b. The absence of "subscriptionDelete. serverAddress. device" SHALL be treated as if it was
	//    present and set to the recipient's "device" address part.

	var clientAddress model.FeatureAddressType
	util.DeepCopy(data.ClientAddress, &clientAddress)
	if data.ClientAddress.Device == nil {
		clientAddress.Device = remoteDevice.Address()
	}

	clientFeature := remoteDevice.FeatureByAddress(data.ClientAddress)
	if clientFeature == nil {
		return fmt.Errorf("client feature '%s' in remote device '%s' not found", data.ClientAddress, *remoteDevice.Address())
	}

	serverFeature := c.localDevice.FeatureByAddress(data.ServerAddress)
	if serverFeature == nil {
		return fmt.Errorf("server feature '%s' in local device '%s' not found", data.ServerAddress, *c.localDevice.Address())
	}

	c.mux.Lock()
	defer c.mux.Unlock()

	for _, item := range c.subscriptionEntries {
		itemAddress := item.ClientFeature.Address()

		if !reflect.DeepEqual(itemAddress.Device, clientAddress.Device) ||
			!reflect.DeepEqual(itemAddress.Entity, clientAddress.Entity) ||
			!reflect.DeepEqual(itemAddress.Feature, clientAddress.Feature) ||
			!reflect.DeepEqual(item.ServerFeature, serverFeature) {
			newSubscriptionEntries = append(newSubscriptionEntries, item)
		}
	}

	if len(newSubscriptionEntries) == len(c.subscriptionEntries) {
		return errors.New("could not find requested SubscriptionId to be removed")
	}

	c.subscriptionEntries = newSubscriptionEntries

	payload := api.EventPayload{
		Ski:          remoteDevice.Ski(),
		EventType:    api.EventTypeSubscriptionChange,
		ChangeType:   api.ElementChangeRemove,
		Data:         data,
		Device:       remoteDevice,
		Entity:       clientFeature.Entity(),
		Feature:      clientFeature,
		LocalFeature: serverFeature,
	}
	Events.Publish(payload)

	return nil
}

// Remove all existing subscriptions for a given remote device
func (c *SubscriptionManager) RemoveSubscriptionsForDevice(remoteDevice api.DeviceRemoteInterface) {
	if remoteDevice == nil {
		return
	}

	for _, entity := range remoteDevice.Entities() {
		c.RemoveSubscriptionsForEntity(entity)
	}
}

// Remove all existing subscriptions for a given remote device entity
func (c *SubscriptionManager) RemoveSubscriptionsForEntity(remoteEntity api.EntityRemoteInterface) {
	if remoteEntity == nil {
		return
	}

	c.mux.Lock()
	defer c.mux.Unlock()

	var newSubscriptionEntries []*api.SubscriptionEntry
	for _, item := range c.subscriptionEntries {
		if !reflect.DeepEqual(item.ClientFeature.Address().Device, remoteEntity.Address().Device) ||
			!reflect.DeepEqual(item.ClientFeature.Address().Entity, remoteEntity.Address().Entity) {
			newSubscriptionEntries = append(newSubscriptionEntries, item)
			continue
		}

		serverFeature := c.localDevice.FeatureByAddress(item.ServerFeature.Address())
		clientFeature := remoteEntity.FeatureOfAddress(item.ClientFeature.Address().Feature)
		payload := api.EventPayload{
			Ski:          remoteEntity.Device().Ski(),
			EventType:    api.EventTypeSubscriptionChange,
			ChangeType:   api.ElementChangeRemove,
			Device:       remoteEntity.Device(),
			Entity:       remoteEntity,
			Feature:      clientFeature,
			LocalFeature: serverFeature,
		}
		Events.Publish(payload)
	}

	c.subscriptionEntries = newSubscriptionEntries
}

func (c *SubscriptionManager) Subscriptions(remoteDevice api.DeviceRemoteInterface) []*api.SubscriptionEntry {
	var result []*api.SubscriptionEntry

	c.mux.Lock()
	defer c.mux.Unlock()

	linq.From(c.subscriptionEntries).WhereT(func(s *api.SubscriptionEntry) bool {
		return s.ClientFeature.Device().Ski() == remoteDevice.Ski()
	}).ToSlice(&result)

	return result
}

func (c *SubscriptionManager) SubscriptionsOnFeature(featureAddress model.FeatureAddressType) []*api.SubscriptionEntry {
	var result []*api.SubscriptionEntry

	c.mux.Lock()
	defer c.mux.Unlock()

	linq.From(c.subscriptionEntries).WhereT(func(s *api.SubscriptionEntry) bool {
		return reflect.DeepEqual(*s.ServerFeature.Address(), featureAddress)
	}).ToSlice(&result)

	return result
}

func (c *SubscriptionManager) subscriptionId() uint64 {
	i := atomic.AddUint64(&c.subscriptionNum, 1)
	return i
}

func (c *SubscriptionManager) checkRoleAndType(feature api.FeatureInterface, role model.RoleType, featureType model.FeatureTypeType) error {
	if feature.Role() != model.RoleTypeSpecial && feature.Role() != role {
		return fmt.Errorf("found feature %s is not matching required role %s", feature.Type(), role)
	}

	if feature.Type() != featureType && feature.Type() != model.FeatureTypeTypeGeneric {
		return fmt.Errorf("found feature %s is not matching required type %s", feature.Type(), featureType)
	}

	return nil
}
