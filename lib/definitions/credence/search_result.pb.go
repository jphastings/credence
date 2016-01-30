// Code generated by protoc-gen-go.
// source: search_result.proto
// DO NOT EDIT!

package credence

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// A List of the keys searched for and the breakdown of Credence'd
// belief of the various creds underneath them
type SearchResult struct {
	Results []*SearchResult_SourceBreakdown `protobuf:"bytes,1,rep,name=results" json:"results,omitempty"`
}

func (m *SearchResult) Reset()                    { *m = SearchResult{} }
func (m *SearchResult) String() string            { return proto.CompactTextString(m) }
func (*SearchResult) ProtoMessage()               {}
func (*SearchResult) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

func (m *SearchResult) GetResults() []*SearchResult_SourceBreakdown {
	if m != nil {
		return m.Results
	}
	return nil
}

type SearchResult_AssertionBreakdown struct {
	StatementHash string `protobuf:"bytes,1,opt,name=statement_hash" json:"statement_hash,omitempty"`
	NoComment     int64  `protobuf:"varint,2,opt,name=no_comment" json:"no_comment,omitempty"`
	IsTrue        int64  `protobuf:"varint,3,opt,name=is_true" json:"is_true,omitempty"`
	IsFalse       int64  `protobuf:"varint,4,opt,name=is_false" json:"is_false,omitempty"`
	IsAmbiguous   int64  `protobuf:"varint,5,opt,name=is_ambiguous" json:"is_ambiguous,omitempty"`
	Recognised    int64  `protobuf:"varint,6,opt,name=recognised" json:"recognised,omitempty"`
	// TODO: include this directly from cred?
	//
	// Types that are valid to be assigned to Statement:
	//	*SearchResult_AssertionBreakdown_HumanReadable
	//	*SearchResult_AssertionBreakdown_CredSeen
	//	*SearchResult_AssertionBreakdown_ApplicationSpecific
	//	*SearchResult_AssertionBreakdown_IdentityDeclaration
	Statement isSearchResult_AssertionBreakdown_Statement `protobuf_oneof:"statement"`
}

func (m *SearchResult_AssertionBreakdown) Reset()         { *m = SearchResult_AssertionBreakdown{} }
func (m *SearchResult_AssertionBreakdown) String() string { return proto.CompactTextString(m) }
func (*SearchResult_AssertionBreakdown) ProtoMessage()    {}
func (*SearchResult_AssertionBreakdown) Descriptor() ([]byte, []int) {
	return fileDescriptor3, []int{0, 0}
}

type isSearchResult_AssertionBreakdown_Statement interface {
	isSearchResult_AssertionBreakdown_Statement()
}

type SearchResult_AssertionBreakdown_HumanReadable struct {
	HumanReadable *Cred_HumanReadableStatement `protobuf:"bytes,7,opt,name=human_readable,oneof"`
}
type SearchResult_AssertionBreakdown_CredSeen struct {
	CredSeen *Cred_CredSeenStatement `protobuf:"bytes,8,opt,name=cred_seen,oneof"`
}
type SearchResult_AssertionBreakdown_ApplicationSpecific struct {
	ApplicationSpecific *Cred_ApplicationSpecificStatement `protobuf:"bytes,9,opt,name=application_specific,oneof"`
}
type SearchResult_AssertionBreakdown_IdentityDeclaration struct {
	IdentityDeclaration *Cred_IdentityDeclarationStatement `protobuf:"bytes,10,opt,name=identity_declaration,oneof"`
}

func (*SearchResult_AssertionBreakdown_HumanReadable) isSearchResult_AssertionBreakdown_Statement() {}
func (*SearchResult_AssertionBreakdown_CredSeen) isSearchResult_AssertionBreakdown_Statement()      {}
func (*SearchResult_AssertionBreakdown_ApplicationSpecific) isSearchResult_AssertionBreakdown_Statement() {
}
func (*SearchResult_AssertionBreakdown_IdentityDeclaration) isSearchResult_AssertionBreakdown_Statement() {
}

func (m *SearchResult_AssertionBreakdown) GetStatement() isSearchResult_AssertionBreakdown_Statement {
	if m != nil {
		return m.Statement
	}
	return nil
}

func (m *SearchResult_AssertionBreakdown) GetHumanReadable() *Cred_HumanReadableStatement {
	if x, ok := m.GetStatement().(*SearchResult_AssertionBreakdown_HumanReadable); ok {
		return x.HumanReadable
	}
	return nil
}

func (m *SearchResult_AssertionBreakdown) GetCredSeen() *Cred_CredSeenStatement {
	if x, ok := m.GetStatement().(*SearchResult_AssertionBreakdown_CredSeen); ok {
		return x.CredSeen
	}
	return nil
}

func (m *SearchResult_AssertionBreakdown) GetApplicationSpecific() *Cred_ApplicationSpecificStatement {
	if x, ok := m.GetStatement().(*SearchResult_AssertionBreakdown_ApplicationSpecific); ok {
		return x.ApplicationSpecific
	}
	return nil
}

