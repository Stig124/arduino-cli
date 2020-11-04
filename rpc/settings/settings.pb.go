// This file is part of arduino-cli.
//
// Copyright 2020 ARDUINO SA (http://www.arduino.cc/)
//
// This software is released under the GNU General Public License version 3,
// which covers the main part of arduino-cli.
// The terms of this license can be found at:
// https://www.gnu.org/licenses/gpl-3.0.en.html
//
// You can be released from the requirements of the above licenses by purchasing
// a commercial license. Buying such a license is mandatory if you want to
// modify or otherwise use the software for commercial activities involving the
// Arduino software without disclosing the source code of your own applications.
// To purchase a commercial license, send an email to license@arduino.cc.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: settings/settings.proto

package settings

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type RawData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The settings, in JSON format.
	JsonData string `protobuf:"bytes,1,opt,name=jsonData,proto3" json:"jsonData,omitempty"`
}

func (x *RawData) Reset() {
	*x = RawData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_settings_settings_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RawData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RawData) ProtoMessage() {}

func (x *RawData) ProtoReflect() protoreflect.Message {
	mi := &file_settings_settings_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RawData.ProtoReflect.Descriptor instead.
func (*RawData) Descriptor() ([]byte, []int) {
	return file_settings_settings_proto_rawDescGZIP(), []int{0}
}

func (x *RawData) GetJsonData() string {
	if x != nil {
		return x.JsonData
	}
	return ""
}

type Value struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The key of the setting.
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// The setting, in JSON format.
	JsonData string `protobuf:"bytes,2,opt,name=jsonData,proto3" json:"jsonData,omitempty"`
}

func (x *Value) Reset() {
	*x = Value{}
	if protoimpl.UnsafeEnabled {
		mi := &file_settings_settings_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Value) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Value) ProtoMessage() {}

func (x *Value) ProtoReflect() protoreflect.Message {
	mi := &file_settings_settings_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Value.ProtoReflect.Descriptor instead.
func (*Value) Descriptor() ([]byte, []int) {
	return file_settings_settings_proto_rawDescGZIP(), []int{1}
}

func (x *Value) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Value) GetJsonData() string {
	if x != nil {
		return x.JsonData
	}
	return ""
}

type GetAllRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetAllRequest) Reset() {
	*x = GetAllRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_settings_settings_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllRequest) ProtoMessage() {}

func (x *GetAllRequest) ProtoReflect() protoreflect.Message {
	mi := &file_settings_settings_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllRequest.ProtoReflect.Descriptor instead.
func (*GetAllRequest) Descriptor() ([]byte, []int) {
	return file_settings_settings_proto_rawDescGZIP(), []int{2}
}

type GetValueRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The key of the setting.
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *GetValueRequest) Reset() {
	*x = GetValueRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_settings_settings_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetValueRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetValueRequest) ProtoMessage() {}

