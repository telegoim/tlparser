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

// GetVideoChatAvailableParticipantsRequest represents TL type `getVideoChatAvailableParticipants#c45da305`.
type GetVideoChatAvailableParticipantsRequest struct {
	// Chat identifier
	ChatID int64
}

// GetVideoChatAvailableParticipantsRequestTypeID is TL type id of GetVideoChatAvailableParticipantsRequest.
const GetVideoChatAvailableParticipantsRequestTypeID = 0xc45da305

// Ensuring interfaces in compile-time for GetVideoChatAvailableParticipantsRequest.
var (
	_ bin.Encoder     = &GetVideoChatAvailableParticipantsRequest{}
	_ bin.Decoder     = &GetVideoChatAvailableParticipantsRequest{}
	_ bin.BareEncoder = &GetVideoChatAvailableParticipantsRequest{}
	_ bin.BareDecoder = &GetVideoChatAvailableParticipantsRequest{}
)

func (g *GetVideoChatAvailableParticipantsRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.ChatID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetVideoChatAvailableParticipantsRequest) String() string {
	if g == nil {
		return "GetVideoChatAvailableParticipantsRequest(nil)"
	}
	type Alias GetVideoChatAvailableParticipantsRequest
	return fmt.Sprintf("GetVideoChatAvailableParticipantsRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetVideoChatAvailableParticipantsRequest) TypeID() uint32 {
	return GetVideoChatAvailableParticipantsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetVideoChatAvailableParticipantsRequest) TypeName() string {
	return "getVideoChatAvailableParticipants"
}

// TypeInfo returns info about TL type.
func (g *GetVideoChatAvailableParticipantsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getVideoChatAvailableParticipants",
		ID:   GetVideoChatAvailableParticipantsRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ChatID",
			SchemaName: "chat_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *GetVideoChatAvailableParticipantsRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getVideoChatAvailableParticipants#c45da305 as nil")
	}
	b.PutID(GetVideoChatAvailableParticipantsRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetVideoChatAvailableParticipantsRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getVideoChatAvailableParticipants#c45da305 as nil")
	}
	b.PutInt53(g.ChatID)
	return nil
}

// Decode implements bin.Decoder.
func (g *GetVideoChatAvailableParticipantsRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getVideoChatAvailableParticipants#c45da305 to nil")
	}
	if err := b.ConsumeID(GetVideoChatAvailableParticipantsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getVideoChatAvailableParticipants#c45da305: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetVideoChatAvailableParticipantsRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getVideoChatAvailableParticipants#c45da305 to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode getVideoChatAvailableParticipants#c45da305: field chat_id: %w", err)
		}
		g.ChatID = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (g *GetVideoChatAvailableParticipantsRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode getVideoChatAvailableParticipants#c45da305 as nil")
	}
	b.ObjStart()
	b.PutID("getVideoChatAvailableParticipants")
	b.FieldStart("chat_id")
	b.PutInt53(g.ChatID)
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (g *GetVideoChatAvailableParticipantsRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if g == nil {
		return fmt.Errorf("can't decode getVideoChatAvailableParticipants#c45da305 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("getVideoChatAvailableParticipants"); err != nil {
				return fmt.Errorf("unable to decode getVideoChatAvailableParticipants#c45da305: %w", err)
			}
		case "chat_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode getVideoChatAvailableParticipants#c45da305: field chat_id: %w", err)
			}
			g.ChatID = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetChatID returns value of ChatID field.
func (g *GetVideoChatAvailableParticipantsRequest) GetChatID() (value int64) {
	if g == nil {
		return
	}
	return g.ChatID
}

// GetVideoChatAvailableParticipants invokes method getVideoChatAvailableParticipants#c45da305 returning error if any.
func (c *Client) GetVideoChatAvailableParticipants(ctx context.Context, chatid int64) (*MessageSenders, error) {
	var result MessageSenders

	request := &GetVideoChatAvailableParticipantsRequest{
		ChatID: chatid,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
