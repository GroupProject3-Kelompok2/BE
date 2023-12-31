// Code generated by mockery v2.30.1. DO NOT EDIT.

package mocks

import (
	review "github.com/GroupProject3-Kelompok2/BE/features/review"
	mock "github.com/stretchr/testify/mock"
)

// ReviewData is an autogenerated mock type for the ReviewData type
type ReviewData struct {
	mock.Mock
}

// AddReview provides a mock function with given fields: userId, request
func (_m *ReviewData) AddReview(userId string, request review.ReviewCore) error {
	ret := _m.Called(userId, request)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, review.ReviewCore) error); ok {
		r0 = rf(userId, request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteReview provides a mock function with given fields: userId, reviewId
func (_m *ReviewData) DeleteReview(userId string, reviewId string) error {
	ret := _m.Called(userId, reviewId)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(userId, reviewId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// EditReview provides a mock function with given fields: userId, reviewId, request
func (_m *ReviewData) EditReview(userId string, reviewId string, request review.ReviewCore) error {
	ret := _m.Called(userId, reviewId, request)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, review.ReviewCore) error); ok {
		r0 = rf(userId, reviewId, request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewReviewData creates a new instance of ReviewData. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewReviewData(t interface {
	mock.TestingT
	Cleanup(func())
}) *ReviewData {
	mock := &ReviewData{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
