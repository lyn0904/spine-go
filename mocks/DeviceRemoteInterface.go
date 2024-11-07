// Code generated by mockery v2.45.0. DO NOT EDIT.

package mocks

import (
	api "github.com/lyn0904/spine-go/api"
	mock "github.com/stretchr/testify/mock"

	model "github.com/lyn0904/spine-go/model"
)

// DeviceRemoteInterface is an autogenerated mock type for the DeviceRemoteInterface type
type DeviceRemoteInterface struct {
	mock.Mock
}

type DeviceRemoteInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *DeviceRemoteInterface) EXPECT() *DeviceRemoteInterface_Expecter {
	return &DeviceRemoteInterface_Expecter{mock: &_m.Mock}
}

// AddEntity provides a mock function with given fields: entity
func (_m *DeviceRemoteInterface) AddEntity(entity api.EntityRemoteInterface) {
	_m.Called(entity)
}

// DeviceRemoteInterface_AddEntity_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddEntity'
type DeviceRemoteInterface_AddEntity_Call struct {
	*mock.Call
}

// AddEntity is a helper method to define mock.On call
//   - entity api.EntityRemoteInterface
func (_e *DeviceRemoteInterface_Expecter) AddEntity(entity interface{}) *DeviceRemoteInterface_AddEntity_Call {
	return &DeviceRemoteInterface_AddEntity_Call{Call: _e.mock.On("AddEntity", entity)}
}

func (_c *DeviceRemoteInterface_AddEntity_Call) Run(run func(entity api.EntityRemoteInterface)) *DeviceRemoteInterface_AddEntity_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(api.EntityRemoteInterface))
	})
	return _c
}

func (_c *DeviceRemoteInterface_AddEntity_Call) Return() *DeviceRemoteInterface_AddEntity_Call {
	_c.Call.Return()
	return _c
}

func (_c *DeviceRemoteInterface_AddEntity_Call) RunAndReturn(run func(api.EntityRemoteInterface)) *DeviceRemoteInterface_AddEntity_Call {
	_c.Call.Return(run)
	return _c
}

// AddEntityAndFeatures provides a mock function with given fields: initialData, data
func (_m *DeviceRemoteInterface) AddEntityAndFeatures(initialData bool, data *model.NodeManagementDetailedDiscoveryDataType) ([]api.EntityRemoteInterface, error) {
	ret := _m.Called(initialData, data)

	if len(ret) == 0 {
		panic("no return value specified for AddEntityAndFeatures")
	}

	var r0 []api.EntityRemoteInterface
	var r1 error
	if rf, ok := ret.Get(0).(func(bool, *model.NodeManagementDetailedDiscoveryDataType) ([]api.EntityRemoteInterface, error)); ok {
		return rf(initialData, data)
	}
	if rf, ok := ret.Get(0).(func(bool, *model.NodeManagementDetailedDiscoveryDataType) []api.EntityRemoteInterface); ok {
		r0 = rf(initialData, data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]api.EntityRemoteInterface)
		}
	}

	if rf, ok := ret.Get(1).(func(bool, *model.NodeManagementDetailedDiscoveryDataType) error); ok {
		r1 = rf(initialData, data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeviceRemoteInterface_AddEntityAndFeatures_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddEntityAndFeatures'
type DeviceRemoteInterface_AddEntityAndFeatures_Call struct {
	*mock.Call
}

// AddEntityAndFeatures is a helper method to define mock.On call
//   - initialData bool
//   - data *model.NodeManagementDetailedDiscoveryDataType
func (_e *DeviceRemoteInterface_Expecter) AddEntityAndFeatures(initialData interface{}, data interface{}) *DeviceRemoteInterface_AddEntityAndFeatures_Call {
	return &DeviceRemoteInterface_AddEntityAndFeatures_Call{Call: _e.mock.On("AddEntityAndFeatures", initialData, data)}
}

func (_c *DeviceRemoteInterface_AddEntityAndFeatures_Call) Run(run func(initialData bool, data *model.NodeManagementDetailedDiscoveryDataType)) *DeviceRemoteInterface_AddEntityAndFeatures_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(bool), args[1].(*model.NodeManagementDetailedDiscoveryDataType))
	})
	return _c
}

