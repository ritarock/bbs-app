// Code generated by MockGen. DO NOT EDIT.
// Source: ./port/comment.go
//
// Generated by this command:
//
//	mockgen -source=./port/comment.go -destination=./testing/mock/comment.go -package=mock
//

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	domain "github.com/ritarock/bbs-app/domain"
	gomock "go.uber.org/mock/gomock"
)

// MockCommentRepository is a mock of CommentRepository interface.
type MockCommentRepository struct {
	ctrl     *gomock.Controller
	recorder *MockCommentRepositoryMockRecorder
}

// MockCommentRepositoryMockRecorder is the mock recorder for MockCommentRepository.
type MockCommentRepositoryMockRecorder struct {
	mock *MockCommentRepository
}

// NewMockCommentRepository creates a new mock instance.
func NewMockCommentRepository(ctrl *gomock.Controller) *MockCommentRepository {
	mock := &MockCommentRepository{ctrl: ctrl}
	mock.recorder = &MockCommentRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCommentRepository) EXPECT() *MockCommentRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockCommentRepository) Create(ctx context.Context, comment *domain.Comment) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, comment)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockCommentRepositoryMockRecorder) Create(ctx, comment any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockCommentRepository)(nil).Create), ctx, comment)
}

// GetAll mocks base method.
func (m *MockCommentRepository) GetAll(ctx context.Context, postID int) ([]*domain.Comment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", ctx, postID)
	ret0, _ := ret[0].([]*domain.Comment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockCommentRepositoryMockRecorder) GetAll(ctx, postID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockCommentRepository)(nil).GetAll), ctx, postID)
}

// MockCommentUsecase is a mock of CommentUsecase interface.
type MockCommentUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockCommentUsecaseMockRecorder
}

// MockCommentUsecaseMockRecorder is the mock recorder for MockCommentUsecase.
type MockCommentUsecaseMockRecorder struct {
	mock *MockCommentUsecase
}

// NewMockCommentUsecase creates a new mock instance.
func NewMockCommentUsecase(ctrl *gomock.Controller) *MockCommentUsecase {
	mock := &MockCommentUsecase{ctrl: ctrl}
	mock.recorder = &MockCommentUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCommentUsecase) EXPECT() *MockCommentUsecaseMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockCommentUsecase) Create(ctx context.Context, comment *domain.Comment) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, comment)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockCommentUsecaseMockRecorder) Create(ctx, comment any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockCommentUsecase)(nil).Create), ctx, comment)
}

// GetAll mocks base method.
func (m *MockCommentUsecase) GetAll(ctx context.Context, postID int) ([]*domain.Comment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", ctx, postID)
	ret0, _ := ret[0].([]*domain.Comment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockCommentUsecaseMockRecorder) GetAll(ctx, postID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockCommentUsecase)(nil).GetAll), ctx, postID)
}