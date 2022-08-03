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

// InputInvoiceMessage represents TL type `inputInvoiceMessage#58dcea10`.
type InputInvoiceMessage struct {
	// Chat identifier of the message
	ChatID int64
	// Message identifier
	MessageID int64
}

// InputInvoiceMessageTypeID is TL type id of InputInvoiceMessage.
const InputInvoiceMessageTypeID = 0x58dcea10

// construct implements constructor of InputInvoiceClass.
func (i InputInvoiceMessage) construct() InputInvoiceClass { return &i }

// Ensuring interfaces in compile-time for InputInvoiceMessage.
var (
	_ bin.Encoder     = &InputInvoiceMessage{}
	_ bin.Decoder     = &InputInvoiceMessage{}
	_ bin.BareEncoder = &InputInvoiceMessage{}
	_ bin.BareDecoder = &InputInvoiceMessage{}

	_ InputInvoiceClass = &InputInvoiceMessage{}
)

func (i *InputInvoiceMessage) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.ChatID == 0) {
		return false
	}
	if !(i.MessageID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputInvoiceMessage) String() string {
	if i == nil {
		return "InputInvoiceMessage(nil)"
	}
	type Alias InputInvoiceMessage
	return fmt.Sprintf("InputInvoiceMessage%+v", Alias(*i))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InputInvoiceMessage) TypeID() uint32 {
	return InputInvoiceMessageTypeID
}

// TypeName returns name of type in TL schema.
func (*InputInvoiceMessage) TypeName() string {
	return "inputInvoiceMessage"
}

// TypeInfo returns info about TL type.
func (i *InputInvoiceMessage) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inputInvoiceMessage",
		ID:   InputInvoiceMessageTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ChatID",
			SchemaName: "chat_id",
		},
		{
			Name:       "MessageID",
			SchemaName: "message_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (i *InputInvoiceMessage) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputInvoiceMessage#58dcea10 as nil")
	}
	b.PutID(InputInvoiceMessageTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InputInvoiceMessage) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputInvoiceMessage#58dcea10 as nil")
	}
	b.PutInt53(i.ChatID)
	b.PutInt53(i.MessageID)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputInvoiceMessage) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputInvoiceMessage#58dcea10 to nil")
	}
	if err := b.ConsumeID(InputInvoiceMessageTypeID); err != nil {
		return fmt.Errorf("unable to decode inputInvoiceMessage#58dcea10: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InputInvoiceMessage) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputInvoiceMessage#58dcea10 to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode inputInvoiceMessage#58dcea10: field chat_id: %w", err)
		}
		i.ChatID = value
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode inputInvoiceMessage#58dcea10: field message_id: %w", err)
		}
		i.MessageID = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (i *InputInvoiceMessage) EncodeTDLibJSON(b tdjson.Encoder) error {
	if i == nil {
		return fmt.Errorf("can't encode inputInvoiceMessage#58dcea10 as nil")
	}
	b.ObjStart()
	b.PutID("inputInvoiceMessage")
	b.Comma()
	b.FieldStart("chat_id")
	b.PutInt53(i.ChatID)
	b.Comma()
	b.FieldStart("message_id")
	b.PutInt53(i.MessageID)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (i *InputInvoiceMessage) DecodeTDLibJSON(b tdjson.Decoder) error {
	if i == nil {
		return fmt.Errorf("can't decode inputInvoiceMessage#58dcea10 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("inputInvoiceMessage"); err != nil {
				return fmt.Errorf("unable to decode inputInvoiceMessage#58dcea10: %w", err)
			}
		case "chat_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode inputInvoiceMessage#58dcea10: field chat_id: %w", err)
			}
			i.ChatID = value
		case "message_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode inputInvoiceMessage#58dcea10: field message_id: %w", err)
			}
			i.MessageID = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetChatID returns value of ChatID field.
func (i *InputInvoiceMessage) GetChatID() (value int64) {
	if i == nil {
		return
	}
	return i.ChatID
}

// GetMessageID returns value of MessageID field.
func (i *InputInvoiceMessage) GetMessageID() (value int64) {
	if i == nil {
		return
	}
	return i.MessageID
}

// InputInvoiceName represents TL type `inputInvoiceName#b1ca16f3`.
type InputInvoiceName struct {
	// Name of the invoice
	Name string
}

// InputInvoiceNameTypeID is TL type id of InputInvoiceName.
const InputInvoiceNameTypeID = 0xb1ca16f3

// construct implements constructor of InputInvoiceClass.
func (i InputInvoiceName) construct() InputInvoiceClass { return &i }

// Ensuring interfaces in compile-time for InputInvoiceName.
var (
	_ bin.Encoder     = &InputInvoiceName{}
	_ bin.Decoder     = &InputInvoiceName{}
	_ bin.BareEncoder = &InputInvoiceName{}
	_ bin.BareDecoder = &InputInvoiceName{}

	_ InputInvoiceClass = &InputInvoiceName{}
)

func (i *InputInvoiceName) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.Name == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputInvoiceName) String() string {
	if i == nil {
		return "InputInvoiceName(nil)"
	}
	type Alias InputInvoiceName
	return fmt.Sprintf("InputInvoiceName%+v", Alias(*i))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InputInvoiceName) TypeID() uint32 {
	return InputInvoiceNameTypeID
}

// TypeName returns name of type in TL schema.
func (*InputInvoiceName) TypeName() string {
	return "inputInvoiceName"
}

