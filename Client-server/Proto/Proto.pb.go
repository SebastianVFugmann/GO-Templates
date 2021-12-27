//Command: protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative Proto/Proto.proto
//Run from Client-server folder. If you are in the proto folder, remove the first "Proto/"

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.0
// source: Client-server/Proto/Proto.proto

package Proto

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

// The request message containing the user's message
type Object struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Object) Reset() {
	*x = Object{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Client_server_Proto_Proto_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Object) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Object) ProtoMessage() {}

func (x *Object) ProtoReflect() protoreflect.Message {
	mi := &file_Client_server_Proto_Proto_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Object.ProtoReflect.Descriptor instead.
func (*Object) Descriptor() ([]byte, []int) {
	return file_Client_server_Proto_Proto_proto_rawDescGZIP(), []int{0}
}

type ServerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HttpStatusCode int32 `protobuf:"varint,1,opt,name=HttpStatusCode,proto3" json:"HttpStatusCode,omitempty"`
	Time           int32 `protobuf:"varint,2,opt,name=time,proto3" json:"time,omitempty"`
}

func (x *ServerResponse) Reset() {
	*x = ServerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Client_server_Proto_Proto_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerResponse) ProtoMessage() {}

func (x *ServerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_Client_server_Proto_Proto_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerResponse.ProtoReflect.Descriptor instead.
func (*ServerResponse) Descriptor() ([]byte, []int) {
	return file_Client_server_Proto_Proto_proto_rawDescGZIP(), []int{1}
}

func (x *ServerResponse) GetHttpStatusCode() int32 {
	if x != nil {
		return x.HttpStatusCode
	}
	return 0
}

func (x *ServerResponse) GetTime() int32 {
	if x != nil {
		return x.Time
	}
	return 0
}

type JoinRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Time int32 `protobuf:"varint,2,opt,name=time,proto3" json:"time,omitempty"`
}

func (x *JoinRequest) Reset() {
	*x = JoinRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Client_server_Proto_Proto_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JoinRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinRequest) ProtoMessage() {}

func (x *JoinRequest) ProtoReflect() protoreflect.Message {
	mi := &file_Client_server_Proto_Proto_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinRequest.ProtoReflect.Descriptor instead.
func (*JoinRequest) Descriptor() ([]byte, []int) {
	return file_Client_server_Proto_Proto_proto_rawDescGZIP(), []int{2}
}

func (x *JoinRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *JoinRequest) GetTime() int32 {
	if x != nil {
		return x.Time
	}
	return 0
}

type LeaveRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Time int32 `protobuf:"varint,2,opt,name=time,proto3" json:"time,omitempty"`
}

func (x *LeaveRequest) Reset() {
	*x = LeaveRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Client_server_Proto_Proto_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LeaveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LeaveRequest) ProtoMessage() {}

func (x *LeaveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_Client_server_Proto_Proto_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LeaveRequest.ProtoReflect.Descriptor instead.
func (*LeaveRequest) Descriptor() ([]byte, []int) {
	return file_Client_server_Proto_Proto_proto_rawDescGZIP(), []int{3}
}

func (x *LeaveRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *LeaveRequest) GetTime() int32 {
	if x != nil {
		return x.Time
	}
	return 0
}

type NodeId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *NodeId) Reset() {
	*x = NodeId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Client_server_Proto_Proto_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeId) ProtoMessage() {}

