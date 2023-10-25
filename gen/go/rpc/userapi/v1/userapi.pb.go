// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: rpc/userapi/v1/userapi.proto

package userapiv1

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type Bid struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TxHash      string `protobuf:"bytes,1,opt,name=tx_hash,json=txHash,proto3" json:"tx_hash,omitempty"`
	Amount      int64  `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
	BlockNumber int64  `protobuf:"varint,3,opt,name=block_number,json=blockNumber,proto3" json:"block_number,omitempty"`
}

func (x *Bid) Reset() {
	*x = Bid{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_userapi_v1_userapi_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Bid) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Bid) ProtoMessage() {}

func (x *Bid) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_userapi_v1_userapi_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Bid.ProtoReflect.Descriptor instead.
func (*Bid) Descriptor() ([]byte, []int) {
	return file_rpc_userapi_v1_userapi_proto_rawDescGZIP(), []int{0}
}

func (x *Bid) GetTxHash() string {
	if x != nil {
		return x.TxHash
	}
	return ""
}

func (x *Bid) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *Bid) GetBlockNumber() int64 {
	if x != nil {
		return x.BlockNumber
	}
	return 0
}

type PreConfirmation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TxHash                   string `protobuf:"bytes,1,opt,name=tx_hash,json=txHash,proto3" json:"tx_hash,omitempty"`
	Amount                   int64  `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
	BlockNumber              int64  `protobuf:"varint,3,opt,name=block_number,json=blockNumber,proto3" json:"block_number,omitempty"`
	BidDigest                string `protobuf:"bytes,4,opt,name=bid_digest,json=bidDigest,proto3" json:"bid_digest,omitempty"`
	BidSignature             string `protobuf:"bytes,5,opt,name=bid_signature,json=bidSignature,proto3" json:"bid_signature,omitempty"`
	PreConfirmationDigest    string `protobuf:"bytes,6,opt,name=pre_confirmation_digest,json=preConfirmationDigest,proto3" json:"pre_confirmation_digest,omitempty"`
	PreConfirmationSignature string `protobuf:"bytes,7,opt,name=pre_confirmation_signature,json=preConfirmationSignature,proto3" json:"pre_confirmation_signature,omitempty"`
}

func (x *PreConfirmation) Reset() {
	*x = PreConfirmation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_userapi_v1_userapi_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PreConfirmation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PreConfirmation) ProtoMessage() {}

func (x *PreConfirmation) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_userapi_v1_userapi_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PreConfirmation.ProtoReflect.Descriptor instead.
func (*PreConfirmation) Descriptor() ([]byte, []int) {
	return file_rpc_userapi_v1_userapi_proto_rawDescGZIP(), []int{1}
}

func (x *PreConfirmation) GetTxHash() string {
	if x != nil {
		return x.TxHash
	}
	return ""
}

func (x *PreConfirmation) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *PreConfirmation) GetBlockNumber() int64 {
	if x != nil {
		return x.BlockNumber
	}
	return 0
}

func (x *PreConfirmation) GetBidDigest() string {
	if x != nil {
		return x.BidDigest
	}
	return ""
}

func (x *PreConfirmation) GetBidSignature() string {
	if x != nil {
		return x.BidSignature
	}
	return ""
}

func (x *PreConfirmation) GetPreConfirmationDigest() string {
	if x != nil {
		return x.PreConfirmationDigest
	}
	return ""
}

func (x *PreConfirmation) GetPreConfirmationSignature() string {
	if x != nil {
		return x.PreConfirmationSignature
	}
	return ""
}

var File_rpc_userapi_v1_userapi_proto protoreflect.FileDescriptor

