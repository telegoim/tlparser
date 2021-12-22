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

// SetStickerSetThumbnailRequest represents TL type `setStickerSetThumbnail#4952fa88`.
type SetStickerSetThumbnailRequest struct {
	// Sticker set owner
	UserID int64
	// Sticker set name
	Name string
	// Thumbnail to set in PNG or TGS format; pass null to remove the sticker set thumbnail.
	// Animated thumbnail must be set for animated sticker sets and only for them
	Thumbnail InputFileClass
}

// SetStickerSetThumbnailRequestTypeID is TL type id of SetStickerSetThumbnailRequest.
const SetStickerSetThumbnailRequestTypeID = 0x4952fa88

// Ensuring interfaces in compile-time for SetStickerSetThumbnailRequest.
var (
	_ bin.Encoder     = &SetStickerSetThumbnailRequest{}
	_ bin.Decoder     = &SetStickerSetThumbnailRequest{}
	_ bin.BareEncoder = &SetStickerSetThumbnailRequest{}
	_ bin.BareDecoder = &SetStickerSetThumbnailRequest{}
)

func (s *SetStickerSetThumbnailRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.UserID == 0) {
		return false
	}
	if !(s.Name == "") {
		return false
	}
	if !(s.Thumbnail == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *SetStickerSetThumbnailRequest) String() string {
	if s == nil {
		return "SetStickerSetThumbnailRequest(nil)"
	}
	type Alias SetStickerSetThumbnailRequest
	return fmt.Sprintf("SetStickerSetThumbnailRequest%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*SetStickerSetThumbnailRequest) TypeID() uint32 {
	return SetStickerSetThumbnailRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*SetStickerSetThumbnailRequest) TypeName() string {
	return "setStickerSetThumbnail"
}

// TypeInfo returns info about TL type.
func (s *SetStickerSetThumbnailRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "setStickerSetThumbnail",
		ID:   SetStickerSetThumbnailRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "UserID",
			SchemaName: "user_id",
		},
		{
			Name:       "Name",
			SchemaName: "name",
		},
		{
			Name:       "Thumbnail",
			SchemaName: "thumbnail",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *SetStickerSetThumbnailRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode setStickerSetThumbnail#4952fa88 as nil")
	}
	b.PutID(SetStickerSetThumbnailRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *SetStickerSetThumbnailRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode setStickerSetThumbnail#4952fa88 as nil")
	}
	b.PutInt53(s.UserID)
	b.PutString(s.Name)
	if s.Thumbnail == nil {
		return fmt.Errorf("unable to encode setStickerSetThumbnail#4952fa88: field thumbnail is nil")
	}
	if err := s.Thumbnail.Encode(b); err != nil {
		return fmt.Errorf("unable to encode setStickerSetThumbnail#4952fa88: field thumbnail: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *SetStickerSetThumbnailRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode setStickerSetThumbnail#4952fa88 to nil")
	}
	if err := b.ConsumeID(SetStickerSetThumbnailRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode setStickerSetThumbnail#4952fa88: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *SetStickerSetThumbnailRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode setStickerSetThumbnail#4952fa88 to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode setStickerSetThumbnail#4952fa88: field user_id: %w", err)
		}
		s.UserID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode setStickerSetThumbnail#4952fa88: field name: %w", err)
		}
		s.Name = value
	}
	{
		value, err := DecodeInputFile(b)
		if err != nil {
			return fmt.Errorf("unable to decode setStickerSetThumbnail#4952fa88: field thumbnail: %w", err)
		}
		s.Thumbnail = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (s *SetStickerSetThumbnailRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if s == nil {
		return fmt.Errorf("can't encode setStickerSetThumbnail#4952fa88 as nil")
	}
	b.ObjStart()
	b.PutID("setStickerSetThumbnail")
	b.FieldStart("user_id")
	b.PutInt53(s.UserID)
	b.FieldStart("name")
	b.PutString(s.Name)
	b.FieldStart("thumbnail")
	if s.Thumbnail == nil {
		return fmt.Errorf("unable to encode setStickerSetThumbnail#4952fa88: field thumbnail is nil")
	}
	if err := s.Thumbnail.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode setStickerSetThumbnail#4952fa88: field thumbnail: %w", err)
	}
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (s *SetStickerSetThumbnailRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if s == nil {
		return fmt.Errorf("can't decode setStickerSetThumbnail#4952fa88 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("setStickerSetThumbnail"); err != nil {
				return fmt.Errorf("unable to decode setStickerSetThumbnail#4952fa88: %w", err)
			}
		case "user_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode setStickerSetThumbnail#4952fa88: field user_id: %w", err)
			}
			s.UserID = value
		case "name":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode setStickerSetThumbnail#4952fa88: field name: %w", err)
			}
			s.Name = value
		case "thumbnail":
			value, err := DecodeTDLibJSONInputFile(b)
			if err != nil {
				return fmt.Errorf("unable to decode setStickerSetThumbnail#4952fa88: field thumbnail: %w", err)
			}
			s.Thumbnail = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetUserID returns value of UserID field.
func (s *SetStickerSetThumbnailRequest) GetUserID() (value int64) {
	if s == nil {
		return
	}
	return s.UserID
}

// GetName returns value of Name field.
func (s *SetStickerSetThumbnailRequest) GetName() (value string) {
	if s == nil {
		return
	}
	return s.Name
}

// GetThumbnail returns value of Thumbnail field.
func (s *SetStickerSetThumbnailRequest) GetThumbnail() (value InputFileClass) {
	if s == nil {
		return
	}
	return s.Thumbnail
}

// SetStickerSetThumbnail invokes method setStickerSetThumbnail#4952fa88 returning error if any.
func (c *Client) SetStickerSetThumbnail(ctx context.Context, request *SetStickerSetThumbnailRequest) (*StickerSet, error) {
	var result StickerSet

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
