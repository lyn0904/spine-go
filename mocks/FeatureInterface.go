// Code generated by mockery v2.45.0. DO NOT EDIT.

package mocks

import (
	api "github.com/lyn0904/spine-go/api"
	mock "github.com/stretchr/testify/mock"

	model "github.com/lyn0904/spine-go/model"
)

// FeatureInterface is an autogenerated mock type for the FeatureInterface type
type FeatureInterface struct {
	mock.Mock
}

type FeatureInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *FeatureInterface) EXPECT() *FeatureInterface_Expecter {
	return &FeatureInterface_Expecter{mock: &_m.Mock}
}

// Address provides a mock function with given fields:
func (_m *FeatureInterface) Address() *model.FeatureAddressType {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Address")
	}

	var r0 *model.FeatureAddressType
	if rf, ok := ret.Get(0).(func() *model.FeatureAddressType); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.FeatureAddressType)
		}
	}

	return r0
}

// FeatureInterface_Address_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Address'
type FeatureInterface_Address_Call struct {
	*mock.Call
}

// Address is a helper method to define mock.On call
func (_e *FeatureInterface_Expecter) Address() *FeatureInterface_Address_Call {
	return &FeatureInterface_Address_Call{Call: _e.mock.On("Address")}
}

func (_c *FeatureInterface_Address_Call) Run(run func()) *FeatureInterface_Address_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *FeatureInterface_Address_Call) Return(_a0 *model.FeatureAddressType) *FeatureInterface_Address_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *FeatureInterface_Address_Call) RunAndReturn(run func() *model.FeatureAddressType) *FeatureInterface_Address_Call {
	_c.Call.Return(run)
	return _c
}

// Description provides a mock function with given fields:
func (_m *FeatureInterface) Description() *model.DescriptionType {
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

// FeatureInterface_Description_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Description'
type FeatureInterface_Description_Call struct {
	*mock.Call
}

// Description is a helper method to define mock.On call
func (_e *FeatureInterface_Expecter) Description() *FeatureInterface_Description_Call {
	return &FeatureInterface_Description_Call{Call: _e.mock.On("Description")}
}

func (_c *FeatureInterface_Description_Call) Run(run func()) *FeatureInterface_Description_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *FeatureInterface_Description_Call) Return(_a0 *model.DescriptionType) *FeatureInterface_Description_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *FeatureInterface_Description_Call) RunAndReturn(run func() *model.DescriptionType) *FeatureInterface_Description_Call {
	_c.Call.Return(run)
	return _c
}

// Operations provides a mock function with given fields:
func (_m *FeatureInterface) Operations() map[model.FunctionType]api.OperationsInterface {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Operations")
	}

	var r0 map[model.FunctionType]api.OperationsInterface
	if rf, ok := ret.Get(0).(func() map[model.FunctionType]api.OperationsInterface); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[model.FunctionType]api.OperationsInterface)
		}
	}

	return r0
}

// FeatureInterface_Operations_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Operations'
type FeatureInterface_Operations_Call struct {
	*mock.Call
}

// Operations is a helper method to define mock.On call
func (_e *FeatureInterface_Expecter) Operations() *FeatureInterface_Operations_Call {
	return &FeatureInterface_Operations_Call{Call: _e.mock.On("Operations")}
}

func (_c *FeatureInterface_Operations_Call) Run(run func()) *FeatureInterface_Operations_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *FeatureInterface_Operations_Call) Return(_a0 map[model.FunctionType]api.OperationsInterface) *FeatureInterface_Operations_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *FeatureInterface_Operations_Call) RunAndReturn(run func() map[model.FunctionType]api.OperationsInterface) *FeatureInterface_Operations_Call {
	_c.Call.Return(run)
	return _c
}

// Role provides a mock function with given fields:
func (_m *FeatureInterface) Role() model.RoleType {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Role")
	}

	var r0 model.RoleType
	if rf, ok := ret.Get(0).(func() model.RoleType); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(model.RoleType)
	}

	return r0
}

// FeatureInterface_Role_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Role'
type FeatureInterface_Role_Call struct {
	*mock.Call
}

// Role is a helper method to define mock.On call
func (_e *FeatureInterface_Expecter) Role() *FeatureInterface_Role_Call {
	return &FeatureInterface_Role_Call{Call: _e.mock.On("Role")}
}

