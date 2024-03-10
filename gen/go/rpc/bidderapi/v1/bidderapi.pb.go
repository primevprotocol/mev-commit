// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: rpc/bidderapi/v1/bidderapi.proto

package bidderapiv1

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
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

type PrepayRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Amount string `protobuf:"bytes,1,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *PrepayRequest) Reset() {
	*x = PrepayRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_bidderapi_v1_bidderapi_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrepayRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrepayRequest) ProtoMessage() {}

func (x *PrepayRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_bidderapi_v1_bidderapi_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrepayRequest.ProtoReflect.Descriptor instead.
func (*PrepayRequest) Descriptor() ([]byte, []int) {
	return file_rpc_bidderapi_v1_bidderapi_proto_rawDescGZIP(), []int{0}
}

func (x *PrepayRequest) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

type PrepayResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Amount string `protobuf:"bytes,1,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *PrepayResponse) Reset() {
	*x = PrepayResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_bidderapi_v1_bidderapi_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrepayResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrepayResponse) ProtoMessage() {}

func (x *PrepayResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_bidderapi_v1_bidderapi_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrepayResponse.ProtoReflect.Descriptor instead.
func (*PrepayResponse) Descriptor() ([]byte, []int) {
	return file_rpc_bidderapi_v1_bidderapi_proto_rawDescGZIP(), []int{1}
}

func (x *PrepayResponse) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

type EmptyMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyMessage) Reset() {
	*x = EmptyMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_bidderapi_v1_bidderapi_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyMessage) ProtoMessage() {}

func (x *EmptyMessage) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_bidderapi_v1_bidderapi_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyMessage.ProtoReflect.Descriptor instead.
func (*EmptyMessage) Descriptor() ([]byte, []int) {
	return file_rpc_bidderapi_v1_bidderapi_proto_rawDescGZIP(), []int{2}
}

type Bid struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TxHashes            []string `protobuf:"bytes,1,rep,name=tx_hashes,json=txHashes,proto3" json:"tx_hashes,omitempty"`
	Amount              string   `protobuf:"bytes,2,opt,name=amount,proto3" json:"amount,omitempty"`
	BlockNumber         int64    `protobuf:"varint,3,opt,name=block_number,json=blockNumber,proto3" json:"block_number,omitempty"`
	DecayStartTimestamp uint64   `protobuf:"varint,4,opt,name=decay_start_timestamp,json=decayStartTimestamp,proto3" json:"decay_start_timestamp,omitempty"`
	DecayEndTimestamp   uint64   `protobuf:"varint,5,opt,name=decay_end_timestamp,json=decayEndTimestamp,proto3" json:"decay_end_timestamp,omitempty"`
}

func (x *Bid) Reset() {
	*x = Bid{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_bidderapi_v1_bidderapi_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Bid) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Bid) ProtoMessage() {}

func (x *Bid) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_bidderapi_v1_bidderapi_proto_msgTypes[3]
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
	return file_rpc_bidderapi_v1_bidderapi_proto_rawDescGZIP(), []int{3}
}

func (x *Bid) GetTxHashes() []string {
	if x != nil {
		return x.TxHashes
	}
	return nil
}

func (x *Bid) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

func (x *Bid) GetBlockNumber() int64 {
	if x != nil {
		return x.BlockNumber
	}
	return 0
}

func (x *Bid) GetDecayStartTimestamp() uint64 {
	if x != nil {
		return x.DecayStartTimestamp
	}
	return 0
}

func (x *Bid) GetDecayEndTimestamp() uint64 {
	if x != nil {
		return x.DecayEndTimestamp
	}
	return 0
}

type Commitment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TxHashes             []string `protobuf:"bytes,1,rep,name=tx_hashes,json=txHashes,proto3" json:"tx_hashes,omitempty"`
	BidAmount            int64    `protobuf:"varint,2,opt,name=bid_amount,json=bidAmount,proto3" json:"bid_amount,omitempty"`
	BlockNumber          int64    `protobuf:"varint,3,opt,name=block_number,json=blockNumber,proto3" json:"block_number,omitempty"`
	ReceivedBidDigest    string   `protobuf:"bytes,4,opt,name=received_bid_digest,json=receivedBidDigest,proto3" json:"received_bid_digest,omitempty"`
	ReceivedBidSignature string   `protobuf:"bytes,5,opt,name=received_bid_signature,json=receivedBidSignature,proto3" json:"received_bid_signature,omitempty"`
	CommitmentDigest     string   `protobuf:"bytes,6,opt,name=commitment_digest,json=commitmentDigest,proto3" json:"commitment_digest,omitempty"`
	CommitmentSignature  string   `protobuf:"bytes,7,opt,name=commitment_signature,json=commitmentSignature,proto3" json:"commitment_signature,omitempty"`
	ProviderAddress      string   `protobuf:"bytes,8,opt,name=provider_address,json=providerAddress,proto3" json:"provider_address,omitempty"`
	DecayStartTimestamp  uint64   `protobuf:"varint,9,opt,name=decay_start_timestamp,json=decayStartTimestamp,proto3" json:"decay_start_timestamp,omitempty"`
	DecayEndTimestamp    uint64   `protobuf:"varint,10,opt,name=decay_end_timestamp,json=decayEndTimestamp,proto3" json:"decay_end_timestamp,omitempty"`
}

