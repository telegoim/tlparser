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

// SuggestedActionEnableArchiveAndMuteNewChats represents TL type `suggestedActionEnableArchiveAndMuteNewChats#7841ec4f`.
type SuggestedActionEnableArchiveAndMuteNewChats struct {
}

// SuggestedActionEnableArchiveAndMuteNewChatsTypeID is TL type id of SuggestedActionEnableArchiveAndMuteNewChats.
const SuggestedActionEnableArchiveAndMuteNewChatsTypeID = 0x7841ec4f

// construct implements constructor of SuggestedActionClass.
func (s SuggestedActionEnableArchiveAndMuteNewChats) construct() SuggestedActionClass { return &s }

// Ensuring interfaces in compile-time for SuggestedActionEnableArchiveAndMuteNewChats.
var (
	_ bin.Encoder     = &SuggestedActionEnableArchiveAndMuteNewChats{}
	_ bin.Decoder     = &SuggestedActionEnableArchiveAndMuteNewChats{}
	_ bin.BareEncoder = &SuggestedActionEnableArchiveAndMuteNewChats{}
	_ bin.BareDecoder = &SuggestedActionEnableArchiveAndMuteNewChats{}

	_ SuggestedActionClass = &SuggestedActionEnableArchiveAndMuteNewChats{}
)

