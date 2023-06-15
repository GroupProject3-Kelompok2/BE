// Code generated by mockery v2.30.1. DO NOT EDIT.

package mocks

import (
	reservation "github.com/GroupProject3-Kelompok2/BE/features/reservation"
	mock "github.com/stretchr/testify/mock"
)

// ReservationDataInterface is an autogenerated mock type for the ReservationDataInterface type
type ReservationDataInterface struct {
	mock.Mock
}

// Insert provides a mock function with given fields: input
func (_m *ReservationDataInterface) Insert(input reservation.ReservationCore) error {
	ret := _m.Called(input)

	var r0 error
	if rf, ok := ret.Get(0).(func(reservation.ReservationCore) error); ok {
		r0 = rf(input)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewReservationDataInterface creates a new instance of ReservationDataInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewReservationDataInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *ReservationDataInterface {
	mock := &ReservationDataInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
