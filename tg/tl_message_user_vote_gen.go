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

// MessageUserVote represents TL type `messageUserVote#34d247b4`.
// How a user voted in a poll
//
// See https://core.telegram.org/constructor/messageUserVote for reference.
type MessageUserVote struct {
	// User ID
	UserID int64
	// The option chosen by the user
	Option []byte
	// When did the user cast the vote
	Date int
}

// MessageUserVoteTypeID is TL type id of MessageUserVote.
const MessageUserVoteTypeID = 0x34d247b4

// construct implements constructor of MessageUserVoteClass.
func (m MessageUserVote) construct() MessageUserVoteClass { return &m }

// Ensuring interfaces in compile-time for MessageUserVote.
var (
	_ bin.Encoder     = &MessageUserVote{}
	_ bin.Decoder     = &MessageUserVote{}
	_ bin.BareEncoder = &MessageUserVote{}
	_ bin.BareDecoder = &MessageUserVote{}

	_ MessageUserVoteClass = &MessageUserVote{}
)

func (m *MessageUserVote) Zero() bool {
	if m == nil {
		return true
	}
	if !(m.UserID == 0) {
		return false
	}
	if !(m.Option == nil) {
		return false
	}
	if !(m.Date == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (m *MessageUserVote) String() string {
	if m == nil {
		return "MessageUserVote(nil)"
	}
	type Alias MessageUserVote
	return fmt.Sprintf("MessageUserVote%+v", Alias(*m))
}

// FillFrom fills MessageUserVote from given interface.
func (m *MessageUserVote) FillFrom(from interface {
	GetUserID() (value int64)
	GetOption() (value []byte)
	GetDate() (value int)
}) {
	m.UserID = from.GetUserID()
	m.Option = from.GetOption()
	m.Date = from.GetDate()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessageUserVote) TypeID() uint32 {
	return MessageUserVoteTypeID
}

// TypeName returns name of type in TL schema.
func (*MessageUserVote) TypeName() string {
	return "messageUserVote"
}

// TypeInfo returns info about TL type.
func (m *MessageUserVote) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messageUserVote",
		ID:   MessageUserVoteTypeID,
	}
	if m == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "UserID",
			SchemaName: "user_id",
		},
		{
			Name:       "Option",
			SchemaName: "option",
		},
		{
			Name:       "Date",
			SchemaName: "date",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (m *MessageUserVote) Encode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode messageUserVote#34d247b4 as nil")
	}
	b.PutID(MessageUserVoteTypeID)
	return m.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (m *MessageUserVote) EncodeBare(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode messageUserVote#34d247b4 as nil")
	}
	b.PutLong(m.UserID)
	b.PutBytes(m.Option)
	b.PutInt(m.Date)
	return nil
}

// Decode implements bin.Decoder.
func (m *MessageUserVote) Decode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode messageUserVote#34d247b4 to nil")
	}
	if err := b.ConsumeID(MessageUserVoteTypeID); err != nil {
		return fmt.Errorf("unable to decode messageUserVote#34d247b4: %w", err)
	}
	return m.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (m *MessageUserVote) DecodeBare(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode messageUserVote#34d247b4 to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode messageUserVote#34d247b4: field user_id: %w", err)
		}
		m.UserID = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode messageUserVote#34d247b4: field option: %w", err)
		}
		m.Option = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messageUserVote#34d247b4: field date: %w", err)
		}
		m.Date = value
	}
	return nil
}

// GetUserID returns value of UserID field.
func (m *MessageUserVote) GetUserID() (value int64) {
	if m == nil {
		return
	}
	return m.UserID
}

// GetOption returns value of Option field.
func (m *MessageUserVote) GetOption() (value []byte) {
	if m == nil {
		return
	}
	return m.Option
}

// GetDate returns value of Date field.
func (m *MessageUserVote) GetDate() (value int) {
	if m == nil {
		return
	}
	return m.Date
}

// MessageUserVoteInputOption represents TL type `messageUserVoteInputOption#3ca5b0ec`.
// How a user voted in a poll (reduced constructor, returned if an option was provided to
// messages.getPollVotes¹)
//
// Links:
//  1) https://core.telegram.org/method/messages.getPollVotes
//
// See https://core.telegram.org/constructor/messageUserVoteInputOption for reference.
type MessageUserVoteInputOption struct {
	// The user that voted for the queried option
	UserID int64
	// When did the user cast the vote
	Date int
}

// MessageUserVoteInputOptionTypeID is TL type id of MessageUserVoteInputOption.
const MessageUserVoteInputOptionTypeID = 0x3ca5b0ec

// construct implements constructor of MessageUserVoteClass.
func (m MessageUserVoteInputOption) construct() MessageUserVoteClass { return &m }

// Ensuring interfaces in compile-time for MessageUserVoteInputOption.
var (
	_ bin.Encoder     = &MessageUserVoteInputOption{}
	_ bin.Decoder     = &MessageUserVoteInputOption{}
	_ bin.BareEncoder = &MessageUserVoteInputOption{}
	_ bin.BareDecoder = &MessageUserVoteInputOption{}

	_ MessageUserVoteClass = &MessageUserVoteInputOption{}
)

