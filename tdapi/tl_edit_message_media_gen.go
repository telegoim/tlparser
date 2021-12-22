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

// EditMessageMediaRequest represents TL type `editMessageMedia#bb4b8713`.
type EditMessageMediaRequest struct {
	// The chat the message belongs to
	ChatID int64
	// Identifier of the message
	MessageID int64
	// The new message reply markup; pass null if none; for bots only
	ReplyMarkup ReplyMarkupClass
	// New content of the message. Must be one of the following types: inputMessageAnimation,
	// inputMessageAudio, inputMessageDocument, inputMessagePhoto or inputMessageVideo
	InputMessageContent InputMessageContentClass
}

// EditMessageMediaRequestTypeID is TL type id of EditMessageMediaRequest.
const EditMessageMediaRequestTypeID = 0xbb4b8713

// Ensuring interfaces in compile-time for EditMessageMediaRequest.
var (
	_ bin.Encoder     = &EditMessageMediaRequest{}
	_ bin.Decoder     = &EditMessageMediaRequest{}
	_ bin.BareEncoder = &EditMessageMediaRequest{}
	_ bin.BareDecoder = &EditMessageMediaRequest{}
)

func (e *EditMessageMediaRequest) Zero() bool {
	if e == nil {
		return true
	}
	if !(e.ChatID == 0) {
		return false
	}
	if !(e.MessageID == 0) {
		return false
	}
	if !(e.ReplyMarkup == nil) {
		return false
	}
	if !(e.InputMessageContent == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (e *EditMessageMediaRequest) String() string {
	if e == nil {
		return "EditMessageMediaRequest(nil)"
	}
	type Alias EditMessageMediaRequest
	return fmt.Sprintf("EditMessageMediaRequest%+v", Alias(*e))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*EditMessageMediaRequest) TypeID() uint32 {
	return EditMessageMediaRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*EditMessageMediaRequest) TypeName() string {
	return "editMessageMedia"
}

// TypeInfo returns info about TL type.
func (e *EditMessageMediaRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "editMessageMedia",
		ID:   EditMessageMediaRequestTypeID,
	}
	if e == nil {
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
		{
			Name:       "ReplyMarkup",
			SchemaName: "reply_markup",
		},
		{
			Name:       "InputMessageContent",
			SchemaName: "input_message_content",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (e *EditMessageMediaRequest) Encode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode editMessageMedia#bb4b8713 as nil")
	}
	b.PutID(EditMessageMediaRequestTypeID)
	return e.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (e *EditMessageMediaRequest) EncodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode editMessageMedia#bb4b8713 as nil")
	}
	b.PutInt53(e.ChatID)
	b.PutInt53(e.MessageID)
	if e.ReplyMarkup == nil {
		return fmt.Errorf("unable to encode editMessageMedia#bb4b8713: field reply_markup is nil")
	}
	if err := e.ReplyMarkup.Encode(b); err != nil {
		return fmt.Errorf("unable to encode editMessageMedia#bb4b8713: field reply_markup: %w", err)
	}
	if e.InputMessageContent == nil {
		return fmt.Errorf("unable to encode editMessageMedia#bb4b8713: field input_message_content is nil")
	}
	if err := e.InputMessageContent.Encode(b); err != nil {
		return fmt.Errorf("unable to encode editMessageMedia#bb4b8713: field input_message_content: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (e *EditMessageMediaRequest) Decode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode editMessageMedia#bb4b8713 to nil")
	}
	if err := b.ConsumeID(EditMessageMediaRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode editMessageMedia#bb4b8713: %w", err)
	}
	return e.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (e *EditMessageMediaRequest) DecodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode editMessageMedia#bb4b8713 to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode editMessageMedia#bb4b8713: field chat_id: %w", err)
		}
		e.ChatID = value
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode editMessageMedia#bb4b8713: field message_id: %w", err)
		}
		e.MessageID = value
	}
	{
		value, err := DecodeReplyMarkup(b)
		if err != nil {
			return fmt.Errorf("unable to decode editMessageMedia#bb4b8713: field reply_markup: %w", err)
		}
		e.ReplyMarkup = value
	}
	{
		value, err := DecodeInputMessageContent(b)
		if err != nil {
			return fmt.Errorf("unable to decode editMessageMedia#bb4b8713: field input_message_content: %w", err)
		}
		e.InputMessageContent = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (e *EditMessageMediaRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if e == nil {
		return fmt.Errorf("can't encode editMessageMedia#bb4b8713 as nil")
	}
	b.ObjStart()
	b.PutID("editMessageMedia")
	b.FieldStart("chat_id")
	b.PutInt53(e.ChatID)
	b.FieldStart("message_id")
	b.PutInt53(e.MessageID)
	b.FieldStart("reply_markup")
	if e.ReplyMarkup == nil {
		return fmt.Errorf("unable to encode editMessageMedia#bb4b8713: field reply_markup is nil")
	}
	if err := e.ReplyMarkup.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode editMessageMedia#bb4b8713: field reply_markup: %w", err)
	}
	b.FieldStart("input_message_content")
	if e.InputMessageContent == nil {
		return fmt.Errorf("unable to encode editMessageMedia#bb4b8713: field input_message_content is nil")
	}
	if err := e.InputMessageContent.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode editMessageMedia#bb4b8713: field input_message_content: %w", err)
	}
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (e *EditMessageMediaRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if e == nil {
		return fmt.Errorf("can't decode editMessageMedia#bb4b8713 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("editMessageMedia"); err != nil {
				return fmt.Errorf("unable to decode editMessageMedia#bb4b8713: %w", err)
			}
		case "chat_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode editMessageMedia#bb4b8713: field chat_id: %w", err)
			}
			e.ChatID = value
		case "message_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode editMessageMedia#bb4b8713: field message_id: %w", err)
			}
			e.MessageID = value
		case "reply_markup":
			value, err := DecodeTDLibJSONReplyMarkup(b)
			if err != nil {
				return fmt.Errorf("unable to decode editMessageMedia#bb4b8713: field reply_markup: %w", err)
			}
			e.ReplyMarkup = value
		case "input_message_content":
			value, err := DecodeTDLibJSONInputMessageContent(b)
			if err != nil {
				return fmt.Errorf("unable to decode editMessageMedia#bb4b8713: field input_message_content: %w", err)
			}
			e.InputMessageContent = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetChatID returns value of ChatID field.
func (e *EditMessageMediaRequest) GetChatID() (value int64) {
	if e == nil {
		return
	}
	return e.ChatID
}

// GetMessageID returns value of MessageID field.
func (e *EditMessageMediaRequest) GetMessageID() (value int64) {
	if e == nil {
		return
	}
	return e.MessageID
}

// GetReplyMarkup returns value of ReplyMarkup field.
func (e *EditMessageMediaRequest) GetReplyMarkup() (value ReplyMarkupClass) {
	if e == nil {
		return
	}
	return e.ReplyMarkup
}

// GetInputMessageContent returns value of InputMessageContent field.
func (e *EditMessageMediaRequest) GetInputMessageContent() (value InputMessageContentClass) {
	if e == nil {
		return
	}
	return e.InputMessageContent
}

// EditMessageMedia invokes method editMessageMedia#bb4b8713 returning error if any.
func (c *Client) EditMessageMedia(ctx context.Context, request *EditMessageMediaRequest) (*Message, error) {
	var result Message

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
