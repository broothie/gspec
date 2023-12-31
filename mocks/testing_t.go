// Code generated by MockGen. DO NOT EDIT.
// Source: testing.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"
	testing "testing"

	gomock "go.uber.org/mock/gomock"
)

// MocktestingT is a mock of testingT interface.
type MocktestingT struct {
	ctrl     *gomock.Controller
	recorder *MocktestingTMockRecorder
}

// MocktestingTMockRecorder is the mock recorder for MocktestingT.
type MocktestingTMockRecorder struct {
	mock *MocktestingT
}

// NewMocktestingT creates a new mock instance.
func NewMocktestingT(ctrl *gomock.Controller) *MocktestingT {
	mock := &MocktestingT{ctrl: ctrl}
	mock.recorder = &MocktestingTMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MocktestingT) EXPECT() *MocktestingTMockRecorder {
	return m.recorder
}

// Errorf mocks base method.
func (m *MocktestingT) Errorf(format string, args ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{format}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Errorf", varargs...)
}

// Errorf indicates an expected call of Errorf.
func (mr *MocktestingTMockRecorder) Errorf(format interface{}, args ...interface{}) *testingTErrorfCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{format}, args...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Errorf", reflect.TypeOf((*MocktestingT)(nil).Errorf), varargs...)
	return &testingTErrorfCall{Call: call}
}

// testingTErrorfCall wrap *gomock.Call
type testingTErrorfCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *testingTErrorfCall) Return() *testingTErrorfCall {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *testingTErrorfCall) Do(f func(string, ...interface{})) *testingTErrorfCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *testingTErrorfCall) DoAndReturn(f func(string, ...interface{})) *testingTErrorfCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// FailNow mocks base method.
func (m *MocktestingT) FailNow() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "FailNow")
}

// FailNow indicates an expected call of FailNow.
func (mr *MocktestingTMockRecorder) FailNow() *testingTFailNowCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FailNow", reflect.TypeOf((*MocktestingT)(nil).FailNow))
	return &testingTFailNowCall{Call: call}
}

// testingTFailNowCall wrap *gomock.Call
type testingTFailNowCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *testingTFailNowCall) Return() *testingTFailNowCall {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *testingTFailNowCall) Do(f func()) *testingTFailNowCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *testingTFailNowCall) DoAndReturn(f func()) *testingTFailNowCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Helper mocks base method.
func (m *MocktestingT) Helper() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Helper")
}

// Helper indicates an expected call of Helper.
func (mr *MocktestingTMockRecorder) Helper() *testingTHelperCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Helper", reflect.TypeOf((*MocktestingT)(nil).Helper))
	return &testingTHelperCall{Call: call}
}

// testingTHelperCall wrap *gomock.Call
type testingTHelperCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *testingTHelperCall) Return() *testingTHelperCall {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *testingTHelperCall) Do(f func()) *testingTHelperCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *testingTHelperCall) DoAndReturn(f func()) *testingTHelperCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Run mocks base method.
func (m *MocktestingT) Run(name string, f func(*testing.T)) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Run", name, f)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Run indicates an expected call of Run.
func (mr *MocktestingTMockRecorder) Run(name, f interface{}) *testingTRunCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MocktestingT)(nil).Run), name, f)
	return &testingTRunCall{Call: call}
}

// testingTRunCall wrap *gomock.Call
type testingTRunCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *testingTRunCall) Return(arg0 bool) *testingTRunCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *testingTRunCall) Do(f func(string, func(*testing.T)) bool) *testingTRunCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *testingTRunCall) DoAndReturn(f func(string, func(*testing.T)) bool) *testingTRunCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
