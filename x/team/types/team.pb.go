// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: kyve/team/v1beta1/team.proto

package types

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Authority ...
type Authority struct {
	// total inflation rewards is the total amount of rewards the authority has received ever
	TotalRewards uint64 `protobuf:"varint,1,opt,name=total_rewards,json=totalRewards,proto3" json:"total_rewards,omitempty"`
	// claimed is the amount of inflation rewards claimed by the authority
	RewardsClaimed uint64 `protobuf:"varint,2,opt,name=rewards_claimed,json=rewardsClaimed,proto3" json:"rewards_claimed,omitempty"`
}

func (m *Authority) Reset()         { *m = Authority{} }
func (m *Authority) String() string { return proto.CompactTextString(m) }
func (*Authority) ProtoMessage()    {}
func (*Authority) Descriptor() ([]byte, []int) {
	return fileDescriptor_a9a907d008be83cf, []int{0}
}
func (m *Authority) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Authority) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Authority.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Authority) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Authority.Merge(m, src)
}
func (m *Authority) XXX_Size() int {
	return m.Size()
}
func (m *Authority) XXX_DiscardUnknown() {
	xxx_messageInfo_Authority.DiscardUnknown(m)
}

var xxx_messageInfo_Authority proto.InternalMessageInfo

func (m *Authority) GetTotalRewards() uint64 {
	if m != nil {
		return m.TotalRewards
	}
	return 0
}

func (m *Authority) GetRewardsClaimed() uint64 {
	if m != nil {
		return m.RewardsClaimed
	}
	return 0
}

// TeamVestingAccount ...
type TeamVestingAccount struct {
	// id is a unique identify for each vesting account, tied to a single team member.
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// total_allocation is the number of tokens reserved for this team member.
	TotalAllocation uint64 `protobuf:"varint,2,opt,name=total_allocation,json=totalAllocation,proto3" json:"total_allocation,omitempty"`
	// commencement is the unix timestamp of the member's official start date in seconds
	Commencement uint64 `protobuf:"varint,3,opt,name=commencement,proto3" json:"commencement,omitempty"`
	// clawback is a unix timestamp of a clawback in seconds. If timestamp is zero
	// it means that the account has not received a clawback
	Clawback uint64 `protobuf:"varint,4,opt,name=clawback,proto3" json:"clawback,omitempty"`
	// unlocked_claimed is the amount of $KYVE already claimed by the account holder
	UnlockedClaimed uint64 `protobuf:"varint,5,opt,name=unlocked_claimed,json=unlockedClaimed,proto3" json:"unlocked_claimed,omitempty"`
	// the last time the unlocked amount was claimed
	LastClaimedTime uint64 `protobuf:"varint,6,opt,name=last_claimed_time,json=lastClaimedTime,proto3" json:"last_claimed_time,omitempty"`
	// total rewards is the total amount of rewards the account has received ever
	TotalRewards uint64 `protobuf:"varint,7,opt,name=total_rewards,json=totalRewards,proto3" json:"total_rewards,omitempty"`
	// rewards claimed is the amount inflation rewards claimed by account holder
	RewardsClaimed uint64 `protobuf:"varint,8,opt,name=rewards_claimed,json=rewardsClaimed,proto3" json:"rewards_claimed,omitempty"`
}

func (m *TeamVestingAccount) Reset()         { *m = TeamVestingAccount{} }
func (m *TeamVestingAccount) String() string { return proto.CompactTextString(m) }
func (*TeamVestingAccount) ProtoMessage()    {}
func (*TeamVestingAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_a9a907d008be83cf, []int{1}
}
func (m *TeamVestingAccount) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TeamVestingAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TeamVestingAccount.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TeamVestingAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TeamVestingAccount.Merge(m, src)
}
func (m *TeamVestingAccount) XXX_Size() int {
	return m.Size()
}
func (m *TeamVestingAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_TeamVestingAccount.DiscardUnknown(m)
}

var xxx_messageInfo_TeamVestingAccount proto.InternalMessageInfo

func (m *TeamVestingAccount) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *TeamVestingAccount) GetTotalAllocation() uint64 {
	if m != nil {
		return m.TotalAllocation
	}
	return 0
}

func (m *TeamVestingAccount) GetCommencement() uint64 {
	if m != nil {
		return m.Commencement
	}
	return 0
}

func (m *TeamVestingAccount) GetClawback() uint64 {
	if m != nil {
		return m.Clawback
	}
	return 0
}

func (m *TeamVestingAccount) GetUnlockedClaimed() uint64 {
	if m != nil {
		return m.UnlockedClaimed
	}
	return 0
}

func (m *TeamVestingAccount) GetLastClaimedTime() uint64 {
	if m != nil {
		return m.LastClaimedTime
	}
	return 0
}

