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

// PhotoSizeEmpty represents TL type `photoSizeEmpty#e17e23c`.
// Empty constructor. Image with this thumbnail is unavailable.
//
// See https://core.telegram.org/constructor/photoSizeEmpty for reference.
type PhotoSizeEmpty struct {
	// Thumbnail type (see. photoSize¹)
	//
	// Links:
	//  1) https://core.telegram.org/constructor/photoSize
	Type string
}

// PhotoSizeEmptyTypeID is TL type id of PhotoSizeEmpty.
const PhotoSizeEmptyTypeID = 0xe17e23c

// construct implements constructor of PhotoSizeClass.
func (p PhotoSizeEmpty) construct() PhotoSizeClass { return &p }

// Ensuring interfaces in compile-time for PhotoSizeEmpty.
var (
	_ bin.Encoder     = &PhotoSizeEmpty{}
	_ bin.Decoder     = &PhotoSizeEmpty{}
	_ bin.BareEncoder = &PhotoSizeEmpty{}
	_ bin.BareDecoder = &PhotoSizeEmpty{}

	_ PhotoSizeClass = &PhotoSizeEmpty{}
)

func (p *PhotoSizeEmpty) Zero() bool {
	if p == nil {
		return true
	}
	if !(p.Type == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (p *PhotoSizeEmpty) String() string {
	if p == nil {
		return "PhotoSizeEmpty(nil)"
	}
	type Alias PhotoSizeEmpty
	return fmt.Sprintf("PhotoSizeEmpty%+v", Alias(*p))
}

// FillFrom fills PhotoSizeEmpty from given interface.
func (p *PhotoSizeEmpty) FillFrom(from interface {
	GetType() (value string)
}) {
	p.Type = from.GetType()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PhotoSizeEmpty) TypeID() uint32 {
	return PhotoSizeEmptyTypeID
}

// TypeName returns name of type in TL schema.
func (*PhotoSizeEmpty) TypeName() string {
	return "photoSizeEmpty"
}

// TypeInfo returns info about TL type.
func (p *PhotoSizeEmpty) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "photoSizeEmpty",
		ID:   PhotoSizeEmptyTypeID,
	}
	if p == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Type",
			SchemaName: "type",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (p *PhotoSizeEmpty) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode photoSizeEmpty#e17e23c as nil")
	}
	b.PutID(PhotoSizeEmptyTypeID)
	return p.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (p *PhotoSizeEmpty) EncodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode photoSizeEmpty#e17e23c as nil")
	}
	b.PutString(p.Type)
	return nil
}

// Decode implements bin.Decoder.
func (p *PhotoSizeEmpty) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode photoSizeEmpty#e17e23c to nil")
	}
	if err := b.ConsumeID(PhotoSizeEmptyTypeID); err != nil {
		return fmt.Errorf("unable to decode photoSizeEmpty#e17e23c: %w", err)
	}
	return p.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (p *PhotoSizeEmpty) DecodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode photoSizeEmpty#e17e23c to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode photoSizeEmpty#e17e23c: field type: %w", err)
		}
		p.Type = value
	}
	return nil
}

// GetType returns value of Type field.
func (p *PhotoSizeEmpty) GetType() (value string) {
	if p == nil {
		return
	}
	return p.Type
}

// PhotoSize represents TL type `photoSize#75c78e60`.
// Image description.
//
// See https://core.telegram.org/constructor/photoSize for reference.
type PhotoSize struct {
	// Thumbnail type
	Type string
	// Image width
	W int
	// Image height
	H int
	// File size
	Size int
}

// PhotoSizeTypeID is TL type id of PhotoSize.
const PhotoSizeTypeID = 0x75c78e60

// construct implements constructor of PhotoSizeClass.
func (p PhotoSize) construct() PhotoSizeClass { return &p }

// Ensuring interfaces in compile-time for PhotoSize.
var (
	_ bin.Encoder     = &PhotoSize{}
	_ bin.Decoder     = &PhotoSize{}
	_ bin.BareEncoder = &PhotoSize{}
	_ bin.BareDecoder = &PhotoSize{}

	_ PhotoSizeClass = &PhotoSize{}
)

