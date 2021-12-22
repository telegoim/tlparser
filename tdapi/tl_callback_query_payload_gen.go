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

// CallbackQueryPayloadData represents TL type `callbackQueryPayloadData#8a1e3c66`.
type CallbackQueryPayloadData struct {
	// Data that was attached to the callback button
	Data []byte
}

// CallbackQueryPayloadDataTypeID is TL type id of CallbackQueryPayloadData.
const CallbackQueryPayloadDataTypeID = 0x8a1e3c66

// construct implements constructor of CallbackQueryPayloadClass.
func (c CallbackQueryPayloadData) construct() CallbackQueryPayloadClass { return &c }

// Ensuring interfaces in compile-time for CallbackQueryPayloadData.
var (
	_ bin.Encoder     = &CallbackQueryPayloadData{}
	_ bin.Decoder     = &CallbackQueryPayloadData{}
	_ bin.BareEncoder = &CallbackQueryPayloadData{}
	_ bin.BareDecoder = &CallbackQueryPayloadData{}

	_ CallbackQueryPayloadClass = &CallbackQueryPayloadData{}
)

func (c *CallbackQueryPayloadData) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.Data == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *CallbackQueryPayloadData) String() string {
	if c == nil {
		return "CallbackQueryPayloadData(nil)"
	}
	type Alias CallbackQueryPayloadData
	return fmt.Sprintf("CallbackQueryPayloadData%+v", Alias(*c))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*CallbackQueryPayloadData) TypeID() uint32 {
	return CallbackQueryPayloadDataTypeID
}

// TypeName returns name of type in TL schema.
func (*CallbackQueryPayloadData) TypeName() string {
	return "callbackQueryPayloadData"
}

// TypeInfo returns info about TL type.
func (c *CallbackQueryPayloadData) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "callbackQueryPayloadData",
		ID:   CallbackQueryPayloadDataTypeID,
	}
	if c == nil {
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
func (c *CallbackQueryPayloadData) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode callbackQueryPayloadData#8a1e3c66 as nil")
	}
	b.PutID(CallbackQueryPayloadDataTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *CallbackQueryPayloadData) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode callbackQueryPayloadData#8a1e3c66 as nil")
	}
	b.PutBytes(c.Data)
	return nil
}

