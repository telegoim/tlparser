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

// InputPrivacyKeyStatusTimestamp represents TL type `inputPrivacyKeyStatusTimestamp#4f96cb18`.
// Whether we can see the exact last online timestamp of the user
//
// See https://core.telegram.org/constructor/inputPrivacyKeyStatusTimestamp for reference.
type InputPrivacyKeyStatusTimestamp struct {
}

// InputPrivacyKeyStatusTimestampTypeID is TL type id of InputPrivacyKeyStatusTimestamp.
const InputPrivacyKeyStatusTimestampTypeID = 0x4f96cb18

// construct implements constructor of InputPrivacyKeyClass.
func (i InputPrivacyKeyStatusTimestamp) construct() InputPrivacyKeyClass { return &i }

// Ensuring interfaces in compile-time for InputPrivacyKeyStatusTimestamp.
var (
	_ bin.Encoder     = &InputPrivacyKeyStatusTimestamp{}
	_ bin.Decoder     = &InputPrivacyKeyStatusTimestamp{}
	_ bin.BareEncoder = &InputPrivacyKeyStatusTimestamp{}
	_ bin.BareDecoder = &InputPrivacyKeyStatusTimestamp{}

	_ InputPrivacyKeyClass = &InputPrivacyKeyStatusTimestamp{}
)

func (i *InputPrivacyKeyStatusTimestamp) Zero() bool {
	if i == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputPrivacyKeyStatusTimestamp) String() string {
	if i == nil {
		return "InputPrivacyKeyStatusTimestamp(nil)"
	}
	type Alias InputPrivacyKeyStatusTimestamp
	return fmt.Sprintf("InputPrivacyKeyStatusTimestamp%+v", Alias(*i))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InputPrivacyKeyStatusTimestamp) TypeID() uint32 {
	return InputPrivacyKeyStatusTimestampTypeID
}

// TypeName returns name of type in TL schema.
func (*InputPrivacyKeyStatusTimestamp) TypeName() string {
	return "inputPrivacyKeyStatusTimestamp"
}

// TypeInfo returns info about TL type.
func (i *InputPrivacyKeyStatusTimestamp) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inputPrivacyKeyStatusTimestamp",
		ID:   InputPrivacyKeyStatusTimestampTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (i *InputPrivacyKeyStatusTimestamp) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputPrivacyKeyStatusTimestamp#4f96cb18 as nil")
	}
	b.PutID(InputPrivacyKeyStatusTimestampTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InputPrivacyKeyStatusTimestamp) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputPrivacyKeyStatusTimestamp#4f96cb18 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (i *InputPrivacyKeyStatusTimestamp) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputPrivacyKeyStatusTimestamp#4f96cb18 to nil")
	}
	if err := b.ConsumeID(InputPrivacyKeyStatusTimestampTypeID); err != nil {
		return fmt.Errorf("unable to decode inputPrivacyKeyStatusTimestamp#4f96cb18: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InputPrivacyKeyStatusTimestamp) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputPrivacyKeyStatusTimestamp#4f96cb18 to nil")
	}
	return nil
}

// InputPrivacyKeyChatInvite represents TL type `inputPrivacyKeyChatInvite#bdfb0426`.
// Whether the user can be invited to chats
//
// See https://core.telegram.org/constructor/inputPrivacyKeyChatInvite for reference.
type InputPrivacyKeyChatInvite struct {
}

// InputPrivacyKeyChatInviteTypeID is TL type id of InputPrivacyKeyChatInvite.
const InputPrivacyKeyChatInviteTypeID = 0xbdfb0426

// construct implements constructor of InputPrivacyKeyClass.
func (i InputPrivacyKeyChatInvite) construct() InputPrivacyKeyClass { return &i }