func (x *Commitment) Reset() {
	*x = Commitment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_bidderapi_v1_bidderapi_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Commitment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Commitment) ProtoMessage() {}

func (x *Commitment) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_bidderapi_v1_bidderapi_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Commitment.ProtoReflect.Descriptor instead.
func (*Commitment) Descriptor() ([]byte, []int) {
	return file_rpc_bidderapi_v1_bidderapi_proto_rawDescGZIP(), []int{4}
}

func (x *Commitment) GetTxHashes() []string {
	if x != nil {
		return x.TxHashes
	}
	return nil
}

func (x *Commitment) GetBidAmount() int64 {
	if x != nil {
		return x.BidAmount
	}
	return 0
}

func (x *Commitment) GetBlockNumber() int64 {
	if x != nil {
		return x.BlockNumber
	}
	return 0
}

func (x *Commitment) GetReceivedBidDigest() string {
	if x != nil {
		return x.ReceivedBidDigest
	}
	return ""
}

func (x *Commitment) GetReceivedBidSignature() string {
	if x != nil {
		return x.ReceivedBidSignature
	}
	return ""
}

func (x *Commitment) GetCommitmentDigest() string {
	if x != nil {
		return x.CommitmentDigest
	}
	return ""
}

func (x *Commitment) GetCommitmentSignature() string {
	if x != nil {
		return x.CommitmentSignature
	}
	return ""
}

func (x *Commitment) GetProviderAddress() string {
	if x != nil {
		return x.ProviderAddress
	}
	return ""
}

func (x *Commitment) GetDecayStartTimestamp() uint64 {
	if x != nil {
		return x.DecayStartTimestamp
	}
	return 0
}

func (x *Commitment) GetDecayEndTimestamp() uint64 {
	if x != nil {
		return x.DecayEndTimestamp
	}
	return 0
}

var File_rpc_bidderapi_v1_bidderapi_proto protoreflect.FileDescriptor

