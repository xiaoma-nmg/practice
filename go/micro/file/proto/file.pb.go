// Code generated by protoc-gen-go. DO NOT EDIT.
// source: file.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// 定义文件传输切片
// 是stream类型通信的最小单位
type FileSlice struct {
	Byte                 []byte   `protobuf:"bytes,1,opt,name=byte,proto3" json:"byte,omitempty"`
	Len                  int64    `protobuf:"varint,2,opt,name=len,proto3" json:"len,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FileSlice) Reset()         { *m = FileSlice{} }
func (m *FileSlice) String() string { return proto.CompactTextString(m) }
func (*FileSlice) ProtoMessage()    {}
func (*FileSlice) Descriptor() ([]byte, []int) {
	return fileDescriptor_file_8484b9eb3383a247, []int{0}
}
func (m *FileSlice) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileSlice.Unmarshal(m, b)
}
func (m *FileSlice) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileSlice.Marshal(b, m, deterministic)
}
func (dst *FileSlice) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileSlice.Merge(dst, src)
}
func (m *FileSlice) XXX_Size() int {
	return xxx_messageInfo_FileSlice.Size(m)
}
func (m *FileSlice) XXX_DiscardUnknown() {
	xxx_messageInfo_FileSlice.DiscardUnknown(m)
}

var xxx_messageInfo_FileSlice proto.InternalMessageInfo

func (m *FileSlice) GetByte() []byte {
	if m != nil {
		return m.Byte
	}
	return nil
}

func (m *FileSlice) GetLen() int64 {
	if m != nil {
		return m.Len
	}
	return 0
}

type FileSliceMsg struct {
	FileName             string   `protobuf:"bytes,1,opt,name=file_name,json=fileName,proto3" json:"file_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FileSliceMsg) Reset()         { *m = FileSliceMsg{} }
func (m *FileSliceMsg) String() string { return proto.CompactTextString(m) }
func (*FileSliceMsg) ProtoMessage()    {}
func (*FileSliceMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_file_8484b9eb3383a247, []int{1}
}
func (m *FileSliceMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileSliceMsg.Unmarshal(m, b)
}
func (m *FileSliceMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileSliceMsg.Marshal(b, m, deterministic)
}
func (dst *FileSliceMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileSliceMsg.Merge(dst, src)
}
func (m *FileSliceMsg) XXX_Size() int {
	return xxx_messageInfo_FileSliceMsg.Size(m)
}
func (m *FileSliceMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_FileSliceMsg.DiscardUnknown(m)
}

var xxx_messageInfo_FileSliceMsg proto.InternalMessageInfo

func (m *FileSliceMsg) GetFileName() string {
	if m != nil {
		return m.FileName
	}
	return ""
}

func init() {
	proto.RegisterType((*FileSlice)(nil), "pb.FileSlice")
	proto.RegisterType((*FileSliceMsg)(nil), "pb.FileSliceMsg")
}

func init() { proto.RegisterFile("file.proto", fileDescriptor_file_8484b9eb3383a247) }

var fileDescriptor_file_8484b9eb3383a247 = []byte{
	// 148 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0xcb, 0xcc, 0x49,
	0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x52, 0x32, 0xe4, 0xe2, 0x74, 0xcb,
	0xcc, 0x49, 0x0d, 0xce, 0xc9, 0x4c, 0x4e, 0x15, 0x12, 0xe2, 0x62, 0x49, 0xaa, 0x2c, 0x49, 0x95,
	0x60, 0x54, 0x60, 0xd4, 0xe0, 0x09, 0x02, 0xb3, 0x85, 0x04, 0xb8, 0x98, 0x73, 0x52, 0xf3, 0x24,
	0x98, 0x14, 0x18, 0x35, 0x98, 0x83, 0x40, 0x4c, 0x25, 0x6d, 0x2e, 0x1e, 0xb8, 0x16, 0xdf, 0xe2,
	0x74, 0x21, 0x69, 0x2e, 0x4e, 0x90, 0xa1, 0xf1, 0x79, 0x89, 0xb9, 0x10, 0xad, 0x9c, 0x41, 0x1c,
	0x20, 0x01, 0xbf, 0xc4, 0xdc, 0x54, 0x23, 0x63, 0x2e, 0x16, 0x90, 0x62, 0x21, 0x6d, 0x28, 0xcd,
	0xab, 0x57, 0x90, 0xa4, 0x07, 0xd7, 0x2e, 0x25, 0x80, 0xc2, 0xf5, 0x2d, 0x4e, 0x57, 0x62, 0xd0,
	0x60, 0x4c, 0x62, 0x03, 0xbb, 0xcf, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x48, 0x82, 0x69, 0x82,
	0xad, 0x00, 0x00, 0x00,
}