// Ensuring interfaces in compile-time for InputPrivacyKeyChatInvite.
var (
	_ bin.Encoder     = &InputPrivacyKeyChatInvite{}
	_ bin.Decoder     = &InputPrivacyKeyChatInvite{}
	_ bin.BareEncoder = &InputPrivacyKeyChatInvite{}
	_ bin.BareDecoder = &InputPrivacyKeyChatInvite{}

	_ InputPrivacyKeyClass = &InputPrivacyKeyChatInvite{}
)

func (i *InputPrivacyKeyChatInvite) Zero() bool {
	if i == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputPrivacyKeyChatInvite) String() string {
	if i == nil {
		return "InputPrivacyKeyChatInvite(nil)"
	}
	type Alias InputPrivacyKeyChatInvite
	return fmt.Sprintf("InputPrivacyKeyChatInvite%+v", Alias(*i))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InputPrivacyKeyChatInvite) TypeID() uint32 {
	return InputPrivacyKeyChatInviteTypeID
}

// TypeName returns name of type in TL schema.
func (*InputPrivacyKeyChatInvite) TypeName() string {
	return "inputPrivacyKeyChatInvite"
}

// TypeInfo returns info about TL type.
func (i *InputPrivacyKeyChatInvite) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inputPrivacyKeyChatInvite",
		ID:   InputPrivacyKeyChatInviteTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (i *InputPrivacyKeyChatInvite) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputPrivacyKeyChatInvite#bdfb0426 as nil")
	}
	b.PutID(InputPrivacyKeyChatInviteTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InputPrivacyKeyChatInvite) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputPrivacyKeyChatInvite#bdfb0426 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (i *InputPrivacyKeyChatInvite) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputPrivacyKeyChatInvite#bdfb0426 to nil")
	}
	if err := b.ConsumeID(InputPrivacyKeyChatInviteTypeID); err != nil {
		return fmt.Errorf("unable to decode inputPrivacyKeyChatInvite#bdfb0426: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InputPrivacyKeyChatInvite) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputPrivacyKeyChatInvite#bdfb0426 to nil")
	}
	return nil
}

// InputPrivacyKeyPhoneCall represents TL type `inputPrivacyKeyPhoneCall#fabadc5f`.
// Whether the user will accept phone calls
//
// See https://core.telegram.org/constructor/inputPrivacyKeyPhoneCall for reference.
type InputPrivacyKeyPhoneCall struct {
}

// InputPrivacyKeyPhoneCallTypeID is TL type id of InputPrivacyKeyPhoneCall.
const InputPrivacyKeyPhoneCallTypeID = 0xfabadc5f

// construct implements constructor of InputPrivacyKeyClass.
func (i InputPrivacyKeyPhoneCall) construct() InputPrivacyKeyClass { return &i }

// Ensuring interfaces in compile-time for InputPrivacyKeyPhoneCall.
var (
	_ bin.Encoder     = &InputPrivacyKeyPhoneCall{}
	_ bin.Decoder     = &InputPrivacyKeyPhoneCall{}
	_ bin.BareEncoder = &InputPrivacyKeyPhoneCall{}
	_ bin.BareDecoder = &InputPrivacyKeyPhoneCall{}

	_ InputPrivacyKeyClass = &InputPrivacyKeyPhoneCall{}
)

func (i *InputPrivacyKeyPhoneCall) Zero() bool {
	if i == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputPrivacyKeyPhoneCall) String() string {
	if i == nil {
		return "InputPrivacyKeyPhoneCall(nil)"
	}
	type Alias InputPrivacyKeyPhoneCall
	return fmt.Sprintf("InputPrivacyKeyPhoneCall%+v", Alias(*i))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InputPrivacyKeyPhoneCall) TypeID() uint32 {
	return InputPrivacyKeyPhoneCallTypeID
}

// TypeName returns name of type in TL schema.
func (*InputPrivacyKeyPhoneCall) TypeName() string {
	return "inputPrivacyKeyPhoneCall"
}