func (p *PhotoSize) Zero() bool {
	if p == nil {
		return true
	}
	if !(p.Type == "") {
		return false
	}
	if !(p.W == 0) {
		return false
	}
	if !(p.H == 0) {
		return false
	}
	if !(p.Size == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (p *PhotoSize) String() string {
	if p == nil {
		return "PhotoSize(nil)"
	}
	type Alias PhotoSize
	return fmt.Sprintf("PhotoSize%+v", Alias(*p))
}

// FillFrom fills PhotoSize from given interface.
func (p *PhotoSize) FillFrom(from interface {
	GetType() (value string)
	GetW() (value int)
	GetH() (value int)
	GetSize() (value int)
}) {
	p.Type = from.GetType()
	p.W = from.GetW()
	p.H = from.GetH()
	p.Size = from.GetSize()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PhotoSize) TypeID() uint32 {
	return PhotoSizeTypeID
}

// TypeName returns name of type in TL schema.
func (*PhotoSize) TypeName() string {
	return "photoSize"
}

// TypeInfo returns info about TL type.
func (p *PhotoSize) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "photoSize",
		ID:   PhotoSizeTypeID,
	}
	if p == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Type",
			SchemaName: "type",
		},
		{
			Name:       "W",
			SchemaName: "w",
		},
		{
			Name:       "H",
			SchemaName: "h",
		},
		{
			Name:       "Size",
			SchemaName: "size",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (p *PhotoSize) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode photoSize#75c78e60 as nil")
	}
	b.PutID(PhotoSizeTypeID)
	return p.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (p *PhotoSize) EncodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode photoSize#75c78e60 as nil")
	}
	b.PutString(p.Type)
	b.PutInt(p.W)
	b.PutInt(p.H)
	b.PutInt(p.Size)
	return nil
}

// Decode implements bin.Decoder.
func (p *PhotoSize) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode photoSize#75c78e60 to nil")
	}
	if err := b.ConsumeID(PhotoSizeTypeID); err != nil {
		return fmt.Errorf("unable to decode photoSize#75c78e60: %w", err)
	}
	return p.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (p *PhotoSize) DecodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode photoSize#75c78e60 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode photoSize#75c78e60: field type: %w", err)
		}
		p.Type = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode photoSize#75c78e60: field w: %w", err)
		}
		p.W = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode photoSize#75c78e60: field h: %w", err)
		}
		p.H = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode photoSize#75c78e60: field size: %w", err)
		}
		p.Size = value
	}
	return nil
}

// GetType returns value of Type field.
func (p *PhotoSize) GetType() (value string) {
	if p == nil {
		return
	}
	return p.Type
}

// GetW returns value of W field.
func (p *PhotoSize) GetW() (value int) {
	if p == nil {
		return
	}
	return p.W
}

// GetH returns value of H field.
func (p *PhotoSize) GetH() (value int) {
	if p == nil {
		return
	}
	return p.H
}

// GetSize returns value of Size field.
func (p *PhotoSize) GetSize() (value int) {
	if p == nil {
		return
	}
	return p.Size
}

// PhotoCachedSize represents TL type `photoCachedSize#21e1ad6`.
// Description of an image and its content.
//
// See https://core.telegram.org/constructor/photoCachedSize for reference.
type PhotoCachedSize struct {
	// Thumbnail type
	Type string
	// Image width
	W int
	// Image height
	H int
	// Binary data, file content
	Bytes []byte
}

// PhotoCachedSizeTypeID is TL type id of PhotoCachedSize.
const PhotoCachedSizeTypeID = 0x21e1ad6

// construct implements constructor of PhotoSizeClass.
func (p PhotoCachedSize) construct() PhotoSizeClass { return &p }

// Ensuring interfaces in compile-time for PhotoCachedSize.
var (
	_ bin.Encoder     = &PhotoCachedSize{}
	_ bin.Decoder     = &PhotoCachedSize{}
	_ bin.BareEncoder = &PhotoCachedSize{}
	_ bin.BareDecoder = &PhotoCachedSize{}

	_ PhotoSizeClass = &PhotoCachedSize{}
)

