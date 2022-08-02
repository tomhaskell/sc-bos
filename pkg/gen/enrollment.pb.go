// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.2
// source: enrollment.proto

package gen

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

// Enrollment is metadata associated with the enrollment of a target node with a management node.
//
//The enrollment binds the target node to the management node's public key infrastructure.
//A given target node can have at most one Enrollment at a time, so an Enrollment does not need an identifier.
//
//The Enrollment Connection is the gRPC connection from the management node to the target node, opened for the purpose of
//calling CreateEnrollment.
type Enrollment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The Smart Core name that should be adopted by the target node as its root name when it accepts this enrollment.
	TargetName string `protobuf:"bytes,1,opt,name=target_name,json=targetName,proto3" json:"target_name,omitempty"`
	// The address that the management node will use to connect to the target node, in the form "host:port"
	TargetAddress string `protobuf:"bytes,2,opt,name=target_address,json=targetAddress,proto3" json:"target_address,omitempty"`
	// The Smart Core root name of the node which will manage the target node
	ManagerName string `protobuf:"bytes,3,opt,name=manager_name,json=managerName,proto3" json:"manager_name,omitempty"`
	// The address where the management node's Smart Core gRPC server can be found, in the form "host:port".
	//
	//The host must either be a DNS name or an IP address. When the target node connects to this address using gRPC with
	//TLS, the management node MUST use a certificate signed by one of the Certificate Authorities present in root_cas,
	//and that certificate MUST contain the host as a Subject Alternative Name. This is so the target node can verify
	//the identity of the management node.
	ManagerAddress string `protobuf:"bytes,4,opt,name=manager_address,json=managerAddress,proto3" json:"manager_address,omitempty"`
	// An X.509 certificate chain issued by the management node to the target node, in DER-encoded ASN.1 in a PEM container.
	//
	//If more than once certificate is present, they should be concatenated.
	//The certificate chain MUST be in leaf-first order; the leaf certificate is the certificate issued to the target node.
	//The leaf certificate's public key MUST be the target node's public key.
	//Each certificate in the chain MUST be signed by the next certificate in the chain.
	//The final certificate in the chain MUST be signed by the one of the Certificate Authorities whose certificate is
	//present in root_cas.
	//
	//The leaf certificate's Subject Common Name SHOULD be a human-readable name for the target node.
	//The leaf certificate MUST contain target_name as a URI Subject Alternative Name in the form "smart-core:<target_name>"
	//If the enrollment connection was opened by resolving a DNS name, then the leaf certificate MUST contain that DNS name
	//as a Subject Alternative Name.
	//If the enrollment connection was opened by directly connecting to an IP address, then the leaf certificate MUST
	//contain that IP address as a Subject Alternative Name.
	Certificate []byte `protobuf:"bytes,5,opt,name=certificate,proto3" json:"certificate,omitempty"`
	// One or more X.509 certificates, in DER-encoded ASN.1 in a PEM container.
	//
	//If more than one certificate is present, they should be concatenated.
	//These are the Root Certificate Authorities for the enrollment. Each MUST be self-signed.
	//The target node SHOULD use these certificate authorities whenever it communicates with another Smart Core node,
	//to verify that the other node is also enrolled with the same manager.
	RootCas []byte `protobuf:"bytes,6,opt,name=root_cas,json=rootCas,proto3" json:"root_cas,omitempty"`
}

func (x *Enrollment) Reset() {
	*x = Enrollment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_enrollment_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Enrollment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Enrollment) ProtoMessage() {}

func (x *Enrollment) ProtoReflect() protoreflect.Message {
	mi := &file_enrollment_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Enrollment.ProtoReflect.Descriptor instead.
func (*Enrollment) Descriptor() ([]byte, []int) {
	return file_enrollment_proto_rawDescGZIP(), []int{0}
}

func (x *Enrollment) GetTargetName() string {
	if x != nil {
		return x.TargetName
	}
	return ""
}

func (x *Enrollment) GetTargetAddress() string {
	if x != nil {
		return x.TargetAddress
	}
	return ""
}

func (x *Enrollment) GetManagerName() string {
	if x != nil {
		return x.ManagerName
	}
	return ""
}

func (x *Enrollment) GetManagerAddress() string {
	if x != nil {
		return x.ManagerAddress
	}
	return ""
}

func (x *Enrollment) GetCertificate() []byte {
	if x != nil {
		return x.Certificate
	}
	return nil
}

func (x *Enrollment) GetRootCas() []byte {
	if x != nil {
		return x.RootCas
	}
	return nil
}

type GetEnrollmentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetEnrollmentRequest) Reset() {
	*x = GetEnrollmentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_enrollment_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEnrollmentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEnrollmentRequest) ProtoMessage() {}