func (_c *DeviceRemoteInterface_AddEntityAndFeatures_Call) Return(_a0 []api.EntityRemoteInterface, _a1 error) *DeviceRemoteInterface_AddEntityAndFeatures_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *DeviceRemoteInterface_AddEntityAndFeatures_Call) RunAndReturn(run func(bool, *model.NodeManagementDetailedDiscoveryDataType) ([]api.EntityRemoteInterface, error)) *DeviceRemoteInterface_AddEntityAndFeatures_Call {
	_c.Call.Return(run)
	return _c
}

// Address provides a mock function with given fields:
func (_m *DeviceRemoteInterface) Address() *model.AddressDeviceType {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Address")
	}

	var r0 *model.AddressDeviceType
	if rf, ok := ret.Get(0).(func() *model.AddressDeviceType); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.AddressDeviceType)
		}
	}

	return r0
}

// DeviceRemoteInterface_Address_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Address'
type DeviceRemoteInterface_Address_Call struct {
	*mock.Call
}

// Address is a helper method to define mock.On call
func (_e *DeviceRemoteInterface_Expecter) Address() *DeviceRemoteInterface_Address_Call {
	return &DeviceRemoteInterface_Address_Call{Call: _e.mock.On("Address")}
}

func (_c *DeviceRemoteInterface_Address_Call) Run(run func()) *DeviceRemoteInterface_Address_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *DeviceRemoteInterface_Address_Call) Return(_a0 *model.AddressDeviceType) *DeviceRemoteInterface_Address_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *DeviceRemoteInterface_Address_Call) RunAndReturn(run func() *model.AddressDeviceType) *DeviceRemoteInterface_Address_Call {
	_c.Call.Return(run)
	return _c
}

// CheckEntityInformation provides a mock function with given fields: initialData, entity
func (_m *DeviceRemoteInterface) CheckEntityInformation(initialData bool, entity model.NodeManagementDetailedDiscoveryEntityInformationType) error {
	ret := _m.Called(initialData, entity)

	if len(ret) == 0 {
		panic("no return value specified for CheckEntityInformation")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(bool, model.NodeManagementDetailedDiscoveryEntityInformationType) error); ok {
		r0 = rf(initialData, entity)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeviceRemoteInterface_CheckEntityInformation_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CheckEntityInformation'
type DeviceRemoteInterface_CheckEntityInformation_Call struct {
	*mock.Call
}

// CheckEntityInformation is a helper method to define mock.On call
//   - initialData bool
//   - entity model.NodeManagementDetailedDiscoveryEntityInformationType
func (_e *DeviceRemoteInterface_Expecter) CheckEntityInformation(initialData interface{}, entity interface{}) *DeviceRemoteInterface_CheckEntityInformation_Call {
	return &DeviceRemoteInterface_CheckEntityInformation_Call{Call: _e.mock.On("CheckEntityInformation", initialData, entity)}
}

func (_c *DeviceRemoteInterface_CheckEntityInformation_Call) Run(run func(initialData bool, entity model.NodeManagementDetailedDiscoveryEntityInformationType)) *DeviceRemoteInterface_CheckEntityInformation_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(bool), args[1].(model.NodeManagementDetailedDiscoveryEntityInformationType))
	})
	return _c
}

func (_c *DeviceRemoteInterface_CheckEntityInformation_Call) Return(_a0 error) *DeviceRemoteInterface_CheckEntityInformation_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *DeviceRemoteInterface_CheckEntityInformation_Call) RunAndReturn(run func(bool, model.NodeManagementDetailedDiscoveryEntityInformationType) error) *DeviceRemoteInterface_CheckEntityInformation_Call {
	_c.Call.Return(run)
	return _c
}

