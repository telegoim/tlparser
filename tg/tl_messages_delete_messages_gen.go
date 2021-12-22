// Code generated by gotdgen, DO NOT EDIT.

package tg

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

// MessagesDeleteMessagesRequest represents TL type `messages.deleteMessages#e58e95d2`.
// Deletes messages by their identifiers.
//
// See https://core.telegram.org/method/messages.deleteMessages for reference.
type MessagesDeleteMessagesRequest struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Whether to delete messages for all participants of the chat
	Revoke bool
	// Message ID list
	ID []int
}

// MessagesDeleteMessagesRequestTypeID is TL type id of MessagesDeleteMessagesRequest.
const MessagesDeleteMessagesRequestTypeID = 0xe58e95d2

// Ensuring interfaces in compile-time for MessagesDeleteMessagesRequest.
var (
	_ bin.Encoder     = &MessagesDeleteMessagesRequest{}
	_ bin.Decoder     = &MessagesDeleteMessagesRequest{}
	_ bin.BareEncoder = &MessagesDeleteMessagesRequest{}
	_ bin.BareDecoder = &MessagesDeleteMessagesRequest{}
)

func (d *MessagesDeleteMessagesRequest) Zero() bool {
	if d == nil {
		return true
	}
	if !(d.Flags.Zero()) {
		return false
	}
	if !(d.Revoke == false) {
		return false
	}
	if !(d.ID == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (d *MessagesDeleteMessagesRequest) String() string {
	if d == nil {
		return "MessagesDeleteMessagesRequest(nil)"
	}
	type Alias MessagesDeleteMessagesRequest
	return fmt.Sprintf("MessagesDeleteMessagesRequest%+v", Alias(*d))
}

// FillFrom fills MessagesDeleteMessagesRequest from given interface.
func (d *MessagesDeleteMessagesRequest) FillFrom(from interface {
	GetRevoke() (value bool)
	GetID() (value []int)
}) {
	d.Revoke = from.GetRevoke()
	d.ID = from.GetID()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesDeleteMessagesRequest) TypeID() uint32 {
	return MessagesDeleteMessagesRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesDeleteMessagesRequest) TypeName() string {
	return "messages.deleteMessages"
}

// TypeInfo returns info about TL type.
func (d *MessagesDeleteMessagesRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.deleteMessages",
		ID:   MessagesDeleteMessagesRequestTypeID,
	}
	if d == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Revoke",
			SchemaName: "revoke",
			Null:       !d.Flags.Has(0),
		},
		{
			Name:       "ID",
			SchemaName: "id",
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (d *MessagesDeleteMessagesRequest) SetFlags() {
	if !(d.Revoke == false) {
		d.Flags.Set(0)
	}
}

// Encode implements bin.Encoder.
func (d *MessagesDeleteMessagesRequest) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode messages.deleteMessages#e58e95d2 as nil")
	}
	b.PutID(MessagesDeleteMessagesRequestTypeID)
	return d.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (d *MessagesDeleteMessagesRequest) EncodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode messages.deleteMessages#e58e95d2 as nil")
	}
	d.SetFlags()
	if err := d.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.deleteMessages#e58e95d2: field flags: %w", err)
	}
	b.PutVectorHeader(len(d.ID))
	for _, v := range d.ID {
		b.PutInt(v)
	}
	return nil
}

// Decode implements bin.Decoder.
func (d *MessagesDeleteMessagesRequest) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode messages.deleteMessages#e58e95d2 to nil")
	}
	if err := b.ConsumeID(MessagesDeleteMessagesRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.deleteMessages#e58e95d2: %w", err)
	}
	return d.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (d *MessagesDeleteMessagesRequest) DecodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode messages.deleteMessages#e58e95d2 to nil")
	}
	{
		if err := d.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messages.deleteMessages#e58e95d2: field flags: %w", err)
		}
	}
	d.Revoke = d.Flags.Has(0)
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.deleteMessages#e58e95d2: field id: %w", err)
		}

		if headerLen > 0 {
			d.ID = make([]int, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Int()
			if err != nil {
				return fmt.Errorf("unable to decode messages.deleteMessages#e58e95d2: field id: %w", err)
			}
			d.ID = append(d.ID, value)
		}
	}
	return nil
}

// SetRevoke sets value of Revoke conditional field.
func (d *MessagesDeleteMessagesRequest) SetRevoke(value bool) {
	if value {
		d.Flags.Set(0)
		d.Revoke = true
	} else {
		d.Flags.Unset(0)
		d.Revoke = false
	}
}

// GetRevoke returns value of Revoke conditional field.
func (d *MessagesDeleteMessagesRequest) GetRevoke() (value bool) {
	if d == nil {
		return
	}
	return d.Flags.Has(0)
}

// GetID returns value of ID field.
func (d *MessagesDeleteMessagesRequest) GetID() (value []int) {
	if d == nil {
		return
	}
	return d.ID
}

// MessagesDeleteMessages invokes method messages.deleteMessages#e58e95d2 returning error if any.
// Deletes messages by their identifiers.
//
// Possible errors:
//  403 MESSAGE_DELETE_FORBIDDEN: You can't delete one of the messages you tried to delete, most likely because it is a service message.
//
// See https://core.telegram.org/method/messages.deleteMessages for reference.
// Can be used by bots.
func (c *Client) MessagesDeleteMessages(ctx context.Context, request *MessagesDeleteMessagesRequest) (*MessagesAffectedMessages, error) {
	var result MessagesAffectedMessages

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
