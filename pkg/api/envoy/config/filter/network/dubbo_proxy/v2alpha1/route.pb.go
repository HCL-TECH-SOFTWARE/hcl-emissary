// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/filter/network/dubbo_proxy/v2alpha1/route.proto

package envoy_config_filter_network_dubbo_proxy_v2alpha1

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	route "github.com/datawire/ambassador/pkg/api/envoy/api/v2/route"
	_type "github.com/datawire/ambassador/pkg/api/envoy/type"
	matcher "github.com/datawire/ambassador/pkg/api/envoy/type/matcher"
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

type RouteConfiguration struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Interface            string   `protobuf:"bytes,2,opt,name=interface,proto3" json:"interface,omitempty"`
	Group                string   `protobuf:"bytes,3,opt,name=group,proto3" json:"group,omitempty"`
	Version              string   `protobuf:"bytes,4,opt,name=version,proto3" json:"version,omitempty"`
	Routes               []*Route `protobuf:"bytes,5,rep,name=routes,proto3" json:"routes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RouteConfiguration) Reset()         { *m = RouteConfiguration{} }
func (m *RouteConfiguration) String() string { return proto.CompactTextString(m) }
func (*RouteConfiguration) ProtoMessage()    {}
func (*RouteConfiguration) Descriptor() ([]byte, []int) {
	return fileDescriptor_74a572433a3292e0, []int{0}
}

func (m *RouteConfiguration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RouteConfiguration.Unmarshal(m, b)
}
func (m *RouteConfiguration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RouteConfiguration.Marshal(b, m, deterministic)
}
func (m *RouteConfiguration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouteConfiguration.Merge(m, src)
}
func (m *RouteConfiguration) XXX_Size() int {
	return xxx_messageInfo_RouteConfiguration.Size(m)
}
func (m *RouteConfiguration) XXX_DiscardUnknown() {
	xxx_messageInfo_RouteConfiguration.DiscardUnknown(m)
}

var xxx_messageInfo_RouteConfiguration proto.InternalMessageInfo

func (m *RouteConfiguration) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RouteConfiguration) GetInterface() string {
	if m != nil {
		return m.Interface
	}
	return ""
}

func (m *RouteConfiguration) GetGroup() string {
	if m != nil {
		return m.Group
	}
	return ""
}

func (m *RouteConfiguration) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *RouteConfiguration) GetRoutes() []*Route {
	if m != nil {
		return m.Routes
	}
	return nil
}

type Route struct {
	Match                *RouteMatch  `protobuf:"bytes,1,opt,name=match,proto3" json:"match,omitempty"`
	Route                *RouteAction `protobuf:"bytes,2,opt,name=route,proto3" json:"route,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Route) Reset()         { *m = Route{} }
func (m *Route) String() string { return proto.CompactTextString(m) }
func (*Route) ProtoMessage()    {}
func (*Route) Descriptor() ([]byte, []int) {
	return fileDescriptor_74a572433a3292e0, []int{1}
}

func (m *Route) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Route.Unmarshal(m, b)
}
func (m *Route) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Route.Marshal(b, m, deterministic)
}
func (m *Route) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Route.Merge(m, src)
}
func (m *Route) XXX_Size() int {
	return xxx_messageInfo_Route.Size(m)
}
func (m *Route) XXX_DiscardUnknown() {
	xxx_messageInfo_Route.DiscardUnknown(m)
}

var xxx_messageInfo_Route proto.InternalMessageInfo

func (m *Route) GetMatch() *RouteMatch {
	if m != nil {
		return m.Match
	}
	return nil
}

func (m *Route) GetRoute() *RouteAction {
	if m != nil {
		return m.Route
	}
	return nil
}