// TypeInfo returns info about TL type.
func (i *InputPrivacyKeyPhoneCall) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inputPrivacyKeyPhoneCall",
		ID:   InputPrivacyKeyPhoneCallTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (i *InputPrivacyKeyPhoneCall) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputPrivacyKeyPhoneCall#fabadc5f as nil")
	}
	b.PutID(InputPrivacyKeyPhoneCallTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InputPrivacyKeyPhoneCall) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputPrivacyKeyPhoneCall#fabadc5f as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (i *InputPrivacyKeyPhoneCall) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputPrivacyKeyPhoneCall#fabadc5f to nil")
	}
	if err := b.ConsumeID(InputPrivacyKeyPhoneCallTypeID); err != nil {
		return fmt.Errorf("unable to decode inputPrivacyKeyPhoneCall#fabadc5f: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InputPrivacyKeyPhoneCall) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputPrivacyKeyPhoneCall#fabadc5f to nil")
	}
	return nil
}

// InputPrivacyKeyPhoneP2P represents TL type `inputPrivacyKeyPhoneP2P#db9e70d2`.
// Whether the user allows P2P communication during VoIP calls
//
// See https://core.telegram.org/constructor/inputPrivacyKeyPhoneP2P for reference.
type InputPrivacyKeyPhoneP2P struct {
}

// InputPrivacyKeyPhoneP2PTypeID is TL type id of InputPrivacyKeyPhoneP2P.
const InputPrivacyKeyPhoneP2PTypeID = 0xdb9e70d2

// construct implements constructor of InputPrivacyKeyClass.
func (i InputPrivacyKeyPhoneP2P) construct() InputPrivacyKeyClass { return &i }

// Ensuring interfaces in compile-time for InputPrivacyKeyPhoneP2P.
var (
	_ bin.Encoder     = &InputPrivacyKeyPhoneP2P{}
	_ bin.Decoder     = &InputPrivacyKeyPhoneP2P{}
	_ bin.BareEncoder = &InputPrivacyKeyPhoneP2P{}
	_ bin.BareDecoder = &InputPrivacyKeyPhoneP2P{}

	_ InputPrivacyKeyClass = &InputPrivacyKeyPhoneP2P{}
)

func (i *InputPrivacyKeyPhoneP2P) Zero() bool {
	if i == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputPrivacyKeyPhoneP2P) String() string {
	if i == nil {
		return "InputPrivacyKeyPhoneP2P(nil)"
	}
	type Alias InputPrivacyKeyPhoneP2P
	return fmt.Sprintf("InputPrivacyKeyPhoneP2P%+v", Alias(*i))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InputPrivacyKeyPhoneP2P) TypeID() uint32 {
	return InputPrivacyKeyPhoneP2PTypeID
}

// TypeName returns name of type in TL schema.
func (*InputPrivacyKeyPhoneP2P) TypeName() string {
	return "inputPrivacyKeyPhoneP2P"
}

// TypeInfo returns info about TL type.
func (i *InputPrivacyKeyPhoneP2P) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inputPrivacyKeyPhoneP2P",
		ID:   InputPrivacyKeyPhoneP2PTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (i *InputPrivacyKeyPhoneP2P) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputPrivacyKeyPhoneP2P#db9e70d2 as nil")
	}
	b.PutID(InputPrivacyKeyPhoneP2PTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InputPrivacyKeyPhoneP2P) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputPrivacyKeyPhoneP2P#db9e70d2 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (i *InputPrivacyKeyPhoneP2P) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputPrivacyKeyPhoneP2P#db9e70d2 to nil")
	}
	if err := b.ConsumeID(InputPrivacyKeyPhoneP2PTypeID); err != nil {
		return fmt.Errorf("unable to decode inputPrivacyKeyPhoneP2P#db9e70d2: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InputPrivacyKeyPhoneP2P) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputPrivacyKeyPhoneP2P#db9e70d2 to nil")
	}
	return nil
}

