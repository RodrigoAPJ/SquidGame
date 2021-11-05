// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: pozo.proto

package grpc_pozo

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

type RequestMonto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Request string `protobuf:"bytes,1,opt,name=request,proto3" json:"request,omitempty"`
}

func (x *RequestMonto) Reset() {
	*x = RequestMonto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pozo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestMonto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestMonto) ProtoMessage() {}

func (x *RequestMonto) ProtoReflect() protoreflect.Message {
	mi := &file_pozo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestMonto.ProtoReflect.Descriptor instead.
func (*RequestMonto) Descriptor() ([]byte, []int) {
	return file_pozo_proto_rawDescGZIP(), []int{0}
}

func (x *RequestMonto) GetRequest() string {
	if x != nil {
		return x.Request
	}
	return ""
}

type Monto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MontoPozo string `protobuf:"bytes,1,opt,name=montoPozo,proto3" json:"montoPozo,omitempty"`
}

func (x *Monto) Reset() {
	*x = Monto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pozo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Monto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Monto) ProtoMessage() {}

func (x *Monto) ProtoReflect() protoreflect.Message {
	mi := &file_pozo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Monto.ProtoReflect.Descriptor instead.
func (*Monto) Descriptor() ([]byte, []int) {
	return file_pozo_proto_rawDescGZIP(), []int{1}
}

func (x *Monto) GetMontoPozo() string {
	if x != nil {
		return x.MontoPozo
	}
	return ""
}

var File_pozo_proto protoreflect.FileDescriptor

var file_pozo_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x70, 0x6f, 0x7a, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x28, 0x0a, 0x0c,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x6f, 0x6e, 0x74, 0x6f, 0x12, 0x18, 0x0a, 0x07,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x25, 0x0a, 0x05, 0x4d, 0x6f, 0x6e, 0x74, 0x6f, 0x12,
	0x1c, 0x0a, 0x09, 0x6d, 0x6f, 0x6e, 0x74, 0x6f, 0x50, 0x6f, 0x7a, 0x6f, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x6d, 0x6f, 0x6e, 0x74, 0x6f, 0x50, 0x6f, 0x7a, 0x6f, 0x32, 0x2b, 0x0a,
	0x04, 0x50, 0x6f, 0x7a, 0x6f, 0x12, 0x23, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x6e, 0x74,
	0x6f, 0x12, 0x0d, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x6f, 0x6e, 0x74, 0x6f,
	0x1a, 0x06, 0x2e, 0x4d, 0x6f, 0x6e, 0x74, 0x6f, 0x22, 0x00, 0x42, 0x0e, 0x5a, 0x0c, 0x2e, 0x2e,
	0x2f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x70, 0x6f, 0x7a, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_pozo_proto_rawDescOnce sync.Once
	file_pozo_proto_rawDescData = file_pozo_proto_rawDesc
)

func file_pozo_proto_rawDescGZIP() []byte {
	file_pozo_proto_rawDescOnce.Do(func() {
		file_pozo_proto_rawDescData = protoimpl.X.CompressGZIP(file_pozo_proto_rawDescData)
	})
	return file_pozo_proto_rawDescData
}

var file_pozo_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pozo_proto_goTypes = []interface{}{
	(*RequestMonto)(nil), // 0: RequestMonto
	(*Monto)(nil),        // 1: Monto
}
var file_pozo_proto_depIdxs = []int32{
	0, // 0: Pozo.GetMonto:input_type -> RequestMonto
	1, // 1: Pozo.GetMonto:output_type -> Monto
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pozo_proto_init() }
func file_pozo_proto_init() {
	if File_pozo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pozo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestMonto); i {
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
		file_pozo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Monto); i {
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
			RawDescriptor: file_pozo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pozo_proto_goTypes,
		DependencyIndexes: file_pozo_proto_depIdxs,
		MessageInfos:      file_pozo_proto_msgTypes,
	}.Build()
	File_pozo_proto = out.File
	file_pozo_proto_rawDesc = nil
	file_pozo_proto_goTypes = nil
	file_pozo_proto_depIdxs = nil
}