type RouteMatch struct {
	Method               *MethodMatch           `protobuf:"bytes,1,opt,name=method,proto3" json:"method,omitempty"`
	Headers              []*route.HeaderMatcher `protobuf:"bytes,2,rep,name=headers,proto3" json:"headers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *RouteMatch) Reset()         { *m = RouteMatch{} }
func (m *RouteMatch) String() string { return proto.CompactTextString(m) }
func (*RouteMatch) ProtoMessage()    {}
func (*RouteMatch) Descriptor() ([]byte, []int) {
	return fileDescriptor_74a572433a3292e0, []int{2}
}

func (m *RouteMatch) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RouteMatch.Unmarshal(m, b)
}
func (m *RouteMatch) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RouteMatch.Marshal(b, m, deterministic)
}
func (m *RouteMatch) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouteMatch.Merge(m, src)
}
func (m *RouteMatch) XXX_Size() int {
	return xxx_messageInfo_RouteMatch.Size(m)
}
func (m *RouteMatch) XXX_DiscardUnknown() {
	xxx_messageInfo_RouteMatch.DiscardUnknown(m)
}

var xxx_messageInfo_RouteMatch proto.InternalMessageInfo

func (m *RouteMatch) GetMethod() *MethodMatch {
	if m != nil {
		return m.Method
	}
	return nil
}

func (m *RouteMatch) GetHeaders() []*route.HeaderMatcher {
	if m != nil {
		return m.Headers
	}
	return nil
}

type RouteAction struct {
	// Types that are valid to be assigned to ClusterSpecifier:
	//	*RouteAction_Cluster
	//	*RouteAction_WeightedClusters
	ClusterSpecifier     isRouteAction_ClusterSpecifier `protobuf_oneof:"cluster_specifier"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *RouteAction) Reset()         { *m = RouteAction{} }
func (m *RouteAction) String() string { return proto.CompactTextString(m) }
func (*RouteAction) ProtoMessage()    {}
func (*RouteAction) Descriptor() ([]byte, []int) {
	return fileDescriptor_74a572433a3292e0, []int{3}
}

func (m *RouteAction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RouteAction.Unmarshal(m, b)
}
func (m *RouteAction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RouteAction.Marshal(b, m, deterministic)
}
func (m *RouteAction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouteAction.Merge(m, src)
}
func (m *RouteAction) XXX_Size() int {
	return xxx_messageInfo_RouteAction.Size(m)
}
func (m *RouteAction) XXX_DiscardUnknown() {
	xxx_messageInfo_RouteAction.DiscardUnknown(m)
}

var xxx_messageInfo_RouteAction proto.InternalMessageInfo

type isRouteAction_ClusterSpecifier interface {
	isRouteAction_ClusterSpecifier()
}

type RouteAction_Cluster struct {
	Cluster string `protobuf:"bytes,1,opt,name=cluster,proto3,oneof"`
}

type RouteAction_WeightedClusters struct {
	WeightedClusters *route.WeightedCluster `protobuf:"bytes,2,opt,name=weighted_clusters,json=weightedClusters,proto3,oneof"`
}

func (*RouteAction_Cluster) isRouteAction_ClusterSpecifier() {}

func (*RouteAction_WeightedClusters) isRouteAction_ClusterSpecifier() {}

func (m *RouteAction) GetClusterSpecifier() isRouteAction_ClusterSpecifier {
	if m != nil {
		return m.ClusterSpecifier
	}
	return nil
}

func (m *RouteAction) GetCluster() string {
	if x, ok := m.GetClusterSpecifier().(*RouteAction_Cluster); ok {
		return x.Cluster
	}
	return ""
}

func (m *RouteAction) GetWeightedClusters() *route.WeightedCluster {
	if x, ok := m.GetClusterSpecifier().(*RouteAction_WeightedClusters); ok {
		return x.WeightedClusters
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*RouteAction) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*RouteAction_Cluster)(nil),
		(*RouteAction_WeightedClusters)(nil),
	}
}

type MethodMatch struct {
	Name                 *matcher.StringMatcher                          `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	ParamsMatch          map[uint32]*MethodMatch_ParameterMatchSpecifier `protobuf:"bytes,2,rep,name=params_match,json=paramsMatch,proto3" json:"params_match,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                                        `json:"-"`
	XXX_unrecognized     []byte                                          `json:"-"`
	XXX_sizecache        int32                                           `json:"-"`
}

func (m *MethodMatch) Reset()         { *m = MethodMatch{} }
func (m *MethodMatch) String() string { return proto.CompactTextString(m) }
func (*MethodMatch) ProtoMessage()    {}
func (*MethodMatch) Descriptor() ([]byte, []int) {
	return fileDescriptor_74a572433a3292e0, []int{4}
}

