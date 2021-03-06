// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/product/product.proto

package product

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type ProductInput struct {
	Sku                  string   `protobuf:"bytes,1,opt,name=sku,proto3" json:"sku,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	ListPrice            int32    `protobuf:"varint,3,opt,name=ListPrice,proto3" json:"ListPrice,omitempty"`
	Price                int32    `protobuf:"varint,4,opt,name=Price,proto3" json:"Price,omitempty"`
	PriceWithCC          int32    `protobuf:"varint,5,opt,name=PriceWithCC,proto3" json:"PriceWithCC,omitempty"`
	Link                 string   `protobuf:"bytes,6,opt,name=link,proto3" json:"link,omitempty"`
	Images               []string `protobuf:"bytes,7,rep,name=images,proto3" json:"images,omitempty"`
	HasStock             bool     `protobuf:"varint,8,opt,name=hasStock,proto3" json:"hasStock,omitempty"`
	Category             string   `protobuf:"bytes,9,opt,name=category,proto3" json:"category,omitempty"`
	StoreId              int32    `protobuf:"varint,10,opt,name=storeId,proto3" json:"storeId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProductInput) Reset()         { *m = ProductInput{} }
func (m *ProductInput) String() string { return proto.CompactTextString(m) }
func (*ProductInput) ProtoMessage()    {}
func (*ProductInput) Descriptor() ([]byte, []int) {
	return fileDescriptor_14fbc13de7c38f78, []int{0}
}

func (m *ProductInput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProductInput.Unmarshal(m, b)
}
func (m *ProductInput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProductInput.Marshal(b, m, deterministic)
}
func (m *ProductInput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProductInput.Merge(m, src)
}
func (m *ProductInput) XXX_Size() int {
	return xxx_messageInfo_ProductInput.Size(m)
}
func (m *ProductInput) XXX_DiscardUnknown() {
	xxx_messageInfo_ProductInput.DiscardUnknown(m)
}

var xxx_messageInfo_ProductInput proto.InternalMessageInfo

func (m *ProductInput) GetSku() string {
	if m != nil {
		return m.Sku
	}
	return ""
}

func (m *ProductInput) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ProductInput) GetListPrice() int32 {
	if m != nil {
		return m.ListPrice
	}
	return 0
}

func (m *ProductInput) GetPrice() int32 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *ProductInput) GetPriceWithCC() int32 {
	if m != nil {
		return m.PriceWithCC
	}
	return 0
}

func (m *ProductInput) GetLink() string {
	if m != nil {
		return m.Link
	}
	return ""
}

func (m *ProductInput) GetImages() []string {
	if m != nil {
		return m.Images
	}
	return nil
}

func (m *ProductInput) GetHasStock() bool {
	if m != nil {
		return m.HasStock
	}
	return false
}

func (m *ProductInput) GetCategory() string {
	if m != nil {
		return m.Category
	}
	return ""
}

func (m *ProductInput) GetStoreId() int32 {
	if m != nil {
		return m.StoreId
	}
	return 0
}

type SuccessOutput struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SuccessOutput) Reset()         { *m = SuccessOutput{} }
func (m *SuccessOutput) String() string { return proto.CompactTextString(m) }
func (*SuccessOutput) ProtoMessage()    {}
func (*SuccessOutput) Descriptor() ([]byte, []int) {
	return fileDescriptor_14fbc13de7c38f78, []int{1}
}

func (m *SuccessOutput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SuccessOutput.Unmarshal(m, b)
}
func (m *SuccessOutput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SuccessOutput.Marshal(b, m, deterministic)
}
func (m *SuccessOutput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SuccessOutput.Merge(m, src)
}
func (m *SuccessOutput) XXX_Size() int {
	return xxx_messageInfo_SuccessOutput.Size(m)
}
func (m *SuccessOutput) XXX_DiscardUnknown() {
	xxx_messageInfo_SuccessOutput.DiscardUnknown(m)
}

var xxx_messageInfo_SuccessOutput proto.InternalMessageInfo

func (m *SuccessOutput) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *SuccessOutput) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*ProductInput)(nil), "product.ProductInput")
	proto.RegisterType((*SuccessOutput)(nil), "product.SuccessOutput")
}