// InputPrivacyKeyForwards represents TL type `inputPrivacyKeyForwards#a4dd4c08`.
// Whether messages forwarded from this user will be anonymous¹
//
// Links:
//  1. https://telegram.org/blog/unsend-privacy-emoji#anonymous-forwarding
//
// See https://core.telegram.org/constructor/inputPrivacyKeyForwards for reference.
type InputPrivacyKeyForwards struct {
}

// InputPrivacyKeyForwardsTypeID is TL type id of InputPrivacyKeyForwards.
const InputPrivacyKeyForwardsTypeID = 0xa4dd4c08

// construct implements constructor of InputPrivacyKeyClass.
func (i InputPrivacyKeyForwards) construct() InputPrivacyKeyClass { return &i }

// Ensuring interfaces in compile-time for InputPrivacyKeyForwards.
var (
	_ bin.Encoder     = &InputPrivacyKeyForwards{}
	_ bin.Decoder     = &InputPrivacyKeyForwards{}
	_ bin.BareEncoder = &InputPrivacyKeyForwards{}
	_ bin.BareDecoder = &InputPrivacyKeyForwards{}

	_ InputPrivacyKeyClass = &InputPrivacyKeyForwards{}
)

func (i *InputPrivacyKeyForwards) Zero() bool {
	if i == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputPrivacyKeyForwards) String() string {
	if i == nil {
		return "InputPrivacyKeyForwards(nil)"
	}
	type Alias InputPrivacyKeyForwards
	return fmt.Sprintf("InputPrivacyKeyForwards%+v", Alias(*i))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InputPrivacyKeyForwards) TypeID() uint32 {
	return InputPrivacyKeyForwardsTypeID
}

// TypeName returns name of type in TL schema.
func (*InputPrivacyKeyForwards) TypeName() string {
	return "inputPrivacyKeyForwards"
}

// TypeInfo returns info about TL type.
func (i *InputPrivacyKeyForwards) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inputPrivacyKeyForwards",
		ID:   InputPrivacyKeyForwardsTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (i *InputPrivacyKeyForwards) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputPrivacyKeyForwards#a4dd4c08 as nil")
	}
	b.PutID(InputPrivacyKeyForwardsTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InputPrivacyKeyForwards) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputPrivacyKeyForwards#a4dd4c08 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (i *InputPrivacyKeyForwards) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputPrivacyKeyForwards#a4dd4c08 to nil")
	}
	if err := b.ConsumeID(InputPrivacyKeyForwardsTypeID); err != nil {
		return fmt.Errorf("unable to decode inputPrivacyKeyForwards#a4dd4c08: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InputPrivacyKeyForwards) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputPrivacyKeyForwards#a4dd4c08 to nil")
	}
	return nil
}

// InputPrivacyKeyProfilePhoto represents TL type `inputPrivacyKeyProfilePhoto#5719bacc`.
// Whether people will be able to see the user's profile picture
//
// See https://core.telegram.org/constructor/inputPrivacyKeyProfilePhoto for reference.
type InputPrivacyKeyProfilePhoto struct {
}

// InputPrivacyKeyProfilePhotoTypeID is TL type id of InputPrivacyKeyProfilePhoto.
const InputPrivacyKeyProfilePhotoTypeID = 0x5719bacc

// construct implements constructor of InputPrivacyKeyClass.
func (i InputPrivacyKeyProfilePhoto) construct() InputPrivacyKeyClass { return &i }

// Ensuring interfaces in compile-time for InputPrivacyKeyProfilePhoto.
var (
	_ bin.Encoder     = &InputPrivacyKeyProfilePhoto{}
	_ bin.Decoder     = &InputPrivacyKeyProfilePhoto{}
	_ bin.BareEncoder = &InputPrivacyKeyProfilePhoto{}
	_ bin.BareDecoder = &InputPrivacyKeyProfilePhoto{}

	_ InputPrivacyKeyClass = &InputPrivacyKeyProfilePhoto{}
)

