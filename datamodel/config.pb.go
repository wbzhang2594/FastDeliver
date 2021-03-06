// Code generated by protoc-gen-go.
// source: config.proto
// DO NOT EDIT!

/*
Package datamodel is a generated protocol buffer package.

It is generated from these files:
	config.proto
	models.proto

It has these top-level messages:
	ClientInfo
	FileDeliverTask
	MessageAck
	KeepAlive
	StartSendFile
*/
package datamodel

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ClientInfo struct {
	Type             *uint32 `protobuf:"varint,1,req,name=Type" json:"Type,omitempty"`
	ClientId         *string `protobuf:"bytes,2,req,name=ClientId" json:"ClientId,omitempty"`
	CmdReqPort       *uint32 `protobuf:"varint,3,req,name=CmdReqPort,def=9000" json:"CmdReqPort,omitempty"`
	DataPullPort     *uint32 `protobuf:"varint,4,req,name=DataPullPort,def=10000" json:"DataPullPort,omitempty"`
	DataReqPort      *uint32 `protobuf:"varint,5,req,name=DataReqPort,def=10100" json:"DataReqPort,omitempty"`
	Ip               *string `protobuf:"bytes,6,req,name=Ip" json:"Ip,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ClientInfo) Reset()                    { *m = ClientInfo{} }
func (m *ClientInfo) String() string            { return proto.CompactTextString(m) }
func (*ClientInfo) ProtoMessage()               {}
func (*ClientInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

const Default_ClientInfo_CmdReqPort uint32 = 9000
const Default_ClientInfo_DataPullPort uint32 = 10000
const Default_ClientInfo_DataReqPort uint32 = 10100

func (m *ClientInfo) GetType() uint32 {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return 0
}

func (m *ClientInfo) GetClientId() string {
	if m != nil && m.ClientId != nil {
		return *m.ClientId
	}
	return ""
}

func (m *ClientInfo) GetCmdReqPort() uint32 {
	if m != nil && m.CmdReqPort != nil {
		return *m.CmdReqPort
	}
	return Default_ClientInfo_CmdReqPort
}

func (m *ClientInfo) GetDataPullPort() uint32 {
	if m != nil && m.DataPullPort != nil {
		return *m.DataPullPort
	}
	return Default_ClientInfo_DataPullPort
}

func (m *ClientInfo) GetDataReqPort() uint32 {
	if m != nil && m.DataReqPort != nil {
		return *m.DataReqPort
	}
	return Default_ClientInfo_DataReqPort
}

func (m *ClientInfo) GetIp() string {
	if m != nil && m.Ip != nil {
		return *m.Ip
	}
	return ""
}

func init() {
	proto.RegisterType((*ClientInfo)(nil), "datamodel.ClientInfo")
}

var fileDescriptor0 = []byte{
	// 155 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x49, 0xce, 0xcf, 0x4b,
	0xcb, 0x4c, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x4c, 0x49, 0x2c, 0x49, 0xcc, 0xcd,
	0x4f, 0x49, 0xcd, 0x51, 0xea, 0x66, 0xe4, 0xe2, 0x72, 0xce, 0xc9, 0x4c, 0xcd, 0x2b, 0xf1, 0xcc,
	0x4b, 0xcb, 0x17, 0xe2, 0xe1, 0x62, 0x09, 0xa9, 0x2c, 0x48, 0x95, 0x60, 0x54, 0x60, 0xd2, 0xe0,
	0x15, 0x12, 0xe0, 0xe2, 0x80, 0xca, 0xa5, 0x48, 0x30, 0x01, 0x45, 0x38, 0x85, 0x24, 0x80, 0xaa,
	0x73, 0x53, 0x82, 0x52, 0x0b, 0x03, 0xf2, 0x8b, 0x4a, 0x24, 0x98, 0x41, 0xaa, 0xac, 0x58, 0x2c,
	0x0d, 0x0c, 0x0c, 0x84, 0xa4, 0xb9, 0x78, 0x5c, 0x80, 0xa6, 0x06, 0x94, 0xe6, 0xe4, 0x80, 0xe5,
	0x58, 0xc0, 0x72, 0xac, 0x86, 0x06, 0x20, 0x49, 0x29, 0x2e, 0x6e, 0x90, 0x24, 0x4c, 0x1f, 0x2b,
	0x4c, 0x0e, 0x28, 0x2b, 0xc4, 0xc5, 0xc5, 0xe4, 0x59, 0x20, 0xc1, 0x06, 0x32, 0x1e, 0x10, 0x00,
	0x00, 0xff, 0xff, 0x79, 0x66, 0x4e, 0x6b, 0xa7, 0x00, 0x00, 0x00,
}