// DestinationData provides a mock function with given fields:
func (_m *DeviceRemoteInterface) DestinationData() model.NodeManagementDestinationDataType {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for DestinationData")
	}

	var r0 model.NodeManagementDestinationDataType
	if rf, ok := ret.Get(0).(func() model.NodeManagementDestinationDataType); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(model.NodeManagementDestinationDataType)
	}

	return r0
}

// DeviceRemoteInterface_DestinationData_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DestinationData'
type DeviceRemoteInterface_DestinationData_Call struct {
	*mock.Call
}

// DestinationData is a helper method to define mock.On call
func (_e *DeviceRemoteInterface_Expecter) DestinationData() *DeviceRemoteInterface_DestinationData_Call {
	return &DeviceRemoteInterface_DestinationData_Call{Call: _e.mock.On("DestinationData")}
}

func (_c *DeviceRemoteInterface_DestinationData_Call) Run(run func()) *DeviceRemoteInterface_DestinationData_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *DeviceRemoteInterface_DestinationData_Call) Return(_a0 model.NodeManagementDestinationDataType) *DeviceRemoteInterface_DestinationData_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *DeviceRemoteInterface_DestinationData_Call) RunAndReturn(run func() model.NodeManagementDestinationDataType) *DeviceRemoteInterface_DestinationData_Call {
	_c.Call.Return(run)
	return _c
}

// DeviceType provides a mock function with given fields:
func (_m *DeviceRemoteInterface) DeviceType() *model.DeviceTypeType {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for DeviceType")
	}

	var r0 *model.DeviceTypeType
	if rf, ok := ret.Get(0).(func() *model.DeviceTypeType); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.DeviceTypeType)
		}
	}

	return r0
}

// DeviceRemoteInterface_DeviceType_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeviceType'
type DeviceRemoteInterface_DeviceType_Call struct {
	*mock.Call
}

// DeviceType is a helper method to define mock.On call
func (_e *DeviceRemoteInterface_Expecter) DeviceType() *DeviceRemoteInterface_DeviceType_Call {
	return &DeviceRemoteInterface_DeviceType_Call{Call: _e.mock.On("DeviceType")}
}

func (_c *DeviceRemoteInterface_DeviceType_Call) Run(run func()) *DeviceRemoteInterface_DeviceType_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *DeviceRemoteInterface_DeviceType_Call) Return(_a0 *model.DeviceTypeType) *DeviceRemoteInterface_DeviceType_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *DeviceRemoteInterface_DeviceType_Call) RunAndReturn(run func() *model.DeviceTypeType) *DeviceRemoteInterface_DeviceType_Call {
	_c.Call.Return(run)
	return _c
}

// Entities provides a mock function with given fields:
func (_m *DeviceRemoteInterface) Entities() []api.EntityRemoteInterface {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Entities")
	}

	var r0 []api.EntityRemoteInterface
	if rf, ok := ret.Get(0).(func() []api.EntityRemoteInterface); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]api.EntityRemoteInterface)
		}
	}

	return r0
}

// DeviceRemoteInterface_Entities_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Entities'
type DeviceRemoteInterface_Entities_Call struct {
	*mock.Call
}

// Entities is a helper method to define mock.On call
func (_e *DeviceRemoteInterface_Expecter) Entities() *DeviceRemoteInterface_Entities_Call {
	return &DeviceRemoteInterface_Entities_Call{Call: _e.mock.On("Entities")}
}

func (_c *DeviceRemoteInterface_Entities_Call) Run(run func()) *DeviceRemoteInterface_Entities_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *DeviceRemoteInterface_Entities_Call) Return(_a0 []api.EntityRemoteInterface) *DeviceRemoteInterface_Entities_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *DeviceRemoteInterface_Entities_Call) RunAndReturn(run func() []api.EntityRemoteInterface) *DeviceRemoteInterface_Entities_Call {
	_c.Call.Return(run)
	return _c
}

// Entity provides a mock function with given fields: id
func (_m *DeviceRemoteInterface) Entity(id []model.AddressEntityType) api.EntityRemoteInterface {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for Entity")
	}

	var r0 api.EntityRemoteInterface
	if rf, ok := ret.Get(0).(func([]model.AddressEntityType) api.EntityRemoteInterface); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(api.EntityRemoteInterface)
		}
	}

	return r0
}

