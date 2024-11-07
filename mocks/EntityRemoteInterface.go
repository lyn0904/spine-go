// Code generated by mockery v2.45.0. DO NOT EDIT.

package mocks

import (
	api "github.com/lyn0904/spine-go/api"
	mock "github.com/stretchr/testify/mock"

	model "github.com/lyn0904/spine-go/model"
)

// EntityRemoteInterface is an autogenerated mock type for the EntityRemoteInterface type
type EntityRemoteInterface struct {
	mock.Mock
}

type EntityRemoteInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *EntityRemoteInterface) EXPECT() *EntityRemoteInterface_Expecter {
	return &EntityRemoteInterface_Expecter{mock: &_m.Mock}
}

// AddFeature provides a mock function with given fields: f
func (_m *EntityRemoteInterface) AddFeature(f api.FeatureRemoteInterface) {
	_m.Called(f)
}

// EntityRemoteInterface_AddFeature_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddFeature'
type EntityRemoteInterface_AddFeature_Call struct {
	*mock.Call
}

// AddFeature is a helper method to define mock.On call
//   - f api.FeatureRemoteInterface
func (_e *EntityRemoteInterface_Expecter) AddFeature(f interface{}) *EntityRemoteInterface_AddFeature_Call {
	return &EntityRemoteInterface_AddFeature_Call{Call: _e.mock.On("AddFeature", f)}
}

func (_c *EntityRemoteInterface_AddFeature_Call) Run(run func(f api.FeatureRemoteInterface)) *EntityRemoteInterface_AddFeature_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(api.FeatureRemoteInterface))
	})
	return _c
}

func (_c *EntityRemoteInterface_AddFeature_Call) Return() *EntityRemoteInterface_AddFeature_Call {
	_c.Call.Return()
	return _c
}

func (_c *EntityRemoteInterface_AddFeature_Call) RunAndReturn(run func(api.FeatureRemoteInterface)) *EntityRemoteInterface_AddFeature_Call {
	_c.Call.Return(run)
	return _c
}

// Address provides a mock function with given fields:
func (_m *EntityRemoteInterface) Address() *model.EntityAddressType {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Address")
	}

	var r0 *model.EntityAddressType
	if rf, ok := ret.Get(0).(func() *model.EntityAddressType); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.EntityAddressType)
		}
	}

	return r0
}

// EntityRemoteInterface_Address_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Address'
type EntityRemoteInterface_Address_Call struct {
	*mock.Call
}

// Address is a helper method to define mock.On call
func (_e *EntityRemoteInterface_Expecter) Address() *EntityRemoteInterface_Address_Call {
	return &EntityRemoteInterface_Address_Call{Call: _e.mock.On("Address")}
}

func (_c *EntityRemoteInterface_Address_Call) Run(run func()) *EntityRemoteInterface_Address_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *EntityRemoteInterface_Address_Call) Return(_a0 *model.EntityAddressType) *EntityRemoteInterface_Address_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *EntityRemoteInterface_Address_Call) RunAndReturn(run func() *model.EntityAddressType) *EntityRemoteInterface_Address_Call {
	_c.Call.Return(run)
	return _c
}

// Description provides a mock function with given fields:
func (_m *EntityRemoteInterface) Description() *model.DescriptionType {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Description")
	}

	var r0 *model.DescriptionType
	if rf, ok := ret.Get(0).(func() *model.DescriptionType); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.DescriptionType)
		}
	}

	return r0
}

// EntityRemoteInterface_Description_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Description'
type EntityRemoteInterface_Description_Call struct {
	*mock.Call
}

// Description is a helper method to define mock.On call
func (_e *EntityRemoteInterface_Expecter) Description() *EntityRemoteInterface_Description_Call {
	return &EntityRemoteInterface_Description_Call{Call: _e.mock.On("Description")}
}

func (_c *EntityRemoteInterface_Description_Call) Run(run func()) *EntityRemoteInterface_Description_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *EntityRemoteInterface_Description_Call) Return(_a0 *model.DescriptionType) *EntityRemoteInterface_Description_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *EntityRemoteInterface_Description_Call) RunAndReturn(run func() *model.DescriptionType) *EntityRemoteInterface_Description_Call {
	_c.Call.Return(run)
	return _c
}