// Decode implements bin.Decoder.
func (c *CallbackQueryPayloadData) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode callbackQueryPayloadData#8a1e3c66 to nil")
	}
	if err := b.ConsumeID(CallbackQueryPayloadDataTypeID); err != nil {
		return fmt.Errorf("unable to decode callbackQueryPayloadData#8a1e3c66: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *CallbackQueryPayloadData) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode callbackQueryPayloadData#8a1e3c66 to nil")
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode callbackQueryPayloadData#8a1e3c66: field data: %w", err)
		}
		c.Data = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (c *CallbackQueryPayloadData) EncodeTDLibJSON(b tdjson.Encoder) error {
	if c == nil {
		return fmt.Errorf("can't encode callbackQueryPayloadData#8a1e3c66 as nil")
	}
	b.ObjStart()
	b.PutID("callbackQueryPayloadData")
	b.FieldStart("data")
	b.PutBytes(c.Data)
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (c *CallbackQueryPayloadData) DecodeTDLibJSON(b tdjson.Decoder) error {
	if c == nil {
		return fmt.Errorf("can't decode callbackQueryPayloadData#8a1e3c66 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("callbackQueryPayloadData"); err != nil {
				return fmt.Errorf("unable to decode callbackQueryPayloadData#8a1e3c66: %w", err)
			}
		case "data":
			value, err := b.Bytes()
			if err != nil {
				return fmt.Errorf("unable to decode callbackQueryPayloadData#8a1e3c66: field data: %w", err)
			}
			c.Data = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetData returns value of Data field.
func (c *CallbackQueryPayloadData) GetData() (value []byte) {
	if c == nil {
		return
	}
	return c.Data
}

// CallbackQueryPayloadDataWithPassword represents TL type `callbackQueryPayloadDataWithPassword#4fe2d8f2`.
type CallbackQueryPayloadDataWithPassword struct {
	// The password for the current user
	Password string
	// Data that was attached to the callback button
	Data []byte
}

// CallbackQueryPayloadDataWithPasswordTypeID is TL type id of CallbackQueryPayloadDataWithPassword.
const CallbackQueryPayloadDataWithPasswordTypeID = 0x4fe2d8f2

// construct implements constructor of CallbackQueryPayloadClass.
func (c CallbackQueryPayloadDataWithPassword) construct() CallbackQueryPayloadClass { return &c }

// Ensuring interfaces in compile-time for CallbackQueryPayloadDataWithPassword.
var (
	_ bin.Encoder     = &CallbackQueryPayloadDataWithPassword{}
	_ bin.Decoder     = &CallbackQueryPayloadDataWithPassword{}
	_ bin.BareEncoder = &CallbackQueryPayloadDataWithPassword{}
	_ bin.BareDecoder = &CallbackQueryPayloadDataWithPassword{}

	_ CallbackQueryPayloadClass = &CallbackQueryPayloadDataWithPassword{}
)

func (c *CallbackQueryPayloadDataWithPassword) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.Password == "") {
		return false
	}
	if !(c.Data == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *CallbackQueryPayloadDataWithPassword) String() string {
	if c == nil {
		return "CallbackQueryPayloadDataWithPassword(nil)"
	}
	type Alias CallbackQueryPayloadDataWithPassword
	return fmt.Sprintf("CallbackQueryPayloadDataWithPassword%+v", Alias(*c))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*CallbackQueryPayloadDataWithPassword) TypeID() uint32 {
	return CallbackQueryPayloadDataWithPasswordTypeID
}

// TypeName returns name of type in TL schema.
func (*CallbackQueryPayloadDataWithPassword) TypeName() string {
	return "callbackQueryPayloadDataWithPassword"
}

// TypeInfo returns info about TL type.
func (c *CallbackQueryPayloadDataWithPassword) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "callbackQueryPayloadDataWithPassword",
		ID:   CallbackQueryPayloadDataWithPasswordTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Password",
			SchemaName: "password",
		},
		{
			Name:       "Data",
			SchemaName: "data",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *CallbackQueryPayloadDataWithPassword) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode callbackQueryPayloadDataWithPassword#4fe2d8f2 as nil")
	}
	b.PutID(CallbackQueryPayloadDataWithPasswordTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *CallbackQueryPayloadDataWithPassword) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode callbackQueryPayloadDataWithPassword#4fe2d8f2 as nil")
	}
	b.PutString(c.Password)
	b.PutBytes(c.Data)
	return nil
}