// DeviceRemoteInterface_Entity_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Entity'
type DeviceRemoteInterface_Entity_Call struct {
	*mock.Call
}

// Entity is a helper method to define mock.On call
//   - id []model.AddressEntityType
func (_e *DeviceRemoteInterface_Expecter) Entity(id interface{}) *DeviceRemoteInterface_Entity_Call {
	return &DeviceRemoteInterface_Entity_Call{Call: _e.mock.On("Entity", id)}
}

func (_c *DeviceRemoteInterface_Entity_Call) Run(run func(id []model.AddressEntityType)) *DeviceRemoteInterface_Entity_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]model.AddressEntityType))
	})
	return _c
}

func (_c *DeviceRemoteInterface_Entity_Call) Return(_a0 api.EntityRemoteInterface) *DeviceRemoteInterface_Entity_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *DeviceRemoteInterface_Entity_Call) RunAndReturn(run func([]model.AddressEntityType) api.EntityRemoteInterface) *DeviceRemoteInterface_Entity_Call {
	_c.Call.Return(run)
	return _c
}

// FeatureByAddress provides a mock function with given fields: address
func (_m *DeviceRemoteInterface) FeatureByAddress(address *model.FeatureAddressType) api.FeatureRemoteInterface {
	ret := _m.Called(address)

	if len(ret) == 0 {
		panic("no return value specified for FeatureByAddress")
	}

	var r0 api.FeatureRemoteInterface
	if rf, ok := ret.Get(0).(func(*model.FeatureAddressType) api.FeatureRemoteInterface); ok {
		r0 = rf(address)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(api.FeatureRemoteInterface)
		}
	}

	return r0
}

// DeviceRemoteInterface_FeatureByAddress_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FeatureByAddress'
type DeviceRemoteInterface_FeatureByAddress_Call struct {
	*mock.Call
}

// FeatureByAddress is a helper method to define mock.On call
//   - address *model.FeatureAddressType
func (_e *DeviceRemoteInterface_Expecter) FeatureByAddress(address interface{}) *DeviceRemoteInterface_FeatureByAddress_Call {
	return &DeviceRemoteInterface_FeatureByAddress_Call{Call: _e.mock.On("FeatureByAddress", address)}
}

func (_c *DeviceRemoteInterface_FeatureByAddress_Call) Run(run func(address *model.FeatureAddressType)) *DeviceRemoteInterface_FeatureByAddress_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*model.FeatureAddressType))
	})
	return _c
}

func (_c *DeviceRemoteInterface_FeatureByAddress_Call) Return(_a0 api.FeatureRemoteInterface) *DeviceRemoteInterface_FeatureByAddress_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *DeviceRemoteInterface_FeatureByAddress_Call) RunAndReturn(run func(*model.FeatureAddressType) api.FeatureRemoteInterface) *DeviceRemoteInterface_FeatureByAddress_Call {
	_c.Call.Return(run)
	return _c
}

// FeatureByEntityTypeAndRole provides a mock function with given fields: entity, featureType, role
func (_m *DeviceRemoteInterface) FeatureByEntityTypeAndRole(entity api.EntityRemoteInterface, featureType model.FeatureTypeType, role model.RoleType) api.FeatureRemoteInterface {
	ret := _m.Called(entity, featureType, role)

	if len(ret) == 0 {
		panic("no return value specified for FeatureByEntityTypeAndRole")
	}

	var r0 api.FeatureRemoteInterface
	if rf, ok := ret.Get(0).(func(api.EntityRemoteInterface, model.FeatureTypeType, model.RoleType) api.FeatureRemoteInterface); ok {
		r0 = rf(entity, featureType, role)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(api.FeatureRemoteInterface)
		}
	}

	return r0
}

// DeviceRemoteInterface_FeatureByEntityTypeAndRole_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FeatureByEntityTypeAndRole'
type DeviceRemoteInterface_FeatureByEntityTypeAndRole_Call struct {
	*mock.Call
}

