// Code generated by protoc-gen-go. DO NOT EDIT.
// source: HouseContractListInput.proto

/*
Package example is a generated protocol buffer package.

It is generated from these files:
	HouseContractListInput.proto
	HouseContractListOutput.proto

It has these top-level messages:
	BaseInput
	HouseContractListInput
	BaseOutput
	HouseContractListDto
	HouseContractListOutput
*/
package example

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

type SourceWayEnum int32

const (
	SourceWayEnum_NoSettingSourceWay SourceWayEnum = 0
	SourceWayEnum_IOS                SourceWayEnum = 20
	SourceWayEnum_Android            SourceWayEnum = 30
	SourceWayEnum_Web                SourceWayEnum = 40
)

var SourceWayEnum_name = map[int32]string{
	0:  "NoSettingSourceWay",
	20: "IOS",
	30: "Android",
	40: "Web",
}
var SourceWayEnum_value = map[string]int32{
	"NoSettingSourceWay": 0,
	"IOS":                20,
	"Android":            30,
	"Web":                40,
}

func (x SourceWayEnum) String() string {
	return proto.EnumName(SourceWayEnum_name, int32(x))
}
func (SourceWayEnum) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type BaseInput struct {
	AppVersion    string        `protobuf:"bytes,1,opt,name=AppVersion" json:"AppVersion,omitempty"`
	SourceWayEnum SourceWayEnum `protobuf:"varint,2,opt,name=SourceWayEnum,enum=example.SourceWayEnum" json:"SourceWayEnum,omitempty"`
	TerminalExt   string        `protobuf:"bytes,3,opt,name=TerminalExt" json:"TerminalExt,omitempty"`
	IMEI          string        `protobuf:"bytes,4,opt,name=IMEI" json:"IMEI,omitempty"`
	OpenId        string        `protobuf:"bytes,5,opt,name=OpenId" json:"OpenId,omitempty"`
}

func (m *BaseInput) Reset()                    { *m = BaseInput{} }
func (m *BaseInput) String() string            { return proto.CompactTextString(m) }
func (*BaseInput) ProtoMessage()               {}
func (*BaseInput) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *BaseInput) GetAppVersion() string {
	if m != nil {
		return m.AppVersion
	}
	return ""
}

func (m *BaseInput) GetSourceWayEnum() SourceWayEnum {
	if m != nil {
		return m.SourceWayEnum
	}
	return SourceWayEnum_NoSettingSourceWay
}

func (m *BaseInput) GetTerminalExt() string {
	if m != nil {
		return m.TerminalExt
	}
	return ""
}

func (m *BaseInput) GetIMEI() string {
	if m != nil {
		return m.IMEI
	}
	return ""
}

func (m *BaseInput) GetOpenId() string {
	if m != nil {
		return m.OpenId
	}
	return ""
}

type HouseContractListInput struct {
	BaseInput     *BaseInput `protobuf:"bytes,1,opt,name=BaseInput" json:"BaseInput,omitempty"`
	PageIndex     int32      `protobuf:"varint,2,opt,name=PageIndex" json:"PageIndex,omitempty"`
	PageSize      int32      `protobuf:"varint,3,opt,name=PageSize" json:"PageSize,omitempty"`
	StatusId      int32      `protobuf:"varint,4,opt,name=StatusId" json:"StatusId,omitempty"`
	KeyWordName   string     `protobuf:"bytes,5,opt,name=KeyWordName" json:"KeyWordName,omitempty"`
	UnixBeginDate int64      `protobuf:"varint,6,opt,name=UnixBeginDate" json:"UnixBeginDate,omitempty"`
	UnixEndDate   int64      `protobuf:"varint,7,opt,name=UnixEndDate" json:"UnixEndDate,omitempty"`
}

func (m *HouseContractListInput) Reset()                    { *m = HouseContractListInput{} }
func (m *HouseContractListInput) String() string            { return proto.CompactTextString(m) }
func (*HouseContractListInput) ProtoMessage()               {}
func (*HouseContractListInput) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *HouseContractListInput) GetBaseInput() *BaseInput {
	if m != nil {
		return m.BaseInput
	}
	return nil
}

func (m *HouseContractListInput) GetPageIndex() int32 {
	if m != nil {
		return m.PageIndex
	}
	return 0
}

func (m *HouseContractListInput) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *HouseContractListInput) GetStatusId() int32 {
	if m != nil {
		return m.StatusId
	}
	return 0
}

func (m *HouseContractListInput) GetKeyWordName() string {
	if m != nil {
		return m.KeyWordName
	}
	return ""
}

func (m *HouseContractListInput) GetUnixBeginDate() int64 {
	if m != nil {
		return m.UnixBeginDate
	}
	return 0
}

func (m *HouseContractListInput) GetUnixEndDate() int64 {
	if m != nil {
		return m.UnixEndDate
	}
	return 0
}