// Decode implements bin.Decoder.
func (c *CallbackQueryPayloadDataWithPassword) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode callbackQueryPayloadDataWithPassword#4fe2d8f2 to nil")
	}
	if err := b.ConsumeID(CallbackQueryPayloadDataWithPasswordTypeID); err != nil {
		return fmt.Errorf("unable to decode callbackQueryPayloadDataWithPassword#4fe2d8f2: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *CallbackQueryPayloadDataWithPassword) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode callbackQueryPayloadDataWithPassword#4fe2d8f2 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode callbackQueryPayloadDataWithPassword#4fe2d8f2: field password: %w", err)
		}
		c.Password = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode callbackQueryPayloadDataWithPassword#4fe2d8f2: field data: %w", err)
		}
		c.Data = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (c *CallbackQueryPayloadDataWithPassword) EncodeTDLibJSON(b tdjson.Encoder) error {
	if c == nil {
		return fmt.Errorf("can't encode callbackQueryPayloadDataWithPassword#4fe2d8f2 as nil")
	}
	b.ObjStart()
	b.PutID("callbackQueryPayloadDataWithPassword")
	b.FieldStart("password")
	b.PutString(c.Password)
	b.FieldStart("data")
	b.PutBytes(c.Data)
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (c *CallbackQueryPayloadDataWithPassword) DecodeTDLibJSON(b tdjson.Decoder) error {
	if c == nil {
		return fmt.Errorf("can't decode callbackQueryPayloadDataWithPassword#4fe2d8f2 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("callbackQueryPayloadDataWithPassword"); err != nil {
				return fmt.Errorf("unable to decode callbackQueryPayloadDataWithPassword#4fe2d8f2: %w", err)
			}
		case "password":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode callbackQueryPayloadDataWithPassword#4fe2d8f2: field password: %w", err)
			}
			c.Password = value
		case "data":
			value, err := b.Bytes()
			if err != nil {
				return fmt.Errorf("unable to decode callbackQueryPayloadDataWithPassword#4fe2d8f2: field data: %w", err)
			}
			c.Data = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetPassword returns value of Password field.
func (c *CallbackQueryPayloadDataWithPassword) GetPassword() (value string) {
	if c == nil {
		return
	}
	return c.Password
}

// GetData returns value of Data field.
func (c *CallbackQueryPayloadDataWithPassword) GetData() (value []byte) {
	if c == nil {
		return
	}
	return c.Data
}

// CallbackQueryPayloadGame represents TL type `callbackQueryPayloadGame#4db2ec38`.
type CallbackQueryPayloadGame struct {
	// A short name of the game that was attached to the callback button
	GameShortName string
}

// CallbackQueryPayloadGameTypeID is TL type id of CallbackQueryPayloadGame.
const CallbackQueryPayloadGameTypeID = 0x4db2ec38

// construct implements constructor of CallbackQueryPayloadClass.
func (c CallbackQueryPayloadGame) construct() CallbackQueryPayloadClass { return &c }

// Ensuring interfaces in compile-time for CallbackQueryPayloadGame.
var (
	_ bin.Encoder     = &CallbackQueryPayloadGame{}
	_ bin.Decoder     = &CallbackQueryPayloadGame{}
	_ bin.BareEncoder = &CallbackQueryPayloadGame{}
	_ bin.BareDecoder = &CallbackQueryPayloadGame{}

	_ CallbackQueryPayloadClass = &CallbackQueryPayloadGame{}
)

func (c *CallbackQueryPayloadGame) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.GameShortName == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *CallbackQueryPayloadGame) String() string {
	if c == nil {
		return "CallbackQueryPayloadGame(nil)"
	}
	type Alias CallbackQueryPayloadGame
	return fmt.Sprintf("CallbackQueryPayloadGame%+v", Alias(*c))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*CallbackQueryPayloadGame) TypeID() uint32 {
	return CallbackQueryPayloadGameTypeID
}

// TypeName returns name of type in TL schema.
func (*CallbackQueryPayloadGame) TypeName() string {
	return "callbackQueryPayloadGame"
}

// TypeInfo returns info about TL type.
func (c *CallbackQueryPayloadGame) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "callbackQueryPayloadGame",
		ID:   CallbackQueryPayloadGameTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "GameShortName",
			SchemaName: "game_short_name",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *CallbackQueryPayloadGame) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode callbackQueryPayloadGame#4db2ec38 as nil")
	}
	b.PutID(CallbackQueryPayloadGameTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *CallbackQueryPayloadGame) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode callbackQueryPayloadGame#4db2ec38 as nil")
	}
	b.PutString(c.GameShortName)
	return nil
}

