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

// StickersCreateStickerSetRequest represents TL type `stickers.createStickerSet#9021ab67`.
// Create a stickerset, bots only.
//
// See https://core.telegram.org/method/stickers.createStickerSet for reference.
type StickersCreateStickerSetRequest struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Whether this is a mask stickerset
	Masks bool
	// Whether this is an animated stickerset
	Animated bool
	// Videos field of StickersCreateStickerSetRequest.
	Videos bool
	// Stickerset owner
	UserID InputUserClass
	// Stickerset name, 1-64 chars
	Title string
	// Sticker set name. Can contain only English letters, digits and underscores. Must end
	// with "by" ( is case insensitive); 1-64 characters
	ShortName string
	// Thumbnail
	//
	// Use SetThumb and GetThumb helpers.
	Thumb InputDocumentClass
	// Stickers
	Stickers []InputStickerSetItem
	// Used when importing stickers using the sticker import SDKs¹, specifies the name of
	// the software that created the stickers
	//
	// Links:
	//  1) https://core.telegram.org/import-stickers
	//
	// Use SetSoftware and GetSoftware helpers.
	Software string
}

// StickersCreateStickerSetRequestTypeID is TL type id of StickersCreateStickerSetRequest.
const StickersCreateStickerSetRequestTypeID = 0x9021ab67

// Ensuring interfaces in compile-time for StickersCreateStickerSetRequest.
var (
	_ bin.Encoder     = &StickersCreateStickerSetRequest{}
	_ bin.Decoder     = &StickersCreateStickerSetRequest{}
	_ bin.BareEncoder = &StickersCreateStickerSetRequest{}
	_ bin.BareDecoder = &StickersCreateStickerSetRequest{}
)

func (c *StickersCreateStickerSetRequest) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.Flags.Zero()) {
		return false
	}
	if !(c.Masks == false) {
		return false
	}
	if !(c.Animated == false) {
		return false
	}
	if !(c.Videos == false) {
		return false
	}
	if !(c.UserID == nil) {
		return false
	}
	if !(c.Title == "") {
		return false
	}
	if !(c.ShortName == "") {
		return false
	}
	if !(c.Thumb == nil) {
		return false
	}
	if !(c.Stickers == nil) {
		return false
	}
	if !(c.Software == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *StickersCreateStickerSetRequest) String() string {
	if c == nil {
		return "StickersCreateStickerSetRequest(nil)"
	}
	type Alias StickersCreateStickerSetRequest
	return fmt.Sprintf("StickersCreateStickerSetRequest%+v", Alias(*c))
}

// FillFrom fills StickersCreateStickerSetRequest from given interface.
func (c *StickersCreateStickerSetRequest) FillFrom(from interface {
	GetMasks() (value bool)
	GetAnimated() (value bool)
	GetVideos() (value bool)
	GetUserID() (value InputUserClass)
	GetTitle() (value string)
	GetShortName() (value string)
	GetThumb() (value InputDocumentClass, ok bool)
	GetStickers() (value []InputStickerSetItem)
	GetSoftware() (value string, ok bool)
}) {
	c.Masks = from.GetMasks()
	c.Animated = from.GetAnimated()
	c.Videos = from.GetVideos()
	c.UserID = from.GetUserID()
	c.Title = from.GetTitle()
	c.ShortName = from.GetShortName()
	if val, ok := from.GetThumb(); ok {
		c.Thumb = val
	}

	c.Stickers = from.GetStickers()
	if val, ok := from.GetSoftware(); ok {
		c.Software = val
	}

}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*StickersCreateStickerSetRequest) TypeID() uint32 {
	return StickersCreateStickerSetRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*StickersCreateStickerSetRequest) TypeName() string {
	return "stickers.createStickerSet"
}