func (x *GetValueRequest) ProtoReflect() protoreflect.Message {
	mi := &file_settings_settings_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetValueRequest.ProtoReflect.Descriptor instead.
func (*GetValueRequest) Descriptor() ([]byte, []int) {
	return file_settings_settings_proto_rawDescGZIP(), []int{3}
}

func (x *GetValueRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type MergeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MergeResponse) Reset() {
	*x = MergeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_settings_settings_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MergeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MergeResponse) ProtoMessage() {}

func (x *MergeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_settings_settings_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MergeResponse.ProtoReflect.Descriptor instead.
func (*MergeResponse) Descriptor() ([]byte, []int) {
	return file_settings_settings_proto_rawDescGZIP(), []int{4}
}

type SetValueResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SetValueResponse) Reset() {
	*x = SetValueResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_settings_settings_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetValueResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetValueResponse) ProtoMessage() {}

func (x *SetValueResponse) ProtoReflect() protoreflect.Message {
	mi := &file_settings_settings_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetValueResponse.ProtoReflect.Descriptor instead.
func (*SetValueResponse) Descriptor() ([]byte, []int) {
	return file_settings_settings_proto_rawDescGZIP(), []int{5}
}

var File_settings_settings_proto protoreflect.FileDescriptor

var file_settings_settings_proto_rawDesc = []byte{
	0x0a, 0x17, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2f, 0x73, 0x65, 0x74, 0x74, 0x69,
	0x6e, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x63, 0x63, 0x2e, 0x61, 0x72,
	0x64, 0x75, 0x69, 0x6e, 0x6f, 0x2e, 0x63, 0x6c, 0x69, 0x2e, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e,
	0x67, 0x73, 0x22, 0x25, 0x0a, 0x07, 0x52, 0x61, 0x77, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1a, 0x0a,
	0x08, 0x6a, 0x73, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x6a, 0x73, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x22, 0x35, 0x0a, 0x05, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x6a, 0x73, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6a, 0x73, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61,
	0x22, 0x0f, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x22, 0x23, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x0f, 0x0a, 0x0d, 0x4d, 0x65, 0x72, 0x67, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x12, 0x0a, 0x10, 0x53, 0x65, 0x74, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xde, 0x02, 0x0a, 0x08,
	0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x52, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x41,
	0x6c, 0x6c, 0x12, 0x26, 0x2e, 0x63, 0x63, 0x2e, 0x61, 0x72, 0x64, 0x75, 0x69, 0x6e, 0x6f, 0x2e,
	0x63, 0x6c, 0x69, 0x2e, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x47, 0x65, 0x74,
	0x41, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x63, 0x63, 0x2e,
	0x61, 0x72, 0x64, 0x75, 0x69, 0x6e, 0x6f, 0x2e, 0x63, 0x6c, 0x69, 0x2e, 0x73, 0x65, 0x74, 0x74,
	0x69, 0x6e, 0x67, 0x73, 0x2e, 0x52, 0x61, 0x77, 0x44, 0x61, 0x74, 0x61, 0x12, 0x51, 0x0a, 0x05,
	0x4d, 0x65, 0x72, 0x67, 0x65, 0x12, 0x20, 0x2e, 0x63, 0x63, 0x2e, 0x61, 0x72, 0x64, 0x75, 0x69,
	0x6e, 0x6f, 0x2e, 0x63, 0x6c, 0x69, 0x2e, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e,
	0x52, 0x61, 0x77, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x26, 0x2e, 0x63, 0x63, 0x2e, 0x61, 0x72, 0x64,
	0x75, 0x69, 0x6e, 0x6f, 0x2e, 0x63, 0x6c, 0x69, 0x2e, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67,
	0x73, 0x2e, 0x4d, 0x65, 0x72, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x54, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x28, 0x2e, 0x63, 0x63,
	0x2e, 0x61, 0x72, 0x64, 0x75, 0x69, 0x6e, 0x6f, 0x2e, 0x63, 0x6c, 0x69, 0x2e, 0x73, 0x65, 0x74,
	0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x63, 0x63, 0x2e, 0x61, 0x72, 0x64, 0x75, 0x69,
	0x6e, 0x6f, 0x2e, 0x63, 0x6c, 0x69, 0x2e, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x55, 0x0a, 0x08, 0x53, 0x65, 0x74, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x12, 0x1e, 0x2e, 0x63, 0x63, 0x2e, 0x61, 0x72, 0x64, 0x75, 0x69, 0x6e, 0x6f, 0x2e, 0x63,
	0x6c, 0x69, 0x2e, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x1a, 0x29, 0x2e, 0x63, 0x63, 0x2e, 0x61, 0x72, 0x64, 0x75, 0x69, 0x6e, 0x6f, 0x2e, 0x63,
	0x6c, 0x69, 0x2e, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x53, 0x65, 0x74, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x2d, 0x5a, 0x2b,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x72, 0x64, 0x75, 0x69,
	0x6e, 0x6f, 0x2f, 0x61, 0x72, 0x64, 0x75, 0x69, 0x6e, 0x6f, 0x2d, 0x63, 0x6c, 0x69, 0x2f, 0x72,
	0x70, 0x63, 0x2f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_settings_settings_proto_rawDescOnce sync.Once
	file_settings_settings_proto_rawDescData = file_settings_settings_proto_rawDesc
)

func file_settings_settings_proto_rawDescGZIP() []byte {
	file_settings_settings_proto_rawDescOnce.Do(func() {
		file_settings_settings_proto_rawDescData = protoimpl.X.CompressGZIP(file_settings_settings_proto_rawDescData)
	})
	return file_settings_settings_proto_rawDescData
}

var file_settings_settings_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_settings_settings_proto_goTypes = []interface{}{
	(*RawData)(nil),          // 0: cc.arduino.cli.settings.RawData
	(*Value)(nil),            // 1: cc.arduino.cli.settings.Value
	(*GetAllRequest)(nil),    // 2: cc.arduino.cli.settings.GetAllRequest
	(*GetValueRequest)(nil),  // 3: cc.arduino.cli.settings.GetValueRequest
	(*MergeResponse)(nil),    // 4: cc.arduino.cli.settings.MergeResponse
	(*SetValueResponse)(nil), // 5: cc.arduino.cli.settings.SetValueResponse
}
var file_settings_settings_proto_depIdxs = []int32{
	2, // 0: cc.arduino.cli.settings.Settings.GetAll:input_type -> cc.arduino.cli.settings.GetAllRequest
	0, // 1: cc.arduino.cli.settings.Settings.Merge:input_type -> cc.arduino.cli.settings.RawData
	3, // 2: cc.arduino.cli.settings.Settings.GetValue:input_type -> cc.arduino.cli.settings.GetValueRequest
	1, // 3: cc.arduino.cli.settings.Settings.SetValue:input_type -> cc.arduino.cli.settings.Value
	0, // 4: cc.arduino.cli.settings.Settings.GetAll:output_type -> cc.arduino.cli.settings.RawData
	4, // 5: cc.arduino.cli.settings.Settings.Merge:output_type -> cc.arduino.cli.settings.MergeResponse
	1, // 6: cc.arduino.cli.settings.Settings.GetValue:output_type -> cc.arduino.cli.settings.Value
	5, // 7: cc.arduino.cli.settings.Settings.SetValue:output_type -> cc.arduino.cli.settings.SetValueResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_settings_settings_proto_init() }
func file_settings_settings_proto_init() {
	if File_settings_settings_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_settings_settings_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RawData); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_settings_settings_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Value); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_settings_settings_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_settings_settings_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetValueRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_settings_settings_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MergeResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_settings_settings_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetValueResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_settings_settings_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_settings_settings_proto_goTypes,
		DependencyIndexes: file_settings_settings_proto_depIdxs,
		MessageInfos:      file_settings_settings_proto_msgTypes,
	}.Build()
	File_settings_settings_proto = out.File
	file_settings_settings_proto_rawDesc = nil
	file_settings_settings_proto_goTypes = nil
	file_settings_settings_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SettingsClient is the client API for Settings service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SettingsClient interface {
	// List all the settings.
	GetAll(ctx context.Context, in *GetAllRequest, opts ...grpc.CallOption) (*RawData, error)
	// Set multiple settings values at once.
	Merge(ctx context.Context, in *RawData, opts ...grpc.CallOption) (*MergeResponse, error)
	// Get the value of a specific setting.
	GetValue(ctx context.Context, in *GetValueRequest, opts ...grpc.CallOption) (*Value, error)
	// Set the value of a specific setting.
	SetValue(ctx context.Context, in *Value, opts ...grpc.CallOption) (*SetValueResponse, error)
}

type settingsClient struct {
	cc grpc.ClientConnInterface
}

func NewSettingsClient(cc grpc.ClientConnInterface) SettingsClient {
	return &settingsClient{cc}
}

func (c *settingsClient) GetAll(ctx context.Context, in *GetAllRequest, opts ...grpc.CallOption) (*RawData, error) {
	out := new(RawData)
	err := c.cc.Invoke(ctx, "/cc.arduino.cli.settings.Settings/GetAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *settingsClient) Merge(ctx context.Context, in *RawData, opts ...grpc.CallOption) (*MergeResponse, error) {
	out := new(MergeResponse)
	err := c.cc.Invoke(ctx, "/cc.arduino.cli.settings.Settings/Merge", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *settingsClient) GetValue(ctx context.Context, in *GetValueRequest, opts ...grpc.CallOption) (*Value, error) {
	out := new(Value)
	err := c.cc.Invoke(ctx, "/cc.arduino.cli.settings.Settings/GetValue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *settingsClient) SetValue(ctx context.Context, in *Value, opts ...grpc.CallOption) (*SetValueResponse, error) {
	out := new(SetValueResponse)
	err := c.cc.Invoke(ctx, "/cc.arduino.cli.settings.Settings/SetValue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SettingsServer is the server API for Settings service.
type SettingsServer interface {
	// List all the settings.
	GetAll(context.Context, *GetAllRequest) (*RawData, error)
	// Set multiple settings values at once.
	Merge(context.Context, *RawData) (*MergeResponse, error)
	// Get the value of a specific setting.
	GetValue(context.Context, *GetValueRequest) (*Value, error)
	// Set the value of a specific setting.
	SetValue(context.Context, *Value) (*SetValueResponse, error)
}

// UnimplementedSettingsServer can be embedded to have forward compatible implementations.
type UnimplementedSettingsServer struct {
}

func (*UnimplementedSettingsServer) GetAll(context.Context, *GetAllRequest) (*RawData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (*UnimplementedSettingsServer) Merge(context.Context, *RawData) (*MergeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Merge not implemented")
}
func (*UnimplementedSettingsServer) GetValue(context.Context, *GetValueRequest) (*Value, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetValue not implemented")
}
func (*UnimplementedSettingsServer) SetValue(context.Context, *Value) (*SetValueResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetValue not implemented")
}

func RegisterSettingsServer(s *grpc.Server, srv SettingsServer) {
	s.RegisterService(&_Settings_serviceDesc, srv)
}

func _Settings_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SettingsServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cc.arduino.cli.settings.Settings/GetAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SettingsServer).GetAll(ctx, req.(*GetAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Settings_Merge_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RawData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SettingsServer).Merge(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cc.arduino.cli.settings.Settings/Merge",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SettingsServer).Merge(ctx, req.(*RawData))
	}
	return interceptor(ctx, in, info, handler)
}

func _Settings_GetValue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetValueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SettingsServer).GetValue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cc.arduino.cli.settings.Settings/GetValue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SettingsServer).GetValue(ctx, req.(*GetValueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Settings_SetValue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Value)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SettingsServer).SetValue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cc.arduino.cli.settings.Settings/SetValue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SettingsServer).SetValue(ctx, req.(*Value))
	}
	return interceptor(ctx, in, info, handler)
}

var _Settings_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cc.arduino.cli.settings.Settings",
	HandlerType: (*SettingsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAll",
			Handler:    _Settings_GetAll_Handler,
		},
		{
			MethodName: "Merge",
			Handler:    _Settings_Merge_Handler,
		},
		{
			MethodName: "GetValue",
			Handler:    _Settings_GetValue_Handler,
		},
		{
			MethodName: "SetValue",
			Handler:    _Settings_SetValue_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "settings/settings.proto",
}