func (m *TeamVestingAccount) GetTotalRewards() uint64 {
	if m != nil {
		return m.TotalRewards
	}
	return 0
}

func (m *TeamVestingAccount) GetRewardsClaimed() uint64 {
	if m != nil {
		return m.RewardsClaimed
	}
	return 0
}

func init() {
	proto.RegisterType((*Authority)(nil), "kyve.team.v1beta1.Authority")
	proto.RegisterType((*TeamVestingAccount)(nil), "kyve.team.v1beta1.TeamVestingAccount")
}

func init() { proto.RegisterFile("kyve/team/v1beta1/team.proto", fileDescriptor_a9a907d008be83cf) }

var fileDescriptor_a9a907d008be83cf = []byte{
	// 338 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xc1, 0x4e, 0xc2, 0x40,
	0x10, 0x86, 0x69, 0x45, 0xc4, 0x0d, 0x82, 0xec, 0xa9, 0x31, 0xa6, 0x31, 0x78, 0x50, 0x3c, 0xd0,
	0x10, 0x9f, 0x00, 0x89, 0x27, 0x13, 0x0f, 0x84, 0x90, 0xe0, 0x85, 0x6c, 0xb7, 0x13, 0xd8, 0xb4,
	0xdb, 0x25, 0xed, 0x14, 0xe4, 0x2d, 0x7c, 0x18, 0x1f, 0xc2, 0x23, 0x47, 0x8f, 0x06, 0x5e, 0xc4,
	0xb0, 0xbb, 0x10, 0x13, 0x3d, 0x78, 0x9c, 0xef, 0xff, 0x33, 0xf3, 0xe7, 0x1f, 0x72, 0x19, 0xaf,
	0x16, 0x10, 0x20, 0x30, 0x19, 0x2c, 0xba, 0x21, 0x20, 0xeb, 0xea, 0xa1, 0x33, 0xcf, 0x14, 0x2a,
	0xda, 0xdc, 0xa9, 0x1d, 0x0d, 0xac, 0xda, 0x1a, 0x93, 0xd3, 0x5e, 0x81, 0x33, 0x95, 0x09, 0x5c,
	0xd1, 0x6b, 0x72, 0x86, 0x0a, 0x59, 0x32, 0xc9, 0x60, 0xc9, 0xb2, 0x28, 0xf7, 0x9c, 0x2b, 0xe7,
	0xb6, 0x3c, 0xa8, 0x69, 0x38, 0x30, 0x8c, 0xde, 0x90, 0x86, 0x95, 0x27, 0x3c, 0x61, 0x42, 0x42,
	0xe4, 0xb9, 0xda, 0x56, 0xb7, 0xb8, 0x6f, 0x68, 0xeb, 0xdd, 0x25, 0x74, 0x08, 0x4c, 0x8e, 0x20,
	0x47, 0x91, 0x4e, 0x7b, 0x9c, 0xab, 0x22, 0x45, 0x5a, 0x27, 0xae, 0x88, 0xec, 0x66, 0x57, 0x44,
	0xb4, 0x4d, 0xce, 0xcd, 0x51, 0x96, 0x24, 0x8a, 0x33, 0x14, 0x2a, 0xb5, 0x0b, 0x1b, 0x9a, 0xf7,
	0x0e, 0x98, 0xb6, 0x48, 0x8d, 0x2b, 0x29, 0x21, 0xe5, 0x20, 0x21, 0x45, 0xef, 0xc8, 0xc4, 0xfb,
	0xc9, 0xe8, 0x05, 0xa9, 0xf2, 0x84, 0x2d, 0x43, 0xc6, 0x63, 0xaf, 0xac, 0xf5, 0xc3, 0xbc, 0x3b,
	0x55, 0xa4, 0x89, 0xe2, 0x31, 0x44, 0x87, 0xec, 0xc7, 0xe6, 0xd4, 0x9e, 0xdb, 0xf0, 0xf4, 0x8e,
	0x34, 0x13, 0x96, 0xe3, 0xde, 0x36, 0x41, 0x21, 0xc1, 0xab, 0x18, 0xef, 0x4e, 0xb0, 0xbe, 0xa1,
	0x90, 0xf0, 0xbb, 0xb6, 0x93, 0xff, 0xd5, 0x56, 0xfd, 0xab, 0xb6, 0x87, 0xfe, 0xc7, 0xc6, 0x77,
	0xd6, 0x1b, 0xdf, 0xf9, 0xda, 0xf8, 0xce, 0xdb, 0xd6, 0x2f, 0xad, 0xb7, 0x7e, 0xe9, 0x73, 0xeb,
	0x97, 0x5e, 0xda, 0x53, 0x81, 0xb3, 0x22, 0xec, 0x70, 0x25, 0x83, 0xa7, 0xf1, 0xe8, 0xf1, 0x19,
	0x70, 0xa9, 0xb2, 0x38, 0xe0, 0x33, 0x26, 0xd2, 0xe0, 0xd5, 0xbc, 0x1d, 0x57, 0x73, 0xc8, 0xc3,
	0x8a, 0x7e, 0xf8, 0xfd, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x7f, 0xc5, 0x6b, 0x77, 0x10, 0x02,
	0x00, 0x00,
}

