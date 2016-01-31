// Code generated by protoc-gen-go.
// source: search_request.proto
// DO NOT EDIT!

package credence

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// A search request defines a request of a cred store
// to look for any creds with the given key and publish them.
type SearchRequest struct {
	// A list of keys that the requesting system would like repeated
	Keys []string `protobuf:"bytes,1,rep,name=keys" json:"keys,omitempty"`
	// The time the request was issued (used for determining if a response is worthwhile)
	Timestamp int64 `protobuf:"varint,2,opt,name=timestamp" json:"timestamp,omitempty"`
	// The number of times this request has been passed on
	Proximity int32 `protobuf:"varint,3,opt,name=proximity" json:"proximity,omitempty"`
}

func (m *SearchRequest) Reset()                    { *m = SearchRequest{} }
func (m *SearchRequest) String() string            { return proto.CompactTextString(m) }
func (*SearchRequest) ProtoMessage()               {}
func (*SearchRequest) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

func init() {
	proto.RegisterType((*SearchRequest)(nil), "credence.SearchRequest")
}

var fileDescriptor3 = []byte{
	// 118 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x29, 0x4e, 0x4d, 0x2c,
	0x4a, 0xce, 0x88, 0x2f, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0x17, 0xe2, 0x48, 0x2e, 0x4a, 0x4d, 0x49, 0xcd, 0x4b, 0x4e, 0x55, 0x72, 0xe6, 0xe2, 0x0d, 0x06,
	0xab, 0x08, 0x82, 0x28, 0x10, 0xe2, 0xe1, 0x62, 0xc9, 0x4e, 0xad, 0x2c, 0x96, 0x60, 0x54, 0x60,
	0xd6, 0xe0, 0x14, 0x12, 0xe4, 0xe2, 0x2c, 0xc9, 0xcc, 0x05, 0x8a, 0x27, 0xe6, 0x16, 0x48, 0x30,
	0x29, 0x30, 0x6a, 0x30, 0x83, 0x84, 0x80, 0x86, 0x54, 0x64, 0xe6, 0x66, 0x96, 0x54, 0x4a, 0x30,
	0x03, 0x85, 0x58, 0x93, 0xd8, 0xc0, 0xa6, 0x1a, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x47, 0x80,
	0xd5, 0x2f, 0x6d, 0x00, 0x00, 0x00,
}