func (p *PhotoCachedSize) Zero() bool {
	if p == nil {
		return true
	}
	if !(p.Type == "") {
		return false
	}
	if !(p.W == 0) {
		return false
	}
	if !(p.H == 0) {
		return false
	}
	if !(p.Bytes == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (p *PhotoCachedSize) String() string {
	if p == nil {
		return "PhotoCachedSize(nil)"
	}
	type Alias PhotoCachedSize
	return fmt.Sprintf("PhotoCachedSize%+v", Alias(*p))
}

// FillFrom fills PhotoCachedSize from given interface.
func (p *PhotoCachedSize) FillFrom(from interface {
	GetType() (value string)
	GetW() (value int)
	GetH() (value int)
	GetBytes() (value []byte)
}) {
	p.Type = from.GetType()
	p.W = from.GetW()
	p.H = from.GetH()
	p.Bytes = from.GetBytes()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PhotoCachedSize) TypeID() uint32 {
	return PhotoCachedSizeTypeID
}

// TypeName returns name of type in TL schema.
func (*PhotoCachedSize) TypeName() string {
	return "photoCachedSize"
}

// TypeInfo returns info about TL type.
func (p *PhotoCachedSize) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "photoCachedSize",
		ID:   PhotoCachedSizeTypeID,
	}
	if p == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Type",
			SchemaName: "type",
		},
		{
			Name:       "W",
			SchemaName: "w",
		},
		{
			Name:       "H",
			SchemaName: "h",
		},
		{
			Name:       "Bytes",
			SchemaName: "bytes",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (p *PhotoCachedSize) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode photoCachedSize#21e1ad6 as nil")
	}
	b.PutID(PhotoCachedSizeTypeID)
	return p.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (p *PhotoCachedSize) EncodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode photoCachedSize#21e1ad6 as nil")
	}
	b.PutString(p.Type)
	b.PutInt(p.W)
	b.PutInt(p.H)
	b.PutBytes(p.Bytes)
	return nil
}

// Decode implements bin.Decoder.
func (p *PhotoCachedSize) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode photoCachedSize#21e1ad6 to nil")
	}
	if err := b.ConsumeID(PhotoCachedSizeTypeID); err != nil {
		return fmt.Errorf("unable to decode photoCachedSize#21e1ad6: %w", err)
	}
	return p.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (p *PhotoCachedSize) DecodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode photoCachedSize#21e1ad6 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode photoCachedSize#21e1ad6: field type: %w", err)
		}
		p.Type = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode photoCachedSize#21e1ad6: field w: %w", err)
		}
		p.W = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode photoCachedSize#21e1ad6: field h: %w", err)
		}
		p.H = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode photoCachedSize#21e1ad6: field bytes: %w", err)
		}
		p.Bytes = value
	}
	return nil
}

// GetType returns value of Type field.
func (p *PhotoCachedSize) GetType() (value string) {
	if p == nil {
		return
	}
	return p.Type
}

// GetW returns value of W field.
func (p *PhotoCachedSize) GetW() (value int) {
	if p == nil {
		return
	}
	return p.W
}

// GetH returns value of H field.
func (p *PhotoCachedSize) GetH() (value int) {
	if p == nil {
		return
	}
	return p.H
}

// GetBytes returns value of Bytes field.
func (p *PhotoCachedSize) GetBytes() (value []byte) {
	if p == nil {
		return
	}
	return p.Bytes
}

// PhotoStrippedSize represents TL type `photoStrippedSize#e0b0bc2e`.
// A low-resolution compressed JPG payload
//
// See https://core.telegram.org/constructor/photoStrippedSize for reference.
type PhotoStrippedSize struct {
	// Thumbnail type
	Type string
	// Thumbnail data, see here for more info on decompression »¹
	//
	// Links:
	//  1) https://core.telegram.org/api/files#stripped-thumbnails
	Bytes []byte
}

// PhotoStrippedSizeTypeID is TL type id of PhotoStrippedSize.
const PhotoStrippedSizeTypeID = 0xe0b0bc2e

// construct implements constructor of PhotoSizeClass.
func (p PhotoStrippedSize) construct() PhotoSizeClass { return &p }

// Ensuring interfaces in compile-time for PhotoStrippedSize.
var (
	_ bin.Encoder     = &PhotoStrippedSize{}
	_ bin.Decoder     = &PhotoStrippedSize{}
	_ bin.BareEncoder = &PhotoStrippedSize{}
	_ bin.BareDecoder = &PhotoStrippedSize{}

	_ PhotoSizeClass = &PhotoStrippedSize{}
)