func (m *MessageUserVoteInputOption) Zero() bool {
	if m == nil {
		return true
	}
	if !(m.UserID == 0) {
		return false
	}
	if !(m.Date == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (m *MessageUserVoteInputOption) String() string {
	if m == nil {
		return "MessageUserVoteInputOption(nil)"
	}
	type Alias MessageUserVoteInputOption
	return fmt.Sprintf("MessageUserVoteInputOption%+v", Alias(*m))
}

// FillFrom fills MessageUserVoteInputOption from given interface.
func (m *MessageUserVoteInputOption) FillFrom(from interface {
	GetUserID() (value int64)
	GetDate() (value int)
}) {
	m.UserID = from.GetUserID()
	m.Date = from.GetDate()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessageUserVoteInputOption) TypeID() uint32 {
	return MessageUserVoteInputOptionTypeID
}

// TypeName returns name of type in TL schema.
func (*MessageUserVoteInputOption) TypeName() string {
	return "messageUserVoteInputOption"
}

// TypeInfo returns info about TL type.
func (m *MessageUserVoteInputOption) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messageUserVoteInputOption",
		ID:   MessageUserVoteInputOptionTypeID,
	}
	if m == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "UserID",
			SchemaName: "user_id",
		},
		{
			Name:       "Date",
			SchemaName: "date",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (m *MessageUserVoteInputOption) Encode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode messageUserVoteInputOption#3ca5b0ec as nil")
	}
	b.PutID(MessageUserVoteInputOptionTypeID)
	return m.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (m *MessageUserVoteInputOption) EncodeBare(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode messageUserVoteInputOption#3ca5b0ec as nil")
	}
	b.PutLong(m.UserID)
	b.PutInt(m.Date)
	return nil
}

// Decode implements bin.Decoder.
func (m *MessageUserVoteInputOption) Decode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode messageUserVoteInputOption#3ca5b0ec to nil")
	}
	if err := b.ConsumeID(MessageUserVoteInputOptionTypeID); err != nil {
		return fmt.Errorf("unable to decode messageUserVoteInputOption#3ca5b0ec: %w", err)
	}
	return m.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (m *MessageUserVoteInputOption) DecodeBare(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode messageUserVoteInputOption#3ca5b0ec to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode messageUserVoteInputOption#3ca5b0ec: field user_id: %w", err)
		}
		m.UserID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messageUserVoteInputOption#3ca5b0ec: field date: %w", err)
		}
		m.Date = value
	}
	return nil
}

// GetUserID returns value of UserID field.
func (m *MessageUserVoteInputOption) GetUserID() (value int64) {
	if m == nil {
		return
	}
	return m.UserID
}

// GetDate returns value of Date field.
func (m *MessageUserVoteInputOption) GetDate() (value int) {
	if m == nil {
		return
	}
	return m.Date
}

// MessageUserVoteMultiple represents TL type `messageUserVoteMultiple#8a65e557`.
// How a user voted in a multiple-choice poll
//
// See https://core.telegram.org/constructor/messageUserVoteMultiple for reference.
type MessageUserVoteMultiple struct {
	// User ID
	UserID int64
	// Options chosen by the user
	Options [][]byte
	// When did the user cast their votes
	Date int
}

// MessageUserVoteMultipleTypeID is TL type id of MessageUserVoteMultiple.
const MessageUserVoteMultipleTypeID = 0x8a65e557

// construct implements constructor of MessageUserVoteClass.
func (m MessageUserVoteMultiple) construct() MessageUserVoteClass { return &m }

// Ensuring interfaces in compile-time for MessageUserVoteMultiple.
var (
	_ bin.Encoder     = &MessageUserVoteMultiple{}
	_ bin.Decoder     = &MessageUserVoteMultiple{}
	_ bin.BareEncoder = &MessageUserVoteMultiple{}
	_ bin.BareDecoder = &MessageUserVoteMultiple{}

	_ MessageUserVoteClass = &MessageUserVoteMultiple{}
)

