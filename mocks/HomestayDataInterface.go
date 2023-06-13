// Code generated by mockery v2.29.0. DO NOT EDIT.

package mocks

import (
	homestay "github.com/GroupProject3-Kelompok2/BE/features/homestay"
	mock "github.com/stretchr/testify/mock"
)

// HomestayDataInterface is an autogenerated mock type for the HomestayDataInterface type
type HomestayDataInterface struct {
	mock.Mock
}

// DeleteById provides a mock function with given fields: userId, homestayId
func (_m *HomestayDataInterface) DeleteById(userId string, homestayId string) error {
	ret := _m.Called(userId, homestayId)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(userId, homestayId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// HomestayPictures provides a mock function with given fields: homestayId, req
func (_m *HomestayDataInterface) HomestayPictures(homestayId string, req homestay.HomestayPictureCore) error {
	ret := _m.Called(homestayId, req)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, homestay.HomestayPictureCore) error); ok {
		r0 = rf(homestayId, req)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Insert provides a mock function with given fields: input
func (_m *HomestayDataInterface) Insert(input homestay.HomestayCore) error {
	ret := _m.Called(input)

	var r0 error
	if rf, ok := ret.Get(0).(func(homestay.HomestayCore) error); ok {
		r0 = rf(input)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SelectAll provides a mock function with given fields:
func (_m *HomestayDataInterface) SelectAll() ([]homestay.HomestayCore, error) {
	ret := _m.Called()

	var r0 []homestay.HomestayCore
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]homestay.HomestayCore, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []homestay.HomestayCore); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]homestay.HomestayCore)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SelectById provides a mock function with given fields: homestayID
func (_m *HomestayDataInterface) SelectById(homestayID string) (homestay.HomestayCore, error) {
	ret := _m.Called(homestayID)

	var r0 homestay.HomestayCore
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (homestay.HomestayCore, error)); ok {
		return rf(homestayID)
	}
	if rf, ok := ret.Get(0).(func(string) homestay.HomestayCore); ok {
		r0 = rf(homestayID)
	} else {
		r0 = ret.Get(0).(homestay.HomestayCore)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(homestayID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateById provides a mock function with given fields: userId, homestayId, input
func (_m *HomestayDataInterface) UpdateById(userId string, homestayId string, input homestay.HomestayCore) error {
	ret := _m.Called(userId, homestayId, input)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, homestay.HomestayCore) error); ok {
		r0 = rf(userId, homestayId, input)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewHomestayDataInterface interface {
	mock.TestingT
	Cleanup(func())
}

// NewHomestayDataInterface creates a new instance of HomestayDataInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewHomestayDataInterface(t mockConstructorTestingTNewHomestayDataInterface) *HomestayDataInterface {
	mock := &HomestayDataInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
