// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/cluster/v4alpha/circuit_breaker.proto

package envoy_config_cluster_v4alpha

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	v4alpha "github.com/datawire/ambassador/pkg/api/envoy/config/core/v4alpha"
	v3 "github.com/datawire/ambassador/pkg/api/envoy/type/v3"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

type CircuitBreakers struct {
	Thresholds           []*CircuitBreakers_Thresholds `protobuf:"bytes,1,rep,name=thresholds,proto3" json:"thresholds,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *CircuitBreakers) Reset()         { *m = CircuitBreakers{} }
func (m *CircuitBreakers) String() string { return proto.CompactTextString(m) }
func (*CircuitBreakers) ProtoMessage()    {}
func (*CircuitBreakers) Descriptor() ([]byte, []int) {
	return fileDescriptor_7e788e4bcf428d89, []int{0}
}

func (m *CircuitBreakers) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CircuitBreakers.Unmarshal(m, b)
}
func (m *CircuitBreakers) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CircuitBreakers.Marshal(b, m, deterministic)
}
func (m *CircuitBreakers) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CircuitBreakers.Merge(m, src)
}
func (m *CircuitBreakers) XXX_Size() int {
	return xxx_messageInfo_CircuitBreakers.Size(m)
}
func (m *CircuitBreakers) XXX_DiscardUnknown() {
	xxx_messageInfo_CircuitBreakers.DiscardUnknown(m)
}

var xxx_messageInfo_CircuitBreakers proto.InternalMessageInfo

func (m *CircuitBreakers) GetThresholds() []*CircuitBreakers_Thresholds {
	if m != nil {
		return m.Thresholds
	}
	return nil
}

type CircuitBreakers_Thresholds struct {
	Priority             v4alpha.RoutingPriority                 `protobuf:"varint,1,opt,name=priority,proto3,enum=envoy.config.core.v4alpha.RoutingPriority" json:"priority,omitempty"`
	MaxConnections       *wrappers.UInt32Value                   `protobuf:"bytes,2,opt,name=max_connections,json=maxConnections,proto3" json:"max_connections,omitempty"`
	MaxPendingRequests   *wrappers.UInt32Value                   `protobuf:"bytes,3,opt,name=max_pending_requests,json=maxPendingRequests,proto3" json:"max_pending_requests,omitempty"`
	MaxRequests          *wrappers.UInt32Value                   `protobuf:"bytes,4,opt,name=max_requests,json=maxRequests,proto3" json:"max_requests,omitempty"`
	MaxRetries           *wrappers.UInt32Value                   `protobuf:"bytes,5,opt,name=max_retries,json=maxRetries,proto3" json:"max_retries,omitempty"`
	RetryBudget          *CircuitBreakers_Thresholds_RetryBudget `protobuf:"bytes,8,opt,name=retry_budget,json=retryBudget,proto3" json:"retry_budget,omitempty"`
	TrackRemaining       bool                                    `protobuf:"varint,6,opt,name=track_remaining,json=trackRemaining,proto3" json:"track_remaining,omitempty"`
	MaxConnectionPools   *wrappers.UInt32Value                   `protobuf:"bytes,7,opt,name=max_connection_pools,json=maxConnectionPools,proto3" json:"max_connection_pools,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                `json:"-"`
	XXX_unrecognized     []byte                                  `json:"-"`
	XXX_sizecache        int32                                   `json:"-"`
}

func (m *CircuitBreakers_Thresholds) Reset()         { *m = CircuitBreakers_Thresholds{} }
func (m *CircuitBreakers_Thresholds) String() string { return proto.CompactTextString(m) }
func (*CircuitBreakers_Thresholds) ProtoMessage()    {}
func (*CircuitBreakers_Thresholds) Descriptor() ([]byte, []int) {
	return fileDescriptor_7e788e4bcf428d89, []int{0, 0}
}

func (m *CircuitBreakers_Thresholds) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CircuitBreakers_Thresholds.Unmarshal(m, b)
}
func (m *CircuitBreakers_Thresholds) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CircuitBreakers_Thresholds.Marshal(b, m, deterministic)
}
func (m *CircuitBreakers_Thresholds) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CircuitBreakers_Thresholds.Merge(m, src)
}
func (m *CircuitBreakers_Thresholds) XXX_Size() int {
	return xxx_messageInfo_CircuitBreakers_Thresholds.Size(m)
}
func (m *CircuitBreakers_Thresholds) XXX_DiscardUnknown() {
	xxx_messageInfo_CircuitBreakers_Thresholds.DiscardUnknown(m)
}