func (p *PhotoStrippedSize) Zero() bool {
	if p == nil {
		return true
	}
	if !(p.Type == "") {
		return false
	}
	if !(p.Bytes == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (p *PhotoStrippedSize) String() string {
	if p == nil {
		return "PhotoStrippedSize(nil)"
	}
	type Alias PhotoStrippedSize
	return fmt.Sprintf("PhotoStrippedSize%+v", Alias(*p))
}

// FillFrom fills PhotoStrippedSize from given interface.
func (p *PhotoStrippedSize) FillFrom(from interface {
	GetType() (value string)
	GetBytes() (value []byte)
}) {
	p.Type = from.GetType()
	p.Bytes = from.GetBytes()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PhotoStrippedSize) TypeID() uint32 {
	return PhotoStrippedSizeTypeID
}

// TypeName returns name of type in TL schema.
func (*PhotoStrippedSize) TypeName() string {
	return "photoStrippedSize"
}

// TypeInfo returns info about TL type.
func (p *PhotoStrippedSize) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "photoStrippedSize",
		ID:   PhotoStrippedSizeTypeID,
	}
	if p == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Type",
			SchemaName: "type",
		},
		{
			Name:       "Bytes",
			SchemaName: "bytes",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (p *PhotoStrippedSize) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode photoStrippedSize#e0b0bc2e as nil")
	}
	b.PutID(PhotoStrippedSizeTypeID)
	return p.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (p *PhotoStrippedSize) EncodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode photoStrippedSize#e0b0bc2e as nil")
	}
	b.PutString(p.Type)
	b.PutBytes(p.Bytes)
	return nil
}

// Decode implements bin.Decoder.
func (p *PhotoStrippedSize) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode photoStrippedSize#e0b0bc2e to nil")
	}
	if err := b.ConsumeID(PhotoStrippedSizeTypeID); err != nil {
		return fmt.Errorf("unable to decode photoStrippedSize#e0b0bc2e: %w", err)
	}
	return p.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (p *PhotoStrippedSize) DecodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode photoStrippedSize#e0b0bc2e to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode photoStrippedSize#e0b0bc2e: field type: %w", err)
		}
		p.Type = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode photoStrippedSize#e0b0bc2e: field bytes: %w", err)
		}
		p.Bytes = value
	}
	return nil
}

// GetType returns value of Type field.
func (p *PhotoStrippedSize) GetType() (value string) {
	if p == nil {
		return
	}
	return p.Type
}

// GetBytes returns value of Bytes field.
func (p *PhotoStrippedSize) GetBytes() (value []byte) {
	if p == nil {
		return
	}
	return p.Bytes
}

// PhotoSizeProgressive represents TL type `photoSizeProgressive#fa3efb95`.
// Progressively encoded photosize
//
// See https://core.telegram.org/constructor/photoSizeProgressive for reference.
type PhotoSizeProgressive struct {
	// Photosize type
	Type string
	// Photo width
	W int
	// Photo height
	H int
	// Sizes of progressive JPEG file prefixes, which can be used to preliminarily show the
	// image.
	Sizes []int
}

// PhotoSizeProgressiveTypeID is TL type id of PhotoSizeProgressive.
const PhotoSizeProgressiveTypeID = 0xfa3efb95

// construct implements constructor of PhotoSizeClass.
func (p PhotoSizeProgressive) construct() PhotoSizeClass { return &p }

// Ensuring interfaces in compile-time for PhotoSizeProgressive.
var (
	_ bin.Encoder     = &PhotoSizeProgressive{}
	_ bin.Decoder     = &PhotoSizeProgressive{}
	_ bin.BareEncoder = &PhotoSizeProgressive{}
	_ bin.BareDecoder = &PhotoSizeProgressive{}

	_ PhotoSizeClass = &PhotoSizeProgressive{}
)

