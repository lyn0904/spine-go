// Code generated by mockery v2.45.0. DO NOT EDIT.

package mocks

import (
	api "github.com/lyn0904/spine-go/api"
	mock "github.com/stretchr/testify/mock"

	model "github.com/lyn0904/spine-go/model"
)

// BindingManagerInterface is an autogenerated mock type for the BindingManagerInterface type
type BindingManagerInterface struct {
	mock.Mock
}

type BindingManagerInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *BindingManagerInterface) EXPECT() *BindingManagerInterface_Expecter {
	return &BindingManagerInterface_Expecter{mock: &_m.Mock}
}

// AddBinding provides a mock function with given fields: remoteDevice, data
func (_m *BindingManagerInterface) AddBinding(remoteDevice api.DeviceRemoteInterface, data model.BindingManagementRequestCallType) error {
	ret := _m.Called(remoteDevice, data)

	if len(ret) == 0 {
		panic("no return value specified for AddBinding")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(api.DeviceRemoteInterface, model.BindingManagementRequestCallType) error); ok {
		r0 = rf(remoteDevice, data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// BindingManagerInterface_AddBinding_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddBinding'
type BindingManagerInterface_AddBinding_Call struct {
	*mock.Call
}

// AddBinding is a helper method to define mock.On call
//   - remoteDevice api.DeviceRemoteInterface
//   - data model.BindingManagementRequestCallType
func (_e *BindingManagerInterface_Expecter) AddBinding(remoteDevice interface{}, data interface{}) *BindingManagerInterface_AddBinding_Call {
	return &BindingManagerInterface_AddBinding_Call{Call: _e.mock.On("AddBinding", remoteDevice, data)}
}

func (_c *BindingManagerInterface_AddBinding_Call) Run(run func(remoteDevice api.DeviceRemoteInterface, data model.BindingManagementRequestCallType)) *BindingManagerInterface_AddBinding_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(api.DeviceRemoteInterface), args[1].(model.BindingManagementRequestCallType))
	})
	return _c
}

func (_c *BindingManagerInterface_AddBinding_Call) Return(_a0 error) *BindingManagerInterface_AddBinding_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *BindingManagerInterface_AddBinding_Call) RunAndReturn(run func(api.DeviceRemoteInterface, model.BindingManagementRequestCallType) error) *BindingManagerInterface_AddBinding_Call {
	_c.Call.Return(run)
	return _c
}

// Bindings provides a mock function with given fields: remoteDevice
func (_m *BindingManagerInterface) Bindings(remoteDevice api.DeviceRemoteInterface) []*api.BindingEntry {
	ret := _m.Called(remoteDevice)

	if len(ret) == 0 {
		panic("no return value specified for Bindings")
	}

	var r0 []*api.BindingEntry
	if rf, ok := ret.Get(0).(func(api.DeviceRemoteInterface) []*api.BindingEntry); ok {
		r0 = rf(remoteDevice)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*api.BindingEntry)
		}
	}

	return r0
}

// BindingManagerInterface_Bindings_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Bindings'
type BindingManagerInterface_Bindings_Call struct {
	*mock.Call
}

// Bindings is a helper method to define mock.On call
//   - remoteDevice api.DeviceRemoteInterface
func (_e *BindingManagerInterface_Expecter) Bindings(remoteDevice interface{}) *BindingManagerInterface_Bindings_Call {
	return &BindingManagerInterface_Bindings_Call{Call: _e.mock.On("Bindings", remoteDevice)}
}

func (_c *BindingManagerInterface_Bindings_Call) Run(run func(remoteDevice api.DeviceRemoteInterface)) *BindingManagerInterface_Bindings_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(api.DeviceRemoteInterface))
	})
	return _c
}

func (_c *BindingManagerInterface_Bindings_Call) Return(_a0 []*api.BindingEntry) *BindingManagerInterface_Bindings_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *BindingManagerInterface_Bindings_Call) RunAndReturn(run func(api.DeviceRemoteInterface) []*api.BindingEntry) *BindingManagerInterface_Bindings_Call {
	_c.Call.Return(run)
	return _c
}

// BindingsOnFeature provides a mock function with given fields: featureAddress
func (_m *BindingManagerInterface) BindingsOnFeature(featureAddress model.FeatureAddressType) []*api.BindingEntry {
	ret := _m.Called(featureAddress)

	if len(ret) == 0 {
		panic("no return value specified for BindingsOnFeature")
	}

	var r0 []*api.BindingEntry
	if rf, ok := ret.Get(0).(func(model.FeatureAddressType) []*api.BindingEntry); ok {
		r0 = rf(featureAddress)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*api.BindingEntry)
		}
	}

	return r0
}

