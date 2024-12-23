// Code generated by MockGen. DO NOT EDIT.
// Source: handler.go

// Package rating is a generated GoMock package.
package rating

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// Mockstorage is a mock of storage interface.
type Mockstorage struct {
	ctrl     *gomock.Controller
	recorder *MockstorageMockRecorder
}

// MockstorageMockRecorder is the mock recorder for Mockstorage.
type MockstorageMockRecorder struct {
	mock *Mockstorage
}

// NewMockstorage creates a new mock instance.
func NewMockstorage(ctrl *gomock.Controller) *Mockstorage {
	mock := &Mockstorage{ctrl: ctrl}
	mock.recorder = &MockstorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Mockstorage) EXPECT() *MockstorageMockRecorder {
	return m.recorder
}

// CreateRatingRecord mocks base method.
func (m *Mockstorage) CreateRatingRecord(ctx context.Context, record *ratingRecord) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateRatingRecord", ctx, record)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateRatingRecord indicates an expected call of CreateRatingRecord.
func (mr *MockstorageMockRecorder) CreateRatingRecord(ctx, record interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRatingRecord", reflect.TypeOf((*Mockstorage)(nil).CreateRatingRecord), ctx, record)
}

// GetRatingRecord mocks base method.
func (m *Mockstorage) GetRatingRecord(ctx context.Context, username string) (ratingRecord, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRatingRecord", ctx, username)
	ret0, _ := ret[0].(ratingRecord)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRatingRecord indicates an expected call of GetRatingRecord.
func (mr *MockstorageMockRecorder) GetRatingRecord(ctx, username interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRatingRecord", reflect.TypeOf((*Mockstorage)(nil).GetRatingRecord), ctx, username)
}

// UpdateRatingRecord mocks base method.
func (m *Mockstorage) UpdateRatingRecord(ctx context.Context, userName string, record *ratingRecord) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateRatingRecord", ctx, userName, record)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateRatingRecord indicates an expected call of UpdateRatingRecord.
func (mr *MockstorageMockRecorder) UpdateRatingRecord(ctx, userName, record interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateRatingRecord", reflect.TypeOf((*Mockstorage)(nil).UpdateRatingRecord), ctx, userName, record)
}
