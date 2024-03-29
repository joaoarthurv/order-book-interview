// Code generated by mockery v2.22.1. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	model "gitlab.com/projetos/orderbook/orders-management-service/internal/model"
)

// IOrdersDao is an autogenerated mock type for the IOrdersDao type
type IOrdersDao struct {
	mock.Mock
}

type IOrdersDao_Expecter struct {
	mock *mock.Mock
}

func (_m *IOrdersDao) EXPECT() *IOrdersDao_Expecter {
	return &IOrdersDao_Expecter{mock: &_m.Mock}
}

// DeleteOrderById provides a mock function with given fields: ctx, orderId
func (_m *IOrdersDao) DeleteOrderById(ctx context.Context, orderId string) error {
	ret := _m.Called(ctx, orderId)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, orderId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// IOrdersDao_DeleteOrderById_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteOrderById'
type IOrdersDao_DeleteOrderById_Call struct {
	*mock.Call
}

// DeleteOrderById is a helper method to define mock.On call
//   - ctx context.Context
//   - orderId string
func (_e *IOrdersDao_Expecter) DeleteOrderById(ctx interface{}, orderId interface{}) *IOrdersDao_DeleteOrderById_Call {
	return &IOrdersDao_DeleteOrderById_Call{Call: _e.mock.On("DeleteOrderById", ctx, orderId)}
}

func (_c *IOrdersDao_DeleteOrderById_Call) Run(run func(ctx context.Context, orderId string)) *IOrdersDao_DeleteOrderById_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *IOrdersDao_DeleteOrderById_Call) Return(_a0 error) *IOrdersDao_DeleteOrderById_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *IOrdersDao_DeleteOrderById_Call) RunAndReturn(run func(context.Context, string) error) *IOrdersDao_DeleteOrderById_Call {
	_c.Call.Return(run)
	return _c
}

// InsertOrder provides a mock function with given fields: ctx, order
func (_m *IOrdersDao) InsertOrder(ctx context.Context, order *model.OrdersDataRow) error {
	ret := _m.Called(ctx, order)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.OrdersDataRow) error); ok {
		r0 = rf(ctx, order)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// IOrdersDao_InsertOrder_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'InsertOrder'
type IOrdersDao_InsertOrder_Call struct {
	*mock.Call
}

// InsertOrder is a helper method to define mock.On call
//   - ctx context.Context
//   - order *model.OrdersDataRow
func (_e *IOrdersDao_Expecter) InsertOrder(ctx interface{}, order interface{}) *IOrdersDao_InsertOrder_Call {
	return &IOrdersDao_InsertOrder_Call{Call: _e.mock.On("InsertOrder", ctx, order)}
}

func (_c *IOrdersDao_InsertOrder_Call) Run(run func(ctx context.Context, order *model.OrdersDataRow)) *IOrdersDao_InsertOrder_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*model.OrdersDataRow))
	})
	return _c
}

func (_c *IOrdersDao_InsertOrder_Call) Return(_a0 error) *IOrdersDao_InsertOrder_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *IOrdersDao_InsertOrder_Call) RunAndReturn(run func(context.Context, *model.OrdersDataRow) error) *IOrdersDao_InsertOrder_Call {
	_c.Call.Return(run)
	return _c
}

// SelectBuyOrdersByOrderType provides a mock function with given fields: ctx
func (_m *IOrdersDao) SelectBuyOrdersByOrderType(ctx context.Context) ([]model.OrdersDataRow, error) {
	ret := _m.Called(ctx)

	var r0 []model.OrdersDataRow
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]model.OrdersDataRow, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []model.OrdersDataRow); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.OrdersDataRow)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IOrdersDao_SelectBuyOrdersByOrderType_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SelectBuyOrdersByOrderType'
type IOrdersDao_SelectBuyOrdersByOrderType_Call struct {
	*mock.Call
}

// SelectBuyOrdersByOrderType is a helper method to define mock.On call
//   - ctx context.Context
func (_e *IOrdersDao_Expecter) SelectBuyOrdersByOrderType(ctx interface{}) *IOrdersDao_SelectBuyOrdersByOrderType_Call {
	return &IOrdersDao_SelectBuyOrdersByOrderType_Call{Call: _e.mock.On("SelectBuyOrdersByOrderType", ctx)}
}

