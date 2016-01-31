// Code generated by protoc-gen-go.
// source: identity_assertion.proto
// DO NOT EDIT!

package credence

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// A statement that there is a public/private key pair that a person is using
// with Credence.
// - The public key in question is stated outright.
// - The person claiming to own the associated private key can be identified
//   by the identity_uri (and possibly the name).
// - The proof_uri might be a link to a post on twitter by the user saying "I am
//   a Credence User with fingerprint 00112233445566778899aabbccddeeff" allowing
//   other users to make a judgement call as to whether that is the person they
//   are expecting.
// - Any user can sign the assertion that this public key is associated to this identity.
//   - When self-signing, the fingerprint should be the fingerprint of the given
//     public key, and the associated private key used to make the signature.
//   - When counter-signing, the fingerprint should be of the counter-signing user
//     and the signature created using their private key.
type IdentityAssertion struct {
	// Their public key
	PublicKey []byte `protobuf:"bytes,1,opt,name=public_key,proto3" json:"public_key,omitempty"`
	// The name this user suggests they should go by within Credence.
	Name string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	// The main URI that can be used by others to distinguish this user from others with the same name. eg. a Twitter profile
	IdentityUri string `protobuf:"bytes,3,opt,name=identity_uri" json:"identity_uri,omitempty"`
	// A URI that attempts to prove that the paired private key is owned by the identity stated.
	ProofUri string `protobuf:"bytes,4,opt,name=proof_uri" json:"proof_uri,omitempty"`
	// The fingerprint of the user asserting this name and public key are related.
	Fingerprint []byte `protobuf:"bytes,5,opt,name=fingerprint,proto3" json:"fingerprint,omitempty"`
	// A signature of this message (with this fields blanked) by the private key referenced by the fingerprint
	Signature []byte `protobuf:"bytes,6,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (m *IdentityAssertion) Reset()                    { *m = IdentityAssertion{} }
func (m *IdentityAssertion) String() string            { return proto.CompactTextString(m) }
func (*IdentityAssertion) ProtoMessage()               {}
func (*IdentityAssertion) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func init() {
	proto.RegisterType((*IdentityAssertion)(nil), "credence.IdentityAssertion")
}

var fileDescriptor1 = []byte{
	// 162 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x3c, 0x8e, 0x31, 0x0e, 0x82, 0x50,
	0x0c, 0x40, 0x83, 0x22, 0x91, 0xca, 0xc2, 0xd7, 0xa1, 0xa3, 0x71, 0x72, 0x72, 0xf1, 0x04, 0x8e,
	0x9e, 0x82, 0x00, 0x16, 0xd2, 0xa8, 0xfd, 0xa4, 0x94, 0x81, 0x0b, 0x78, 0x6e, 0xbf, 0x3f, 0xea,
	0xd8, 0xd7, 0x97, 0xd7, 0x02, 0xf2, 0x8d, 0xc4, 0xd8, 0xe6, 0xaa, 0x1e, 0x47, 0x52, 0x63, 0x2f,
	0xa7, 0x41, 0xbd, 0x79, 0xb7, 0x6e, 0x95, 0xc2, 0xae, 0xa5, 0xc3, 0x2b, 0x81, 0xf2, 0xfa, 0xd5,
	0x2e, 0x3f, 0xcb, 0x39, 0x80, 0x61, 0x6a, 0x1e, 0xdc, 0x56, 0x77, 0x9a, 0x31, 0xd9, 0x27, 0xc7,
	0xc2, 0x15, 0x90, 0x4a, 0xfd, 0x24, 0x5c, 0x84, 0x29, 0x77, 0x3b, 0x28, 0xfe, 0xf5, 0x49, 0x19,
	0x97, 0x91, 0x96, 0x90, 0x87, 0x03, 0xbe, 0x8b, 0x28, 0x8d, 0x68, 0x0b, 0x9b, 0x8e, 0xa5, 0x27,
	0x1d, 0x94, 0xc5, 0x70, 0x15, 0x5b, 0xc1, 0x1b, 0xb9, 0x97, 0xda, 0x26, 0x25, 0xcc, 0x3e, 0xa8,
	0xc9, 0xe2, 0x67, 0xe7, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0xdd, 0x40, 0x07, 0x2b, 0xb5, 0x00,
	0x00, 0x00,
}