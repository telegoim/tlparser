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

// SetChatDiscussionGroupRequest represents TL type `setChatDiscussionGroup#c93c32b8`.
type SetChatDiscussionGroupRequest struct {
	// Identifier of the channel chat. Pass 0 to remove a link from the supergroup passed in
	// the second argument to a linked channel chat (requires can_pin_messages rights in the
	// supergroup)
	ChatID int64
	// Identifier of a new channel's discussion group. Use 0 to remove the discussion group.
	DiscussionChatID int64
}

// SetChatDiscussionGroupRequestTypeID is TL type id of SetChatDiscussionGroupRequest.
const SetChatDiscussionGroupRequestTypeID = 0xc93c32b8

// Ensuring interfaces in compile-time for SetChatDiscussionGroupRequest.
var (
	_ bin.Encoder     = &SetChatDiscussionGroupRequest{}
	_ bin.Decoder     = &SetChatDiscussionGroupRequest{}
	_ bin.BareEncoder = &SetChatDiscussionGroupRequest{}
	_ bin.BareDecoder = &SetChatDiscussionGroupRequest{}
)

func (s *SetChatDiscussionGroupRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.ChatID == 0) {
		return false
	}
	if !(s.DiscussionChatID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *SetChatDiscussionGroupRequest) String() string {
	if s == nil {
		return "SetChatDiscussionGroupRequest(nil)"
	}
	type Alias SetChatDiscussionGroupRequest
	return fmt.Sprintf("SetChatDiscussionGroupRequest%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*SetChatDiscussionGroupRequest) TypeID() uint32 {
	return SetChatDiscussionGroupRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*SetChatDiscussionGroupRequest) TypeName() string {
	return "setChatDiscussionGroup"
}

// TypeInfo returns info about TL type.
func (s *SetChatDiscussionGroupRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "setChatDiscussionGroup",
		ID:   SetChatDiscussionGroupRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ChatID",
			SchemaName: "chat_id",
		},
		{
			Name:       "DiscussionChatID",
			SchemaName: "discussion_chat_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *SetChatDiscussionGroupRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode setChatDiscussionGroup#c93c32b8 as nil")
	}
	b.PutID(SetChatDiscussionGroupRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *SetChatDiscussionGroupRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode setChatDiscussionGroup#c93c32b8 as nil")
	}
	b.PutInt53(s.ChatID)
	b.PutInt53(s.DiscussionChatID)
	return nil
}

// Decode implements bin.Decoder.
func (s *SetChatDiscussionGroupRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode setChatDiscussionGroup#c93c32b8 to nil")
	}
	if err := b.ConsumeID(SetChatDiscussionGroupRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode setChatDiscussionGroup#c93c32b8: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *SetChatDiscussionGroupRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode setChatDiscussionGroup#c93c32b8 to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode setChatDiscussionGroup#c93c32b8: field chat_id: %w", err)
		}
		s.ChatID = value
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode setChatDiscussionGroup#c93c32b8: field discussion_chat_id: %w", err)
		}
		s.DiscussionChatID = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (s *SetChatDiscussionGroupRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if s == nil {
		return fmt.Errorf("can't encode setChatDiscussionGroup#c93c32b8 as nil")
	}
	b.ObjStart()
	b.PutID("setChatDiscussionGroup")
	b.FieldStart("chat_id")
	b.PutInt53(s.ChatID)
	b.FieldStart("discussion_chat_id")
	b.PutInt53(s.DiscussionChatID)
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (s *SetChatDiscussionGroupRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if s == nil {
		return fmt.Errorf("can't decode setChatDiscussionGroup#c93c32b8 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("setChatDiscussionGroup"); err != nil {
				return fmt.Errorf("unable to decode setChatDiscussionGroup#c93c32b8: %w", err)
			}
		case "chat_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode setChatDiscussionGroup#c93c32b8: field chat_id: %w", err)
			}
			s.ChatID = value
		case "discussion_chat_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode setChatDiscussionGroup#c93c32b8: field discussion_chat_id: %w", err)
			}
			s.DiscussionChatID = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetChatID returns value of ChatID field.
func (s *SetChatDiscussionGroupRequest) GetChatID() (value int64) {
	if s == nil {
		return
	}
	return s.ChatID
}

// GetDiscussionChatID returns value of DiscussionChatID field.
func (s *SetChatDiscussionGroupRequest) GetDiscussionChatID() (value int64) {
	if s == nil {
		return
	}
	return s.DiscussionChatID
}

// SetChatDiscussionGroup invokes method setChatDiscussionGroup#c93c32b8 returning error if any.
func (c *Client) SetChatDiscussionGroup(ctx context.Context, request *SetChatDiscussionGroupRequest) error {
	var ok Ok

	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
