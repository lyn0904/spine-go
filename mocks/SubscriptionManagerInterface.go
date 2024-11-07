// Code generated by mockery v2.45.0. DO NOT EDIT.

package mocks

import (
	api "github.com/lyn0904/spine-go/api"
	mock "github.com/stretchr/testify/mock"

	model "github.com/lyn0904/spine-go/model"
)

// SubscriptionManagerInterface is an autogenerated mock type for the SubscriptionManagerInterface type
type SubscriptionManagerInterface struct {
	mock.Mock
}

type SubscriptionManagerInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *SubscriptionManagerInterface) EXPECT() *SubscriptionManagerInterface_Expecter {
	return &SubscriptionManagerInterface_Expecter{mock: &_m.Mock}
}

// AddSubscription provides a mock function with given fields: remoteDevice, data
func (_m *SubscriptionManagerInterface) AddSubscription(remoteDevice api.DeviceRemoteInterface, data model.SubscriptionManagementRequestCallType) error {
	ret := _m.Called(remoteDevice, data)

	if len(ret) == 0 {
		panic("no return value specified for AddSubscription")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(api.DeviceRemoteInterface, model.SubscriptionManagementRequestCallType) error); ok {
		r0 = rf(remoteDevice, data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SubscriptionManagerInterface_AddSubscription_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddSubscription'
type SubscriptionManagerInterface_AddSubscription_Call struct {
	*mock.Call
}

// AddSubscription is a helper method to define mock.On call
//   - remoteDevice api.DeviceRemoteInterface
//   - data model.SubscriptionManagementRequestCallType
func (_e *SubscriptionManagerInterface_Expecter) AddSubscription(remoteDevice interface{}, data interface{}) *SubscriptionManagerInterface_AddSubscription_Call {
	return &SubscriptionManagerInterface_AddSubscription_Call{Call: _e.mock.On("AddSubscription", remoteDevice, data)}
}

func (_c *SubscriptionManagerInterface_AddSubscription_Call) Run(run func(remoteDevice api.DeviceRemoteInterface, data model.SubscriptionManagementRequestCallType)) *SubscriptionManagerInterface_AddSubscription_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(api.DeviceRemoteInterface), args[1].(model.SubscriptionManagementRequestCallType))
	})
	return _c
}

func (_c *SubscriptionManagerInterface_AddSubscription_Call) Return(_a0 error) *SubscriptionManagerInterface_AddSubscription_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *SubscriptionManagerInterface_AddSubscription_Call) RunAndReturn(run func(api.DeviceRemoteInterface, model.SubscriptionManagementRequestCallType) error) *SubscriptionManagerInterface_AddSubscription_Call {
	_c.Call.Return(run)
	return _c
}

// RemoveSubscription provides a mock function with given fields: data, remoteDevice
func (_m *SubscriptionManagerInterface) RemoveSubscription(data model.SubscriptionManagementDeleteCallType, remoteDevice api.DeviceRemoteInterface) error {
	ret := _m.Called(data, remoteDevice)

	if len(ret) == 0 {
		panic("no return value specified for RemoveSubscription")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(model.SubscriptionManagementDeleteCallType, api.DeviceRemoteInterface) error); ok {
		r0 = rf(data, remoteDevice)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SubscriptionManagerInterface_RemoveSubscription_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RemoveSubscription'
type SubscriptionManagerInterface_RemoveSubscription_Call struct {
	*mock.Call
}

// RemoveSubscription is a helper method to define mock.On call
//   - data model.SubscriptionManagementDeleteCallType
//   - remoteDevice api.DeviceRemoteInterface
func (_e *SubscriptionManagerInterface_Expecter) RemoveSubscription(data interface{}, remoteDevice interface{}) *SubscriptionManagerInterface_RemoveSubscription_Call {
	return &SubscriptionManagerInterface_RemoveSubscription_Call{Call: _e.mock.On("RemoveSubscription", data, remoteDevice)}
}

func (_c *SubscriptionManagerInterface_RemoveSubscription_Call) Run(run func(data model.SubscriptionManagementDeleteCallType, remoteDevice api.DeviceRemoteInterface)) *SubscriptionManagerInterface_RemoveSubscription_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(model.SubscriptionManagementDeleteCallType), args[1].(api.DeviceRemoteInterface))
	})
	return _c
}

