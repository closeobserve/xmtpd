// Signature is a generic structure for public key signatures.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        (unknown)
// source: message_contents/signature.proto

package message_contents

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

// Signature represents a generalized public key signature,
// defined as a union to support cryptographic algorithm agility.
type Signature struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to Union:
	//
	//	*Signature_EcdsaCompact
	//	*Signature_WalletEcdsaCompact
	Union         isSignature_Union `protobuf_oneof:"union"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Signature) Reset() {
	*x = Signature{}
	mi := &file_message_contents_signature_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Signature) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Signature) ProtoMessage() {}

func (x *Signature) ProtoReflect() protoreflect.Message {
	mi := &file_message_contents_signature_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Signature.ProtoReflect.Descriptor instead.
func (*Signature) Descriptor() ([]byte, []int) {
	return file_message_contents_signature_proto_rawDescGZIP(), []int{0}
}

func (x *Signature) GetUnion() isSignature_Union {
	if x != nil {
		return x.Union
	}
	return nil
}

func (x *Signature) GetEcdsaCompact() *Signature_ECDSACompact {
	if x != nil {
		if x, ok := x.Union.(*Signature_EcdsaCompact); ok {
			return x.EcdsaCompact
		}
	}
	return nil
}

func (x *Signature) GetWalletEcdsaCompact() *Signature_WalletECDSACompact {
	if x != nil {
		if x, ok := x.Union.(*Signature_WalletEcdsaCompact); ok {
			return x.WalletEcdsaCompact
		}
	}
	return nil
}

type isSignature_Union interface {
	isSignature_Union()
}

type Signature_EcdsaCompact struct {
	EcdsaCompact *Signature_ECDSACompact `protobuf:"bytes,1,opt,name=ecdsa_compact,json=ecdsaCompact,proto3,oneof"`
}

type Signature_WalletEcdsaCompact struct {
	WalletEcdsaCompact *Signature_WalletECDSACompact `protobuf:"bytes,2,opt,name=wallet_ecdsa_compact,json=walletEcdsaCompact,proto3,oneof"`
}

func (*Signature_EcdsaCompact) isSignature_Union() {}

func (*Signature_WalletEcdsaCompact) isSignature_Union() {}

// ECDSA signature bytes and the recovery bit
type Signature_ECDSACompact struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Bytes         []byte                 `protobuf:"bytes,1,opt,name=bytes,proto3" json:"bytes,omitempty"`        // compact representation [ R || S ], 64 bytes
	Recovery      uint32                 `protobuf:"varint,2,opt,name=recovery,proto3" json:"recovery,omitempty"` // recovery bit
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Signature_ECDSACompact) Reset() {
	*x = Signature_ECDSACompact{}
	mi := &file_message_contents_signature_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Signature_ECDSACompact) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Signature_ECDSACompact) ProtoMessage() {}

func (x *Signature_ECDSACompact) ProtoReflect() protoreflect.Message {
	mi := &file_message_contents_signature_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Signature_ECDSACompact.ProtoReflect.Descriptor instead.
func (*Signature_ECDSACompact) Descriptor() ([]byte, []int) {
	return file_message_contents_signature_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Signature_ECDSACompact) GetBytes() []byte {
	if x != nil {
		return x.Bytes
	}
	return nil
}

func (x *Signature_ECDSACompact) GetRecovery() uint32 {
	if x != nil {
		return x.Recovery
	}
	return 0
}

// ECDSA signature bytes and the recovery bit
// produced by xmtp-js::PublicKey.signWithWallet function, i.e.
// EIP-191 signature of a "Create Identity" message with the key embedded.
// Used to sign identity keys.
type Signature_WalletECDSACompact struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Bytes         []byte                 `protobuf:"bytes,1,opt,name=bytes,proto3" json:"bytes,omitempty"`        // compact representation [ R || S ], 64 bytes
	Recovery      uint32                 `protobuf:"varint,2,opt,name=recovery,proto3" json:"recovery,omitempty"` // recovery bit
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Signature_WalletECDSACompact) Reset() {
	*x = Signature_WalletECDSACompact{}
	mi := &file_message_contents_signature_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Signature_WalletECDSACompact) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Signature_WalletECDSACompact) ProtoMessage() {}

func (x *Signature_WalletECDSACompact) ProtoReflect() protoreflect.Message {
	mi := &file_message_contents_signature_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Signature_WalletECDSACompact.ProtoReflect.Descriptor instead.
func (*Signature_WalletECDSACompact) Descriptor() ([]byte, []int) {
	return file_message_contents_signature_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Signature_WalletECDSACompact) GetBytes() []byte {
	if x != nil {
		return x.Bytes
	}
	return nil
}

func (x *Signature_WalletECDSACompact) GetRecovery() uint32 {
	if x != nil {
		return x.Recovery
	}
	return 0
}

var File_message_contents_signature_proto protoreflect.FileDescriptor

var file_message_contents_signature_proto_rawDesc = []byte{
	0x0a, 0x20, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x73, 0x2f, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x15, 0x78, 0x6d, 0x74, 0x70, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x22, 0xdd, 0x02, 0x0a, 0x09, 0x53, 0x69,
	0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x54, 0x0a, 0x0d, 0x65, 0x63, 0x64, 0x73, 0x61,
	0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d,
	0x2e, 0x78, 0x6d, 0x74, 0x70, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x2e, 0x45, 0x43, 0x44, 0x53, 0x41, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x63, 0x74, 0x48, 0x00, 0x52,
	0x0c, 0x65, 0x63, 0x64, 0x73, 0x61, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x63, 0x74, 0x12, 0x67, 0x0a,
	0x14, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x5f, 0x65, 0x63, 0x64, 0x73, 0x61, 0x5f, 0x63, 0x6f,
	0x6d, 0x70, 0x61, 0x63, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x78, 0x6d,
	0x74, 0x70, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x73, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x57, 0x61,
	0x6c, 0x6c, 0x65, 0x74, 0x45, 0x43, 0x44, 0x53, 0x41, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x63, 0x74,
	0x48, 0x00, 0x52, 0x12, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x45, 0x63, 0x64, 0x73, 0x61, 0x43,
	0x6f, 0x6d, 0x70, 0x61, 0x63, 0x74, 0x1a, 0x40, 0x0a, 0x0c, 0x45, 0x43, 0x44, 0x53, 0x41, 0x43,
	0x6f, 0x6d, 0x70, 0x61, 0x63, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x62, 0x79, 0x74, 0x65, 0x73, 0x12, 0x1a, 0x0a, 0x08,
	0x72, 0x65, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08,
	0x72, 0x65, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x1a, 0x46, 0x0a, 0x12, 0x57, 0x61, 0x6c, 0x6c,
	0x65, 0x74, 0x45, 0x43, 0x44, 0x53, 0x41, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x63, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x62,
	0x79, 0x74, 0x65, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x72, 0x65, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79,
	0x42, 0x07, 0x0a, 0x05, 0x75, 0x6e, 0x69, 0x6f, 0x6e, 0x42, 0xce, 0x01, 0x0a, 0x19, 0x63, 0x6f,
	0x6d, 0x2e, 0x78, 0x6d, 0x74, 0x70, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x42, 0x0e, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x78, 0x6d, 0x74, 0x70, 0x2f, 0x78, 0x6d, 0x74, 0x70, 0x64,
	0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0xa2, 0x02, 0x03, 0x58, 0x4d,
	0x58, 0xaa, 0x02, 0x14, 0x58, 0x6d, 0x74, 0x70, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0xca, 0x02, 0x14, 0x58, 0x6d, 0x74, 0x70, 0x5c,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0xe2,
	0x02, 0x20, 0x58, 0x6d, 0x74, 0x70, 0x5c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0xea, 0x02, 0x15, 0x58, 0x6d, 0x74, 0x70, 0x3a, 0x3a, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_message_contents_signature_proto_rawDescOnce sync.Once
	file_message_contents_signature_proto_rawDescData = file_message_contents_signature_proto_rawDesc
)

func file_message_contents_signature_proto_rawDescGZIP() []byte {
	file_message_contents_signature_proto_rawDescOnce.Do(func() {
		file_message_contents_signature_proto_rawDescData = protoimpl.X.CompressGZIP(file_message_contents_signature_proto_rawDescData)
	})
	return file_message_contents_signature_proto_rawDescData
}

var file_message_contents_signature_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_message_contents_signature_proto_goTypes = []any{
	(*Signature)(nil),                    // 0: xmtp.message_contents.Signature
	(*Signature_ECDSACompact)(nil),       // 1: xmtp.message_contents.Signature.ECDSACompact
	(*Signature_WalletECDSACompact)(nil), // 2: xmtp.message_contents.Signature.WalletECDSACompact
}
var file_message_contents_signature_proto_depIdxs = []int32{
	1, // 0: xmtp.message_contents.Signature.ecdsa_compact:type_name -> xmtp.message_contents.Signature.ECDSACompact
	2, // 1: xmtp.message_contents.Signature.wallet_ecdsa_compact:type_name -> xmtp.message_contents.Signature.WalletECDSACompact
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_message_contents_signature_proto_init() }
func file_message_contents_signature_proto_init() {
	if File_message_contents_signature_proto != nil {
		return
	}
	file_message_contents_signature_proto_msgTypes[0].OneofWrappers = []any{
		(*Signature_EcdsaCompact)(nil),
		(*Signature_WalletEcdsaCompact)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_message_contents_signature_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_message_contents_signature_proto_goTypes,
		DependencyIndexes: file_message_contents_signature_proto_depIdxs,
		MessageInfos:      file_message_contents_signature_proto_msgTypes,
	}.Build()
	File_message_contents_signature_proto = out.File
	file_message_contents_signature_proto_rawDesc = nil
	file_message_contents_signature_proto_goTypes = nil
	file_message_contents_signature_proto_depIdxs = nil
}