// TypeInfo returns info about TL type.
func (c *StickersCreateStickerSetRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "stickers.createStickerSet",
		ID:   StickersCreateStickerSetRequestTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Masks",
			SchemaName: "masks",
			Null:       !c.Flags.Has(0),
		},
		{
			Name:       "Animated",
			SchemaName: "animated",
			Null:       !c.Flags.Has(1),
		},
		{
			Name:       "Videos",
			SchemaName: "videos",
			Null:       !c.Flags.Has(4),
		},
		{
			Name:       "UserID",
			SchemaName: "user_id",
		},
		{
			Name:       "Title",
			SchemaName: "title",
		},
		{
			Name:       "ShortName",
			SchemaName: "short_name",
		},
		{
			Name:       "Thumb",
			SchemaName: "thumb",
			Null:       !c.Flags.Has(2),
		},
		{
			Name:       "Stickers",
			SchemaName: "stickers",
		},
		{
			Name:       "Software",
			SchemaName: "software",
			Null:       !c.Flags.Has(3),
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (c *StickersCreateStickerSetRequest) SetFlags() {
	if !(c.Masks == false) {
		c.Flags.Set(0)
	}
	if !(c.Animated == false) {
		c.Flags.Set(1)
	}
	if !(c.Videos == false) {
		c.Flags.Set(4)
	}
	if !(c.Thumb == nil) {
		c.Flags.Set(2)
	}
	if !(c.Software == "") {
		c.Flags.Set(3)
	}
}

// Encode implements bin.Encoder.
func (c *StickersCreateStickerSetRequest) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode stickers.createStickerSet#9021ab67 as nil")
	}
	b.PutID(StickersCreateStickerSetRequestTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *StickersCreateStickerSetRequest) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode stickers.createStickerSet#9021ab67 as nil")
	}
	c.SetFlags()
	if err := c.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode stickers.createStickerSet#9021ab67: field flags: %w", err)
	}
	if c.UserID == nil {
		return fmt.Errorf("unable to encode stickers.createStickerSet#9021ab67: field user_id is nil")
	}
	if err := c.UserID.Encode(b); err != nil {
		return fmt.Errorf("unable to encode stickers.createStickerSet#9021ab67: field user_id: %w", err)
	}
	b.PutString(c.Title)
	b.PutString(c.ShortName)
	if c.Flags.Has(2) {
		if c.Thumb == nil {
			return fmt.Errorf("unable to encode stickers.createStickerSet#9021ab67: field thumb is nil")
		}
		if err := c.Thumb.Encode(b); err != nil {
			return fmt.Errorf("unable to encode stickers.createStickerSet#9021ab67: field thumb: %w", err)
		}
	}
	b.PutVectorHeader(len(c.Stickers))
	for idx, v := range c.Stickers {
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode stickers.createStickerSet#9021ab67: field stickers element with index %d: %w", idx, err)
		}
	}
	if c.Flags.Has(3) {
		b.PutString(c.Software)
	}
	return nil
}

// Decode implements bin.Decoder.
func (c *StickersCreateStickerSetRequest) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode stickers.createStickerSet#9021ab67 to nil")
	}
	if err := b.ConsumeID(StickersCreateStickerSetRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode stickers.createStickerSet#9021ab67: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *StickersCreateStickerSetRequest) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode stickers.createStickerSet#9021ab67 to nil")
	}
	{
		if err := c.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode stickers.createStickerSet#9021ab67: field flags: %w", err)
		}
	}
	c.Masks = c.Flags.Has(0)
	c.Animated = c.Flags.Has(1)
	c.Videos = c.Flags.Has(4)
	{
		value, err := DecodeInputUser(b)
		if err != nil {
			return fmt.Errorf("unable to decode stickers.createStickerSet#9021ab67: field user_id: %w", err)
		}
		c.UserID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode stickers.createStickerSet#9021ab67: field title: %w", err)
		}
		c.Title = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode stickers.createStickerSet#9021ab67: field short_name: %w", err)
		}
		c.ShortName = value
	}
	if c.Flags.Has(2) {
		value, err := DecodeInputDocument(b)
		if err != nil {
			return fmt.Errorf("unable to decode stickers.createStickerSet#9021ab67: field thumb: %w", err)
		}
		c.Thumb = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode stickers.createStickerSet#9021ab67: field stickers: %w", err)
		}

		if headerLen > 0 {
			c.Stickers = make([]InputStickerSetItem, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value InputStickerSetItem
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode stickers.createStickerSet#9021ab67: field stickers: %w", err)
			}
			c.Stickers = append(c.Stickers, value)
		}
	}
	if c.Flags.Has(3) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode stickers.createStickerSet#9021ab67: field software: %w", err)
		}
		c.Software = value
	}
	return nil
}

// SetMasks sets value of Masks conditional field.
func (c *StickersCreateStickerSetRequest) SetMasks(value bool) {
	if value {
		c.Flags.Set(0)
		c.Masks = true
	} else {
		c.Flags.Unset(0)
		c.Masks = false
	}
}

// GetMasks returns value of Masks conditional field.
func (c *StickersCreateStickerSetRequest) GetMasks() (value bool) {
	if c == nil {
		return
	}
	return c.Flags.Has(0)
}

// SetAnimated sets value of Animated conditional field.
func (c *StickersCreateStickerSetRequest) SetAnimated(value bool) {
	if value {
		c.Flags.Set(1)
		c.Animated = true
	} else {
		c.Flags.Unset(1)
		c.Animated = false
	}
}

