// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/numaproj/numaflow-go/pkg/apis/proto/function/v1 (interfaces: UserDefinedFunctionClient)

// Package funcmock is a generated GoMock package.
package funcmock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "github.com/numaproj/numaflow-go/pkg/apis/proto/function/v1"
	grpc "google.golang.org/grpc"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// MockUserDefinedFunctionClient is a mock of UserDefinedFunctionClient interface.
type MockUserDefinedFunctionClient struct {
	ctrl     *gomock.Controller
	recorder *MockUserDefinedFunctionClientMockRecorder
}

// MockUserDefinedFunctionClientMockRecorder is the mock recorder for MockUserDefinedFunctionClient.
type MockUserDefinedFunctionClientMockRecorder struct {
	mock *MockUserDefinedFunctionClient
}

// NewMockUserDefinedFunctionClient creates a new mock instance.
func NewMockUserDefinedFunctionClient(ctrl *gomock.Controller) *MockUserDefinedFunctionClient {
	mock := &MockUserDefinedFunctionClient{ctrl: ctrl}
	mock.recorder = &MockUserDefinedFunctionClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserDefinedFunctionClient) EXPECT() *MockUserDefinedFunctionClientMockRecorder {
	return m.recorder
}

// DoFn mocks base method.
func (m *MockUserDefinedFunctionClient) DoFn(arg0 context.Context, arg1 *v1.Datum, arg2 ...grpc.CallOption) (*v1.DatumList, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DoFn", varargs...)
	ret0, _ := ret[0].(*v1.DatumList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DoFn indicates an expected call of DoFn.
func (mr *MockUserDefinedFunctionClientMockRecorder) DoFn(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DoFn", reflect.TypeOf((*MockUserDefinedFunctionClient)(nil).DoFn), varargs...)
}

// IsReady mocks base method.
func (m *MockUserDefinedFunctionClient) IsReady(arg0 context.Context, arg1 *emptypb.Empty, arg2 ...grpc.CallOption) (*v1.ReadyResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "IsReady", varargs...)
	ret0, _ := ret[0].(*v1.ReadyResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsReady indicates an expected call of IsReady.
func (mr *MockUserDefinedFunctionClientMockRecorder) IsReady(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsReady", reflect.TypeOf((*MockUserDefinedFunctionClient)(nil).IsReady), varargs...)
}

// ReduceFn mocks base method.
func (m *MockUserDefinedFunctionClient) ReduceFn(arg0 context.Context, arg1 ...grpc.CallOption) (v1.UserDefinedFunction_ReduceFnClient, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ReduceFn", varargs...)
	ret0, _ := ret[0].(v1.UserDefinedFunction_ReduceFnClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReduceFn indicates an expected call of ReduceFn.
func (mr *MockUserDefinedFunctionClientMockRecorder) ReduceFn(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReduceFn", reflect.TypeOf((*MockUserDefinedFunctionClient)(nil).ReduceFn), varargs...)
}
