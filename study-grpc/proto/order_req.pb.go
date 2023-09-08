// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.20.1
// 	protoc        v3.15.4
// source: order_req.proto

package proto

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type CreateOrderReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Amount int64               `protobuf:"varint,1,opt,name=amount,proto3" json:"amount,omitempty"`
	User   *CreateOrderUser    `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	Goods  []*CreateOrderGoods `protobuf:"bytes,3,rep,name=goods,proto3" json:"goods,omitempty"`
}

func (x *CreateOrderReq) Reset() {
	*x = CreateOrderReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_req_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateOrderReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOrderReq) ProtoMessage() {}

func (x *CreateOrderReq) ProtoReflect() protoreflect.Message {
	mi := &file_order_req_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOrderReq.ProtoReflect.Descriptor instead.
func (*CreateOrderReq) Descriptor() ([]byte, []int) {
	return file_order_req_proto_rawDescGZIP(), []int{0}
}

func (x *CreateOrderReq) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *CreateOrderReq) GetUser() *CreateOrderUser {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *CreateOrderReq) GetGoods() []*CreateOrderGoods {
	if x != nil {
		return x.Goods
	}
	return nil
}

type CreateOrderUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId  int64  `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	Phone   string `protobuf:"bytes,2,opt,name=phone,proto3" json:"phone,omitempty"`
	Address string `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	AreaId  int64  `protobuf:"varint,4,opt,name=areaId,proto3" json:"areaId,omitempty"`
}

func (x *CreateOrderUser) Reset() {
	*x = CreateOrderUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_req_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateOrderUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOrderUser) ProtoMessage() {}

func (x *CreateOrderUser) ProtoReflect() protoreflect.Message {
	mi := &file_order_req_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOrderUser.ProtoReflect.Descriptor instead.
func (*CreateOrderUser) Descriptor() ([]byte, []int) {
	return file_order_req_proto_rawDescGZIP(), []int{1}
}

func (x *CreateOrderUser) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *CreateOrderUser) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *CreateOrderUser) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *CreateOrderUser) GetAreaId() int64 {
	if x != nil {
		return x.AreaId
	}
	return 0
}

type CreateOrderGoods struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SkuId  string `protobuf:"bytes,1,opt,name=skuId,proto3" json:"skuId,omitempty"`
	Number int32  `protobuf:"varint,2,opt,name=number,proto3" json:"number,omitempty"`
	Price  int64  `protobuf:"varint,3,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *CreateOrderGoods) Reset() {
	*x = CreateOrderGoods{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_req_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateOrderGoods) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOrderGoods) ProtoMessage() {}

func (x *CreateOrderGoods) ProtoReflect() protoreflect.Message {
	mi := &file_order_req_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOrderGoods.ProtoReflect.Descriptor instead.
func (*CreateOrderGoods) Descriptor() ([]byte, []int) {
	return file_order_req_proto_rawDescGZIP(), []int{2}
}

func (x *CreateOrderGoods) GetSkuId() string {
	if x != nil {
		return x.SkuId
	}
	return ""
}

func (x *CreateOrderGoods) GetNumber() int32 {
	if x != nil {
		return x.Number
	}
	return 0
}

func (x *CreateOrderGoods) GetPrice() int64 {
	if x != nil {
		return x.Price
	}
	return 0
}

var File_order_req_proto protoreflect.FileDescriptor

var file_order_req_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x72, 0x65, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x83, 0x01, 0x0a, 0x0e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x61,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x61, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x2a, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12,
	0x2d, 0x0a, 0x05, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x52, 0x05, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x22, 0x71,
	0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x55, 0x73, 0x65,
	0x72, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f,
	0x6e, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x72, 0x65,
	0x61, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x61, 0x72, 0x65, 0x61, 0x49,
	0x64, 0x22, 0x56, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x47, 0x6f, 0x6f, 0x64, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x6b, 0x75, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x6b, 0x75, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x42, 0x12, 0x5a, 0x10, 0x73, 0x74, 0x75,
	0x64, 0x79, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_order_req_proto_rawDescOnce sync.Once
	file_order_req_proto_rawDescData = file_order_req_proto_rawDesc
)

func file_order_req_proto_rawDescGZIP() []byte {
	file_order_req_proto_rawDescOnce.Do(func() {
		file_order_req_proto_rawDescData = protoimpl.X.CompressGZIP(file_order_req_proto_rawDescData)
	})
	return file_order_req_proto_rawDescData
}

var file_order_req_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_order_req_proto_goTypes = []interface{}{
	(*CreateOrderReq)(nil),   // 0: proto.CreateOrderReq
	(*CreateOrderUser)(nil),  // 1: proto.CreateOrderUser
	(*CreateOrderGoods)(nil), // 2: proto.CreateOrderGoods
}
var file_order_req_proto_depIdxs = []int32{
	1, // 0: proto.CreateOrderReq.user:type_name -> proto.CreateOrderUser
	2, // 1: proto.CreateOrderReq.goods:type_name -> proto.CreateOrderGoods
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_order_req_proto_init() }
func file_order_req_proto_init() {
	if File_order_req_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_order_req_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateOrderReq); i {
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
		file_order_req_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateOrderUser); i {
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
		file_order_req_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateOrderGoods); i {
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
			RawDescriptor: file_order_req_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_order_req_proto_goTypes,
		DependencyIndexes: file_order_req_proto_depIdxs,
		MessageInfos:      file_order_req_proto_msgTypes,
	}.Build()
	File_order_req_proto = out.File
	file_order_req_proto_rawDesc = nil
	file_order_req_proto_goTypes = nil
	file_order_req_proto_depIdxs = nil
}
