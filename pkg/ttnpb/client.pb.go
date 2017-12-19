// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/TheThingsNetwork/ttn/api/client.proto

package ttnpb

import proto "github.com/gogo/protobuf/proto"
import golang_proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import _ "github.com/gogo/protobuf/types"

import time "time"

import strconv "strconv"

import github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// ReviewingState enum defines all the possible admin reviewing states that a
// request can be at. For example a third-party client creation request.
type ReviewingState int32

const (
	// Denotes that the request is pending to review.
	STATE_PENDING ReviewingState = 0
	// Denotes that the request has been reviewed and approved by an admin.
	STATE_APPROVED ReviewingState = 1
	// Denotes that the request has been reviewed and rejected by an admin.
	STATE_REJECTED ReviewingState = 2
)

var ReviewingState_name = map[int32]string{
	0: "STATE_PENDING",
	1: "STATE_APPROVED",
	2: "STATE_REJECTED",
}
var ReviewingState_value = map[string]int32{
	"STATE_PENDING":  0,
	"STATE_APPROVED": 1,
	"STATE_REJECTED": 2,
}

func (ReviewingState) EnumDescriptor() ([]byte, []int) { return fileDescriptorClient, []int{0} }

// Client is the message that defines a third-party client on the network.
type Client struct {
	// id is the client's unique and immutable ID.
	ClientIdentifier `protobuf:"bytes,1,opt,name=id,embedded=id" json:"id"`
	// description is the description of the client.
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// secret is the secret used to prove the client identity.
	// This is a read-only field.
	Secret string `protobuf:"bytes,3,opt,name=secret,proto3" json:"secret,omitempty"`
	// redirect_uri is the OAuth redirect URI of the client.
	RedirectURI string `protobuf:"bytes,4,opt,name=redirect_uri,json=redirectUri,proto3" json:"redirect_uri,omitempty"`
	// state denotes the reviewing state of the client by admin.
	// This field can only be modified by admins.
	State ReviewingState `protobuf:"varint,5,opt,name=state,proto3,enum=ttn.v3.ReviewingState" json:"state,omitempty"`
	// official_labeled denotes if a client has been labeled as an official
	// third-party client by the tenant admin.
	// This field can only be modified by admins.
	OfficialLabeled bool `protobuf:"varint,6,opt,name=official_labeled,json=officialLabeled,proto3" json:"official_labeled,omitempty"`
	// grants denotes which OAuth2 flows can the client use to get a token.
	// This field can only be modified by admins.
	Grants []GrantType `protobuf:"varint,7,rep,packed,name=grants,enum=ttn.v3.GrantType" json:"grants,omitempty"`
	// Rights denotes what rights the client will have access to.
	Rights []Right `protobuf:"varint,8,rep,packed,name=rights,enum=ttn.v3.Right" json:"rights,omitempty"`
	// creator is the user ID of the user who created the client.
	// This is an immutable and read-only field.
	Creator UserIdentifier `protobuf:"bytes,9,opt,name=creator" json:"creator"`
	// created_at denotes when the client was created.
	// This a read-only field.
	CreatedAt time.Time `protobuf:"bytes,10,opt,name=created_at,json=createdAt,stdtime" json:"created_at"`
	// updated_at is the last time the client was updated.
	// This is a read-only field.
	UpdatedAt time.Time `protobuf:"bytes,11,opt,name=updated_at,json=updatedAt,stdtime" json:"updated_at"`
}

func (m *Client) Reset()                    { *m = Client{} }
func (m *Client) String() string            { return proto.CompactTextString(m) }
func (*Client) ProtoMessage()               {}
func (*Client) Descriptor() ([]byte, []int) { return fileDescriptorClient, []int{0} }

func (m *Client) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Client) GetSecret() string {
	if m != nil {
		return m.Secret
	}
	return ""
}

func (m *Client) GetRedirectURI() string {
	if m != nil {
		return m.RedirectURI
	}
	return ""
}

func (m *Client) GetState() ReviewingState {
	if m != nil {
		return m.State
	}
	return STATE_PENDING
}

func (m *Client) GetOfficialLabeled() bool {
	if m != nil {
		return m.OfficialLabeled
	}
	return false
}

func (m *Client) GetGrants() []GrantType {
	if m != nil {
		return m.Grants
	}
	return nil
}