func (p *PhotoSizeProgressive) Zero() bool {
	if p == nil {
		return true
	}
	if !(p.Type == "") {
		return false
	}
	if !(p.W == 0) {
		return false
	}
	if !(p.H == 0) {
		return false
	}
	if !(p.Sizes == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (p *PhotoSizeProgressive) String() string {
	if p == nil {
		return "PhotoSizeProgressive(nil)"
	}
	type Alias PhotoSizeProgressive
	return fmt.Sprintf("PhotoSizeProgressive%+v", Alias(*p))
}

// FillFrom fills PhotoSizeProgressive from given interface.
func (p *PhotoSizeProgressive) FillFrom(from interface {
	GetType() (value string)
	GetW() (value int)
	GetH() (value int)
	GetSizes() (value []int)
}) {
	p.Type = from.GetType()
	p.W = from.GetW()
	p.H = from.GetH()
	p.Sizes = from.GetSizes()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PhotoSizeProgressive) TypeID() uint32 {
	return PhotoSizeProgressiveTypeID
}

// TypeName returns name of type in TL schema.
func (*PhotoSizeProgressive) TypeName() string {
	return "photoSizeProgressive"
}

// TypeInfo returns info about TL type.
func (p *PhotoSizeProgressive) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "photoSizeProgressive",
		ID:   PhotoSizeProgressiveTypeID,
	}
	if p == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Type",
			SchemaName: "type",
		},
		{
			Name:       "W",
			SchemaName: "w",
		},
		{
			Name:       "H",
			SchemaName: "h",
		},
		{
			Name:       "Sizes",
			SchemaName: "sizes",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (p *PhotoSizeProgressive) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode photoSizeProgressive#fa3efb95 as nil")
	}
	b.PutID(PhotoSizeProgressiveTypeID)
	return p.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (p *PhotoSizeProgressive) EncodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode photoSizeProgressive#fa3efb95 as nil")
	}
	b.PutString(p.Type)
	b.PutInt(p.W)
	b.PutInt(p.H)
	b.PutVectorHeader(len(p.Sizes))
	for _, v := range p.Sizes {
		b.PutInt(v)
	}
	return nil
}

// Decode implements bin.Decoder.
func (p *PhotoSizeProgressive) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode photoSizeProgressive#fa3efb95 to nil")
	}
	if err := b.ConsumeID(PhotoSizeProgressiveTypeID); err != nil {
		return fmt.Errorf("unable to decode photoSizeProgressive#fa3efb95: %w", err)
	}
	return p.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (p *PhotoSizeProgressive) DecodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode photoSizeProgressive#fa3efb95 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode photoSizeProgressive#fa3efb95: field type: %w", err)
		}
		p.Type = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode photoSizeProgressive#fa3efb95: field w: %w", err)
		}
		p.W = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode photoSizeProgressive#fa3efb95: field h: %w", err)
		}
		p.H = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode photoSizeProgressive#fa3efb95: field sizes: %w", err)
		}

		if headerLen > 0 {
			p.Sizes = make([]int, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Int()
			if err != nil {
				return fmt.Errorf("unable to decode photoSizeProgressive#fa3efb95: field sizes: %w", err)
			}
			p.Sizes = append(p.Sizes, value)
		}
	}
	return nil
}

// GetType returns value of Type field.
func (p *PhotoSizeProgressive) GetType() (value string) {
	if p == nil {
		return
	}
	return p.Type
}

// GetW returns value of W field.
func (p *PhotoSizeProgressive) GetW() (value int) {
	if p == nil {
		return
	}
	return p.W
}

// GetH returns value of H field.
func (p *PhotoSizeProgressive) GetH() (value int) {
	if p == nil {
		return
	}
	return p.H
}

// GetSizes returns value of Sizes field.
func (p *PhotoSizeProgressive) GetSizes() (value []int) {
	if p == nil {
		return
	}
	return p.Sizes
}

// PhotoPathSize represents TL type `photoPathSize#d8214d41`.
// Messages with animated stickers can have a compressed svg (< 300 bytes) to show the
// outline of the sticker before fetching the actual lottie animation.
//
// See https://core.telegram.org/constructor/photoPathSize for reference.
type PhotoPathSize struct {
	// Always j
	Type string
	// Compressed SVG path payload, see here for decompression instructions¹
	//
	// Links:
	//  1) https://core.telegram.org/api/files#vector-thumbnails
	Bytes []byte
}

