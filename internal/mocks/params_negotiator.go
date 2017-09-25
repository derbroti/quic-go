// Code generated by MockGen. DO NOT EDIT.
// Source: ../handshake/params_negotiator_base.go

package mocks

import (
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	protocol "github.com/lucas-clemente/quic-go/internal/protocol"
)

// MockParamsNegotiator is a mock of ParamsNegotiator interface
type MockParamsNegotiator struct {
	ctrl     *gomock.Controller
	recorder *MockParamsNegotiatorMockRecorder
}

// MockParamsNegotiatorMockRecorder is the mock recorder for MockParamsNegotiator
type MockParamsNegotiatorMockRecorder struct {
	mock *MockParamsNegotiator
}

// NewMockParamsNegotiator creates a new mock instance
func NewMockParamsNegotiator(ctrl *gomock.Controller) *MockParamsNegotiator {
	mock := &MockParamsNegotiator{ctrl: ctrl}
	mock.recorder = &MockParamsNegotiatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockParamsNegotiator) EXPECT() *MockParamsNegotiatorMockRecorder {
	return _m.recorder
}

// GetSendStreamFlowControlWindow mocks base method
func (_m *MockParamsNegotiator) GetSendStreamFlowControlWindow() protocol.ByteCount {
	ret := _m.ctrl.Call(_m, "GetSendStreamFlowControlWindow")
	ret0, _ := ret[0].(protocol.ByteCount)
	return ret0
}

// GetSendStreamFlowControlWindow indicates an expected call of GetSendStreamFlowControlWindow
func (_mr *MockParamsNegotiatorMockRecorder) GetSendStreamFlowControlWindow() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "GetSendStreamFlowControlWindow", reflect.TypeOf((*MockParamsNegotiator)(nil).GetSendStreamFlowControlWindow))
}

// GetSendConnectionFlowControlWindow mocks base method
func (_m *MockParamsNegotiator) GetSendConnectionFlowControlWindow() protocol.ByteCount {
	ret := _m.ctrl.Call(_m, "GetSendConnectionFlowControlWindow")
	ret0, _ := ret[0].(protocol.ByteCount)
	return ret0
}

// GetSendConnectionFlowControlWindow indicates an expected call of GetSendConnectionFlowControlWindow
func (_mr *MockParamsNegotiatorMockRecorder) GetSendConnectionFlowControlWindow() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "GetSendConnectionFlowControlWindow", reflect.TypeOf((*MockParamsNegotiator)(nil).GetSendConnectionFlowControlWindow))
}

// GetReceiveStreamFlowControlWindow mocks base method
func (_m *MockParamsNegotiator) GetReceiveStreamFlowControlWindow() protocol.ByteCount {
	ret := _m.ctrl.Call(_m, "GetReceiveStreamFlowControlWindow")
	ret0, _ := ret[0].(protocol.ByteCount)
	return ret0
}

// GetReceiveStreamFlowControlWindow indicates an expected call of GetReceiveStreamFlowControlWindow
func (_mr *MockParamsNegotiatorMockRecorder) GetReceiveStreamFlowControlWindow() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "GetReceiveStreamFlowControlWindow", reflect.TypeOf((*MockParamsNegotiator)(nil).GetReceiveStreamFlowControlWindow))
}

// GetMaxReceiveStreamFlowControlWindow mocks base method
func (_m *MockParamsNegotiator) GetMaxReceiveStreamFlowControlWindow() protocol.ByteCount {
	ret := _m.ctrl.Call(_m, "GetMaxReceiveStreamFlowControlWindow")
	ret0, _ := ret[0].(protocol.ByteCount)
	return ret0
}

// GetMaxReceiveStreamFlowControlWindow indicates an expected call of GetMaxReceiveStreamFlowControlWindow
func (_mr *MockParamsNegotiatorMockRecorder) GetMaxReceiveStreamFlowControlWindow() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "GetMaxReceiveStreamFlowControlWindow", reflect.TypeOf((*MockParamsNegotiator)(nil).GetMaxReceiveStreamFlowControlWindow))
}

