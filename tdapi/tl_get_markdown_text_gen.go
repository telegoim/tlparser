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

// GetMarkdownTextRequest represents TL type `getMarkdownText#9ce7228`.
type GetMarkdownTextRequest struct {
	// The text
	Text FormattedText
}

// GetMarkdownTextRequestTypeID is TL type id of GetMarkdownTextRequest.
const GetMarkdownTextRequestTypeID = 0x9ce7228

// Ensuring interfaces in compile-time for GetMarkdownTextRequest.
var (
	_ bin.Encoder     = &GetMarkdownTextRequest{}
	_ bin.Decoder     = &GetMarkdownTextRequest{}
	_ bin.BareEncoder = &GetMarkdownTextRequest{}
	_ bin.BareDecoder = &GetMarkdownTextRequest{}
)

func (g *GetMarkdownTextRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Text.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetMarkdownTextRequest) String() string {
	if g == nil {
		return "GetMarkdownTextRequest(nil)"
	}
	type Alias GetMarkdownTextRequest
	return fmt.Sprintf("GetMarkdownTextRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetMarkdownTextRequest) TypeID() uint32 {
	return GetMarkdownTextRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetMarkdownTextRequest) TypeName() string {
	return "getMarkdownText"
}

// TypeInfo returns info about TL type.
func (g *GetMarkdownTextRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getMarkdownText",
		ID:   GetMarkdownTextRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Text",
			SchemaName: "text",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *GetMarkdownTextRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getMarkdownText#9ce7228 as nil")
	}
	b.PutID(GetMarkdownTextRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetMarkdownTextRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getMarkdownText#9ce7228 as nil")
	}
	if err := g.Text.Encode(b); err != nil {
		return fmt.Errorf("unable to encode getMarkdownText#9ce7228: field text: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (g *GetMarkdownTextRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getMarkdownText#9ce7228 to nil")
	}
	if err := b.ConsumeID(GetMarkdownTextRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getMarkdownText#9ce7228: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetMarkdownTextRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getMarkdownText#9ce7228 to nil")
	}
	{
		if err := g.Text.Decode(b); err != nil {
			return fmt.Errorf("unable to decode getMarkdownText#9ce7228: field text: %w", err)
		}
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (g *GetMarkdownTextRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode getMarkdownText#9ce7228 as nil")
	}
	b.ObjStart()
	b.PutID("getMarkdownText")
	b.FieldStart("text")
	if err := g.Text.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode getMarkdownText#9ce7228: field text: %w", err)
	}
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (g *GetMarkdownTextRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if g == nil {
		return fmt.Errorf("can't decode getMarkdownText#9ce7228 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("getMarkdownText"); err != nil {
				return fmt.Errorf("unable to decode getMarkdownText#9ce7228: %w", err)
			}
		case "text":
			if err := g.Text.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode getMarkdownText#9ce7228: field text: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetText returns value of Text field.
func (g *GetMarkdownTextRequest) GetText() (value FormattedText) {
	if g == nil {
		return
	}
	return g.Text
}

// GetMarkdownText invokes method getMarkdownText#9ce7228 returning error if any.
func (c *Client) GetMarkdownText(ctx context.Context, text FormattedText) (*FormattedText, error) {
	var result FormattedText

	request := &GetMarkdownTextRequest{
		Text: text,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
