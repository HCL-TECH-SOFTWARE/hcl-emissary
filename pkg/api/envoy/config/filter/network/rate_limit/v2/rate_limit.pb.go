// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/filter/network/rate_limit/v2/rate_limit.proto

package envoy_config_filter_network_rate_limit_v2

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	ratelimit "github.com/datawire/ambassador/pkg/api/envoy/api/v2/ratelimit"
	v2 "github.com/datawire/ambassador/pkg/api/envoy/config/ratelimit/v2"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
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

type RateLimit struct {
	StatPrefix           string                           `protobuf:"bytes,1,opt,name=stat_prefix,json=statPrefix,proto3" json:"stat_prefix,omitempty"`
	Domain               string                           `protobuf:"bytes,2,opt,name=domain,proto3" json:"domain,omitempty"`
	Descriptors          []*ratelimit.RateLimitDescriptor `protobuf:"bytes,3,rep,name=descriptors,proto3" json:"descriptors,omitempty"`
	Timeout              *duration.Duration               `protobuf:"bytes,4,opt,name=timeout,proto3" json:"timeout,omitempty"`
	FailureModeDeny      bool                             `protobuf:"varint,5,opt,name=failure_mode_deny,json=failureModeDeny,proto3" json:"failure_mode_deny,omitempty"`
	RateLimitService     *v2.RateLimitServiceConfig       `protobuf:"bytes,6,opt,name=rate_limit_service,json=rateLimitService,proto3" json:"rate_limit_service,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *RateLimit) Reset()         { *m = RateLimit{} }
func (m *RateLimit) String() string { return proto.CompactTextString(m) }
func (*RateLimit) ProtoMessage()    {}
func (*RateLimit) Descriptor() ([]byte, []int) {
	return fileDescriptor_34e9a222968daa71, []int{0}
}

func (m *RateLimit) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RateLimit.Unmarshal(m, b)
}
func (m *RateLimit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RateLimit.Marshal(b, m, deterministic)
}
func (m *RateLimit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RateLimit.Merge(m, src)
}
func (m *RateLimit) XXX_Size() int {
	return xxx_messageInfo_RateLimit.Size(m)
}
func (m *RateLimit) XXX_DiscardUnknown() {
	xxx_messageInfo_RateLimit.DiscardUnknown(m)
}

var xxx_messageInfo_RateLimit proto.InternalMessageInfo

func (m *RateLimit) GetStatPrefix() string {
	if m != nil {
		return m.StatPrefix
	}
	return ""
}

func (m *RateLimit) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

func (m *RateLimit) GetDescriptors() []*ratelimit.RateLimitDescriptor {
	if m != nil {
		return m.Descriptors
	}
	return nil
}

func (m *RateLimit) GetTimeout() *duration.Duration {
	if m != nil {
		return m.Timeout
	}
	return nil
}

func (m *RateLimit) GetFailureModeDeny() bool {
	if m != nil {
		return m.FailureModeDeny
	}
	return false
}

func (m *RateLimit) GetRateLimitService() *v2.RateLimitServiceConfig {
	if m != nil {
		return m.RateLimitService
	}
	return nil
}

func init() {
	proto.RegisterType((*RateLimit)(nil), "envoy.config.filter.network.rate_limit.v2.RateLimit")
}

func init() {
	proto.RegisterFile("envoy/config/filter/network/rate_limit/v2/rate_limit.proto", fileDescriptor_34e9a222968daa71)
}

var fileDescriptor_34e9a222968daa71 = []byte{
	// 464 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0x3f, 0x8f, 0xd3, 0x30,
	0x14, 0x97, 0x73, 0xd7, 0x5e, 0xcf, 0x95, 0xa0, 0x78, 0x21, 0x9c, 0x04, 0x44, 0x20, 0xa1, 0x00,
	0xc2, 0x16, 0xe9, 0x80, 0xc4, 0x18, 0x3a, 0x82, 0x54, 0x85, 0x81, 0x31, 0xf2, 0x35, 0x2f, 0x95,
	0x45, 0x62, 0x47, 0x8e, 0x13, 0xda, 0x8d, 0x11, 0xb1, 0x20, 0x31, 0xf1, 0x59, 0xf8, 0x04, 0xac,
	0x7c, 0x15, 0xc6, 0x43, 0x42, 0x28, 0x71, 0xd2, 0x34, 0xb0, 0xdc, 0x66, 0xfb, 0xf7, 0xe7, 0xbd,
	0xf7, 0xf3, 0xc3, 0x2f, 0x41, 0xd6, 0x6a, 0xcf, 0x36, 0x4a, 0xa6, 0x62, 0xcb, 0x52, 0x91, 0x19,
	0xd0, 0x4c, 0x82, 0xf9, 0xa0, 0xf4, 0x7b, 0xa6, 0xb9, 0x81, 0x38, 0x13, 0xb9, 0x30, 0xac, 0x0e,
	0x8e, 0x6e, 0xb4, 0xd0, 0xca, 0x28, 0xf2, 0xb8, 0xd5, 0x52, 0xab, 0xa5, 0x56, 0x4b, 0x3b, 0x2d,
	0x3d, 0x62, 0xd7, 0xc1, 0xc5, 0x23, 0x5b, 0x86, 0x17, 0xa2, 0x77, 0xb2, 0xb6, 0x87, 0x93, 0xb5,
	0xbc, 0x78, 0x38, 0x6a, 0x67, 0xe0, 0x35, 0xa2, 0xac, 0xec, 0x48, 0xf7, 0xb6, 0x4a, 0x6d, 0x33,
	0x60, 0xed, 0xed, 0xb2, 0x4a, 0x59, 0x52, 0x69, 0x6e, 0x84, 0x92, 0x3d, 0x5e, 0x25, 0x05, 0x67,
	0x5c, 0x4a, 0x65, 0xda, 0xe7, 0x92, 0xe5, 0x62, 0xdb, 0x78, 0x75, 0xf8, 0xdd, 0xff, 0xf0, 0xd2,
	0x70, 0x53, 0xf5, 0xf6, 0xb7, 0x6b, 0x9e, 0x89, 0x84, 0x1b, 0x60, 0xfd, 0xc1, 0x02, 0x0f, 0x7e,
	0x3b, 0xf8, 0x3c, 0xe2, 0x06, 0x5e, 0x37, 0x2d, 0x11, 0x1f, 0xcf, 0x1b, 0x59, 0x5c, 0x68, 0x48,
	0xc5, 0xce, 0x45, 0x1e, 0xf2, 0xcf, 0xc3, 0xb3, 0xab, 0xf0, 0x54, 0x3b, 0x1e, 0x8a, 0x70, 0x83,
	0xad, 0x5b, 0x88, 0xdc, 0xc7, 0xd3, 0x44, 0xe5, 0x5c, 0x48, 0xd7, 0x19, 0x93, 0xba, 0x67, 0xf2,
	0x0e, 0xcf, 0x13, 0x28, 0x37, 0x5a, 0x14, 0x46, 0xe9, 0xd2, 0x3d, 0xf1, 0x4e, 0xfc, 0x79, 0xf0,
	0x94, 0xda, 0x78, 0x79, 0x21, 0x68, 0x1d, 0xd0, 0x21, 0xa9, 0x43, 0x0b, 0xab, 0x83, 0x26, 0x9c,
	0x5d, 0x85, 0x93, 0xaf, 0xc8, 0x99, 0xa1, 0xe8, 0xd8, 0x89, 0x2c, 0xf1, 0x99, 0x11, 0x39, 0xa8,
	0xca, 0xb8, 0xa7, 0x1e, 0xf2, 0xe7, 0xc1, 0x1d, 0x6a, 0xb3, 0xa3, 0x7d, 0x76, 0x74, 0xd5, 0x65,
	0x17, 0xf5, 0x4c, 0xf2, 0x04, 0xdf, 0x4a, 0xb9, 0xc8, 0x2a, 0x0d, 0x71, 0xae, 0x12, 0x88, 0x13,
	0x90, 0x7b, 0x77, 0xe2, 0x21, 0x7f, 0x16, 0xdd, 0xec, 0x80, 0x37, 0x2a, 0x81, 0x15, 0xc8, 0x3d,
	0x11, 0x98, 0x0c, 0x1f, 0x1d, 0x97, 0xa0, 0x6b, 0xb1, 0x01, 0x77, 0xda, 0xd6, 0x7a, 0x4e, 0x47,
	0xfb, 0x31, 0x0c, 0x50, 0x07, 0xc3, 0x0c, 0x6f, 0xad, 0xe4, 0x55, 0xcb, 0x69, 0xc7, 0xf8, 0x8c,
	0x9c, 0x05, 0x8a, 0x16, 0xfa, 0x1f, 0x46, 0xf8, 0x09, 0xfd, 0xfa, 0xf6, 0xe7, 0xcb, 0x84, 0x91,
	0x67, 0xd6, 0x16, 0x76, 0x06, 0x64, 0xd9, 0x7c, 0x5f, 0xb7, 0x7a, 0xe5, 0x68, 0xf7, 0xba, 0x52,
	0xcb, 0xef, 0x1f, 0x7f, 0xfc, 0x9c, 0x3a, 0x0b, 0x84, 0x5f, 0x08, 0x65, 0x1b, 0x2a, 0xb4, 0xda,
	0xed, 0xe9, 0xb5, 0x77, 0x37, 0xbc, 0x71, 0x68, 0x76, 0xdd, 0x84, 0xb6, 0x46, 0x97, 0xd3, 0x36,
	0xbd, 0xe5, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x7d, 0xf8, 0x35, 0x5d, 0x3d, 0x03, 0x00, 0x00,
}
