// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.1
// source: domain/msg/rpc/service/msg.proto

package service

import (
	_ "common/idl/base"
	msg "common/idl/domain/msg"
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

type SendMsgReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MsgData *msg.MsgDTO `protobuf:"bytes,1,opt,name=msgData,proto3" json:"msgData,omitempty"`
}

func (x *SendMsgReq) Reset() {
	*x = SendMsgReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_domain_msg_rpc_service_msg_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendMsgReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMsgReq) ProtoMessage() {}

func (x *SendMsgReq) ProtoReflect() protoreflect.Message {
	mi := &file_domain_msg_rpc_service_msg_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendMsgReq.ProtoReflect.Descriptor instead.
func (*SendMsgReq) Descriptor() ([]byte, []int) {
	return file_domain_msg_rpc_service_msg_proto_rawDescGZIP(), []int{0}
}

func (x *SendMsgReq) GetMsgData() *msg.MsgDTO {
	if x != nil {
		return x.MsgData
	}
	return nil
}

type SendMsgResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServerMsgID string `protobuf:"bytes,1,opt,name=serverMsgID,proto3" json:"serverMsgID,omitempty"`
	ClientMsgID string `protobuf:"bytes,2,opt,name=clientMsgID,proto3" json:"clientMsgID,omitempty"`
	SendTime    int64  `protobuf:"varint,3,opt,name=sendTime,proto3" json:"sendTime,omitempty"`
}

func (x *SendMsgResp) Reset() {
	*x = SendMsgResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_domain_msg_rpc_service_msg_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendMsgResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMsgResp) ProtoMessage() {}

func (x *SendMsgResp) ProtoReflect() protoreflect.Message {
	mi := &file_domain_msg_rpc_service_msg_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendMsgResp.ProtoReflect.Descriptor instead.
func (*SendMsgResp) Descriptor() ([]byte, []int) {
	return file_domain_msg_rpc_service_msg_proto_rawDescGZIP(), []int{1}
}

func (x *SendMsgResp) GetServerMsgID() string {
	if x != nil {
		return x.ServerMsgID
	}
	return ""
}

func (x *SendMsgResp) GetClientMsgID() string {
	if x != nil {
		return x.ClientMsgID
	}
	return ""
}

func (x *SendMsgResp) GetSendTime() int64 {
	if x != nil {
		return x.SendTime
	}
	return 0
}

var File_domain_msg_rpc_service_msg_proto protoreflect.FileDescriptor

var file_domain_msg_rpc_service_msg_proto_rawDesc = []byte{
	0x0a, 0x20, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x6d, 0x73, 0x67, 0x2f, 0x72, 0x70, 0x63,
	0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x6d, 0x73, 0x67, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x1a, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2f, 0x69, 0x64, 0x6c, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x62, 0x61, 0x73,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f,
	0x69, 0x64, 0x6c, 0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x6d, 0x73, 0x67, 0x2f, 0x6d,
	0x73, 0x67, 0x5f, 0x64, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x33, 0x0a, 0x0a,
	0x53, 0x65, 0x6e, 0x64, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x71, 0x12, 0x25, 0x0a, 0x07, 0x6d, 0x73,
	0x67, 0x44, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x6d, 0x73,
	0x67, 0x2e, 0x6d, 0x73, 0x67, 0x44, 0x54, 0x4f, 0x52, 0x07, 0x6d, 0x73, 0x67, 0x44, 0x61, 0x74,
	0x61, 0x22, 0x6d, 0x0a, 0x0b, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x73, 0x70,
	0x12, 0x20, 0x0a, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4d, 0x73, 0x67, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4d, 0x73, 0x67,
	0x49, 0x44, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4d, 0x73, 0x67, 0x49,
	0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4d,
	0x73, 0x67, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x73, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65,
	0x32, 0x3b, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x34, 0x0a, 0x07, 0x53, 0x65, 0x6e, 0x64, 0x4d,
	0x73, 0x67, 0x12, 0x13, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x65, 0x6e,
	0x64, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x71, 0x1a, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x73, 0x70, 0x42, 0x18, 0x5a,
	0x16, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x6d, 0x73, 0x67, 0x2f, 0x72, 0x70, 0x63, 0x2f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_domain_msg_rpc_service_msg_proto_rawDescOnce sync.Once
	file_domain_msg_rpc_service_msg_proto_rawDescData = file_domain_msg_rpc_service_msg_proto_rawDesc
)

func file_domain_msg_rpc_service_msg_proto_rawDescGZIP() []byte {
	file_domain_msg_rpc_service_msg_proto_rawDescOnce.Do(func() {
		file_domain_msg_rpc_service_msg_proto_rawDescData = protoimpl.X.CompressGZIP(file_domain_msg_rpc_service_msg_proto_rawDescData)
	})
	return file_domain_msg_rpc_service_msg_proto_rawDescData
}

var file_domain_msg_rpc_service_msg_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_domain_msg_rpc_service_msg_proto_goTypes = []interface{}{
	(*SendMsgReq)(nil),  // 0: service.SendMsgReq
	(*SendMsgResp)(nil), // 1: service.SendMsgResp
	(*msg.MsgDTO)(nil),  // 2: msg.msgDTO
}
var file_domain_msg_rpc_service_msg_proto_depIdxs = []int32{
	2, // 0: service.SendMsgReq.msgData:type_name -> msg.msgDTO
	0, // 1: service.msg.SendMsg:input_type -> service.SendMsgReq
	1, // 2: service.msg.SendMsg:output_type -> service.SendMsgResp
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_domain_msg_rpc_service_msg_proto_init() }
func file_domain_msg_rpc_service_msg_proto_init() {
	if File_domain_msg_rpc_service_msg_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_domain_msg_rpc_service_msg_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendMsgReq); i {
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
		file_domain_msg_rpc_service_msg_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendMsgResp); i {
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
			RawDescriptor: file_domain_msg_rpc_service_msg_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_domain_msg_rpc_service_msg_proto_goTypes,
		DependencyIndexes: file_domain_msg_rpc_service_msg_proto_depIdxs,
		MessageInfos:      file_domain_msg_rpc_service_msg_proto_msgTypes,
	}.Build()
	File_domain_msg_rpc_service_msg_proto = out.File
	file_domain_msg_rpc_service_msg_proto_rawDesc = nil
	file_domain_msg_rpc_service_msg_proto_goTypes = nil
	file_domain_msg_rpc_service_msg_proto_depIdxs = nil
}
