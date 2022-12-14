// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.20.1
// source: pkg/driver/axiomxa/rpc/axiomxa.proto

package rpc

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

type SaveQRCredentialRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	QrBody  []byte `protobuf:"bytes,2,opt,name=qr_body,json=qrBody,proto3" json:"qr_body,omitempty"`
	Account string `protobuf:"bytes,3,opt,name=account,proto3" json:"account,omitempty"`
}

func (x *SaveQRCredentialRequest) Reset() {
	*x = SaveQRCredentialRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_driver_axiomxa_rpc_axiomxa_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaveQRCredentialRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveQRCredentialRequest) ProtoMessage() {}

func (x *SaveQRCredentialRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_driver_axiomxa_rpc_axiomxa_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveQRCredentialRequest.ProtoReflect.Descriptor instead.
func (*SaveQRCredentialRequest) Descriptor() ([]byte, []int) {
	return file_pkg_driver_axiomxa_rpc_axiomxa_proto_rawDescGZIP(), []int{0}
}

func (x *SaveQRCredentialRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SaveQRCredentialRequest) GetQrBody() []byte {
	if x != nil {
		return x.QrBody
	}
	return nil
}

func (x *SaveQRCredentialRequest) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

type SaveQRCredentialResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SaveQRCredentialResponse) Reset() {
	*x = SaveQRCredentialResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_driver_axiomxa_rpc_axiomxa_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaveQRCredentialResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveQRCredentialResponse) ProtoMessage() {}

func (x *SaveQRCredentialResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_driver_axiomxa_rpc_axiomxa_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveQRCredentialResponse.ProtoReflect.Descriptor instead.
func (*SaveQRCredentialResponse) Descriptor() ([]byte, []int) {
	return file_pkg_driver_axiomxa_rpc_axiomxa_proto_rawDescGZIP(), []int{1}
}

var File_pkg_driver_axiomxa_rpc_axiomxa_proto protoreflect.FileDescriptor

var file_pkg_driver_axiomxa_rpc_axiomxa_proto_rawDesc = []byte{
	0x0a, 0x24, 0x70, 0x6b, 0x67, 0x2f, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x2f, 0x61, 0x78, 0x69,
	0x6f, 0x6d, 0x78, 0x61, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x61, 0x78, 0x69, 0x6f, 0x6d, 0x78, 0x61,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x76, 0x61, 0x6e, 0x74, 0x69, 0x2e, 0x62, 0x73,
	0x70, 0x2e, 0x65, 0x77, 0x2e, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x2e, 0x61, 0x78, 0x69, 0x6f,
	0x6d, 0x78, 0x61, 0x22, 0x60, 0x0a, 0x17, 0x53, 0x61, 0x76, 0x65, 0x51, 0x52, 0x43, 0x72, 0x65,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x71, 0x72, 0x5f, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x06, 0x71, 0x72, 0x42, 0x6f, 0x64, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x1a, 0x0a, 0x18, 0x53, 0x61, 0x76, 0x65, 0x51, 0x52, 0x43,
	0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x32, 0x97, 0x01, 0x0a, 0x14, 0x41, 0x78, 0x69, 0x6f, 0x6d, 0x58, 0x61, 0x44, 0x72, 0x69,
	0x76, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x7f, 0x0a, 0x10, 0x53, 0x61,
	0x76, 0x65, 0x51, 0x52, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x12, 0x34,
	0x2e, 0x76, 0x61, 0x6e, 0x74, 0x69, 0x2e, 0x62, 0x73, 0x70, 0x2e, 0x65, 0x77, 0x2e, 0x64, 0x72,
	0x69, 0x76, 0x65, 0x72, 0x2e, 0x61, 0x78, 0x69, 0x6f, 0x6d, 0x78, 0x61, 0x2e, 0x53, 0x61, 0x76,
	0x65, 0x51, 0x52, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x35, 0x2e, 0x76, 0x61, 0x6e, 0x74, 0x69, 0x2e, 0x62, 0x73, 0x70,
	0x2e, 0x65, 0x77, 0x2e, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x2e, 0x61, 0x78, 0x69, 0x6f, 0x6d,
	0x78, 0x61, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x51, 0x52, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x34, 0x5a, 0x32, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x76, 0x61, 0x6e, 0x74, 0x69, 0x2d,
	0x64, 0x65, 0x76, 0x2f, 0x73, 0x63, 0x2d, 0x62, 0x6f, 0x73, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x64,
	0x72, 0x69, 0x76, 0x65, 0x72, 0x2f, 0x61, 0x78, 0x69, 0x6f, 0x6d, 0x78, 0x61, 0x2f, 0x72, 0x70,
	0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_driver_axiomxa_rpc_axiomxa_proto_rawDescOnce sync.Once
	file_pkg_driver_axiomxa_rpc_axiomxa_proto_rawDescData = file_pkg_driver_axiomxa_rpc_axiomxa_proto_rawDesc
)

func file_pkg_driver_axiomxa_rpc_axiomxa_proto_rawDescGZIP() []byte {
	file_pkg_driver_axiomxa_rpc_axiomxa_proto_rawDescOnce.Do(func() {
		file_pkg_driver_axiomxa_rpc_axiomxa_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_driver_axiomxa_rpc_axiomxa_proto_rawDescData)
	})
	return file_pkg_driver_axiomxa_rpc_axiomxa_proto_rawDescData
}

var file_pkg_driver_axiomxa_rpc_axiomxa_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pkg_driver_axiomxa_rpc_axiomxa_proto_goTypes = []interface{}{
	(*SaveQRCredentialRequest)(nil),  // 0: vanti.bsp.ew.driver.axiomxa.SaveQRCredentialRequest
	(*SaveQRCredentialResponse)(nil), // 1: vanti.bsp.ew.driver.axiomxa.SaveQRCredentialResponse
}
var file_pkg_driver_axiomxa_rpc_axiomxa_proto_depIdxs = []int32{
	0, // 0: vanti.bsp.ew.driver.axiomxa.AxiomXaDriverService.SaveQRCredential:input_type -> vanti.bsp.ew.driver.axiomxa.SaveQRCredentialRequest
	1, // 1: vanti.bsp.ew.driver.axiomxa.AxiomXaDriverService.SaveQRCredential:output_type -> vanti.bsp.ew.driver.axiomxa.SaveQRCredentialResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pkg_driver_axiomxa_rpc_axiomxa_proto_init() }
func file_pkg_driver_axiomxa_rpc_axiomxa_proto_init() {
	if File_pkg_driver_axiomxa_rpc_axiomxa_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_driver_axiomxa_rpc_axiomxa_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaveQRCredentialRequest); i {
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
		file_pkg_driver_axiomxa_rpc_axiomxa_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaveQRCredentialResponse); i {
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
			RawDescriptor: file_pkg_driver_axiomxa_rpc_axiomxa_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_driver_axiomxa_rpc_axiomxa_proto_goTypes,
		DependencyIndexes: file_pkg_driver_axiomxa_rpc_axiomxa_proto_depIdxs,
		MessageInfos:      file_pkg_driver_axiomxa_rpc_axiomxa_proto_msgTypes,
	}.Build()
	File_pkg_driver_axiomxa_rpc_axiomxa_proto = out.File
	file_pkg_driver_axiomxa_rpc_axiomxa_proto_rawDesc = nil
	file_pkg_driver_axiomxa_rpc_axiomxa_proto_goTypes = nil
	file_pkg_driver_axiomxa_rpc_axiomxa_proto_depIdxs = nil
}