func (x *GetEnrollmentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_enrollment_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEnrollmentRequest.ProtoReflect.Descriptor instead.
func (*GetEnrollmentRequest) Descriptor() ([]byte, []int) {
	return file_enrollment_proto_rawDescGZIP(), []int{1}
}

type CreateEnrollmentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Enrollment *Enrollment `protobuf:"bytes,1,opt,name=enrollment,proto3" json:"enrollment,omitempty"`
}

func (x *CreateEnrollmentRequest) Reset() {
	*x = CreateEnrollmentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_enrollment_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateEnrollmentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateEnrollmentRequest) ProtoMessage() {}

func (x *CreateEnrollmentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_enrollment_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateEnrollmentRequest.ProtoReflect.Descriptor instead.
func (*CreateEnrollmentRequest) Descriptor() ([]byte, []int) {
	return file_enrollment_proto_rawDescGZIP(), []int{2}
}

func (x *CreateEnrollmentRequest) GetEnrollment() *Enrollment {
	if x != nil {
		return x.Enrollment
	}
	return nil
}

type DeleteEnrollmentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteEnrollmentRequest) Reset() {
	*x = DeleteEnrollmentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_enrollment_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteEnrollmentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteEnrollmentRequest) ProtoMessage() {}

func (x *DeleteEnrollmentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_enrollment_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteEnrollmentRequest.ProtoReflect.Descriptor instead.
func (*DeleteEnrollmentRequest) Descriptor() ([]byte, []int) {
	return file_enrollment_proto_rawDescGZIP(), []int{3}
}

var File_enrollment_proto protoreflect.FileDescriptor

var file_enrollment_proto_rawDesc = []byte{
	0x0a, 0x10, 0x65, 0x6e, 0x72, 0x6f, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0c, 0x76, 0x61, 0x6e, 0x74, 0x69, 0x2e, 0x62, 0x73, 0x70, 0x2e, 0x65, 0x77,
	0x22, 0xdd, 0x01, 0x0a, 0x0a, 0x45, 0x6e, 0x72, 0x6f, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x12,
	0x1f, 0x0a, 0x0b, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x25, 0x0a, 0x0e, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x72, 0x6f, 0x6f, 0x74, 0x5f, 0x63, 0x61,
	0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x72, 0x6f, 0x6f, 0x74, 0x43, 0x61, 0x73,
	0x22, 0x16, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x45, 0x6e, 0x72, 0x6f, 0x6c, 0x6c, 0x6d, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x53, 0x0a, 0x17, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x45, 0x6e, 0x72, 0x6f, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x38, 0x0a, 0x0a, 0x65, 0x6e, 0x72, 0x6f, 0x6c, 0x6c, 0x6d, 0x65, 0x6e,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x76, 0x61, 0x6e, 0x74, 0x69, 0x2e,
	0x62, 0x73, 0x70, 0x2e, 0x65, 0x77, 0x2e, 0x45, 0x6e, 0x72, 0x6f, 0x6c, 0x6c, 0x6d, 0x65, 0x6e,
	0x74, 0x52, 0x0a, 0x65, 0x6e, 0x72, 0x6f, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x19, 0x0a,
	0x17, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x6e, 0x72, 0x6f, 0x6c, 0x6c, 0x6d, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x32, 0x8e, 0x02, 0x0a, 0x0d, 0x45, 0x6e, 0x72,
	0x6f, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x41, 0x70, 0x69, 0x12, 0x4f, 0x0a, 0x0d, 0x47, 0x65,
	0x74, 0x45, 0x6e, 0x72, 0x6f, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x22, 0x2e, 0x76, 0x61,
	0x6e, 0x74, 0x69, 0x2e, 0x62, 0x73, 0x70, 0x2e, 0x65, 0x77, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x6e,
	0x72, 0x6f, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x18, 0x2e, 0x76, 0x61, 0x6e, 0x74, 0x69, 0x2e, 0x62, 0x73, 0x70, 0x2e, 0x65, 0x77, 0x2e, 0x45,
	0x6e, 0x72, 0x6f, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x00, 0x12, 0x55, 0x0a, 0x10, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x6e, 0x72, 0x6f, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x12,
	0x25, 0x2e, 0x76, 0x61, 0x6e, 0x74, 0x69, 0x2e, 0x62, 0x73, 0x70, 0x2e, 0x65, 0x77, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x6e, 0x72, 0x6f, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x76, 0x61, 0x6e, 0x74, 0x69, 0x2e, 0x62,
	0x73, 0x70, 0x2e, 0x65, 0x77, 0x2e, 0x45, 0x6e, 0x72, 0x6f, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74,
	0x22, 0x00, 0x12, 0x55, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x6e, 0x72, 0x6f,
	0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x25, 0x2e, 0x76, 0x61, 0x6e, 0x74, 0x69, 0x2e, 0x62,
	0x73, 0x70, 0x2e, 0x65, 0x77, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x6e, 0x72, 0x6f,
	0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e,
	0x76, 0x61, 0x6e, 0x74, 0x69, 0x2e, 0x62, 0x73, 0x70, 0x2e, 0x65, 0x77, 0x2e, 0x45, 0x6e, 0x72,
	0x6f, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x00, 0x42, 0x25, 0x5a, 0x23, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x76, 0x61, 0x6e, 0x74, 0x69, 0x2d, 0x64, 0x65,
	0x76, 0x2f, 0x62, 0x73, 0x70, 0x2d, 0x65, 0x77, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x67, 0x65, 0x6e,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_enrollment_proto_rawDescOnce sync.Once
	file_enrollment_proto_rawDescData = file_enrollment_proto_rawDesc
)

func file_enrollment_proto_rawDescGZIP() []byte {
	file_enrollment_proto_rawDescOnce.Do(func() {
		file_enrollment_proto_rawDescData = protoimpl.X.CompressGZIP(file_enrollment_proto_rawDescData)
	})
	return file_enrollment_proto_rawDescData
}

var file_enrollment_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_enrollment_proto_goTypes = []interface{}{
	(*Enrollment)(nil),              // 0: vanti.bsp.ew.Enrollment
	(*GetEnrollmentRequest)(nil),    // 1: vanti.bsp.ew.GetEnrollmentRequest
	(*CreateEnrollmentRequest)(nil), // 2: vanti.bsp.ew.CreateEnrollmentRequest
	(*DeleteEnrollmentRequest)(nil), // 3: vanti.bsp.ew.DeleteEnrollmentRequest
}
var file_enrollment_proto_depIdxs = []int32{
	0, // 0: vanti.bsp.ew.CreateEnrollmentRequest.enrollment:type_name -> vanti.bsp.ew.Enrollment
	1, // 1: vanti.bsp.ew.EnrollmentApi.GetEnrollment:input_type -> vanti.bsp.ew.GetEnrollmentRequest
	2, // 2: vanti.bsp.ew.EnrollmentApi.CreateEnrollment:input_type -> vanti.bsp.ew.CreateEnrollmentRequest
	3, // 3: vanti.bsp.ew.EnrollmentApi.DeleteEnrollment:input_type -> vanti.bsp.ew.DeleteEnrollmentRequest
	0, // 4: vanti.bsp.ew.EnrollmentApi.GetEnrollment:output_type -> vanti.bsp.ew.Enrollment
	0, // 5: vanti.bsp.ew.EnrollmentApi.CreateEnrollment:output_type -> vanti.bsp.ew.Enrollment
	0, // 6: vanti.bsp.ew.EnrollmentApi.DeleteEnrollment:output_type -> vanti.bsp.ew.Enrollment
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_enrollment_proto_init() }
func file_enrollment_proto_init() {
	if File_enrollment_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_enrollment_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Enrollment); i {
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
		file_enrollment_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEnrollmentRequest); i {
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
		file_enrollment_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateEnrollmentRequest); i {
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
		file_enrollment_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteEnrollmentRequest); i {
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
			RawDescriptor: file_enrollment_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_enrollment_proto_goTypes,
		DependencyIndexes: file_enrollment_proto_depIdxs,
		MessageInfos:      file_enrollment_proto_msgTypes,
	}.Build()
	File_enrollment_proto = out.File
	file_enrollment_proto_rawDesc = nil
	file_enrollment_proto_goTypes = nil
	file_enrollment_proto_depIdxs = nil
}