// Device provides a mock function with given fields:
func (_m *EntityRemoteInterface) Device() api.DeviceRemoteInterface {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Device")
	}

	var r0 api.DeviceRemoteInterface
	if rf, ok := ret.Get(0).(func() api.DeviceRemoteInterface); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(api.DeviceRemoteInterface)
		}
	}

	return r0
}

// EntityRemoteInterface_Device_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Device'
type EntityRemoteInterface_Device_Call struct {
	*mock.Call
}

// Device is a helper method to define mock.On call
func (_e *EntityRemoteInterface_Expecter) Device() *EntityRemoteInterface_Device_Call {
	return &EntityRemoteInterface_Device_Call{Call: _e.mock.On("Device")}
}

func (_c *EntityRemoteInterface_Device_Call) Run(run func()) *EntityRemoteInterface_Device_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *EntityRemoteInterface_Device_Call) Return(_a0 api.DeviceRemoteInterface) *EntityRemoteInterface_Device_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *EntityRemoteInterface_Device_Call) RunAndReturn(run func() api.DeviceRemoteInterface) *EntityRemoteInterface_Device_Call {
	_c.Call.Return(run)
	return _c
}

// EntityType provides a mock function with given fields:
func (_m *EntityRemoteInterface) EntityType() model.EntityTypeType {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for EntityType")
	}

	var r0 model.EntityTypeType
	if rf, ok := ret.Get(0).(func() model.EntityTypeType); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(model.EntityTypeType)
	}

	return r0
}

// EntityRemoteInterface_EntityType_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'EntityType'
type EntityRemoteInterface_EntityType_Call struct {
	*mock.Call
}

// EntityType is a helper method to define mock.On call
func (_e *EntityRemoteInterface_Expecter) EntityType() *EntityRemoteInterface_EntityType_Call {
	return &EntityRemoteInterface_EntityType_Call{Call: _e.mock.On("EntityType")}
}

func (_c *EntityRemoteInterface_EntityType_Call) Run(run func()) *EntityRemoteInterface_EntityType_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *EntityRemoteInterface_EntityType_Call) Return(_a0 model.EntityTypeType) *EntityRemoteInterface_EntityType_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *EntityRemoteInterface_EntityType_Call) RunAndReturn(run func() model.EntityTypeType) *EntityRemoteInterface_EntityType_Call {
	_c.Call.Return(run)
	return _c
}

// FeatureOfAddress provides a mock function with given fields: addressFeature
func (_m *EntityRemoteInterface) FeatureOfAddress(addressFeature *model.AddressFeatureType) api.FeatureRemoteInterface {
	ret := _m.Called(addressFeature)

	if len(ret) == 0 {
		panic("no return value specified for FeatureOfAddress")
	}

	var r0 api.FeatureRemoteInterface
	if rf, ok := ret.Get(0).(func(*model.AddressFeatureType) api.FeatureRemoteInterface); ok {
		r0 = rf(addressFeature)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(api.FeatureRemoteInterface)
		}
	}

	return r0
}

// EntityRemoteInterface_FeatureOfAddress_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FeatureOfAddress'
type EntityRemoteInterface_FeatureOfAddress_Call struct {
	*mock.Call
}

// FeatureOfAddress is a helper method to define mock.On call
//   - addressFeature *model.AddressFeatureType
func (_e *EntityRemoteInterface_Expecter) FeatureOfAddress(addressFeature interface{}) *EntityRemoteInterface_FeatureOfAddress_Call {
	return &EntityRemoteInterface_FeatureOfAddress_Call{Call: _e.mock.On("FeatureOfAddress", addressFeature)}
}

func (_c *EntityRemoteInterface_FeatureOfAddress_Call) Run(run func(addressFeature *model.AddressFeatureType)) *EntityRemoteInterface_FeatureOfAddress_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*model.AddressFeatureType))
	})
	return _c
}

func (_c *EntityRemoteInterface_FeatureOfAddress_Call) Return(_a0 api.FeatureRemoteInterface) *EntityRemoteInterface_FeatureOfAddress_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *EntityRemoteInterface_FeatureOfAddress_Call) RunAndReturn(run func(*model.AddressFeatureType) api.FeatureRemoteInterface) *EntityRemoteInterface_FeatureOfAddress_Call {
	_c.Call.Return(run)
	return _c
}

