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

// CreatePrivateChatRequest represents TL type `createPrivateChat#c7825b09`.
type CreatePrivateChatRequest struct {
	// User identifier
	UserID int64
	// If true, the chat will be created without network request. In this case all
	// information about the chat except its type, title and photo can be incorrect
	Force bool
}

// CreatePrivateChatRequestTypeID is TL type id of CreatePrivateChatRequest.
const CreatePrivateChatRequestTypeID = 0xc7825b09

// Ensuring interfaces in compile-time for CreatePrivateChatRequest.
var (
	_ bin.Encoder     = &CreatePrivateChatRequest{}
	_ bin.Decoder     = &CreatePrivateChatRequest{}
	_ bin.BareEncoder = &CreatePrivateChatRequest{}
	_ bin.BareDecoder = &CreatePrivateChatRequest{}
)

func (c *CreatePrivateChatRequest) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.UserID == 0) {
		return false
	}
	if !(c.Force == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *CreatePrivateChatRequest) String() string {
	if c == nil {
		return "CreatePrivateChatRequest(nil)"
	}
	type Alias CreatePrivateChatRequest
	return fmt.Sprintf("CreatePrivateChatRequest%+v", Alias(*c))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*CreatePrivateChatRequest) TypeID() uint32 {
	return CreatePrivateChatRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*CreatePrivateChatRequest) TypeName() string {
	return "createPrivateChat"
}

// TypeInfo returns info about TL type.
func (c *CreatePrivateChatRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "createPrivateChat",
		ID:   CreatePrivateChatRequestTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "UserID",
			SchemaName: "user_id",
		},
		{
			Name:       "Force",
			SchemaName: "force",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *CreatePrivateChatRequest) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode createPrivateChat#c7825b09 as nil")
	}
	b.PutID(CreatePrivateChatRequestTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *CreatePrivateChatRequest) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode createPrivateChat#c7825b09 as nil")
	}
	b.PutInt53(c.UserID)
	b.PutBool(c.Force)
	return nil
}

// Decode implements bin.Decoder.
func (c *CreatePrivateChatRequest) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode createPrivateChat#c7825b09 to nil")
	}
	if err := b.ConsumeID(CreatePrivateChatRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode createPrivateChat#c7825b09: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *CreatePrivateChatRequest) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode createPrivateChat#c7825b09 to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode createPrivateChat#c7825b09: field user_id: %w", err)
		}
		c.UserID = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode createPrivateChat#c7825b09: field force: %w", err)
		}
		c.Force = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (c *CreatePrivateChatRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if c == nil {
		return fmt.Errorf("can't encode createPrivateChat#c7825b09 as nil")
	}
	b.ObjStart()
	b.PutID("createPrivateChat")
	b.FieldStart("user_id")
	b.PutInt53(c.UserID)
	b.FieldStart("force")
	b.PutBool(c.Force)
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (c *CreatePrivateChatRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if c == nil {
		return fmt.Errorf("can't decode createPrivateChat#c7825b09 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("createPrivateChat"); err != nil {
				return fmt.Errorf("unable to decode createPrivateChat#c7825b09: %w", err)
			}
		case "user_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode createPrivateChat#c7825b09: field user_id: %w", err)
			}
			c.UserID = value
		case "force":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode createPrivateChat#c7825b09: field force: %w", err)
			}
			c.Force = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetUserID returns value of UserID field.
func (c *CreatePrivateChatRequest) GetUserID() (value int64) {
	if c == nil {
		return
	}
	return c.UserID
}

// GetForce returns value of Force field.
func (c *CreatePrivateChatRequest) GetForce() (value bool) {
	if c == nil {
		return
	}
	return c.Force
}

// CreatePrivateChat invokes method createPrivateChat#c7825b09 returning error if any.
func (c *Client) CreatePrivateChat(ctx context.Context, request *CreatePrivateChatRequest) (*Chat, error) {
	var result Chat

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
