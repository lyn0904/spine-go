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

type BindingManager struct {
	localDevice api.DeviceLocalInterface

	bindingNum     uint64
	bindingEntries []*api.BindingEntry

	mux sync.Mutex
	// TODO: add persistence
}

func NewBindingManager(localDevice api.DeviceLocalInterface) *BindingManager {
	c := &BindingManager{
		bindingNum:  0,
		localDevice: localDevice,
	}

	return c
}

// is sent from the client (remote device) to the server (local device)
func (c *BindingManager) AddBinding(remoteDevice api.DeviceRemoteInterface, data model.BindingManagementRequestCallType) error {
	serverFeature := c.localDevice.FeatureByAddress(data.ServerAddress)
	if serverFeature == nil {
		return fmt.Errorf("server feature '%s' in local device '%s' not found", data.ServerAddress, *c.localDevice.Address())
	}
	if data.ServerFeatureType == nil {
		return errors.New("serverFeatureType is missing but required")
	}
	if err := c.checkRoleAndType(serverFeature, model.RoleTypeServer, *data.ServerFeatureType); err != nil {
		return err
	}

	// a local feature can only have one remote binding
	bindings := c.BindingsOnFeature(*serverFeature.Address())
	if len(bindings) > 0 {
		return errors.New("the server feature already has a binding")
	}

	clientFeature := remoteDevice.FeatureByAddress(data.ClientAddress)
	if clientFeature == nil {
		return fmt.Errorf("client feature '%s' in remote device '%s' not found", data.ClientAddress, *remoteDevice.Address())
	}
	if err := c.checkRoleAndType(clientFeature, model.RoleTypeClient, *data.ServerFeatureType); err != nil {
		return err
	}

	bindingEntry := &api.BindingEntry{
		Id:            c.bindingId(),
		ServerFeature: serverFeature,
		ClientFeature: clientFeature,
	}

	c.mux.Lock()
	defer c.mux.Unlock()

	c.bindingEntries = append(c.bindingEntries, bindingEntry)

	payload := api.EventPayload{
		Ski:          remoteDevice.Ski(),
		EventType:    api.EventTypeBindingChange,
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

func (c *BindingManager) RemoveBinding(data model.BindingManagementDeleteCallType, remoteDevice api.DeviceRemoteInterface) error {
	var newBindingEntries []*api.BindingEntry

	// according to the spec 7.4.4
	// a. The absence of "bindingDelete. clientAddress. device" SHALL be treated as if it was
	//    present and set to the sender's "device" address part.
	// b. The absence of "bindingDelete. serverAddress. device" SHALL be treated as if it was
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

	if err := c.checkRoleAndType(serverFeature, model.RoleTypeServer, serverFeature.Type()); err != nil {
		return err
	}

	if !c.HasLocalFeatureRemoteBinding(serverFeature.Address(), clientFeature.Address()) {
		return fmt.Errorf("the feature '%s' address has no binding", data.ClientAddress)
	}

	c.mux.Lock()
	defer c.mux.Unlock()

	for _, item := range c.bindingEntries {
		itemAddress := item.ClientFeature.Address()

		if !reflect.DeepEqual(*itemAddress, clientAddress) &&
			!reflect.DeepEqual(item.ServerFeature, serverFeature) {
			newBindingEntries = append(newBindingEntries, item)
		}
	}

	if len(newBindingEntries) == len(c.bindingEntries) {
		return errors.New("could not find requested binding to be removed")
	}

	c.bindingEntries = newBindingEntries

	payload := api.EventPayload{
		Ski:          remoteDevice.Ski(),
		EventType:    api.EventTypeBindingChange,
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

// Remove all existing bindings for a given remote device
func (c *BindingManager) RemoveBindingsForDevice(remoteDevice api.DeviceRemoteInterface) {
	if remoteDevice == nil {
		return
	}

	for _, entity := range remoteDevice.Entities() {
		c.RemoveBindingsForEntity(entity)
	}
}

// Remove all existing bindings for a given remote device entity
func (c *BindingManager) RemoveBindingsForEntity(remoteEntity api.EntityRemoteInterface) {
	if remoteEntity == nil {
		return
	}

	c.mux.Lock()
	defer c.mux.Unlock()

	var newBindingEntries []*api.BindingEntry
	for _, item := range c.bindingEntries {
		if !reflect.DeepEqual(item.ClientFeature.Address().Entity, remoteEntity.Address().Entity) {
			newBindingEntries = append(newBindingEntries, item)
			continue
		}

		serverFeature := c.localDevice.FeatureByAddress(item.ServerFeature.Address())
		clientFeature := remoteEntity.FeatureOfAddress(item.ClientFeature.Address().Feature)
		payload := api.EventPayload{
			Ski:          remoteEntity.Device().Ski(),
			EventType:    api.EventTypeBindingChange,
			ChangeType:   api.ElementChangeRemove,
			Device:       remoteEntity.Device(),
			Entity:       remoteEntity,
			Feature:      clientFeature,
			LocalFeature: serverFeature,
		}
		Events.Publish(payload)
	}

	c.bindingEntries = newBindingEntries
}

func (c *BindingManager) Bindings(remoteDevice api.DeviceRemoteInterface) []*api.BindingEntry {
	var result []*api.BindingEntry

	c.mux.Lock()
	defer c.mux.Unlock()

	linq.From(c.bindingEntries).WhereT(func(s *api.BindingEntry) bool {
		return s.ClientFeature.Device().Ski() == remoteDevice.Ski()
	}).ToSlice(&result)

	return result
}

// checks if a remote address has a binding on the local feature
func (c *BindingManager) HasLocalFeatureRemoteBinding(localAddress, remoteAddress *model.FeatureAddressType) bool {
	bindings := c.BindingsOnFeature(*localAddress)

	for _, item := range bindings {
		if reflect.DeepEqual(item.ClientFeature.Address(), remoteAddress) {
			return true
		}
	}

	return false
}

func (c *BindingManager) BindingsOnFeature(featureAddress model.FeatureAddressType) []*api.BindingEntry {
	var result []*api.BindingEntry

	c.mux.Lock()
	defer c.mux.Unlock()

	linq.From(c.bindingEntries).WhereT(func(s *api.BindingEntry) bool {
		return reflect.DeepEqual(*s.ServerFeature.Address(), featureAddress)
	}).ToSlice(&result)

	return result
}

func (c *BindingManager) bindingId() uint64 {
	i := atomic.AddUint64(&c.bindingNum, 1)
	return i
}

func (c *BindingManager) checkRoleAndType(feature api.FeatureInterface, role model.RoleType, featureType model.FeatureTypeType) error {
	if feature.Role() != model.RoleTypeSpecial && feature.Role() != role {
		return fmt.Errorf("found feature %s is not matching required role %s", feature.Type(), role)
	}

	if feature.Type() != featureType && feature.Type() != model.FeatureTypeTypeGeneric {
		return fmt.Errorf("found feature %s is not matching required type %s", feature.Type(), featureType)
	}

	return nil
}
