// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pylons/execution.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
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

type ItemRecord struct {
	ID      string           `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Doubles []DoubleKeyValue `protobuf:"bytes,2,rep,name=doubles,proto3" json:"doubles"`
	Longs   []LongKeyValue   `protobuf:"bytes,3,rep,name=longs,proto3" json:"longs"`
	Strings []StringKeyValue `protobuf:"bytes,4,rep,name=strings,proto3" json:"strings"`
}

func (m *ItemRecord) Reset()         { *m = ItemRecord{} }
func (m *ItemRecord) String() string { return proto.CompactTextString(m) }
func (*ItemRecord) ProtoMessage()    {}
func (*ItemRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_00479ca2f8a1ba2c, []int{0}
}
func (m *ItemRecord) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ItemRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ItemRecord.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ItemRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ItemRecord.Merge(m, src)
}
func (m *ItemRecord) XXX_Size() int {
	return m.Size()
}
func (m *ItemRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_ItemRecord.DiscardUnknown(m)
}

var xxx_messageInfo_ItemRecord proto.InternalMessageInfo

func (m *ItemRecord) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *ItemRecord) GetDoubles() []DoubleKeyValue {
	if m != nil {
		return m.Doubles
	}
	return nil
}

func (m *ItemRecord) GetLongs() []LongKeyValue {
	if m != nil {
		return m.Longs
	}
	return nil
}

func (m *ItemRecord) GetStrings() []StringKeyValue {
	if m != nil {
		return m.Strings
	}
	return nil
}

type Execution struct {
	Creator             string                                   `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	ID                  string                                   `protobuf:"bytes,2,opt,name=ID,proto3" json:"ID,omitempty"`
	RecipeID            string                                   `protobuf:"bytes,3,opt,name=recipeID,proto3" json:"recipeID,omitempty"`
	CookbookID          string                                   `protobuf:"bytes,4,opt,name=cookbookID,proto3" json:"cookbookID,omitempty"`
	RecipeVersion       string                                   `protobuf:"bytes,5,opt,name=recipeVersion,proto3" json:"recipeVersion,omitempty"`
	NodeVersion         uint64                                   `protobuf:"varint,6,opt,name=nodeVersion,proto3" json:"nodeVersion,omitempty"`
	BlockHeight         int64                                    `protobuf:"varint,7,opt,name=blockHeight,proto3" json:"blockHeight,omitempty"`
	ItemInputs          []ItemRecord                             `protobuf:"bytes,8,rep,name=itemInputs,proto3" json:"itemInputs"`
	CoinInputs          github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,9,rep,name=coinInputs,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"coinInputs"`
	CoinOutputs         github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,10,rep,name=coinOutputs,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"coinOutputs"`
	ItemOutputIDs       []string                                 `protobuf:"bytes,11,rep,name=itemOutputIDs,proto3" json:"itemOutputIDs,omitempty"`
	ItemModifyOutputIDs []string                                 `protobuf:"bytes,12,rep,name=itemModifyOutputIDs,proto3" json:"itemModifyOutputIDs,omitempty"`
}

func (m *Execution) Reset()         { *m = Execution{} }
func (m *Execution) String() string { return proto.CompactTextString(m) }
func (*Execution) ProtoMessage()    {}
func (*Execution) Descriptor() ([]byte, []int) {
	return fileDescriptor_00479ca2f8a1ba2c, []int{1}
}
func (m *Execution) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Execution) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Execution.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Execution) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Execution.Merge(m, src)
}
func (m *Execution) XXX_Size() int {
	return m.Size()
}
func (m *Execution) XXX_DiscardUnknown() {
	xxx_messageInfo_Execution.DiscardUnknown(m)
}

var xxx_messageInfo_Execution proto.InternalMessageInfo

func (m *Execution) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *Execution) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *Execution) GetRecipeID() string {
	if m != nil {
		return m.RecipeID
	}
	return ""
}

func (m *Execution) GetCookbookID() string {
	if m != nil {
		return m.CookbookID
	}
	return ""
}

func (m *Execution) GetRecipeVersion() string {
	if m != nil {
		return m.RecipeVersion
	}
	return ""
}