// Decode implements bin.Decoder.
func (c *CallbackQueryPayloadGame) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode callbackQueryPayloadGame#4db2ec38 to nil")
	}
	if err := b.ConsumeID(CallbackQueryPayloadGameTypeID); err != nil {
		return fmt.Errorf("unable to decode callbackQueryPayloadGame#4db2ec38: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *CallbackQueryPayloadGame) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode callbackQueryPayloadGame#4db2ec38 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode callbackQueryPayloadGame#4db2ec38: field game_short_name: %w", err)
		}
		c.GameShortName = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (c *CallbackQueryPayloadGame) EncodeTDLibJSON(b tdjson.Encoder) error {
	if c == nil {
		return fmt.Errorf("can't encode callbackQueryPayloadGame#4db2ec38 as nil")
	}
	b.ObjStart()
	b.PutID("callbackQueryPayloadGame")
	b.FieldStart("game_short_name")
	b.PutString(c.GameShortName)
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (c *CallbackQueryPayloadGame) DecodeTDLibJSON(b tdjson.Decoder) error {
	if c == nil {
		return fmt.Errorf("can't decode callbackQueryPayloadGame#4db2ec38 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("callbackQueryPayloadGame"); err != nil {
				return fmt.Errorf("unable to decode callbackQueryPayloadGame#4db2ec38: %w", err)
			}
		case "game_short_name":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode callbackQueryPayloadGame#4db2ec38: field game_short_name: %w", err)
			}
			c.GameShortName = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetGameShortName returns value of GameShortName field.
func (c *CallbackQueryPayloadGame) GetGameShortName() (value string) {
	if c == nil {
		return
	}
	return c.GameShortName
}

// CallbackQueryPayloadClassName is schema name of CallbackQueryPayloadClass.
const CallbackQueryPayloadClassName = "CallbackQueryPayload"

// CallbackQueryPayloadClass represents CallbackQueryPayload generic type.
//
// Example:
//  g, err := tdapi.DecodeCallbackQueryPayload(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *tdapi.CallbackQueryPayloadData: // callbackQueryPayloadData#8a1e3c66
//  case *tdapi.CallbackQueryPayloadDataWithPassword: // callbackQueryPayloadDataWithPassword#4fe2d8f2
//  case *tdapi.CallbackQueryPayloadGame: // callbackQueryPayloadGame#4db2ec38
//  default: panic(v)
//  }
type CallbackQueryPayloadClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() CallbackQueryPayloadClass

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

	EncodeTDLibJSON(b tdjson.Encoder) error
	DecodeTDLibJSON(b tdjson.Decoder) error
}

// DecodeCallbackQueryPayload implements binary de-serialization for CallbackQueryPayloadClass.
func DecodeCallbackQueryPayload(buf *bin.Buffer) (CallbackQueryPayloadClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case CallbackQueryPayloadDataTypeID:
		// Decoding callbackQueryPayloadData#8a1e3c66.
		v := CallbackQueryPayloadData{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode CallbackQueryPayloadClass: %w", err)
		}
		return &v, nil
	case CallbackQueryPayloadDataWithPasswordTypeID:
		// Decoding callbackQueryPayloadDataWithPassword#4fe2d8f2.
		v := CallbackQueryPayloadDataWithPassword{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode CallbackQueryPayloadClass: %w", err)
		}
		return &v, nil
	case CallbackQueryPayloadGameTypeID:
		// Decoding callbackQueryPayloadGame#4db2ec38.
		v := CallbackQueryPayloadGame{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode CallbackQueryPayloadClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode CallbackQueryPayloadClass: %w", bin.NewUnexpectedID(id))
	}
}

// DecodeTDLibJSONCallbackQueryPayload implements binary de-serialization for CallbackQueryPayloadClass.
func DecodeTDLibJSONCallbackQueryPayload(buf tdjson.Decoder) (CallbackQueryPayloadClass, error) {
	id, err := buf.FindTypeID()
	if err != nil {
		return nil, err
	}
	switch id {
	case "callbackQueryPayloadData":
		// Decoding callbackQueryPayloadData#8a1e3c66.
		v := CallbackQueryPayloadData{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode CallbackQueryPayloadClass: %w", err)
		}
		return &v, nil
	case "callbackQueryPayloadDataWithPassword":
		// Decoding callbackQueryPayloadDataWithPassword#4fe2d8f2.
		v := CallbackQueryPayloadDataWithPassword{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode CallbackQueryPayloadClass: %w", err)
		}
		return &v, nil
	case "callbackQueryPayloadGame":
		// Decoding callbackQueryPayloadGame#4db2ec38.
		v := CallbackQueryPayloadGame{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode CallbackQueryPayloadClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode CallbackQueryPayloadClass: %w", tdjson.NewUnexpectedID(id))
	}
}

// CallbackQueryPayload boxes the CallbackQueryPayloadClass providing a helper.
type CallbackQueryPayloadBox struct {
	CallbackQueryPayload CallbackQueryPayloadClass
}

// Decode implements bin.Decoder for CallbackQueryPayloadBox.
func (b *CallbackQueryPayloadBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode CallbackQueryPayloadBox to nil")
	}
	v, err := DecodeCallbackQueryPayload(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.CallbackQueryPayload = v
	return nil
}

// Encode implements bin.Encode for CallbackQueryPayloadBox.
func (b *CallbackQueryPayloadBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.CallbackQueryPayload == nil {
		return fmt.Errorf("unable to encode CallbackQueryPayloadClass as nil")
	}
	return b.CallbackQueryPayload.Encode(buf)
}

// DecodeTDLibJSON implements bin.Decoder for CallbackQueryPayloadBox.
func (b *CallbackQueryPayloadBox) DecodeTDLibJSON(buf tdjson.Decoder) error {
	if b == nil {
		return fmt.Errorf("unable to decode CallbackQueryPayloadBox to nil")
	}
	v, err := DecodeTDLibJSONCallbackQueryPayload(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.CallbackQueryPayload = v
	return nil
}

// EncodeTDLibJSON implements bin.Encode for CallbackQueryPayloadBox.
func (b *CallbackQueryPayloadBox) EncodeTDLibJSON(buf tdjson.Encoder) error {
	if b == nil || b.CallbackQueryPayload == nil {
		return fmt.Errorf("unable to encode CallbackQueryPayloadClass as nil")
	}
	return b.CallbackQueryPayload.EncodeTDLibJSON(buf)
}