func (_c *FeatureInterface_Role_Call) Run(run func()) *FeatureInterface_Role_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *FeatureInterface_Role_Call) Return(_a0 model.RoleType) *FeatureInterface_Role_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *FeatureInterface_Role_Call) RunAndReturn(run func() model.RoleType) *FeatureInterface_Role_Call {
	_c.Call.Return(run)
	return _c
}

// SetDescription provides a mock function with given fields: desc
func (_m *FeatureInterface) SetDescription(desc *model.DescriptionType) {
	_m.Called(desc)
}

// FeatureInterface_SetDescription_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetDescription'
type FeatureInterface_SetDescription_Call struct {
	*mock.Call
}

// SetDescription is a helper method to define mock.On call
//   - desc *model.DescriptionType
func (_e *FeatureInterface_Expecter) SetDescription(desc interface{}) *FeatureInterface_SetDescription_Call {
	return &FeatureInterface_SetDescription_Call{Call: _e.mock.On("SetDescription", desc)}
}

func (_c *FeatureInterface_SetDescription_Call) Run(run func(desc *model.DescriptionType)) *FeatureInterface_SetDescription_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*model.DescriptionType))
	})
	return _c
}

func (_c *FeatureInterface_SetDescription_Call) Return() *FeatureInterface_SetDescription_Call {
	_c.Call.Return()
	return _c
}

func (_c *FeatureInterface_SetDescription_Call) RunAndReturn(run func(*model.DescriptionType)) *FeatureInterface_SetDescription_Call {
	_c.Call.Return(run)
	return _c
}

// SetDescriptionString provides a mock function with given fields: s
func (_m *FeatureInterface) SetDescriptionString(s string) {
	_m.Called(s)
}

// FeatureInterface_SetDescriptionString_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetDescriptionString'
type FeatureInterface_SetDescriptionString_Call struct {
	*mock.Call
}

// SetDescriptionString is a helper method to define mock.On call
//   - s string
func (_e *FeatureInterface_Expecter) SetDescriptionString(s interface{}) *FeatureInterface_SetDescriptionString_Call {
	return &FeatureInterface_SetDescriptionString_Call{Call: _e.mock.On("SetDescriptionString", s)}
}

func (_c *FeatureInterface_SetDescriptionString_Call) Run(run func(s string)) *FeatureInterface_SetDescriptionString_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *FeatureInterface_SetDescriptionString_Call) Return() *FeatureInterface_SetDescriptionString_Call {
	_c.Call.Return()
	return _c
}

func (_c *FeatureInterface_SetDescriptionString_Call) RunAndReturn(run func(string)) *FeatureInterface_SetDescriptionString_Call {
	_c.Call.Return(run)
	return _c
}

// String provides a mock function with given fields:
func (_m *FeatureInterface) String() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for String")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// FeatureInterface_String_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'String'
type FeatureInterface_String_Call struct {
	*mock.Call
}

// String is a helper method to define mock.On call
func (_e *FeatureInterface_Expecter) String() *FeatureInterface_String_Call {
	return &FeatureInterface_String_Call{Call: _e.mock.On("String")}
}

func (_c *FeatureInterface_String_Call) Run(run func()) *FeatureInterface_String_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *FeatureInterface_String_Call) Return(_a0 string) *FeatureInterface_String_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *FeatureInterface_String_Call) RunAndReturn(run func() string) *FeatureInterface_String_Call {
	_c.Call.Return(run)
	return _c
}

// Type provides a mock function with given fields:
func (_m *FeatureInterface) Type() model.FeatureTypeType {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Type")
	}

	var r0 model.FeatureTypeType
	if rf, ok := ret.Get(0).(func() model.FeatureTypeType); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(model.FeatureTypeType)
	}

	return r0
}

// FeatureInterface_Type_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Type'
type FeatureInterface_Type_Call struct {
	*mock.Call
}

// Type is a helper method to define mock.On call
func (_e *FeatureInterface_Expecter) Type() *FeatureInterface_Type_Call {
	return &FeatureInterface_Type_Call{Call: _e.mock.On("Type")}
}

func (_c *FeatureInterface_Type_Call) Run(run func()) *FeatureInterface_Type_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *FeatureInterface_Type_Call) Return(_a0 model.FeatureTypeType) *FeatureInterface_Type_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *FeatureInterface_Type_Call) RunAndReturn(run func() model.FeatureTypeType) *FeatureInterface_Type_Call {
	_c.Call.Return(run)
	return _c
}

// NewFeatureInterface creates a new instance of FeatureInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewFeatureInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *FeatureInterface {
	mock := &FeatureInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
