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

// OrderInfo represents TL type `orderInfo#2ebad96e`.
type OrderInfo struct {
	// Name of the user
	Name string
	// Phone number of the user
	PhoneNumber string
	// Email address of the user
	EmailAddress string
	// Shipping address for this order; may be null
	ShippingAddress Address
}

// OrderInfoTypeID is TL type id of OrderInfo.
const OrderInfoTypeID = 0x2ebad96e

// Ensuring interfaces in compile-time for OrderInfo.
var (
	_ bin.Encoder     = &OrderInfo{}
	_ bin.Decoder     = &OrderInfo{}
	_ bin.BareEncoder = &OrderInfo{}
	_ bin.BareDecoder = &OrderInfo{}
)

func (o *OrderInfo) Zero() bool {
	if o == nil {
		return true
	}
	if !(o.Name == "") {
		return false
	}
	if !(o.PhoneNumber == "") {
		return false
	}
	if !(o.EmailAddress == "") {
		return false
	}
	if !(o.ShippingAddress.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (o *OrderInfo) String() string {
	if o == nil {
		return "OrderInfo(nil)"
	}
	type Alias OrderInfo
	return fmt.Sprintf("OrderInfo%+v", Alias(*o))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*OrderInfo) TypeID() uint32 {
	return OrderInfoTypeID
}

// TypeName returns name of type in TL schema.
func (*OrderInfo) TypeName() string {
	return "orderInfo"
}

// TypeInfo returns info about TL type.
func (o *OrderInfo) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "orderInfo",
		ID:   OrderInfoTypeID,
	}
	if o == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Name",
			SchemaName: "name",
		},
		{
			Name:       "PhoneNumber",
			SchemaName: "phone_number",
		},
		{
			Name:       "EmailAddress",
			SchemaName: "email_address",
		},
		{
			Name:       "ShippingAddress",
			SchemaName: "shipping_address",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (o *OrderInfo) Encode(b *bin.Buffer) error {
	if o == nil {
		return fmt.Errorf("can't encode orderInfo#2ebad96e as nil")
	}
	b.PutID(OrderInfoTypeID)
	return o.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (o *OrderInfo) EncodeBare(b *bin.Buffer) error {
	if o == nil {
		return fmt.Errorf("can't encode orderInfo#2ebad96e as nil")
	}
	b.PutString(o.Name)
	b.PutString(o.PhoneNumber)
	b.PutString(o.EmailAddress)
	if err := o.ShippingAddress.Encode(b); err != nil {
		return fmt.Errorf("unable to encode orderInfo#2ebad96e: field shipping_address: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (o *OrderInfo) Decode(b *bin.Buffer) error {
	if o == nil {
		return fmt.Errorf("can't decode orderInfo#2ebad96e to nil")
	}
	if err := b.ConsumeID(OrderInfoTypeID); err != nil {
		return fmt.Errorf("unable to decode orderInfo#2ebad96e: %w", err)
	}
	return o.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (o *OrderInfo) DecodeBare(b *bin.Buffer) error {
	if o == nil {
		return fmt.Errorf("can't decode orderInfo#2ebad96e to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode orderInfo#2ebad96e: field name: %w", err)
		}
		o.Name = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode orderInfo#2ebad96e: field phone_number: %w", err)
		}
		o.PhoneNumber = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode orderInfo#2ebad96e: field email_address: %w", err)
		}
		o.EmailAddress = value
	}
	{
		if err := o.ShippingAddress.Decode(b); err != nil {
			return fmt.Errorf("unable to decode orderInfo#2ebad96e: field shipping_address: %w", err)
		}
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (o *OrderInfo) EncodeTDLibJSON(b tdjson.Encoder) error {
	if o == nil {
		return fmt.Errorf("can't encode orderInfo#2ebad96e as nil")
	}
	b.ObjStart()
	b.PutID("orderInfo")
	b.FieldStart("name")
	b.PutString(o.Name)
	b.FieldStart("phone_number")
	b.PutString(o.PhoneNumber)
	b.FieldStart("email_address")
	b.PutString(o.EmailAddress)
	b.FieldStart("shipping_address")
	if err := o.ShippingAddress.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode orderInfo#2ebad96e: field shipping_address: %w", err)
	}
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (o *OrderInfo) DecodeTDLibJSON(b tdjson.Decoder) error {
	if o == nil {
		return fmt.Errorf("can't decode orderInfo#2ebad96e to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("orderInfo"); err != nil {
				return fmt.Errorf("unable to decode orderInfo#2ebad96e: %w", err)
			}
		case "name":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode orderInfo#2ebad96e: field name: %w", err)
			}
			o.Name = value
		case "phone_number":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode orderInfo#2ebad96e: field phone_number: %w", err)
			}
			o.PhoneNumber = value
		case "email_address":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode orderInfo#2ebad96e: field email_address: %w", err)
			}
			o.EmailAddress = value
		case "shipping_address":
			if err := o.ShippingAddress.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode orderInfo#2ebad96e: field shipping_address: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetName returns value of Name field.
func (o *OrderInfo) GetName() (value string) {
	if o == nil {
		return
	}
	return o.Name
}

// GetPhoneNumber returns value of PhoneNumber field.
func (o *OrderInfo) GetPhoneNumber() (value string) {
	if o == nil {
		return
	}
	return o.PhoneNumber
}

// GetEmailAddress returns value of EmailAddress field.
func (o *OrderInfo) GetEmailAddress() (value string) {
	if o == nil {
		return
	}
	return o.EmailAddress
}

// GetShippingAddress returns value of ShippingAddress field.
func (o *OrderInfo) GetShippingAddress() (value Address) {
	if o == nil {
		return
	}
	return o.ShippingAddress
}
