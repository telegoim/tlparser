// Code generated by gotdgen, DO NOT EDIT.

package tg

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"

	"go.uber.org/multierr"

	"github.com/gotd/td/bin"
	"github.com/gotd/td/tdjson"
	"github.com/gotd/td/tdp"
	"github.com/gotd/td/tgerr"
)

// No-op definition for keeping imports.
var (
	_ = bin.Buffer{}
	_ = context.Background()
	_ = fmt.Stringer(nil)
	_ = strings.Builder{}
	_ = errors.Is
	_ = multierr.AppendInto
	_ = sort.Ints
	_ = tdp.Format
	_ = tgerr.Error{}
	_ = tdjson.Encoder{}
)

// InputChannelEmpty represents TL type `inputChannelEmpty#ee8c1e86`.
// Represents the absence of a channel
//
// See https://core.telegram.org/constructor/inputChannelEmpty for reference.
type InputChannelEmpty struct {
}

// InputChannelEmptyTypeID is TL type id of InputChannelEmpty.
const InputChannelEmptyTypeID = 0xee8c1e86

// construct implements constructor of InputChannelClass.
func (i InputChannelEmpty) construct() InputChannelClass { return &i }

// Ensuring interfaces in compile-time for InputChannelEmpty.
var (
	_ bin.Encoder     = &InputChannelEmpty{}
	_ bin.Decoder     = &InputChannelEmpty{}
	_ bin.BareEncoder = &InputChannelEmpty{}
	_ bin.BareDecoder = &InputChannelEmpty{}

	_ InputChannelClass = &InputChannelEmpty{}
)

func (i *InputChannelEmpty) Zero() bool {
	if i == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputChannelEmpty) String() string {
	if i == nil {
		return "InputChannelEmpty(nil)"
	}
	type Alias InputChannelEmpty
	return fmt.Sprintf("InputChannelEmpty%+v", Alias(*i))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InputChannelEmpty) TypeID() uint32 {
	return InputChannelEmptyTypeID
}

// TypeName returns name of type in TL schema.
func (*InputChannelEmpty) TypeName() string {
	return "inputChannelEmpty"
}

// TypeInfo returns info about TL type.
func (i *InputChannelEmpty) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inputChannelEmpty",
		ID:   InputChannelEmptyTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (i *InputChannelEmpty) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputChannelEmpty#ee8c1e86 as nil")
	}
	b.PutID(InputChannelEmptyTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InputChannelEmpty) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputChannelEmpty#ee8c1e86 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (i *InputChannelEmpty) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputChannelEmpty#ee8c1e86 to nil")
	}
	if err := b.ConsumeID(InputChannelEmptyTypeID); err != nil {
		return fmt.Errorf("unable to decode inputChannelEmpty#ee8c1e86: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InputChannelEmpty) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputChannelEmpty#ee8c1e86 to nil")
	}
	return nil
}

// InputChannel represents TL type `inputChannel#f35aec28`.
// Represents a channel
//
// See https://core.telegram.org/constructor/inputChannel for reference.
type InputChannel struct {
	// Channel ID
	ChannelID int64
	// Access hash taken from the channel¹ constructor
	//
	// Links:
	//  1) https://core.telegram.org/constructor/channel
	AccessHash int64
}

// InputChannelTypeID is TL type id of InputChannel.
const InputChannelTypeID = 0xf35aec28

// construct implements constructor of InputChannelClass.
func (i InputChannel) construct() InputChannelClass { return &i }

// Ensuring interfaces in compile-time for InputChannel.
var (
	_ bin.Encoder     = &InputChannel{}
	_ bin.Decoder     = &InputChannel{}
	_ bin.BareEncoder = &InputChannel{}
	_ bin.BareDecoder = &InputChannel{}

	_ InputChannelClass = &InputChannel{}
)

func (i *InputChannel) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.ChannelID == 0) {
		return false
	}
	if !(i.AccessHash == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputChannel) String() string {
	if i == nil {
		return "InputChannel(nil)"
	}
	type Alias InputChannel
	return fmt.Sprintf("InputChannel%+v", Alias(*i))
}

// FillFrom fills InputChannel from given interface.
func (i *InputChannel) FillFrom(from interface {
	GetChannelID() (value int64)
	GetAccessHash() (value int64)
}) {
	i.ChannelID = from.GetChannelID()
	i.AccessHash = from.GetAccessHash()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InputChannel) TypeID() uint32 {
	return InputChannelTypeID
}

// TypeName returns name of type in TL schema.
func (*InputChannel) TypeName() string {
	return "inputChannel"
}