var file_rpc_bidderapi_v1_bidderapi_proto_rawDesc = []byte{
	0x0a, 0x20, 0x72, 0x70, 0x63, 0x2f, 0x62, 0x69, 0x64, 0x64, 0x65, 0x72, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x31, 0x2f, 0x62, 0x69, 0x64, 0x64, 0x65, 0x72, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x10, 0x72, 0x70, 0x63, 0x2e, 0x62, 0x69, 0x64, 0x64, 0x65, 0x72, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x31, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e,
	0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xb3, 0x02, 0x0a, 0x0d, 0x50, 0x72, 0x65, 0x70, 0x61, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0xa5, 0x01, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x8c, 0x01, 0x92, 0x41, 0x2e, 0x32, 0x23, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x20, 0x6f, 0x66, 0x20, 0x45, 0x54, 0x48, 0x20, 0x74, 0x6f, 0x20, 0x62, 0x65, 0x20, 0x70, 0x72,
	0x65, 0x70, 0x61, 0x69, 0x64, 0x20, 0x69, 0x6e, 0x20, 0x77, 0x65, 0x69, 0x2e, 0x8a, 0x01, 0x06,
	0x5b, 0x30, 0x2d, 0x39, 0x5d, 0x2b, 0xba, 0x48, 0x58, 0xba, 0x01, 0x55, 0x0a, 0x06, 0x61, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x20, 0x6d, 0x75, 0x73,
	0x74, 0x20, 0x62, 0x65, 0x20, 0x61, 0x20, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x20, 0x69, 0x6e, 0x74,
	0x65, 0x67, 0x65, 0x72, 0x2e, 0x1a, 0x2a, 0x74, 0x68, 0x69, 0x73, 0x2e, 0x6d, 0x61, 0x74, 0x63,
	0x68, 0x65, 0x73, 0x28, 0x27, 0x5e, 0x5b, 0x30, 0x2d, 0x39, 0x5d, 0x2b, 0x24, 0x27, 0x29, 0x20,
	0x26, 0x26, 0x20, 0x75, 0x69, 0x6e, 0x74, 0x28, 0x74, 0x68, 0x69, 0x73, 0x29, 0x20, 0x3e, 0x20,
	0x30, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x3a, 0x7a, 0x92, 0x41, 0x77, 0x0a, 0x51,
	0x2a, 0x0e, 0x50, 0x72, 0x65, 0x70, 0x61, 0x79, 0x20, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x32, 0x36, 0x50, 0x72, 0x65, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x20, 0x66, 0x6f, 0x72,
	0x20, 0x62, 0x69, 0x64, 0x73, 0x20, 0x74, 0x6f, 0x20, 0x62, 0x65, 0x20, 0x69, 0x73, 0x73, 0x75,
	0x65, 0x64, 0x20, 0x62, 0x79, 0x20, 0x74, 0x68, 0x65, 0x20, 0x62, 0x69, 0x64, 0x64, 0x65, 0x72,
	0x20, 0x69, 0x6e, 0x20, 0x77, 0x65, 0x69, 0x2e, 0xd2, 0x01, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x32, 0x22, 0x7b, 0x22, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x3a, 0x20, 0x22, 0x31,
	0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
	0x30, 0x30, 0x22, 0x20, 0x7d, 0x22, 0x9e, 0x01, 0x0a, 0x0e, 0x50, 0x72, 0x65, 0x70, 0x61, 0x79,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x3a, 0x74, 0x92, 0x41, 0x71, 0x0a, 0x4b, 0x2a, 0x0f, 0x50, 0x72, 0x65, 0x70, 0x61, 0x79, 0x20,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x38, 0x47, 0x65, 0x74, 0x20, 0x70, 0x72,
	0x65, 0x70, 0x61, 0x69, 0x64, 0x20, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x61, 0x6e, 0x63, 0x65, 0x20,
	0x66, 0x6f, 0x72, 0x20, 0x62, 0x69, 0x64, 0x64, 0x65, 0x72, 0x20, 0x69, 0x6e, 0x20, 0x74, 0x68,
	0x65, 0x20, 0x62, 0x69, 0x64, 0x64, 0x65, 0x72, 0x20, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72,
	0x79, 0x2e, 0x32, 0x22, 0x7b, 0x22, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x3a, 0x20, 0x22,
	0x31, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
	0x30, 0x30, 0x30, 0x22, 0x20, 0x7d, 0x22, 0x0e, 0x0a, 0x0c, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0xa2, 0x0b, 0x0a, 0x03, 0x42, 0x69, 0x64, 0x12, 0xa3,
	0x02, 0x0a, 0x09, 0x74, 0x78, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x09, 0x42, 0x85, 0x02, 0x92, 0x41, 0x78, 0x32, 0x64, 0x48, 0x65, 0x78, 0x20, 0x73, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x20, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x20, 0x6f, 0x66,
	0x20, 0x74, 0x68, 0x65, 0x20, 0x68, 0x61, 0x73, 0x68, 0x65, 0x73, 0x20, 0x6f, 0x66, 0x20, 0x74,
	0x68, 0x65, 0x20, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x20,
	0x74, 0x68, 0x61, 0x74, 0x20, 0x74, 0x68, 0x65, 0x20, 0x62, 0x69, 0x64, 0x64, 0x65, 0x72, 0x20,
	0x77, 0x61, 0x6e, 0x74, 0x73, 0x20, 0x74, 0x6f, 0x20, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65,
	0x20, 0x69, 0x6e, 0x20, 0x74, 0x68, 0x65, 0x20, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x8a, 0x01,
	0x0f, 0x5b, 0x61, 0x2d, 0x66, 0x41, 0x2d, 0x46, 0x30, 0x2d, 0x39, 0x5d, 0x7b, 0x36, 0x34, 0x7d,
	0xba, 0x48, 0x86, 0x01, 0xba, 0x01, 0x82, 0x01, 0x0a, 0x09, 0x74, 0x78, 0x5f, 0x68, 0x61, 0x73,
	0x68, 0x65, 0x73, 0x12, 0x36, 0x74, 0x78, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x65, 0x73, 0x20, 0x6d,
	0x75, 0x73, 0x74, 0x20, 0x62, 0x65, 0x20, 0x61, 0x20, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x20, 0x61,
	0x72, 0x72, 0x61, 0x79, 0x20, 0x6f, 0x66, 0x20, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x20, 0x68, 0x61, 0x73, 0x68, 0x65, 0x73, 0x2e, 0x1a, 0x3d, 0x74, 0x68, 0x69,
	0x73, 0x2e, 0x61, 0x6c, 0x6c, 0x28, 0x72, 0x2c, 0x20, 0x72, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68,
	0x65, 0x73, 0x28, 0x27, 0x5e, 0x5b, 0x61, 0x2d, 0x66, 0x41, 0x2d, 0x46, 0x30, 0x2d, 0x39, 0x5d,
	0x7b, 0x36, 0x34, 0x7d, 0x24, 0x27, 0x29, 0x29, 0x20, 0x26, 0x26, 0x20, 0x73, 0x69, 0x7a, 0x65,
	0x28, 0x74, 0x68, 0x69, 0x73, 0x29, 0x20, 0x3e, 0x20, 0x30, 0x52, 0x08, 0x74, 0x78, 0x48, 0x61,
	0x73, 0x68, 0x65, 0x73, 0x12, 0xed, 0x01, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0xd4, 0x01, 0x92, 0x41, 0x76, 0x32, 0x6b, 0x41, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x20, 0x6f, 0x66, 0x20, 0x45, 0x54, 0x48, 0x20, 0x74, 0x68, 0x61, 0x74, 0x20,
	0x74, 0x68, 0x65, 0x20, 0x62, 0x69, 0x64, 0x64, 0x65, 0x72, 0x20, 0x69, 0x73, 0x20, 0x77, 0x69,
	0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x20, 0x74, 0x6f, 0x20, 0x70, 0x61, 0x79, 0x20, 0x74, 0x6f, 0x20,
	0x74, 0x68, 0x65, 0x20, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x20, 0x66, 0x6f, 0x72,
	0x20, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x69, 0x6e, 0x67, 0x20, 0x74, 0x68, 0x65, 0x20, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x69, 0x6e, 0x20, 0x74, 0x68,
	0x65, 0x20, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x8a, 0x01, 0x06, 0x5b, 0x30, 0x2d, 0x39, 0x5d,
	0x2b, 0xba, 0x48, 0x58, 0xba, 0x01, 0x55, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x1f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x20, 0x6d, 0x75, 0x73, 0x74, 0x20, 0x62, 0x65, 0x20,
	0x61, 0x20, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x20, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x65, 0x72, 0x2e,
	0x1a, 0x2a, 0x74, 0x68, 0x69, 0x73, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x28, 0x27,
	0x5e, 0x5b, 0x30, 0x2d, 0x39, 0x5d, 0x2b, 0x24, 0x27, 0x29, 0x20, 0x26, 0x26, 0x20, 0x75, 0x69,
	0x6e, 0x74, 0x28, 0x74, 0x68, 0x69, 0x73, 0x29, 0x20, 0x3e, 0x20, 0x30, 0x52, 0x06, 0x61, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0xb9, 0x01, 0x0a, 0x0c, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x6e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x42, 0x95, 0x01, 0x92, 0x41,
	0x47, 0x32, 0x45, 0x4d, 0x61, 0x78, 0x20, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x20, 0x6e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x20, 0x74, 0x68, 0x61, 0x74, 0x20, 0x74, 0x68, 0x65, 0x20, 0x62, 0x69, 0x64,
	0x64, 0x65, 0x72, 0x20, 0x77, 0x61, 0x6e, 0x74, 0x73, 0x20, 0x74, 0x6f, 0x20, 0x69, 0x6e, 0x63,
	0x6c, 0x75, 0x64, 0x65, 0x20, 0x74, 0x68, 0x65, 0x20, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x20, 0x69, 0x6e, 0x2e, 0xba, 0x48, 0x48, 0xba, 0x01, 0x45, 0x0a, 0x0c,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x25, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x20, 0x6d, 0x75, 0x73, 0x74, 0x20,
	0x62, 0x65, 0x20, 0x61, 0x20, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x20, 0x69, 0x6e, 0x74, 0x65, 0x67,
	0x65, 0x72, 0x2e, 0x1a, 0x0e, 0x75, 0x69, 0x6e, 0x74, 0x28, 0x74, 0x68, 0x69, 0x73, 0x29, 0x20,
	0x3e, 0x20, 0x30, 0x52, 0x0b, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x12, 0xc2, 0x01, 0x0a, 0x15, 0x64, 0x65, 0x63, 0x61, 0x79, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04,
	0x42, 0x8d, 0x01, 0x92, 0x41, 0x2d, 0x32, 0x2b, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x20, 0x61, 0x74, 0x20, 0x77, 0x68, 0x69, 0x63, 0x68, 0x20, 0x74, 0x68, 0x65, 0x20, 0x62,
	0x69, 0x64, 0x20, 0x73, 0x74, 0x61, 0x72, 0x74, 0x73, 0x20, 0x64, 0x65, 0x63, 0x61, 0x79, 0x69,
	0x6e, 0x67, 0x2e, 0xba, 0x48, 0x5a, 0xba, 0x01, 0x57, 0x0a, 0x15, 0x64, 0x65, 0x63, 0x61, 0x79,
	0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x12, 0x2e, 0x64, 0x65, 0x63, 0x61, 0x79, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x20, 0x6d, 0x75, 0x73, 0x74, 0x20, 0x62, 0x65, 0x20,
	0x61, 0x20, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x20, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x65, 0x72, 0x2e,
	0x1a, 0x0e, 0x75, 0x69, 0x6e, 0x74, 0x28, 0x74, 0x68, 0x69, 0x73, 0x29, 0x20, 0x3e, 0x20, 0x30,
	0x52, 0x13, 0x64, 0x65, 0x63, 0x61, 0x79, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0xb8, 0x01, 0x0a, 0x13, 0x64, 0x65, 0x63, 0x61, 0x79, 0x5f,
	0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x04, 0x42, 0x87, 0x01, 0x92, 0x41, 0x2b, 0x32, 0x29, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x20, 0x61, 0x74, 0x20, 0x77, 0x68, 0x69, 0x63, 0x68, 0x20, 0x74, 0x68,
	0x65, 0x20, 0x62, 0x69, 0x64, 0x20, 0x65, 0x6e, 0x64, 0x73, 0x20, 0x64, 0x65, 0x63, 0x61, 0x79,
	0x69, 0x6e, 0x67, 0x2e, 0xba, 0x48, 0x56, 0xba, 0x01, 0x53, 0x0a, 0x13, 0x64, 0x65, 0x63, 0x61,
	0x79, 0x5f, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12,
	0x2c, 0x64, 0x65, 0x63, 0x61, 0x79, 0x5f, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x20, 0x6d, 0x75, 0x73, 0x74, 0x20, 0x62, 0x65, 0x20, 0x61, 0x20, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x20, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x65, 0x72, 0x2e, 0x1a, 0x0e, 0x75,
	0x69, 0x6e, 0x74, 0x28, 0x74, 0x68, 0x69, 0x73, 0x29, 0x20, 0x3e, 0x20, 0x30, 0x52, 0x11, 0x64,
	0x65, 0x63, 0x61, 0x79, 0x45, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x3a, 0xc8, 0x02, 0x92, 0x41, 0xc4, 0x02, 0x0a, 0x71, 0x2a, 0x0b, 0x42, 0x69, 0x64, 0x20, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x40, 0x55, 0x6e, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64,
	0x20, 0x62, 0x69, 0x64, 0x20, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x20, 0x66, 0x72, 0x6f,
	0x6d, 0x20, 0x62, 0x69, 0x64, 0x64, 0x65, 0x72, 0x73, 0x20, 0x74, 0x6f, 0x20, 0x74, 0x68, 0x65,
	0x20, 0x62, 0x69, 0x64, 0x64, 0x65, 0x72, 0x20, 0x6d, 0x65, 0x76, 0x2d, 0x63, 0x6f, 0x6d, 0x6d,
	0x69, 0x74, 0x20, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0xd2, 0x01, 0x08, 0x74, 0x78, 0x48, 0x61, 0x73,
	0x68, 0x65, 0x73, 0xd2, 0x01, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0xd2, 0x01, 0x0b, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x32, 0xce, 0x01, 0x7b, 0x22, 0x74,
	0x78, 0x48, 0x61, 0x73, 0x68, 0x65, 0x73, 0x22, 0x3a, 0x20, 0x5b, 0x22, 0x66, 0x65, 0x34, 0x63,
	0x62, 0x34, 0x37, 0x64, 0x62, 0x33, 0x36, 0x33, 0x30, 0x35, 0x35, 0x31, 0x62, 0x65, 0x65, 0x64,
	0x66, 0x62, 0x64, 0x30, 0x32, 0x61, 0x37, 0x31, 0x65, 0x63, 0x63, 0x36, 0x39, 0x66, 0x64, 0x35,
	0x39, 0x37, 0x35, 0x38, 0x65, 0x32, 0x62, 0x61, 0x36, 0x39, 0x39, 0x36, 0x30, 0x36, 0x65, 0x32,
	0x64, 0x35, 0x63, 0x37, 0x34, 0x32, 0x38, 0x34, 0x66, 0x66, 0x61, 0x37, 0x22, 0x2c, 0x20, 0x22,
	0x37, 0x31, 0x63, 0x31, 0x33, 0x34, 0x38, 0x66, 0x32, 0x64, 0x37, 0x66, 0x66, 0x37, 0x65, 0x38,
	0x31, 0x34, 0x66, 0x39, 0x63, 0x33, 0x36, 0x31, 0x37, 0x39, 0x38, 0x33, 0x37, 0x30, 0x33, 0x34,
	0x33, 0x35, 0x65, 0x61, 0x37, 0x34, 0x34, 0x36, 0x64, 0x65, 0x34, 0x32, 0x30, 0x61, 0x65, 0x61,
	0x63, 0x34, 0x38, 0x38, 0x62, 0x66, 0x31, 0x64, 0x65, 0x33, 0x35, 0x37, 0x33, 0x37, 0x65, 0x38,
	0x22, 0x5d, 0x2c, 0x20, 0x22, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x3a, 0x20, 0x22, 0x31,
	0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
	0x30, 0x30, 0x22, 0x2c, 0x20, 0x22, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x22, 0x3a, 0x20, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x7d, 0x22, 0xf7, 0x09, 0x0a, 0x0a,
	0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x95, 0x01, 0x0a, 0x09, 0x74,
	0x78, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x42, 0x78,
	0x92, 0x41, 0x75, 0x32, 0x61, 0x48, 0x65, 0x78, 0x20, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x20,
	0x65, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x20, 0x6f, 0x66, 0x20, 0x74, 0x68, 0x65, 0x20,
	0x68, 0x61, 0x73, 0x68, 0x20, 0x6f, 0x66, 0x20, 0x74, 0x68, 0x65, 0x20, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x74, 0x68, 0x61, 0x74, 0x20, 0x74, 0x68, 0x65,
	0x20, 0x62, 0x69, 0x64, 0x64, 0x65, 0x72, 0x20, 0x77, 0x61, 0x6e, 0x74, 0x73, 0x20, 0x74, 0x6f,
	0x20, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x20, 0x69, 0x6e, 0x20, 0x74, 0x68, 0x65, 0x20,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x8a, 0x01, 0x0f, 0x5b, 0x61, 0x2d, 0x66, 0x41, 0x2d, 0x46,
	0x30, 0x2d, 0x39, 0x5d, 0x7b, 0x36, 0x34, 0x7d, 0x52, 0x08, 0x74, 0x78, 0x48, 0x61, 0x73, 0x68,
	0x65, 0x73, 0x12, 0x8f, 0x01, 0x0a, 0x0a, 0x62, 0x69, 0x64, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x42, 0x70, 0x92, 0x41, 0x6d, 0x32, 0x6b, 0x41, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x20, 0x6f, 0x66, 0x20, 0x45, 0x54, 0x48, 0x20, 0x74, 0x68, 0x61, 0x74,
	0x20, 0x74, 0x68, 0x65, 0x20, 0x62, 0x69, 0x64, 0x64, 0x65, 0x72, 0x20, 0x68, 0x61, 0x73, 0x20,
	0x61, 0x67, 0x72, 0x65, 0x65, 0x64, 0x20, 0x74, 0x6f, 0x20, 0x70, 0x61, 0x79, 0x20, 0x74, 0x6f,
	0x20, 0x74, 0x68, 0x65, 0x20, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x20, 0x66, 0x6f,
	0x72, 0x20, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x69, 0x6e, 0x67, 0x20, 0x74, 0x68, 0x65, 0x20,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x69, 0x6e, 0x20, 0x74,
	0x68, 0x65, 0x20, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x52, 0x09, 0x62, 0x69, 0x64, 0x41, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x6d, 0x0a, 0x0c, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x6e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x42, 0x4a, 0x92, 0x41, 0x47, 0x32,
	0x45, 0x4d, 0x61, 0x78, 0x20, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x20, 0x6e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x20, 0x74, 0x68, 0x61, 0x74, 0x20, 0x74, 0x68, 0x65, 0x20, 0x62, 0x69, 0x64, 0x64, 0x65,
	0x72, 0x20, 0x77, 0x61, 0x6e, 0x74, 0x73, 0x20, 0x74, 0x6f, 0x20, 0x69, 0x6e, 0x63, 0x6c, 0x75,
	0x64, 0x65, 0x20, 0x74, 0x68, 0x65, 0x20, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x20, 0x69, 0x6e, 0x2e, 0x52, 0x0b, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x12, 0x7b, 0x0a, 0x13, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x5f,
	0x62, 0x69, 0x64, 0x5f, 0x64, 0x69, 0x67, 0x65, 0x73, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x4b, 0x92, 0x41, 0x48, 0x32, 0x46, 0x48, 0x65, 0x78, 0x20, 0x73, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x20, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x20, 0x6f, 0x66, 0x20, 0x64, 0x69,
	0x67, 0x65, 0x73, 0x74, 0x20, 0x6f, 0x66, 0x20, 0x74, 0x68, 0x65, 0x20, 0x62, 0x69, 0x64, 0x20,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x20, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x20, 0x62,
	0x79, 0x20, 0x74, 0x68, 0x65, 0x20, 0x62, 0x69, 0x64, 0x64, 0x65, 0x72, 0x2e, 0x52, 0x11, 0x72,
	0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x42, 0x69, 0x64, 0x44, 0x69, 0x67, 0x65, 0x73, 0x74,
	0x12, 0x7d, 0x0a, 0x16, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x5f, 0x62, 0x69, 0x64,
	0x5f, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x47, 0x92, 0x41, 0x44, 0x32, 0x42, 0x48, 0x65, 0x78, 0x20, 0x73, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x20, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x20, 0x6f, 0x66, 0x20, 0x73, 0x69,
	0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x20, 0x6f, 0x66, 0x20, 0x74, 0x68, 0x65, 0x20, 0x62,
	0x69, 0x64, 0x64, 0x65, 0x72, 0x20, 0x74, 0x68, 0x61, 0x74, 0x20, 0x73, 0x65, 0x6e, 0x74, 0x20,
	0x74, 0x68, 0x69, 0x73, 0x20, 0x62, 0x69, 0x64, 0x2e, 0x52, 0x14, 0x72, 0x65, 0x63, 0x65, 0x69,
	0x76, 0x65, 0x64, 0x42, 0x69, 0x64, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12,
	0x62, 0x0a, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x64, 0x69,
	0x67, 0x65, 0x73, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x42, 0x35, 0x92, 0x41, 0x32, 0x32,
	0x30, 0x48, 0x65, 0x78, 0x20, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x20, 0x65, 0x6e, 0x63, 0x6f,
	0x64, 0x69, 0x6e, 0x67, 0x20, 0x6f, 0x66, 0x20, 0x64, 0x69, 0x67, 0x65, 0x73, 0x74, 0x20, 0x6f,
	0x66, 0x20, 0x74, 0x68, 0x65, 0x20, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x52, 0x10, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x44, 0x69, 0x67,
	0x65, 0x73, 0x74, 0x12, 0x9e, 0x01, 0x0a, 0x14, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x6d, 0x65,
	0x6e, 0x74, 0x5f, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x6b, 0x92, 0x41, 0x68, 0x32, 0x66, 0x48, 0x65, 0x78, 0x20, 0x73, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x20, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x20, 0x6f, 0x66, 0x20,
	0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x20, 0x6f, 0x66, 0x20, 0x74, 0x68, 0x65,
	0x20, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x20, 0x73, 0x69, 0x67, 0x6e,
	0x65, 0x64, 0x20, 0x62, 0x79, 0x20, 0x74, 0x68, 0x65, 0x20, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x20, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x69, 0x6e, 0x67, 0x20, 0x74, 0x68,
	0x69, 0x73, 0x20, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52,
	0x13, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x69, 0x67, 0x6e, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x12, 0x88, 0x01, 0x0a, 0x10, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x5d, 0x92, 0x41, 0x5a, 0x32, 0x58, 0x48, 0x65, 0x78, 0x20, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x20, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x20, 0x6f, 0x66, 0x20, 0x74, 0x68, 0x65,
	0x20, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x20, 0x6f, 0x66, 0x20, 0x74, 0x68, 0x65, 0x20,
	0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x20, 0x74, 0x68, 0x61, 0x74, 0x20, 0x73, 0x69,
	0x67, 0x6e, 0x65, 0x64, 0x20, 0x74, 0x68, 0x65, 0x20, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x6d,
	0x65, 0x6e, 0x74, 0x20, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x52, 0x0f,
	0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12,
	0x64, 0x0a, 0x15, 0x64, 0x65, 0x63, 0x61, 0x79, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x09, 0x20, 0x01, 0x28, 0x04, 0x42, 0x30,
	0x92, 0x41, 0x2d, 0x32, 0x2b, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x20, 0x61,
	0x74, 0x20, 0x77, 0x68, 0x69, 0x63, 0x68, 0x20, 0x74, 0x68, 0x65, 0x20, 0x62, 0x69, 0x64, 0x20,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x73, 0x20, 0x64, 0x65, 0x63, 0x61, 0x79, 0x69, 0x6e, 0x67, 0x2e,
	0x52, 0x13, 0x64, 0x65, 0x63, 0x61, 0x79, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x5e, 0x0a, 0x13, 0x64, 0x65, 0x63, 0x61, 0x79, 0x5f, 0x65,
	0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x04, 0x42, 0x2e, 0x92, 0x41, 0x2b, 0x32, 0x29, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x20, 0x61, 0x74, 0x20, 0x77, 0x68, 0x69, 0x63, 0x68, 0x20, 0x74, 0x68, 0x65, 0x20,
	0x62, 0x69, 0x64, 0x20, 0x65, 0x6e, 0x64, 0x73, 0x20, 0x64, 0x65, 0x63, 0x61, 0x79, 0x69, 0x6e,
	0x67, 0x2e, 0x52, 0x11, 0x64, 0x65, 0x63, 0x61, 0x79, 0x45, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x32, 0xce, 0x03, 0x0a, 0x06, 0x42, 0x69, 0x64, 0x64, 0x65, 0x72,
	0x12, 0x5b, 0x0a, 0x07, 0x53, 0x65, 0x6e, 0x64, 0x42, 0x69, 0x64, 0x12, 0x15, 0x2e, 0x72, 0x70,
	0x63, 0x2e, 0x62, 0x69, 0x64, 0x64, 0x65, 0x72, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x42,
	0x69, 0x64, 0x1a, 0x1c, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x62, 0x69, 0x64, 0x64, 0x65, 0x72, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x6d, 0x65, 0x6e, 0x74,
	0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x3a, 0x01, 0x2a, 0x22, 0x0e, 0x2f, 0x76, 0x31,
	0x2f, 0x62, 0x69, 0x64, 0x64, 0x65, 0x72, 0x2f, 0x62, 0x69, 0x64, 0x30, 0x01, 0x12, 0x78, 0x0a,
	0x0f, 0x50, 0x72, 0x65, 0x70, 0x61, 0x79, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x61, 0x6e, 0x63, 0x65,
	0x12, 0x1f, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x62, 0x69, 0x64, 0x64, 0x65, 0x72, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x65, 0x70, 0x61, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x20, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x62, 0x69, 0x64, 0x64, 0x65, 0x72, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x65, 0x70, 0x61, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x22, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x22, 0x1a, 0x2f, 0x76, 0x31,
	0x2f, 0x62, 0x69, 0x64, 0x64, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x65, 0x70, 0x61, 0x79, 0x2f, 0x7b,
	0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x7d, 0x12, 0x72, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x41, 0x6c,
	0x6c, 0x6f, 0x77, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x1e, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x62, 0x69,
	0x64, 0x64, 0x65, 0x72, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x20, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x62, 0x69,
	0x64, 0x64, 0x65, 0x72, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x65, 0x70, 0x61,
	0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x1a, 0x12, 0x18, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x69, 0x64, 0x64, 0x65, 0x72, 0x2f, 0x67, 0x65,
	0x74, 0x5f, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x79, 0x0a, 0x0f, 0x47,
	0x65, 0x74, 0x4d, 0x69, 0x6e, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x1e,
	0x2e, 0x72, 0x70, 0x63, 0x2e, 0x62, 0x69, 0x64, 0x64, 0x65, 0x72, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x31, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x20,
	0x2e, 0x72, 0x70, 0x63, 0x2e, 0x62, 0x69, 0x64, 0x64, 0x65, 0x72, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x31, 0x2e, 0x50, 0x72, 0x65, 0x70, 0x61, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x24, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x12, 0x1c, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x69,
	0x64, 0x64, 0x65, 0x72, 0x2f, 0x67, 0x65, 0x74, 0x5f, 0x6d, 0x69, 0x6e, 0x5f, 0x61, 0x6c, 0x6c,
	0x6f, 0x77, 0x61, 0x6e, 0x63, 0x65, 0x42, 0xcc, 0x02, 0x92, 0x41, 0x7a, 0x12, 0x78, 0x0a, 0x0a,
	0x42, 0x69, 0x64, 0x64, 0x65, 0x72, 0x20, 0x41, 0x50, 0x49, 0x2a, 0x5d, 0x0a, 0x1b, 0x42, 0x75,
	0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x20, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x20, 0x4c, 0x69,
	0x63, 0x65, 0x6e, 0x73, 0x65, 0x20, 0x31, 0x2e, 0x31, 0x12, 0x3e, 0x68, 0x74, 0x74, 0x70, 0x73,
	0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x72,
	0x69, 0x6d, 0x65, 0x76, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x6d, 0x65, 0x76,
	0x2d, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x2f, 0x62, 0x6c, 0x6f, 0x62, 0x2f, 0x6d, 0x61, 0x69,
	0x6e, 0x2f, 0x4c, 0x49, 0x43, 0x45, 0x4e, 0x53, 0x45, 0x32, 0x0b, 0x31, 0x2e, 0x30, 0x2e, 0x30,
	0x2d, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x0a, 0x14, 0x63, 0x6f, 0x6d, 0x2e, 0x72, 0x70, 0x63, 0x2e,
	0x62, 0x69, 0x64, 0x64, 0x65, 0x72, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x42, 0x0e, 0x42, 0x69,
	0x64, 0x64, 0x65, 0x72, 0x61, 0x70, 0x69, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x45,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x72, 0x69, 0x6d, 0x65,
	0x76, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x6d, 0x65, 0x76, 0x2d, 0x63, 0x6f,
	0x6d, 0x6d, 0x69, 0x74, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x62, 0x69, 0x64,
	0x64, 0x65, 0x72, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x3b, 0x62, 0x69, 0x64, 0x64, 0x65, 0x72,
	0x61, 0x70, 0x69, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x52, 0x42, 0x58, 0xaa, 0x02, 0x10, 0x52, 0x70,
	0x63, 0x2e, 0x42, 0x69, 0x64, 0x64, 0x65, 0x72, 0x61, 0x70, 0x69, 0x2e, 0x56, 0x31, 0xca, 0x02,
	0x10, 0x52, 0x70, 0x63, 0x5c, 0x42, 0x69, 0x64, 0x64, 0x65, 0x72, 0x61, 0x70, 0x69, 0x5c, 0x56,
	0x31, 0xe2, 0x02, 0x1c, 0x52, 0x70, 0x63, 0x5c, 0x42, 0x69, 0x64, 0x64, 0x65, 0x72, 0x61, 0x70,
	0x69, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0xea, 0x02, 0x12, 0x52, 0x70, 0x63, 0x3a, 0x3a, 0x42, 0x69, 0x64, 0x64, 0x65, 0x72, 0x61, 0x70,
	0x69, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpc_bidderapi_v1_bidderapi_proto_rawDescOnce sync.Once
	file_rpc_bidderapi_v1_bidderapi_proto_rawDescData = file_rpc_bidderapi_v1_bidderapi_proto_rawDesc
)