// TypeInfo returns info about TL type.
func (i *InputInvoiceName) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inputInvoiceName",
		ID:   InputInvoiceNameTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Name",
			SchemaName: "name",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (i *InputInvoiceName) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputInvoiceName#b1ca16f3 as nil")
	}
	b.PutID(InputInvoiceNameTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InputInvoiceName) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputInvoiceName#b1ca16f3 as nil")
	}
	b.PutString(i.Name)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputInvoiceName) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputInvoiceName#b1ca16f3 to nil")
	}
	if err := b.ConsumeID(InputInvoiceNameTypeID); err != nil {
		return fmt.Errorf("unable to decode inputInvoiceName#b1ca16f3: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InputInvoiceName) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputInvoiceName#b1ca16f3 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode inputInvoiceName#b1ca16f3: field name: %w", err)
		}
		i.Name = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (i *InputInvoiceName) EncodeTDLibJSON(b tdjson.Encoder) error {
	if i == nil {
		return fmt.Errorf("can't encode inputInvoiceName#b1ca16f3 as nil")
	}
	b.ObjStart()
	b.PutID("inputInvoiceName")
	b.Comma()
	b.FieldStart("name")
	b.PutString(i.Name)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (i *InputInvoiceName) DecodeTDLibJSON(b tdjson.Decoder) error {
	if i == nil {
		return fmt.Errorf("can't decode inputInvoiceName#b1ca16f3 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("inputInvoiceName"); err != nil {
				return fmt.Errorf("unable to decode inputInvoiceName#b1ca16f3: %w", err)
			}
		case "name":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode inputInvoiceName#b1ca16f3: field name: %w", err)
			}
			i.Name = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetName returns value of Name field.
func (i *InputInvoiceName) GetName() (value string) {
	if i == nil {
		return
	}
	return i.Name
}

// InputInvoiceClassName is schema name of InputInvoiceClass.
const InputInvoiceClassName = "InputInvoice"

// InputInvoiceClass represents InputInvoice generic type.
//
// Example:
//
//	g, err := tdapi.DecodeInputInvoice(buf)
//	if err != nil {
//	    panic(err)
//	}
//	switch v := g.(type) {
//	case *tdapi.InputInvoiceMessage: // inputInvoiceMessage#58dcea10
//	case *tdapi.InputInvoiceName: // inputInvoiceName#b1ca16f3
//	default: panic(v)
//	}
type InputInvoiceClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() InputInvoiceClass

	// TypeID returns type id in TL schema.
	//
	// See https://core.telegram.org/mtproto/TL-tl#remarks.
	TypeID() uint32
	// TypeName returns name of type in TL schema.
	TypeName() string
	// String implements fmt.Stringer.
	String() string
	// Zero returns true if current object has a zero value.
	Zero() bool

	EncodeTDLibJSON(b tdjson.Encoder) error
	DecodeTDLibJSON(b tdjson.Decoder) error
}

// DecodeInputInvoice implements binary de-serialization for InputInvoiceClass.
func DecodeInputInvoice(buf *bin.Buffer) (InputInvoiceClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case InputInvoiceMessageTypeID:
		// Decoding inputInvoiceMessage#58dcea10.
		v := InputInvoiceMessage{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputInvoiceClass: %w", err)
		}
		return &v, nil
	case InputInvoiceNameTypeID:
		// Decoding inputInvoiceName#b1ca16f3.
		v := InputInvoiceName{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputInvoiceClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode InputInvoiceClass: %w", bin.NewUnexpectedID(id))
	}
}

// DecodeTDLibJSONInputInvoice implements binary de-serialization for InputInvoiceClass.
func DecodeTDLibJSONInputInvoice(buf tdjson.Decoder) (InputInvoiceClass, error) {
	id, err := buf.FindTypeID()
	if err != nil {
		return nil, err
	}
	switch id {
	case "inputInvoiceMessage":
		// Decoding inputInvoiceMessage#58dcea10.
		v := InputInvoiceMessage{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputInvoiceClass: %w", err)
		}
		return &v, nil
	case "inputInvoiceName":
		// Decoding inputInvoiceName#b1ca16f3.
		v := InputInvoiceName{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputInvoiceClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode InputInvoiceClass: %w", tdjson.NewUnexpectedID(id))
	}
}

// InputInvoice boxes the InputInvoiceClass providing a helper.
type InputInvoiceBox struct {
	InputInvoice InputInvoiceClass
}

// Decode implements bin.Decoder for InputInvoiceBox.
func (b *InputInvoiceBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode InputInvoiceBox to nil")
	}
	v, err := DecodeInputInvoice(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.InputInvoice = v
	return nil
}

// Encode implements bin.Encode for InputInvoiceBox.
func (b *InputInvoiceBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.InputInvoice == nil {
		return fmt.Errorf("unable to encode InputInvoiceClass as nil")
	}
	return b.InputInvoice.Encode(buf)
}

// DecodeTDLibJSON implements bin.Decoder for InputInvoiceBox.
func (b *InputInvoiceBox) DecodeTDLibJSON(buf tdjson.Decoder) error {
	if b == nil {
		return fmt.Errorf("unable to decode InputInvoiceBox to nil")
	}
	v, err := DecodeTDLibJSONInputInvoice(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.InputInvoice = v
	return nil
}

// EncodeTDLibJSON implements bin.Encode for InputInvoiceBox.
func (b *InputInvoiceBox) EncodeTDLibJSON(buf tdjson.Encoder) error {
	if b == nil || b.InputInvoice == nil {
		return fmt.Errorf("unable to encode InputInvoiceClass as nil")
	}
	return b.InputInvoice.EncodeTDLibJSON(buf)
}
