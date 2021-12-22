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

// GetFileDownloadedPrefixSizeRequest represents TL type `getFileDownloadedPrefixSize#9c8724a0`.
type GetFileDownloadedPrefixSizeRequest struct {
	// Identifier of the file
	FileID int32
	// Offset from which downloaded prefix size needs to be calculated
	Offset int32
}

// GetFileDownloadedPrefixSizeRequestTypeID is TL type id of GetFileDownloadedPrefixSizeRequest.
const GetFileDownloadedPrefixSizeRequestTypeID = 0x9c8724a0

// Ensuring interfaces in compile-time for GetFileDownloadedPrefixSizeRequest.
var (
	_ bin.Encoder     = &GetFileDownloadedPrefixSizeRequest{}
	_ bin.Decoder     = &GetFileDownloadedPrefixSizeRequest{}
	_ bin.BareEncoder = &GetFileDownloadedPrefixSizeRequest{}
	_ bin.BareDecoder = &GetFileDownloadedPrefixSizeRequest{}
)

func (g *GetFileDownloadedPrefixSizeRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.FileID == 0) {
		return false
	}
	if !(g.Offset == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetFileDownloadedPrefixSizeRequest) String() string {
	if g == nil {
		return "GetFileDownloadedPrefixSizeRequest(nil)"
	}
	type Alias GetFileDownloadedPrefixSizeRequest
	return fmt.Sprintf("GetFileDownloadedPrefixSizeRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetFileDownloadedPrefixSizeRequest) TypeID() uint32 {
	return GetFileDownloadedPrefixSizeRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetFileDownloadedPrefixSizeRequest) TypeName() string {
	return "getFileDownloadedPrefixSize"
}

// TypeInfo returns info about TL type.
func (g *GetFileDownloadedPrefixSizeRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getFileDownloadedPrefixSize",
		ID:   GetFileDownloadedPrefixSizeRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "FileID",
			SchemaName: "file_id",
		},
		{
			Name:       "Offset",
			SchemaName: "offset",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *GetFileDownloadedPrefixSizeRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getFileDownloadedPrefixSize#9c8724a0 as nil")
	}
	b.PutID(GetFileDownloadedPrefixSizeRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetFileDownloadedPrefixSizeRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getFileDownloadedPrefixSize#9c8724a0 as nil")
	}
	b.PutInt32(g.FileID)
	b.PutInt32(g.Offset)
	return nil
}

// Decode implements bin.Decoder.
func (g *GetFileDownloadedPrefixSizeRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getFileDownloadedPrefixSize#9c8724a0 to nil")
	}
	if err := b.ConsumeID(GetFileDownloadedPrefixSizeRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getFileDownloadedPrefixSize#9c8724a0: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetFileDownloadedPrefixSizeRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getFileDownloadedPrefixSize#9c8724a0 to nil")
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode getFileDownloadedPrefixSize#9c8724a0: field file_id: %w", err)
		}
		g.FileID = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode getFileDownloadedPrefixSize#9c8724a0: field offset: %w", err)
		}
		g.Offset = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (g *GetFileDownloadedPrefixSizeRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode getFileDownloadedPrefixSize#9c8724a0 as nil")
	}
	b.ObjStart()
	b.PutID("getFileDownloadedPrefixSize")
	b.FieldStart("file_id")
	b.PutInt32(g.FileID)
	b.FieldStart("offset")
	b.PutInt32(g.Offset)
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (g *GetFileDownloadedPrefixSizeRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if g == nil {
		return fmt.Errorf("can't decode getFileDownloadedPrefixSize#9c8724a0 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("getFileDownloadedPrefixSize"); err != nil {
				return fmt.Errorf("unable to decode getFileDownloadedPrefixSize#9c8724a0: %w", err)
			}
		case "file_id":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode getFileDownloadedPrefixSize#9c8724a0: field file_id: %w", err)
			}
			g.FileID = value
		case "offset":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode getFileDownloadedPrefixSize#9c8724a0: field offset: %w", err)
			}
			g.Offset = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetFileID returns value of FileID field.
func (g *GetFileDownloadedPrefixSizeRequest) GetFileID() (value int32) {
	if g == nil {
		return
	}
	return g.FileID
}

// GetOffset returns value of Offset field.
func (g *GetFileDownloadedPrefixSizeRequest) GetOffset() (value int32) {
	if g == nil {
		return
	}
	return g.Offset
}

// GetFileDownloadedPrefixSize invokes method getFileDownloadedPrefixSize#9c8724a0 returning error if any.
func (c *Client) GetFileDownloadedPrefixSize(ctx context.Context, request *GetFileDownloadedPrefixSizeRequest) (*Count, error) {
	var result Count

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