func (m *SearchResult_AssertionBreakdown) GetIdentityDeclaration() *Cred_IdentityDeclarationStatement {
	if x, ok := m.GetStatement().(*SearchResult_AssertionBreakdown_IdentityDeclaration); ok {
		return x.IdentityDeclaration
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*SearchResult_AssertionBreakdown) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _SearchResult_AssertionBreakdown_OneofMarshaler, _SearchResult_AssertionBreakdown_OneofUnmarshaler, _SearchResult_AssertionBreakdown_OneofSizer, []interface{}{
		(*SearchResult_AssertionBreakdown_HumanReadable)(nil),
		(*SearchResult_AssertionBreakdown_CredSeen)(nil),
		(*SearchResult_AssertionBreakdown_ApplicationSpecific)(nil),
		(*SearchResult_AssertionBreakdown_IdentityDeclaration)(nil),
	}
}

func _SearchResult_AssertionBreakdown_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*SearchResult_AssertionBreakdown)
	// statement
	switch x := m.Statement.(type) {
	case *SearchResult_AssertionBreakdown_HumanReadable:
		b.EncodeVarint(7<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.HumanReadable); err != nil {
			return err
		}
	case *SearchResult_AssertionBreakdown_CredSeen:
		b.EncodeVarint(8<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.CredSeen); err != nil {
			return err
		}
	case *SearchResult_AssertionBreakdown_ApplicationSpecific:
		b.EncodeVarint(9<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ApplicationSpecific); err != nil {
			return err
		}
	case *SearchResult_AssertionBreakdown_IdentityDeclaration:
		b.EncodeVarint(10<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.IdentityDeclaration); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("SearchResult_AssertionBreakdown.Statement has unexpected type %T", x)
	}
	return nil
}

func _SearchResult_AssertionBreakdown_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*SearchResult_AssertionBreakdown)
	switch tag {
	case 7: // statement.human_readable
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Cred_HumanReadableStatement)
		err := b.DecodeMessage(msg)
		m.Statement = &SearchResult_AssertionBreakdown_HumanReadable{msg}
		return true, err
	case 8: // statement.cred_seen
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Cred_CredSeenStatement)
		err := b.DecodeMessage(msg)
		m.Statement = &SearchResult_AssertionBreakdown_CredSeen{msg}
		return true, err
	case 9: // statement.application_specific
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Cred_ApplicationSpecificStatement)
		err := b.DecodeMessage(msg)
		m.Statement = &SearchResult_AssertionBreakdown_ApplicationSpecific{msg}
		return true, err
	case 10: // statement.identity_declaration
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Cred_IdentityDeclarationStatement)
		err := b.DecodeMessage(msg)
		m.Statement = &SearchResult_AssertionBreakdown_IdentityDeclaration{msg}
		return true, err
	default:
		return false, nil
	}
}