func (i *InputPrivacyKeyProfilePhoto) Zero() bool {
	if i == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputPrivacyKeyProfilePhoto) String() string {
	if i == nil {
		return "InputPrivacyKeyProfilePhoto(nil)"
	}
	type Alias InputPrivacyKeyProfilePhoto
	return fmt.Sprintf("InputPrivacyKeyProfilePhoto%+v", Alias(*i))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InputPrivacyKeyProfilePhoto) TypeID() uint32 {
	return InputPrivacyKeyProfilePhotoTypeID
}

// TypeName returns name of type in TL schema.
func (*InputPrivacyKeyProfilePhoto) TypeName() string {
	return "inputPrivacyKeyProfilePhoto"
}

// TypeInfo returns info about TL type.
func (i *InputPrivacyKeyProfilePhoto) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inputPrivacyKeyProfilePhoto",
		ID:   InputPrivacyKeyProfilePhotoTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (i *InputPrivacyKeyProfilePhoto) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputPrivacyKeyProfilePhoto#5719bacc as nil")
	}
	b.PutID(InputPrivacyKeyProfilePhotoTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InputPrivacyKeyProfilePhoto) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputPrivacyKeyProfilePhoto#5719bacc as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (i *InputPrivacyKeyProfilePhoto) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputPrivacyKeyProfilePhoto#5719bacc to nil")
	}
	if err := b.ConsumeID(InputPrivacyKeyProfilePhotoTypeID); err != nil {
		return fmt.Errorf("unable to decode inputPrivacyKeyProfilePhoto#5719bacc: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InputPrivacyKeyProfilePhoto) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputPrivacyKeyProfilePhoto#5719bacc to nil")
	}
	return nil
}

// InputPrivacyKeyPhoneNumber represents TL type `inputPrivacyKeyPhoneNumber#352dafa`.
// Whether people will be able to see the user's phone number
//
// See https://core.telegram.org/constructor/inputPrivacyKeyPhoneNumber for reference.
type InputPrivacyKeyPhoneNumber struct {
}

// InputPrivacyKeyPhoneNumberTypeID is TL type id of InputPrivacyKeyPhoneNumber.
const InputPrivacyKeyPhoneNumberTypeID = 0x352dafa

// construct implements constructor of InputPrivacyKeyClass.
func (i InputPrivacyKeyPhoneNumber) construct() InputPrivacyKeyClass { return &i }

// Ensuring interfaces in compile-time for InputPrivacyKeyPhoneNumber.
var (
	_ bin.Encoder     = &InputPrivacyKeyPhoneNumber{}
	_ bin.Decoder     = &InputPrivacyKeyPhoneNumber{}
	_ bin.BareEncoder = &InputPrivacyKeyPhoneNumber{}
	_ bin.BareDecoder = &InputPrivacyKeyPhoneNumber{}

	_ InputPrivacyKeyClass = &InputPrivacyKeyPhoneNumber{}
)

func (i *InputPrivacyKeyPhoneNumber) Zero() bool {
	if i == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputPrivacyKeyPhoneNumber) String() string {
	if i == nil {
		return "InputPrivacyKeyPhoneNumber(nil)"
	}
	type Alias InputPrivacyKeyPhoneNumber
	return fmt.Sprintf("InputPrivacyKeyPhoneNumber%+v", Alias(*i))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InputPrivacyKeyPhoneNumber) TypeID() uint32 {
	return InputPrivacyKeyPhoneNumberTypeID
}

// TypeName returns name of type in TL schema.
func (*InputPrivacyKeyPhoneNumber) TypeName() string {
	return "inputPrivacyKeyPhoneNumber"
}