// GetAnimated returns value of Animated conditional field.
func (c *StickersCreateStickerSetRequest) GetAnimated() (value bool) {
	if c == nil {
		return
	}
	return c.Flags.Has(1)
}

// SetVideos sets value of Videos conditional field.
func (c *StickersCreateStickerSetRequest) SetVideos(value bool) {
	if value {
		c.Flags.Set(4)
		c.Videos = true
	} else {
		c.Flags.Unset(4)
		c.Videos = false
	}
}

// GetVideos returns value of Videos conditional field.
func (c *StickersCreateStickerSetRequest) GetVideos() (value bool) {
	if c == nil {
		return
	}
	return c.Flags.Has(4)
}

// GetUserID returns value of UserID field.
func (c *StickersCreateStickerSetRequest) GetUserID() (value InputUserClass) {
	if c == nil {
		return
	}
	return c.UserID
}

// GetTitle returns value of Title field.
func (c *StickersCreateStickerSetRequest) GetTitle() (value string) {
	if c == nil {
		return
	}
	return c.Title
}

// GetShortName returns value of ShortName field.
func (c *StickersCreateStickerSetRequest) GetShortName() (value string) {
	if c == nil {
		return
	}
	return c.ShortName
}

// SetThumb sets value of Thumb conditional field.
func (c *StickersCreateStickerSetRequest) SetThumb(value InputDocumentClass) {
	c.Flags.Set(2)
	c.Thumb = value
}

// GetThumb returns value of Thumb conditional field and
// boolean which is true if field was set.
func (c *StickersCreateStickerSetRequest) GetThumb() (value InputDocumentClass, ok bool) {
	if c == nil {
		return
	}
	if !c.Flags.Has(2) {
		return value, false
	}
	return c.Thumb, true
}

// GetStickers returns value of Stickers field.
func (c *StickersCreateStickerSetRequest) GetStickers() (value []InputStickerSetItem) {
	if c == nil {
		return
	}
	return c.Stickers
}

// SetSoftware sets value of Software conditional field.
func (c *StickersCreateStickerSetRequest) SetSoftware(value string) {
	c.Flags.Set(3)
	c.Software = value
}

// GetSoftware returns value of Software conditional field and
// boolean which is true if field was set.
func (c *StickersCreateStickerSetRequest) GetSoftware() (value string, ok bool) {
	if c == nil {
		return
	}
	if !c.Flags.Has(3) {
		return value, false
	}
	return c.Software, true
}

// GetThumbAsNotEmpty returns mapped value of Thumb conditional field and
// boolean which is true if field was set.
func (c *StickersCreateStickerSetRequest) GetThumbAsNotEmpty() (*InputDocument, bool) {
	if value, ok := c.GetThumb(); ok {
		return value.AsNotEmpty()
	}
	return nil, false
}

// StickersCreateStickerSet invokes method stickers.createStickerSet#9021ab67 returning error if any.
// Create a stickerset, bots only.
//
// Possible errors:
//  400 BOT_MISSING: This method can only be run by a bot.
//  400 PACK_SHORT_NAME_INVALID: Short pack name invalid.
//  400 PACK_SHORT_NAME_OCCUPIED: A stickerpack with this name already exists.
//  400 PACK_TITLE_INVALID: The stickerpack title is invalid.
//  400 PEER_ID_INVALID: The provided peer id is invalid.
//  400 SHORTNAME_OCCUPY_FAILED: An internal error occurred.
//  400 STICKERS_EMPTY: No sticker provided.
//  400 STICKER_EMOJI_INVALID: Sticker emoji invalid.
//  400 STICKER_FILE_INVALID: Sticker file invalid.
//  400 STICKER_PNG_DIMENSIONS: Sticker png dimensions invalid.
//  400 STICKER_PNG_NOPNG: One of the specified stickers is not a valid PNG file.
//  400 STICKER_TGS_NODOC: Incorrect document type for sticker.
//  400 STICKER_TGS_NOTGS: Invalid TGS sticker provided.
//  400 STICKER_THUMB_PNG_NOPNG: Incorrect stickerset thumb file provided, PNG / WEBP expected.
//  400 STICKER_THUMB_TGS_NOTGS: Incorrect stickerset TGS thumb file provided.
//  400 USER_ID_INVALID: The provided user ID is invalid.
//
// See https://core.telegram.org/method/stickers.createStickerSet for reference.
// Can be used by bots.
func (c *Client) StickersCreateStickerSet(ctx context.Context, request *StickersCreateStickerSetRequest) (MessagesStickerSetClass, error) {
	var result MessagesStickerSetBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.StickerSet, nil
}
