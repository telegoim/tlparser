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

// HelpEditUserInfoRequest represents TL type `help.editUserInfo#66b91b70`.
// Internal use
//
// See https://core.telegram.org/method/help.editUserInfo for reference.
type HelpEditUserInfoRequest struct {
	// User
	UserID InputUserClass
	// Message
	Message string
	// Message entities for styled text¹
	//
	// Links:
	//  1) https://core.telegram.org/api/entities
	Entities []MessageEntityClass
}

// HelpEditUserInfoRequestTypeID is TL type id of HelpEditUserInfoRequest.
const HelpEditUserInfoRequestTypeID = 0x66b91b70

// Ensuring interfaces in compile-time for HelpEditUserInfoRequest.
var (
	_ bin.Encoder     = &HelpEditUserInfoRequest{}
	_ bin.Decoder     = &HelpEditUserInfoRequest{}
	_ bin.BareEncoder = &HelpEditUserInfoRequest{}
	_ bin.BareDecoder = &HelpEditUserInfoRequest{}
)

func (e *HelpEditUserInfoRequest) Zero() bool {
	if e == nil {
		return true
	}
	if !(e.UserID == nil) {
		return false
	}
	if !(e.Message == "") {
		return false
	}
	if !(e.Entities == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (e *HelpEditUserInfoRequest) String() string {
	if e == nil {
		return "HelpEditUserInfoRequest(nil)"
	}
	type Alias HelpEditUserInfoRequest
	return fmt.Sprintf("HelpEditUserInfoRequest%+v", Alias(*e))
}

// FillFrom fills HelpEditUserInfoRequest from given interface.
func (e *HelpEditUserInfoRequest) FillFrom(from interface {
	GetUserID() (value InputUserClass)
	GetMessage() (value string)
	GetEntities() (value []MessageEntityClass)
}) {
	e.UserID = from.GetUserID()
	e.Message = from.GetMessage()
	e.Entities = from.GetEntities()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*HelpEditUserInfoRequest) TypeID() uint32 {
	return HelpEditUserInfoRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*HelpEditUserInfoRequest) TypeName() string {
	return "help.editUserInfo"
}

// TypeInfo returns info about TL type.
func (e *HelpEditUserInfoRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "help.editUserInfo",
		ID:   HelpEditUserInfoRequestTypeID,
	}
	if e == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "UserID",
			SchemaName: "user_id",
		},
		{
			Name:       "Message",
			SchemaName: "message",
		},
		{
			Name:       "Entities",
			SchemaName: "entities",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (e *HelpEditUserInfoRequest) Encode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode help.editUserInfo#66b91b70 as nil")
	}
	b.PutID(HelpEditUserInfoRequestTypeID)
	return e.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (e *HelpEditUserInfoRequest) EncodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode help.editUserInfo#66b91b70 as nil")
	}
	if e.UserID == nil {
		return fmt.Errorf("unable to encode help.editUserInfo#66b91b70: field user_id is nil")
	}
	if err := e.UserID.Encode(b); err != nil {
		return fmt.Errorf("unable to encode help.editUserInfo#66b91b70: field user_id: %w", err)
	}
	b.PutString(e.Message)
	b.PutVectorHeader(len(e.Entities))
	for idx, v := range e.Entities {
		if v == nil {
			return fmt.Errorf("unable to encode help.editUserInfo#66b91b70: field entities element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode help.editUserInfo#66b91b70: field entities element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (e *HelpEditUserInfoRequest) Decode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode help.editUserInfo#66b91b70 to nil")
	}
	if err := b.ConsumeID(HelpEditUserInfoRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode help.editUserInfo#66b91b70: %w", err)
	}
	return e.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (e *HelpEditUserInfoRequest) DecodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode help.editUserInfo#66b91b70 to nil")
	}
	{
		value, err := DecodeInputUser(b)
		if err != nil {
			return fmt.Errorf("unable to decode help.editUserInfo#66b91b70: field user_id: %w", err)
		}
		e.UserID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode help.editUserInfo#66b91b70: field message: %w", err)
		}
		e.Message = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode help.editUserInfo#66b91b70: field entities: %w", err)
		}

		if headerLen > 0 {
			e.Entities = make([]MessageEntityClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeMessageEntity(b)
			if err != nil {
				return fmt.Errorf("unable to decode help.editUserInfo#66b91b70: field entities: %w", err)
			}
			e.Entities = append(e.Entities, value)
		}
	}
	return nil
}

// GetUserID returns value of UserID field.
func (e *HelpEditUserInfoRequest) GetUserID() (value InputUserClass) {
	if e == nil {
		return
	}
	return e.UserID
}

// GetMessage returns value of Message field.
func (e *HelpEditUserInfoRequest) GetMessage() (value string) {
	if e == nil {
		return
	}
	return e.Message
}

// GetEntities returns value of Entities field.
func (e *HelpEditUserInfoRequest) GetEntities() (value []MessageEntityClass) {
	if e == nil {
		return
	}
	return e.Entities
}

// MapEntities returns field Entities wrapped in MessageEntityClassArray helper.
func (e *HelpEditUserInfoRequest) MapEntities() (value MessageEntityClassArray) {
	return MessageEntityClassArray(e.Entities)
}

// HelpEditUserInfo invokes method help.editUserInfo#66b91b70 returning error if any.
// Internal use
//
// Possible errors:
//  400 USER_INVALID: Invalid user provided.
//
// See https://core.telegram.org/method/help.editUserInfo for reference.
func (c *Client) HelpEditUserInfo(ctx context.Context, request *HelpEditUserInfoRequest) (HelpUserInfoClass, error) {
	var result HelpUserInfoBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.UserInfo, nil
}
