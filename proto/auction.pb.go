// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.7
// source: proto/auction.proto

package auction

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

type Status int32

const (
	Status_FAIL      Status = 0
	Status_SUCCESS   Status = 1
	Status_EXECPTION Status = 2
)

// Enum value maps for Status.
var (
	Status_name = map[int32]string{
		0: "FAIL",
		1: "SUCCESS",
		2: "EXECPTION",
	}
	Status_value = map[string]int32{
		"FAIL":      0,
		"SUCCESS":   1,
		"EXECPTION": 2,
	}
)

func (x Status) Enum() *Status {
	p := new(Status)
	*p = x
	return p
}

func (x Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Status) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_auction_proto_enumTypes[0].Descriptor()
}

func (Status) Type() protoreflect.EnumType {
	return &file_proto_auction_proto_enumTypes[0]
}

func (x Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Status.Descriptor instead.
func (Status) EnumDescriptor() ([]byte, []int) {
	return file_proto_auction_proto_rawDescGZIP(), []int{0}
}

type BidRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Amount   int32 `protobuf:"varint,1,opt,name=amount,proto3" json:"amount,omitempty"`
	BidderID int32 `protobuf:"varint,2,opt,name=BidderID,proto3" json:"BidderID,omitempty"`
}

func (x *BidRequest) Reset() {
	*x = BidRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_auction_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BidRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BidRequest) ProtoMessage() {}

func (x *BidRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_auction_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BidRequest.ProtoReflect.Descriptor instead.
func (*BidRequest) Descriptor() ([]byte, []int) {
	return file_proto_auction_proto_rawDescGZIP(), []int{0}
}

func (x *BidRequest) GetAmount() int32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *BidRequest) GetBidderID() int32 {
	if x != nil {
		return x.BidderID
	}
	return 0
}

type BidResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status Status `protobuf:"varint,1,opt,name=status,proto3,enum=auction_system.Status" json:"status,omitempty"`
}

func (x *BidResponse) Reset() {
	*x = BidResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_auction_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BidResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BidResponse) ProtoMessage() {}

func (x *BidResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_auction_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BidResponse.ProtoReflect.Descriptor instead.
func (*BidResponse) Descriptor() ([]byte, []int) {
	return file_proto_auction_proto_rawDescGZIP(), []int{1}
}

func (x *BidResponse) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_FAIL
}

type ResultRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientId int32 `protobuf:"varint,1,opt,name=clientId,proto3" json:"clientId,omitempty"`
}

func (x *ResultRequest) Reset() {
	*x = ResultRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_auction_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResultRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResultRequest) ProtoMessage() {}

func (x *ResultRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_auction_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResultRequest.ProtoReflect.Descriptor instead.
func (*ResultRequest) Descriptor() ([]byte, []int) {
	return file_proto_auction_proto_rawDescGZIP(), []int{2}
}

func (x *ResultRequest) GetClientId() int32 {
	if x != nil {
		return x.ClientId
	}
	return 0
}

type ResultResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HighestBidderId int32 `protobuf:"varint,1,opt,name=highestBidderId,proto3" json:"highestBidderId,omitempty"`
	HighestBid      int32 `protobuf:"varint,2,opt,name=highestBid,proto3" json:"highestBid,omitempty"`
}

func (x *ResultResponse) Reset() {
	*x = ResultResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_auction_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResultResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResultResponse) ProtoMessage() {}