func (s *SuggestedActionEnableArchiveAndMuteNewChats) Zero() bool {
	if s == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (s *SuggestedActionEnableArchiveAndMuteNewChats) String() string {
	if s == nil {
		return "SuggestedActionEnableArchiveAndMuteNewChats(nil)"
	}
	type Alias SuggestedActionEnableArchiveAndMuteNewChats
	return fmt.Sprintf("SuggestedActionEnableArchiveAndMuteNewChats%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*SuggestedActionEnableArchiveAndMuteNewChats) TypeID() uint32 {
	return SuggestedActionEnableArchiveAndMuteNewChatsTypeID
}

// TypeName returns name of type in TL schema.
func (*SuggestedActionEnableArchiveAndMuteNewChats) TypeName() string {
	return "suggestedActionEnableArchiveAndMuteNewChats"
}

// TypeInfo returns info about TL type.
func (s *SuggestedActionEnableArchiveAndMuteNewChats) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "suggestedActionEnableArchiveAndMuteNewChats",
		ID:   SuggestedActionEnableArchiveAndMuteNewChatsTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (s *SuggestedActionEnableArchiveAndMuteNewChats) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode suggestedActionEnableArchiveAndMuteNewChats#7841ec4f as nil")
	}
	b.PutID(SuggestedActionEnableArchiveAndMuteNewChatsTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *SuggestedActionEnableArchiveAndMuteNewChats) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode suggestedActionEnableArchiveAndMuteNewChats#7841ec4f as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *SuggestedActionEnableArchiveAndMuteNewChats) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode suggestedActionEnableArchiveAndMuteNewChats#7841ec4f to nil")
	}
	if err := b.ConsumeID(SuggestedActionEnableArchiveAndMuteNewChatsTypeID); err != nil {
		return fmt.Errorf("unable to decode suggestedActionEnableArchiveAndMuteNewChats#7841ec4f: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *SuggestedActionEnableArchiveAndMuteNewChats) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode suggestedActionEnableArchiveAndMuteNewChats#7841ec4f to nil")
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (s *SuggestedActionEnableArchiveAndMuteNewChats) EncodeTDLibJSON(b tdjson.Encoder) error {
	if s == nil {
		return fmt.Errorf("can't encode suggestedActionEnableArchiveAndMuteNewChats#7841ec4f as nil")
	}
	b.ObjStart()
	b.PutID("suggestedActionEnableArchiveAndMuteNewChats")
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (s *SuggestedActionEnableArchiveAndMuteNewChats) DecodeTDLibJSON(b tdjson.Decoder) error {
	if s == nil {
		return fmt.Errorf("can't decode suggestedActionEnableArchiveAndMuteNewChats#7841ec4f to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("suggestedActionEnableArchiveAndMuteNewChats"); err != nil {
				return fmt.Errorf("unable to decode suggestedActionEnableArchiveAndMuteNewChats#7841ec4f: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// SuggestedActionCheckPassword represents TL type `suggestedActionCheckPassword#71e072b7`.
type SuggestedActionCheckPassword struct {
}

// SuggestedActionCheckPasswordTypeID is TL type id of SuggestedActionCheckPassword.
const SuggestedActionCheckPasswordTypeID = 0x71e072b7

// construct implements constructor of SuggestedActionClass.
func (s SuggestedActionCheckPassword) construct() SuggestedActionClass { return &s }

// Ensuring interfaces in compile-time for SuggestedActionCheckPassword.
var (
	_ bin.Encoder     = &SuggestedActionCheckPassword{}
	_ bin.Decoder     = &SuggestedActionCheckPassword{}
	_ bin.BareEncoder = &SuggestedActionCheckPassword{}
	_ bin.BareDecoder = &SuggestedActionCheckPassword{}

	_ SuggestedActionClass = &SuggestedActionCheckPassword{}
)

func (s *SuggestedActionCheckPassword) Zero() bool {
	if s == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (s *SuggestedActionCheckPassword) String() string {
	if s == nil {
		return "SuggestedActionCheckPassword(nil)"
	}
	type Alias SuggestedActionCheckPassword
	return fmt.Sprintf("SuggestedActionCheckPassword%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*SuggestedActionCheckPassword) TypeID() uint32 {
	return SuggestedActionCheckPasswordTypeID
}

// TypeName returns name of type in TL schema.
func (*SuggestedActionCheckPassword) TypeName() string {
	return "suggestedActionCheckPassword"
}

// TypeInfo returns info about TL type.
func (s *SuggestedActionCheckPassword) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "suggestedActionCheckPassword",
		ID:   SuggestedActionCheckPasswordTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (s *SuggestedActionCheckPassword) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode suggestedActionCheckPassword#71e072b7 as nil")
	}
	b.PutID(SuggestedActionCheckPasswordTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *SuggestedActionCheckPassword) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode suggestedActionCheckPassword#71e072b7 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *SuggestedActionCheckPassword) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode suggestedActionCheckPassword#71e072b7 to nil")
	}
	if err := b.ConsumeID(SuggestedActionCheckPasswordTypeID); err != nil {
		return fmt.Errorf("unable to decode suggestedActionCheckPassword#71e072b7: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *SuggestedActionCheckPassword) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode suggestedActionCheckPassword#71e072b7 to nil")
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (s *SuggestedActionCheckPassword) EncodeTDLibJSON(b tdjson.Encoder) error {
	if s == nil {
		return fmt.Errorf("can't encode suggestedActionCheckPassword#71e072b7 as nil")
	}
	b.ObjStart()
	b.PutID("suggestedActionCheckPassword")
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (s *SuggestedActionCheckPassword) DecodeTDLibJSON(b tdjson.Decoder) error {
	if s == nil {
		return fmt.Errorf("can't decode suggestedActionCheckPassword#71e072b7 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("suggestedActionCheckPassword"); err != nil {
				return fmt.Errorf("unable to decode suggestedActionCheckPassword#71e072b7: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// SuggestedActionCheckPhoneNumber represents TL type `suggestedActionCheckPhoneNumber#26ab77eb`.
type SuggestedActionCheckPhoneNumber struct {
}

// SuggestedActionCheckPhoneNumberTypeID is TL type id of SuggestedActionCheckPhoneNumber.
const SuggestedActionCheckPhoneNumberTypeID = 0x26ab77eb

// construct implements constructor of SuggestedActionClass.
func (s SuggestedActionCheckPhoneNumber) construct() SuggestedActionClass { return &s }

// Ensuring interfaces in compile-time for SuggestedActionCheckPhoneNumber.
var (
	_ bin.Encoder     = &SuggestedActionCheckPhoneNumber{}
	_ bin.Decoder     = &SuggestedActionCheckPhoneNumber{}
	_ bin.BareEncoder = &SuggestedActionCheckPhoneNumber{}
	_ bin.BareDecoder = &SuggestedActionCheckPhoneNumber{}

	_ SuggestedActionClass = &SuggestedActionCheckPhoneNumber{}
)

func (s *SuggestedActionCheckPhoneNumber) Zero() bool {
	if s == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (s *SuggestedActionCheckPhoneNumber) String() string {
	if s == nil {
		return "SuggestedActionCheckPhoneNumber(nil)"
	}
	type Alias SuggestedActionCheckPhoneNumber
	return fmt.Sprintf("SuggestedActionCheckPhoneNumber%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*SuggestedActionCheckPhoneNumber) TypeID() uint32 {
	return SuggestedActionCheckPhoneNumberTypeID
}

// TypeName returns name of type in TL schema.
func (*SuggestedActionCheckPhoneNumber) TypeName() string {
	return "suggestedActionCheckPhoneNumber"
}

// TypeInfo returns info about TL type.
func (s *SuggestedActionCheckPhoneNumber) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "suggestedActionCheckPhoneNumber",
		ID:   SuggestedActionCheckPhoneNumberTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (s *SuggestedActionCheckPhoneNumber) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode suggestedActionCheckPhoneNumber#26ab77eb as nil")
	}
	b.PutID(SuggestedActionCheckPhoneNumberTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *SuggestedActionCheckPhoneNumber) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode suggestedActionCheckPhoneNumber#26ab77eb as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *SuggestedActionCheckPhoneNumber) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode suggestedActionCheckPhoneNumber#26ab77eb to nil")
	}
	if err := b.ConsumeID(SuggestedActionCheckPhoneNumberTypeID); err != nil {
		return fmt.Errorf("unable to decode suggestedActionCheckPhoneNumber#26ab77eb: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *SuggestedActionCheckPhoneNumber) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode suggestedActionCheckPhoneNumber#26ab77eb to nil")
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (s *SuggestedActionCheckPhoneNumber) EncodeTDLibJSON(b tdjson.Encoder) error {
	if s == nil {
		return fmt.Errorf("can't encode suggestedActionCheckPhoneNumber#26ab77eb as nil")
	}
	b.ObjStart()
	b.PutID("suggestedActionCheckPhoneNumber")
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (s *SuggestedActionCheckPhoneNumber) DecodeTDLibJSON(b tdjson.Decoder) error {
	if s == nil {
		return fmt.Errorf("can't decode suggestedActionCheckPhoneNumber#26ab77eb to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("suggestedActionCheckPhoneNumber"); err != nil {
				return fmt.Errorf("unable to decode suggestedActionCheckPhoneNumber#26ab77eb: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// SuggestedActionViewChecksHint represents TL type `suggestedActionViewChecksHint#35203547`.
type SuggestedActionViewChecksHint struct {
}

// SuggestedActionViewChecksHintTypeID is TL type id of SuggestedActionViewChecksHint.
const SuggestedActionViewChecksHintTypeID = 0x35203547

// construct implements constructor of SuggestedActionClass.
func (s SuggestedActionViewChecksHint) construct() SuggestedActionClass { return &s }

// Ensuring interfaces in compile-time for SuggestedActionViewChecksHint.
var (
	_ bin.Encoder     = &SuggestedActionViewChecksHint{}
	_ bin.Decoder     = &SuggestedActionViewChecksHint{}
	_ bin.BareEncoder = &SuggestedActionViewChecksHint{}
	_ bin.BareDecoder = &SuggestedActionViewChecksHint{}

	_ SuggestedActionClass = &SuggestedActionViewChecksHint{}
)

func (s *SuggestedActionViewChecksHint) Zero() bool {
	if s == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (s *SuggestedActionViewChecksHint) String() string {
	if s == nil {
		return "SuggestedActionViewChecksHint(nil)"
	}
	type Alias SuggestedActionViewChecksHint
	return fmt.Sprintf("SuggestedActionViewChecksHint%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*SuggestedActionViewChecksHint) TypeID() uint32 {
	return SuggestedActionViewChecksHintTypeID
}

// TypeName returns name of type in TL schema.
func (*SuggestedActionViewChecksHint) TypeName() string {
	return "suggestedActionViewChecksHint"
}

// TypeInfo returns info about TL type.
func (s *SuggestedActionViewChecksHint) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "suggestedActionViewChecksHint",
		ID:   SuggestedActionViewChecksHintTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (s *SuggestedActionViewChecksHint) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode suggestedActionViewChecksHint#35203547 as nil")
	}
	b.PutID(SuggestedActionViewChecksHintTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *SuggestedActionViewChecksHint) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode suggestedActionViewChecksHint#35203547 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *SuggestedActionViewChecksHint) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode suggestedActionViewChecksHint#35203547 to nil")
	}
	if err := b.ConsumeID(SuggestedActionViewChecksHintTypeID); err != nil {
		return fmt.Errorf("unable to decode suggestedActionViewChecksHint#35203547: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *SuggestedActionViewChecksHint) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode suggestedActionViewChecksHint#35203547 to nil")
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (s *SuggestedActionViewChecksHint) EncodeTDLibJSON(b tdjson.Encoder) error {
	if s == nil {
		return fmt.Errorf("can't encode suggestedActionViewChecksHint#35203547 as nil")
	}
	b.ObjStart()
	b.PutID("suggestedActionViewChecksHint")
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (s *SuggestedActionViewChecksHint) DecodeTDLibJSON(b tdjson.Decoder) error {
	if s == nil {
		return fmt.Errorf("can't decode suggestedActionViewChecksHint#35203547 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("suggestedActionViewChecksHint"); err != nil {
				return fmt.Errorf("unable to decode suggestedActionViewChecksHint#35203547: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// SuggestedActionConvertToBroadcastGroup represents TL type `suggestedActionConvertToBroadcastGroup#c67a2e38`.
type SuggestedActionConvertToBroadcastGroup struct {
	// Supergroup identifier
	SupergroupID int64
}

// SuggestedActionConvertToBroadcastGroupTypeID is TL type id of SuggestedActionConvertToBroadcastGroup.
const SuggestedActionConvertToBroadcastGroupTypeID = 0xc67a2e38

// construct implements constructor of SuggestedActionClass.
func (s SuggestedActionConvertToBroadcastGroup) construct() SuggestedActionClass { return &s }

// Ensuring interfaces in compile-time for SuggestedActionConvertToBroadcastGroup.
var (
	_ bin.Encoder     = &SuggestedActionConvertToBroadcastGroup{}
	_ bin.Decoder     = &SuggestedActionConvertToBroadcastGroup{}
	_ bin.BareEncoder = &SuggestedActionConvertToBroadcastGroup{}
	_ bin.BareDecoder = &SuggestedActionConvertToBroadcastGroup{}

	_ SuggestedActionClass = &SuggestedActionConvertToBroadcastGroup{}
)

func (s *SuggestedActionConvertToBroadcastGroup) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.SupergroupID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *SuggestedActionConvertToBroadcastGroup) String() string {
	if s == nil {
		return "SuggestedActionConvertToBroadcastGroup(nil)"
	}
	type Alias SuggestedActionConvertToBroadcastGroup
	return fmt.Sprintf("SuggestedActionConvertToBroadcastGroup%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*SuggestedActionConvertToBroadcastGroup) TypeID() uint32 {
	return SuggestedActionConvertToBroadcastGroupTypeID
}

// TypeName returns name of type in TL schema.
func (*SuggestedActionConvertToBroadcastGroup) TypeName() string {
	return "suggestedActionConvertToBroadcastGroup"
}

// TypeInfo returns info about TL type.
func (s *SuggestedActionConvertToBroadcastGroup) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "suggestedActionConvertToBroadcastGroup",
		ID:   SuggestedActionConvertToBroadcastGroupTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "SupergroupID",
			SchemaName: "supergroup_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *SuggestedActionConvertToBroadcastGroup) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode suggestedActionConvertToBroadcastGroup#c67a2e38 as nil")
	}
	b.PutID(SuggestedActionConvertToBroadcastGroupTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *SuggestedActionConvertToBroadcastGroup) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode suggestedActionConvertToBroadcastGroup#c67a2e38 as nil")
	}
	b.PutInt53(s.SupergroupID)
	return nil
}

// Decode implements bin.Decoder.
func (s *SuggestedActionConvertToBroadcastGroup) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode suggestedActionConvertToBroadcastGroup#c67a2e38 to nil")
	}
	if err := b.ConsumeID(SuggestedActionConvertToBroadcastGroupTypeID); err != nil {
		return fmt.Errorf("unable to decode suggestedActionConvertToBroadcastGroup#c67a2e38: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *SuggestedActionConvertToBroadcastGroup) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode suggestedActionConvertToBroadcastGroup#c67a2e38 to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode suggestedActionConvertToBroadcastGroup#c67a2e38: field supergroup_id: %w", err)
		}
		s.SupergroupID = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (s *SuggestedActionConvertToBroadcastGroup) EncodeTDLibJSON(b tdjson.Encoder) error {
	if s == nil {
		return fmt.Errorf("can't encode suggestedActionConvertToBroadcastGroup#c67a2e38 as nil")
	}
	b.ObjStart()
	b.PutID("suggestedActionConvertToBroadcastGroup")
	b.Comma()
	b.FieldStart("supergroup_id")
	b.PutInt53(s.SupergroupID)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (s *SuggestedActionConvertToBroadcastGroup) DecodeTDLibJSON(b tdjson.Decoder) error {
	if s == nil {
		return fmt.Errorf("can't decode suggestedActionConvertToBroadcastGroup#c67a2e38 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("suggestedActionConvertToBroadcastGroup"); err != nil {
				return fmt.Errorf("unable to decode suggestedActionConvertToBroadcastGroup#c67a2e38: %w", err)
			}
		case "supergroup_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode suggestedActionConvertToBroadcastGroup#c67a2e38: field supergroup_id: %w", err)
			}
			s.SupergroupID = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetSupergroupID returns value of SupergroupID field.
func (s *SuggestedActionConvertToBroadcastGroup) GetSupergroupID() (value int64) {
	if s == nil {
		return
	}
	return s.SupergroupID
}

// SuggestedActionSetPassword represents TL type `suggestedActionSetPassword#6f147d98`.
type SuggestedActionSetPassword struct {
	// The number of days to pass between consecutive authorizations if the user declines to
	// set password
	AuthorizationDelay int32
}

// SuggestedActionSetPasswordTypeID is TL type id of SuggestedActionSetPassword.
const SuggestedActionSetPasswordTypeID = 0x6f147d98

// construct implements constructor of SuggestedActionClass.
func (s SuggestedActionSetPassword) construct() SuggestedActionClass { return &s }

// Ensuring interfaces in compile-time for SuggestedActionSetPassword.
var (
	_ bin.Encoder     = &SuggestedActionSetPassword{}
	_ bin.Decoder     = &SuggestedActionSetPassword{}
	_ bin.BareEncoder = &SuggestedActionSetPassword{}
	_ bin.BareDecoder = &SuggestedActionSetPassword{}

	_ SuggestedActionClass = &SuggestedActionSetPassword{}
)

func (s *SuggestedActionSetPassword) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.AuthorizationDelay == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *SuggestedActionSetPassword) String() string {
	if s == nil {
		return "SuggestedActionSetPassword(nil)"
	}
	type Alias SuggestedActionSetPassword
	return fmt.Sprintf("SuggestedActionSetPassword%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*SuggestedActionSetPassword) TypeID() uint32 {
	return SuggestedActionSetPasswordTypeID
}

// TypeName returns name of type in TL schema.
func (*SuggestedActionSetPassword) TypeName() string {
	return "suggestedActionSetPassword"
}

// TypeInfo returns info about TL type.
func (s *SuggestedActionSetPassword) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "suggestedActionSetPassword",
		ID:   SuggestedActionSetPasswordTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "AuthorizationDelay",
			SchemaName: "authorization_delay",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *SuggestedActionSetPassword) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode suggestedActionSetPassword#6f147d98 as nil")
	}
	b.PutID(SuggestedActionSetPasswordTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *SuggestedActionSetPassword) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode suggestedActionSetPassword#6f147d98 as nil")
	}
	b.PutInt32(s.AuthorizationDelay)
	return nil
}

// Decode implements bin.Decoder.
func (s *SuggestedActionSetPassword) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode suggestedActionSetPassword#6f147d98 to nil")
	}
	if err := b.ConsumeID(SuggestedActionSetPasswordTypeID); err != nil {
		return fmt.Errorf("unable to decode suggestedActionSetPassword#6f147d98: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *SuggestedActionSetPassword) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode suggestedActionSetPassword#6f147d98 to nil")
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode suggestedActionSetPassword#6f147d98: field authorization_delay: %w", err)
		}
		s.AuthorizationDelay = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (s *SuggestedActionSetPassword) EncodeTDLibJSON(b tdjson.Encoder) error {
	if s == nil {
		return fmt.Errorf("can't encode suggestedActionSetPassword#6f147d98 as nil")
	}
	b.ObjStart()
	b.PutID("suggestedActionSetPassword")
	b.Comma()
	b.FieldStart("authorization_delay")
	b.PutInt32(s.AuthorizationDelay)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (s *SuggestedActionSetPassword) DecodeTDLibJSON(b tdjson.Decoder) error {
	if s == nil {
		return fmt.Errorf("can't decode suggestedActionSetPassword#6f147d98 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("suggestedActionSetPassword"); err != nil {
				return fmt.Errorf("unable to decode suggestedActionSetPassword#6f147d98: %w", err)
			}
		case "authorization_delay":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode suggestedActionSetPassword#6f147d98: field authorization_delay: %w", err)
			}
			s.AuthorizationDelay = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetAuthorizationDelay returns value of AuthorizationDelay field.
func (s *SuggestedActionSetPassword) GetAuthorizationDelay() (value int32) {
	if s == nil {
		return
	}
	return s.AuthorizationDelay
}

// SuggestedActionClassName is schema name of SuggestedActionClass.
const SuggestedActionClassName = "SuggestedAction"

// SuggestedActionClass represents SuggestedAction generic type.
//
// Example:
//
//	g, err := tdapi.DecodeSuggestedAction(buf)
//	if err != nil {
//	    panic(err)
//	}
//	switch v := g.(type) {
//	case *tdapi.SuggestedActionEnableArchiveAndMuteNewChats: // suggestedActionEnableArchiveAndMuteNewChats#7841ec4f
//	case *tdapi.SuggestedActionCheckPassword: // suggestedActionCheckPassword#71e072b7
//	case *tdapi.SuggestedActionCheckPhoneNumber: // suggestedActionCheckPhoneNumber#26ab77eb
//	case *tdapi.SuggestedActionViewChecksHint: // suggestedActionViewChecksHint#35203547
//	case *tdapi.SuggestedActionConvertToBroadcastGroup: // suggestedActionConvertToBroadcastGroup#c67a2e38
//	case *tdapi.SuggestedActionSetPassword: // suggestedActionSetPassword#6f147d98
//	default: panic(v)
//	}
type SuggestedActionClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() SuggestedActionClass

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

// DecodeSuggestedAction implements binary de-serialization for SuggestedActionClass.
func DecodeSuggestedAction(buf *bin.Buffer) (SuggestedActionClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case SuggestedActionEnableArchiveAndMuteNewChatsTypeID:
		// Decoding suggestedActionEnableArchiveAndMuteNewChats#7841ec4f.
		v := SuggestedActionEnableArchiveAndMuteNewChats{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode SuggestedActionClass: %w", err)
		}
		return &v, nil
	case SuggestedActionCheckPasswordTypeID:
		// Decoding suggestedActionCheckPassword#71e072b7.
		v := SuggestedActionCheckPassword{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode SuggestedActionClass: %w", err)
		}
		return &v, nil
	case SuggestedActionCheckPhoneNumberTypeID:
		// Decoding suggestedActionCheckPhoneNumber#26ab77eb.
		v := SuggestedActionCheckPhoneNumber{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode SuggestedActionClass: %w", err)
		}
		return &v, nil
	case SuggestedActionViewChecksHintTypeID:
		// Decoding suggestedActionViewChecksHint#35203547.
		v := SuggestedActionViewChecksHint{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode SuggestedActionClass: %w", err)
		}
		return &v, nil
	case SuggestedActionConvertToBroadcastGroupTypeID:
		// Decoding suggestedActionConvertToBroadcastGroup#c67a2e38.
		v := SuggestedActionConvertToBroadcastGroup{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode SuggestedActionClass: %w", err)
		}
		return &v, nil
	case SuggestedActionSetPasswordTypeID:
		// Decoding suggestedActionSetPassword#6f147d98.
		v := SuggestedActionSetPassword{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode SuggestedActionClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode SuggestedActionClass: %w", bin.NewUnexpectedID(id))
	}
}

// DecodeTDLibJSONSuggestedAction implements binary de-serialization for SuggestedActionClass.
func DecodeTDLibJSONSuggestedAction(buf tdjson.Decoder) (SuggestedActionClass, error) {
	id, err := buf.FindTypeID()
	if err != nil {
		return nil, err
	}
	switch id {
	case "suggestedActionEnableArchiveAndMuteNewChats":
		// Decoding suggestedActionEnableArchiveAndMuteNewChats#7841ec4f.
		v := SuggestedActionEnableArchiveAndMuteNewChats{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode SuggestedActionClass: %w", err)
		}
		return &v, nil
	case "suggestedActionCheckPassword":
		// Decoding suggestedActionCheckPassword#71e072b7.
		v := SuggestedActionCheckPassword{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode SuggestedActionClass: %w", err)
		}
		return &v, nil
	case "suggestedActionCheckPhoneNumber":
		// Decoding suggestedActionCheckPhoneNumber#26ab77eb.
		v := SuggestedActionCheckPhoneNumber{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode SuggestedActionClass: %w", err)
		}
		return &v, nil
	case "suggestedActionViewChecksHint":
		// Decoding suggestedActionViewChecksHint#35203547.
		v := SuggestedActionViewChecksHint{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode SuggestedActionClass: %w", err)
		}
		return &v, nil
	case "suggestedActionConvertToBroadcastGroup":
		// Decoding suggestedActionConvertToBroadcastGroup#c67a2e38.
		v := SuggestedActionConvertToBroadcastGroup{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode SuggestedActionClass: %w", err)
		}
		return &v, nil
	case "suggestedActionSetPassword":
		// Decoding suggestedActionSetPassword#6f147d98.
		v := SuggestedActionSetPassword{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode SuggestedActionClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode SuggestedActionClass: %w", tdjson.NewUnexpectedID(id))
	}
}

// SuggestedAction boxes the SuggestedActionClass providing a helper.
type SuggestedActionBox struct {
	SuggestedAction SuggestedActionClass
}

// Decode implements bin.Decoder for SuggestedActionBox.
func (b *SuggestedActionBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode SuggestedActionBox to nil")
	}
	v, err := DecodeSuggestedAction(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.SuggestedAction = v
	return nil
}

// Encode implements bin.Encode for SuggestedActionBox.
func (b *SuggestedActionBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.SuggestedAction == nil {
		return fmt.Errorf("unable to encode SuggestedActionClass as nil")
	}
	return b.SuggestedAction.Encode(buf)
}

// DecodeTDLibJSON implements bin.Decoder for SuggestedActionBox.
func (b *SuggestedActionBox) DecodeTDLibJSON(buf tdjson.Decoder) error {
	if b == nil {
		return fmt.Errorf("unable to decode SuggestedActionBox to nil")
	}
	v, err := DecodeTDLibJSONSuggestedAction(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.SuggestedAction = v
	return nil
}

// EncodeTDLibJSON implements bin.Encode for SuggestedActionBox.
func (b *SuggestedActionBox) EncodeTDLibJSON(buf tdjson.Encoder) error {
	if b == nil || b.SuggestedAction == nil {
		return fmt.Errorf("unable to encode SuggestedActionClass as nil")
	}
	return b.SuggestedAction.EncodeTDLibJSON(buf)
}
