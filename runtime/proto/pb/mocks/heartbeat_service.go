// Licensed to the Apache Software Foundation (ASF) under one or more
// contributor license agreements.  See the NOTICE file distributed with
// this work for additional information regarding copyright ownership.
// The ASF licenses this file to You under the Apache License, Version 2.0
// (the "License"); you may not use this file except in compliance with
// the License.  You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/apache/incubator-eventmesh/eventmesh-server-go/runtime/proto/pb (interfaces: HeartbeatServiceServer)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	pb "github.com/apache/incubator-eventmesh/eventmesh-server-go/runtime/proto/pb"
	gomock "github.com/golang/mock/gomock"
)

// MockHeartbeatServiceServer is a mock of HeartbeatServiceServer interface.
type MockHeartbeatServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockHeartbeatServiceServerMockRecorder
}

// MockHeartbeatServiceServerMockRecorder is the mock recorder for MockHeartbeatServiceServer.
type MockHeartbeatServiceServerMockRecorder struct {
	mock *MockHeartbeatServiceServer
}

// NewMockHeartbeatServiceServer creates a new mock instance.
func NewMockHeartbeatServiceServer(ctrl *gomock.Controller) *MockHeartbeatServiceServer {
	mock := &MockHeartbeatServiceServer{ctrl: ctrl}
	mock.recorder = &MockHeartbeatServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHeartbeatServiceServer) EXPECT() *MockHeartbeatServiceServerMockRecorder {
	return m.recorder
}

// Heartbeat mocks base method.
func (m *MockHeartbeatServiceServer) Heartbeat(arg0 context.Context, arg1 *pb.Heartbeat) (*pb.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Heartbeat", arg0, arg1)
	ret0, _ := ret[0].(*pb.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Heartbeat indicates an expected call of Heartbeat.
func (mr *MockHeartbeatServiceServerMockRecorder) Heartbeat(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Heartbeat", reflect.TypeOf((*MockHeartbeatServiceServer)(nil).Heartbeat), arg0, arg1)
}

// mustEmbedUnimplementedHeartbeatServiceServer mocks base method.
func (m *MockHeartbeatServiceServer) mustEmbedUnimplementedHeartbeatServiceServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedHeartbeatServiceServer")
}

// mustEmbedUnimplementedHeartbeatServiceServer indicates an expected call of mustEmbedUnimplementedHeartbeatServiceServer.
func (mr *MockHeartbeatServiceServerMockRecorder) mustEmbedUnimplementedHeartbeatServiceServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedHeartbeatServiceServer", reflect.TypeOf((*MockHeartbeatServiceServer)(nil).mustEmbedUnimplementedHeartbeatServiceServer))
}
