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

// RecommendedChatFilter represents TL type `recommendedChatFilter#2260ee2`.
type RecommendedChatFilter struct {
	// The chat filter
	Filter ChatFilter
	// Describes a recommended chat filter
	Description string
}

// RecommendedChatFilterTypeID is TL type id of RecommendedChatFilter.
const RecommendedChatFilterTypeID = 0x2260ee2

// Ensuring interfaces in compile-time for RecommendedChatFilter.
var (
	_ bin.Encoder     = &RecommendedChatFilter{}
	_ bin.Decoder     = &RecommendedChatFilter{}
	_ bin.BareEncoder = &RecommendedChatFilter{}
	_ bin.BareDecoder = &RecommendedChatFilter{}
)

func (r *RecommendedChatFilter) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.Filter.Zero()) {
		return false
	}
	if !(r.Description == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *RecommendedChatFilter) String() string {
	if r == nil {
		return "RecommendedChatFilter(nil)"
	}
	type Alias RecommendedChatFilter
	return fmt.Sprintf("RecommendedChatFilter%+v", Alias(*r))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*RecommendedChatFilter) TypeID() uint32 {
	return RecommendedChatFilterTypeID
}

// TypeName returns name of type in TL schema.
func (*RecommendedChatFilter) TypeName() string {
	return "recommendedChatFilter"
}

// TypeInfo returns info about TL type.
func (r *RecommendedChatFilter) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "recommendedChatFilter",
		ID:   RecommendedChatFilterTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Filter",
			SchemaName: "filter",
		},
		{
			Name:       "Description",
			SchemaName: "description",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (r *RecommendedChatFilter) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode recommendedChatFilter#2260ee2 as nil")
	}
	b.PutID(RecommendedChatFilterTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *RecommendedChatFilter) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode recommendedChatFilter#2260ee2 as nil")
	}
	if err := r.Filter.Encode(b); err != nil {
		return fmt.Errorf("unable to encode recommendedChatFilter#2260ee2: field filter: %w", err)
	}
	b.PutString(r.Description)
	return nil
}

// Decode implements bin.Decoder.
func (r *RecommendedChatFilter) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode recommendedChatFilter#2260ee2 to nil")
	}
	if err := b.ConsumeID(RecommendedChatFilterTypeID); err != nil {
		return fmt.Errorf("unable to decode recommendedChatFilter#2260ee2: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *RecommendedChatFilter) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode recommendedChatFilter#2260ee2 to nil")
	}
	{
		if err := r.Filter.Decode(b); err != nil {
			return fmt.Errorf("unable to decode recommendedChatFilter#2260ee2: field filter: %w", err)
		}
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode recommendedChatFilter#2260ee2: field description: %w", err)
		}
		r.Description = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (r *RecommendedChatFilter) EncodeTDLibJSON(b tdjson.Encoder) error {
	if r == nil {
		return fmt.Errorf("can't encode recommendedChatFilter#2260ee2 as nil")
	}
	b.ObjStart()
	b.PutID("recommendedChatFilter")
	b.FieldStart("filter")
	if err := r.Filter.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode recommendedChatFilter#2260ee2: field filter: %w", err)
	}
	b.FieldStart("description")
	b.PutString(r.Description)
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (r *RecommendedChatFilter) DecodeTDLibJSON(b tdjson.Decoder) error {
	if r == nil {
		return fmt.Errorf("can't decode recommendedChatFilter#2260ee2 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("recommendedChatFilter"); err != nil {
				return fmt.Errorf("unable to decode recommendedChatFilter#2260ee2: %w", err)
			}
		case "filter":
			if err := r.Filter.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode recommendedChatFilter#2260ee2: field filter: %w", err)
			}
		case "description":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode recommendedChatFilter#2260ee2: field description: %w", err)
			}
			r.Description = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetFilter returns value of Filter field.
func (r *RecommendedChatFilter) GetFilter() (value ChatFilter) {
	if r == nil {
		return
	}
	return r.Filter
}

// GetDescription returns value of Description field.
func (r *RecommendedChatFilter) GetDescription() (value string) {
	if r == nil {
		return
	}
	return r.Description
}