func (_c *SubscriptionManagerInterface_RemoveSubscription_Call) Return(_a0 error) *SubscriptionManagerInterface_RemoveSubscription_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *SubscriptionManagerInterface_RemoveSubscription_Call) RunAndReturn(run func(model.SubscriptionManagementDeleteCallType, api.DeviceRemoteInterface) error) *SubscriptionManagerInterface_RemoveSubscription_Call {
	_c.Call.Return(run)
	return _c
}

// RemoveSubscriptionsForDevice provides a mock function with given fields: remoteDevice
func (_m *SubscriptionManagerInterface) RemoveSubscriptionsForDevice(remoteDevice api.DeviceRemoteInterface) {
	_m.Called(remoteDevice)
}

// SubscriptionManagerInterface_RemoveSubscriptionsForDevice_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RemoveSubscriptionsForDevice'
type SubscriptionManagerInterface_RemoveSubscriptionsForDevice_Call struct {
	*mock.Call
}

// RemoveSubscriptionsForDevice is a helper method to define mock.On call
//   - remoteDevice api.DeviceRemoteInterface
func (_e *SubscriptionManagerInterface_Expecter) RemoveSubscriptionsForDevice(remoteDevice interface{}) *SubscriptionManagerInterface_RemoveSubscriptionsForDevice_Call {
	return &SubscriptionManagerInterface_RemoveSubscriptionsForDevice_Call{Call: _e.mock.On("RemoveSubscriptionsForDevice", remoteDevice)}
}

func (_c *SubscriptionManagerInterface_RemoveSubscriptionsForDevice_Call) Run(run func(remoteDevice api.DeviceRemoteInterface)) *SubscriptionManagerInterface_RemoveSubscriptionsForDevice_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(api.DeviceRemoteInterface))
	})
	return _c
}

func (_c *SubscriptionManagerInterface_RemoveSubscriptionsForDevice_Call) Return() *SubscriptionManagerInterface_RemoveSubscriptionsForDevice_Call {
	_c.Call.Return()
	return _c
}

func (_c *SubscriptionManagerInterface_RemoveSubscriptionsForDevice_Call) RunAndReturn(run func(api.DeviceRemoteInterface)) *SubscriptionManagerInterface_RemoveSubscriptionsForDevice_Call {
	_c.Call.Return(run)
	return _c
}

// RemoveSubscriptionsForEntity provides a mock function with given fields: remoteEntity
func (_m *SubscriptionManagerInterface) RemoveSubscriptionsForEntity(remoteEntity api.EntityRemoteInterface) {
	_m.Called(remoteEntity)
}

// SubscriptionManagerInterface_RemoveSubscriptionsForEntity_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RemoveSubscriptionsForEntity'
type SubscriptionManagerInterface_RemoveSubscriptionsForEntity_Call struct {
	*mock.Call
}

// RemoveSubscriptionsForEntity is a helper method to define mock.On call
//   - remoteEntity api.EntityRemoteInterface
func (_e *SubscriptionManagerInterface_Expecter) RemoveSubscriptionsForEntity(remoteEntity interface{}) *SubscriptionManagerInterface_RemoveSubscriptionsForEntity_Call {
	return &SubscriptionManagerInterface_RemoveSubscriptionsForEntity_Call{Call: _e.mock.On("RemoveSubscriptionsForEntity", remoteEntity)}
}

func (_c *SubscriptionManagerInterface_RemoveSubscriptionsForEntity_Call) Run(run func(remoteEntity api.EntityRemoteInterface)) *SubscriptionManagerInterface_RemoveSubscriptionsForEntity_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(api.EntityRemoteInterface))
	})
	return _c
}

func (_c *SubscriptionManagerInterface_RemoveSubscriptionsForEntity_Call) Return() *SubscriptionManagerInterface_RemoveSubscriptionsForEntity_Call {
	_c.Call.Return()
	return _c
}

func (_c *SubscriptionManagerInterface_RemoveSubscriptionsForEntity_Call) RunAndReturn(run func(api.EntityRemoteInterface)) *SubscriptionManagerInterface_RemoveSubscriptionsForEntity_Call {
	_c.Call.Return(run)
	return _c
}

