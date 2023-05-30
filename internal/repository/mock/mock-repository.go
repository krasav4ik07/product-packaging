// Code generated by MockGen. DO NOT EDIT.
// Source: repository.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	os "os"
	core "product-packaging/internal/core"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockDBHandler is a mock of DBHandler interface.
type MockDBHandler struct {
	ctrl     *gomock.Controller
	recorder *MockDBHandlerMockRecorder
}

// MockDBHandlerMockRecorder is the mock recorder for MockDBHandler.
type MockDBHandlerMockRecorder struct {
	mock *MockDBHandler
}

// NewMockDBHandler creates a new mock instance.
func NewMockDBHandler(ctrl *gomock.Controller) *MockDBHandler {
	mock := &MockDBHandler{ctrl: ctrl}
	mock.recorder = &MockDBHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDBHandler) EXPECT() *MockDBHandlerMockRecorder {
	return m.recorder
}

// GetAmountPacksAndVerifyGtin mocks base method.
func (m *MockDBHandler) GetAmountPacksAndVerifyGtin(gtin string) (int, *core.ErrorResponse) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAmountPacksAndVerifyGtin", gtin)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(*core.ErrorResponse)
	return ret0, ret1
}

// GetAmountPacksAndVerifyGtin indicates an expected call of GetAmountPacksAndVerifyGtin.
func (mr *MockDBHandlerMockRecorder) GetAmountPacksAndVerifyGtin(gtin interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAmountPacksAndVerifyGtin", reflect.TypeOf((*MockDBHandler)(nil).GetAmountPacksAndVerifyGtin), gtin)
}

// GetBoxesAndPacksByGtin mocks base method.
func (m *MockDBHandler) GetBoxesAndPacksByGtin(gtin string, csvFile *os.File) *core.ErrorResponse {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBoxesAndPacksByGtin", gtin, csvFile)
	ret0, _ := ret[0].(*core.ErrorResponse)
	return ret0
}

// GetBoxesAndPacksByGtin indicates an expected call of GetBoxesAndPacksByGtin.
func (mr *MockDBHandlerMockRecorder) GetBoxesAndPacksByGtin(gtin, csvFile interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBoxesAndPacksByGtin", reflect.TypeOf((*MockDBHandler)(nil).GetBoxesAndPacksByGtin), gtin, csvFile)
}

// GetBoxesBySgtin mocks base method.
func (m *MockDBHandler) GetBoxesBySgtin(serialNumbers string) (map[string]*string, *core.ErrorResponse) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBoxesBySgtin", serialNumbers)
	ret0, _ := ret[0].(map[string]*string)
	ret1, _ := ret[1].(*core.ErrorResponse)
	return ret0, ret1
}

// GetBoxesBySgtin indicates an expected call of GetBoxesBySgtin.
func (mr *MockDBHandlerMockRecorder) GetBoxesBySgtin(serialNumbers interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBoxesBySgtin", reflect.TypeOf((*MockDBHandler)(nil).GetBoxesBySgtin), serialNumbers)
}

// SaveBox mocks base method.
func (m *MockDBHandler) SaveBox(box core.Box, dataPacks string) *core.ErrorResponse {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveBox", box, dataPacks)
	ret0, _ := ret[0].(*core.ErrorResponse)
	return ret0
}

// SaveBox indicates an expected call of SaveBox.
func (mr *MockDBHandlerMockRecorder) SaveBox(box, dataPacks interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveBox", reflect.TypeOf((*MockDBHandler)(nil).SaveBox), box, dataPacks)
}
