// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	dto "TTD-golang-gin-test/model/dto"
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// IUpdateShopInfoDomain is an autogenerated mock type for the IUpdateShopInfoDomain type
type IUpdateShopInfoDomain struct {
	mock.Mock
}

// UpdateName provides a mock function with given fields: ctx, request
func (_m *IUpdateShopInfoDomain) UpdateName(ctx context.Context, request dto.UpdateNameRequest) error {
	ret := _m.Called(ctx, request)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, dto.UpdateNameRequest) error); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