var xxx_messageInfo_CircuitBreakers_Thresholds proto.InternalMessageInfo

func (m *CircuitBreakers_Thresholds) GetPriority() v4alpha.RoutingPriority {
	if m != nil {
		return m.Priority
	}
	return v4alpha.RoutingPriority_DEFAULT
}

func (m *CircuitBreakers_Thresholds) GetMaxConnections() *wrappers.UInt32Value {
	if m != nil {
		return m.MaxConnections
	}
	return nil
}

func (m *CircuitBreakers_Thresholds) GetMaxPendingRequests() *wrappers.UInt32Value {
	if m != nil {
		return m.MaxPendingRequests
	}
	return nil
}

func (m *CircuitBreakers_Thresholds) GetMaxRequests() *wrappers.UInt32Value {
	if m != nil {
		return m.MaxRequests
	}
	return nil
}

func (m *CircuitBreakers_Thresholds) GetMaxRetries() *wrappers.UInt32Value {
	if m != nil {
		return m.MaxRetries
	}
	return nil
}

func (m *CircuitBreakers_Thresholds) GetRetryBudget() *CircuitBreakers_Thresholds_RetryBudget {
	if m != nil {
		return m.RetryBudget
	}
	return nil
}

func (m *CircuitBreakers_Thresholds) GetTrackRemaining() bool {
	if m != nil {
		return m.TrackRemaining
	}
	return false
}

func (m *CircuitBreakers_Thresholds) GetMaxConnectionPools() *wrappers.UInt32Value {
	if m != nil {
		return m.MaxConnectionPools
	}
	return nil
}

