// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.2
// source: grpc_services.proto

package codeless_grpc

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

type IdDesignRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *IdDesignRequest) Reset() {
	*x = IdDesignRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_services_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IdDesignRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IdDesignRequest) ProtoMessage() {}

func (x *IdDesignRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_services_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IdDesignRequest.ProtoReflect.Descriptor instead.
func (*IdDesignRequest) Descriptor() ([]byte, []int) {
	return file_grpc_services_proto_rawDescGZIP(), []int{0}
}

func (x *IdDesignRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ErrorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error   int32  `protobuf:"varint,1,opt,name=error,proto3" json:"error,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *ErrorResponse) Reset() {
	*x = ErrorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_services_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ErrorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ErrorResponse) ProtoMessage() {}

func (x *ErrorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_services_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ErrorResponse.ProtoReflect.Descriptor instead.
func (*ErrorResponse) Descriptor() ([]byte, []int) {
	return file_grpc_services_proto_rawDescGZIP(), []int{1}
}

func (x *ErrorResponse) GetError() int32 {
	if x != nil {
		return x.Error
	}
	return 0
}

func (x *ErrorResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_grpc_services_proto protoreflect.FileDescriptor

var file_grpc_services_proto_rawDesc = []byte{
	0x0a, 0x13, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x63, 0x6f, 0x64, 0x65, 0x6c, 0x65, 0x73, 0x73, 0x5f,
	0x67, 0x72, 0x70, 0x63, 0x22, 0x21, 0x0a, 0x0f, 0x49, 0x64, 0x44, 0x65, 0x73, 0x69, 0x67, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x3f, 0x0a, 0x0d, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0xb7, 0x01, 0x0a, 0x15, 0x54, 0x65, 0x6c,
	0x65, 0x67, 0x72, 0x61, 0x6d, 0x44, 0x65, 0x73, 0x69, 0x67, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x4e, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x65, 0x73, 0x69,
	0x67, 0x6e, 0x12, 0x1e, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x6c, 0x65, 0x73, 0x73, 0x5f, 0x67, 0x72,
	0x70, 0x63, 0x2e, 0x49, 0x64, 0x44, 0x65, 0x73, 0x69, 0x67, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x6c, 0x65, 0x73, 0x73, 0x5f, 0x67, 0x72,
	0x70, 0x63, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x4e, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x65, 0x73, 0x69,
	0x67, 0x6e, 0x12, 0x1e, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x6c, 0x65, 0x73, 0x73, 0x5f, 0x67, 0x72,
	0x70, 0x63, 0x2e, 0x49, 0x64, 0x44, 0x65, 0x73, 0x69, 0x67, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x6c, 0x65, 0x73, 0x73, 0x5f, 0x67, 0x72,
	0x70, 0x63, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0x14, 0x5a, 0x12, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x6c,
	0x65, 0x73, 0x73, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_grpc_services_proto_rawDescOnce sync.Once
	file_grpc_services_proto_rawDescData = file_grpc_services_proto_rawDesc
)

func file_grpc_services_proto_rawDescGZIP() []byte {
	file_grpc_services_proto_rawDescOnce.Do(func() {
		file_grpc_services_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpc_services_proto_rawDescData)
	})
	return file_grpc_services_proto_rawDescData
}

var file_grpc_services_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_grpc_services_proto_goTypes = []any{
	(*IdDesignRequest)(nil), // 0: codeless_grpc.IdDesignRequest
	(*ErrorResponse)(nil),   // 1: codeless_grpc.ErrorResponse
}
var file_grpc_services_proto_depIdxs = []int32{
	0, // 0: codeless_grpc.TelegramDesignService.CreateDesign:input_type -> codeless_grpc.IdDesignRequest
	0, // 1: codeless_grpc.TelegramDesignService.DeleteDesign:input_type -> codeless_grpc.IdDesignRequest
	1, // 2: codeless_grpc.TelegramDesignService.CreateDesign:output_type -> codeless_grpc.ErrorResponse
	1, // 3: codeless_grpc.TelegramDesignService.DeleteDesign:output_type -> codeless_grpc.ErrorResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_grpc_services_proto_init() }
func file_grpc_services_proto_init() {
	if File_grpc_services_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grpc_services_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*IdDesignRequest); i {
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
		file_grpc_services_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*ErrorResponse); i {
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
			RawDescriptor: file_grpc_services_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_grpc_services_proto_goTypes,
		DependencyIndexes: file_grpc_services_proto_depIdxs,
		MessageInfos:      file_grpc_services_proto_msgTypes,
	}.Build()
	File_grpc_services_proto = out.File
	file_grpc_services_proto_rawDesc = nil
	file_grpc_services_proto_goTypes = nil
	file_grpc_services_proto_depIdxs = nil
}