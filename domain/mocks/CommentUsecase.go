// Code generated by mockery v2.29.0. DO NOT EDIT.

package mocks

import (
	context "context"
	domain "ritarock/bbs-app/domain"

	mock "github.com/stretchr/testify/mock"
)

// CommentUsecase is an autogenerated mock type for the CommentUsecase type
type CommentUsecase struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, comment
func (_m *CommentUsecase) Create(ctx context.Context, comment *domain.Comment) error {
	ret := _m.Called(ctx, comment)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Comment) error); ok {
		r0 = rf(ctx, comment)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllByPost provides a mock function with given fields: ctx, postId
func (_m *CommentUsecase) GetAllByPost(ctx context.Context, postId int) ([]domain.Comment, error) {
	ret := _m.Called(ctx, postId)

	var r0 []domain.Comment
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int) ([]domain.Comment, error)); ok {
		return rf(ctx, postId)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int) []domain.Comment); ok {
		r0 = rf(ctx, postId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Comment)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, postId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewCommentUsecase interface {
	mock.TestingT
	Cleanup(func())
}

// NewCommentUsecase creates a new instance of CommentUsecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCommentUsecase(t mockConstructorTestingTNewCommentUsecase) *CommentUsecase {
	mock := &CommentUsecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
