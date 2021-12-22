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

// GetBackgroundsRequest represents TL type `getBackgrounds#ed88bf9`.
type GetBackgroundsRequest struct {
	// True, if the backgrounds must be ordered for dark theme
	ForDarkTheme bool
}

// GetBackgroundsRequestTypeID is TL type id of GetBackgroundsRequest.
const GetBackgroundsRequestTypeID = 0xed88bf9

// Ensuring interfaces in compile-time for GetBackgroundsRequest.
var (
	_ bin.Encoder     = &GetBackgroundsRequest{}
	_ bin.Decoder     = &GetBackgroundsRequest{}
	_ bin.BareEncoder = &GetBackgroundsRequest{}
	_ bin.BareDecoder = &GetBackgroundsRequest{}
)

func (g *GetBackgroundsRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.ForDarkTheme == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetBackgroundsRequest) String() string {
	if g == nil {
		return "GetBackgroundsRequest(nil)"
	}
	type Alias GetBackgroundsRequest
	return fmt.Sprintf("GetBackgroundsRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetBackgroundsRequest) TypeID() uint32 {
	return GetBackgroundsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetBackgroundsRequest) TypeName() string {
	return "getBackgrounds"
}

// TypeInfo returns info about TL type.
func (g *GetBackgroundsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getBackgrounds",
		ID:   GetBackgroundsRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ForDarkTheme",
			SchemaName: "for_dark_theme",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *GetBackgroundsRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getBackgrounds#ed88bf9 as nil")
	}
	b.PutID(GetBackgroundsRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetBackgroundsRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getBackgrounds#ed88bf9 as nil")
	}
	b.PutBool(g.ForDarkTheme)
	return nil
}

// Decode implements bin.Decoder.
func (g *GetBackgroundsRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getBackgrounds#ed88bf9 to nil")
	}
	if err := b.ConsumeID(GetBackgroundsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getBackgrounds#ed88bf9: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetBackgroundsRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getBackgrounds#ed88bf9 to nil")
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode getBackgrounds#ed88bf9: field for_dark_theme: %w", err)
		}
		g.ForDarkTheme = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (g *GetBackgroundsRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode getBackgrounds#ed88bf9 as nil")
	}
	b.ObjStart()
	b.PutID("getBackgrounds")
	b.FieldStart("for_dark_theme")
	b.PutBool(g.ForDarkTheme)
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (g *GetBackgroundsRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if g == nil {
		return fmt.Errorf("can't decode getBackgrounds#ed88bf9 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("getBackgrounds"); err != nil {
				return fmt.Errorf("unable to decode getBackgrounds#ed88bf9: %w", err)
			}
		case "for_dark_theme":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode getBackgrounds#ed88bf9: field for_dark_theme: %w", err)
			}
			g.ForDarkTheme = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetForDarkTheme returns value of ForDarkTheme field.
func (g *GetBackgroundsRequest) GetForDarkTheme() (value bool) {
	if g == nil {
		return
	}
	return g.ForDarkTheme
}

// GetBackgrounds invokes method getBackgrounds#ed88bf9 returning error if any.
func (c *Client) GetBackgrounds(ctx context.Context, fordarktheme bool) (*Backgrounds, error) {
	var result Backgrounds

	request := &GetBackgroundsRequest{
		ForDarkTheme: fordarktheme,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