func (_c *IOrdersDao_SelectBuyOrdersByOrderType_Call) Run(run func(ctx context.Context)) *IOrdersDao_SelectBuyOrdersByOrderType_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *IOrdersDao_SelectBuyOrdersByOrderType_Call) Return(_a0 []model.OrdersDataRow, _a1 error) *IOrdersDao_SelectBuyOrdersByOrderType_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *IOrdersDao_SelectBuyOrdersByOrderType_Call) RunAndReturn(run func(context.Context) ([]model.OrdersDataRow, error)) *IOrdersDao_SelectBuyOrdersByOrderType_Call {
	_c.Call.Return(run)
	return _c
}

// SelectBuyOrdersToExecute provides a mock function with given fields: ctx, order
func (_m *IOrdersDao) SelectBuyOrdersToExecute(ctx context.Context, order *model.OrdersDataRow) ([]model.OrdersDataRow, error) {
	ret := _m.Called(ctx, order)

	var r0 []model.OrdersDataRow
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.OrdersDataRow) ([]model.OrdersDataRow, error)); ok {
		return rf(ctx, order)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *model.OrdersDataRow) []model.OrdersDataRow); ok {
		r0 = rf(ctx, order)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.OrdersDataRow)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *model.OrdersDataRow) error); ok {
		r1 = rf(ctx, order)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IOrdersDao_SelectBuyOrdersToExecute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SelectBuyOrdersToExecute'
type IOrdersDao_SelectBuyOrdersToExecute_Call struct {
	*mock.Call
}

// SelectBuyOrdersToExecute is a helper method to define mock.On call
//   - ctx context.Context
//   - order *model.OrdersDataRow
func (_e *IOrdersDao_Expecter) SelectBuyOrdersToExecute(ctx interface{}, order interface{}) *IOrdersDao_SelectBuyOrdersToExecute_Call {
	return &IOrdersDao_SelectBuyOrdersToExecute_Call{Call: _e.mock.On("SelectBuyOrdersToExecute", ctx, order)}
}

func (_c *IOrdersDao_SelectBuyOrdersToExecute_Call) Run(run func(ctx context.Context, order *model.OrdersDataRow)) *IOrdersDao_SelectBuyOrdersToExecute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*model.OrdersDataRow))
	})
	return _c
}

func (_c *IOrdersDao_SelectBuyOrdersToExecute_Call) Return(_a0 []model.OrdersDataRow, _a1 error) *IOrdersDao_SelectBuyOrdersToExecute_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *IOrdersDao_SelectBuyOrdersToExecute_Call) RunAndReturn(run func(context.Context, *model.OrdersDataRow) ([]model.OrdersDataRow, error)) *IOrdersDao_SelectBuyOrdersToExecute_Call {
	_c.Call.Return(run)
	return _c
}

// SelectSellOrdersByOrderType provides a mock function with given fields: ctx
func (_m *IOrdersDao) SelectSellOrdersByOrderType(ctx context.Context) ([]model.OrdersDataRow, error) {
	ret := _m.Called(ctx)

	var r0 []model.OrdersDataRow
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]model.OrdersDataRow, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []model.OrdersDataRow); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.OrdersDataRow)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IOrdersDao_SelectSellOrdersByOrderType_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SelectSellOrdersByOrderType'
type IOrdersDao_SelectSellOrdersByOrderType_Call struct {
	*mock.Call
}

// SelectSellOrdersByOrderType is a helper method to define mock.On call
//   - ctx context.Context
func (_e *IOrdersDao_Expecter) SelectSellOrdersByOrderType(ctx interface{}) *IOrdersDao_SelectSellOrdersByOrderType_Call {
	return &IOrdersDao_SelectSellOrdersByOrderType_Call{Call: _e.mock.On("SelectSellOrdersByOrderType", ctx)}
}

func (_c *IOrdersDao_SelectSellOrdersByOrderType_Call) Run(run func(ctx context.Context)) *IOrdersDao_SelectSellOrdersByOrderType_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *IOrdersDao_SelectSellOrdersByOrderType_Call) Return(_a0 []model.OrdersDataRow, _a1 error) *IOrdersDao_SelectSellOrdersByOrderType_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *IOrdersDao_SelectSellOrdersByOrderType_Call) RunAndReturn(run func(context.Context) ([]model.OrdersDataRow, error)) *IOrdersDao_SelectSellOrdersByOrderType_Call {
	_c.Call.Return(run)
	return _c
}