func (m *MethodMatch) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MethodMatch.Unmarshal(m, b)
}
func (m *MethodMatch) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MethodMatch.Marshal(b, m, deterministic)
}
func (m *MethodMatch) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MethodMatch.Merge(m, src)
}
func (m *MethodMatch) XXX_Size() int {
	return xxx_messageInfo_MethodMatch.Size(m)
}
func (m *MethodMatch) XXX_DiscardUnknown() {
	xxx_messageInfo_MethodMatch.DiscardUnknown(m)
}

var xxx_messageInfo_MethodMatch proto.InternalMessageInfo

func (m *MethodMatch) GetName() *matcher.StringMatcher {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *MethodMatch) GetParamsMatch() map[uint32]*MethodMatch_ParameterMatchSpecifier {
	if m != nil {
		return m.ParamsMatch
	}
	return nil
}

type MethodMatch_ParameterMatchSpecifier struct {
	// Types that are valid to be assigned to ParameterMatchSpecifier:
	//	*MethodMatch_ParameterMatchSpecifier_ExactMatch
	//	*MethodMatch_ParameterMatchSpecifier_RangeMatch
	ParameterMatchSpecifier isMethodMatch_ParameterMatchSpecifier_ParameterMatchSpecifier `protobuf_oneof:"parameter_match_specifier"`
	XXX_NoUnkeyedLiteral    struct{}                                                      `json:"-"`
	XXX_unrecognized        []byte                                                        `json:"-"`
	XXX_sizecache           int32                                                         `json:"-"`
}

func (m *MethodMatch_ParameterMatchSpecifier) Reset()         { *m = MethodMatch_ParameterMatchSpecifier{} }
func (m *MethodMatch_ParameterMatchSpecifier) String() string { return proto.CompactTextString(m) }
func (*MethodMatch_ParameterMatchSpecifier) ProtoMessage()    {}
func (*MethodMatch_ParameterMatchSpecifier) Descriptor() ([]byte, []int) {
	return fileDescriptor_74a572433a3292e0, []int{4, 0}
}

func (m *MethodMatch_ParameterMatchSpecifier) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MethodMatch_ParameterMatchSpecifier.Unmarshal(m, b)
}
func (m *MethodMatch_ParameterMatchSpecifier) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MethodMatch_ParameterMatchSpecifier.Marshal(b, m, deterministic)
}
func (m *MethodMatch_ParameterMatchSpecifier) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MethodMatch_ParameterMatchSpecifier.Merge(m, src)
}
func (m *MethodMatch_ParameterMatchSpecifier) XXX_Size() int {
	return xxx_messageInfo_MethodMatch_ParameterMatchSpecifier.Size(m)
}
func (m *MethodMatch_ParameterMatchSpecifier) XXX_DiscardUnknown() {
	xxx_messageInfo_MethodMatch_ParameterMatchSpecifier.DiscardUnknown(m)
}

var xxx_messageInfo_MethodMatch_ParameterMatchSpecifier proto.InternalMessageInfo

type isMethodMatch_ParameterMatchSpecifier_ParameterMatchSpecifier interface {
	isMethodMatch_ParameterMatchSpecifier_ParameterMatchSpecifier()
}

type MethodMatch_ParameterMatchSpecifier_ExactMatch struct {
	ExactMatch string `protobuf:"bytes,3,opt,name=exact_match,json=exactMatch,proto3,oneof"`
}

type MethodMatch_ParameterMatchSpecifier_RangeMatch struct {
	RangeMatch *_type.Int64Range `protobuf:"bytes,4,opt,name=range_match,json=rangeMatch,proto3,oneof"`
}

func (*MethodMatch_ParameterMatchSpecifier_ExactMatch) isMethodMatch_ParameterMatchSpecifier_ParameterMatchSpecifier() {
}

func (*MethodMatch_ParameterMatchSpecifier_RangeMatch) isMethodMatch_ParameterMatchSpecifier_ParameterMatchSpecifier() {
}