func _SearchResult_AssertionBreakdown_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*SearchResult_AssertionBreakdown)
	// statement
	switch x := m.Statement.(type) {
	case *SearchResult_AssertionBreakdown_HumanReadable:
		s := proto.Size(x.HumanReadable)
		n += proto.SizeVarint(7<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *SearchResult_AssertionBreakdown_CredSeen:
		s := proto.Size(x.CredSeen)
		n += proto.SizeVarint(8<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *SearchResult_AssertionBreakdown_ApplicationSpecific:
		s := proto.Size(x.ApplicationSpecific)
		n += proto.SizeVarint(9<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *SearchResult_AssertionBreakdown_IdentityDeclaration:
		s := proto.Size(x.IdentityDeclaration)
		n += proto.SizeVarint(10<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type SearchResult_SourceBreakdown struct {
	SourceUri  string                             `protobuf:"bytes,1,opt,name=source_uri" json:"source_uri,omitempty"`
	Assertions []*SearchResult_AssertionBreakdown `protobuf:"bytes,2,rep,name=assertions" json:"assertions,omitempty"`
}

func (m *SearchResult_SourceBreakdown) Reset()                    { *m = SearchResult_SourceBreakdown{} }
func (m *SearchResult_SourceBreakdown) String() string            { return proto.CompactTextString(m) }
func (*SearchResult_SourceBreakdown) ProtoMessage()               {}
func (*SearchResult_SourceBreakdown) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0, 1} }

func (m *SearchResult_SourceBreakdown) GetAssertions() []*SearchResult_AssertionBreakdown {
	if m != nil {
		return m.Assertions
	}
	return nil
}

func init() {
	proto.RegisterType((*SearchResult)(nil), "credence.SearchResult")
	proto.RegisterType((*SearchResult_AssertionBreakdown)(nil), "credence.SearchResult.AssertionBreakdown")
	proto.RegisterType((*SearchResult_SourceBreakdown)(nil), "credence.SearchResult.SourceBreakdown")
}

var fileDescriptor3 = []byte{
	// 375 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x74, 0x92, 0xdf, 0xca, 0xd3, 0x30,
	0x18, 0xc6, 0xfd, 0xec, 0xe7, 0xb7, 0xf5, 0xdd, 0xd8, 0x24, 0x0e, 0x09, 0x3b, 0x1a, 0x82, 0x32,
	0x11, 0x7a, 0x30, 0x11, 0x8f, 0x44, 0x36, 0x3d, 0x98, 0x78, 0xb6, 0x5e, 0x40, 0xc8, 0xd2, 0x77,
	0x6b, 0xb0, 0x4d, 0x4a, 0x92, 0x22, 0xde, 0x93, 0x97, 0xe4, 0xc5, 0x98, 0x64, 0xdd, 0x1f, 0xaa,
	0xdf, 0x49, 0x29, 0xbf, 0xfc, 0xde, 0x87, 0xf6, 0x7d, 0x02, 0x2f, 0x2c, 0x72, 0x23, 0x4a, 0x66,
	0xd0, 0xb6, 0x95, 0xcb, 0x1a, 0xa3, 0x9d, 0x26, 0x43, 0x61, 0xb0, 0x40, 0x25, 0x70, 0x0e, 0xe1,
	0xed, 0x44, 0x5f, 0xfd, 0xb9, 0x87, 0x71, 0x1e, 0xed, 0x5d, 0x94, 0xc9, 0x47, 0x18, 0x9c, 0xc6,
	0x2c, 0xbd, 0x5b, 0x24, 0xcb, 0xd1, 0xea, 0x4d, 0x76, 0x1e, 0xcc, 0x6e, 0xc5, 0x2c, 0xd7, 0xad,
	0x11, 0xb8, 0x31, 0xc8, 0x7f, 0x14, 0xfa, 0xa7, 0x9a, 0xff, 0x4e, 0x80, 0xac, 0xad, 0x45, 0xe3,
	0xa4, 0x56, 0x17, 0x4c, 0x5e, 0xc2, 0xc4, 0x3a, 0xee, 0xb0, 0x46, 0xe5, 0x58, 0xc9, 0x6d, 0xe9,
	0x63, 0xef, 0x96, 0x29, 0x21, 0x00, 0x4a, 0x33, 0xa1, 0xeb, 0x70, 0x40, 0x9f, 0x7a, 0x96, 0x90,
	0x29, 0x0c, 0xa4, 0x65, 0xce, 0xb4, 0x48, 0x93, 0x08, 0x9e, 0xc3, 0xd0, 0x83, 0x03, 0xaf, 0x2c,
	0xd2, 0xfb, 0x48, 0x66, 0x30, 0xf6, 0x84, 0xd7, 0x7b, 0x79, 0x6c, 0x75, 0x6b, 0xe9, 0xb3, 0x48,
	0x7d, 0x98, 0x41, 0xa1, 0x8f, 0x4a, 0x5a, 0x2c, 0xe8, 0x43, 0x64, 0x9f, 0x61, 0x52, 0xb6, 0x35,
	0x57, 0x7e, 0x0b, 0xbc, 0xe0, 0xfb, 0x0a, 0xe9, 0xc0, 0xf3, 0xd1, 0xea, 0xf5, 0xf5, 0x7f, 0xbe,
	0x84, 0x3d, 0x6c, 0x83, 0xb4, 0xeb, 0x9c, 0xfc, 0xfc, 0xa9, 0xdb, 0x27, 0xe4, 0x03, 0xa4, 0xc1,
	0x64, 0x16, 0x51, 0xd1, 0x61, 0x9c, 0x5d, 0xf4, 0x66, 0xc3, 0x23, 0xf7, 0xc7, 0xb7, 0x63, 0xdf,
	0x61, 0xc6, 0x9b, 0xa6, 0x92, 0x82, 0x87, 0x45, 0x30, 0xdb, 0xa0, 0x90, 0x07, 0x29, 0x68, 0x1a,
	0x13, 0xde, 0xf5, 0x12, 0xd6, 0x57, 0x35, 0xef, 0xcc, 0x5e, 0x98, 0xf4, 0xb6, 0x93, 0xee, 0x17,
	0x2b, 0x50, 0x54, 0xdc, 0x44, 0x95, 0xc2, 0x7f, 0xc3, 0xbe, 0x75, 0xea, 0xd7, 0xab, 0x79, 0x13,
	0xb6, 0x19, 0x41, 0x7a, 0xa9, 0x62, 0x5e, 0xc0, 0xb4, 0xd7, 0x60, 0xd8, 0xa2, 0x8d, 0x88, 0xb5,
	0x46, 0x76, 0x35, 0x7d, 0x02, 0xe0, 0xe7, 0x52, 0xad, 0xaf, 0x29, 0xdc, 0x88, 0xb7, 0x8f, 0xdc,
	0x88, 0x7f, 0xdb, 0xdf, 0x3f, 0xc4, 0x5b, 0xf6, 0xfe, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd9,
	0x11, 0x1a, 0xe2, 0x92, 0x02, 0x00, 0x00,
}
