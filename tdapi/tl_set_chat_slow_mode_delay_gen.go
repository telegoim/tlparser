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

// SetChatSlowModeDelayRequest represents TL type `setChatSlowModeDelay#dfcae63e`.
type SetChatSlowModeDelayRequest struct {
	// Chat identifier
	ChatID int64
	// New slow mode delay for the chat, in seconds; must be one of 0, 10, 30, 60, 300, 900,
	// 3600
	SlowModeDelay int32
}

// SetChatSlowModeDelayRequestTypeID is TL type id of SetChatSlowModeDelayRequest.
const SetChatSlowModeDelayRequestTypeID = 0xdfcae63e

// Ensuring interfaces in compile-time for SetChatSlowModeDelayRequest.
var (
	_ bin.Encoder     = &SetChatSlowModeDelayRequest{}
	_ bin.Decoder     = &SetChatSlowModeDelayRequest{}
	_ bin.BareEncoder = &SetChatSlowModeDelayRequest{}
	_ bin.BareDecoder = &SetChatSlowModeDelayRequest{}
)

func (s *SetChatSlowModeDelayRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.ChatID == 0) {
		return false
	}
	if !(s.SlowModeDelay == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *SetChatSlowModeDelayRequest) String() string {
	if s == nil {
		return "SetChatSlowModeDelayRequest(nil)"
	}
	type Alias SetChatSlowModeDelayRequest
	return fmt.Sprintf("SetChatSlowModeDelayRequest%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*SetChatSlowModeDelayRequest) TypeID() uint32 {
	return SetChatSlowModeDelayRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*SetChatSlowModeDelayRequest) TypeName() string {
	return "setChatSlowModeDelay"
}

// TypeInfo returns info about TL type.
func (s *SetChatSlowModeDelayRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "setChatSlowModeDelay",
		ID:   SetChatSlowModeDelayRequestTypeID,
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
			Name:       "SlowModeDelay",
			SchemaName: "slow_mode_delay",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *SetChatSlowModeDelayRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode setChatSlowModeDelay#dfcae63e as nil")
	}
	b.PutID(SetChatSlowModeDelayRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *SetChatSlowModeDelayRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode setChatSlowModeDelay#dfcae63e as nil")
	}
	b.PutInt53(s.ChatID)
	b.PutInt32(s.SlowModeDelay)
	return nil
}

// Decode implements bin.Decoder.
func (s *SetChatSlowModeDelayRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode setChatSlowModeDelay#dfcae63e to nil")
	}
	if err := b.ConsumeID(SetChatSlowModeDelayRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode setChatSlowModeDelay#dfcae63e: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *SetChatSlowModeDelayRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode setChatSlowModeDelay#dfcae63e to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode setChatSlowModeDelay#dfcae63e: field chat_id: %w", err)
		}
		s.ChatID = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode setChatSlowModeDelay#dfcae63e: field slow_mode_delay: %w", err)
		}
		s.SlowModeDelay = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (s *SetChatSlowModeDelayRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if s == nil {
		return fmt.Errorf("can't encode setChatSlowModeDelay#dfcae63e as nil")
	}
	b.ObjStart()
	b.PutID("setChatSlowModeDelay")
	b.FieldStart("chat_id")
	b.PutInt53(s.ChatID)
	b.FieldStart("slow_mode_delay")
	b.PutInt32(s.SlowModeDelay)
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (s *SetChatSlowModeDelayRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if s == nil {
		return fmt.Errorf("can't decode setChatSlowModeDelay#dfcae63e to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("setChatSlowModeDelay"); err != nil {
				return fmt.Errorf("unable to decode setChatSlowModeDelay#dfcae63e: %w", err)
			}
		case "chat_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode setChatSlowModeDelay#dfcae63e: field chat_id: %w", err)
			}
			s.ChatID = value
		case "slow_mode_delay":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode setChatSlowModeDelay#dfcae63e: field slow_mode_delay: %w", err)
			}
			s.SlowModeDelay = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetChatID returns value of ChatID field.
func (s *SetChatSlowModeDelayRequest) GetChatID() (value int64) {
	if s == nil {
		return
	}
	return s.ChatID
}

// GetSlowModeDelay returns value of SlowModeDelay field.
func (s *SetChatSlowModeDelayRequest) GetSlowModeDelay() (value int32) {
	if s == nil {
		return
	}
	return s.SlowModeDelay
}

// SetChatSlowModeDelay invokes method setChatSlowModeDelay#dfcae63e returning error if any.
func (c *Client) SetChatSlowModeDelay(ctx context.Context, request *SetChatSlowModeDelayRequest) error {
	var ok Ok

	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
