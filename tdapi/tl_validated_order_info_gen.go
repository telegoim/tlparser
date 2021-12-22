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

// ValidatedOrderInfo represents TL type `validatedOrderInfo#ac585f14`.
type ValidatedOrderInfo struct {
	// Temporary identifier of the order information
	OrderInfoID string
	// Available shipping options
	ShippingOptions []ShippingOption
}

// ValidatedOrderInfoTypeID is TL type id of ValidatedOrderInfo.
const ValidatedOrderInfoTypeID = 0xac585f14

// Ensuring interfaces in compile-time for ValidatedOrderInfo.
var (
	_ bin.Encoder     = &ValidatedOrderInfo{}
	_ bin.Decoder     = &ValidatedOrderInfo{}
	_ bin.BareEncoder = &ValidatedOrderInfo{}
	_ bin.BareDecoder = &ValidatedOrderInfo{}
)

func (v *ValidatedOrderInfo) Zero() bool {
	if v == nil {
		return true
	}
	if !(v.OrderInfoID == "") {
		return false
	}
	if !(v.ShippingOptions == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (v *ValidatedOrderInfo) String() string {
	if v == nil {
		return "ValidatedOrderInfo(nil)"
	}
	type Alias ValidatedOrderInfo
	return fmt.Sprintf("ValidatedOrderInfo%+v", Alias(*v))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ValidatedOrderInfo) TypeID() uint32 {
	return ValidatedOrderInfoTypeID
}

// TypeName returns name of type in TL schema.
func (*ValidatedOrderInfo) TypeName() string {
	return "validatedOrderInfo"
}

// TypeInfo returns info about TL type.
func (v *ValidatedOrderInfo) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "validatedOrderInfo",
		ID:   ValidatedOrderInfoTypeID,
	}
	if v == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "OrderInfoID",
			SchemaName: "order_info_id",
		},
		{
			Name:       "ShippingOptions",
			SchemaName: "shipping_options",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (v *ValidatedOrderInfo) Encode(b *bin.Buffer) error {
	if v == nil {
		return fmt.Errorf("can't encode validatedOrderInfo#ac585f14 as nil")
	}
	b.PutID(ValidatedOrderInfoTypeID)
	return v.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (v *ValidatedOrderInfo) EncodeBare(b *bin.Buffer) error {
	if v == nil {
		return fmt.Errorf("can't encode validatedOrderInfo#ac585f14 as nil")
	}
	b.PutString(v.OrderInfoID)
	b.PutInt(len(v.ShippingOptions))
	for idx, v := range v.ShippingOptions {
		if err := v.EncodeBare(b); err != nil {
			return fmt.Errorf("unable to encode bare validatedOrderInfo#ac585f14: field shipping_options element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (v *ValidatedOrderInfo) Decode(b *bin.Buffer) error {
	if v == nil {
		return fmt.Errorf("can't decode validatedOrderInfo#ac585f14 to nil")
	}
	if err := b.ConsumeID(ValidatedOrderInfoTypeID); err != nil {
		return fmt.Errorf("unable to decode validatedOrderInfo#ac585f14: %w", err)
	}
	return v.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (v *ValidatedOrderInfo) DecodeBare(b *bin.Buffer) error {
	if v == nil {
		return fmt.Errorf("can't decode validatedOrderInfo#ac585f14 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode validatedOrderInfo#ac585f14: field order_info_id: %w", err)
		}
		v.OrderInfoID = value
	}
	{
		headerLen, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode validatedOrderInfo#ac585f14: field shipping_options: %w", err)
		}

		if headerLen > 0 {
			v.ShippingOptions = make([]ShippingOption, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value ShippingOption
			if err := value.DecodeBare(b); err != nil {
				return fmt.Errorf("unable to decode bare validatedOrderInfo#ac585f14: field shipping_options: %w", err)
			}
			v.ShippingOptions = append(v.ShippingOptions, value)
		}
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (v *ValidatedOrderInfo) EncodeTDLibJSON(b tdjson.Encoder) error {
	if v == nil {
		return fmt.Errorf("can't encode validatedOrderInfo#ac585f14 as nil")
	}
	b.ObjStart()
	b.PutID("validatedOrderInfo")
	b.FieldStart("order_info_id")
	b.PutString(v.OrderInfoID)
	b.FieldStart("shipping_options")
	b.ArrStart()
	for idx, v := range v.ShippingOptions {
		if err := v.EncodeTDLibJSON(b); err != nil {
			return fmt.Errorf("unable to encode validatedOrderInfo#ac585f14: field shipping_options element with index %d: %w", idx, err)
		}
	}
	b.ArrEnd()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (v *ValidatedOrderInfo) DecodeTDLibJSON(b tdjson.Decoder) error {
	if v == nil {
		return fmt.Errorf("can't decode validatedOrderInfo#ac585f14 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("validatedOrderInfo"); err != nil {
				return fmt.Errorf("unable to decode validatedOrderInfo#ac585f14: %w", err)
			}
		case "order_info_id":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode validatedOrderInfo#ac585f14: field order_info_id: %w", err)
			}
			v.OrderInfoID = value
		case "shipping_options":
			if err := b.Arr(func(b tdjson.Decoder) error {
				var value ShippingOption
				if err := value.DecodeTDLibJSON(b); err != nil {
					return fmt.Errorf("unable to decode validatedOrderInfo#ac585f14: field shipping_options: %w", err)
				}
				v.ShippingOptions = append(v.ShippingOptions, value)
				return nil
			}); err != nil {
				return fmt.Errorf("unable to decode validatedOrderInfo#ac585f14: field shipping_options: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetOrderInfoID returns value of OrderInfoID field.
func (v *ValidatedOrderInfo) GetOrderInfoID() (value string) {
	if v == nil {
		return
	}
	return v.OrderInfoID
}

// GetShippingOptions returns value of ShippingOptions field.
func (v *ValidatedOrderInfo) GetShippingOptions() (value []ShippingOption) {
	if v == nil {
		return
	}
	return v.ShippingOptions
}
