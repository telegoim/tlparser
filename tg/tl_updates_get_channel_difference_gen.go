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

// UpdatesGetChannelDifferenceRequest represents TL type `updates.getChannelDifference#3173d78`.
// Returns the difference between the current state of updates of a certain channel and
// transmitted.
//
// See https://core.telegram.org/method/updates.getChannelDifference for reference.
type UpdatesGetChannelDifferenceRequest struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Set to true to skip some possibly unneeded updates and reduce server-side load
	Force bool
	// The channel
	Channel InputChannelClass
	// Messsage filter
	Filter ChannelMessagesFilterClass
	// Persistent timestamp (see updates¹)
	//
	// Links:
	//  1) https://core.telegram.org/api/updates
	Pts int
	// How many updates to fetch, max 100000Ordinary (non-bot) users are supposed to pass
	// 10-100
	Limit int
}

// UpdatesGetChannelDifferenceRequestTypeID is TL type id of UpdatesGetChannelDifferenceRequest.
const UpdatesGetChannelDifferenceRequestTypeID = 0x3173d78

// Ensuring interfaces in compile-time for UpdatesGetChannelDifferenceRequest.
var (
	_ bin.Encoder     = &UpdatesGetChannelDifferenceRequest{}
	_ bin.Decoder     = &UpdatesGetChannelDifferenceRequest{}
	_ bin.BareEncoder = &UpdatesGetChannelDifferenceRequest{}
	_ bin.BareDecoder = &UpdatesGetChannelDifferenceRequest{}
)

func (g *UpdatesGetChannelDifferenceRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Flags.Zero()) {
		return false
	}
	if !(g.Force == false) {
		return false
	}
	if !(g.Channel == nil) {
		return false
	}
	if !(g.Filter == nil) {
		return false
	}
	if !(g.Pts == 0) {
		return false
	}
	if !(g.Limit == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *UpdatesGetChannelDifferenceRequest) String() string {
	if g == nil {
		return "UpdatesGetChannelDifferenceRequest(nil)"
	}
	type Alias UpdatesGetChannelDifferenceRequest
	return fmt.Sprintf("UpdatesGetChannelDifferenceRequest%+v", Alias(*g))
}

// FillFrom fills UpdatesGetChannelDifferenceRequest from given interface.
func (g *UpdatesGetChannelDifferenceRequest) FillFrom(from interface {
	GetForce() (value bool)
	GetChannel() (value InputChannelClass)
	GetFilter() (value ChannelMessagesFilterClass)
	GetPts() (value int)
	GetLimit() (value int)
}) {
	g.Force = from.GetForce()
	g.Channel = from.GetChannel()
	g.Filter = from.GetFilter()
	g.Pts = from.GetPts()
	g.Limit = from.GetLimit()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*UpdatesGetChannelDifferenceRequest) TypeID() uint32 {
	return UpdatesGetChannelDifferenceRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*UpdatesGetChannelDifferenceRequest) TypeName() string {
	return "updates.getChannelDifference"
}

// TypeInfo returns info about TL type.
func (g *UpdatesGetChannelDifferenceRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "updates.getChannelDifference",
		ID:   UpdatesGetChannelDifferenceRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Force",
			SchemaName: "force",
			Null:       !g.Flags.Has(0),
		},
		{
			Name:       "Channel",
			SchemaName: "channel",
		},
		{
			Name:       "Filter",
			SchemaName: "filter",
		},
		{
			Name:       "Pts",
			SchemaName: "pts",
		},
		{
			Name:       "Limit",
			SchemaName: "limit",
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (g *UpdatesGetChannelDifferenceRequest) SetFlags() {
	if !(g.Force == false) {
		g.Flags.Set(0)
	}
}

// Encode implements bin.Encoder.
func (g *UpdatesGetChannelDifferenceRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode updates.getChannelDifference#3173d78 as nil")
	}
	b.PutID(UpdatesGetChannelDifferenceRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *UpdatesGetChannelDifferenceRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode updates.getChannelDifference#3173d78 as nil")
	}
	g.SetFlags()
	if err := g.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode updates.getChannelDifference#3173d78: field flags: %w", err)
	}
	if g.Channel == nil {
		return fmt.Errorf("unable to encode updates.getChannelDifference#3173d78: field channel is nil")
	}
	if err := g.Channel.Encode(b); err != nil {
		return fmt.Errorf("unable to encode updates.getChannelDifference#3173d78: field channel: %w", err)
	}
	if g.Filter == nil {
		return fmt.Errorf("unable to encode updates.getChannelDifference#3173d78: field filter is nil")
	}
	if err := g.Filter.Encode(b); err != nil {
		return fmt.Errorf("unable to encode updates.getChannelDifference#3173d78: field filter: %w", err)
	}
	b.PutInt(g.Pts)
	b.PutInt(g.Limit)
	return nil
}