func (m *Authority) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Authority) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Authority) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.RewardsClaimed != 0 {
		i = encodeVarintTeam(dAtA, i, uint64(m.RewardsClaimed))
		i--
		dAtA[i] = 0x10
	}
	if m.TotalRewards != 0 {
		i = encodeVarintTeam(dAtA, i, uint64(m.TotalRewards))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *TeamVestingAccount) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TeamVestingAccount) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TeamVestingAccount) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.RewardsClaimed != 0 {
		i = encodeVarintTeam(dAtA, i, uint64(m.RewardsClaimed))
		i--
		dAtA[i] = 0x40
	}
	if m.TotalRewards != 0 {
		i = encodeVarintTeam(dAtA, i, uint64(m.TotalRewards))
		i--
		dAtA[i] = 0x38
	}
	if m.LastClaimedTime != 0 {
		i = encodeVarintTeam(dAtA, i, uint64(m.LastClaimedTime))
		i--
		dAtA[i] = 0x30
	}
	if m.UnlockedClaimed != 0 {
		i = encodeVarintTeam(dAtA, i, uint64(m.UnlockedClaimed))
		i--
		dAtA[i] = 0x28
	}
	if m.Clawback != 0 {
		i = encodeVarintTeam(dAtA, i, uint64(m.Clawback))
		i--
		dAtA[i] = 0x20
	}
	if m.Commencement != 0 {
		i = encodeVarintTeam(dAtA, i, uint64(m.Commencement))
		i--
		dAtA[i] = 0x18
	}
	if m.TotalAllocation != 0 {
		i = encodeVarintTeam(dAtA, i, uint64(m.TotalAllocation))
		i--
		dAtA[i] = 0x10
	}
	if m.Id != 0 {
		i = encodeVarintTeam(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintTeam(dAtA []byte, offset int, v uint64) int {
	offset -= sovTeam(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Authority) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.TotalRewards != 0 {
		n += 1 + sovTeam(uint64(m.TotalRewards))
	}
	if m.RewardsClaimed != 0 {
		n += 1 + sovTeam(uint64(m.RewardsClaimed))
	}
	return n
}

func (m *TeamVestingAccount) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovTeam(uint64(m.Id))
	}
	if m.TotalAllocation != 0 {
		n += 1 + sovTeam(uint64(m.TotalAllocation))
	}
	if m.Commencement != 0 {
		n += 1 + sovTeam(uint64(m.Commencement))
	}
	if m.Clawback != 0 {
		n += 1 + sovTeam(uint64(m.Clawback))
	}
	if m.UnlockedClaimed != 0 {
		n += 1 + sovTeam(uint64(m.UnlockedClaimed))
	}
	if m.LastClaimedTime != 0 {
		n += 1 + sovTeam(uint64(m.LastClaimedTime))
	}
	if m.TotalRewards != 0 {
		n += 1 + sovTeam(uint64(m.TotalRewards))
	}
	if m.RewardsClaimed != 0 {
		n += 1 + sovTeam(uint64(m.RewardsClaimed))
	}
	return n
}

func sovTeam(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTeam(x uint64) (n int) {
	return sovTeam(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Authority) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTeam
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Authority: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Authority: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalRewards", wireType)
			}
			m.TotalRewards = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTeam
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TotalRewards |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RewardsClaimed", wireType)
			}
			m.RewardsClaimed = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTeam
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RewardsClaimed |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTeam(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTeam
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
func (m *TeamVestingAccount) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTeam
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TeamVestingAccount: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TeamVestingAccount: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTeam
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalAllocation", wireType)
			}
			m.TotalAllocation = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTeam
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TotalAllocation |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Commencement", wireType)
			}
			m.Commencement = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTeam
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Commencement |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Clawback", wireType)
			}
			m.Clawback = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTeam
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Clawback |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UnlockedClaimed", wireType)
			}
			m.UnlockedClaimed = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTeam
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UnlockedClaimed |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastClaimedTime", wireType)
			}
			m.LastClaimedTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTeam
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LastClaimedTime |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalRewards", wireType)
			}
			m.TotalRewards = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTeam
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TotalRewards |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RewardsClaimed", wireType)
			}
			m.RewardsClaimed = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTeam
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RewardsClaimed |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTeam(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTeam
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
func skipTeam(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTeam
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
					return 0, ErrIntOverflowTeam
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTeam
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
			if length < 0 {
				return 0, ErrInvalidLengthTeam
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTeam
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTeam
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTeam        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTeam          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTeam = fmt.Errorf("proto: unexpected end of group")
)