// FeatureOfTypeAndRole provides a mock function with given fields: featureType, role
func (_m *EntityRemoteInterface) FeatureOfTypeAndRole(featureType model.FeatureTypeType, role model.RoleType) api.FeatureRemoteInterface {
	ret := _m.Called(featureType, role)

	if len(ret) == 0 {
		panic("no return value specified for FeatureOfTypeAndRole")
	}

	var r0 api.FeatureRemoteInterface
	if rf, ok := ret.Get(0).(func(model.FeatureTypeType, model.RoleType) api.FeatureRemoteInterface); ok {
		r0 = rf(featureType, role)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(api.FeatureRemoteInterface)
		}
	}

	return r0
}

// EntityRemoteInterface_FeatureOfTypeAndRole_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FeatureOfTypeAndRole'
type EntityRemoteInterface_FeatureOfTypeAndRole_Call struct {
	*mock.Call
}

// FeatureOfTypeAndRole is a helper method to define mock.On call
//   - featureType model.FeatureTypeType
//   - role model.RoleType
func (_e *EntityRemoteInterface_Expecter) FeatureOfTypeAndRole(featureType interface{}, role interface{}) *EntityRemoteInterface_FeatureOfTypeAndRole_Call {
	return &EntityRemoteInterface_FeatureOfTypeAndRole_Call{Call: _e.mock.On("FeatureOfTypeAndRole", featureType, role)}
}

func (_c *EntityRemoteInterface_FeatureOfTypeAndRole_Call) Run(run func(featureType model.FeatureTypeType, role model.RoleType)) *EntityRemoteInterface_FeatureOfTypeAndRole_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(model.FeatureTypeType), args[1].(model.RoleType))
	})
	return _c
}

func (_c *EntityRemoteInterface_FeatureOfTypeAndRole_Call) Return(_a0 api.FeatureRemoteInterface) *EntityRemoteInterface_FeatureOfTypeAndRole_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *EntityRemoteInterface_FeatureOfTypeAndRole_Call) RunAndReturn(run func(model.FeatureTypeType, model.RoleType) api.FeatureRemoteInterface) *EntityRemoteInterface_FeatureOfTypeAndRole_Call {
	_c.Call.Return(run)
	return _c
}

// Features provides a mock function with given fields:
func (_m *EntityRemoteInterface) Features() []api.FeatureRemoteInterface {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Features")
	}

	var r0 []api.FeatureRemoteInterface
	if rf, ok := ret.Get(0).(func() []api.FeatureRemoteInterface); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]api.FeatureRemoteInterface)
		}
	}

	return r0
}

// EntityRemoteInterface_Features_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Features'
type EntityRemoteInterface_Features_Call struct {
	*mock.Call
}

// Features is a helper method to define mock.On call
func (_e *EntityRemoteInterface_Expecter) Features() *EntityRemoteInterface_Features_Call {
	return &EntityRemoteInterface_Features_Call{Call: _e.mock.On("Features")}
}

func (_c *EntityRemoteInterface_Features_Call) Run(run func()) *EntityRemoteInterface_Features_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *EntityRemoteInterface_Features_Call) Return(_a0 []api.FeatureRemoteInterface) *EntityRemoteInterface_Features_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *EntityRemoteInterface_Features_Call) RunAndReturn(run func() []api.FeatureRemoteInterface) *EntityRemoteInterface_Features_Call {
	_c.Call.Return(run)
	return _c
}

// NextFeatureId provides a mock function with given fields:
func (_m *EntityRemoteInterface) NextFeatureId() uint {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for NextFeatureId")
	}

	var r0 uint
	if rf, ok := ret.Get(0).(func() uint); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint)
	}

	return r0
}

// EntityRemoteInterface_NextFeatureId_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'NextFeatureId'
type EntityRemoteInterface_NextFeatureId_Call struct {
	*mock.Call
}

// NextFeatureId is a helper method to define mock.On call
func (_e *EntityRemoteInterface_Expecter) NextFeatureId() *EntityRemoteInterface_NextFeatureId_Call {
	return &EntityRemoteInterface_NextFeatureId_Call{Call: _e.mock.On("NextFeatureId")}
}

func (_c *EntityRemoteInterface_NextFeatureId_Call) Run(run func()) *EntityRemoteInterface_NextFeatureId_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *EntityRemoteInterface_NextFeatureId_Call) Return(_a0 uint) *EntityRemoteInterface_NextFeatureId_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *EntityRemoteInterface_NextFeatureId_Call) RunAndReturn(run func() uint) *EntityRemoteInterface_NextFeatureId_Call {
	_c.Call.Return(run)
	return _c
}