func (x *NodeId) ProtoReflect() protoreflect.Message {
	mi := &file_Client_server_Proto_Proto_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeId.ProtoReflect.Descriptor instead.
func (*NodeId) Descriptor() ([]byte, []int) {
	return file_Client_server_Proto_Proto_proto_rawDescGZIP(), []int{4}
}

func (x *NodeId) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_Client_server_Proto_Proto_proto protoreflect.FileDescriptor

var file_Client_server_Proto_Proto_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x05, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x08, 0x0a, 0x06, 0x4f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x22, 0x4c, 0x0a, 0x0e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x48, 0x74, 0x74, 0x70, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x48, 0x74,
	0x74, 0x70, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65,
	0x22, 0x31, 0x0a, 0x0b, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x74,
	0x69, 0x6d, 0x65, 0x22, 0x32, 0x0a, 0x0c, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x22, 0x18, 0x0a, 0x06, 0x4e, 0x6f, 0x64, 0x65, 0x49,
	0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69,
	0x64, 0x32, 0xcb, 0x01, 0x0a, 0x07, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2d, 0x0a,
	0x04, 0x4a, 0x6f, 0x69, 0x6e, 0x12, 0x12, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4a, 0x6f,
	0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x00, 0x30, 0x01, 0x12, 0x30, 0x0a, 0x06,
	0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x0d, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x1a, 0x15, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x35,
	0x0a, 0x05, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x12, 0x13, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x4c, 0x65, 0x61, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x28, 0x0a, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49,
	0x44, 0x12, 0x0d, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x64,
	0x1a, 0x0d, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x42,
	0x3e, 0x5a, 0x3c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x53, 0x65,
	0x62, 0x61, 0x73, 0x74, 0x69, 0x61, 0x6e, 0x56, 0x46, 0x75, 0x67, 0x6d, 0x61, 0x6e, 0x6e, 0x2f,
	0x47, 0x4f, 0x2d, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2f, 0x43, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Client_server_Proto_Proto_proto_rawDescOnce sync.Once
	file_Client_server_Proto_Proto_proto_rawDescData = file_Client_server_Proto_Proto_proto_rawDesc
)

func file_Client_server_Proto_Proto_proto_rawDescGZIP() []byte {
	file_Client_server_Proto_Proto_proto_rawDescOnce.Do(func() {
		file_Client_server_Proto_Proto_proto_rawDescData = protoimpl.X.CompressGZIP(file_Client_server_Proto_Proto_proto_rawDescData)
	})
	return file_Client_server_Proto_Proto_proto_rawDescData
}

var file_Client_server_Proto_Proto_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_Client_server_Proto_Proto_proto_goTypes = []interface{}{
	(*Object)(nil),         // 0: Proto.Object
	(*ServerResponse)(nil), // 1: Proto.ServerResponse
	(*JoinRequest)(nil),    // 2: Proto.JoinRequest
	(*LeaveRequest)(nil),   // 3: Proto.LeaveRequest
	(*NodeId)(nil),         // 4: Proto.NodeId
}
var file_Client_server_Proto_Proto_proto_depIdxs = []int32{
	2, // 0: Proto.Service.Join:input_type -> Proto.JoinRequest
	0, // 1: Proto.Service.Method:input_type -> Proto.Object
	3, // 2: Proto.Service.Leave:input_type -> Proto.LeaveRequest
	4, // 3: Proto.Service.createID:input_type -> Proto.NodeId
	0, // 4: Proto.Service.Join:output_type -> Proto.Object
	1, // 5: Proto.Service.Method:output_type -> Proto.ServerResponse
	1, // 6: Proto.Service.Leave:output_type -> Proto.ServerResponse
	4, // 7: Proto.Service.createID:output_type -> Proto.NodeId
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_Client_server_Proto_Proto_proto_init() }
func file_Client_server_Proto_Proto_proto_init() {
	if File_Client_server_Proto_Proto_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Client_server_Proto_Proto_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Object); i {
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
		file_Client_server_Proto_Proto_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerResponse); i {
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
		file_Client_server_Proto_Proto_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JoinRequest); i {
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
		file_Client_server_Proto_Proto_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LeaveRequest); i {
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
		file_Client_server_Proto_Proto_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeId); i {
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
			RawDescriptor: file_Client_server_Proto_Proto_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_Client_server_Proto_Proto_proto_goTypes,
		DependencyIndexes: file_Client_server_Proto_Proto_proto_depIdxs,
		MessageInfos:      file_Client_server_Proto_Proto_proto_msgTypes,
	}.Build()
	File_Client_server_Proto_Proto_proto = out.File
	file_Client_server_Proto_Proto_proto_rawDesc = nil
	file_Client_server_Proto_Proto_proto_goTypes = nil
	file_Client_server_Proto_Proto_proto_depIdxs = nil
}