type CircuitBreakers_Thresholds_RetryBudget struct {
	BudgetPercent        *v3.Percent           `protobuf:"bytes,1,opt,name=budget_percent,json=budgetPercent,proto3" json:"budget_percent,omitempty"`
	MinRetryConcurrency  *wrappers.UInt32Value `protobuf:"bytes,2,opt,name=min_retry_concurrency,json=minRetryConcurrency,proto3" json:"min_retry_concurrency,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *CircuitBreakers_Thresholds_RetryBudget) Reset() {
	*m = CircuitBreakers_Thresholds_RetryBudget{}
}
func (m *CircuitBreakers_Thresholds_RetryBudget) String() string { return proto.CompactTextString(m) }
func (*CircuitBreakers_Thresholds_RetryBudget) ProtoMessage()    {}
func (*CircuitBreakers_Thresholds_RetryBudget) Descriptor() ([]byte, []int) {
	return fileDescriptor_7e788e4bcf428d89, []int{0, 0, 0}
}

func (m *CircuitBreakers_Thresholds_RetryBudget) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CircuitBreakers_Thresholds_RetryBudget.Unmarshal(m, b)
}
func (m *CircuitBreakers_Thresholds_RetryBudget) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CircuitBreakers_Thresholds_RetryBudget.Marshal(b, m, deterministic)
}
func (m *CircuitBreakers_Thresholds_RetryBudget) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CircuitBreakers_Thresholds_RetryBudget.Merge(m, src)
}
func (m *CircuitBreakers_Thresholds_RetryBudget) XXX_Size() int {
	return xxx_messageInfo_CircuitBreakers_Thresholds_RetryBudget.Size(m)
}
func (m *CircuitBreakers_Thresholds_RetryBudget) XXX_DiscardUnknown() {
	xxx_messageInfo_CircuitBreakers_Thresholds_RetryBudget.DiscardUnknown(m)
}

var xxx_messageInfo_CircuitBreakers_Thresholds_RetryBudget proto.InternalMessageInfo

func (m *CircuitBreakers_Thresholds_RetryBudget) GetBudgetPercent() *v3.Percent {
	if m != nil {
		return m.BudgetPercent
	}
	return nil
}

func (m *CircuitBreakers_Thresholds_RetryBudget) GetMinRetryConcurrency() *wrappers.UInt32Value {
	if m != nil {
		return m.MinRetryConcurrency
	}
	return nil
}

func init() {
	proto.RegisterType((*CircuitBreakers)(nil), "envoy.config.cluster.v4alpha.CircuitBreakers")
	proto.RegisterType((*CircuitBreakers_Thresholds)(nil), "envoy.config.cluster.v4alpha.CircuitBreakers.Thresholds")
	proto.RegisterType((*CircuitBreakers_Thresholds_RetryBudget)(nil), "envoy.config.cluster.v4alpha.CircuitBreakers.Thresholds.RetryBudget")
}

func init() {
	proto.RegisterFile("envoy/config/cluster/v4alpha/circuit_breaker.proto", fileDescriptor_7e788e4bcf428d89)
}

var fileDescriptor_7e788e4bcf428d89 = []byte{
	// 614 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x94, 0x4f, 0x6f, 0xd3, 0x3e,
	0x18, 0xc7, 0x95, 0xed, 0xb7, 0xfe, 0x2a, 0x67, 0x74, 0x53, 0xc6, 0x9f, 0xa8, 0x8c, 0xa9, 0x20,
	0xa4, 0x95, 0x1d, 0x1c, 0xa9, 0xdd, 0x01, 0x26, 0x0d, 0x50, 0xc7, 0x0e, 0x08, 0x09, 0x45, 0x11,
	0x20, 0x6e, 0x91, 0x9b, 0x3e, 0x4b, 0xad, 0x25, 0x76, 0xb0, 0x9d, 0xd0, 0xdc, 0x10, 0x27, 0x5e,
	0x02, 0xe2, 0xa5, 0x70, 0x47, 0xe2, 0xca, 0x8b, 0xe0, 0x4d, 0xec, 0x84, 0x12, 0xa7, 0xc9, 0x36,
	0xd0, 0x14, 0xb8, 0xd5, 0x79, 0xfc, 0xf9, 0x24, 0xcf, 0xd7, 0x8f, 0x8b, 0x46, 0xc0, 0x32, 0x9e,
	0x3b, 0x01, 0x67, 0x27, 0x34, 0x74, 0x82, 0x28, 0x95, 0x0a, 0x84, 0x93, 0xed, 0x93, 0x28, 0x99,
	0x13, 0x27, 0xa0, 0x22, 0x48, 0xa9, 0xf2, 0xa7, 0x02, 0xc8, 0x29, 0x08, 0x9c, 0x08, 0xae, 0xb8,
	0xb5, 0x5d, 0x32, 0x58, 0x33, 0xb8, 0x62, 0x70, 0xc5, 0xf4, 0xef, 0x5f, 0x34, 0x72, 0x01, 0xb5,
	0x6e, 0x4a, 0x24, 0x68, 0x47, 0xff, 0xb6, 0xde, 0xa5, 0xf2, 0x04, 0x9c, 0x6c, 0xec, 0x24, 0x20,
	0x02, 0x60, 0xaa, 0x2a, 0xee, 0x84, 0x9c, 0x87, 0x11, 0x38, 0xe5, 0x6a, 0x9a, 0x9e, 0x38, 0xef,
	0x05, 0x49, 0x12, 0x10, 0xb2, 0xaa, 0xdf, 0x49, 0x67, 0x09, 0x71, 0x08, 0x63, 0x5c, 0x11, 0x45,
	0x39, 0x93, 0x8e, 0x54, 0x44, 0xa5, 0xcb, 0xf2, 0xdd, 0xdf, 0xca, 0x19, 0x08, 0x49, 0x39, 0xa3,
	0x2c, 0xac, 0xb6, 0xdc, 0xca, 0x48, 0x44, 0x67, 0x44, 0x81, 0xb3, 0xfc, 0xa1, 0x0b, 0xf7, 0x3e,
	0x77, 0xd1, 0xc6, 0x91, 0xee, 0x7a, 0xa2, 0x9b, 0x96, 0xd6, 0x5b, 0x84, 0xd4, 0x5c, 0x80, 0x9c,
	0xf3, 0x68, 0x26, 0x6d, 0x63, 0xb0, 0x3a, 0x34, 0x47, 0x0f, 0xf1, 0x55, 0x21, 0xe0, 0x4b, 0x0a,
	0xfc, 0xaa, 0xe6, 0xbd, 0x73, 0xae, 0xfe, 0x59, 0x07, 0xa1, 0xa6, 0x64, 0xb9, 0xa8, 0x9b, 0x08,
	0xca, 0x05, 0x55, 0xb9, 0x6d, 0x0c, 0x8c, 0x61, 0x6f, 0xb4, 0x77, 0xe9, 0x35, 0x5c, 0x40, 0xfd,
	0x0e, 0x8f, 0xa7, 0x8a, 0xb2, 0xd0, 0xad, 0x88, 0x49, 0xf7, 0x6c, 0xb2, 0xf6, 0xd1, 0x58, 0xd9,
	0x34, 0xbc, 0xda, 0x62, 0x1d, 0xa3, 0x8d, 0x98, 0x2c, 0xfc, 0x80, 0x33, 0x06, 0x41, 0x99, 0x85,
	0xbd, 0x32, 0x30, 0x86, 0xe6, 0x68, 0x1b, 0xeb, 0x8c, 0xf1, 0x32, 0x63, 0xfc, 0xfa, 0x39, 0x53,
	0xe3, 0xd1, 0x1b, 0x12, 0xa5, 0xe0, 0xf5, 0x62, 0xb2, 0x38, 0x6a, 0x18, 0xeb, 0x25, 0xba, 0x5e,
	0x68, 0x12, 0x60, 0x33, 0xca, 0x42, 0x5f, 0xc0, 0xbb, 0x14, 0xa4, 0x92, 0xf6, 0x6a, 0x0b, 0x97,
	0x15, 0x93, 0x85, 0xab, 0x41, 0xaf, 0xe2, 0xac, 0x27, 0x68, 0xbd, 0xf0, 0xd5, 0x9e, 0xff, 0x5a,
	0x78, 0xcc, 0x98, 0x2c, 0x6a, 0xc1, 0x21, 0x32, 0xb5, 0x40, 0x09, 0x0a, 0xd2, 0x5e, 0x6b, 0xc1,
	0xa3, 0x92, 0x2f, 0xf7, 0x5b, 0x21, 0x5a, 0x2f, 0xd0, 0xdc, 0x9f, 0xa6, 0xb3, 0x10, 0x94, 0xdd,
	0x2d, 0xf9, 0x67, 0xff, 0x7a, 0xa6, 0xb8, 0xf0, 0xe6, 0x93, 0xd2, 0xe5, 0x99, 0xa2, 0x59, 0x58,
	0xbb, 0x68, 0x43, 0x09, 0x12, 0x9c, 0xfa, 0x02, 0x62, 0x42, 0x8b, 0x01, 0xb4, 0x3b, 0x03, 0x63,
	0xd8, 0xf5, 0x7a, 0xe5, 0x63, 0x6f, 0xf9, 0x74, 0x99, 0x70, 0x73, 0x50, 0x7e, 0xc2, 0x79, 0x24,
	0xed, 0xff, 0x5b, 0x26, 0xdc, 0x9c, 0x96, 0x5b, 0x70, 0xfd, 0x9f, 0x06, 0x32, 0xcf, 0x7d, 0x95,
	0x75, 0x88, 0x7a, 0xba, 0x57, 0xbf, 0xba, 0x6a, 0xe5, 0x80, 0x99, 0xa3, 0x9b, 0x55, 0xcf, 0xc5,
	0x45, 0xc4, 0xd9, 0x18, 0xbb, 0xba, 0xea, 0x5d, 0xd3, 0xbb, 0xab, 0xa5, 0xe5, 0xa2, 0x1b, 0x31,
	0x65, 0xbe, 0x0e, 0x2d, 0xe0, 0x2c, 0x48, 0x85, 0x00, 0x16, 0xe4, 0xad, 0xa6, 0x69, 0x2b, 0xa6,
	0xac, 0xfc, 0x96, 0xa3, 0x06, 0x3c, 0x38, 0xfe, 0xf2, 0xed, 0xd3, 0xce, 0x53, 0xf4, 0xf8, 0xcf,
	0x91, 0x8f, 0x5b, 0xa6, 0x7d, 0xf0, 0xa8, 0xd0, 0xec, 0x57, 0x7f, 0x63, 0x7f, 0xa5, 0x39, 0xc0,
	0x05, 0xfa, 0x00, 0xed, 0xb6, 0x44, 0x27, 0x2f, 0xbe, 0x7e, 0xf8, 0xfe, 0xa3, 0xb3, 0xb2, 0xb9,
	0x8a, 0xf6, 0x28, 0xd7, 0xb1, 0x25, 0x82, 0x2f, 0xf2, 0x2b, 0xa7, 0x66, 0xb2, 0x75, 0x51, 0xe3,
	0x16, 0x01, 0xb9, 0xc6, 0xb4, 0x53, 0x26, 0x35, 0xfe, 0x15, 0x00, 0x00, 0xff, 0xff, 0x53, 0xee,
	0xd2, 0xb2, 0x80, 0x05, 0x00, 0x00,
}
