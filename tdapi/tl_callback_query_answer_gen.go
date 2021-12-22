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

// CallbackQueryAnswer represents TL type `callbackQueryAnswer#1582685d`.
type CallbackQueryAnswer struct {
	// Text of the answer
	Text string
	// True, if an alert must be shown to the user instead of a toast notification
	ShowAlert bool
	// URL to be opened
	URL string
}

// CallbackQueryAnswerTypeID is TL type id of CallbackQueryAnswer.
const CallbackQueryAnswerTypeID = 0x1582685d

// Ensuring interfaces in compile-time for CallbackQueryAnswer.
var (
	_ bin.Encoder     = &CallbackQueryAnswer{}
	_ bin.Decoder     = &CallbackQueryAnswer{}
	_ bin.BareEncoder = &CallbackQueryAnswer{}
	_ bin.BareDecoder = &CallbackQueryAnswer{}
)

func (c *CallbackQueryAnswer) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.Text == "") {
		return false
	}
	if !(c.ShowAlert == false) {
		return false
	}
	if !(c.URL == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *CallbackQueryAnswer) String() string {
	if c == nil {
		return "CallbackQueryAnswer(nil)"
	}
	type Alias CallbackQueryAnswer
	return fmt.Sprintf("CallbackQueryAnswer%+v", Alias(*c))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*CallbackQueryAnswer) TypeID() uint32 {
	return CallbackQueryAnswerTypeID
}

// TypeName returns name of type in TL schema.
func (*CallbackQueryAnswer) TypeName() string {
	return "callbackQueryAnswer"
}

// TypeInfo returns info about TL type.
func (c *CallbackQueryAnswer) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "callbackQueryAnswer",
		ID:   CallbackQueryAnswerTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Text",
			SchemaName: "text",
		},
		{
			Name:       "ShowAlert",
			SchemaName: "show_alert",
		},
		{
			Name:       "URL",
			SchemaName: "url",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *CallbackQueryAnswer) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode callbackQueryAnswer#1582685d as nil")
	}
	b.PutID(CallbackQueryAnswerTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *CallbackQueryAnswer) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode callbackQueryAnswer#1582685d as nil")
	}
	b.PutString(c.Text)
	b.PutBool(c.ShowAlert)
	b.PutString(c.URL)
	return nil
}

// Decode implements bin.Decoder.
func (c *CallbackQueryAnswer) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode callbackQueryAnswer#1582685d to nil")
	}
	if err := b.ConsumeID(CallbackQueryAnswerTypeID); err != nil {
		return fmt.Errorf("unable to decode callbackQueryAnswer#1582685d: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *CallbackQueryAnswer) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode callbackQueryAnswer#1582685d to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode callbackQueryAnswer#1582685d: field text: %w", err)
		}
		c.Text = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode callbackQueryAnswer#1582685d: field show_alert: %w", err)
		}
		c.ShowAlert = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode callbackQueryAnswer#1582685d: field url: %w", err)
		}
		c.URL = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (c *CallbackQueryAnswer) EncodeTDLibJSON(b tdjson.Encoder) error {
	if c == nil {
		return fmt.Errorf("can't encode callbackQueryAnswer#1582685d as nil")
	}
	b.ObjStart()
	b.PutID("callbackQueryAnswer")
	b.FieldStart("text")
	b.PutString(c.Text)
	b.FieldStart("show_alert")
	b.PutBool(c.ShowAlert)
	b.FieldStart("url")
	b.PutString(c.URL)
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (c *CallbackQueryAnswer) DecodeTDLibJSON(b tdjson.Decoder) error {
	if c == nil {
		return fmt.Errorf("can't decode callbackQueryAnswer#1582685d to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("callbackQueryAnswer"); err != nil {
				return fmt.Errorf("unable to decode callbackQueryAnswer#1582685d: %w", err)
			}
		case "text":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode callbackQueryAnswer#1582685d: field text: %w", err)
			}
			c.Text = value
		case "show_alert":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode callbackQueryAnswer#1582685d: field show_alert: %w", err)
			}
			c.ShowAlert = value
		case "url":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode callbackQueryAnswer#1582685d: field url: %w", err)
			}
			c.URL = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetText returns value of Text field.
func (c *CallbackQueryAnswer) GetText() (value string) {
	if c == nil {
		return
	}
	return c.Text
}

// GetShowAlert returns value of ShowAlert field.
func (c *CallbackQueryAnswer) GetShowAlert() (value bool) {
	if c == nil {
		return
	}
	return c.ShowAlert
}

// GetURL returns value of URL field.
func (c *CallbackQueryAnswer) GetURL() (value string) {
	if c == nil {
		return
	}
	return c.URL
}