// BindingManagerInterface_BindingsOnFeature_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'BindingsOnFeature'
type BindingManagerInterface_BindingsOnFeature_Call struct {
	*mock.Call
}

// BindingsOnFeature is a helper method to define mock.On call
//   - featureAddress model.FeatureAddressType
func (_e *BindingManagerInterface_Expecter) BindingsOnFeature(featureAddress interface{}) *BindingManagerInterface_BindingsOnFeature_Call {
	return &BindingManagerInterface_BindingsOnFeature_Call{Call: _e.mock.On("BindingsOnFeature", featureAddress)}
}

func (_c *BindingManagerInterface_BindingsOnFeature_Call) Run(run func(featureAddress model.FeatureAddressType)) *BindingManagerInterface_BindingsOnFeature_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(model.FeatureAddressType))
	})
	return _c
}

func (_c *BindingManagerInterface_BindingsOnFeature_Call) Return(_a0 []*api.BindingEntry) *BindingManagerInterface_BindingsOnFeature_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *BindingManagerInterface_BindingsOnFeature_Call) RunAndReturn(run func(model.FeatureAddressType) []*api.BindingEntry) *BindingManagerInterface_BindingsOnFeature_Call {
	_c.Call.Return(run)
	return _c
}

// HasLocalFeatureRemoteBinding provides a mock function with given fields: localAddress, remoteAddress
func (_m *BindingManagerInterface) HasLocalFeatureRemoteBinding(localAddress *model.FeatureAddressType, remoteAddress *model.FeatureAddressType) bool {
	ret := _m.Called(localAddress, remoteAddress)

	if len(ret) == 0 {
		panic("no return value specified for HasLocalFeatureRemoteBinding")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func(*model.FeatureAddressType, *model.FeatureAddressType) bool); ok {
		r0 = rf(localAddress, remoteAddress)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// BindingManagerInterface_HasLocalFeatureRemoteBinding_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'HasLocalFeatureRemoteBinding'
type BindingManagerInterface_HasLocalFeatureRemoteBinding_Call struct {
	*mock.Call
}

// HasLocalFeatureRemoteBinding is a helper method to define mock.On call
//   - localAddress *model.FeatureAddressType
//   - remoteAddress *model.FeatureAddressType
func (_e *BindingManagerInterface_Expecter) HasLocalFeatureRemoteBinding(localAddress interface{}, remoteAddress interface{}) *BindingManagerInterface_HasLocalFeatureRemoteBinding_Call {
	return &BindingManagerInterface_HasLocalFeatureRemoteBinding_Call{Call: _e.mock.On("HasLocalFeatureRemoteBinding", localAddress, remoteAddress)}
}

func (_c *BindingManagerInterface_HasLocalFeatureRemoteBinding_Call) Run(run func(localAddress *model.FeatureAddressType, remoteAddress *model.FeatureAddressType)) *BindingManagerInterface_HasLocalFeatureRemoteBinding_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*model.FeatureAddressType), args[1].(*model.FeatureAddressType))
	})
	return _c
}

func (_c *BindingManagerInterface_HasLocalFeatureRemoteBinding_Call) Return(_a0 bool) *BindingManagerInterface_HasLocalFeatureRemoteBinding_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *BindingManagerInterface_HasLocalFeatureRemoteBinding_Call) RunAndReturn(run func(*model.FeatureAddressType, *model.FeatureAddressType) bool) *BindingManagerInterface_HasLocalFeatureRemoteBinding_Call {
	_c.Call.Return(run)
	return _c
}

// RemoveBinding provides a mock function with given fields: data, remoteDevice
func (_m *BindingManagerInterface) RemoveBinding(data model.BindingManagementDeleteCallType, remoteDevice api.DeviceRemoteInterface) error {
	ret := _m.Called(data, remoteDevice)

	if len(ret) == 0 {
		panic("no return value specified for RemoveBinding")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(model.BindingManagementDeleteCallType, api.DeviceRemoteInterface) error); ok {
		r0 = rf(data, remoteDevice)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// BindingManagerInterface_RemoveBinding_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RemoveBinding'
type BindingManagerInterface_RemoveBinding_Call struct {
	*mock.Call
}

// RemoveBinding is a helper method to define mock.On call
//   - data model.BindingManagementDeleteCallType
//   - remoteDevice api.DeviceRemoteInterface
func (_e *BindingManagerInterface_Expecter) RemoveBinding(data interface{}, remoteDevice interface{}) *BindingManagerInterface_RemoveBinding_Call {
	return &BindingManagerInterface_RemoveBinding_Call{Call: _e.mock.On("RemoveBinding", data, remoteDevice)}
}

func (_c *BindingManagerInterface_RemoveBinding_Call) Run(run func(data model.BindingManagementDeleteCallType, remoteDevice api.DeviceRemoteInterface)) *BindingManagerInterface_RemoveBinding_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(model.BindingManagementDeleteCallType), args[1].(api.DeviceRemoteInterface))
	})
	return _c
}