// Subscriptions provides a mock function with given fields: remoteDevice
func (_m *SubscriptionManagerInterface) Subscriptions(remoteDevice api.DeviceRemoteInterface) []*api.SubscriptionEntry {
	ret := _m.Called(remoteDevice)

	if len(ret) == 0 {
		panic("no return value specified for Subscriptions")
	}

	var r0 []*api.SubscriptionEntry
	if rf, ok := ret.Get(0).(func(api.DeviceRemoteInterface) []*api.SubscriptionEntry); ok {
		r0 = rf(remoteDevice)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*api.SubscriptionEntry)
		}
	}

	return r0
}

// SubscriptionManagerInterface_Subscriptions_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Subscriptions'
type SubscriptionManagerInterface_Subscriptions_Call struct {
	*mock.Call
}

// Subscriptions is a helper method to define mock.On call
//   - remoteDevice api.DeviceRemoteInterface
func (_e *SubscriptionManagerInterface_Expecter) Subscriptions(remoteDevice interface{}) *SubscriptionManagerInterface_Subscriptions_Call {
	return &SubscriptionManagerInterface_Subscriptions_Call{Call: _e.mock.On("Subscriptions", remoteDevice)}
}

func (_c *SubscriptionManagerInterface_Subscriptions_Call) Run(run func(remoteDevice api.DeviceRemoteInterface)) *SubscriptionManagerInterface_Subscriptions_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(api.DeviceRemoteInterface))
	})
	return _c
}

func (_c *SubscriptionManagerInterface_Subscriptions_Call) Return(_a0 []*api.SubscriptionEntry) *SubscriptionManagerInterface_Subscriptions_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *SubscriptionManagerInterface_Subscriptions_Call) RunAndReturn(run func(api.DeviceRemoteInterface) []*api.SubscriptionEntry) *SubscriptionManagerInterface_Subscriptions_Call {
	_c.Call.Return(run)
	return _c
}

// SubscriptionsOnFeature provides a mock function with given fields: featureAddress
func (_m *SubscriptionManagerInterface) SubscriptionsOnFeature(featureAddress model.FeatureAddressType) []*api.SubscriptionEntry {
	ret := _m.Called(featureAddress)

	if len(ret) == 0 {
		panic("no return value specified for SubscriptionsOnFeature")
	}

	var r0 []*api.SubscriptionEntry
	if rf, ok := ret.Get(0).(func(model.FeatureAddressType) []*api.SubscriptionEntry); ok {
		r0 = rf(featureAddress)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*api.SubscriptionEntry)
		}
	}

	return r0
}

// SubscriptionManagerInterface_SubscriptionsOnFeature_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SubscriptionsOnFeature'
type SubscriptionManagerInterface_SubscriptionsOnFeature_Call struct {
	*mock.Call
}

// SubscriptionsOnFeature is a helper method to define mock.On call
//   - featureAddress model.FeatureAddressType
func (_e *SubscriptionManagerInterface_Expecter) SubscriptionsOnFeature(featureAddress interface{}) *SubscriptionManagerInterface_SubscriptionsOnFeature_Call {
	return &SubscriptionManagerInterface_SubscriptionsOnFeature_Call{Call: _e.mock.On("SubscriptionsOnFeature", featureAddress)}
}

func (_c *SubscriptionManagerInterface_SubscriptionsOnFeature_Call) Run(run func(featureAddress model.FeatureAddressType)) *SubscriptionManagerInterface_SubscriptionsOnFeature_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(model.FeatureAddressType))
	})
	return _c
}

func (_c *SubscriptionManagerInterface_SubscriptionsOnFeature_Call) Return(_a0 []*api.SubscriptionEntry) *SubscriptionManagerInterface_SubscriptionsOnFeature_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *SubscriptionManagerInterface_SubscriptionsOnFeature_Call) RunAndReturn(run func(model.FeatureAddressType) []*api.SubscriptionEntry) *SubscriptionManagerInterface_SubscriptionsOnFeature_Call {
	_c.Call.Return(run)
	return _c
}

// NewSubscriptionManagerInterface creates a new instance of SubscriptionManagerInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewSubscriptionManagerInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *SubscriptionManagerInterface {
	mock := &SubscriptionManagerInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