func (m *Execution) GetNodeVersion() uint64 {
	if m != nil {
		return m.NodeVersion
	}
	return 0
}

func (m *Execution) GetBlockHeight() int64 {
	if m != nil {
		return m.BlockHeight
	}
	return 0
}

func (m *Execution) GetItemInputs() []ItemRecord {
	if m != nil {
		return m.ItemInputs
	}
	return nil
}

func (m *Execution) GetCoinInputs() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.CoinInputs
	}
	return nil
}

func (m *Execution) GetCoinOutputs() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.CoinOutputs
	}
	return nil
}

func (m *Execution) GetItemOutputIDs() []string {
	if m != nil {
		return m.ItemOutputIDs
	}
	return nil
}

func (m *Execution) GetItemModifyOutputIDs() []string {
	if m != nil {
		return m.ItemModifyOutputIDs
	}
	return nil
}

func init() {
	proto.RegisterType((*ItemRecord)(nil), "Pylonstech.pylons.pylons.ItemRecord")
	proto.RegisterType((*Execution)(nil), "Pylonstech.pylons.pylons.Execution")
}

func init() { proto.RegisterFile("pylons/execution.proto", fileDescriptor_00479ca2f8a1ba2c) }

var fileDescriptor_00479ca2f8a1ba2c = []byte{
	// 533 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x53, 0xcb, 0x6e, 0xd3, 0x40,
	0x14, 0x8d, 0x93, 0xf4, 0x91, 0x09, 0x45, 0x62, 0x8a, 0xd0, 0x90, 0x85, 0x6b, 0x55, 0x15, 0xb2,
	0x10, 0xb5, 0x29, 0x48, 0x7c, 0x40, 0x30, 0xa8, 0xe6, 0x21, 0x90, 0x91, 0xba, 0x60, 0x17, 0x4f,
	0x06, 0x67, 0x94, 0xd8, 0xd7, 0xf2, 0x8c, 0x51, 0xf3, 0x17, 0x7c, 0x07, 0x5f, 0xd2, 0x65, 0x97,
	0xac, 0x00, 0x25, 0x0b, 0x7e, 0x80, 0x0f, 0x40, 0xf3, 0x48, 0xea, 0x56, 0x44, 0x6c, 0x58, 0xcd,
	0xdc, 0x33, 0xe7, 0x9e, 0xb9, 0x73, 0xee, 0x1d, 0x74, 0xaf, 0x9c, 0xcf, 0xa0, 0x10, 0x21, 0x3b,
	0x67, 0xb4, 0x96, 0x1c, 0x8a, 0xa0, 0xac, 0x40, 0x02, 0x26, 0xef, 0x35, 0x2e, 0x19, 0x9d, 0x04,
	0x86, 0x62, 0x97, 0xc1, 0xdd, 0x0c, 0x32, 0xd0, 0xa4, 0x50, 0xed, 0x0c, 0x7f, 0xe0, 0x52, 0x10,
	0x39, 0x88, 0x30, 0x1d, 0x09, 0x16, 0x7e, 0x3e, 0x49, 0x99, 0x1c, 0x9d, 0x84, 0x14, 0xb8, 0xd5,
	0x1b, 0xdc, 0xb1, 0xf7, 0x70, 0xc9, 0x72, 0x0b, 0xed, 0x5b, 0xa8, 0x62, 0x94, 0x97, 0xcc, 0x80,
	0x87, 0xbf, 0x1d, 0x84, 0x62, 0xc9, 0xf2, 0x84, 0x51, 0xa8, 0xc6, 0xf8, 0x36, 0x6a, 0xc7, 0x11,
	0x71, 0x3c, 0xc7, 0xef, 0x25, 0xed, 0x38, 0xc2, 0xa7, 0x68, 0x67, 0x0c, 0x75, 0x3a, 0x63, 0x82,
	0xb4, 0xbd, 0x8e, 0xdf, 0x7f, 0xe2, 0x07, 0x9b, 0x0a, 0x0d, 0x22, 0x4d, 0x7c, 0xcd, 0xe6, 0x67,
	0xa3, 0x59, 0xcd, 0x86, 0xdd, 0x8b, 0xef, 0x07, 0xad, 0x64, 0x95, 0x8e, 0x87, 0x68, 0x6b, 0x06,
	0x45, 0x26, 0x48, 0x47, 0xeb, 0x3c, 0xd8, 0xac, 0xf3, 0x06, 0x8a, 0xec, 0x86, 0x8a, 0x49, 0x55,
	0xd5, 0x08, 0x59, 0x71, 0xa5, 0xd2, 0xfd, 0x57, 0x35, 0x1f, 0x34, 0xf1, 0x66, 0x35, 0x36, 0xfd,
	0xf0, 0x57, 0x17, 0xf5, 0x5e, 0xac, 0x5a, 0x80, 0x09, 0xda, 0xa1, 0x15, 0x1b, 0x49, 0xa8, 0xec,
	0xd3, 0x57, 0xa1, 0xf5, 0xa3, 0xbd, 0xf6, 0x63, 0x80, 0x76, 0x8d, 0x7d, 0x71, 0x44, 0x3a, 0x1a,
	0x5d, 0xc7, 0xd8, 0x45, 0x88, 0x02, 0x4c, 0x53, 0x80, 0x69, 0x1c, 0x91, 0xae, 0x3e, 0x6d, 0x20,
	0xf8, 0x08, 0xed, 0x19, 0xee, 0x19, 0xab, 0x04, 0x87, 0x82, 0x6c, 0x69, 0xca, 0x75, 0x10, 0x7b,
	0xa8, 0x5f, 0xc0, 0x78, 0xcd, 0xd9, 0xf6, 0x1c, 0xbf, 0x9b, 0x34, 0x21, 0xc5, 0x48, 0x67, 0x40,
	0xa7, 0xa7, 0x8c, 0x67, 0x13, 0x49, 0x76, 0x3c, 0xc7, 0xef, 0x24, 0x4d, 0x08, 0xbf, 0x42, 0x48,
	0xf5, 0x3d, 0x2e, 0xca, 0x5a, 0x0a, 0xb2, 0xab, 0xad, 0x3a, 0xda, 0x6c, 0xd5, 0x55, 0xff, 0xad,
	0x4d, 0x8d, 0x6c, 0x3c, 0x55, 0xaf, 0xe2, 0x85, 0xd5, 0xea, 0x69, 0xad, 0xfb, 0x81, 0x99, 0xbe,
	0x40, 0x4d, 0x5f, 0x60, 0xa7, 0x2f, 0x78, 0x0e, 0xbc, 0x18, 0x3e, 0x56, 0x02, 0x5f, 0x7f, 0x1c,
	0xf8, 0x19, 0x97, 0x93, 0x3a, 0x0d, 0x28, 0xe4, 0xa1, 0x1d, 0x55, 0xb3, 0x1c, 0x8b, 0xf1, 0x34,
	0x94, 0xf3, 0x92, 0x09, 0x9d, 0x20, 0x92, 0x86, 0x3c, 0xce, 0x51, 0x5f, 0x45, 0xef, 0x6a, 0xa9,
	0x6f, 0x43, 0xff, 0xff, 0xb6, 0xa6, 0x3e, 0x7e, 0x88, 0xf6, 0xd4, 0x4b, 0x4d, 0x18, 0x47, 0x82,
	0xf4, 0xbd, 0x8e, 0xdf, 0xb3, 0x26, 0x5c, 0x3f, 0xc2, 0xcf, 0xd0, 0xbe, 0x02, 0xde, 0xc2, 0x98,
	0x7f, 0x9a, 0x5f, 0x65, 0xdc, 0x6a, 0x64, 0xfc, 0x8d, 0x30, 0x7c, 0x79, 0xb1, 0x70, 0x9d, 0xcb,
	0x85, 0xeb, 0xfc, 0x5c, 0xb8, 0xce, 0x97, 0xa5, 0xdb, 0xba, 0x5c, 0xba, 0xad, 0x6f, 0x4b, 0xb7,
	0xf5, 0xf1, 0x51, 0xa3, 0x68, 0xd3, 0x9b, 0x63, 0xd5, 0x9c, 0xd0, 0x7e, 0xd3, 0xf3, 0xd5, 0x46,
	0x97, 0x9f, 0x6e, 0xeb, 0xff, 0xfa, 0xf4, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x33, 0x43, 0x92,
	0x38, 0x41, 0x04, 0x00, 0x00,
}

