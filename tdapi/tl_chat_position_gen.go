// Code generated by gotdgen, DO NOT EDIT.

package tdapi

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

// ChatPosition represents TL type `chatPosition#dae48755`.
type ChatPosition struct {
	// The chat list
	List ChatListClass
	// A parameter used to determine order of the chat in the chat list. Chats must be sorted
	// by the pair (order, chat.id) in descending order
	Order int64
	// True, if the chat is pinned in the chat list
	IsPinned bool
	// Source of the chat in the chat list; may be null
	Source ChatSourceClass
}

// ChatPositionTypeID is TL type id of ChatPosition.
const ChatPositionTypeID = 0xdae48755

// Ensuring interfaces in compile-time for ChatPosition.
var (
	_ bin.Encoder     = &ChatPosition{}
	_ bin.Decoder     = &ChatPosition{}
	_ bin.BareEncoder = &ChatPosition{}
	_ bin.BareDecoder = &ChatPosition{}
)

func (c *ChatPosition) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.List == nil) {
		return false
	}
	if !(c.Order == 0) {
		return false
	}
	if !(c.IsPinned == false) {
		return false
	}
	if !(c.Source == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *ChatPosition) String() string {
	if c == nil {
		return "ChatPosition(nil)"
	}
	type Alias ChatPosition
	return fmt.Sprintf("ChatPosition%+v", Alias(*c))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChatPosition) TypeID() uint32 {
	return ChatPositionTypeID
}

// TypeName returns name of type in TL schema.
func (*ChatPosition) TypeName() string {
	return "chatPosition"
}

// TypeInfo returns info about TL type.
func (c *ChatPosition) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "chatPosition",
		ID:   ChatPositionTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "List",
			SchemaName: "list",
		},
		{
			Name:       "Order",
			SchemaName: "order",
		},
		{
			Name:       "IsPinned",
			SchemaName: "is_pinned",
		},
		{
			Name:       "Source",
			SchemaName: "source",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *ChatPosition) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatPosition#dae48755 as nil")
	}
	b.PutID(ChatPositionTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *ChatPosition) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatPosition#dae48755 as nil")
	}
	if c.List == nil {
		return fmt.Errorf("unable to encode chatPosition#dae48755: field list is nil")
	}
	if err := c.List.Encode(b); err != nil {
		return fmt.Errorf("unable to encode chatPosition#dae48755: field list: %w", err)
	}
	b.PutLong(c.Order)
	b.PutBool(c.IsPinned)
	if c.Source == nil {
		return fmt.Errorf("unable to encode chatPosition#dae48755: field source is nil")
	}
	if err := c.Source.Encode(b); err != nil {
		return fmt.Errorf("unable to encode chatPosition#dae48755: field source: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (c *ChatPosition) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatPosition#dae48755 to nil")
	}
	if err := b.ConsumeID(ChatPositionTypeID); err != nil {
		return fmt.Errorf("unable to decode chatPosition#dae48755: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *ChatPosition) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatPosition#dae48755 to nil")
	}
	{
		value, err := DecodeChatList(b)
		if err != nil {
			return fmt.Errorf("unable to decode chatPosition#dae48755: field list: %w", err)
		}
		c.List = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode chatPosition#dae48755: field order: %w", err)
		}
		c.Order = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode chatPosition#dae48755: field is_pinned: %w", err)
		}
		c.IsPinned = value
	}
	{
		value, err := DecodeChatSource(b)
		if err != nil {
			return fmt.Errorf("unable to decode chatPosition#dae48755: field source: %w", err)
		}
		c.Source = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (c *ChatPosition) EncodeTDLibJSON(b tdjson.Encoder) error {
	if c == nil {
		return fmt.Errorf("can't encode chatPosition#dae48755 as nil")
	}
	b.ObjStart()
	b.PutID("chatPosition")
	b.FieldStart("list")
	if c.List == nil {
		return fmt.Errorf("unable to encode chatPosition#dae48755: field list is nil")
	}
	if err := c.List.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode chatPosition#dae48755: field list: %w", err)
	}
	b.FieldStart("order")
	b.PutLong(c.Order)
	b.FieldStart("is_pinned")
	b.PutBool(c.IsPinned)
	b.FieldStart("source")
	if c.Source == nil {
		return fmt.Errorf("unable to encode chatPosition#dae48755: field source is nil")
	}
	if err := c.Source.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode chatPosition#dae48755: field source: %w", err)
	}
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (c *ChatPosition) DecodeTDLibJSON(b tdjson.Decoder) error {
	if c == nil {
		return fmt.Errorf("can't decode chatPosition#dae48755 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("chatPosition"); err != nil {
				return fmt.Errorf("unable to decode chatPosition#dae48755: %w", err)
			}
		case "list":
			value, err := DecodeTDLibJSONChatList(b)
			if err != nil {
				return fmt.Errorf("unable to decode chatPosition#dae48755: field list: %w", err)
			}
			c.List = value
		case "order":
			value, err := b.Long()
			if err != nil {
				return fmt.Errorf("unable to decode chatPosition#dae48755: field order: %w", err)
			}
			c.Order = value
		case "is_pinned":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode chatPosition#dae48755: field is_pinned: %w", err)
			}
			c.IsPinned = value
		case "source":
			value, err := DecodeTDLibJSONChatSource(b)
			if err != nil {
				return fmt.Errorf("unable to decode chatPosition#dae48755: field source: %w", err)
			}
			c.Source = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetList returns value of List field.
func (c *ChatPosition) GetList() (value ChatListClass) {
	if c == nil {
		return
	}
	return c.List
}

// GetOrder returns value of Order field.
func (c *ChatPosition) GetOrder() (value int64) {
	if c == nil {
		return
	}
	return c.Order
}

// GetIsPinned returns value of IsPinned field.
func (c *ChatPosition) GetIsPinned() (value bool) {
	if c == nil {
		return
	}
	return c.IsPinned
}

// GetSource returns value of Source field.
func (c *ChatPosition) GetSource() (value ChatSourceClass) {
	if c == nil {
		return
	}
	return c.Source
}