func file_rpc_bidderapi_v1_bidderapi_proto_rawDescGZIP() []byte {
	file_rpc_bidderapi_v1_bidderapi_proto_rawDescOnce.Do(func() {
		file_rpc_bidderapi_v1_bidderapi_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_bidderapi_v1_bidderapi_proto_rawDescData)
	})
	return file_rpc_bidderapi_v1_bidderapi_proto_rawDescData
}

var file_rpc_bidderapi_v1_bidderapi_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_rpc_bidderapi_v1_bidderapi_proto_goTypes = []interface{}{
	(*PrepayRequest)(nil),  // 0: rpc.bidderapi.v1.PrepayRequest
	(*PrepayResponse)(nil), // 1: rpc.bidderapi.v1.PrepayResponse
	(*EmptyMessage)(nil),   // 2: rpc.bidderapi.v1.EmptyMessage
	(*Bid)(nil),            // 3: rpc.bidderapi.v1.Bid
	(*Commitment)(nil),     // 4: rpc.bidderapi.v1.Commitment
}
var file_rpc_bidderapi_v1_bidderapi_proto_depIdxs = []int32{
	3, // 0: rpc.bidderapi.v1.Bidder.SendBid:input_type -> rpc.bidderapi.v1.Bid
	0, // 1: rpc.bidderapi.v1.Bidder.PrepayAllowance:input_type -> rpc.bidderapi.v1.PrepayRequest
	2, // 2: rpc.bidderapi.v1.Bidder.GetAllowance:input_type -> rpc.bidderapi.v1.EmptyMessage
	2, // 3: rpc.bidderapi.v1.Bidder.GetMinAllowance:input_type -> rpc.bidderapi.v1.EmptyMessage
	4, // 4: rpc.bidderapi.v1.Bidder.SendBid:output_type -> rpc.bidderapi.v1.Commitment
	1, // 5: rpc.bidderapi.v1.Bidder.PrepayAllowance:output_type -> rpc.bidderapi.v1.PrepayResponse
	1, // 6: rpc.bidderapi.v1.Bidder.GetAllowance:output_type -> rpc.bidderapi.v1.PrepayResponse
	1, // 7: rpc.bidderapi.v1.Bidder.GetMinAllowance:output_type -> rpc.bidderapi.v1.PrepayResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_rpc_bidderapi_v1_bidderapi_proto_init() }
func file_rpc_bidderapi_v1_bidderapi_proto_init() {
	if File_rpc_bidderapi_v1_bidderapi_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rpc_bidderapi_v1_bidderapi_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrepayRequest); i {
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
		file_rpc_bidderapi_v1_bidderapi_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrepayResponse); i {
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
		file_rpc_bidderapi_v1_bidderapi_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmptyMessage); i {
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
		file_rpc_bidderapi_v1_bidderapi_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_rpc_bidderapi_v1_bidderapi_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Commitment); i {
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
			RawDescriptor: file_rpc_bidderapi_v1_bidderapi_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rpc_bidderapi_v1_bidderapi_proto_goTypes,
		DependencyIndexes: file_rpc_bidderapi_v1_bidderapi_proto_depIdxs,
		MessageInfos:      file_rpc_bidderapi_v1_bidderapi_proto_msgTypes,
	}.Build()
	File_rpc_bidderapi_v1_bidderapi_proto = out.File
	file_rpc_bidderapi_v1_bidderapi_proto_rawDesc = nil
	file_rpc_bidderapi_v1_bidderapi_proto_goTypes = nil
	file_rpc_bidderapi_v1_bidderapi_proto_depIdxs = nil
}
