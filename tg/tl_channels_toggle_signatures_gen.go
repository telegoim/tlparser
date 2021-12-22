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

// ChannelsToggleSignaturesRequest represents TL type `channels.toggleSignatures#1f69b606`.
// Enable/disable message signatures in channels
//
// See https://core.telegram.org/method/channels.toggleSignatures for reference.
type ChannelsToggleSignaturesRequest struct {
	// Channel
	Channel InputChannelClass
	// Value
	Enabled bool
}

// ChannelsToggleSignaturesRequestTypeID is TL type id of ChannelsToggleSignaturesRequest.
const ChannelsToggleSignaturesRequestTypeID = 0x1f69b606

// Ensuring interfaces in compile-time for ChannelsToggleSignaturesRequest.
var (
	_ bin.Encoder     = &ChannelsToggleSignaturesRequest{}
	_ bin.Decoder     = &ChannelsToggleSignaturesRequest{}
	_ bin.BareEncoder = &ChannelsToggleSignaturesRequest{}
	_ bin.BareDecoder = &ChannelsToggleSignaturesRequest{}
)

func (t *ChannelsToggleSignaturesRequest) Zero() bool {
	if t == nil {
		return true
	}
	if !(t.Channel == nil) {
		return false
	}
	if !(t.Enabled == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (t *ChannelsToggleSignaturesRequest) String() string {
	if t == nil {
		return "ChannelsToggleSignaturesRequest(nil)"
	}
	type Alias ChannelsToggleSignaturesRequest
	return fmt.Sprintf("ChannelsToggleSignaturesRequest%+v", Alias(*t))
}

// FillFrom fills ChannelsToggleSignaturesRequest from given interface.
func (t *ChannelsToggleSignaturesRequest) FillFrom(from interface {
	GetChannel() (value InputChannelClass)
	GetEnabled() (value bool)
}) {
	t.Channel = from.GetChannel()
	t.Enabled = from.GetEnabled()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChannelsToggleSignaturesRequest) TypeID() uint32 {
	return ChannelsToggleSignaturesRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*ChannelsToggleSignaturesRequest) TypeName() string {
	return "channels.toggleSignatures"
}

// TypeInfo returns info about TL type.
func (t *ChannelsToggleSignaturesRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "channels.toggleSignatures",
		ID:   ChannelsToggleSignaturesRequestTypeID,
	}
	if t == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Channel",
			SchemaName: "channel",
		},
		{
			Name:       "Enabled",
			SchemaName: "enabled",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (t *ChannelsToggleSignaturesRequest) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode channels.toggleSignatures#1f69b606 as nil")
	}
	b.PutID(ChannelsToggleSignaturesRequestTypeID)
	return t.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (t *ChannelsToggleSignaturesRequest) EncodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode channels.toggleSignatures#1f69b606 as nil")
	}
	if t.Channel == nil {
		return fmt.Errorf("unable to encode channels.toggleSignatures#1f69b606: field channel is nil")
	}
	if err := t.Channel.Encode(b); err != nil {
		return fmt.Errorf("unable to encode channels.toggleSignatures#1f69b606: field channel: %w", err)
	}
	b.PutBool(t.Enabled)
	return nil
}

// Decode implements bin.Decoder.
func (t *ChannelsToggleSignaturesRequest) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode channels.toggleSignatures#1f69b606 to nil")
	}
	if err := b.ConsumeID(ChannelsToggleSignaturesRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode channels.toggleSignatures#1f69b606: %w", err)
	}
	return t.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (t *ChannelsToggleSignaturesRequest) DecodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode channels.toggleSignatures#1f69b606 to nil")
	}
	{
		value, err := DecodeInputChannel(b)
		if err != nil {
			return fmt.Errorf("unable to decode channels.toggleSignatures#1f69b606: field channel: %w", err)
		}
		t.Channel = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode channels.toggleSignatures#1f69b606: field enabled: %w", err)
		}
		t.Enabled = value
	}
	return nil
}

// GetChannel returns value of Channel field.
func (t *ChannelsToggleSignaturesRequest) GetChannel() (value InputChannelClass) {
	if t == nil {
		return
	}
	return t.Channel
}

// GetEnabled returns value of Enabled field.
func (t *ChannelsToggleSignaturesRequest) GetEnabled() (value bool) {
	if t == nil {
		return
	}
	return t.Enabled
}

// GetChannelAsNotEmpty returns mapped value of Channel field.
func (t *ChannelsToggleSignaturesRequest) GetChannelAsNotEmpty() (NotEmptyInputChannel, bool) {
	return t.Channel.AsNotEmpty()
}

// ChannelsToggleSignatures invokes method channels.toggleSignatures#1f69b606 returning error if any.
// Enable/disable message signatures in channels
//
// Possible errors:
//  400 CHANNEL_INVALID: The provided channel is invalid.
//  400 CHAT_ADMIN_REQUIRED: You must be an admin in this chat to do this.
//  400 CHAT_ID_INVALID: The provided chat id is invalid.
//
// See https://core.telegram.org/method/channels.toggleSignatures for reference.
func (c *Client) ChannelsToggleSignatures(ctx context.Context, request *ChannelsToggleSignaturesRequest) (UpdatesClass, error) {
	var result UpdatesBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Updates, nil
}