func init() { proto.RegisterFile("proto/product/product.proto", fileDescriptor_14fbc13de7c38f78) }

var fileDescriptor_14fbc13de7c38f78 = []byte{
	// 294 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x91, 0x5f, 0x4b, 0x84, 0x40,
	0x14, 0xc5, 0x73, 0xff, 0xa9, 0xb7, 0x5a, 0x62, 0xa8, 0x65, 0xd8, 0x7a, 0x10, 0x9f, 0x7c, 0xda,
	0xa0, 0x3e, 0x41, 0xf8, 0xb4, 0x10, 0xec, 0xa2, 0x44, 0xcf, 0xd3, 0x38, 0xb8, 0x62, 0x3a, 0x32,
	0x73, 0x0d, 0xfa, 0x06, 0x7d, 0xec, 0x98, 0xd1, 0xb1, 0xed, 0xc9, 0xf3, 0x3b, 0x47, 0xbc, 0xd7,
	0x73, 0xe1, 0xbe, 0x53, 0x12, 0xe5, 0x63, 0xa7, 0x64, 0xd1, 0x73, 0x74, 0xcf, 0x9d, 0x75, 0x89,
	0x3f, 0x62, 0xfc, 0x33, 0x83, 0xab, 0xe3, 0xa0, 0xf7, 0x6d, 0xd7, 0x23, 0xb9, 0x81, 0xb9, 0xae,
	0x7b, 0xea, 0x45, 0x5e, 0x12, 0x66, 0x46, 0x12, 0x02, 0x8b, 0x96, 0x35, 0x82, 0xce, 0xac, 0x65,
	0x35, 0x79, 0x80, 0xf0, 0xb5, 0xd2, 0x78, 0x54, 0x15, 0x17, 0x74, 0x1e, 0x79, 0xc9, 0x32, 0xfb,
	0x33, 0xc8, 0x2d, 0x2c, 0x87, 0x64, 0x61, 0x93, 0x01, 0x48, 0x04, 0x97, 0x56, 0xbc, 0x57, 0x78,
	0x4a, 0x53, 0xba, 0xb4, 0xd9, 0xb9, 0x65, 0x26, 0x7d, 0x56, 0x6d, 0x4d, 0x57, 0xc3, 0x24, 0xa3,
	0xc9, 0x06, 0x56, 0x55, 0xc3, 0x4a, 0xa1, 0xa9, 0x1f, 0xcd, 0x93, 0x30, 0x1b, 0x89, 0x6c, 0x21,
	0x38, 0x31, 0x9d, 0xa3, 0xe4, 0x35, 0x0d, 0x22, 0x2f, 0x09, 0xb2, 0x89, 0x4d, 0xc6, 0x19, 0x8a,
	0x52, 0xaa, 0x6f, 0x1a, 0xda, 0x6f, 0x4d, 0x4c, 0x28, 0xf8, 0x1a, 0xa5, 0x12, 0xfb, 0x82, 0x82,
	0xdd, 0xc0, 0x61, 0x9c, 0xc2, 0x75, 0xde, 0x73, 0x2e, 0xb4, 0x3e, 0xf4, 0x68, 0xaa, 0x30, 0xaf,
	0x0e, 0x86, 0xad, 0x23, 0xc8, 0x1c, 0x9a, 0xa4, 0x11, 0x5a, 0xb3, 0xd2, 0xb5, 0xe2, 0xf0, 0x29,
	0x87, 0xf5, 0x58, 0x67, 0x2e, 0xd4, 0x97, 0xf9, 0xed, 0x17, 0x58, 0xa7, 0x4a, 0x30, 0x14, 0x07,
	0xf5, 0xd6, 0x15, 0x0c, 0x05, 0xb9, 0xdb, 0xb9, 0x63, 0x9c, 0x37, 0xbf, 0xdd, 0x4c, 0xf6, 0xbf,
	0x35, 0xe2, 0x8b, 0x8f, 0x95, 0x3d, 0xda, 0xf3, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x48, 0x6b,
	0x97, 0xa8, 0xd3, 0x01, 0x00, 0x00,
}