func (m *Client) GetRights() []Right {
	if m != nil {
		return m.Rights
	}
	return nil
}

func (m *Client) GetCreator() UserIdentifier {
	if m != nil {
		return m.Creator
	}
	return UserIdentifier{}
}

func (m *Client) GetCreatedAt() time.Time {
	if m != nil {
		return m.CreatedAt
	}
	return time.Time{}
}

func (m *Client) GetUpdatedAt() time.Time {
	if m != nil {
		return m.UpdatedAt
	}
	return time.Time{}
}

func init() {
	proto.RegisterType((*Client)(nil), "ttn.v3.Client")
	golang_proto.RegisterType((*Client)(nil), "ttn.v3.Client")
	proto.RegisterEnum("ttn.v3.ReviewingState", ReviewingState_name, ReviewingState_value)
	golang_proto.RegisterEnum("ttn.v3.ReviewingState", ReviewingState_name, ReviewingState_value)
}
func (x ReviewingState) String() string {
	s, ok := ReviewingState_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (this *Client) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*Client)
	if !ok {
		that2, ok := that.(Client)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *Client")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *Client but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *Client but is not nil && this == nil")
	}
	if !this.ClientIdentifier.Equal(&that1.ClientIdentifier) {
		return fmt.Errorf("ClientIdentifier this(%v) Not Equal that(%v)", this.ClientIdentifier, that1.ClientIdentifier)
	}
	if this.Description != that1.Description {
		return fmt.Errorf("Description this(%v) Not Equal that(%v)", this.Description, that1.Description)
	}
	if this.Secret != that1.Secret {
		return fmt.Errorf("Secret this(%v) Not Equal that(%v)", this.Secret, that1.Secret)
	}
	if this.RedirectURI != that1.RedirectURI {
		return fmt.Errorf("RedirectURI this(%v) Not Equal that(%v)", this.RedirectURI, that1.RedirectURI)
	}
	if this.State != that1.State {
		return fmt.Errorf("State this(%v) Not Equal that(%v)", this.State, that1.State)
	}
	if this.OfficialLabeled != that1.OfficialLabeled {
		return fmt.Errorf("OfficialLabeled this(%v) Not Equal that(%v)", this.OfficialLabeled, that1.OfficialLabeled)
	}
	if len(this.Grants) != len(that1.Grants) {
		return fmt.Errorf("Grants this(%v) Not Equal that(%v)", len(this.Grants), len(that1.Grants))
	}
	for i := range this.Grants {
		if this.Grants[i] != that1.Grants[i] {
			return fmt.Errorf("Grants this[%v](%v) Not Equal that[%v](%v)", i, this.Grants[i], i, that1.Grants[i])
		}
	}
	if len(this.Rights) != len(that1.Rights) {
		return fmt.Errorf("Rights this(%v) Not Equal that(%v)", len(this.Rights), len(that1.Rights))
	}
	for i := range this.Rights {
		if this.Rights[i] != that1.Rights[i] {
			return fmt.Errorf("Rights this[%v](%v) Not Equal that[%v](%v)", i, this.Rights[i], i, that1.Rights[i])
		}
	}
	if !this.Creator.Equal(&that1.Creator) {
		return fmt.Errorf("Creator this(%v) Not Equal that(%v)", this.Creator, that1.Creator)
	}
	if !this.CreatedAt.Equal(that1.CreatedAt) {
		return fmt.Errorf("CreatedAt this(%v) Not Equal that(%v)", this.CreatedAt, that1.CreatedAt)
	}
	if !this.UpdatedAt.Equal(that1.UpdatedAt) {
		return fmt.Errorf("UpdatedAt this(%v) Not Equal that(%v)", this.UpdatedAt, that1.UpdatedAt)
	}
	return nil
}
func (this *Client) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*Client)
	if !ok {
		that2, ok := that.(Client)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if !this.ClientIdentifier.Equal(&that1.ClientIdentifier) {
		return false
	}
	if this.Description != that1.Description {
		return false
	}
	if this.Secret != that1.Secret {
		return false
	}
	if this.RedirectURI != that1.RedirectURI {
		return false
	}
	if this.State != that1.State {
		return false
	}
	if this.OfficialLabeled != that1.OfficialLabeled {
		return false
	}
	if len(this.Grants) != len(that1.Grants) {
		return false
	}
	for i := range this.Grants {
		if this.Grants[i] != that1.Grants[i] {
			return false
		}
	}
	if len(this.Rights) != len(that1.Rights) {
		return false
	}
	for i := range this.Rights {
		if this.Rights[i] != that1.Rights[i] {
			return false
		}
	}
	if !this.Creator.Equal(&that1.Creator) {
		return false
	}
	if !this.CreatedAt.Equal(that1.CreatedAt) {
		return false
	}
	if !this.UpdatedAt.Equal(that1.UpdatedAt) {
		return false
	}
	return true
}
func (m *Client) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Client) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintClient(dAtA, i, uint64(m.ClientIdentifier.Size()))
	n1, err := m.ClientIdentifier.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	if len(m.Description) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintClient(dAtA, i, uint64(len(m.Description)))
		i += copy(dAtA[i:], m.Description)
	}
	if len(m.Secret) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintClient(dAtA, i, uint64(len(m.Secret)))
		i += copy(dAtA[i:], m.Secret)
	}
	if len(m.RedirectURI) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintClient(dAtA, i, uint64(len(m.RedirectURI)))
		i += copy(dAtA[i:], m.RedirectURI)
	}
	if m.State != 0 {
		dAtA[i] = 0x28
		i++
		i = encodeVarintClient(dAtA, i, uint64(m.State))
	}
	if m.OfficialLabeled {
		dAtA[i] = 0x30
		i++
		if m.OfficialLabeled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if len(m.Grants) > 0 {
		dAtA3 := make([]byte, len(m.Grants)*10)
		var j2 int
		for _, num := range m.Grants {
			for num >= 1<<7 {
				dAtA3[j2] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j2++
			}
			dAtA3[j2] = uint8(num)
			j2++
		}
		dAtA[i] = 0x3a
		i++
		i = encodeVarintClient(dAtA, i, uint64(j2))
		i += copy(dAtA[i:], dAtA3[:j2])
	}
	if len(m.Rights) > 0 {
		dAtA5 := make([]byte, len(m.Rights)*10)
		var j4 int
		for _, num := range m.Rights {
			for num >= 1<<7 {
				dAtA5[j4] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j4++
			}
			dAtA5[j4] = uint8(num)
			j4++
		}
		dAtA[i] = 0x42
		i++
		i = encodeVarintClient(dAtA, i, uint64(j4))
		i += copy(dAtA[i:], dAtA5[:j4])
	}
	dAtA[i] = 0x4a
	i++
	i = encodeVarintClient(dAtA, i, uint64(m.Creator.Size()))
	n6, err := m.Creator.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n6
	dAtA[i] = 0x52
	i++
	i = encodeVarintClient(dAtA, i, uint64(github_com_gogo_protobuf_types.SizeOfStdTime(m.CreatedAt)))
	n7, err := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.CreatedAt, dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n7
	dAtA[i] = 0x5a
	i++
	i = encodeVarintClient(dAtA, i, uint64(github_com_gogo_protobuf_types.SizeOfStdTime(m.UpdatedAt)))
	n8, err := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.UpdatedAt, dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n8
	return i, nil
}