func (m *ItemRecord) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ItemRecord) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ItemRecord) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Strings) > 0 {
		for iNdEx := len(m.Strings) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Strings[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintExecution(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.Longs) > 0 {
		for iNdEx := len(m.Longs) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Longs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintExecution(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Doubles) > 0 {
		for iNdEx := len(m.Doubles) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Doubles[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintExecution(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.ID) > 0 {
		i -= len(m.ID)
		copy(dAtA[i:], m.ID)
		i = encodeVarintExecution(dAtA, i, uint64(len(m.ID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Execution) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Execution) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Execution) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ItemModifyOutputIDs) > 0 {
		for iNdEx := len(m.ItemModifyOutputIDs) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.ItemModifyOutputIDs[iNdEx])
			copy(dAtA[i:], m.ItemModifyOutputIDs[iNdEx])
			i = encodeVarintExecution(dAtA, i, uint64(len(m.ItemModifyOutputIDs[iNdEx])))
			i--
			dAtA[i] = 0x62
		}
	}
	if len(m.ItemOutputIDs) > 0 {
		for iNdEx := len(m.ItemOutputIDs) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.ItemOutputIDs[iNdEx])
			copy(dAtA[i:], m.ItemOutputIDs[iNdEx])
			i = encodeVarintExecution(dAtA, i, uint64(len(m.ItemOutputIDs[iNdEx])))
			i--
			dAtA[i] = 0x5a
		}
	}
	if len(m.CoinOutputs) > 0 {
		for iNdEx := len(m.CoinOutputs) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.CoinOutputs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintExecution(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x52
		}
	}
	if len(m.CoinInputs) > 0 {
		for iNdEx := len(m.CoinInputs) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.CoinInputs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintExecution(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x4a
		}
	}
	if len(m.ItemInputs) > 0 {
		for iNdEx := len(m.ItemInputs) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ItemInputs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintExecution(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x42
		}
	}
	if m.BlockHeight != 0 {
		i = encodeVarintExecution(dAtA, i, uint64(m.BlockHeight))
		i--
		dAtA[i] = 0x38
	}
	if m.NodeVersion != 0 {
		i = encodeVarintExecution(dAtA, i, uint64(m.NodeVersion))
		i--
		dAtA[i] = 0x30
	}
	if len(m.RecipeVersion) > 0 {
		i -= len(m.RecipeVersion)
		copy(dAtA[i:], m.RecipeVersion)
		i = encodeVarintExecution(dAtA, i, uint64(len(m.RecipeVersion)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.CookbookID) > 0 {
		i -= len(m.CookbookID)
		copy(dAtA[i:], m.CookbookID)
		i = encodeVarintExecution(dAtA, i, uint64(len(m.CookbookID)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.RecipeID) > 0 {
		i -= len(m.RecipeID)
		copy(dAtA[i:], m.RecipeID)
		i = encodeVarintExecution(dAtA, i, uint64(len(m.RecipeID)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.ID) > 0 {
		i -= len(m.ID)
		copy(dAtA[i:], m.ID)
		i = encodeVarintExecution(dAtA, i, uint64(len(m.ID)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintExecution(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintExecution(dAtA []byte, offset int, v uint64) int {
	offset -= sovExecution(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ItemRecord) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ID)
	if l > 0 {
		n += 1 + l + sovExecution(uint64(l))
	}
	if len(m.Doubles) > 0 {
		for _, e := range m.Doubles {
			l = e.Size()
			n += 1 + l + sovExecution(uint64(l))
		}
	}
	if len(m.Longs) > 0 {
		for _, e := range m.Longs {
			l = e.Size()
			n += 1 + l + sovExecution(uint64(l))
		}
	}
	if len(m.Strings) > 0 {
		for _, e := range m.Strings {
			l = e.Size()
			n += 1 + l + sovExecution(uint64(l))
		}
	}
	return n
}

func (m *Execution) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovExecution(uint64(l))
	}
	l = len(m.ID)
	if l > 0 {
		n += 1 + l + sovExecution(uint64(l))
	}
	l = len(m.RecipeID)
	if l > 0 {
		n += 1 + l + sovExecution(uint64(l))
	}
	l = len(m.CookbookID)
	if l > 0 {
		n += 1 + l + sovExecution(uint64(l))
	}
	l = len(m.RecipeVersion)
	if l > 0 {
		n += 1 + l + sovExecution(uint64(l))
	}
	if m.NodeVersion != 0 {
		n += 1 + sovExecution(uint64(m.NodeVersion))
	}
	if m.BlockHeight != 0 {
		n += 1 + sovExecution(uint64(m.BlockHeight))
	}
	if len(m.ItemInputs) > 0 {
		for _, e := range m.ItemInputs {
			l = e.Size()
			n += 1 + l + sovExecution(uint64(l))
		}
	}
	if len(m.CoinInputs) > 0 {
		for _, e := range m.CoinInputs {
			l = e.Size()
			n += 1 + l + sovExecution(uint64(l))
		}
	}
	if len(m.CoinOutputs) > 0 {
		for _, e := range m.CoinOutputs {
			l = e.Size()
			n += 1 + l + sovExecution(uint64(l))
		}
	}
	if len(m.ItemOutputIDs) > 0 {
		for _, s := range m.ItemOutputIDs {
			l = len(s)
			n += 1 + l + sovExecution(uint64(l))
		}
	}
	if len(m.ItemModifyOutputIDs) > 0 {
		for _, s := range m.ItemModifyOutputIDs {
			l = len(s)
			n += 1 + l + sovExecution(uint64(l))
		}
	}
	return n
}

func sovExecution(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozExecution(x uint64) (n int) {
	return sovExecution(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ItemRecord) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowExecution
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
			return fmt.Errorf("proto: ItemRecord: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ItemRecord: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExecution
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthExecution
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthExecution
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Doubles", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExecution
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthExecution
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthExecution
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Doubles = append(m.Doubles, DoubleKeyValue{})
			if err := m.Doubles[len(m.Doubles)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Longs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExecution
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthExecution
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthExecution
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Longs = append(m.Longs, LongKeyValue{})
			if err := m.Longs[len(m.Longs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Strings", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExecution
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthExecution
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthExecution
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Strings = append(m.Strings, StringKeyValue{})
			if err := m.Strings[len(m.Strings)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipExecution(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthExecution
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
func (m *Execution) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowExecution
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
			return fmt.Errorf("proto: Execution: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Execution: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExecution
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthExecution
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthExecution
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExecution
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthExecution
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthExecution
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RecipeID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExecution
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthExecution
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthExecution
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RecipeID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CookbookID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExecution
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthExecution
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthExecution
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CookbookID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RecipeVersion", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExecution
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthExecution
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthExecution
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RecipeVersion = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NodeVersion", wireType)
			}
			m.NodeVersion = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExecution
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NodeVersion |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockHeight", wireType)
			}
			m.BlockHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExecution
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BlockHeight |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ItemInputs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExecution
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthExecution
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthExecution
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ItemInputs = append(m.ItemInputs, ItemRecord{})
			if err := m.ItemInputs[len(m.ItemInputs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CoinInputs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExecution
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthExecution
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthExecution
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CoinInputs = append(m.CoinInputs, types.Coin{})
			if err := m.CoinInputs[len(m.CoinInputs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CoinOutputs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExecution
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthExecution
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthExecution
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CoinOutputs = append(m.CoinOutputs, types.Coin{})
			if err := m.CoinOutputs[len(m.CoinOutputs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ItemOutputIDs", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExecution
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthExecution
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthExecution
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ItemOutputIDs = append(m.ItemOutputIDs, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ItemModifyOutputIDs", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExecution
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthExecution
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthExecution
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ItemModifyOutputIDs = append(m.ItemModifyOutputIDs, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipExecution(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthExecution
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
func skipExecution(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowExecution
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
					return 0, ErrIntOverflowExecution
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
					return 0, ErrIntOverflowExecution
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
				return 0, ErrInvalidLengthExecution
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupExecution
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthExecution
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthExecution        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowExecution          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupExecution = fmt.Errorf("proto: unexpected end of group")
)