func (m *MethodMatch_ParameterMatchSpecifier) GetParameterMatchSpecifier() isMethodMatch_ParameterMatchSpecifier_ParameterMatchSpecifier {
	if m != nil {
		return m.ParameterMatchSpecifier
	}
	return nil
}

func (m *MethodMatch_ParameterMatchSpecifier) GetExactMatch() string {
	if x, ok := m.GetParameterMatchSpecifier().(*MethodMatch_ParameterMatchSpecifier_ExactMatch); ok {
		return x.ExactMatch
	}
	return ""
}

func (m *MethodMatch_ParameterMatchSpecifier) GetRangeMatch() *_type.Int64Range {
	if x, ok := m.GetParameterMatchSpecifier().(*MethodMatch_ParameterMatchSpecifier_RangeMatch); ok {
		return x.RangeMatch
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*MethodMatch_ParameterMatchSpecifier) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*MethodMatch_ParameterMatchSpecifier_ExactMatch)(nil),
		(*MethodMatch_ParameterMatchSpecifier_RangeMatch)(nil),
	}
}

func init() {
	proto.RegisterType((*RouteConfiguration)(nil), "envoy.config.filter.network.dubbo_proxy.v2alpha1.RouteConfiguration")
	proto.RegisterType((*Route)(nil), "envoy.config.filter.network.dubbo_proxy.v2alpha1.Route")
	proto.RegisterType((*RouteMatch)(nil), "envoy.config.filter.network.dubbo_proxy.v2alpha1.RouteMatch")
	proto.RegisterType((*RouteAction)(nil), "envoy.config.filter.network.dubbo_proxy.v2alpha1.RouteAction")
	proto.RegisterType((*MethodMatch)(nil), "envoy.config.filter.network.dubbo_proxy.v2alpha1.MethodMatch")
	proto.RegisterMapType((map[uint32]*MethodMatch_ParameterMatchSpecifier)(nil), "envoy.config.filter.network.dubbo_proxy.v2alpha1.MethodMatch.ParamsMatchEntry")
	proto.RegisterType((*MethodMatch_ParameterMatchSpecifier)(nil), "envoy.config.filter.network.dubbo_proxy.v2alpha1.MethodMatch.ParameterMatchSpecifier")
}

func init() {
	proto.RegisterFile("envoy/config/filter/network/dubbo_proxy/v2alpha1/route.proto", fileDescriptor_74a572433a3292e0)
}