// FeatureByEntityTypeAndRole is a helper method to define mock.On call
//   - entity api.EntityRemoteInterface
//   - featureType model.FeatureTypeType
//   - role model.RoleType
func (_e *DeviceRemoteInterface_Expecter) FeatureByEntityTypeAndRole(entity interface{}, featureType interface{}, role interface{}) *DeviceRemoteInterface_FeatureByEntityTypeAndRole_Call {
	return &DeviceRemoteInterface_FeatureByEntityTypeAndRole_Call{Call: _e.mock.On("FeatureByEntityTypeAndRole", entity, featureType, role)}
}

func (_c *DeviceRemoteInterface_FeatureByEntityTypeAndRole_Call) Run(run func(entity api.EntityRemoteInterface, featureType model.FeatureTypeType, role model.RoleType)) *DeviceRemoteInterface_FeatureByEntityTypeAndRole_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(api.EntityRemoteInterface), args[1].(model.FeatureTypeType), args[2].(model.RoleType))
	})
	return _c
}

func (_c *DeviceRemoteInterface_FeatureByEntityTypeAndRole_Call) Return(_a0 api.FeatureRemoteInterface) *DeviceRemoteInterface_FeatureByEntityTypeAndRole_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *DeviceRemoteInterface_FeatureByEntityTypeAndRole_Call) RunAndReturn(run func(api.EntityRemoteInterface, model.FeatureTypeType, model.RoleType) api.FeatureRemoteInterface) *DeviceRemoteInterface_FeatureByEntityTypeAndRole_Call {
	_c.Call.Return(run)
	return _c
}

// FeatureSet provides a mock function with given fields:
func (_m *DeviceRemoteInterface) FeatureSet() *model.NetworkManagementFeatureSetType {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for FeatureSet")
	}

	var r0 *model.NetworkManagementFeatureSetType
	if rf, ok := ret.Get(0).(func() *model.NetworkManagementFeatureSetType); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.NetworkManagementFeatureSetType)
		}
	}

	return r0
}

// DeviceRemoteInterface_FeatureSet_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FeatureSet'
type DeviceRemoteInterface_FeatureSet_Call struct {
	*mock.Call
}

// FeatureSet is a helper method to define mock.On call
func (_e *DeviceRemoteInterface_Expecter) FeatureSet() *DeviceRemoteInterface_FeatureSet_Call {
	return &DeviceRemoteInterface_FeatureSet_Call{Call: _e.mock.On("FeatureSet")}
}

func (_c *DeviceRemoteInterface_FeatureSet_Call) Run(run func()) *DeviceRemoteInterface_FeatureSet_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *DeviceRemoteInterface_FeatureSet_Call) Return(_a0 *model.NetworkManagementFeatureSetType) *DeviceRemoteInterface_FeatureSet_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *DeviceRemoteInterface_FeatureSet_Call) RunAndReturn(run func() *model.NetworkManagementFeatureSetType) *DeviceRemoteInterface_FeatureSet_Call {
	_c.Call.Return(run)
	return _c
}

// HandleSpineMesssage provides a mock function with given fields: message
func (_m *DeviceRemoteInterface) HandleSpineMesssage(message []byte) (*model.MsgCounterType, error) {
	ret := _m.Called(message)

	if len(ret) == 0 {
		panic("no return value specified for HandleSpineMesssage")
	}

	var r0 *model.MsgCounterType
	var r1 error
	if rf, ok := ret.Get(0).(func([]byte) (*model.MsgCounterType, error)); ok {
		return rf(message)
	}
	if rf, ok := ret.Get(0).(func([]byte) *model.MsgCounterType); ok {
		r0 = rf(message)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.MsgCounterType)
		}
	}

	if rf, ok := ret.Get(1).(func([]byte) error); ok {
		r1 = rf(message)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeviceRemoteInterface_HandleSpineMesssage_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'HandleSpineMesssage'
type DeviceRemoteInterface_HandleSpineMesssage_Call struct {
	*mock.Call
}

// HandleSpineMesssage is a helper method to define mock.On call
//   - message []byte
func (_e *DeviceRemoteInterface_Expecter) HandleSpineMesssage(message interface{}) *DeviceRemoteInterface_HandleSpineMesssage_Call {
	return &DeviceRemoteInterface_HandleSpineMesssage_Call{Call: _e.mock.On("HandleSpineMesssage", message)}
}

func (_c *DeviceRemoteInterface_HandleSpineMesssage_Call) Run(run func(message []byte)) *DeviceRemoteInterface_HandleSpineMesssage_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]byte))
	})
	return _c
}