func encodeVarintClient(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func NewPopulatedClient(r randyClient, easy bool) *Client {
	this := &Client{}
	v1 := NewPopulatedClientIdentifier(r, easy)
	this.ClientIdentifier = *v1
	this.Description = randStringClient(r)
	this.Secret = randStringClient(r)
	this.RedirectURI = randStringClient(r)
	this.State = ReviewingState([]int32{0, 1, 2}[r.Intn(3)])
	this.OfficialLabeled = bool(r.Intn(2) == 0)
	v2 := r.Intn(10)
	this.Grants = make([]GrantType, v2)
	for i := 0; i < v2; i++ {
		this.Grants[i] = GrantType([]int32{0, 1, 2}[r.Intn(3)])
	}
	v3 := r.Intn(10)
	this.Rights = make([]Right, v3)
	for i := 0; i < v3; i++ {
		this.Rights[i] = Right([]int32{0, 1, 2, 3, 4, 5, 6, 9, 10, 12, 13, 14, 15, 16, 31, 32, 33, 34, 35, 36, 37, 38, 39, 51, 52, 53, 54, 55, 56, 57, 58}[r.Intn(31)])
	}
	v4 := NewPopulatedUserIdentifier(r, easy)
	this.Creator = *v4
	v5 := github_com_gogo_protobuf_types.NewPopulatedStdTime(r, easy)
	this.CreatedAt = *v5
	v6 := github_com_gogo_protobuf_types.NewPopulatedStdTime(r, easy)
	this.UpdatedAt = *v6
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

type randyClient interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneClient(r randyClient) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringClient(r randyClient) string {
	v7 := r.Intn(100)
	tmps := make([]rune, v7)
	for i := 0; i < v7; i++ {
		tmps[i] = randUTF8RuneClient(r)
	}
	return string(tmps)
}
func randUnrecognizedClient(r randyClient, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldClient(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldClient(dAtA []byte, r randyClient, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateClient(dAtA, uint64(key))
		v8 := r.Int63()
		if r.Intn(2) == 0 {
			v8 *= -1
		}
		dAtA = encodeVarintPopulateClient(dAtA, uint64(v8))
	case 1:
		dAtA = encodeVarintPopulateClient(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateClient(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateClient(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateClient(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateClient(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(v&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *Client) Size() (n int) {
	var l int
	_ = l
	l = m.ClientIdentifier.Size()
	n += 1 + l + sovClient(uint64(l))
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovClient(uint64(l))
	}
	l = len(m.Secret)
	if l > 0 {
		n += 1 + l + sovClient(uint64(l))
	}
	l = len(m.RedirectURI)
	if l > 0 {
		n += 1 + l + sovClient(uint64(l))
	}
	if m.State != 0 {
		n += 1 + sovClient(uint64(m.State))
	}
	if m.OfficialLabeled {
		n += 2
	}
	if len(m.Grants) > 0 {
		l = 0
		for _, e := range m.Grants {
			l += sovClient(uint64(e))
		}
		n += 1 + sovClient(uint64(l)) + l
	}
	if len(m.Rights) > 0 {
		l = 0
		for _, e := range m.Rights {
			l += sovClient(uint64(e))
		}
		n += 1 + sovClient(uint64(l)) + l
	}
	l = m.Creator.Size()
	n += 1 + l + sovClient(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.CreatedAt)
	n += 1 + l + sovClient(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.UpdatedAt)
	n += 1 + l + sovClient(uint64(l))
	return n
}

func sovClient(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozClient(x uint64) (n int) {
	return sovClient((x << 1) ^ uint64((int64(x) >> 63)))
}
func (m *Client) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowClient
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Client: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Client: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClientIdentifier", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClient
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthClient
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ClientIdentifier.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClient
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthClient
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Secret", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClient
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthClient
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Secret = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RedirectURI", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClient
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthClient
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RedirectURI = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field State", wireType)
			}
			m.State = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClient
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.State |= (ReviewingState(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OfficialLabeled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClient
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.OfficialLabeled = bool(v != 0)
		case 7:
			if wireType == 0 {
				var v GrantType
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowClient
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= (GrantType(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.Grants = append(m.Grants, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowClient
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= (int(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthClient
				}
				postIndex := iNdEx + packedLen
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				for iNdEx < postIndex {
					var v GrantType
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowClient
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= (GrantType(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.Grants = append(m.Grants, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Grants", wireType)
			}
		case 8:
			if wireType == 0 {
				var v Right
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowClient
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= (Right(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.Rights = append(m.Rights, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowClient
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= (int(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthClient
				}
				postIndex := iNdEx + packedLen
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				for iNdEx < postIndex {
					var v Right
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowClient
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= (Right(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.Rights = append(m.Rights, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Rights", wireType)
			}
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClient
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthClient
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Creator.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatedAt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClient
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthClient
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.CreatedAt, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UpdatedAt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClient
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthClient
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.UpdatedAt, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipClient(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthClient
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipClient(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowClient
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowClient
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowClient
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthClient
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowClient
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipClient(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthClient = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowClient   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("github.com/TheThingsNetwork/ttn/api/client.proto", fileDescriptorClient)
}
func init() {
	golang_proto.RegisterFile("github.com/TheThingsNetwork/ttn/api/client.proto", fileDescriptorClient)
}

var fileDescriptorClient = []byte{
	// 630 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0x3d, 0x4c, 0xdb, 0x40,
	0x14, 0xbe, 0x0b, 0x21, 0xc0, 0xa5, 0x84, 0x70, 0x03, 0xb2, 0x32, 0xbc, 0x44, 0x95, 0x2a, 0x85,
	0xaa, 0x75, 0xaa, 0xa0, 0x76, 0x0f, 0x10, 0x21, 0xaa, 0x8a, 0xd2, 0x23, 0x74, 0xe8, 0x12, 0x39,
	0xf1, 0xc5, 0x39, 0x11, 0x6c, 0xeb, 0x7c, 0x01, 0x75, 0x63, 0x64, 0x64, 0x6b, 0xb7, 0x56, 0xea,
	0xc2, 0xc8, 0xc8, 0xc8, 0xc8, 0xc8, 0xc8, 0x44, 0xc9, 0x79, 0x61, 0x64, 0x64, 0xac, 0xe2, 0x1f,
	0xa0, 0x53, 0xe9, 0xe4, 0xbb, 0xef, 0xe7, 0xf9, 0xf3, 0xf7, 0x64, 0xf2, 0xc6, 0x11, 0xaa, 0x3f,
	0xec, 0x98, 0x5d, 0x6f, 0xb7, 0xd6, 0xea, 0xf3, 0x56, 0x5f, 0xb8, 0x4e, 0xb0, 0xc1, 0xd5, 0xbe,
	0x27, 0x77, 0x6a, 0x4a, 0xb9, 0x35, 0xcb, 0x17, 0xb5, 0xee, 0x40, 0x70, 0x57, 0x99, 0xbe, 0xf4,
	0x94, 0x47, 0x73, 0x4a, 0xb9, 0xe6, 0xde, 0x52, 0xe9, 0xf5, 0x23, 0xa7, 0xe3, 0x39, 0x5e, 0x2d,
	0xa2, 0x3b, 0xc3, 0x5e, 0x74, 0x8b, 0x2e, 0xd1, 0x29, 0xb6, 0x95, 0xcc, 0xa7, 0xbc, 0xc8, 0x1a,
	0xaa, 0x7e, 0xa2, 0x7f, 0xfb, 0x14, 0xbd, 0xb0, 0xb9, 0xab, 0x44, 0x4f, 0x70, 0x19, 0x24, 0xb6,
	0x27, 0x7d, 0x8f, 0x14, 0x4e, 0x5f, 0xa5, 0x8e, 0xb2, 0xe3, 0x79, 0xce, 0x80, 0x3f, 0xc4, 0x57,
	0x62, 0x97, 0x07, 0xca, 0xda, 0xf5, 0x63, 0xc1, 0xf3, 0x6f, 0x59, 0x92, 0x5b, 0x89, 0x1a, 0xa0,
	0x75, 0x92, 0x11, 0xb6, 0x81, 0x2b, 0xb8, 0x9a, 0xaf, 0x1b, 0x66, 0x5c, 0x84, 0x19, 0x73, 0xeb,
	0xf7, 0x51, 0x96, 0xa7, 0xcf, 0xaf, 0xca, 0xe8, 0xe2, 0xaa, 0x8c, 0x59, 0x46, 0xd8, 0xb4, 0x42,
	0xf2, 0x36, 0x0f, 0xba, 0x52, 0xf8, 0x4a, 0x78, 0xae, 0x91, 0xa9, 0xe0, 0xea, 0x0c, 0x7b, 0x0c,
	0xd1, 0x05, 0x92, 0x0b, 0x78, 0x57, 0x72, 0x65, 0x4c, 0x44, 0x64, 0x72, 0xa3, 0x75, 0xf2, 0x4c,
	0x72, 0x5b, 0x48, 0xde, 0x55, 0xed, 0xa1, 0x14, 0x46, 0x76, 0xcc, 0x2e, 0xcf, 0xe9, 0xab, 0x72,
	0x9e, 0x25, 0xf8, 0x36, 0x5b, 0x67, 0xf9, 0x54, 0xb4, 0x2d, 0x05, 0x7d, 0x45, 0x26, 0x03, 0x65,
	0x29, 0x6e, 0x4c, 0x56, 0x70, 0xb5, 0x50, 0x5f, 0x48, 0x43, 0x32, 0xbe, 0x27, 0xf8, 0xbe, 0x70,
	0x9d, 0xad, 0x31, 0xcb, 0x62, 0x11, 0x5d, 0x24, 0x45, 0xaf, 0xd7, 0x13, 0x5d, 0x61, 0x0d, 0xda,
	0x03, 0xab, 0xc3, 0x07, 0xdc, 0x36, 0x72, 0x15, 0x5c, 0x9d, 0x66, 0x73, 0x29, 0xfe, 0x21, 0x86,
	0xe9, 0x22, 0xc9, 0x39, 0xd2, 0x72, 0x55, 0x60, 0x4c, 0x55, 0x26, 0xaa, 0x85, 0xfa, 0x7c, 0x3a,
	0x79, 0x6d, 0x8c, 0xb6, 0xbe, 0xfa, 0x9c, 0x25, 0x02, 0xfa, 0x82, 0xe4, 0xe2, 0x86, 0x8d, 0xe9,
	0x48, 0x3a, 0x7b, 0x1f, 0x62, 0x8c, 0xb2, 0x84, 0xa4, 0xef, 0xc8, 0x54, 0x57, 0x72, 0x4b, 0x79,
	0xd2, 0x98, 0x89, 0x1a, 0xbd, 0x0f, 0xbb, 0x1d, 0x70, 0xf9, 0xa8, 0xcf, 0xec, 0xb8, 0x4f, 0x96,
	0x8a, 0xe9, 0x0a, 0x21, 0xd1, 0x91, 0xdb, 0x6d, 0x4b, 0x19, 0x24, 0xb2, 0x96, 0xcc, 0x78, 0x8b,
	0x66, 0xba, 0x45, 0xb3, 0x95, 0x6e, 0x31, 0x5e, 0xc7, 0xd1, 0xef, 0x32, 0x66, 0x33, 0x89, 0xaf,
	0xa1, 0xc6, 0x43, 0x86, 0xbe, 0x9d, 0x0e, 0xc9, 0xff, 0xcf, 0x90, 0xc4, 0xd7, 0x50, 0x2f, 0x3f,
	0x91, 0xc2, 0xdf, 0xbd, 0xd2, 0x79, 0x32, 0xbb, 0xd5, 0x6a, 0xb4, 0x9a, 0xed, 0xcd, 0xe6, 0xc6,
	0xea, 0xfa, 0xc6, 0x5a, 0x11, 0x51, 0x4a, 0x0a, 0x31, 0xd4, 0xd8, 0xdc, 0x64, 0x1f, 0x3f, 0x37,
	0x57, 0x8b, 0xf8, 0x01, 0x63, 0xcd, 0xf7, 0xcd, 0x95, 0x56, 0x73, 0xb5, 0x98, 0x29, 0x65, 0x0f,
	0x7f, 0x01, 0x5a, 0xfe, 0x81, 0xcf, 0x47, 0x80, 0x2f, 0x46, 0x80, 0x2f, 0x47, 0x80, 0xaf, 0x47,
	0x80, 0x6f, 0x46, 0x80, 0x6e, 0x47, 0x80, 0xee, 0x46, 0x80, 0x0f, 0x34, 0xa0, 0x43, 0x0d, 0xe8,
	0x58, 0x03, 0x3e, 0xd1, 0x80, 0x4e, 0x35, 0xe0, 0x33, 0x0d, 0xf8, 0x5c, 0x03, 0xbe, 0xd0, 0x80,
	0x2f, 0x35, 0xa0, 0x6b, 0x0d, 0xf8, 0x46, 0x03, 0xba, 0xd5, 0x80, 0xef, 0x34, 0xa0, 0x83, 0x10,
	0xd0, 0x61, 0x08, 0xf8, 0x28, 0x04, 0xf4, 0x3d, 0x04, 0xfc, 0x33, 0x04, 0x74, 0x1c, 0x02, 0x3a,
	0x09, 0x01, 0x9f, 0x86, 0x80, 0xcf, 0x42, 0xc0, 0x5f, 0x16, 0xff, 0xf5, 0xdb, 0xf8, 0x3b, 0xce,
	0xf8, 0xe9, 0x77, 0x3a, 0xb9, 0xa8, 0x9d, 0xa5, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x3b, 0x72,
	0x46, 0x6c, 0x3a, 0x04, 0x00, 0x00,
}
