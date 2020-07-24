// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/filter/network/ext_authz/v2/ext_authz.proto

package envoy_config_filter_network_ext_authz_v2

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	core "github.com/datawire/ambassador/pkg/api/envoy/api/v2/core"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

type ExtAuthz struct {
	StatPrefix             string            `protobuf:"bytes,1,opt,name=stat_prefix,json=statPrefix,proto3" json:"stat_prefix,omitempty"`
	GrpcService            *core.GrpcService `protobuf:"bytes,2,opt,name=grpc_service,json=grpcService,proto3" json:"grpc_service,omitempty"`
	FailureModeAllow       bool              `protobuf:"varint,3,opt,name=failure_mode_allow,json=failureModeAllow,proto3" json:"failure_mode_allow,omitempty"`
	IncludePeerCertificate bool              `protobuf:"varint,4,opt,name=include_peer_certificate,json=includePeerCertificate,proto3" json:"include_peer_certificate,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}          `json:"-"`
	XXX_unrecognized       []byte            `json:"-"`
	XXX_sizecache          int32             `json:"-"`
}

func (m *ExtAuthz) Reset()         { *m = ExtAuthz{} }
func (m *ExtAuthz) String() string { return proto.CompactTextString(m) }
func (*ExtAuthz) ProtoMessage()    {}
func (*ExtAuthz) Descriptor() ([]byte, []int) {
	return fileDescriptor_3ec2615c2696024a, []int{0}
}

func (m *ExtAuthz) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExtAuthz.Unmarshal(m, b)
}
func (m *ExtAuthz) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExtAuthz.Marshal(b, m, deterministic)
}
func (m *ExtAuthz) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExtAuthz.Merge(m, src)
}
func (m *ExtAuthz) XXX_Size() int {
	return xxx_messageInfo_ExtAuthz.Size(m)
}
func (m *ExtAuthz) XXX_DiscardUnknown() {
	xxx_messageInfo_ExtAuthz.DiscardUnknown(m)
}

var xxx_messageInfo_ExtAuthz proto.InternalMessageInfo

func (m *ExtAuthz) GetStatPrefix() string {
	if m != nil {
		return m.StatPrefix
	}
	return ""
}

func (m *ExtAuthz) GetGrpcService() *core.GrpcService {
	if m != nil {
		return m.GrpcService
	}
	return nil
}

func (m *ExtAuthz) GetFailureModeAllow() bool {
	if m != nil {
		return m.FailureModeAllow
	}
	return false
}

func (m *ExtAuthz) GetIncludePeerCertificate() bool {
	if m != nil {
		return m.IncludePeerCertificate
	}
	return false
}

func init() {
	proto.RegisterType((*ExtAuthz)(nil), "envoy.config.filter.network.ext_authz.v2.ExtAuthz")
}

func init() {
	proto.RegisterFile("envoy/config/filter/network/ext_authz/v2/ext_authz.proto", fileDescriptor_3ec2615c2696024a)
}

var fileDescriptor_3ec2615c2696024a = []byte{
	// 384 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xb1, 0xee, 0xd3, 0x30,
	0x10, 0xc6, 0xe5, 0x52, 0x4a, 0x71, 0x41, 0xaa, 0x32, 0x40, 0x54, 0x89, 0x2a, 0x42, 0x0c, 0x19,
	0xc0, 0x96, 0x52, 0x09, 0x75, 0x6d, 0x10, 0x62, 0x42, 0x8a, 0xca, 0x03, 0x44, 0xc6, 0xb9, 0x04,
	0x8b, 0xd4, 0xb6, 0x1c, 0x27, 0x4d, 0x99, 0x98, 0x58, 0x59, 0x79, 0x16, 0x9e, 0x80, 0x95, 0x27,
	0xe0, 0x1d, 0x18, 0x19, 0x10, 0x72, 0x9c, 0x2a, 0x48, 0x30, 0xfc, 0xb7, 0x5c, 0x7e, 0xf7, 0x9d,
	0xef, 0xbe, 0x0f, 0xef, 0x41, 0x76, 0xea, 0x42, 0xb9, 0x92, 0xa5, 0xa8, 0x68, 0x29, 0x6a, 0x0b,
	0x86, 0x4a, 0xb0, 0x67, 0x65, 0xde, 0x53, 0xe8, 0x6d, 0xce, 0x5a, 0xfb, 0xee, 0x03, 0xed, 0x92,
	0xa9, 0x20, 0xda, 0x28, 0xab, 0x82, 0x78, 0x50, 0x12, 0xaf, 0x24, 0x5e, 0x49, 0x46, 0x25, 0x99,
	0x9a, 0xbb, 0x64, 0xf3, 0xc4, 0xbf, 0xc1, 0xb4, 0x70, 0x73, 0xb8, 0x32, 0x40, 0x2b, 0xa3, 0x79,
	0xde, 0x80, 0xe9, 0x04, 0x07, 0x3f, 0x6f, 0xb3, 0x6d, 0x0b, 0xcd, 0x28, 0x93, 0x52, 0x59, 0x66,
	0x85, 0x92, 0x0d, 0x3d, 0x89, 0xca, 0x30, 0x7b, 0xe5, 0x8f, 0xfe, 0xe1, 0x8d, 0x65, 0xb6, 0x6d,
	0x46, 0xfc, 0xb0, 0x63, 0xb5, 0x28, 0x98, 0x05, 0x7a, 0xfd, 0xf0, 0xe0, 0xf1, 0x0f, 0x84, 0x97,
	0x2f, 0x7b, 0x7b, 0x70, 0xdb, 0x04, 0x31, 0x5e, 0x39, 0x55, 0xae, 0x0d, 0x94, 0xa2, 0x0f, 0x51,
	0x84, 0xe2, 0xbb, 0xe9, 0x9d, 0x5f, 0xe9, 0xdc, 0xcc, 0x22, 0x74, 0xc4, 0x8e, 0x65, 0x03, 0x0a,
	0x0e, 0xf8, 0xde, 0xdf, 0x4b, 0x86, 0xb3, 0x08, 0xc5, 0xab, 0x64, 0x4b, 0xfc, 0xd5, 0x4c, 0x0b,
	0xd2, 0x25, 0xc4, 0xdd, 0x42, 0x5e, 0x19, 0xcd, 0xdf, 0xf8, 0xae, 0xe3, 0xaa, 0x9a, 0x8a, 0xe0,
	0x29, 0x0e, 0x4a, 0x26, 0xea, 0xd6, 0x40, 0x7e, 0x52, 0x05, 0xe4, 0xac, 0xae, 0xd5, 0x39, 0xbc,
	0x15, 0xa1, 0x78, 0x79, 0x5c, 0x8f, 0xe4, 0xb5, 0x2a, 0xe0, 0xe0, 0xfe, 0x07, 0x7b, 0x1c, 0x0a,
	0xc9, 0xeb, 0xb6, 0x80, 0x5c, 0x03, 0x98, 0x9c, 0x83, 0xb1, 0xa2, 0x14, 0x9c, 0x59, 0x08, 0xe7,
	0x83, 0xe6, 0xc1, 0xc8, 0x33, 0x00, 0xf3, 0x62, 0xa2, 0xe9, 0x27, 0xf4, 0xf3, 0xcb, 0xef, 0xcf,
	0xb7, 0x69, 0xf0, 0xcc, 0x2f, 0x07, 0xbd, 0x05, 0xd9, 0x38, 0x8b, 0xc6, 0x58, 0x9a, 0xff, 0xe5,
	0xb2, 0xfb, 0xfa, 0xf1, 0xdb, 0xf7, 0xc5, 0x6c, 0x8d, 0xf0, 0x73, 0xa1, 0xfc, 0x59, 0xda, 0xa8,
	0xfe, 0x42, 0x6e, 0x9a, 0x6b, 0x7a, 0xff, 0x6a, 0x6b, 0xe6, 0x8c, 0xce, 0xd0, 0xdb, 0xc5, 0xe0,
	0xf8, 0xee, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xeb, 0x56, 0xa0, 0x9b, 0x55, 0x02, 0x00, 0x00,
}