func (_c *BindingManagerInterface_RemoveBinding_Call) Return(_a0 error) *BindingManagerInterface_RemoveBinding_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *BindingManagerInterface_RemoveBinding_Call) RunAndReturn(run func(model.BindingManagementDeleteCallType, api.DeviceRemoteInterface) error) *BindingManagerInterface_RemoveBinding_Call {
	_c.Call.Return(run)
	return _c
}

// RemoveBindingsForDevice provides a mock function with given fields: remoteDevice
func (_m *BindingManagerInterface) RemoveBindingsForDevice(remoteDevice api.DeviceRemoteInterface) {
	_m.Called(remoteDevice)
}

// BindingManagerInterface_RemoveBindingsForDevice_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RemoveBindingsForDevice'
type BindingManagerInterface_RemoveBindingsForDevice_Call struct {
	*mock.Call
}

// RemoveBindingsForDevice is a helper method to define mock.On call
//   - remoteDevice api.DeviceRemoteInterface
func (_e *BindingManagerInterface_Expecter) RemoveBindingsForDevice(remoteDevice interface{}) *BindingManagerInterface_RemoveBindingsForDevice_Call {
	return &BindingManagerInterface_RemoveBindingsForDevice_Call{Call: _e.mock.On("RemoveBindingsForDevice", remoteDevice)}
}

func (_c *BindingManagerInterface_RemoveBindingsForDevice_Call) Run(run func(remoteDevice api.DeviceRemoteInterface)) *BindingManagerInterface_RemoveBindingsForDevice_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(api.DeviceRemoteInterface))
	})
	return _c
}

func (_c *BindingManagerInterface_RemoveBindingsForDevice_Call) Return() *BindingManagerInterface_RemoveBindingsForDevice_Call {
	_c.Call.Return()
	return _c
}

func (_c *BindingManagerInterface_RemoveBindingsForDevice_Call) RunAndReturn(run func(api.DeviceRemoteInterface)) *BindingManagerInterface_RemoveBindingsForDevice_Call {
	_c.Call.Return(run)
	return _c
}

// RemoveBindingsForEntity provides a mock function with given fields: remoteEntity
func (_m *BindingManagerInterface) RemoveBindingsForEntity(remoteEntity api.EntityRemoteInterface) {
	_m.Called(remoteEntity)
}

// BindingManagerInterface_RemoveBindingsForEntity_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RemoveBindingsForEntity'
type BindingManagerInterface_RemoveBindingsForEntity_Call struct {
	*mock.Call
}

// RemoveBindingsForEntity is a helper method to define mock.On call
//   - remoteEntity api.EntityRemoteInterface
func (_e *BindingManagerInterface_Expecter) RemoveBindingsForEntity(remoteEntity interface{}) *BindingManagerInterface_RemoveBindingsForEntity_Call {
	return &BindingManagerInterface_RemoveBindingsForEntity_Call{Call: _e.mock.On("RemoveBindingsForEntity", remoteEntity)}
}

func (_c *BindingManagerInterface_RemoveBindingsForEntity_Call) Run(run func(remoteEntity api.EntityRemoteInterface)) *BindingManagerInterface_RemoveBindingsForEntity_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(api.EntityRemoteInterface))
	})
	return _c
}

func (_c *BindingManagerInterface_RemoveBindingsForEntity_Call) Return() *BindingManagerInterface_RemoveBindingsForEntity_Call {
	_c.Call.Return()
	return _c
}

func (_c *BindingManagerInterface_RemoveBindingsForEntity_Call) RunAndReturn(run func(api.EntityRemoteInterface)) *BindingManagerInterface_RemoveBindingsForEntity_Call {
	_c.Call.Return(run)
	return _c
}

// NewBindingManagerInterface creates a new instance of BindingManagerInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewBindingManagerInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *BindingManagerInterface {
	mock := &BindingManagerInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