// Decode implements bin.Decoder.
func (g *UpdatesGetChannelDifferenceRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode updates.getChannelDifference#3173d78 to nil")
	}
	if err := b.ConsumeID(UpdatesGetChannelDifferenceRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode updates.getChannelDifference#3173d78: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *UpdatesGetChannelDifferenceRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode updates.getChannelDifference#3173d78 to nil")
	}
	{
		if err := g.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode updates.getChannelDifference#3173d78: field flags: %w", err)
		}
	}
	g.Force = g.Flags.Has(0)
	{
		value, err := DecodeInputChannel(b)
		if err != nil {
			return fmt.Errorf("unable to decode updates.getChannelDifference#3173d78: field channel: %w", err)
		}
		g.Channel = value
	}
	{
		value, err := DecodeChannelMessagesFilter(b)
		if err != nil {
			return fmt.Errorf("unable to decode updates.getChannelDifference#3173d78: field filter: %w", err)
		}
		g.Filter = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode updates.getChannelDifference#3173d78: field pts: %w", err)
		}
		g.Pts = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode updates.getChannelDifference#3173d78: field limit: %w", err)
		}
		g.Limit = value
	}
	return nil
}

// SetForce sets value of Force conditional field.
func (g *UpdatesGetChannelDifferenceRequest) SetForce(value bool) {
	if value {
		g.Flags.Set(0)
		g.Force = true
	} else {
		g.Flags.Unset(0)
		g.Force = false
	}
}

// GetForce returns value of Force conditional field.
func (g *UpdatesGetChannelDifferenceRequest) GetForce() (value bool) {
	if g == nil {
		return
	}
	return g.Flags.Has(0)
}

// GetChannel returns value of Channel field.
func (g *UpdatesGetChannelDifferenceRequest) GetChannel() (value InputChannelClass) {
	if g == nil {
		return
	}
	return g.Channel
}

// GetFilter returns value of Filter field.
func (g *UpdatesGetChannelDifferenceRequest) GetFilter() (value ChannelMessagesFilterClass) {
	if g == nil {
		return
	}
	return g.Filter
}

// GetPts returns value of Pts field.
func (g *UpdatesGetChannelDifferenceRequest) GetPts() (value int) {
	if g == nil {
		return
	}
	return g.Pts
}

// GetLimit returns value of Limit field.
func (g *UpdatesGetChannelDifferenceRequest) GetLimit() (value int) {
	if g == nil {
		return
	}
	return g.Limit
}

// GetChannelAsNotEmpty returns mapped value of Channel field.
func (g *UpdatesGetChannelDifferenceRequest) GetChannelAsNotEmpty() (NotEmptyInputChannel, bool) {
	return g.Channel.AsNotEmpty()
}

// GetFilterAsNotEmpty returns mapped value of Filter field.
func (g *UpdatesGetChannelDifferenceRequest) GetFilterAsNotEmpty() (*ChannelMessagesFilter, bool) {
	return g.Filter.AsNotEmpty()
}

// UpdatesGetChannelDifference invokes method updates.getChannelDifference#3173d78 returning error if any.
// Returns the difference between the current state of updates of a certain channel and
// transmitted.
//
// Possible errors:
//
//	400 CHANNEL_INVALID: The provided channel is invalid.
//	400 CHANNEL_PRIVATE: You haven't joined this channel/supergroup.
//	403 CHANNEL_PUBLIC_GROUP_NA: channel/supergroup not available.
//	403 CHAT_WRITE_FORBIDDEN: You can't write in this chat.
//	400 FROM_MESSAGE_BOT_DISABLED: Bots can't use fromMessage min constructors.
//	400 MSG_ID_INVALID: Invalid message ID provided.
//	400 PERSISTENT_TIMESTAMP_EMPTY: Persistent timestamp empty.
//	400 PERSISTENT_TIMESTAMP_INVALID: Persistent timestamp invalid.
//	500 PERSISTENT_TIMESTAMP_OUTDATED: Channel internal replication issues, try again later (treat this like an RPC_CALL_FAIL).
//	400 PINNED_DIALOGS_TOO_MUCH: Too many pinned dialogs.
//	400 RANGES_INVALID: Invalid range provided.
//	400 USER_BANNED_IN_CHANNEL: You're banned from sending messages in supergroups/channels.
//
// See https://core.telegram.org/method/updates.getChannelDifference for reference.
// Can be used by bots.
func (c *Client) UpdatesGetChannelDifference(ctx context.Context, request *UpdatesGetChannelDifferenceRequest) (UpdatesChannelDifferenceClass, error) {
	var result UpdatesChannelDifferenceBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.ChannelDifference, nil
}