// GetReceiveConnectionFlowControlWindow mocks base method
func (_m *MockParamsNegotiator) GetReceiveConnectionFlowControlWindow() protocol.ByteCount {
	ret := _m.ctrl.Call(_m, "GetReceiveConnectionFlowControlWindow")
	ret0, _ := ret[0].(protocol.ByteCount)
	return ret0
}

// GetReceiveConnectionFlowControlWindow indicates an expected call of GetReceiveConnectionFlowControlWindow
func (_mr *MockParamsNegotiatorMockRecorder) GetReceiveConnectionFlowControlWindow() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "GetReceiveConnectionFlowControlWindow", reflect.TypeOf((*MockParamsNegotiator)(nil).GetReceiveConnectionFlowControlWindow))
}

// GetMaxReceiveConnectionFlowControlWindow mocks base method
func (_m *MockParamsNegotiator) GetMaxReceiveConnectionFlowControlWindow() protocol.ByteCount {
	ret := _m.ctrl.Call(_m, "GetMaxReceiveConnectionFlowControlWindow")
	ret0, _ := ret[0].(protocol.ByteCount)
	return ret0
}

// GetMaxReceiveConnectionFlowControlWindow indicates an expected call of GetMaxReceiveConnectionFlowControlWindow
func (_mr *MockParamsNegotiatorMockRecorder) GetMaxReceiveConnectionFlowControlWindow() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "GetMaxReceiveConnectionFlowControlWindow", reflect.TypeOf((*MockParamsNegotiator)(nil).GetMaxReceiveConnectionFlowControlWindow))
}

// GetMaxOutgoingStreams mocks base method
func (_m *MockParamsNegotiator) GetMaxOutgoingStreams() uint32 {
	ret := _m.ctrl.Call(_m, "GetMaxOutgoingStreams")
	ret0, _ := ret[0].(uint32)
	return ret0
}

// GetMaxOutgoingStreams indicates an expected call of GetMaxOutgoingStreams
func (_mr *MockParamsNegotiatorMockRecorder) GetMaxOutgoingStreams() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "GetMaxOutgoingStreams", reflect.TypeOf((*MockParamsNegotiator)(nil).GetMaxOutgoingStreams))
}

// GetMaxIncomingStreams mocks base method
func (_m *MockParamsNegotiator) GetMaxIncomingStreams() uint32 {
	ret := _m.ctrl.Call(_m, "GetMaxIncomingStreams")
	ret0, _ := ret[0].(uint32)
	return ret0
}

// GetMaxIncomingStreams indicates an expected call of GetMaxIncomingStreams
func (_mr *MockParamsNegotiatorMockRecorder) GetMaxIncomingStreams() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "GetMaxIncomingStreams", reflect.TypeOf((*MockParamsNegotiator)(nil).GetMaxIncomingStreams))
}

// GetIdleConnectionStateLifetime mocks base method
func (_m *MockParamsNegotiator) GetIdleConnectionStateLifetime() time.Duration {
	ret := _m.ctrl.Call(_m, "GetIdleConnectionStateLifetime")
	ret0, _ := ret[0].(time.Duration)
	return ret0
}

// GetIdleConnectionStateLifetime indicates an expected call of GetIdleConnectionStateLifetime
func (_mr *MockParamsNegotiatorMockRecorder) GetIdleConnectionStateLifetime() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "GetIdleConnectionStateLifetime", reflect.TypeOf((*MockParamsNegotiator)(nil).GetIdleConnectionStateLifetime))
}

// TruncateConnectionID mocks base method
func (_m *MockParamsNegotiator) TruncateConnectionID() bool {
	ret := _m.ctrl.Call(_m, "TruncateConnectionID")
	ret0, _ := ret[0].(bool)
	return ret0
}

// TruncateConnectionID indicates an expected call of TruncateConnectionID
func (_mr *MockParamsNegotiatorMockRecorder) TruncateConnectionID() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "TruncateConnectionID", reflect.TypeOf((*MockParamsNegotiator)(nil).TruncateConnectionID))
}
