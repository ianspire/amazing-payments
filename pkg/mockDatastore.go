// Code generated by MockGen. DO NOT EDIT.
// Source: store.go

package pkg

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockDatastore is a mock of Datastore interface
type MockDatastore struct {
	ctrl     *gomock.Controller
	recorder *MockDatastoreMockRecorder
}

// MockDatastoreMockRecorder is the mock recorder for MockDatastore
type MockDatastoreMockRecorder struct {
	mock *MockDatastore
}

// NewMockDatastore creates a new mock instance
func NewMockDatastore(ctrl *gomock.Controller) *MockDatastore {
	mock := &MockDatastore{ctrl: ctrl}
	mock.recorder = &MockDatastoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDatastore) EXPECT() *MockDatastoreMockRecorder {
	return m.recorder
}

// HealthCheck mocks base method
func (m *MockDatastore) HealthCheck() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HealthCheck")
	ret0, _ := ret[0].(error)
	return ret0
}

// HealthCheck indicates an expected call of HealthCheck
func (mr *MockDatastoreMockRecorder) HealthCheck() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HealthCheck", reflect.TypeOf((*MockDatastore)(nil).HealthCheck))
}

// InsertCustomer mocks base method
func (m *MockDatastore) InsertCustomer(ctx context.Context, name, email, stripeChargeDate, customerKey string) (*CustomerRecord, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertCustomer", ctx, name, email, stripeChargeDate, customerKey)
	ret0, _ := ret[0].(*CustomerRecord)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertCustomer indicates an expected call of InsertCustomer
func (mr *MockDatastoreMockRecorder) InsertCustomer(ctx, name, email, stripeChargeDate, customerKey interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertCustomer", reflect.TypeOf((*MockDatastore)(nil).InsertCustomer), ctx, name, email, stripeChargeDate, customerKey)
}

// GetCustomer mocks base method
func (m *MockDatastore) GetCustomer(ctx context.Context, customerID int64) (*CustomerRecord, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCustomer", ctx, customerID)
	ret0, _ := ret[0].(*CustomerRecord)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCustomer indicates an expected call of GetCustomer
func (mr *MockDatastoreMockRecorder) GetCustomer(ctx, customerID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCustomer", reflect.TypeOf((*MockDatastore)(nil).GetCustomer), ctx, customerID)
}
