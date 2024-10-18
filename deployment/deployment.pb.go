/*
 * Licensed to the TerrAPI under one or more contributor
 * license agreements. See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership. The TerrAPI licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.2
// source: deployment/deployment.proto

package deployment

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Action int32

const (
	Action_ACTION_INIT    Action = 0
	Action_ACTION_PLAN    Action = 1
	Action_ACTION_APPLY   Action = 2
	Action_ACTION_DESTROY Action = 3
)

// Enum value maps for Action.
var (
	Action_name = map[int32]string{
		0: "ACTION_INIT",
		1: "ACTION_PLAN",
		2: "ACTION_APPLY",
		3: "ACTION_DESTROY",
	}
	Action_value = map[string]int32{
		"ACTION_INIT":    0,
		"ACTION_PLAN":    1,
		"ACTION_APPLY":   2,
		"ACTION_DESTROY": 3,
	}
)

func (x Action) Enum() *Action {
	p := new(Action)
	*p = x
	return p
}

func (x Action) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Action) Descriptor() protoreflect.EnumDescriptor {
	return file_deployment_deployment_proto_enumTypes[0].Descriptor()
}

func (Action) Type() protoreflect.EnumType {
	return &file_deployment_deployment_proto_enumTypes[0]
}

func (x Action) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Action.Descriptor instead.
func (Action) EnumDescriptor() ([]byte, []int) {
	return file_deployment_deployment_proto_rawDescGZIP(), []int{0}
}

type Status_State int32

const (
	Status_unknown   Status_State = 0
	Status_pending   Status_State = 1
	Status_running   Status_State = 2
	Status_failed    Status_State = 3
	Status_succeeded Status_State = 4
)

// Enum value maps for Status_State.
var (
	Status_State_name = map[int32]string{
		0: "unknown",
		1: "pending",
		2: "running",
		3: "failed",
		4: "succeeded",
	}
	Status_State_value = map[string]int32{
		"unknown":   0,
		"pending":   1,
		"running":   2,
		"failed":    3,
		"succeeded": 4,
	}
)

func (x Status_State) Enum() *Status_State {
	p := new(Status_State)
	*p = x
	return p
}

func (x Status_State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Status_State) Descriptor() protoreflect.EnumDescriptor {
	return file_deployment_deployment_proto_enumTypes[1].Descriptor()
}

func (Status_State) Type() protoreflect.EnumType {
	return &file_deployment_deployment_proto_enumTypes[1]
}

func (x Status_State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Status_State.Descriptor instead.
func (Status_State) EnumDescriptor() ([]byte, []int) {
	return file_deployment_deployment_proto_rawDescGZIP(), []int{5, 0}
}

type RetrieveRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *RetrieveRequest) Reset() {
	*x = RetrieveRequest{}
	mi := &file_deployment_deployment_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RetrieveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RetrieveRequest) ProtoMessage() {}

func (x *RetrieveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_deployment_deployment_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RetrieveRequest.ProtoReflect.Descriptor instead.
func (*RetrieveRequest) Descriptor() ([]byte, []int) {
	return file_deployment_deployment_proto_rawDescGZIP(), []int{0}
}

func (x *RetrieveRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type SetStatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeploymentId int32  `protobuf:"varint,1,opt,name=deployment_id,json=deploymentId,proto3" json:"deployment_id,omitempty"`
	Action       Action `protobuf:"varint,2,opt,name=action,proto3,enum=terrapi.deployment.v1.Action" json:"action,omitempty"`
}

func (x *SetStatusRequest) Reset() {
	*x = SetStatusRequest{}
	mi := &file_deployment_deployment_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SetStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetStatusRequest) ProtoMessage() {}

func (x *SetStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_deployment_deployment_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetStatusRequest.ProtoReflect.Descriptor instead.
func (*SetStatusRequest) Descriptor() ([]byte, []int) {
	return file_deployment_deployment_proto_rawDescGZIP(), []int{1}
}

func (x *SetStatusRequest) GetDeploymentId() int32 {
	if x != nil {
		return x.DeploymentId
	}
	return 0
}

func (x *SetStatusRequest) GetAction() Action {
	if x != nil {
		return x.Action
	}
	return Action_ACTION_INIT
}

type Module struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Address  string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Path     string `protobuf:"bytes,3,opt,name=path,proto3" json:"path,omitempty"`
	Username string `protobuf:"bytes,4,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,5,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *Module) Reset() {
	*x = Module{}
	mi := &file_deployment_deployment_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Module) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Module) ProtoMessage() {}

func (x *Module) ProtoReflect() protoreflect.Message {
	mi := &file_deployment_deployment_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Module.ProtoReflect.Descriptor instead.
func (*Module) Descriptor() ([]byte, []int) {
	return file_deployment_deployment_proto_rawDescGZIP(), []int{2}
}

func (x *Module) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Module) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Module) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *Module) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *Module) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Action    Action             `protobuf:"varint,1,opt,name=action,proto3,enum=terrapi.deployment.v1.Action" json:"action,omitempty"`
	Variables []*RequestVariable `protobuf:"bytes,2,rep,name=variables,proto3" json:"variables,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	mi := &file_deployment_deployment_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_deployment_deployment_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_deployment_deployment_proto_rawDescGZIP(), []int{3}
}

func (x *Request) GetAction() Action {
	if x != nil {
		return x.Action
	}
	return Action_ACTION_INIT
}

func (x *Request) GetVariables() []*RequestVariable {
	if x != nil {
		return x.Variables
	}
	return nil
}

type RequestVariable struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Value  string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Secret bool   `protobuf:"varint,3,opt,name=secret,proto3" json:"secret,omitempty"`
}

func (x *RequestVariable) Reset() {
	*x = RequestVariable{}
	mi := &file_deployment_deployment_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RequestVariable) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestVariable) ProtoMessage() {}

func (x *RequestVariable) ProtoReflect() protoreflect.Message {
	mi := &file_deployment_deployment_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestVariable.ProtoReflect.Descriptor instead.
func (*RequestVariable) Descriptor() ([]byte, []int) {
	return file_deployment_deployment_proto_rawDescGZIP(), []int{4}
}

func (x *RequestVariable) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RequestVariable) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *RequestVariable) GetSecret() bool {
	if x != nil {
		return x.Secret
	}
	return false
}

type Status struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	State     Status_State `protobuf:"varint,1,opt,name=state,proto3,enum=terrapi.deployment.v1.Status_State" json:"state,omitempty"`
	Message   string       `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Timestamp int32        `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *Status) Reset() {
	*x = Status{}
	mi := &file_deployment_deployment_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Status) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Status) ProtoMessage() {}

func (x *Status) ProtoReflect() protoreflect.Message {
	mi := &file_deployment_deployment_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Status.ProtoReflect.Descriptor instead.
func (*Status) Descriptor() ([]byte, []int) {
	return file_deployment_deployment_proto_rawDescGZIP(), []int{5}
}

func (x *Status) GetState() Status_State {
	if x != nil {
		return x.State
	}
	return Status_unknown
}

func (x *Status) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Status) GetTimestamp() int32 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

type SetStatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Module  *Module  `protobuf:"bytes,1,opt,name=module,proto3" json:"module,omitempty"`
	Request *Request `protobuf:"bytes,2,opt,name=request,proto3" json:"request,omitempty"`
	Status  *Status  `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *SetStatusResponse) Reset() {
	*x = SetStatusResponse{}
	mi := &file_deployment_deployment_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SetStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetStatusResponse) ProtoMessage() {}

func (x *SetStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_deployment_deployment_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetStatusResponse.ProtoReflect.Descriptor instead.
func (*SetStatusResponse) Descriptor() ([]byte, []int) {
	return file_deployment_deployment_proto_rawDescGZIP(), []int{6}
}

func (x *SetStatusResponse) GetModule() *Module {
	if x != nil {
		return x.Module
	}
	return nil
}

func (x *SetStatusResponse) GetRequest() *Request {
	if x != nil {
		return x.Request
	}
	return nil
}

func (x *SetStatusResponse) GetStatus() *Status {
	if x != nil {
		return x.Status
	}
	return nil
}

var File_deployment_deployment_proto protoreflect.FileDescriptor

var file_deployment_deployment_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x64, 0x65, 0x70,
	0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x74,
	0x65, 0x72, 0x72, 0x61, 0x70, 0x69, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x2e, 0x76, 0x31, 0x22, 0x21, 0x0a, 0x0f, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x6e, 0x0a, 0x10, 0x53, 0x65, 0x74, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x64,
	0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0c, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64,
	0x12, 0x35, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x1d, 0x2e, 0x74, 0x65, 0x72, 0x72, 0x61, 0x70, 0x69, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x82, 0x01, 0x0a, 0x06, 0x4d, 0x6f, 0x64, 0x75,
	0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x70, 0x61, 0x74, 0x68, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x86, 0x01, 0x0a,
	0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x35, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x74, 0x65, 0x72, 0x72, 0x61,
	0x70, 0x69, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31,
	0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x44, 0x0a, 0x09, 0x76, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x26, 0x2e, 0x74, 0x65, 0x72, 0x72, 0x61, 0x70, 0x69, 0x2e, 0x64, 0x65, 0x70,
	0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x56, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x09, 0x76, 0x61, 0x72, 0x69,
	0x61, 0x62, 0x6c, 0x65, 0x73, 0x22, 0x53, 0x0a, 0x0f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x56, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x06, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x22, 0xc6, 0x01, 0x0a, 0x06, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x39, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x23, 0x2e, 0x74, 0x65, 0x72, 0x72, 0x61, 0x70, 0x69, 0x2e, 0x64,
	0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x49, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x12, 0x0b, 0x0a, 0x07, 0x75, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x10, 0x00, 0x12, 0x0b,
	0x0a, 0x07, 0x70, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x72,
	0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x66, 0x61, 0x69, 0x6c,
	0x65, 0x64, 0x10, 0x03, 0x12, 0x0d, 0x0a, 0x09, 0x73, 0x75, 0x63, 0x63, 0x65, 0x65, 0x64, 0x65,
	0x64, 0x10, 0x04, 0x22, 0xbb, 0x01, 0x0a, 0x11, 0x53, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x06, 0x6d, 0x6f, 0x64,
	0x75, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x74, 0x65, 0x72, 0x72,
	0x61, 0x70, 0x69, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76,
	0x31, 0x2e, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x06, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65,
	0x12, 0x38, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1e, 0x2e, 0x74, 0x65, 0x72, 0x72, 0x61, 0x70, 0x69, 0x2e, 0x64, 0x65, 0x70, 0x6c,
	0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x52, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x35, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x74, 0x65, 0x72,
	0x72, 0x61, 0x70, 0x69, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x2a, 0x50, 0x0a, 0x06, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0f, 0x0a, 0x0b, 0x41,
	0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x49, 0x4e, 0x49, 0x54, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b,
	0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x50, 0x4c, 0x41, 0x4e, 0x10, 0x01, 0x12, 0x10, 0x0a,
	0x0c, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x41, 0x50, 0x50, 0x4c, 0x59, 0x10, 0x02, 0x12,
	0x12, 0x0a, 0x0e, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x44, 0x45, 0x53, 0x54, 0x52, 0x4f,
	0x59, 0x10, 0x03, 0x32, 0xd0, 0x01, 0x0a, 0x11, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x59, 0x0a, 0x03, 0x47, 0x65, 0x74,
	0x12, 0x26, 0x2e, 0x74, 0x65, 0x72, 0x72, 0x61, 0x70, 0x69, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x74, 0x65, 0x72, 0x72, 0x61,
	0x70, 0x69, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31,
	0x2e, 0x53, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x60, 0x0a, 0x09, 0x53, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x27, 0x2e, 0x74, 0x65, 0x72, 0x72, 0x61, 0x70, 0x69, 0x2e, 0x64, 0x65, 0x70, 0x6c,
	0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x74, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x74, 0x65, 0x72,
	0x72, 0x61, 0x70, 0x69, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x65, 0x72, 0x72, 0x61, 0x70, 0x69, 0x2d, 0x73, 0x6f, 0x6c,
	0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x64,
	0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_deployment_deployment_proto_rawDescOnce sync.Once
	file_deployment_deployment_proto_rawDescData = file_deployment_deployment_proto_rawDesc
)

func file_deployment_deployment_proto_rawDescGZIP() []byte {
	file_deployment_deployment_proto_rawDescOnce.Do(func() {
		file_deployment_deployment_proto_rawDescData = protoimpl.X.CompressGZIP(file_deployment_deployment_proto_rawDescData)
	})
	return file_deployment_deployment_proto_rawDescData
}

var file_deployment_deployment_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_deployment_deployment_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_deployment_deployment_proto_goTypes = []any{
	(Action)(0),               // 0: terrapi.deployment.v1.Action
	(Status_State)(0),         // 1: terrapi.deployment.v1.Status.State
	(*RetrieveRequest)(nil),   // 2: terrapi.deployment.v1.RetrieveRequest
	(*SetStatusRequest)(nil),  // 3: terrapi.deployment.v1.SetStatusRequest
	(*Module)(nil),            // 4: terrapi.deployment.v1.Module
	(*Request)(nil),           // 5: terrapi.deployment.v1.Request
	(*RequestVariable)(nil),   // 6: terrapi.deployment.v1.RequestVariable
	(*Status)(nil),            // 7: terrapi.deployment.v1.Status
	(*SetStatusResponse)(nil), // 8: terrapi.deployment.v1.SetStatusResponse
}
var file_deployment_deployment_proto_depIdxs = []int32{
	0, // 0: terrapi.deployment.v1.SetStatusRequest.action:type_name -> terrapi.deployment.v1.Action
	0, // 1: terrapi.deployment.v1.Request.action:type_name -> terrapi.deployment.v1.Action
	6, // 2: terrapi.deployment.v1.Request.variables:type_name -> terrapi.deployment.v1.RequestVariable
	1, // 3: terrapi.deployment.v1.Status.state:type_name -> terrapi.deployment.v1.Status.State
	4, // 4: terrapi.deployment.v1.SetStatusResponse.module:type_name -> terrapi.deployment.v1.Module
	5, // 5: terrapi.deployment.v1.SetStatusResponse.request:type_name -> terrapi.deployment.v1.Request
	7, // 6: terrapi.deployment.v1.SetStatusResponse.status:type_name -> terrapi.deployment.v1.Status
	2, // 7: terrapi.deployment.v1.DeploymentService.Get:input_type -> terrapi.deployment.v1.RetrieveRequest
	3, // 8: terrapi.deployment.v1.DeploymentService.SetStatus:input_type -> terrapi.deployment.v1.SetStatusRequest
	8, // 9: terrapi.deployment.v1.DeploymentService.Get:output_type -> terrapi.deployment.v1.SetStatusResponse
	8, // 10: terrapi.deployment.v1.DeploymentService.SetStatus:output_type -> terrapi.deployment.v1.SetStatusResponse
	9, // [9:11] is the sub-list for method output_type
	7, // [7:9] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_deployment_deployment_proto_init() }
func file_deployment_deployment_proto_init() {
	if File_deployment_deployment_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_deployment_deployment_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_deployment_deployment_proto_goTypes,
		DependencyIndexes: file_deployment_deployment_proto_depIdxs,
		EnumInfos:         file_deployment_deployment_proto_enumTypes,
		MessageInfos:      file_deployment_deployment_proto_msgTypes,
	}.Build()
	File_deployment_deployment_proto = out.File
	file_deployment_deployment_proto_rawDesc = nil
	file_deployment_deployment_proto_goTypes = nil
	file_deployment_deployment_proto_depIdxs = nil
}
