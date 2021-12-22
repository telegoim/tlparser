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

// Supergroup represents TL type `supergroup#d4f3e735`.
type Supergroup struct {
	// Supergroup or channel identifier
	ID int64
	// Username of the supergroup or channel; empty for private supergroups or channels
	Username string
	// Point in time (Unix timestamp) when the current user joined, or the point in time when
	// the supergroup or channel was created, in case the user is not a member
	Date int32
	// Status of the current user in the supergroup or channel; custom title will be always
	// empty
	Status ChatMemberStatusClass
	// Number of members in the supergroup or channel; 0 if unknown. Currently, it is
	// guaranteed to be known only if the supergroup or channel was received through
	// searchPublicChats, searchChatsNearby, getInactiveSupergroupChats,
	// getSuitableDiscussionChats, getGroupsInCommon, or getUserPrivacySettingRules
	MemberCount int32
	// True, if the channel has a discussion group, or the supergroup is the designated
	// discussion group for a channel
	HasLinkedChat bool
	// True, if the supergroup is connected to a location, i.e. the supergroup is a
	// location-based supergroup
	HasLocation bool
	// True, if messages sent to the channel need to contain information about the sender.
	// This field is only applicable to channels
	SignMessages bool
	// True, if the slow mode is enabled in the supergroup
	IsSlowModeEnabled bool
	// True, if the supergroup is a channel
	IsChannel bool
	// True, if the supergroup is a broadcast group, i.e. only administrators can send
	// messages and there is no limit on number of members
	IsBroadcastGroup bool
	// True, if the supergroup or channel is verified
	IsVerified bool
	// If non-empty, contains a human-readable description of the reason why access to this
	// supergroup or channel must be restricted
	RestrictionReason string
	// True, if many users reported this supergroup or channel as a scam
	IsScam bool
	// True, if many users reported this supergroup or channel as a fake account
	IsFake bool
}

// SupergroupTypeID is TL type id of Supergroup.
const SupergroupTypeID = 0xd4f3e735

// Ensuring interfaces in compile-time for Supergroup.
var (
	_ bin.Encoder     = &Supergroup{}
	_ bin.Decoder     = &Supergroup{}
	_ bin.BareEncoder = &Supergroup{}
	_ bin.BareDecoder = &Supergroup{}
)

func (s *Supergroup) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.ID == 0) {
		return false
	}
	if !(s.Username == "") {
		return false
	}
	if !(s.Date == 0) {
		return false
	}
	if !(s.Status == nil) {
		return false
	}
	if !(s.MemberCount == 0) {
		return false
	}
	if !(s.HasLinkedChat == false) {
		return false
	}
	if !(s.HasLocation == false) {
		return false
	}
	if !(s.SignMessages == false) {
		return false
	}
	if !(s.IsSlowModeEnabled == false) {
		return false
	}
	if !(s.IsChannel == false) {
		return false
	}
	if !(s.IsBroadcastGroup == false) {
		return false
	}
	if !(s.IsVerified == false) {
		return false
	}
	if !(s.RestrictionReason == "") {
		return false
	}
	if !(s.IsScam == false) {
		return false
	}
	if !(s.IsFake == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *Supergroup) String() string {
	if s == nil {
		return "Supergroup(nil)"
	}
	type Alias Supergroup
	return fmt.Sprintf("Supergroup%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*Supergroup) TypeID() uint32 {
	return SupergroupTypeID
}

// TypeName returns name of type in TL schema.
func (*Supergroup) TypeName() string {
	return "supergroup"
}

// TypeInfo returns info about TL type.
func (s *Supergroup) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "supergroup",
		ID:   SupergroupTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ID",
			SchemaName: "id",
		},
		{
			Name:       "Username",
			SchemaName: "username",
		},
		{
			Name:       "Date",
			SchemaName: "date",
		},
		{
			Name:       "Status",
			SchemaName: "status",
		},
		{
			Name:       "MemberCount",
			SchemaName: "member_count",
		},
		{
			Name:       "HasLinkedChat",
			SchemaName: "has_linked_chat",
		},
		{
			Name:       "HasLocation",
			SchemaName: "has_location",
		},
		{
			Name:       "SignMessages",
			SchemaName: "sign_messages",
		},
		{
			Name:       "IsSlowModeEnabled",
			SchemaName: "is_slow_mode_enabled",
		},
		{
			Name:       "IsChannel",
			SchemaName: "is_channel",
		},
		{
			Name:       "IsBroadcastGroup",
			SchemaName: "is_broadcast_group",
		},
		{
			Name:       "IsVerified",
			SchemaName: "is_verified",
		},
		{
			Name:       "RestrictionReason",
			SchemaName: "restriction_reason",
		},
		{
			Name:       "IsScam",
			SchemaName: "is_scam",
		},
		{
			Name:       "IsFake",
			SchemaName: "is_fake",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *Supergroup) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode supergroup#d4f3e735 as nil")
	}
	b.PutID(SupergroupTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *Supergroup) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode supergroup#d4f3e735 as nil")
	}
	b.PutInt53(s.ID)
	b.PutString(s.Username)
	b.PutInt32(s.Date)
	if s.Status == nil {
		return fmt.Errorf("unable to encode supergroup#d4f3e735: field status is nil")
	}
	if err := s.Status.Encode(b); err != nil {
		return fmt.Errorf("unable to encode supergroup#d4f3e735: field status: %w", err)
	}
	b.PutInt32(s.MemberCount)
	b.PutBool(s.HasLinkedChat)
	b.PutBool(s.HasLocation)
	b.PutBool(s.SignMessages)
	b.PutBool(s.IsSlowModeEnabled)
	b.PutBool(s.IsChannel)
	b.PutBool(s.IsBroadcastGroup)
	b.PutBool(s.IsVerified)
	b.PutString(s.RestrictionReason)
	b.PutBool(s.IsScam)
	b.PutBool(s.IsFake)
	return nil
}

