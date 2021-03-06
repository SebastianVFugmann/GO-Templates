//Command: protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative Proto/Service.proto
//Run from Active_replication folder. If you are in the proto folder, remove the first "Proto/"

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.0
// source: Proto/Service.proto

package Service

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

type IncrementRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value int32 `protobuf:"varint,1,opt,name=Value,proto3" json:"Value,omitempty"`
}

func (x *IncrementRequest) Reset() {
	*x = IncrementRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Proto_Service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IncrementRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IncrementRequest) ProtoMessage() {}

func (x *IncrementRequest) ProtoReflect() protoreflect.Message {
	mi := &file_Proto_Service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IncrementRequest.ProtoReflect.Descriptor instead.
func (*IncrementRequest) Descriptor() ([]byte, []int) {
	return file_Proto_Service_proto_rawDescGZIP(), []int{0}
}

func (x *IncrementRequest) GetValue() int32 {
	if x != nil {
		return x.Value
	}
	return 0
}

type IncrementReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success     bool  `protobuf:"varint,1,opt,name=Success,proto3" json:"Success,omitempty"`
	ValueBefore int32 `protobuf:"varint,2,opt,name=ValueBefore,proto3" json:"ValueBefore,omitempty"`
}

func (x *IncrementReply) Reset() {
	*x = IncrementReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Proto_Service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IncrementReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IncrementReply) ProtoMessage() {}

func (x *IncrementReply) ProtoReflect() protoreflect.Message {
	mi := &file_Proto_Service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IncrementReply.ProtoReflect.Descriptor instead.
func (*IncrementReply) Descriptor() ([]byte, []int) {
	return file_Proto_Service_proto_rawDescGZIP(), []int{1}
}

func (x *IncrementReply) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *IncrementReply) GetValueBefore() int32 {
	if x != nil {
		return x.ValueBefore
	}
	return 0
}

var File_Proto_Service_proto protoreflect.FileDescriptor

var file_Proto_Service_proto_rawDesc = []byte{
	0x0a, 0x13, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x28,
	0x0a, 0x10, 0x49, 0x6e, 0x63, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x4c, 0x0a, 0x0e, 0x49, 0x6e, 0x63, 0x72,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x53, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x65, 0x66,
	0x6f, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x42, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x32, 0x4c, 0x0a, 0x07, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x41, 0x0a, 0x09, 0x69, 0x6e, 0x63, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x19,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x49, 0x6e, 0x63, 0x72, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x49, 0x6e, 0x63, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x00, 0x42, 0x46, 0x5a, 0x44, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x53, 0x65, 0x62, 0x61, 0x73, 0x74, 0x69, 0x61, 0x6e, 0x56, 0x46, 0x75, 0x67,
	0x6d, 0x61, 0x6e, 0x6e, 0x2f, 0x47, 0x4f, 0x2d, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x73, 0x2f, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Proto_Service_proto_rawDescOnce sync.Once
	file_Proto_Service_proto_rawDescData = file_Proto_Service_proto_rawDesc
)

func file_Proto_Service_proto_rawDescGZIP() []byte {
	file_Proto_Service_proto_rawDescOnce.Do(func() {
		file_Proto_Service_proto_rawDescData = protoimpl.X.CompressGZIP(file_Proto_Service_proto_rawDescData)
	})
	return file_Proto_Service_proto_rawDescData
}

var file_Proto_Service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_Proto_Service_proto_goTypes = []interface{}{
	(*IncrementRequest)(nil), // 0: service.IncrementRequest
	(*IncrementReply)(nil),   // 1: service.IncrementReply
}
var file_Proto_Service_proto_depIdxs = []int32{
	0, // 0: service.Service.increment:input_type -> service.IncrementRequest
	1, // 1: service.Service.increment:output_type -> service.IncrementReply
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_Proto_Service_proto_init() }
func file_Proto_Service_proto_init() {
	if File_Proto_Service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Proto_Service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IncrementRequest); i {
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
		file_Proto_Service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IncrementReply); i {
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
			RawDescriptor: file_Proto_Service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_Proto_Service_proto_goTypes,
		DependencyIndexes: file_Proto_Service_proto_depIdxs,
		MessageInfos:      file_Proto_Service_proto_msgTypes,
	}.Build()
	File_Proto_Service_proto = out.File
	file_Proto_Service_proto_rawDesc = nil
	file_Proto_Service_proto_goTypes = nil
	file_Proto_Service_proto_depIdxs = nil
}
