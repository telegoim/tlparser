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

// CheckAuthenticationCodeRequest represents TL type `checkAuthenticationCode#edfe44aa`.
type CheckAuthenticationCodeRequest struct {
	// Authentication code to check
	Code string
}

// CheckAuthenticationCodeRequestTypeID is TL type id of CheckAuthenticationCodeRequest.
const CheckAuthenticationCodeRequestTypeID = 0xedfe44aa

// Ensuring interfaces in compile-time for CheckAuthenticationCodeRequest.
var (
	_ bin.Encoder     = &CheckAuthenticationCodeRequest{}
	_ bin.Decoder     = &CheckAuthenticationCodeRequest{}
	_ bin.BareEncoder = &CheckAuthenticationCodeRequest{}
	_ bin.BareDecoder = &CheckAuthenticationCodeRequest{}
)

func (c *CheckAuthenticationCodeRequest) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.Code == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *CheckAuthenticationCodeRequest) String() string {
	if c == nil {
		return "CheckAuthenticationCodeRequest(nil)"
	}
	type Alias CheckAuthenticationCodeRequest
	return fmt.Sprintf("CheckAuthenticationCodeRequest%+v", Alias(*c))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*CheckAuthenticationCodeRequest) TypeID() uint32 {
	return CheckAuthenticationCodeRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*CheckAuthenticationCodeRequest) TypeName() string {
	return "checkAuthenticationCode"
}

// TypeInfo returns info about TL type.
func (c *CheckAuthenticationCodeRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "checkAuthenticationCode",
		ID:   CheckAuthenticationCodeRequestTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Code",
			SchemaName: "code",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *CheckAuthenticationCodeRequest) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode checkAuthenticationCode#edfe44aa as nil")
	}
	b.PutID(CheckAuthenticationCodeRequestTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *CheckAuthenticationCodeRequest) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode checkAuthenticationCode#edfe44aa as nil")
	}
	b.PutString(c.Code)
	return nil
}

// Decode implements bin.Decoder.
func (c *CheckAuthenticationCodeRequest) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode checkAuthenticationCode#edfe44aa to nil")
	}
	if err := b.ConsumeID(CheckAuthenticationCodeRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode checkAuthenticationCode#edfe44aa: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *CheckAuthenticationCodeRequest) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode checkAuthenticationCode#edfe44aa to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode checkAuthenticationCode#edfe44aa: field code: %w", err)
		}
		c.Code = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (c *CheckAuthenticationCodeRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if c == nil {
		return fmt.Errorf("can't encode checkAuthenticationCode#edfe44aa as nil")
	}
	b.ObjStart()
	b.PutID("checkAuthenticationCode")
	b.FieldStart("code")
	b.PutString(c.Code)
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (c *CheckAuthenticationCodeRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if c == nil {
		return fmt.Errorf("can't decode checkAuthenticationCode#edfe44aa to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("checkAuthenticationCode"); err != nil {
				return fmt.Errorf("unable to decode checkAuthenticationCode#edfe44aa: %w", err)
			}
		case "code":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode checkAuthenticationCode#edfe44aa: field code: %w", err)
			}
			c.Code = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetCode returns value of Code field.
func (c *CheckAuthenticationCodeRequest) GetCode() (value string) {
	if c == nil {
		return
	}
	return c.Code
}

// CheckAuthenticationCode invokes method checkAuthenticationCode#edfe44aa returning error if any.
func (c *Client) CheckAuthenticationCode(ctx context.Context, code string) error {
	var ok Ok

	request := &CheckAuthenticationCodeRequest{
		Code: code,
	}
	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
