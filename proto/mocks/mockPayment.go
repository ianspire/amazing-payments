// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ianspire/amazing-payments/proto (interfaces: PaymentServiceClient,PaymentServiceServer)

// Package mock_proto is a generated GoMock package.
package mock_proto

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	payment "github.com/ianspire/amazing-payments/proto"
	grpc "google.golang.org/grpc"
	reflect "reflect"
)

// MockPaymentServiceClient is a mock of PaymentServiceClient interface
type MockPaymentServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockPaymentServiceClientMockRecorder
}

// MockPaymentServiceClientMockRecorder is the mock recorder for MockPaymentServiceClient
type MockPaymentServiceClientMockRecorder struct {
	mock *MockPaymentServiceClient
}

// NewMockPaymentServiceClient creates a new mock instance
func NewMockPaymentServiceClient(ctrl *gomock.Controller) *MockPaymentServiceClient {
	mock := &MockPaymentServiceClient{ctrl: ctrl}
	mock.recorder = &MockPaymentServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPaymentServiceClient) EXPECT() *MockPaymentServiceClientMockRecorder {
	return m.recorder
}

// CreateCustomer mocks base method
func (m *MockPaymentServiceClient) CreateCustomer(arg0 context.Context, arg1 *payment.CreateCustomerRequest, arg2 ...grpc.CallOption) (*payment.Customer, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateCustomer", varargs...)
	ret0, _ := ret[0].(*payment.Customer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateCustomer indicates an expected call of CreateCustomer
func (mr *MockPaymentServiceClientMockRecorder) CreateCustomer(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCustomer", reflect.TypeOf((*MockPaymentServiceClient)(nil).CreateCustomer), varargs...)
}

// GetCustomer mocks base method
func (m *MockPaymentServiceClient) GetCustomer(arg0 context.Context, arg1 *payment.GetCustomerRequest, arg2 ...grpc.CallOption) (*payment.Customer, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetCustomer", varargs...)
	ret0, _ := ret[0].(*payment.Customer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCustomer indicates an expected call of GetCustomer
func (mr *MockPaymentServiceClientMockRecorder) GetCustomer(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCustomer", reflect.TypeOf((*MockPaymentServiceClient)(nil).GetCustomer), varargs...)
}

// HealthCheck mocks base method
func (m *MockPaymentServiceClient) HealthCheck(arg0 context.Context, arg1 *payment.HealthCheckRequest, arg2 ...grpc.CallOption) (*payment.HealthCheckResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "HealthCheck", varargs...)
	ret0, _ := ret[0].(*payment.HealthCheckResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HealthCheck indicates an expected call of HealthCheck
func (mr *MockPaymentServiceClientMockRecorder) HealthCheck(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HealthCheck", reflect.TypeOf((*MockPaymentServiceClient)(nil).HealthCheck), varargs...)
}

// MockPaymentServiceServer is a mock of PaymentServiceServer interface
type MockPaymentServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockPaymentServiceServerMockRecorder
}

// MockPaymentServiceServerMockRecorder is the mock recorder for MockPaymentServiceServer
type MockPaymentServiceServerMockRecorder struct {
	mock *MockPaymentServiceServer
}

// NewMockPaymentServiceServer creates a new mock instance
func NewMockPaymentServiceServer(ctrl *gomock.Controller) *MockPaymentServiceServer {
	mock := &MockPaymentServiceServer{ctrl: ctrl}
	mock.recorder = &MockPaymentServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPaymentServiceServer) EXPECT() *MockPaymentServiceServerMockRecorder {
	return m.recorder
}

// CreateCustomer mocks base method
func (m *MockPaymentServiceServer) CreateCustomer(arg0 context.Context, arg1 *payment.CreateCustomerRequest) (*payment.Customer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCustomer", arg0, arg1)
	ret0, _ := ret[0].(*payment.Customer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateCustomer indicates an expected call of CreateCustomer
func (mr *MockPaymentServiceServerMockRecorder) CreateCustomer(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCustomer", reflect.TypeOf((*MockPaymentServiceServer)(nil).CreateCustomer), arg0, arg1)
}

// GetCustomer mocks base method
func (m *MockPaymentServiceServer) GetCustomer(arg0 context.Context, arg1 *payment.GetCustomerRequest) (*payment.Customer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCustomer", arg0, arg1)
	ret0, _ := ret[0].(*payment.Customer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCustomer indicates an expected call of GetCustomer
func (mr *MockPaymentServiceServerMockRecorder) GetCustomer(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCustomer", reflect.TypeOf((*MockPaymentServiceServer)(nil).GetCustomer), arg0, arg1)
}

// HealthCheck mocks base method
func (m *MockPaymentServiceServer) HealthCheck(arg0 context.Context, arg1 *payment.HealthCheckRequest) (*payment.HealthCheckResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HealthCheck", arg0, arg1)
	ret0, _ := ret[0].(*payment.HealthCheckResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HealthCheck indicates an expected call of HealthCheck
func (mr *MockPaymentServiceServerMockRecorder) HealthCheck(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HealthCheck", reflect.TypeOf((*MockPaymentServiceServer)(nil).HealthCheck), arg0, arg1)
}