// Decode implements bin.Decoder.
func (s *Supergroup) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode supergroup#d4f3e735 to nil")
	}
	if err := b.ConsumeID(SupergroupTypeID); err != nil {
		return fmt.Errorf("unable to decode supergroup#d4f3e735: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *Supergroup) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode supergroup#d4f3e735 to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode supergroup#d4f3e735: field id: %w", err)
		}
		s.ID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode supergroup#d4f3e735: field username: %w", err)
		}
		s.Username = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode supergroup#d4f3e735: field date: %w", err)
		}
		s.Date = value
	}
	{
		value, err := DecodeChatMemberStatus(b)
		if err != nil {
			return fmt.Errorf("unable to decode supergroup#d4f3e735: field status: %w", err)
		}
		s.Status = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode supergroup#d4f3e735: field member_count: %w", err)
		}
		s.MemberCount = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode supergroup#d4f3e735: field has_linked_chat: %w", err)
		}
		s.HasLinkedChat = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode supergroup#d4f3e735: field has_location: %w", err)
		}
		s.HasLocation = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode supergroup#d4f3e735: field sign_messages: %w", err)
		}
		s.SignMessages = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode supergroup#d4f3e735: field is_slow_mode_enabled: %w", err)
		}
		s.IsSlowModeEnabled = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode supergroup#d4f3e735: field is_channel: %w", err)
		}
		s.IsChannel = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode supergroup#d4f3e735: field is_broadcast_group: %w", err)
		}
		s.IsBroadcastGroup = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode supergroup#d4f3e735: field is_verified: %w", err)
		}
		s.IsVerified = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode supergroup#d4f3e735: field restriction_reason: %w", err)
		}
		s.RestrictionReason = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode supergroup#d4f3e735: field is_scam: %w", err)
		}
		s.IsScam = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode supergroup#d4f3e735: field is_fake: %w", err)
		}
		s.IsFake = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (s *Supergroup) EncodeTDLibJSON(b tdjson.Encoder) error {
	if s == nil {
		return fmt.Errorf("can't encode supergroup#d4f3e735 as nil")
	}
	b.ObjStart()
	b.PutID("supergroup")
	b.FieldStart("id")
	b.PutInt53(s.ID)
	b.FieldStart("username")
	b.PutString(s.Username)
	b.FieldStart("date")
	b.PutInt32(s.Date)
	b.FieldStart("status")
	if s.Status == nil {
		return fmt.Errorf("unable to encode supergroup#d4f3e735: field status is nil")
	}
	if err := s.Status.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode supergroup#d4f3e735: field status: %w", err)
	}
	b.FieldStart("member_count")
	b.PutInt32(s.MemberCount)
	b.FieldStart("has_linked_chat")
	b.PutBool(s.HasLinkedChat)
	b.FieldStart("has_location")
	b.PutBool(s.HasLocation)
	b.FieldStart("sign_messages")
	b.PutBool(s.SignMessages)
	b.FieldStart("is_slow_mode_enabled")
	b.PutBool(s.IsSlowModeEnabled)
	b.FieldStart("is_channel")
	b.PutBool(s.IsChannel)
	b.FieldStart("is_broadcast_group")
	b.PutBool(s.IsBroadcastGroup)
	b.FieldStart("is_verified")
	b.PutBool(s.IsVerified)
	b.FieldStart("restriction_reason")
	b.PutString(s.RestrictionReason)
	b.FieldStart("is_scam")
	b.PutBool(s.IsScam)
	b.FieldStart("is_fake")
	b.PutBool(s.IsFake)
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (s *Supergroup) DecodeTDLibJSON(b tdjson.Decoder) error {
	if s == nil {
		return fmt.Errorf("can't decode supergroup#d4f3e735 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("supergroup"); err != nil {
				return fmt.Errorf("unable to decode supergroup#d4f3e735: %w", err)
			}
		case "id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode supergroup#d4f3e735: field id: %w", err)
			}
			s.ID = value
		case "username":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode supergroup#d4f3e735: field username: %w", err)
			}
			s.Username = value
		case "date":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode supergroup#d4f3e735: field date: %w", err)
			}
			s.Date = value
		case "status":
			value, err := DecodeTDLibJSONChatMemberStatus(b)
			if err != nil {
				return fmt.Errorf("unable to decode supergroup#d4f3e735: field status: %w", err)
			}
			s.Status = value
		case "member_count":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode supergroup#d4f3e735: field member_count: %w", err)
			}
			s.MemberCount = value
		case "has_linked_chat":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode supergroup#d4f3e735: field has_linked_chat: %w", err)
			}
			s.HasLinkedChat = value
		case "has_location":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode supergroup#d4f3e735: field has_location: %w", err)
			}
			s.HasLocation = value
		case "sign_messages":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode supergroup#d4f3e735: field sign_messages: %w", err)
			}
			s.SignMessages = value
		case "is_slow_mode_enabled":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode supergroup#d4f3e735: field is_slow_mode_enabled: %w", err)
			}
			s.IsSlowModeEnabled = value
		case "is_channel":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode supergroup#d4f3e735: field is_channel: %w", err)
			}
			s.IsChannel = value
		case "is_broadcast_group":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode supergroup#d4f3e735: field is_broadcast_group: %w", err)
			}
			s.IsBroadcastGroup = value
		case "is_verified":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode supergroup#d4f3e735: field is_verified: %w", err)
			}
			s.IsVerified = value
		case "restriction_reason":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode supergroup#d4f3e735: field restriction_reason: %w", err)
			}
			s.RestrictionReason = value
		case "is_scam":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode supergroup#d4f3e735: field is_scam: %w", err)
			}
			s.IsScam = value
		case "is_fake":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode supergroup#d4f3e735: field is_fake: %w", err)
			}
			s.IsFake = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetID returns value of ID field.