// RemoveAllFeatures provides a mock function with given fields:
func (_m *EntityRemoteInterface) RemoveAllFeatures() {
	_m.Called()
}

// EntityRemoteInterface_RemoveAllFeatures_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RemoveAllFeatures'
type EntityRemoteInterface_RemoveAllFeatures_Call struct {
	*mock.Call
}

// RemoveAllFeatures is a helper method to define mock.On call
func (_e *EntityRemoteInterface_Expecter) RemoveAllFeatures() *EntityRemoteInterface_RemoveAllFeatures_Call {
	return &EntityRemoteInterface_RemoveAllFeatures_Call{Call: _e.mock.On("RemoveAllFeatures")}
}

func (_c *EntityRemoteInterface_RemoveAllFeatures_Call) Run(run func()) *EntityRemoteInterface_RemoveAllFeatures_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *EntityRemoteInterface_RemoveAllFeatures_Call) Return() *EntityRemoteInterface_RemoveAllFeatures_Call {
	_c.Call.Return()
	return _c
}

func (_c *EntityRemoteInterface_RemoveAllFeatures_Call) RunAndReturn(run func()) *EntityRemoteInterface_RemoveAllFeatures_Call {
	_c.Call.Return(run)
	return _c
}

// SetDescription provides a mock function with given fields: d
func (_m *EntityRemoteInterface) SetDescription(d *model.DescriptionType) {
	_m.Called(d)
}

// EntityRemoteInterface_SetDescription_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetDescription'
type EntityRemoteInterface_SetDescription_Call struct {
	*mock.Call
}

// SetDescription is a helper method to define mock.On call
//   - d *model.DescriptionType
func (_e *EntityRemoteInterface_Expecter) SetDescription(d interface{}) *EntityRemoteInterface_SetDescription_Call {
	return &EntityRemoteInterface_SetDescription_Call{Call: _e.mock.On("SetDescription", d)}
}

func (_c *EntityRemoteInterface_SetDescription_Call) Run(run func(d *model.DescriptionType)) *EntityRemoteInterface_SetDescription_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*model.DescriptionType))
	})
	return _c
}

func (_c *EntityRemoteInterface_SetDescription_Call) Return() *EntityRemoteInterface_SetDescription_Call {
	_c.Call.Return()
	return _c
}

func (_c *EntityRemoteInterface_SetDescription_Call) RunAndReturn(run func(*model.DescriptionType)) *EntityRemoteInterface_SetDescription_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateDeviceAddress provides a mock function with given fields: address
func (_m *EntityRemoteInterface) UpdateDeviceAddress(address model.AddressDeviceType) {
	_m.Called(address)
}

// EntityRemoteInterface_UpdateDeviceAddress_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateDeviceAddress'
type EntityRemoteInterface_UpdateDeviceAddress_Call struct {
	*mock.Call
}

// UpdateDeviceAddress is a helper method to define mock.On call
//   - address model.AddressDeviceType
func (_e *EntityRemoteInterface_Expecter) UpdateDeviceAddress(address interface{}) *EntityRemoteInterface_UpdateDeviceAddress_Call {
	return &EntityRemoteInterface_UpdateDeviceAddress_Call{Call: _e.mock.On("UpdateDeviceAddress", address)}
}

func (_c *EntityRemoteInterface_UpdateDeviceAddress_Call) Run(run func(address model.AddressDeviceType)) *EntityRemoteInterface_UpdateDeviceAddress_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(model.AddressDeviceType))
	})
	return _c
}

func (_c *EntityRemoteInterface_UpdateDeviceAddress_Call) Return() *EntityRemoteInterface_UpdateDeviceAddress_Call {
	_c.Call.Return()
	return _c
}

func (_c *EntityRemoteInterface_UpdateDeviceAddress_Call) RunAndReturn(run func(model.AddressDeviceType)) *EntityRemoteInterface_UpdateDeviceAddress_Call {
	_c.Call.Return(run)
	return _c
}

// NewEntityRemoteInterface creates a new instance of EntityRemoteInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewEntityRemoteInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *EntityRemoteInterface {
	mock := &EntityRemoteInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
