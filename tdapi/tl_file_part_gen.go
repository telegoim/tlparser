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

// FilePart represents TL type `filePart#36594c36`.
type FilePart struct {
	// File bytes
	Data []byte
}

// FilePartTypeID is TL type id of FilePart.
const FilePartTypeID = 0x36594c36

// Ensuring interfaces in compile-time for FilePart.
var (
	_ bin.Encoder     = &FilePart{}
	_ bin.Decoder     = &FilePart{}
	_ bin.BareEncoder = &FilePart{}
	_ bin.BareDecoder = &FilePart{}
)

func (f *FilePart) Zero() bool {
	if f == nil {
		return true
	}
	if !(f.Data == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (f *FilePart) String() string {
	if f == nil {
		return "FilePart(nil)"
	}
	type Alias FilePart
	return fmt.Sprintf("FilePart%+v", Alias(*f))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*FilePart) TypeID() uint32 {
	return FilePartTypeID
}

// TypeName returns name of type in TL schema.
func (*FilePart) TypeName() string {
	return "filePart"
}

// TypeInfo returns info about TL type.
func (f *FilePart) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "filePart",
		ID:   FilePartTypeID,
	}
	if f == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Data",
			SchemaName: "data",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (f *FilePart) Encode(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't encode filePart#36594c36 as nil")
	}
	b.PutID(FilePartTypeID)
	return f.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (f *FilePart) EncodeBare(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't encode filePart#36594c36 as nil")
	}
	b.PutBytes(f.Data)
	return nil
}

// Decode implements bin.Decoder.
func (f *FilePart) Decode(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't decode filePart#36594c36 to nil")
	}
	if err := b.ConsumeID(FilePartTypeID); err != nil {
		return fmt.Errorf("unable to decode filePart#36594c36: %w", err)
	}
	return f.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (f *FilePart) DecodeBare(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't decode filePart#36594c36 to nil")
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode filePart#36594c36: field data: %w", err)
		}
		f.Data = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (f *FilePart) EncodeTDLibJSON(b tdjson.Encoder) error {
	if f == nil {
		return fmt.Errorf("can't encode filePart#36594c36 as nil")
	}
	b.ObjStart()
	b.PutID("filePart")
	b.FieldStart("data")
	b.PutBytes(f.Data)
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (f *FilePart) DecodeTDLibJSON(b tdjson.Decoder) error {
	if f == nil {
		return fmt.Errorf("can't decode filePart#36594c36 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("filePart"); err != nil {
				return fmt.Errorf("unable to decode filePart#36594c36: %w", err)
			}
		case "data":
			value, err := b.Bytes()
			if err != nil {
				return fmt.Errorf("unable to decode filePart#36594c36: field data: %w", err)
			}
			f.Data = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetData returns value of Data field.
func (f *FilePart) GetData() (value []byte) {
	if f == nil {
		return
	}
	return f.Data
}