func (_c *DeviceRemoteInterface_HandleSpineMesssage_Call) Return(_a0 *model.MsgCounterType, _a1 error) *DeviceRemoteInterface_HandleSpineMesssage_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *DeviceRemoteInterface_HandleSpineMesssage_Call) RunAndReturn(run func([]byte) (*model.MsgCounterType, error)) *DeviceRemoteInterface_HandleSpineMesssage_Call {
	_c.Call.Return(run)
	return _c
}

// RemoveEntityByAddress provides a mock function with given fields: addr
func (_m *DeviceRemoteInterface) RemoveEntityByAddress(addr []model.AddressEntityType) api.EntityRemoteInterface {
	ret := _m.Called(addr)

	if len(ret) == 0 {
		panic("no return value specified for RemoveEntityByAddress")
	}

	var r0 api.EntityRemoteInterface
	if rf, ok := ret.Get(0).(func([]model.AddressEntityType) api.EntityRemoteInterface); ok {
		r0 = rf(addr)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(api.EntityRemoteInterface)
		}
	}

	return r0
}

// DeviceRemoteInterface_RemoveEntityByAddress_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RemoveEntityByAddress'
type DeviceRemoteInterface_RemoveEntityByAddress_Call struct {
	*mock.Call
}

// RemoveEntityByAddress is a helper method to define mock.On call
//   - addr []model.AddressEntityType
func (_e *DeviceRemoteInterface_Expecter) RemoveEntityByAddress(addr interface{}) *DeviceRemoteInterface_RemoveEntityByAddress_Call {
	return &DeviceRemoteInterface_RemoveEntityByAddress_Call{Call: _e.mock.On("RemoveEntityByAddress", addr)}
}

func (_c *DeviceRemoteInterface_RemoveEntityByAddress_Call) Run(run func(addr []model.AddressEntityType)) *DeviceRemoteInterface_RemoveEntityByAddress_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]model.AddressEntityType))
	})
	return _c
}

func (_c *DeviceRemoteInterface_RemoveEntityByAddress_Call) Return(_a0 api.EntityRemoteInterface) *DeviceRemoteInterface_RemoveEntityByAddress_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *DeviceRemoteInterface_RemoveEntityByAddress_Call) RunAndReturn(run func([]model.AddressEntityType) api.EntityRemoteInterface) *DeviceRemoteInterface_RemoveEntityByAddress_Call {
	_c.Call.Return(run)
	return _c
}

// Sender provides a mock function with given fields:
func (_m *DeviceRemoteInterface) Sender() api.SenderInterface {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Sender")
	}

	var r0 api.SenderInterface
	if rf, ok := ret.Get(0).(func() api.SenderInterface); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(api.SenderInterface)
		}
	}

	return r0
}

// DeviceRemoteInterface_Sender_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Sender'
type DeviceRemoteInterface_Sender_Call struct {
	*mock.Call
}

// Sender is a helper method to define mock.On call
func (_e *DeviceRemoteInterface_Expecter) Sender() *DeviceRemoteInterface_Sender_Call {
	return &DeviceRemoteInterface_Sender_Call{Call: _e.mock.On("Sender")}
}

func (_c *DeviceRemoteInterface_Sender_Call) Run(run func()) *DeviceRemoteInterface_Sender_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *DeviceRemoteInterface_Sender_Call) Return(_a0 api.SenderInterface) *DeviceRemoteInterface_Sender_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *DeviceRemoteInterface_Sender_Call) RunAndReturn(run func() api.SenderInterface) *DeviceRemoteInterface_Sender_Call {
	_c.Call.Return(run)
	return _c
}