// TypeInfo returns info about TL type.
func (i *InputPrivacyKeyPhoneNumber) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inputPrivacyKeyPhoneNumber",
		ID:   InputPrivacyKeyPhoneNumberTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (i *InputPrivacyKeyPhoneNumber) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputPrivacyKeyPhoneNumber#352dafa as nil")
	}
	b.PutID(InputPrivacyKeyPhoneNumberTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InputPrivacyKeyPhoneNumber) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputPrivacyKeyPhoneNumber#352dafa as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (i *InputPrivacyKeyPhoneNumber) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputPrivacyKeyPhoneNumber#352dafa to nil")
	}
	if err := b.ConsumeID(InputPrivacyKeyPhoneNumberTypeID); err != nil {
		return fmt.Errorf("unable to decode inputPrivacyKeyPhoneNumber#352dafa: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InputPrivacyKeyPhoneNumber) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputPrivacyKeyPhoneNumber#352dafa to nil")
	}
	return nil
}

// InputPrivacyKeyAddedByPhone represents TL type `inputPrivacyKeyAddedByPhone#d1219bdd`.
// Whether people can add you to their contact list by your phone number
//
// See https://core.telegram.org/constructor/inputPrivacyKeyAddedByPhone for reference.
type InputPrivacyKeyAddedByPhone struct {
}

// InputPrivacyKeyAddedByPhoneTypeID is TL type id of InputPrivacyKeyAddedByPhone.
const InputPrivacyKeyAddedByPhoneTypeID = 0xd1219bdd

// construct implements constructor of InputPrivacyKeyClass.
func (i InputPrivacyKeyAddedByPhone) construct() InputPrivacyKeyClass { return &i }

// Ensuring interfaces in compile-time for InputPrivacyKeyAddedByPhone.
var (
	_ bin.Encoder     = &InputPrivacyKeyAddedByPhone{}
	_ bin.Decoder     = &InputPrivacyKeyAddedByPhone{}
	_ bin.BareEncoder = &InputPrivacyKeyAddedByPhone{}
	_ bin.BareDecoder = &InputPrivacyKeyAddedByPhone{}

	_ InputPrivacyKeyClass = &InputPrivacyKeyAddedByPhone{}
)

func (i *InputPrivacyKeyAddedByPhone) Zero() bool {
	if i == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputPrivacyKeyAddedByPhone) String() string {
	if i == nil {
		return "InputPrivacyKeyAddedByPhone(nil)"
	}
	type Alias InputPrivacyKeyAddedByPhone
	return fmt.Sprintf("InputPrivacyKeyAddedByPhone%+v", Alias(*i))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InputPrivacyKeyAddedByPhone) TypeID() uint32 {
	return InputPrivacyKeyAddedByPhoneTypeID
}

// TypeName returns name of type in TL schema.
func (*InputPrivacyKeyAddedByPhone) TypeName() string {
	return "inputPrivacyKeyAddedByPhone"
}

// TypeInfo returns info about TL type.
func (i *InputPrivacyKeyAddedByPhone) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inputPrivacyKeyAddedByPhone",
		ID:   InputPrivacyKeyAddedByPhoneTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (i *InputPrivacyKeyAddedByPhone) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputPrivacyKeyAddedByPhone#d1219bdd as nil")
	}
	b.PutID(InputPrivacyKeyAddedByPhoneTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InputPrivacyKeyAddedByPhone) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputPrivacyKeyAddedByPhone#d1219bdd as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (i *InputPrivacyKeyAddedByPhone) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputPrivacyKeyAddedByPhone#d1219bdd to nil")
	}
	if err := b.ConsumeID(InputPrivacyKeyAddedByPhoneTypeID); err != nil {
		return fmt.Errorf("unable to decode inputPrivacyKeyAddedByPhone#d1219bdd: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InputPrivacyKeyAddedByPhone) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputPrivacyKeyAddedByPhone#d1219bdd to nil")
	}
	return nil
}

// InputPrivacyKeyVoiceMessages represents TL type `inputPrivacyKeyVoiceMessages#aee69d68`.
//
// See https://core.telegram.org/constructor/inputPrivacyKeyVoiceMessages for reference.
type InputPrivacyKeyVoiceMessages struct {
}

