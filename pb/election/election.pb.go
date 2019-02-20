// Code generated by protoc-gen-go. DO NOT EDIT.
// source: election.proto

package election

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import duration "github.com/golang/protobuf/ptypes/duration"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Vote struct {
	Voter                []byte               `protobuf:"bytes,1,opt,name=voter,proto3" json:"voter,omitempty"`
	Candidate            []byte               `protobuf:"bytes,2,opt,name=candidate,proto3" json:"candidate,omitempty"`
	Amount               []byte               `protobuf:"bytes,3,opt,name=amount,proto3" json:"amount,omitempty"`
	WeightedAmount       []byte               `protobuf:"bytes,4,opt,name=weightedAmount,proto3" json:"weightedAmount,omitempty"`
	StartTime            *timestamp.Timestamp `protobuf:"bytes,5,opt,name=startTime,proto3" json:"startTime,omitempty"`
	Duration             *duration.Duration   `protobuf:"bytes,6,opt,name=duration,proto3" json:"duration,omitempty"`
	Decay                bool                 `protobuf:"varint,7,opt,name=decay,proto3" json:"decay,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Vote) Reset()         { *m = Vote{} }
func (m *Vote) String() string { return proto.CompactTextString(m) }
func (*Vote) ProtoMessage()    {}
func (*Vote) Descriptor() ([]byte, []int) {
	return fileDescriptor_election_6c8a4ffb36b48953, []int{0}
}
func (m *Vote) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Vote.Unmarshal(m, b)
}
func (m *Vote) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Vote.Marshal(b, m, deterministic)
}
func (dst *Vote) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Vote.Merge(dst, src)
}
func (m *Vote) XXX_Size() int {
	return xxx_messageInfo_Vote.Size(m)
}
func (m *Vote) XXX_DiscardUnknown() {
	xxx_messageInfo_Vote.DiscardUnknown(m)
}

var xxx_messageInfo_Vote proto.InternalMessageInfo

func (m *Vote) GetVoter() []byte {
	if m != nil {
		return m.Voter
	}
	return nil
}

func (m *Vote) GetCandidate() []byte {
	if m != nil {
		return m.Candidate
	}
	return nil
}

func (m *Vote) GetAmount() []byte {
	if m != nil {
		return m.Amount
	}
	return nil
}

func (m *Vote) GetWeightedAmount() []byte {
	if m != nil {
		return m.WeightedAmount
	}
	return nil
}

func (m *Vote) GetStartTime() *timestamp.Timestamp {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (m *Vote) GetDuration() *duration.Duration {
	if m != nil {
		return m.Duration
	}
	return nil
}

func (m *Vote) GetDecay() bool {
	if m != nil {
		return m.Decay
	}
	return false
}

type VoteList struct {
	Votes                []*Vote  `protobuf:"bytes,1,rep,name=votes,proto3" json:"votes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VoteList) Reset()         { *m = VoteList{} }
func (m *VoteList) String() string { return proto.CompactTextString(m) }
func (*VoteList) ProtoMessage()    {}
func (*VoteList) Descriptor() ([]byte, []int) {
	return fileDescriptor_election_6c8a4ffb36b48953, []int{1}
}
func (m *VoteList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VoteList.Unmarshal(m, b)
}
func (m *VoteList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VoteList.Marshal(b, m, deterministic)
}
func (dst *VoteList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VoteList.Merge(dst, src)
}
func (m *VoteList) XXX_Size() int {
	return xxx_messageInfo_VoteList.Size(m)
}
func (m *VoteList) XXX_DiscardUnknown() {
	xxx_messageInfo_VoteList.DiscardUnknown(m)
}

var xxx_messageInfo_VoteList proto.InternalMessageInfo

func (m *VoteList) GetVotes() []*Vote {
	if m != nil {
		return m.Votes
	}
	return nil
}

type Candidate struct {
	Name                 []byte   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Address              []byte   `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	OperatorAddress      []byte   `protobuf:"bytes,3,opt,name=operatorAddress,proto3" json:"operatorAddress,omitempty"`
	RewardAddress        []byte   `protobuf:"bytes,4,opt,name=rewardAddress,proto3" json:"rewardAddress,omitempty"`
	SelfStakingWeight    uint32   `protobuf:"varint,5,opt,name=selfStakingWeight,proto3" json:"selfStakingWeight,omitempty"`
	Score                []byte   `protobuf:"bytes,6,opt,name=score,proto3" json:"score,omitempty"`
	SelfStakingScore     []byte   `protobuf:"bytes,7,opt,name=selfStakingScore,proto3" json:"selfStakingScore,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Candidate) Reset()         { *m = Candidate{} }
func (m *Candidate) String() string { return proto.CompactTextString(m) }
func (*Candidate) ProtoMessage()    {}
func (*Candidate) Descriptor() ([]byte, []int) {
	return fileDescriptor_election_6c8a4ffb36b48953, []int{2}
}
func (m *Candidate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Candidate.Unmarshal(m, b)
}
func (m *Candidate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Candidate.Marshal(b, m, deterministic)
}
func (dst *Candidate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Candidate.Merge(dst, src)
}
func (m *Candidate) XXX_Size() int {
	return xxx_messageInfo_Candidate.Size(m)
}
func (m *Candidate) XXX_DiscardUnknown() {
	xxx_messageInfo_Candidate.DiscardUnknown(m)
}

var xxx_messageInfo_Candidate proto.InternalMessageInfo

func (m *Candidate) GetName() []byte {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *Candidate) GetAddress() []byte {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *Candidate) GetOperatorAddress() []byte {
	if m != nil {
		return m.OperatorAddress
	}
	return nil
}

func (m *Candidate) GetRewardAddress() []byte {
	if m != nil {
		return m.RewardAddress
	}
	return nil
}

func (m *Candidate) GetSelfStakingWeight() uint32 {
	if m != nil {
		return m.SelfStakingWeight
	}
	return 0
}

func (m *Candidate) GetScore() []byte {
	if m != nil {
		return m.Score
	}
	return nil
}

func (m *Candidate) GetSelfStakingScore() []byte {
	if m != nil {
		return m.SelfStakingScore
	}
	return nil
}

type ElectionResult struct {
	Timestamp            *timestamp.Timestamp `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Delegates            []*Candidate         `protobuf:"bytes,2,rep,name=delegates,proto3" json:"delegates,omitempty"`
	DelegateVotes        []*VoteList          `protobuf:"bytes,3,rep,name=delegateVotes,proto3" json:"delegateVotes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ElectionResult) Reset()         { *m = ElectionResult{} }
func (m *ElectionResult) String() string { return proto.CompactTextString(m) }
func (*ElectionResult) ProtoMessage()    {}
func (*ElectionResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_election_6c8a4ffb36b48953, []int{3}
}
func (m *ElectionResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ElectionResult.Unmarshal(m, b)
}
func (m *ElectionResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ElectionResult.Marshal(b, m, deterministic)
}
func (dst *ElectionResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ElectionResult.Merge(dst, src)
}
func (m *ElectionResult) XXX_Size() int {
	return xxx_messageInfo_ElectionResult.Size(m)
}
func (m *ElectionResult) XXX_DiscardUnknown() {
	xxx_messageInfo_ElectionResult.DiscardUnknown(m)
}

var xxx_messageInfo_ElectionResult proto.InternalMessageInfo

func (m *ElectionResult) GetTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *ElectionResult) GetDelegates() []*Candidate {
	if m != nil {
		return m.Delegates
	}
	return nil
}

func (m *ElectionResult) GetDelegateVotes() []*VoteList {
	if m != nil {
		return m.DelegateVotes
	}
	return nil
}

func init() {
	proto.RegisterType((*Vote)(nil), "election.Vote")
	proto.RegisterType((*VoteList)(nil), "election.VoteList")
	proto.RegisterType((*Candidate)(nil), "election.Candidate")
	proto.RegisterType((*ElectionResult)(nil), "election.ElectionResult")
}

func init() { proto.RegisterFile("election.proto", fileDescriptor_election_6c8a4ffb36b48953) }

var fileDescriptor_election_6c8a4ffb36b48953 = []byte{
	// 423 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0x86, 0xe5, 0x24, 0xcd, 0xc7, 0xb4, 0x09, 0x30, 0x20, 0xb4, 0x44, 0x08, 0xa2, 0xa8, 0x42,
	0x16, 0x42, 0x2e, 0x14, 0x21, 0xf5, 0x5a, 0x01, 0x37, 0x4e, 0xdb, 0xaa, 0x9c, 0xb7, 0xd9, 0xa9,
	0xb1, 0xb0, 0xbd, 0xd1, 0xee, 0x9a, 0x8a, 0x23, 0xbf, 0x8a, 0xdf, 0xc6, 0x0d, 0x79, 0xd6, 0x1f,
	0x24, 0x39, 0x70, 0xf3, 0x3b, 0xf3, 0x8c, 0x34, 0xfb, 0x8c, 0x61, 0x41, 0x39, 0x6d, 0x7c, 0x66,
	0xca, 0x64, 0x6b, 0x8d, 0x37, 0x38, 0x6d, 0xf3, 0xf2, 0x45, 0x6a, 0x4c, 0x9a, 0xd3, 0x19, 0xd7,
	0x6f, 0xab, 0xbb, 0x33, 0x5d, 0x59, 0xd5, 0x93, 0xcb, 0x97, 0xfb, 0x7d, 0x9f, 0x15, 0xe4, 0xbc,
	0x2a, 0xb6, 0x01, 0x58, 0xff, 0x1a, 0xc0, 0xe8, 0xc6, 0x78, 0xc2, 0x27, 0x70, 0xf4, 0xc3, 0x78,
	0xb2, 0x22, 0x5a, 0x45, 0xf1, 0x89, 0x0c, 0x01, 0x9f, 0xc3, 0x6c, 0xa3, 0x4a, 0x9d, 0x69, 0xe5,
	0x49, 0x0c, 0xb8, 0xd3, 0x17, 0xf0, 0x29, 0x8c, 0x55, 0x61, 0xaa, 0xd2, 0x8b, 0x21, 0xb7, 0x9a,
	0x84, 0xaf, 0x60, 0x71, 0x4f, 0x59, 0xfa, 0xcd, 0x93, 0xbe, 0x0c, 0xfd, 0x11, 0xf7, 0xf7, 0xaa,
	0x78, 0x01, 0x33, 0xe7, 0x95, 0xf5, 0xd7, 0x59, 0x41, 0xe2, 0x68, 0x15, 0xc5, 0xc7, 0xe7, 0xcb,
	0x24, 0x6c, 0x9c, 0xb4, 0x1b, 0x27, 0xd7, 0xed, 0xc6, 0xb2, 0x87, 0xf1, 0x03, 0x4c, 0xdb, 0x97,
	0x8a, 0x31, 0x0f, 0x3e, 0x3b, 0x18, 0xfc, 0xd4, 0x00, 0xb2, 0x43, 0xeb, 0x47, 0x6a, 0xda, 0xa8,
	0x9f, 0x62, 0xb2, 0x8a, 0xe2, 0xa9, 0x0c, 0x61, 0xfd, 0x16, 0xa6, 0xb5, 0x82, 0x2f, 0x99, 0xf3,
	0x78, 0x1a, 0x34, 0x38, 0x11, 0xad, 0x86, 0xf1, 0xf1, 0xf9, 0x22, 0xe9, 0xd4, 0xd7, 0x48, 0xd0,
	0xe2, 0xd6, 0x7f, 0x22, 0x98, 0x7d, 0xec, 0x34, 0x20, 0x8c, 0x4a, 0x55, 0x50, 0x63, 0x8e, 0xbf,
	0x51, 0xc0, 0x44, 0x69, 0x6d, 0xc9, 0xb9, 0x46, 0x5b, 0x1b, 0x31, 0x86, 0x07, 0x66, 0x4b, 0x56,
	0x79, 0x63, 0x2f, 0x1b, 0x22, 0xd8, 0xdb, 0x2f, 0xe3, 0x29, 0xcc, 0x2d, 0xdd, 0x2b, 0xab, 0x5b,
	0x2e, 0x58, 0xdc, 0x2d, 0xe2, 0x1b, 0x78, 0xe4, 0x28, 0xbf, 0xbb, 0xf2, 0xea, 0x7b, 0x56, 0xa6,
	0x5f, 0xd9, 0x30, 0xcb, 0x9c, 0xcb, 0xc3, 0x46, 0x6d, 0xc0, 0x6d, 0x8c, 0x25, 0xb6, 0x76, 0x22,
	0x43, 0xc0, 0xd7, 0xf0, 0xf0, 0x1f, 0xf4, 0x8a, 0x81, 0x09, 0x03, 0x07, 0xf5, 0xf5, 0xef, 0x08,
	0x16, 0x9f, 0x1b, 0x29, 0x92, 0x5c, 0x95, 0xf3, 0x1d, 0xbb, 0xff, 0x8a, 0x2d, 0xfc, 0xe7, 0x8e,
	0x1d, 0x8c, 0xef, 0x60, 0xa6, 0x29, 0xa7, 0x54, 0xd5, 0xca, 0x07, 0xac, 0xfc, 0x71, 0xaf, 0xbc,
	0x53, 0x2c, 0x7b, 0x0a, 0x2f, 0x60, 0xde, 0x86, 0x1b, 0xbe, 0xd4, 0x90, 0xc7, 0x70, 0xf7, 0x52,
	0xf5, 0x31, 0xe5, 0x2e, 0x78, 0x3b, 0xe6, 0x5d, 0xde, 0xff, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x59,
	0x64, 0xe0, 0x5d, 0x4f, 0x03, 0x00, 0x00,
}
