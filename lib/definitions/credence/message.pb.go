// Code generated by protoc-gen-go.
// source: message.proto
// DO NOT EDIT!

package credence

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// A container for all the messages which can be sent
type Message struct {
	// Types that are valid to be assigned to Type:
	//	*Message_Cred
	//	*Message_SearchRequest
	Type isMessage_Type `protobuf_oneof:"type"`
}

func (m *Message) Reset()                    { *m = Message{} }
func (m *Message) String() string            { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()               {}
func (*Message) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

type isMessage_Type interface {
	isMessage_Type()
}

type Message_Cred struct {
	Cred *Cred `protobuf:"bytes,1,opt,name=cred,oneof"`
}
type Message_SearchRequest struct {
	SearchRequest *SearchRequest `protobuf:"bytes,2,opt,name=search_request,oneof"`
}

func (*Message_Cred) isMessage_Type()          {}
func (*Message_SearchRequest) isMessage_Type() {}

func (m *Message) GetType() isMessage_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *Message) GetCred() *Cred {
	if x, ok := m.GetType().(*Message_Cred); ok {
		return x.Cred
	}
	return nil
}

func (m *Message) GetSearchRequest() *SearchRequest {
	if x, ok := m.GetType().(*Message_SearchRequest); ok {
		return x.SearchRequest
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Message) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Message_OneofMarshaler, _Message_OneofUnmarshaler, _Message_OneofSizer, []interface{}{
		(*Message_Cred)(nil),
		(*Message_SearchRequest)(nil),
	}
}

func _Message_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Message)
	// type
	switch x := m.Type.(type) {
	case *Message_Cred:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Cred); err != nil {
			return err
		}
	case *Message_SearchRequest:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.SearchRequest); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Message.Type has unexpected type %T", x)
	}
	return nil
}

func _Message_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Message)
	switch tag {
	case 1: // type.cred
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Cred)
		err := b.DecodeMessage(msg)
		m.Type = &Message_Cred{msg}
		return true, err
	case 2: // type.search_request
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(SearchRequest)
		err := b.DecodeMessage(msg)
		m.Type = &Message_SearchRequest{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Message_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Message)
	// type
	switch x := m.Type.(type) {
	case *Message_Cred:
		s := proto.Size(x.Cred)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Message_SearchRequest:
		s := proto.Size(x.SearchRequest)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*Message)(nil), "credence.Message")
}

var fileDescriptor2 = []byte{
	// 137 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0xcd, 0x4d, 0x2d, 0x2e,
	0x4e, 0x4c, 0x4f, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x48, 0x2e, 0x4a, 0x4d, 0x49,
	0xcd, 0x4b, 0x4e, 0x95, 0xe2, 0x02, 0xb1, 0x20, 0xa2, 0x52, 0x22, 0xc5, 0xa9, 0x89, 0x45, 0xc9,
	0x19, 0xf1, 0x45, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x10, 0x51, 0xa5, 0x14, 0x2e, 0x76, 0x5f,
	0x88, 0x66, 0x21, 0x39, 0x2e, 0x16, 0x90, 0x72, 0x09, 0x46, 0x05, 0x46, 0x0d, 0x6e, 0x23, 0x3e,
	0x3d, 0x98, 0x29, 0x7a, 0xce, 0x40, 0x86, 0x07, 0x83, 0x90, 0x21, 0x17, 0x1f, 0xaa, 0x11, 0x12,
	0x4c, 0x60, 0x95, 0xe2, 0x08, 0x95, 0xc1, 0x60, 0xf9, 0x20, 0x88, 0xb4, 0x07, 0x83, 0x13, 0x1b,
	0x17, 0x4b, 0x49, 0x65, 0x41, 0x6a, 0x12, 0x1b, 0xd8, 0x32, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x37, 0x9f, 0x28, 0xa1, 0xa9, 0x00, 0x00, 0x00,
}