// InputPrivacyKeyVoiceMessagesTypeID is TL type id of InputPrivacyKeyVoiceMessages.
const InputPrivacyKeyVoiceMessagesTypeID = 0xaee69d68

// construct implements constructor of InputPrivacyKeyClass.
func (i InputPrivacyKeyVoiceMessages) construct() InputPrivacyKeyClass { return &i }

// Ensuring interfaces in compile-time for InputPrivacyKeyVoiceMessages.
var (
	_ bin.Encoder     = &InputPrivacyKeyVoiceMessages{}
	_ bin.Decoder     = &InputPrivacyKeyVoiceMessages{}
	_ bin.BareEncoder = &InputPrivacyKeyVoiceMessages{}
	_ bin.BareDecoder = &InputPrivacyKeyVoiceMessages{}

	_ InputPrivacyKeyClass = &InputPrivacyKeyVoiceMessages{}
)

func (i *InputPrivacyKeyVoiceMessages) Zero() bool {
	if i == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputPrivacyKeyVoiceMessages) String() string {
	if i == nil {
		return "InputPrivacyKeyVoiceMessages(nil)"
	}
	type Alias InputPrivacyKeyVoiceMessages
	return fmt.Sprintf("InputPrivacyKeyVoiceMessages%+v", Alias(*i))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InputPrivacyKeyVoiceMessages) TypeID() uint32 {
	return InputPrivacyKeyVoiceMessagesTypeID
}

// TypeName returns name of type in TL schema.
func (*InputPrivacyKeyVoiceMessages) TypeName() string {
	return "inputPrivacyKeyVoiceMessages"
}

// TypeInfo returns info about TL type.
func (i *InputPrivacyKeyVoiceMessages) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inputPrivacyKeyVoiceMessages",
		ID:   InputPrivacyKeyVoiceMessagesTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (i *InputPrivacyKeyVoiceMessages) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputPrivacyKeyVoiceMessages#aee69d68 as nil")
	}
	b.PutID(InputPrivacyKeyVoiceMessagesTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InputPrivacyKeyVoiceMessages) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputPrivacyKeyVoiceMessages#aee69d68 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (i *InputPrivacyKeyVoiceMessages) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputPrivacyKeyVoiceMessages#aee69d68 to nil")
	}
	if err := b.ConsumeID(InputPrivacyKeyVoiceMessagesTypeID); err != nil {
		return fmt.Errorf("unable to decode inputPrivacyKeyVoiceMessages#aee69d68: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InputPrivacyKeyVoiceMessages) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputPrivacyKeyVoiceMessages#aee69d68 to nil")
	}
	return nil
}

// InputPrivacyKeyClassName is schema name of InputPrivacyKeyClass.
const InputPrivacyKeyClassName = "InputPrivacyKey"

// InputPrivacyKeyClass represents InputPrivacyKey generic type.
//
// See https://core.telegram.org/type/InputPrivacyKey for reference.
//
// Example:
//
//	g, err := tg.DecodeInputPrivacyKey(buf)
//	if err != nil {
//	    panic(err)
//	}
//	switch v := g.(type) {
//	case *tg.InputPrivacyKeyStatusTimestamp: // inputPrivacyKeyStatusTimestamp#4f96cb18
//	case *tg.InputPrivacyKeyChatInvite: // inputPrivacyKeyChatInvite#bdfb0426
//	case *tg.InputPrivacyKeyPhoneCall: // inputPrivacyKeyPhoneCall#fabadc5f
//	case *tg.InputPrivacyKeyPhoneP2P: // inputPrivacyKeyPhoneP2P#db9e70d2
//	case *tg.InputPrivacyKeyForwards: // inputPrivacyKeyForwards#a4dd4c08
//	case *tg.InputPrivacyKeyProfilePhoto: // inputPrivacyKeyProfilePhoto#5719bacc
//	case *tg.InputPrivacyKeyPhoneNumber: // inputPrivacyKeyPhoneNumber#352dafa
//	case *tg.InputPrivacyKeyAddedByPhone: // inputPrivacyKeyAddedByPhone#d1219bdd
//	case *tg.InputPrivacyKeyVoiceMessages: // inputPrivacyKeyVoiceMessages#aee69d68
//	default: panic(v)
//	}
type InputPrivacyKeyClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() InputPrivacyKeyClass

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
}