var file_rpc_userapi_v1_userapi_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x72, 0x70, 0x63, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31,
	0x2f, 0x75, 0x73, 0x65, 0x72, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11,
	0x72, 0x70, 0x63, 0x2e, 0x73, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x31, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70,
	0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xf7, 0x04, 0x0a, 0x03, 0x42, 0x69, 0x64, 0x12, 0x8f, 0x01, 0x0a, 0x07, 0x74, 0x78, 0x5f, 0x68,
	0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x76, 0x92, 0x41, 0x73, 0x32, 0x5f,
	0x48, 0x65, 0x78, 0x20, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x20, 0x65, 0x6e, 0x63, 0x6f, 0x64,
	0x69, 0x6e, 0x67, 0x20, 0x6f, 0x66, 0x20, 0x74, 0x68, 0x65, 0x20, 0x68, 0x61, 0x73, 0x68, 0x20,
	0x6f, 0x66, 0x20, 0x74, 0x68, 0x65, 0x20, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x20, 0x74, 0x68, 0x61, 0x74, 0x20, 0x74, 0x68, 0x65, 0x20, 0x75, 0x73, 0x65, 0x72,
	0x20, 0x77, 0x61, 0x6e, 0x74, 0x73, 0x20, 0x74, 0x6f, 0x20, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64,
	0x65, 0x20, 0x69, 0x6e, 0x20, 0x74, 0x68, 0x65, 0x20, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x8a,
	0x01, 0x0f, 0x5b, 0x61, 0x2d, 0x66, 0x41, 0x2d, 0x46, 0x30, 0x2d, 0x39, 0x5d, 0x7b, 0x34, 0x30,
	0x7d, 0x52, 0x06, 0x74, 0x78, 0x48, 0x61, 0x73, 0x68, 0x12, 0x86, 0x01, 0x0a, 0x06, 0x61, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x42, 0x6e, 0x92, 0x41, 0x6b, 0x32,
	0x69, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x20, 0x6f, 0x66, 0x20, 0x45, 0x54, 0x48, 0x20, 0x74,
	0x68, 0x61, 0x74, 0x20, 0x74, 0x68, 0x65, 0x20, 0x75, 0x73, 0x65, 0x72, 0x20, 0x69, 0x73, 0x20,
	0x77, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x20, 0x74, 0x6f, 0x20, 0x70, 0x61, 0x79, 0x20, 0x74,
	0x6f, 0x20, 0x74, 0x68, 0x65, 0x20, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x20, 0x66,
	0x6f, 0x72, 0x20, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x69, 0x6e, 0x67, 0x20, 0x74, 0x68, 0x65,
	0x20, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x69, 0x6e, 0x20,
	0x74, 0x68, 0x65, 0x20, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x6b, 0x0a, 0x0c, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x6e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x42, 0x48, 0x92, 0x41, 0x45, 0x32, 0x43, 0x4d,
	0x61, 0x78, 0x20, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x20, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x20,
	0x74, 0x68, 0x61, 0x74, 0x20, 0x74, 0x68, 0x65, 0x20, 0x75, 0x73, 0x65, 0x72, 0x20, 0x77, 0x61,
	0x6e, 0x74, 0x73, 0x20, 0x74, 0x6f, 0x20, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x20, 0x74,
	0x68, 0x65, 0x20, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x69,
	0x6e, 0x2e, 0x52, 0x0b, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x3a,
	0xe7, 0x01, 0x92, 0x41, 0xe3, 0x01, 0x0a, 0x6d, 0x2a, 0x0b, 0x42, 0x69, 0x64, 0x20, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x3c, 0x55, 0x6e, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x20,
	0x62, 0x69, 0x64, 0x20, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x20, 0x66, 0x72, 0x6f, 0x6d,
	0x20, 0x75, 0x73, 0x65, 0x72, 0x73, 0x20, 0x74, 0x6f, 0x20, 0x74, 0x68, 0x65, 0x20, 0x75, 0x73,
	0x65, 0x72, 0x20, 0x6d, 0x65, 0x76, 0x2d, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x20, 0x6e, 0x6f,
	0x64, 0x65, 0x2e, 0xd2, 0x01, 0x07, 0x74, 0x78, 0x5f, 0x68, 0x61, 0x73, 0x68, 0xd2, 0x01, 0x06,
	0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0xd2, 0x01, 0x0c, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x6e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x32, 0x72, 0x7b, 0x22, 0x74, 0x78, 0x5f, 0x68, 0x61, 0x73, 0x68,
	0x22, 0x3a, 0x20, 0x22, 0x39, 0x31, 0x61, 0x38, 0x39, 0x42, 0x36, 0x33, 0x33, 0x31, 0x39, 0x34,
	0x63, 0x30, 0x44, 0x38, 0x36, 0x43, 0x35, 0x33, 0x39, 0x41, 0x31, 0x41, 0x35, 0x42, 0x31, 0x34,
	0x44, 0x43, 0x43, 0x61, 0x63, 0x66, 0x44, 0x34, 0x37, 0x30, 0x39, 0x34, 0x22, 0x2c, 0x20, 0x22,
	0x62, 0x69, 0x64, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x3a, 0x20, 0x31, 0x30, 0x30,
	0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
	0x2c, 0x20, 0x22, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22,
	0x3a, 0x20, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x7d, 0x22, 0x96, 0x07, 0x0a, 0x0f, 0x50, 0x72,
	0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x8f, 0x01,
	0x0a, 0x07, 0x74, 0x78, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x76, 0x92, 0x41, 0x73, 0x32, 0x5f, 0x48, 0x65, 0x78, 0x20, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x20, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x20, 0x6f, 0x66, 0x20, 0x74, 0x68, 0x65,
	0x20, 0x68, 0x61, 0x73, 0x68, 0x20, 0x6f, 0x66, 0x20, 0x74, 0x68, 0x65, 0x20, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x74, 0x68, 0x61, 0x74, 0x20, 0x74, 0x68,
	0x65, 0x20, 0x75, 0x73, 0x65, 0x72, 0x20, 0x77, 0x61, 0x6e, 0x74, 0x73, 0x20, 0x74, 0x6f, 0x20,
	0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x20, 0x69, 0x6e, 0x20, 0x74, 0x68, 0x65, 0x20, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x8a, 0x01, 0x0f, 0x5b, 0x61, 0x2d, 0x66, 0x41, 0x2d, 0x46, 0x30,
	0x2d, 0x39, 0x5d, 0x7b, 0x34, 0x30, 0x7d, 0x52, 0x06, 0x74, 0x78, 0x48, 0x61, 0x73, 0x68, 0x12,
	0x86, 0x01, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x42, 0x6e, 0x92, 0x41, 0x6b, 0x32, 0x69, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x20, 0x6f, 0x66,
	0x20, 0x45, 0x54, 0x48, 0x20, 0x74, 0x68, 0x61, 0x74, 0x20, 0x74, 0x68, 0x65, 0x20, 0x75, 0x73,
	0x65, 0x72, 0x20, 0x68, 0x61, 0x73, 0x20, 0x61, 0x67, 0x72, 0x65, 0x65, 0x64, 0x20, 0x74, 0x6f,
	0x20, 0x70, 0x61, 0x79, 0x20, 0x74, 0x6f, 0x20, 0x74, 0x68, 0x65, 0x20, 0x70, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x69,
	0x6e, 0x67, 0x20, 0x74, 0x68, 0x65, 0x20, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x20, 0x69, 0x6e, 0x20, 0x74, 0x68, 0x65, 0x20, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x2e,
	0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x6b, 0x0a, 0x0c, 0x62, 0x6c, 0x6f, 0x63,
	0x6b, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x42, 0x48,
	0x92, 0x41, 0x45, 0x32, 0x43, 0x4d, 0x61, 0x78, 0x20, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x20, 0x6e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x20, 0x74, 0x68, 0x61, 0x74, 0x20, 0x74, 0x68, 0x65, 0x20, 0x75,
	0x73, 0x65, 0x72, 0x20, 0x77, 0x61, 0x6e, 0x74, 0x73, 0x20, 0x74, 0x6f, 0x20, 0x69, 0x6e, 0x63,
	0x6c, 0x75, 0x64, 0x65, 0x20, 0x74, 0x68, 0x65, 0x20, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x20, 0x69, 0x6e, 0x2e, 0x52, 0x0b, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x68, 0x0a, 0x0a, 0x62, 0x69, 0x64, 0x5f, 0x64, 0x69, 0x67,
	0x65, 0x73, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x49, 0x92, 0x41, 0x46, 0x32, 0x44,
	0x48, 0x65, 0x78, 0x20, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x20, 0x65, 0x6e, 0x63, 0x6f, 0x64,
	0x69, 0x6e, 0x67, 0x20, 0x6f, 0x66, 0x20, 0x64, 0x69, 0x67, 0x65, 0x73, 0x74, 0x20, 0x6f, 0x66,
	0x20, 0x74, 0x68, 0x65, 0x20, 0x62, 0x69, 0x64, 0x20, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x20, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x20, 0x62, 0x79, 0x20, 0x74, 0x68, 0x65, 0x20, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x52, 0x09, 0x62, 0x69, 0x64, 0x44, 0x69, 0x67, 0x65, 0x73, 0x74, 0x12,
	0x6a, 0x0a, 0x0d, 0x62, 0x69, 0x64, 0x5f, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x45, 0x92, 0x41, 0x42, 0x32, 0x40, 0x48, 0x65, 0x78,
	0x20, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x20, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67,
	0x20, 0x6f, 0x66, 0x20, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x20, 0x6f, 0x66,
	0x20, 0x74, 0x68, 0x65, 0x20, 0x75, 0x73, 0x65, 0x72, 0x20, 0x74, 0x68, 0x61, 0x74, 0x20, 0x73,
	0x65, 0x6e, 0x74, 0x20, 0x74, 0x68, 0x69, 0x73, 0x20, 0x62, 0x69, 0x64, 0x2e, 0x52, 0x0c, 0x62,
	0x69, 0x64, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x73, 0x0a, 0x17, 0x70,
	0x72, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x64, 0x69, 0x67, 0x65, 0x73, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x42, 0x3b, 0x92, 0x41,
	0x38, 0x32, 0x36, 0x48, 0x65, 0x78, 0x20, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x20, 0x65, 0x6e,
	0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x20, 0x6f, 0x66, 0x20, 0x64, 0x69, 0x67, 0x65, 0x73, 0x74,
	0x20, 0x6f, 0x66, 0x20, 0x74, 0x68, 0x65, 0x20, 0x70, 0x72, 0x65, 0x2d, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x15, 0x70, 0x72, 0x65, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x69, 0x67, 0x65, 0x73, 0x74,
	0x12, 0xaf, 0x01, 0x0a, 0x1a, 0x70, 0x72, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x42, 0x71, 0x92, 0x41, 0x6e, 0x32, 0x6c, 0x48, 0x65, 0x78, 0x20,
	0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x20, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x20,
	0x6f, 0x66, 0x20, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x20, 0x6f, 0x66, 0x20,
	0x74, 0x68, 0x65, 0x20, 0x70, 0x72, 0x65, 0x2d, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x20, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x20, 0x62, 0x79, 0x20, 0x74,
	0x68, 0x65, 0x20, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x20, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x72, 0x6d, 0x69, 0x6e, 0x67, 0x20, 0x74, 0x68, 0x69, 0x73, 0x20, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x18, 0x70, 0x72, 0x65, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x32, 0x68, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x60, 0x0a, 0x07, 0x53, 0x65,
	0x6e, 0x64, 0x42, 0x69, 0x64, 0x12, 0x16, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x73, 0x65, 0x61, 0x63,
	0x68, 0x65, 0x72, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x69, 0x64, 0x1a, 0x22, 0x2e,
	0x72, 0x70, 0x63, 0x2e, 0x73, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x31, 0x2e, 0x50, 0x72, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x3a, 0x01, 0x2a, 0x22, 0x0c, 0x2f, 0x76,
	0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x62, 0x69, 0x64, 0x30, 0x01, 0x42, 0xbd, 0x01, 0x92,
	0x41, 0x78, 0x12, 0x76, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x20, 0x41, 0x50, 0x49, 0x2a, 0x5d,
	0x0a, 0x1b, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x20, 0x53, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x20, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x20, 0x31, 0x2e, 0x31, 0x12, 0x3e, 0x68,
	0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x76, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x2f, 0x6d, 0x65, 0x76, 0x2d, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x2f, 0x62, 0x6c, 0x6f, 0x62,
	0x2f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x4c, 0x49, 0x43, 0x45, 0x4e, 0x53, 0x45, 0x32, 0x0b, 0x31,
	0x2e, 0x30, 0x2e, 0x30, 0x2d, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x5a, 0x40, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x76, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x6d, 0x65, 0x76, 0x2d, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74,
	0x2f, 0x72, 0x70, 0x63, 0x2f, 0x73, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x31, 0x3b, 0x75, 0x73, 0x65, 0x72, 0x61, 0x70, 0x69, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpc_userapi_v1_userapi_proto_rawDescOnce sync.Once
	file_rpc_userapi_v1_userapi_proto_rawDescData = file_rpc_userapi_v1_userapi_proto_rawDesc
)