// PhotoPathSizeTypeID is TL type id of PhotoPathSize.
const PhotoPathSizeTypeID = 0xd8214d41

// construct implements constructor of PhotoSizeClass.
func (p PhotoPathSize) construct() PhotoSizeClass { return &p }

// Ensuring interfaces in compile-time for PhotoPathSize.
var (
	_ bin.Encoder     = &PhotoPathSize{}
	_ bin.Decoder     = &PhotoPathSize{}
	_ bin.BareEncoder = &PhotoPathSize{}
	_ bin.BareDecoder = &PhotoPathSize{}

	_ PhotoSizeClass = &PhotoPathSize{}
)

func (p *PhotoPathSize) Zero() bool {
	if p == nil {
		return true
	}
	if !(p.Type == "") {
		return false
	}
	if !(p.Bytes == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (p *PhotoPathSize) String() string {
	if p == nil {
		return "PhotoPathSize(nil)"
	}
	type Alias PhotoPathSize
	return fmt.Sprintf("PhotoPathSize%+v", Alias(*p))
}

// FillFrom fills PhotoPathSize from given interface.
func (p *PhotoPathSize) FillFrom(from interface {
	GetType() (value string)
	GetBytes() (value []byte)
}) {
	p.Type = from.GetType()
	p.Bytes = from.GetBytes()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PhotoPathSize) TypeID() uint32 {
	return PhotoPathSizeTypeID
}

// TypeName returns name of type in TL schema.
func (*PhotoPathSize) TypeName() string {
	return "photoPathSize"
}

// TypeInfo returns info about TL type.
func (p *PhotoPathSize) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "photoPathSize",
		ID:   PhotoPathSizeTypeID,
	}
	if p == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Type",
			SchemaName: "type",
		},
		{
			Name:       "Bytes",
			SchemaName: "bytes",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (p *PhotoPathSize) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode photoPathSize#d8214d41 as nil")
	}
	b.PutID(PhotoPathSizeTypeID)
	return p.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (p *PhotoPathSize) EncodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode photoPathSize#d8214d41 as nil")
	}
	b.PutString(p.Type)
	b.PutBytes(p.Bytes)
	return nil
}

// Decode implements bin.Decoder.
func (p *PhotoPathSize) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode photoPathSize#d8214d41 to nil")
	}
	if err := b.ConsumeID(PhotoPathSizeTypeID); err != nil {
		return fmt.Errorf("unable to decode photoPathSize#d8214d41: %w", err)
	}
	return p.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (p *PhotoPathSize) DecodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode photoPathSize#d8214d41 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode photoPathSize#d8214d41: field type: %w", err)
		}
		p.Type = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode photoPathSize#d8214d41: field bytes: %w", err)
		}
		p.Bytes = value
	}
	return nil
}

// GetType returns value of Type field.
func (p *PhotoPathSize) GetType() (value string) {
	if p == nil {
		return
	}
	return p.Type
}

// GetBytes returns value of Bytes field.
func (p *PhotoPathSize) GetBytes() (value []byte) {
	if p == nil {
		return
	}
	return p.Bytes
}

// PhotoSizeClassName is schema name of PhotoSizeClass.
const PhotoSizeClassName = "PhotoSize"

// PhotoSizeClass represents PhotoSize generic type.
//
// See https://core.telegram.org/type/PhotoSize for reference.
//
// Example:
//  g, err := tg.DecodePhotoSize(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *tg.PhotoSizeEmpty: // photoSizeEmpty#e17e23c
//  case *tg.PhotoSize: // photoSize#75c78e60
//  case *tg.PhotoCachedSize: // photoCachedSize#21e1ad6
//  case *tg.PhotoStrippedSize: // photoStrippedSize#e0b0bc2e
//  case *tg.PhotoSizeProgressive: // photoSizeProgressive#fa3efb95
//  case *tg.PhotoPathSize: // photoPathSize#d8214d41
//  default: panic(v)
//  }
type PhotoSizeClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() PhotoSizeClass

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

	// Thumbnail type (see. photoSize¹)
	//
	// Links:
	//  1) https://core.telegram.org/constructor/photoSize
	GetType() (value string)

	// AsNotEmpty tries to map PhotoSizeClass to NotEmptyPhotoSize.
	AsNotEmpty() (NotEmptyPhotoSize, bool)
}