// SelectSellOrdersToExecute provides a mock function with given fields: ctx, order
func (_m *IOrdersDao) SelectSellOrdersToExecute(ctx context.Context, order *model.OrdersDataRow) ([]model.OrdersDataRow, error) {
	ret := _m.Called(ctx, order)

	var r0 []model.OrdersDataRow
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.OrdersDataRow) ([]model.OrdersDataRow, error)); ok {
		return rf(ctx, order)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *model.OrdersDataRow) []model.OrdersDataRow); ok {
		r0 = rf(ctx, order)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.OrdersDataRow)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *model.OrdersDataRow) error); ok {
		r1 = rf(ctx, order)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IOrdersDao_SelectSellOrdersToExecute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SelectSellOrdersToExecute'
type IOrdersDao_SelectSellOrdersToExecute_Call struct {
	*mock.Call
}

// SelectSellOrdersToExecute is a helper method to define mock.On call
//   - ctx context.Context
//   - order *model.OrdersDataRow
func (_e *IOrdersDao_Expecter) SelectSellOrdersToExecute(ctx interface{}, order interface{}) *IOrdersDao_SelectSellOrdersToExecute_Call {
	return &IOrdersDao_SelectSellOrdersToExecute_Call{Call: _e.mock.On("SelectSellOrdersToExecute", ctx, order)}
}

func (_c *IOrdersDao_SelectSellOrdersToExecute_Call) Run(run func(ctx context.Context, order *model.OrdersDataRow)) *IOrdersDao_SelectSellOrdersToExecute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*model.OrdersDataRow))
	})
	return _c
}

func (_c *IOrdersDao_SelectSellOrdersToExecute_Call) Return(_a0 []model.OrdersDataRow, _a1 error) *IOrdersDao_SelectSellOrdersToExecute_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *IOrdersDao_SelectSellOrdersToExecute_Call) RunAndReturn(run func(context.Context, *model.OrdersDataRow) ([]model.OrdersDataRow, error)) *IOrdersDao_SelectSellOrdersToExecute_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateOrderQuantityById provides a mock function with given fields: ctx, orderId, quantity
func (_m *IOrdersDao) UpdateOrderQuantityById(ctx context.Context, orderId string, quantity int) error {
	ret := _m.Called(ctx, orderId, quantity)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, int) error); ok {
		r0 = rf(ctx, orderId, quantity)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// IOrdersDao_UpdateOrderQuantityById_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateOrderQuantityById'
type IOrdersDao_UpdateOrderQuantityById_Call struct {
	*mock.Call
}

// UpdateOrderQuantityById is a helper method to define mock.On call
//   - ctx context.Context
//   - orderId string
//   - quantity int
func (_e *IOrdersDao_Expecter) UpdateOrderQuantityById(ctx interface{}, orderId interface{}, quantity interface{}) *IOrdersDao_UpdateOrderQuantityById_Call {
	return &IOrdersDao_UpdateOrderQuantityById_Call{Call: _e.mock.On("UpdateOrderQuantityById", ctx, orderId, quantity)}
}

func (_c *IOrdersDao_UpdateOrderQuantityById_Call) Run(run func(ctx context.Context, orderId string, quantity int)) *IOrdersDao_UpdateOrderQuantityById_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(int))
	})
	return _c
}

func (_c *IOrdersDao_UpdateOrderQuantityById_Call) Return(_a0 error) *IOrdersDao_UpdateOrderQuantityById_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *IOrdersDao_UpdateOrderQuantityById_Call) RunAndReturn(run func(context.Context, string, int) error) *IOrdersDao_UpdateOrderQuantityById_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewIOrdersDao interface {
	mock.TestingT
	Cleanup(func())
}

// NewIOrdersDao creates a new instance of IOrdersDao. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewIOrdersDao(t mockConstructorTestingTNewIOrdersDao) *IOrdersDao {
	mock := &IOrdersDao{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