var fileDescriptor_74a572433a3292e0 = []byte{
	// 720 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0x4d, 0x6f, 0xd3, 0x4a,
	0x14, 0xed, 0x24, 0x75, 0xfa, 0x7a, 0xfd, 0x9e, 0x94, 0x8e, 0x9e, 0x5a, 0xbf, 0x3c, 0x3e, 0xda,
	0xb0, 0x29, 0x1b, 0x9b, 0xa6, 0x7c, 0x53, 0x90, 0x48, 0x85, 0x54, 0x16, 0x85, 0xca, 0x55, 0xc5,
	0x06, 0x14, 0x4d, 0x9d, 0x49, 0x62, 0x35, 0x99, 0x31, 0xe3, 0x71, 0xda, 0xec, 0x58, 0x23, 0x21,
	0x36, 0x20, 0x58, 0xf2, 0x3b, 0x58, 0x23, 0xc1, 0x96, 0x3f, 0xc1, 0x0f, 0x60, 0x85, 0x58, 0x20,
	0xe4, 0x3b, 0x63, 0x12, 0x68, 0x37, 0x2d, 0x6c, 0xac, 0xf1, 0x9d, 0x7b, 0xce, 0xb9, 0xf7, 0xe8,
	0xde, 0x81, 0x35, 0x2e, 0x86, 0x72, 0x14, 0x44, 0x52, 0x74, 0xe2, 0x6e, 0xd0, 0x89, 0xfb, 0x9a,
	0xab, 0x40, 0x70, 0xbd, 0x2f, 0xd5, 0x5e, 0xd0, 0xce, 0x76, 0x77, 0x65, 0x2b, 0x51, 0xf2, 0x60,
	0x14, 0x0c, 0x1b, 0xac, 0x9f, 0xf4, 0xd8, 0x4a, 0xa0, 0x64, 0xa6, 0xb9, 0x9f, 0x28, 0xa9, 0x25,
	0xbd, 0x80, 0x68, 0xdf, 0xa0, 0x7d, 0x83, 0xf6, 0x2d, 0xda, 0x9f, 0x40, 0xfb, 0x05, 0xba, 0x76,
	0xde, 0xe8, 0xb1, 0x24, 0x0e, 0x86, 0x0d, 0xc3, 0x65, 0xbe, 0xad, 0x48, 0x0e, 0x12, 0x29, 0xb8,
	0xd0, 0xa9, 0x21, 0xaf, 0x9d, 0x35, 0xa9, 0x7a, 0x94, 0xf0, 0x60, 0xc0, 0x74, 0xd4, 0xe3, 0x2a,
	0x48, 0xb5, 0x8a, 0x45, 0xd7, 0x26, 0xcc, 0x4f, 0x24, 0x28, 0x26, 0xba, 0xb6, 0xaa, 0xda, 0x99,
	0xac, 0x9d, 0xb0, 0x80, 0x09, 0x21, 0x35, 0xd3, 0xb1, 0x14, 0x69, 0x30, 0x88, 0xbb, 0x8a, 0x15,
	0x55, 0xd7, 0x4e, 0x1f, 0xba, 0x4f, 0x35, 0xd3, 0x59, 0xa1, 0xbb, 0x30, 0x64, 0xfd, 0xb8, 0xcd,
	0x34, 0x0f, 0x8a, 0x83, 0xb9, 0xa8, 0xbf, 0x27, 0x40, 0xc3, 0xbc, 0xd6, 0x75, 0xec, 0x37, 0x53,
	0x88, 0xa6, 0x14, 0xa6, 0x05, 0x1b, 0x70, 0x8f, 0x2c, 0x92, 0xe5, 0xd9, 0x10, 0xcf, 0xf4, 0x14,
	0xcc, 0xc6, 0x42, 0x73, 0xd5, 0x61, 0x11, 0xf7, 0x4a, 0x78, 0x31, 0x0e, 0xd0, 0x7f, 0xc1, 0xe9,
	0x2a, 0x99, 0x25, 0x5e, 0x19, 0x6f, 0xcc, 0x0f, 0xf5, 0x60, 0x66, 0xc8, 0x55, 0x1a, 0x4b, 0xe1,
	0x4d, 0x63, 0xbc, 0xf8, 0xa5, 0xf7, 0xa1, 0x82, 0x1e, 0xa5, 0x9e, 0xb3, 0x58, 0x5e, 0x76, 0x1b,
	0x57, 0xfc, 0xe3, 0xfa, 0xee, 0x63, 0xdd, 0xa1, 0xa5, 0xa9, 0xbf, 0x23, 0xe0, 0x60, 0x84, 0x3e,
	0x04, 0x07, 0xbd, 0xc5, 0xea, 0xdd, 0xc6, 0xda, 0x09, 0x99, 0x37, 0x73, 0x8e, 0xe6, 0x5f, 0x5f,
	0x9b, 0xce, 0x53, 0x52, 0xaa, 0x92, 0xd0, 0x90, 0xd2, 0x47, 0xe0, 0xa0, 0x22, 0x5a, 0xe0, 0x36,
	0x6e, 0x9e, 0x90, 0xfd, 0x76, 0x94, 0x1b, 0x3d, 0x49, 0x8f, 0xac, 0xf5, 0x37, 0x04, 0x60, 0x2c,
	0x4f, 0x77, 0xa0, 0x32, 0xe0, 0xba, 0x27, 0xdb, 0xb6, 0x99, 0x13, 0xc8, 0x6d, 0x22, 0x1e, 0xe9,
	0x42, 0x4b, 0x46, 0x6f, 0xc0, 0x4c, 0x8f, 0xb3, 0x36, 0x57, 0xa9, 0x57, 0x42, 0xfb, 0x97, 0x2c,
	0x2f, 0x4b, 0x62, 0x7f, 0xd8, 0xf0, 0xcd, 0x42, 0x6c, 0x60, 0xca, 0xa6, 0x99, 0xd3, 0xb0, 0x40,
	0xd4, 0x5f, 0x11, 0x70, 0x27, 0x7a, 0xa0, 0x35, 0x98, 0x89, 0xfa, 0x59, 0xaa, 0xb9, 0x32, 0xf3,
	0xb2, 0x31, 0x15, 0x16, 0x01, 0x1a, 0xc2, 0xdc, 0x3e, 0x8f, 0xbb, 0x3d, 0xcd, 0xdb, 0x2d, 0x1b,
	0x4b, 0xad, 0x73, 0xe7, 0x8e, 0x92, 0x7c, 0x60, 0x93, 0xd7, 0x4d, 0xee, 0xc6, 0x54, 0x58, 0xdd,
	0xff, 0x39, 0x94, 0x36, 0x3d, 0x98, 0xb3, 0x54, 0xad, 0x34, 0xe1, 0x51, 0xdc, 0x89, 0xb9, 0xa2,
	0xe5, 0x2f, 0x4d, 0x52, 0xff, 0x54, 0x06, 0x77, 0xa2, 0x5d, 0x7a, 0x69, 0x62, 0x8c, 0xc7, 0x3d,
	0xe6, 0xcb, 0xe5, 0xdb, 0xed, 0xf3, 0xb7, 0x71, 0xfb, 0x8a, 0x1e, 0xcd, 0xa4, 0x3f, 0x86, 0xbf,
	0x13, 0xa6, 0xd8, 0x20, 0x6d, 0x99, 0x39, 0x32, 0x16, 0xdd, 0xfb, 0x2d, 0xeb, 0xfd, 0x2d, 0x64,
	0xc4, 0xf3, 0x1d, 0xa1, 0xd5, 0x28, 0x74, 0x93, 0x71, 0xa4, 0xf6, 0x82, 0xc0, 0x02, 0x66, 0x70,
	0x6d, 0x1d, 0xdf, 0xfe, 0xd1, 0xda, 0x12, 0xb8, 0xfc, 0x80, 0x45, 0xda, 0x56, 0x53, 0xb6, 0x1e,
	0x03, 0x06, 0x4d, 0xa3, 0xd7, 0xc0, 0xc5, 0xd7, 0xc2, 0xa6, 0x4c, 0x63, 0xbf, 0xf3, 0x93, 0xfd,
	0xde, 0x15, 0xfa, 0xf2, 0xc5, 0x30, 0xcf, 0xc9, 0xa1, 0x98, 0x6c, 0x06, 0xfc, 0x7f, 0xf8, 0x2f,
	0x29, 0x84, 0x0d, 0x7c, 0xec, 0x6a, 0xed, 0x25, 0x81, 0xea, 0xaf, 0x85, 0xd3, 0x2a, 0x94, 0xf7,
	0xf8, 0x08, 0x4d, 0xfd, 0x27, 0xcc, 0x8f, 0x74, 0x0f, 0x9c, 0x21, 0xeb, 0x67, 0xc5, 0x4e, 0xec,
	0xfc, 0x01, 0xa7, 0x0e, 0xfb, 0x10, 0x1a, 0x8d, 0xeb, 0xa5, 0xab, 0xa4, 0xf9, 0x8c, 0x7c, 0x7e,
	0xfd, 0xed, 0xb9, 0xb3, 0x42, 0x03, 0xa3, 0xc4, 0x0f, 0x34, 0x17, 0xf9, 0xbb, 0x92, 0x5a, 0xb5,
	0xf4, 0x68, 0xb9, 0xd5, 0xb7, 0x4f, 0x3e, 0x7c, 0xac, 0x94, 0xaa, 0x04, 0x6e, 0xc5, 0xd2, 0x54,
	0x69, 0x6e, 0x8e, 0x5b, 0x70, 0xd3, 0x2c, 0xe9, 0x56, 0xfe, 0x88, 0x6e, 0x91, 0xdd, 0x0a, 0xbe,
	0xa6, 0xab, 0xdf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x9f, 0xdc, 0xf1, 0xc1, 0x7b, 0x06, 0x00, 0x00,
}