// NotEmptyPhotoSize represents NotEmpty subset of PhotoSizeClass.
type NotEmptyPhotoSize interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() PhotoSizeClass

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

	// Thumbnail type
	GetType() (value string)
}

// AsNotEmpty tries to map PhotoSizeEmpty to NotEmptyPhotoSize.
func (p *PhotoSizeEmpty) AsNotEmpty() (NotEmptyPhotoSize, bool) {
	value, ok := (PhotoSizeClass(p)).(NotEmptyPhotoSize)
	return value, ok
}

// AsNotEmpty tries to map PhotoSize to NotEmptyPhotoSize.
func (p *PhotoSize) AsNotEmpty() (NotEmptyPhotoSize, bool) {
	value, ok := (PhotoSizeClass(p)).(NotEmptyPhotoSize)
	return value, ok
}

// AsNotEmpty tries to map PhotoCachedSize to NotEmptyPhotoSize.
func (p *PhotoCachedSize) AsNotEmpty() (NotEmptyPhotoSize, bool) {
	value, ok := (PhotoSizeClass(p)).(NotEmptyPhotoSize)
	return value, ok
}

// AsNotEmpty tries to map PhotoStrippedSize to NotEmptyPhotoSize.
func (p *PhotoStrippedSize) AsNotEmpty() (NotEmptyPhotoSize, bool) {
	value, ok := (PhotoSizeClass(p)).(NotEmptyPhotoSize)
	return value, ok
}

// AsNotEmpty tries to map PhotoSizeProgressive to NotEmptyPhotoSize.
func (p *PhotoSizeProgressive) AsNotEmpty() (NotEmptyPhotoSize, bool) {
	value, ok := (PhotoSizeClass(p)).(NotEmptyPhotoSize)
	return value, ok
}

// AsNotEmpty tries to map PhotoPathSize to NotEmptyPhotoSize.
func (p *PhotoPathSize) AsNotEmpty() (NotEmptyPhotoSize, bool) {
	value, ok := (PhotoSizeClass(p)).(NotEmptyPhotoSize)
	return value, ok
}

// DecodePhotoSize implements binary de-serialization for PhotoSizeClass.
func DecodePhotoSize(buf *bin.Buffer) (PhotoSizeClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case PhotoSizeEmptyTypeID:
		// Decoding photoSizeEmpty#e17e23c.
		v := PhotoSizeEmpty{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode PhotoSizeClass: %w", err)
		}
		return &v, nil
	case PhotoSizeTypeID:
		// Decoding photoSize#75c78e60.
		v := PhotoSize{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode PhotoSizeClass: %w", err)
		}
		return &v, nil
	case PhotoCachedSizeTypeID:
		// Decoding photoCachedSize#21e1ad6.
		v := PhotoCachedSize{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode PhotoSizeClass: %w", err)
		}
		return &v, nil
	case PhotoStrippedSizeTypeID:
		// Decoding photoStrippedSize#e0b0bc2e.
		v := PhotoStrippedSize{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode PhotoSizeClass: %w", err)
		}
		return &v, nil
	case PhotoSizeProgressiveTypeID:
		// Decoding photoSizeProgressive#fa3efb95.
		v := PhotoSizeProgressive{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode PhotoSizeClass: %w", err)
		}
		return &v, nil
	case PhotoPathSizeTypeID:
		// Decoding photoPathSize#d8214d41.
		v := PhotoPathSize{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode PhotoSizeClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode PhotoSizeClass: %w", bin.NewUnexpectedID(id))
	}
}

// PhotoSize boxes the PhotoSizeClass providing a helper.
type PhotoSizeBox struct {
	PhotoSize PhotoSizeClass
}

// Decode implements bin.Decoder for PhotoSizeBox.
func (b *PhotoSizeBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode PhotoSizeBox to nil")
	}
	v, err := DecodePhotoSize(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.PhotoSize = v
	return nil
}

// Encode implements bin.Encode for PhotoSizeBox.
func (b *PhotoSizeBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.PhotoSize == nil {
		return fmt.Errorf("unable to encode PhotoSizeClass as nil")
	}
	return b.PhotoSize.Encode(buf)
}
