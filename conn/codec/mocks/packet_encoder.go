// Code generated by MockGen. DO NOT EDIT.
// Source: conn/codec/packet_encoder.go

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	packet "github.com/woshihaomei/pitaya/conn/packet"
	reflect "reflect"
)

// MockPacketEncoder is a mock of PacketEncoder interface
type MockPacketEncoder struct {
	ctrl     *gomock.Controller
	recorder *MockPacketEncoderMockRecorder
}

// MockPacketEncoderMockRecorder is the mock recorder for MockPacketEncoder
type MockPacketEncoderMockRecorder struct {
	mock *MockPacketEncoder
}

// NewMockPacketEncoder creates a new mock instance
func NewMockPacketEncoder(ctrl *gomock.Controller) *MockPacketEncoder {
	mock := &MockPacketEncoder{ctrl: ctrl}
	mock.recorder = &MockPacketEncoderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPacketEncoder) EXPECT() *MockPacketEncoderMockRecorder {
	return m.recorder
}

// Encode mocks base method
func (m *MockPacketEncoder) Encode(typ packet.Type, data []byte) ([]byte, error) {
	ret := m.ctrl.Call(m, "Encode", typ, data)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Encode indicates an expected call of Encode
func (mr *MockPacketEncoderMockRecorder) Encode(typ, data interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Encode", reflect.TypeOf((*MockPacketEncoder)(nil).Encode), typ, data)
}