func init() {
	proto.RegisterType((*BaseInput)(nil), "example.BaseInput")
	proto.RegisterType((*HouseContractListInput)(nil), "example.HouseContractListInput")
	proto.RegisterEnum("example.SourceWayEnum", SourceWayEnum_name, SourceWayEnum_value)
}

func init() { proto.RegisterFile("HouseContractListInput.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 347 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0x51, 0x6b, 0xe2, 0x40,
	0x14, 0x85, 0x37, 0x6a, 0xcc, 0x7a, 0xc5, 0x25, 0x5c, 0x96, 0x10, 0x16, 0x91, 0x20, 0xfb, 0x10,
	0xf6, 0x41, 0x16, 0xfb, 0xda, 0x17, 0x6d, 0x53, 0x1a, 0xda, 0x6a, 0x49, 0xda, 0xfa, 0x3c, 0x9a,
	0x8b, 0x0c, 0x98, 0x99, 0x90, 0x4c, 0x20, 0xf6, 0x27, 0xf4, 0x0f, 0xf5, 0xef, 0x95, 0x0c, 0x92,
	0xc6, 0xd2, 0xb7, 0x9c, 0xef, 0x5c, 0x0e, 0xe7, 0x90, 0x81, 0xf1, 0xad, 0x2c, 0x0b, 0xba, 0x92,
	0x42, 0xe5, 0x6c, 0xa7, 0xee, 0x79, 0xa1, 0x42, 0x91, 0x95, 0x6a, 0x96, 0xe5, 0x52, 0x49, 0xb4,
	0xa8, 0x62, 0x69, 0x76, 0xa0, 0xe9, 0xbb, 0x01, 0x83, 0x25, 0x2b, 0x48, 0x9b, 0x38, 0x01, 0x58,
	0x64, 0xd9, 0x0b, 0xe5, 0x05, 0x97, 0xc2, 0x35, 0x3c, 0xc3, 0x1f, 0x44, 0x2d, 0x82, 0x97, 0x30,
	0x8a, 0x65, 0x99, 0xef, 0x68, 0xc3, 0x8e, 0x81, 0x28, 0x53, 0xb7, 0xe3, 0x19, 0xfe, 0xaf, 0xb9,
	0x33, 0x3b, 0xc5, 0xcd, 0xce, 0xdc, 0xe8, 0xfc, 0x18, 0x3d, 0x18, 0x3e, 0x51, 0x9e, 0x72, 0xc1,
	0x0e, 0x41, 0xa5, 0xdc, 0xae, 0x8e, 0x6f, 0x23, 0x44, 0xe8, 0x85, 0x0f, 0x41, 0xe8, 0xf6, 0xb4,
	0xa5, 0xbf, 0xd1, 0x81, 0xfe, 0x3a, 0x23, 0x11, 0x26, 0xae, 0xa9, 0xe9, 0x49, 0x4d, 0xdf, 0x3a,
	0xe0, 0x7c, 0xbf, 0x11, 0xff, 0xb7, 0x36, 0xe9, 0x15, 0xc3, 0x39, 0x36, 0x15, 0x1b, 0x27, 0x6a,
	0x0d, 0x1f, 0xc3, 0xe0, 0x91, 0xed, 0x29, 0x14, 0x09, 0x55, 0x7a, 0x94, 0x19, 0x7d, 0x02, 0xfc,
	0x03, 0x3f, 0x6b, 0x11, 0xf3, 0x57, 0xd2, 0xad, 0xcd, 0xa8, 0xd1, 0xb5, 0x17, 0x2b, 0xa6, 0xca,
	0x22, 0x4c, 0x74, 0x6d, 0x33, 0x6a, 0x74, 0x3d, 0xf8, 0x8e, 0x8e, 0x1b, 0x99, 0x27, 0x2b, 0x96,
	0xd2, 0xa9, 0x7f, 0x1b, 0xe1, 0x5f, 0x18, 0x3d, 0x0b, 0x5e, 0x2d, 0x69, 0xcf, 0xc5, 0x35, 0x53,
	0xe4, 0xf6, 0x3d, 0xc3, 0xef, 0x46, 0xe7, 0xb0, 0xce, 0xa9, 0x41, 0x20, 0x12, 0x7d, 0x63, 0xe9,
	0x9b, 0x36, 0xfa, 0x77, 0xf3, 0xe5, 0xc7, 0xa0, 0x03, 0xb8, 0x92, 0x31, 0x29, 0xc5, 0xc5, 0xbe,
	0x71, 0xec, 0x1f, 0x68, 0x41, 0x37, 0x5c, 0xc7, 0xf6, 0x6f, 0x1c, 0x82, 0xb5, 0x10, 0x49, 0x2e,
	0x79, 0x62, 0x4f, 0x6a, 0xba, 0xa1, 0xad, 0xed, 0x6f, 0xfb, 0xfa, 0x79, 0x5c, 0x7c, 0x04, 0x00,
	0x00, 0xff, 0xff, 0x15, 0xda, 0x82, 0xf0, 0x3e, 0x02, 0x00, 0x00,
}
