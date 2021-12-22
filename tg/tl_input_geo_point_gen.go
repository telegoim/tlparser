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

// InputGeoPointEmpty represents TL type `inputGeoPointEmpty#e4c123d6`.
// Empty GeoPoint constructor.
//
// See https://core.telegram.org/constructor/inputGeoPointEmpty for reference.
type InputGeoPointEmpty struct {
}

// InputGeoPointEmptyTypeID is TL type id of InputGeoPointEmpty.
const InputGeoPointEmptyTypeID = 0xe4c123d6

// construct implements constructor of InputGeoPointClass.
func (i InputGeoPointEmpty) construct() InputGeoPointClass { return &i }

// Ensuring interfaces in compile-time for InputGeoPointEmpty.
var (
	_ bin.Encoder     = &InputGeoPointEmpty{}
	_ bin.Decoder     = &InputGeoPointEmpty{}
	_ bin.BareEncoder = &InputGeoPointEmpty{}
	_ bin.BareDecoder = &InputGeoPointEmpty{}

	_ InputGeoPointClass = &InputGeoPointEmpty{}
)

func (i *InputGeoPointEmpty) Zero() bool {
	if i == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputGeoPointEmpty) String() string {
	if i == nil {
		return "InputGeoPointEmpty(nil)"
	}
	type Alias InputGeoPointEmpty
	return fmt.Sprintf("InputGeoPointEmpty%+v", Alias(*i))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InputGeoPointEmpty) TypeID() uint32 {
	return InputGeoPointEmptyTypeID
}

// TypeName returns name of type in TL schema.
func (*InputGeoPointEmpty) TypeName() string {
	return "inputGeoPointEmpty"
}

// TypeInfo returns info about TL type.
func (i *InputGeoPointEmpty) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inputGeoPointEmpty",
		ID:   InputGeoPointEmptyTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (i *InputGeoPointEmpty) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputGeoPointEmpty#e4c123d6 as nil")
	}
	b.PutID(InputGeoPointEmptyTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InputGeoPointEmpty) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputGeoPointEmpty#e4c123d6 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (i *InputGeoPointEmpty) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputGeoPointEmpty#e4c123d6 to nil")
	}
	if err := b.ConsumeID(InputGeoPointEmptyTypeID); err != nil {
		return fmt.Errorf("unable to decode inputGeoPointEmpty#e4c123d6: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InputGeoPointEmpty) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputGeoPointEmpty#e4c123d6 to nil")
	}
	return nil
}

// InputGeoPoint represents TL type `inputGeoPoint#48222faf`.
// Defines a GeoPoint by its coordinates.
//
// See https://core.telegram.org/constructor/inputGeoPoint for reference.
type InputGeoPoint struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Latitide
	Lat float64
	// Longtitude
	Long float64
	// The estimated horizontal accuracy of the location, in meters; as defined by the sender.
	//
	// Use SetAccuracyRadius and GetAccuracyRadius helpers.
	AccuracyRadius int
}

// InputGeoPointTypeID is TL type id of InputGeoPoint.
const InputGeoPointTypeID = 0x48222faf

// construct implements constructor of InputGeoPointClass.
func (i InputGeoPoint) construct() InputGeoPointClass { return &i }

// Ensuring interfaces in compile-time for InputGeoPoint.
var (
	_ bin.Encoder     = &InputGeoPoint{}
	_ bin.Decoder     = &InputGeoPoint{}
	_ bin.BareEncoder = &InputGeoPoint{}
	_ bin.BareDecoder = &InputGeoPoint{}

	_ InputGeoPointClass = &InputGeoPoint{}
)

func (i *InputGeoPoint) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.Flags.Zero()) {
		return false
	}
	if !(i.Lat == 0) {
		return false
	}
	if !(i.Long == 0) {
		return false
	}
	if !(i.AccuracyRadius == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputGeoPoint) String() string {
	if i == nil {
		return "InputGeoPoint(nil)"
	}
	type Alias InputGeoPoint
	return fmt.Sprintf("InputGeoPoint%+v", Alias(*i))
}

// FillFrom fills InputGeoPoint from given interface.
func (i *InputGeoPoint) FillFrom(from interface {
	GetLat() (value float64)
	GetLong() (value float64)
	GetAccuracyRadius() (value int, ok bool)
}) {
	i.Lat = from.GetLat()
	i.Long = from.GetLong()
	if val, ok := from.GetAccuracyRadius(); ok {
		i.AccuracyRadius = val
	}

}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InputGeoPoint) TypeID() uint32 {
	return InputGeoPointTypeID
}

// TypeName returns name of type in TL schema.
func (*InputGeoPoint) TypeName() string {
	return "inputGeoPoint"
}