// TypeInfo returns info about TL type.
func (i *InputChannel) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inputChannel",
		ID:   InputChannelTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ChannelID",
			SchemaName: "channel_id",
		},
		{
			Name:       "AccessHash",
			SchemaName: "access_hash",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (i *InputChannel) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputChannel#f35aec28 as nil")
	}
	b.PutID(InputChannelTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InputChannel) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputChannel#f35aec28 as nil")
	}
	b.PutLong(i.ChannelID)
	b.PutLong(i.AccessHash)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputChannel) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputChannel#f35aec28 to nil")
	}
	if err := b.ConsumeID(InputChannelTypeID); err != nil {
		return fmt.Errorf("unable to decode inputChannel#f35aec28: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InputChannel) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputChannel#f35aec28 to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode inputChannel#f35aec28: field channel_id: %w", err)
		}
		i.ChannelID = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode inputChannel#f35aec28: field access_hash: %w", err)
		}
		i.AccessHash = value
	}
	return nil
}

// GetChannelID returns value of ChannelID field.
func (i *InputChannel) GetChannelID() (value int64) {
	if i == nil {
		return
	}
	return i.ChannelID
}

// GetAccessHash returns value of AccessHash field.
func (i *InputChannel) GetAccessHash() (value int64) {
	if i == nil {
		return
	}
	return i.AccessHash
}

// InputChannelFromMessage represents TL type `inputChannelFromMessage#5b934f9d`.
// Defines a min¹ channel that was seen in a certain message of a certain chat.
//
// Links:
//  1) https://core.telegram.org/api/min
//
// See https://core.telegram.org/constructor/inputChannelFromMessage for reference.
type InputChannelFromMessage struct {
	// The chat where the channel was seen
	Peer InputPeerClass
	// The message ID in the chat where the channel was seen
	MsgID int
	// The channel ID
	ChannelID int64
}

// InputChannelFromMessageTypeID is TL type id of InputChannelFromMessage.
const InputChannelFromMessageTypeID = 0x5b934f9d

// construct implements constructor of InputChannelClass.
func (i InputChannelFromMessage) construct() InputChannelClass { return &i }

// Ensuring interfaces in compile-time for InputChannelFromMessage.
var (
	_ bin.Encoder     = &InputChannelFromMessage{}
	_ bin.Decoder     = &InputChannelFromMessage{}
	_ bin.BareEncoder = &InputChannelFromMessage{}
	_ bin.BareDecoder = &InputChannelFromMessage{}

	_ InputChannelClass = &InputChannelFromMessage{}
)

func (i *InputChannelFromMessage) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.Peer == nil) {
		return false
	}
	if !(i.MsgID == 0) {
		return false
	}
	if !(i.ChannelID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputChannelFromMessage) String() string {
	if i == nil {
		return "InputChannelFromMessage(nil)"
	}
	type Alias InputChannelFromMessage
	return fmt.Sprintf("InputChannelFromMessage%+v", Alias(*i))
}

// FillFrom fills InputChannelFromMessage from given interface.
func (i *InputChannelFromMessage) FillFrom(from interface {
	GetPeer() (value InputPeerClass)
	GetMsgID() (value int)
	GetChannelID() (value int64)
}) {
	i.Peer = from.GetPeer()
	i.MsgID = from.GetMsgID()
	i.ChannelID = from.GetChannelID()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InputChannelFromMessage) TypeID() uint32 {
	return InputChannelFromMessageTypeID
}

// TypeName returns name of type in TL schema.
func (*InputChannelFromMessage) TypeName() string {
	return "inputChannelFromMessage"
}

// TypeInfo returns info about TL type.
func (i *InputChannelFromMessage) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inputChannelFromMessage",
		ID:   InputChannelFromMessageTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Peer",
			SchemaName: "peer",
		},
		{
			Name:       "MsgID",
			SchemaName: "msg_id",
		},
		{
			Name:       "ChannelID",
			SchemaName: "channel_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (i *InputChannelFromMessage) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputChannelFromMessage#5b934f9d as nil")
	}
	b.PutID(InputChannelFromMessageTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InputChannelFromMessage) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputChannelFromMessage#5b934f9d as nil")
	}
	if i.Peer == nil {
		return fmt.Errorf("unable to encode inputChannelFromMessage#5b934f9d: field peer is nil")
	}
	if err := i.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode inputChannelFromMessage#5b934f9d: field peer: %w", err)
	}
	b.PutInt(i.MsgID)
	b.PutLong(i.ChannelID)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputChannelFromMessage) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputChannelFromMessage#5b934f9d to nil")
	}
	if err := b.ConsumeID(InputChannelFromMessageTypeID); err != nil {
		return fmt.Errorf("unable to decode inputChannelFromMessage#5b934f9d: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InputChannelFromMessage) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputChannelFromMessage#5b934f9d to nil")
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode inputChannelFromMessage#5b934f9d: field peer: %w", err)
		}
		i.Peer = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode inputChannelFromMessage#5b934f9d: field msg_id: %w", err)
		}
		i.MsgID = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode inputChannelFromMessage#5b934f9d: field channel_id: %w", err)
		}
		i.ChannelID = value
	}
	return nil
}