// DecodeInputPrivacyKey implements binary de-serialization for InputPrivacyKeyClass.
func DecodeInputPrivacyKey(buf *bin.Buffer) (InputPrivacyKeyClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case InputPrivacyKeyStatusTimestampTypeID:
		// Decoding inputPrivacyKeyStatusTimestamp#4f96cb18.
		v := InputPrivacyKeyStatusTimestamp{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputPrivacyKeyClass: %w", err)
		}
		return &v, nil
	case InputPrivacyKeyChatInviteTypeID:
		// Decoding inputPrivacyKeyChatInvite#bdfb0426.
		v := InputPrivacyKeyChatInvite{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputPrivacyKeyClass: %w", err)
		}
		return &v, nil
	case InputPrivacyKeyPhoneCallTypeID:
		// Decoding inputPrivacyKeyPhoneCall#fabadc5f.
		v := InputPrivacyKeyPhoneCall{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputPrivacyKeyClass: %w", err)
		}
		return &v, nil
	case InputPrivacyKeyPhoneP2PTypeID:
		// Decoding inputPrivacyKeyPhoneP2P#db9e70d2.
		v := InputPrivacyKeyPhoneP2P{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputPrivacyKeyClass: %w", err)
		}
		return &v, nil
	case InputPrivacyKeyForwardsTypeID:
		// Decoding inputPrivacyKeyForwards#a4dd4c08.
		v := InputPrivacyKeyForwards{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputPrivacyKeyClass: %w", err)
		}
		return &v, nil
	case InputPrivacyKeyProfilePhotoTypeID:
		// Decoding inputPrivacyKeyProfilePhoto#5719bacc.
		v := InputPrivacyKeyProfilePhoto{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputPrivacyKeyClass: %w", err)
		}
		return &v, nil
	case InputPrivacyKeyPhoneNumberTypeID:
		// Decoding inputPrivacyKeyPhoneNumber#352dafa.
		v := InputPrivacyKeyPhoneNumber{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputPrivacyKeyClass: %w", err)
		}
		return &v, nil
	case InputPrivacyKeyAddedByPhoneTypeID:
		// Decoding inputPrivacyKeyAddedByPhone#d1219bdd.
		v := InputPrivacyKeyAddedByPhone{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputPrivacyKeyClass: %w", err)
		}
		return &v, nil
	case InputPrivacyKeyVoiceMessagesTypeID:
		// Decoding inputPrivacyKeyVoiceMessages#aee69d68.
		v := InputPrivacyKeyVoiceMessages{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputPrivacyKeyClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode InputPrivacyKeyClass: %w", bin.NewUnexpectedID(id))
	}
}

// InputPrivacyKey boxes the InputPrivacyKeyClass providing a helper.
type InputPrivacyKeyBox struct {
	InputPrivacyKey InputPrivacyKeyClass
}

// Decode implements bin.Decoder for InputPrivacyKeyBox.
func (b *InputPrivacyKeyBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode InputPrivacyKeyBox to nil")
	}
	v, err := DecodeInputPrivacyKey(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.InputPrivacyKey = v
	return nil
}

// Encode implements bin.Encode for InputPrivacyKeyBox.
func (b *InputPrivacyKeyBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.InputPrivacyKey == nil {
		return fmt.Errorf("unable to encode InputPrivacyKeyClass as nil")
	}
	return b.InputPrivacyKey.Encode(buf)
}