func (x *ResultResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_auction_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResultResponse.ProtoReflect.Descriptor instead.
func (*ResultResponse) Descriptor() ([]byte, []int) {
	return file_proto_auction_proto_rawDescGZIP(), []int{3}
}

func (x *ResultResponse) GetHighestBidderId() int32 {
	if x != nil {
		return x.HighestBidderId
	}
	return 0
}

func (x *ResultResponse) GetHighestBid() int32 {
	if x != nil {
		return x.HighestBid
	}
	return 0
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_auction_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_proto_auction_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_proto_auction_proto_rawDescGZIP(), []int{4}
}

var File_proto_auction_proto protoreflect.FileDescriptor

var file_proto_auction_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x22, 0x40, 0x0a, 0x0a, 0x42, 0x69, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x42,
	0x69, 0x64, 0x64, 0x65, 0x72, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x42,
	0x69, 0x64, 0x64, 0x65, 0x72, 0x49, 0x44, 0x22, 0x3d, 0x0a, 0x0b, 0x42, 0x69, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x2b, 0x0a, 0x0d, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x49, 0x64, 0x22, 0x5a, 0x0a, 0x0e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x68, 0x69, 0x67, 0x68, 0x65, 0x73, 0x74,
	0x42, 0x69, 0x64, 0x64, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f,
	0x68, 0x69, 0x67, 0x68, 0x65, 0x73, 0x74, 0x42, 0x69, 0x64, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x1e, 0x0a, 0x0a, 0x68, 0x69, 0x67, 0x68, 0x65, 0x73, 0x74, 0x42, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0a, 0x68, 0x69, 0x67, 0x68, 0x65, 0x73, 0x74, 0x42, 0x69, 0x64, 0x22,
	0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x2a, 0x2e, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x08, 0x0a, 0x04, 0x46, 0x41, 0x49, 0x4c, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07,
	0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x45, 0x58, 0x45,
	0x43, 0x50, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x02, 0x32, 0xe4, 0x01, 0x0a, 0x0e, 0x41, 0x75, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x40, 0x0a, 0x03, 0x42,
	0x69, 0x64, 0x12, 0x1a, 0x2e, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x2e, 0x42, 0x69, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b,
	0x2e, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e,
	0x42, 0x69, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x49, 0x0a,
	0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x1d, 0x2e, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x45, 0x0a, 0x0a, 0x45, 0x6e, 0x64, 0x41,
	0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x15, 0x2e, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1e, 0x2e,
	0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42,
	0x0f, 0x5a, 0x0d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_auction_proto_rawDescOnce sync.Once
	file_proto_auction_proto_rawDescData = file_proto_auction_proto_rawDesc
)

func file_proto_auction_proto_rawDescGZIP() []byte {
	file_proto_auction_proto_rawDescOnce.Do(func() {
		file_proto_auction_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_auction_proto_rawDescData)
	})
	return file_proto_auction_proto_rawDescData
}

var file_proto_auction_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_auction_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_auction_proto_goTypes = []interface{}{
	(Status)(0),            // 0: auction_system.Status
	(*BidRequest)(nil),     // 1: auction_system.BidRequest
	(*BidResponse)(nil),    // 2: auction_system.BidResponse
	(*ResultRequest)(nil),  // 3: auction_system.ResultRequest
	(*ResultResponse)(nil), // 4: auction_system.ResultResponse
	(*Empty)(nil),          // 5: auction_system.Empty
}
var file_proto_auction_proto_depIdxs = []int32{
	0, // 0: auction_system.BidResponse.status:type_name -> auction_system.Status
	1, // 1: auction_system.AuctionService.Bid:input_type -> auction_system.BidRequest
	3, // 2: auction_system.AuctionService.Result:input_type -> auction_system.ResultRequest
	5, // 3: auction_system.AuctionService.EndAuction:input_type -> auction_system.Empty
	2, // 4: auction_system.AuctionService.Bid:output_type -> auction_system.BidResponse
	4, // 5: auction_system.AuctionService.Result:output_type -> auction_system.ResultResponse
	4, // 6: auction_system.AuctionService.EndAuction:output_type -> auction_system.ResultResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_auction_proto_init() }
func file_proto_auction_proto_init() {
	if File_proto_auction_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_auction_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BidRequest); i {
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
		file_proto_auction_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BidResponse); i {
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
		file_proto_auction_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResultRequest); i {
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
		file_proto_auction_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResultResponse); i {
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
		file_proto_auction_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
			RawDescriptor: file_proto_auction_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_auction_proto_goTypes,
		DependencyIndexes: file_proto_auction_proto_depIdxs,
		EnumInfos:         file_proto_auction_proto_enumTypes,
		MessageInfos:      file_proto_auction_proto_msgTypes,
	}.Build()
	File_proto_auction_proto = out.File
	file_proto_auction_proto_rawDesc = nil
	file_proto_auction_proto_goTypes = nil
	file_proto_auction_proto_depIdxs = nil
}