func (m *MessageUserVoteMultiple) Zero() bool {
	if m == nil {
		return true
	}
	if !(m.UserID == 0) {
		return false
	}
	if !(m.Options == nil) {
		return false
	}
	if !(m.Date == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (m *MessageUserVoteMultiple) String() string {
	if m == nil {
		return "MessageUserVoteMultiple(nil)"
	}
	type Alias MessageUserVoteMultiple
	return fmt.Sprintf("MessageUserVoteMultiple%+v", Alias(*m))
}

// FillFrom fills MessageUserVoteMultiple from given interface.
func (m *MessageUserVoteMultiple) FillFrom(from interface {
	GetUserID() (value int64)
	GetOptions() (value [][]byte)
	GetDate() (value int)
}) {
	m.UserID = from.GetUserID()
	m.Options = from.GetOptions()
	m.Date = from.GetDate()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessageUserVoteMultiple) TypeID() uint32 {
	return MessageUserVoteMultipleTypeID
}

// TypeName returns name of type in TL schema.
func (*MessageUserVoteMultiple) TypeName() string {
	return "messageUserVoteMultiple"
}

// TypeInfo returns info about TL type.
func (m *MessageUserVoteMultiple) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messageUserVoteMultiple",
		ID:   MessageUserVoteMultipleTypeID,
	}
	if m == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "UserID",
			SchemaName: "user_id",
		},
		{
			Name:       "Options",
			SchemaName: "options",
		},
		{
			Name:       "Date",
			SchemaName: "date",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (m *MessageUserVoteMultiple) Encode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode messageUserVoteMultiple#8a65e557 as nil")
	}
	b.PutID(MessageUserVoteMultipleTypeID)
	return m.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (m *MessageUserVoteMultiple) EncodeBare(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode messageUserVoteMultiple#8a65e557 as nil")
	}
	b.PutLong(m.UserID)
	b.PutVectorHeader(len(m.Options))
	for _, v := range m.Options {
		b.PutBytes(v)
	}
	b.PutInt(m.Date)
	return nil
}

// Decode implements bin.Decoder.
func (m *MessageUserVoteMultiple) Decode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode messageUserVoteMultiple#8a65e557 to nil")
	}
	if err := b.ConsumeID(MessageUserVoteMultipleTypeID); err != nil {
		return fmt.Errorf("unable to decode messageUserVoteMultiple#8a65e557: %w", err)
	}
	return m.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (m *MessageUserVoteMultiple) DecodeBare(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode messageUserVoteMultiple#8a65e557 to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode messageUserVoteMultiple#8a65e557: field user_id: %w", err)
		}
		m.UserID = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messageUserVoteMultiple#8a65e557: field options: %w", err)
		}

		if headerLen > 0 {
			m.Options = make([][]byte, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Bytes()
			if err != nil {
				return fmt.Errorf("unable to decode messageUserVoteMultiple#8a65e557: field options: %w", err)
			}
			m.Options = append(m.Options, value)
		}
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messageUserVoteMultiple#8a65e557: field date: %w", err)
		}
		m.Date = value
	}
	return nil
}

// GetUserID returns value of UserID field.
func (m *MessageUserVoteMultiple) GetUserID() (value int64) {
	if m == nil {
		return
	}
	return m.UserID
}

// GetOptions returns value of Options field.
func (m *MessageUserVoteMultiple) GetOptions() (value [][]byte) {
	if m == nil {
		return
	}
	return m.Options
}

// GetDate returns value of Date field.
func (m *MessageUserVoteMultiple) GetDate() (value int) {
	if m == nil {
		return
	}
	return m.Date
}

// MessageUserVoteClassName is schema name of MessageUserVoteClass.
const MessageUserVoteClassName = "MessageUserVote"

// MessageUserVoteClass represents MessageUserVote generic type.
//
// See https://core.telegram.org/type/MessageUserVote for reference.
//
// Example:
//  g, err := tg.DecodeMessageUserVote(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *tg.MessageUserVote: // messageUserVote#34d247b4
//  case *tg.MessageUserVoteInputOption: // messageUserVoteInputOption#3ca5b0ec
//  case *tg.MessageUserVoteMultiple: // messageUserVoteMultiple#8a65e557
//  default: panic(v)
//  }
type MessageUserVoteClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() MessageUserVoteClass

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

	// User ID
	GetUserID() (value int64)

	// When did the user cast the vote
	GetDate() (value int)
}

// DecodeMessageUserVote implements binary de-serialization for MessageUserVoteClass.
func DecodeMessageUserVote(buf *bin.Buffer) (MessageUserVoteClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case MessageUserVoteTypeID:
		// Decoding messageUserVote#34d247b4.
		v := MessageUserVote{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode MessageUserVoteClass: %w", err)
		}
		return &v, nil
	case MessageUserVoteInputOptionTypeID:
		// Decoding messageUserVoteInputOption#3ca5b0ec.
		v := MessageUserVoteInputOption{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode MessageUserVoteClass: %w", err)
		}
		return &v, nil
	case MessageUserVoteMultipleTypeID:
		// Decoding messageUserVoteMultiple#8a65e557.
		v := MessageUserVoteMultiple{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode MessageUserVoteClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode MessageUserVoteClass: %w", bin.NewUnexpectedID(id))
	}
}

// MessageUserVote boxes the MessageUserVoteClass providing a helper.
type MessageUserVoteBox struct {
	MessageUserVote MessageUserVoteClass
}

// Decode implements bin.Decoder for MessageUserVoteBox.
func (b *MessageUserVoteBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode MessageUserVoteBox to nil")
	}
	v, err := DecodeMessageUserVote(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.MessageUserVote = v
	return nil
}

// Encode implements bin.Encode for MessageUserVoteBox.
func (b *MessageUserVoteBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.MessageUserVote == nil {
		return fmt.Errorf("unable to encode MessageUserVoteClass as nil")
	}
	return b.MessageUserVote.Encode(buf)
}