func (s *Supergroup) GetID() (value int64) {
	if s == nil {
		return
	}
	return s.ID
}

// GetUsername returns value of Username field.
func (s *Supergroup) GetUsername() (value string) {
	if s == nil {
		return
	}
	return s.Username
}

// GetDate returns value of Date field.
func (s *Supergroup) GetDate() (value int32) {
	if s == nil {
		return
	}
	return s.Date
}

// GetStatus returns value of Status field.
func (s *Supergroup) GetStatus() (value ChatMemberStatusClass) {
	if s == nil {
		return
	}
	return s.Status
}

// GetMemberCount returns value of MemberCount field.
func (s *Supergroup) GetMemberCount() (value int32) {
	if s == nil {
		return
	}
	return s.MemberCount
}

// GetHasLinkedChat returns value of HasLinkedChat field.
func (s *Supergroup) GetHasLinkedChat() (value bool) {
	if s == nil {
		return
	}
	return s.HasLinkedChat
}

// GetHasLocation returns value of HasLocation field.
func (s *Supergroup) GetHasLocation() (value bool) {
	if s == nil {
		return
	}
	return s.HasLocation
}

// GetSignMessages returns value of SignMessages field.
func (s *Supergroup) GetSignMessages() (value bool) {
	if s == nil {
		return
	}
	return s.SignMessages
}

// GetIsSlowModeEnabled returns value of IsSlowModeEnabled field.
func (s *Supergroup) GetIsSlowModeEnabled() (value bool) {
	if s == nil {
		return
	}
	return s.IsSlowModeEnabled
}

// GetIsChannel returns value of IsChannel field.
func (s *Supergroup) GetIsChannel() (value bool) {
	if s == nil {
		return
	}
	return s.IsChannel
}

// GetIsBroadcastGroup returns value of IsBroadcastGroup field.
func (s *Supergroup) GetIsBroadcastGroup() (value bool) {
	if s == nil {
		return
	}
	return s.IsBroadcastGroup
}

// GetIsVerified returns value of IsVerified field.
func (s *Supergroup) GetIsVerified() (value bool) {
	if s == nil {
		return
	}
	return s.IsVerified
}

// GetRestrictionReason returns value of RestrictionReason field.
func (s *Supergroup) GetRestrictionReason() (value string) {
	if s == nil {
		return
	}
	return s.RestrictionReason
}

// GetIsScam returns value of IsScam field.
func (s *Supergroup) GetIsScam() (value bool) {
	if s == nil {
		return
	}
	return s.IsScam
}

// GetIsFake returns value of IsFake field.
func (s *Supergroup) GetIsFake() (value bool) {
	if s == nil {
		return
	}
	return s.IsFake
}
