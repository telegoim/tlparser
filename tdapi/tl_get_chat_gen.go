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

// GetChatRequest represents TL type `getChat#6f421440`.
type GetChatRequest struct {
	// Chat identifier
	ChatID int64
}

// GetChatRequestTypeID is TL type id of GetChatRequest.
const GetChatRequestTypeID = 0x6f421440

// Ensuring interfaces in compile-time for GetChatRequest.
var (
	_ bin.Encoder     = &GetChatRequest{}
	_ bin.Decoder     = &GetChatRequest{}
	_ bin.BareEncoder = &GetChatRequest{}
	_ bin.BareDecoder = &GetChatRequest{}
)

func (g *GetChatRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.ChatID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetChatRequest) String() string {
	if g == nil {
		return "GetChatRequest(nil)"
	}
	type Alias GetChatRequest
	return fmt.Sprintf("GetChatRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetChatRequest) TypeID() uint32 {
	return GetChatRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetChatRequest) TypeName() string {
	return "getChat"
}

// TypeInfo returns info about TL type.
func (g *GetChatRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getChat",
		ID:   GetChatRequestTypeID,
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
func (g *GetChatRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getChat#6f421440 as nil")
	}
	b.PutID(GetChatRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetChatRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getChat#6f421440 as nil")
	}
	b.PutInt53(g.ChatID)
	return nil
}

// Decode implements bin.Decoder.
func (g *GetChatRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getChat#6f421440 to nil")
	}
	if err := b.ConsumeID(GetChatRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getChat#6f421440: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetChatRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getChat#6f421440 to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode getChat#6f421440: field chat_id: %w", err)
		}
		g.ChatID = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (g *GetChatRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode getChat#6f421440 as nil")
	}
	b.ObjStart()
	b.PutID("getChat")
	b.FieldStart("chat_id")
	b.PutInt53(g.ChatID)
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (g *GetChatRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if g == nil {
		return fmt.Errorf("can't decode getChat#6f421440 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("getChat"); err != nil {
				return fmt.Errorf("unable to decode getChat#6f421440: %w", err)
			}
		case "chat_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode getChat#6f421440: field chat_id: %w", err)
			}
			g.ChatID = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetChatID returns value of ChatID field.
func (g *GetChatRequest) GetChatID() (value int64) {
	if g == nil {
		return
	}
	return g.ChatID
}

// GetChat invokes method getChat#6f421440 returning error if any.
func (c *Client) GetChat(ctx context.Context, chatid int64) (*Chat, error) {
	var result Chat

	request := &GetChatRequest{
		ChatID: chatid,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
