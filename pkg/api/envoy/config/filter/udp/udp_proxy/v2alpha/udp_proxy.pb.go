// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/filter/udp/udp_proxy/v2alpha/udp_proxy.proto

package envoy_config_filter_udp_udp_proxy_v2alpha

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
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

type UdpProxyConfig struct {
	StatPrefix string `protobuf:"bytes,1,opt,name=stat_prefix,json=statPrefix,proto3" json:"stat_prefix,omitempty"`
	// Types that are valid to be assigned to RouteSpecifier:
	//	*UdpProxyConfig_Cluster
	RouteSpecifier       isUdpProxyConfig_RouteSpecifier `protobuf_oneof:"route_specifier"`
	IdleTimeout          *duration.Duration              `protobuf:"bytes,3,opt,name=idle_timeout,json=idleTimeout,proto3" json:"idle_timeout,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *UdpProxyConfig) Reset()         { *m = UdpProxyConfig{} }
func (m *UdpProxyConfig) String() string { return proto.CompactTextString(m) }
func (*UdpProxyConfig) ProtoMessage()    {}
func (*UdpProxyConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_82ee653a514c2977, []int{0}
}

func (m *UdpProxyConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UdpProxyConfig.Unmarshal(m, b)
}
func (m *UdpProxyConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UdpProxyConfig.Marshal(b, m, deterministic)
}
func (m *UdpProxyConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UdpProxyConfig.Merge(m, src)
}
func (m *UdpProxyConfig) XXX_Size() int {
	return xxx_messageInfo_UdpProxyConfig.Size(m)
}
func (m *UdpProxyConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_UdpProxyConfig.DiscardUnknown(m)
}

var xxx_messageInfo_UdpProxyConfig proto.InternalMessageInfo

func (m *UdpProxyConfig) GetStatPrefix() string {
	if m != nil {
		return m.StatPrefix
	}
	return ""
}

type isUdpProxyConfig_RouteSpecifier interface {
	isUdpProxyConfig_RouteSpecifier()
}

type UdpProxyConfig_Cluster struct {
	Cluster string `protobuf:"bytes,2,opt,name=cluster,proto3,oneof"`
}

func (*UdpProxyConfig_Cluster) isUdpProxyConfig_RouteSpecifier() {}

func (m *UdpProxyConfig) GetRouteSpecifier() isUdpProxyConfig_RouteSpecifier {
	if m != nil {
		return m.RouteSpecifier
	}
	return nil
}

func (m *UdpProxyConfig) GetCluster() string {
	if x, ok := m.GetRouteSpecifier().(*UdpProxyConfig_Cluster); ok {
		return x.Cluster
	}
	return ""
}

func (m *UdpProxyConfig) GetIdleTimeout() *duration.Duration {
	if m != nil {
		return m.IdleTimeout
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*UdpProxyConfig) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*UdpProxyConfig_Cluster)(nil),
	}
}

func init() {
	proto.RegisterType((*UdpProxyConfig)(nil), "envoy.config.filter.udp.udp_proxy.v2alpha.UdpProxyConfig")
}

func init() {
	proto.RegisterFile("envoy/config/filter/udp/udp_proxy/v2alpha/udp_proxy.proto", fileDescriptor_82ee653a514c2977)
}

var fileDescriptor_82ee653a514c2977 = []byte{
	// 316 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0xcf, 0x4a, 0x33, 0x31,
	0x14, 0xc5, 0xbf, 0x4c, 0x3f, 0x5a, 0x4c, 0xfd, 0xc7, 0x2c, 0xb4, 0x16, 0x94, 0xa2, 0x9b, 0xba,
	0x49, 0xa0, 0x2e, 0x44, 0x70, 0x35, 0xba, 0x70, 0x39, 0x0c, 0xba, 0x1e, 0xd2, 0x26, 0x53, 0x03,
	0x71, 0x6e, 0xc8, 0x24, 0xa5, 0xdd, 0xf9, 0x52, 0x6e, 0x7c, 0x02, 0xb7, 0xbe, 0x8d, 0xb8, 0x92,
	0x24, 0x33, 0x14, 0x77, 0x2e, 0x02, 0xe1, 0x9e, 0xf3, 0x4b, 0xee, 0x39, 0xf8, 0x46, 0xd4, 0x2b,
	0xd8, 0xd0, 0x05, 0xd4, 0x95, 0x5c, 0xd2, 0x4a, 0x2a, 0x2b, 0x0c, 0x75, 0x5c, 0xfb, 0x53, 0x6a,
	0x03, 0xeb, 0x0d, 0x5d, 0xcd, 0x98, 0xd2, 0xcf, 0x6c, 0x3b, 0x21, 0xda, 0x80, 0x85, 0xf4, 0x32,
	0xa0, 0x24, 0xa2, 0x24, 0xa2, 0xc4, 0x71, 0x4d, 0xb6, 0xc6, 0x16, 0x1d, 0x9f, 0x2d, 0x01, 0x96,
	0x4a, 0xd0, 0x00, 0xce, 0x5d, 0x45, 0xb9, 0x33, 0xcc, 0x4a, 0xa8, 0xe3, 0x53, 0xe3, 0x53, 0xc7,
	0x35, 0xa3, 0xac, 0xae, 0xc1, 0x86, 0x71, 0x43, 0x1b, 0xcb, 0xac, 0x6b, 0x5a, 0xf9, 0x78, 0xc5,
	0x94, 0xe4, 0xcc, 0x0a, 0xda, 0x5d, 0xa2, 0x70, 0xfe, 0x86, 0xf0, 0xfe, 0x13, 0xd7, 0xb9, 0xff,
	0xec, 0x2e, 0xec, 0x91, 0x4e, 0xf1, 0xd0, 0xb3, 0xa5, 0x36, 0xa2, 0x92, 0xeb, 0x11, 0x9a, 0xa0,
	0xe9, 0x4e, 0x36, 0xf8, 0xce, 0xfe, 0x9b, 0x64, 0x82, 0x0a, 0xec, 0xb5, 0x3c, 0x48, 0xe9, 0x05,
	0x1e, 0x2c, 0x94, 0x6b, 0xac, 0x30, 0xa3, 0xe4, 0x97, 0xeb, 0xe1, 0x5f, 0xd1, 0x29, 0xe9, 0x2d,
	0xde, 0x95, 0x5c, 0x89, 0xd2, 0xca, 0x17, 0x01, 0xce, 0x8e, 0x7a, 0x13, 0x34, 0x1d, 0xce, 0x4e,
	0x48, 0x0c, 0x44, 0xba, 0x40, 0xe4, 0xbe, 0x0d, 0x54, 0x0c, 0xbd, 0xfd, 0x31, 0xba, 0xb3, 0x23,
	0x7c, 0x60, 0xc0, 0x59, 0x51, 0x36, 0x5a, 0x2c, 0x64, 0x25, 0x85, 0x49, 0x7b, 0x5f, 0x19, 0xca,
	0x8a, 0xf7, 0xd7, 0x8f, 0xcf, 0x7e, 0x72, 0x98, 0xe0, 0x6b, 0x09, 0x24, 0xf4, 0x18, 0x0b, 0xfb,
	0x73, 0xa5, 0xd9, 0x5e, 0x97, 0x3b, 0xf7, 0x2b, 0xe4, 0x68, 0xde, 0x0f, 0xbb, 0x5c, 0xfd, 0x04,
	0x00, 0x00, 0xff, 0xff, 0x2d, 0xd4, 0xf4, 0x81, 0xd2, 0x01, 0x00, 0x00,
}