// GetPeer returns value of Peer field.
func (i *InputChannelFromMessage) GetPeer() (value InputPeerClass) {
	if i == nil {
		return
	}
	return i.Peer
}

// GetMsgID returns value of MsgID field.
func (i *InputChannelFromMessage) GetMsgID() (value int) {
	if i == nil {
		return
	}
	return i.MsgID
}

// GetChannelID returns value of ChannelID field.
func (i *InputChannelFromMessage) GetChannelID() (value int64) {
	if i == nil {
		return
	}
	return i.ChannelID
}

// InputChannelClassName is schema name of InputChannelClass.
const InputChannelClassName = "InputChannel"

// InputChannelClass represents InputChannel generic type.
//
// See https://core.telegram.org/type/InputChannel for reference.
//
// Example:
//  g, err := tg.DecodeInputChannel(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *tg.InputChannelEmpty: // inputChannelEmpty#ee8c1e86
//  case *tg.InputChannel: // inputChannel#f35aec28
//  case *tg.InputChannelFromMessage: // inputChannelFromMessage#5b934f9d
//  default: panic(v)
//  }
type InputChannelClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() InputChannelClass

	// TypeID returns type id in TL schema.
	//
	// See https://core.telegram.org/mtproto/TL-tl#remarks.
	TypeID() uint32
	// TypeName returns name of type in TL schema.
	TypeName() string
	// String implements fmt.Stringer.
	String() string
	// Zero returns true if current object has a zero value.
	Zero() bool

	// AsNotEmpty tries to map InputChannelClass to NotEmptyInputChannel.
	AsNotEmpty() (NotEmptyInputChannel, bool)
}

// NotEmptyInputChannel represents NotEmpty subset of InputChannelClass.
type NotEmptyInputChannel interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() InputChannelClass

	// TypeID returns type id in TL schema.
	//
	// See https://core.telegram.org/mtproto/TL-tl#remarks.
	TypeID() uint32
	// TypeName returns name of type in TL schema.
	TypeName() string
	// String implements fmt.Stringer.
	String() string
	// Zero returns true if current object has a zero value.
	Zero() bool

	// Channel ID
	GetChannelID() (value int64)
}

// AsNotEmpty tries to map InputChannelEmpty to NotEmptyInputChannel.
func (i *InputChannelEmpty) AsNotEmpty() (NotEmptyInputChannel, bool) {
	value, ok := (InputChannelClass(i)).(NotEmptyInputChannel)
	return value, ok
}

// AsNotEmpty tries to map InputChannel to NotEmptyInputChannel.
func (i *InputChannel) AsNotEmpty() (NotEmptyInputChannel, bool) {
	value, ok := (InputChannelClass(i)).(NotEmptyInputChannel)
	return value, ok
}

// AsNotEmpty tries to map InputChannelFromMessage to NotEmptyInputChannel.
func (i *InputChannelFromMessage) AsNotEmpty() (NotEmptyInputChannel, bool) {
	value, ok := (InputChannelClass(i)).(NotEmptyInputChannel)
	return value, ok
}

// DecodeInputChannel implements binary de-serialization for InputChannelClass.
func DecodeInputChannel(buf *bin.Buffer) (InputChannelClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case InputChannelEmptyTypeID:
		// Decoding inputChannelEmpty#ee8c1e86.
		v := InputChannelEmpty{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputChannelClass: %w", err)
		}
		return &v, nil
	case InputChannelTypeID:
		// Decoding inputChannel#f35aec28.
		v := InputChannel{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputChannelClass: %w", err)
		}
		return &v, nil
	case InputChannelFromMessageTypeID:
		// Decoding inputChannelFromMessage#5b934f9d.
		v := InputChannelFromMessage{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputChannelClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode InputChannelClass: %w", bin.NewUnexpectedID(id))
	}
}

// InputChannel boxes the InputChannelClass providing a helper.
type InputChannelBox struct {
	InputChannel InputChannelClass
}

// Decode implements bin.Decoder for InputChannelBox.
func (b *InputChannelBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode InputChannelBox to nil")
	}
	v, err := DecodeInputChannel(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.InputChannel = v
	return nil
}

// Encode implements bin.Encode for InputChannelBox.
func (b *InputChannelBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.InputChannel == nil {
		return fmt.Errorf("unable to encode InputChannelClass as nil")
	}
	return b.InputChannel.Encode(buf)
}
