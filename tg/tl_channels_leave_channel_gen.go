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

// ChannelsLeaveChannelRequest represents TL type `channels.leaveChannel#f836aa95`.
// Leave a channel/supergroup¹
//
// Links:
//  1) https://core.telegram.org/api/channel
//
// See https://core.telegram.org/method/channels.leaveChannel for reference.
type ChannelsLeaveChannelRequest struct {
	// Channel/supergroup¹ to leave
	//
	// Links:
	//  1) https://core.telegram.org/api/channel
	Channel InputChannelClass
}

// ChannelsLeaveChannelRequestTypeID is TL type id of ChannelsLeaveChannelRequest.
const ChannelsLeaveChannelRequestTypeID = 0xf836aa95

// Ensuring interfaces in compile-time for ChannelsLeaveChannelRequest.
var (
	_ bin.Encoder     = &ChannelsLeaveChannelRequest{}
	_ bin.Decoder     = &ChannelsLeaveChannelRequest{}
	_ bin.BareEncoder = &ChannelsLeaveChannelRequest{}
	_ bin.BareDecoder = &ChannelsLeaveChannelRequest{}
)

func (l *ChannelsLeaveChannelRequest) Zero() bool {
	if l == nil {
		return true
	}
	if !(l.Channel == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (l *ChannelsLeaveChannelRequest) String() string {
	if l == nil {
		return "ChannelsLeaveChannelRequest(nil)"
	}
	type Alias ChannelsLeaveChannelRequest
	return fmt.Sprintf("ChannelsLeaveChannelRequest%+v", Alias(*l))
}

// FillFrom fills ChannelsLeaveChannelRequest from given interface.
func (l *ChannelsLeaveChannelRequest) FillFrom(from interface {
	GetChannel() (value InputChannelClass)
}) {
	l.Channel = from.GetChannel()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChannelsLeaveChannelRequest) TypeID() uint32 {
	return ChannelsLeaveChannelRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*ChannelsLeaveChannelRequest) TypeName() string {
	return "channels.leaveChannel"
}

// TypeInfo returns info about TL type.
func (l *ChannelsLeaveChannelRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "channels.leaveChannel",
		ID:   ChannelsLeaveChannelRequestTypeID,
	}
	if l == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Channel",
			SchemaName: "channel",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (l *ChannelsLeaveChannelRequest) Encode(b *bin.Buffer) error {
	if l == nil {
		return fmt.Errorf("can't encode channels.leaveChannel#f836aa95 as nil")
	}
	b.PutID(ChannelsLeaveChannelRequestTypeID)
	return l.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (l *ChannelsLeaveChannelRequest) EncodeBare(b *bin.Buffer) error {
	if l == nil {
		return fmt.Errorf("can't encode channels.leaveChannel#f836aa95 as nil")
	}
	if l.Channel == nil {
		return fmt.Errorf("unable to encode channels.leaveChannel#f836aa95: field channel is nil")
	}
	if err := l.Channel.Encode(b); err != nil {
		return fmt.Errorf("unable to encode channels.leaveChannel#f836aa95: field channel: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (l *ChannelsLeaveChannelRequest) Decode(b *bin.Buffer) error {
	if l == nil {
		return fmt.Errorf("can't decode channels.leaveChannel#f836aa95 to nil")
	}
	if err := b.ConsumeID(ChannelsLeaveChannelRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode channels.leaveChannel#f836aa95: %w", err)
	}
	return l.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (l *ChannelsLeaveChannelRequest) DecodeBare(b *bin.Buffer) error {
	if l == nil {
		return fmt.Errorf("can't decode channels.leaveChannel#f836aa95 to nil")
	}
	{
		value, err := DecodeInputChannel(b)
		if err != nil {
			return fmt.Errorf("unable to decode channels.leaveChannel#f836aa95: field channel: %w", err)
		}
		l.Channel = value
	}
	return nil
}

// GetChannel returns value of Channel field.
func (l *ChannelsLeaveChannelRequest) GetChannel() (value InputChannelClass) {
	if l == nil {
		return
	}
	return l.Channel
}

// GetChannelAsNotEmpty returns mapped value of Channel field.
func (l *ChannelsLeaveChannelRequest) GetChannelAsNotEmpty() (NotEmptyInputChannel, bool) {
	return l.Channel.AsNotEmpty()
}

// ChannelsLeaveChannel invokes method channels.leaveChannel#f836aa95 returning error if any.
// Leave a channel/supergroup¹
//
// Links:
//  1) https://core.telegram.org/api/channel
//
// Possible errors:
//  400 CHANNEL_INVALID: The provided channel is invalid.
//  400 CHANNEL_PRIVATE: You haven't joined this channel/supergroup.
//  403 CHANNEL_PUBLIC_GROUP_NA: channel/supergroup not available.
//  400 MSG_ID_INVALID: Invalid message ID provided.
//  400 USER_CREATOR: You can't leave this channel, because you're its creator.
//  400 USER_NOT_PARTICIPANT: You're not a member of this supergroup/channel.
//
// See https://core.telegram.org/method/channels.leaveChannel for reference.
// Can be used by bots.
func (c *Client) ChannelsLeaveChannel(ctx context.Context, channel InputChannelClass) (UpdatesClass, error) {
	var result UpdatesBox

	request := &ChannelsLeaveChannelRequest{
		Channel: channel,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Updates, nil
}