func file_rpc_userapi_v1_userapi_proto_rawDescGZIP() []byte {
	file_rpc_userapi_v1_userapi_proto_rawDescOnce.Do(func() {
		file_rpc_userapi_v1_userapi_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_userapi_v1_userapi_proto_rawDescData)
	})
	return file_rpc_userapi_v1_userapi_proto_rawDescData
}

var file_rpc_userapi_v1_userapi_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_rpc_userapi_v1_userapi_proto_goTypes = []interface{}{
	(*Bid)(nil),             // 0: rpc.seacherapi.v1.Bid
	(*PreConfirmation)(nil), // 1: rpc.seacherapi.v1.PreConfirmation
}
var file_rpc_userapi_v1_userapi_proto_depIdxs = []int32{
	0, // 0: rpc.seacherapi.v1.User.SendBid:input_type -> rpc.seacherapi.v1.Bid
	1, // 1: rpc.seacherapi.v1.User.SendBid:output_type -> rpc.seacherapi.v1.PreConfirmation
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_rpc_userapi_v1_userapi_proto_init() }
func file_rpc_userapi_v1_userapi_proto_init() {
	if File_rpc_userapi_v1_userapi_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rpc_userapi_v1_userapi_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Bid); i {
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
		file_rpc_userapi_v1_userapi_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PreConfirmation); i {
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
			RawDescriptor: file_rpc_userapi_v1_userapi_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rpc_userapi_v1_userapi_proto_goTypes,
		DependencyIndexes: file_rpc_userapi_v1_userapi_proto_depIdxs,
		MessageInfos:      file_rpc_userapi_v1_userapi_proto_msgTypes,
	}.Build()
	File_rpc_userapi_v1_userapi_proto = out.File
	file_rpc_userapi_v1_userapi_proto_rawDesc = nil
	file_rpc_userapi_v1_userapi_proto_goTypes = nil
	file_rpc_userapi_v1_userapi_proto_depIdxs = nil
}