// TypeInfo returns info about TL type.
func (i *InputGeoPoint) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inputGeoPoint",
		ID:   InputGeoPointTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Lat",
			SchemaName: "lat",
		},
		{
			Name:       "Long",
			SchemaName: "long",
		},
		{
			Name:       "AccuracyRadius",
			SchemaName: "accuracy_radius",
			Null:       !i.Flags.Has(0),
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (i *InputGeoPoint) SetFlags() {
	if !(i.AccuracyRadius == 0) {
		i.Flags.Set(0)
	}
}

// Encode implements bin.Encoder.
func (i *InputGeoPoint) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputGeoPoint#48222faf as nil")
	}
	b.PutID(InputGeoPointTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InputGeoPoint) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputGeoPoint#48222faf as nil")
	}
	i.SetFlags()
	if err := i.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode inputGeoPoint#48222faf: field flags: %w", err)
	}
	b.PutDouble(i.Lat)
	b.PutDouble(i.Long)
	if i.Flags.Has(0) {
		b.PutInt(i.AccuracyRadius)
	}
	return nil
}

// Decode implements bin.Decoder.
func (i *InputGeoPoint) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputGeoPoint#48222faf to nil")
	}
	if err := b.ConsumeID(InputGeoPointTypeID); err != nil {
		return fmt.Errorf("unable to decode inputGeoPoint#48222faf: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InputGeoPoint) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputGeoPoint#48222faf to nil")
	}
	{
		if err := i.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode inputGeoPoint#48222faf: field flags: %w", err)
		}
	}
	{
		value, err := b.Double()
		if err != nil {
			return fmt.Errorf("unable to decode inputGeoPoint#48222faf: field lat: %w", err)
		}
		i.Lat = value
	}
	{
		value, err := b.Double()
		if err != nil {
			return fmt.Errorf("unable to decode inputGeoPoint#48222faf: field long: %w", err)
		}
		i.Long = value
	}
	if i.Flags.Has(0) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode inputGeoPoint#48222faf: field accuracy_radius: %w", err)
		}
		i.AccuracyRadius = value
	}
	return nil
}

// GetLat returns value of Lat field.
func (i *InputGeoPoint) GetLat() (value float64) {
	if i == nil {
		return
	}
	return i.Lat
}

// GetLong returns value of Long field.
func (i *InputGeoPoint) GetLong() (value float64) {
	if i == nil {
		return
	}
	return i.Long
}

// SetAccuracyRadius sets value of AccuracyRadius conditional field.
func (i *InputGeoPoint) SetAccuracyRadius(value int) {
	i.Flags.Set(0)
	i.AccuracyRadius = value
}

// GetAccuracyRadius returns value of AccuracyRadius conditional field and
// boolean which is true if field was set.
func (i *InputGeoPoint) GetAccuracyRadius() (value int, ok bool) {
	if i == nil {
		return
	}
	if !i.Flags.Has(0) {
		return value, false
	}
	return i.AccuracyRadius, true
}

// InputGeoPointClassName is schema name of InputGeoPointClass.
const InputGeoPointClassName = "InputGeoPoint"

// InputGeoPointClass represents InputGeoPoint generic type.
//
// See https://core.telegram.org/type/InputGeoPoint for reference.
//
// Example:
//  g, err := tg.DecodeInputGeoPoint(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *tg.InputGeoPointEmpty: // inputGeoPointEmpty#e4c123d6
//  case *tg.InputGeoPoint: // inputGeoPoint#48222faf
//  default: panic(v)
//  }
type InputGeoPointClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() InputGeoPointClass

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

	// AsNotEmpty tries to map InputGeoPointClass to InputGeoPoint.
	AsNotEmpty() (*InputGeoPoint, bool)
}

// AsNotEmpty tries to map InputGeoPointEmpty to InputGeoPoint.
func (i *InputGeoPointEmpty) AsNotEmpty() (*InputGeoPoint, bool) {
	return nil, false
}

// AsNotEmpty tries to map InputGeoPoint to InputGeoPoint.
func (i *InputGeoPoint) AsNotEmpty() (*InputGeoPoint, bool) {
	return i, true
}

// DecodeInputGeoPoint implements binary de-serialization for InputGeoPointClass.
func DecodeInputGeoPoint(buf *bin.Buffer) (InputGeoPointClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case InputGeoPointEmptyTypeID:
		// Decoding inputGeoPointEmpty#e4c123d6.
		v := InputGeoPointEmpty{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputGeoPointClass: %w", err)
		}
		return &v, nil
	case InputGeoPointTypeID:
		// Decoding inputGeoPoint#48222faf.
		v := InputGeoPoint{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputGeoPointClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode InputGeoPointClass: %w", bin.NewUnexpectedID(id))
	}
}

// InputGeoPoint boxes the InputGeoPointClass providing a helper.
type InputGeoPointBox struct {
	InputGeoPoint InputGeoPointClass
}

// Decode implements bin.Decoder for InputGeoPointBox.
func (b *InputGeoPointBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode InputGeoPointBox to nil")
	}
	v, err := DecodeInputGeoPoint(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.InputGeoPoint = v
	return nil
}

// Encode implements bin.Encode for InputGeoPointBox.
func (b *InputGeoPointBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.InputGeoPoint == nil {
		return fmt.Errorf("unable to encode InputGeoPointClass as nil")
	}
	return b.InputGeoPoint.Encode(buf)
}