// Ski provides a mock function with given fields:
func (_m *DeviceRemoteInterface) Ski() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Ski")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// DeviceRemoteInterface_Ski_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Ski'
type DeviceRemoteInterface_Ski_Call struct {
	*mock.Call
}

// Ski is a helper method to define mock.On call
func (_e *DeviceRemoteInterface_Expecter) Ski() *DeviceRemoteInterface_Ski_Call {
	return &DeviceRemoteInterface_Ski_Call{Call: _e.mock.On("Ski")}
}

func (_c *DeviceRemoteInterface_Ski_Call) Run(run func()) *DeviceRemoteInterface_Ski_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *DeviceRemoteInterface_Ski_Call) Return(_a0 string) *DeviceRemoteInterface_Ski_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *DeviceRemoteInterface_Ski_Call) RunAndReturn(run func() string) *DeviceRemoteInterface_Ski_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateDevice provides a mock function with given fields: description
func (_m *DeviceRemoteInterface) UpdateDevice(description *model.NetworkManagementDeviceDescriptionDataType) {
	_m.Called(description)
}

// DeviceRemoteInterface_UpdateDevice_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateDevice'
type DeviceRemoteInterface_UpdateDevice_Call struct {
	*mock.Call
}

// UpdateDevice is a helper method to define mock.On call
//   - description *model.NetworkManagementDeviceDescriptionDataType
func (_e *DeviceRemoteInterface_Expecter) UpdateDevice(description interface{}) *DeviceRemoteInterface_UpdateDevice_Call {
	return &DeviceRemoteInterface_UpdateDevice_Call{Call: _e.mock.On("UpdateDevice", description)}
}

func (_c *DeviceRemoteInterface_UpdateDevice_Call) Run(run func(description *model.NetworkManagementDeviceDescriptionDataType)) *DeviceRemoteInterface_UpdateDevice_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*model.NetworkManagementDeviceDescriptionDataType))
	})
	return _c
}

func (_c *DeviceRemoteInterface_UpdateDevice_Call) Return() *DeviceRemoteInterface_UpdateDevice_Call {
	_c.Call.Return()
	return _c
}

func (_c *DeviceRemoteInterface_UpdateDevice_Call) RunAndReturn(run func(*model.NetworkManagementDeviceDescriptionDataType)) *DeviceRemoteInterface_UpdateDevice_Call {
	_c.Call.Return(run)
	return _c
}

// UseCases provides a mock function with given fields:
func (_m *DeviceRemoteInterface) UseCases() []model.UseCaseInformationDataType {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for UseCases")
	}

	var r0 []model.UseCaseInformationDataType
	if rf, ok := ret.Get(0).(func() []model.UseCaseInformationDataType); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.UseCaseInformationDataType)
		}
	}

	return r0
}

// DeviceRemoteInterface_UseCases_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UseCases'
type DeviceRemoteInterface_UseCases_Call struct {
	*mock.Call
}

// UseCases is a helper method to define mock.On call
func (_e *DeviceRemoteInterface_Expecter) UseCases() *DeviceRemoteInterface_UseCases_Call {
	return &DeviceRemoteInterface_UseCases_Call{Call: _e.mock.On("UseCases")}
}

func (_c *DeviceRemoteInterface_UseCases_Call) Run(run func()) *DeviceRemoteInterface_UseCases_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *DeviceRemoteInterface_UseCases_Call) Return(_a0 []model.UseCaseInformationDataType) *DeviceRemoteInterface_UseCases_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *DeviceRemoteInterface_UseCases_Call) RunAndReturn(run func() []model.UseCaseInformationDataType) *DeviceRemoteInterface_UseCases_Call {
	_c.Call.Return(run)
	return _c
}

// NewDeviceRemoteInterface creates a new instance of DeviceRemoteInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewDeviceRemoteInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *DeviceRemoteInterface {
	mock := &DeviceRemoteInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
