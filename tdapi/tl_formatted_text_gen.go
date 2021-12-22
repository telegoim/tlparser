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

// FormattedText represents TL type `formattedText#a38d39ee`.
type FormattedText struct {
	// The text
	Text string
	// Entities contained in the text. Entities can be nested, but must not mutually
	// intersect with each other.
	Entities []TextEntity
}

// FormattedTextTypeID is TL type id of FormattedText.
const FormattedTextTypeID = 0xa38d39ee

// Ensuring interfaces in compile-time for FormattedText.
var (
	_ bin.Encoder     = &FormattedText{}
	_ bin.Decoder     = &FormattedText{}
	_ bin.BareEncoder = &FormattedText{}
	_ bin.BareDecoder = &FormattedText{}
)

func (f *FormattedText) Zero() bool {
	if f == nil {
		return true
	}
	if !(f.Text == "") {
		return false
	}
	if !(f.Entities == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (f *FormattedText) String() string {
	if f == nil {
		return "FormattedText(nil)"
	}
	type Alias FormattedText
	return fmt.Sprintf("FormattedText%+v", Alias(*f))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*FormattedText) TypeID() uint32 {
	return FormattedTextTypeID
}

// TypeName returns name of type in TL schema.
func (*FormattedText) TypeName() string {
	return "formattedText"
}

// TypeInfo returns info about TL type.
func (f *FormattedText) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "formattedText",
		ID:   FormattedTextTypeID,
	}
	if f == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Text",
			SchemaName: "text",
		},
		{
			Name:       "Entities",
			SchemaName: "entities",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (f *FormattedText) Encode(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't encode formattedText#a38d39ee as nil")
	}
	b.PutID(FormattedTextTypeID)
	return f.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (f *FormattedText) EncodeBare(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't encode formattedText#a38d39ee as nil")
	}
	b.PutString(f.Text)
	b.PutInt(len(f.Entities))
	for idx, v := range f.Entities {
		if err := v.EncodeBare(b); err != nil {
			return fmt.Errorf("unable to encode bare formattedText#a38d39ee: field entities element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (f *FormattedText) Decode(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't decode formattedText#a38d39ee to nil")
	}
	if err := b.ConsumeID(FormattedTextTypeID); err != nil {
		return fmt.Errorf("unable to decode formattedText#a38d39ee: %w", err)
	}
	return f.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (f *FormattedText) DecodeBare(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't decode formattedText#a38d39ee to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode formattedText#a38d39ee: field text: %w", err)
		}
		f.Text = value
	}
	{
		headerLen, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode formattedText#a38d39ee: field entities: %w", err)
		}

		if headerLen > 0 {
			f.Entities = make([]TextEntity, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value TextEntity
			if err := value.DecodeBare(b); err != nil {
				return fmt.Errorf("unable to decode bare formattedText#a38d39ee: field entities: %w", err)
			}
			f.Entities = append(f.Entities, value)
		}
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (f *FormattedText) EncodeTDLibJSON(b tdjson.Encoder) error {
	if f == nil {
		return fmt.Errorf("can't encode formattedText#a38d39ee as nil")
	}
	b.ObjStart()
	b.PutID("formattedText")
	b.FieldStart("text")
	b.PutString(f.Text)
	b.FieldStart("entities")
	b.ArrStart()
	for idx, v := range f.Entities {
		if err := v.EncodeTDLibJSON(b); err != nil {
			return fmt.Errorf("unable to encode formattedText#a38d39ee: field entities element with index %d: %w", idx, err)
		}
	}
	b.ArrEnd()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (f *FormattedText) DecodeTDLibJSON(b tdjson.Decoder) error {
	if f == nil {
		return fmt.Errorf("can't decode formattedText#a38d39ee to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("formattedText"); err != nil {
				return fmt.Errorf("unable to decode formattedText#a38d39ee: %w", err)
			}
		case "text":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode formattedText#a38d39ee: field text: %w", err)
			}
			f.Text = value
		case "entities":
			if err := b.Arr(func(b tdjson.Decoder) error {
				var value TextEntity
				if err := value.DecodeTDLibJSON(b); err != nil {
					return fmt.Errorf("unable to decode formattedText#a38d39ee: field entities: %w", err)
				}
				f.Entities = append(f.Entities, value)
				return nil
			}); err != nil {
				return fmt.Errorf("unable to decode formattedText#a38d39ee: field entities: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetText returns value of Text field.
func (f *FormattedText) GetText() (value string) {
	if f == nil {
		return
	}
	return f.Text
}

// GetEntities returns value of Entities field.
func (f *FormattedText) GetEntities() (value []TextEntity) {
	if f == nil {
		return
	}
	return f.Entities
}
