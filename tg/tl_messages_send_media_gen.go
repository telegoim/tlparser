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

// MessagesSendMediaRequest represents TL type `messages.sendMedia#e25ff8e0`.
// Send a media
//
// See https://core.telegram.org/method/messages.sendMedia for reference.
type MessagesSendMediaRequest struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Send message silently (no notification should be triggered)
	Silent bool
	// Send message in background
	Background bool
	// Clear the draft
	ClearDraft bool
	// Only for bots, disallows forwarding and saving of the messages, even if the
	// destination chat doesn't have content protection¹ enabled
	//
	// Links:
	//  1) https://telegram.org/blog/protected-content-delete-by-date-and-more
	Noforwards bool
	// Destination
	Peer InputPeerClass
	// Message ID to which this message should reply to
	//
	// Use SetReplyToMsgID and GetReplyToMsgID helpers.
	ReplyToMsgID int
	// Attached media
	Media InputMediaClass
	// Caption
	Message string
	// Random ID to avoid resending the same message
	RandomID int64
	// Reply markup for bot keyboards
	//
	// Use SetReplyMarkup and GetReplyMarkup helpers.
	ReplyMarkup ReplyMarkupClass
	// Message entities¹ for styled text
	//
	// Links:
	//  1) https://core.telegram.org/api/entities
	//
	// Use SetEntities and GetEntities helpers.
	Entities []MessageEntityClass
	// Scheduled message date for scheduled messages¹
	//
	// Links:
	//  1) https://core.telegram.org/api/scheduled-messages
	//
	// Use SetScheduleDate and GetScheduleDate helpers.
	ScheduleDate int
	// Send this message as the specified peer
	//
	// Use SetSendAs and GetSendAs helpers.
	SendAs InputPeerClass
}

// MessagesSendMediaRequestTypeID is TL type id of MessagesSendMediaRequest.
const MessagesSendMediaRequestTypeID = 0xe25ff8e0

// Ensuring interfaces in compile-time for MessagesSendMediaRequest.
var (
	_ bin.Encoder     = &MessagesSendMediaRequest{}
	_ bin.Decoder     = &MessagesSendMediaRequest{}
	_ bin.BareEncoder = &MessagesSendMediaRequest{}
	_ bin.BareDecoder = &MessagesSendMediaRequest{}
)

func (s *MessagesSendMediaRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Flags.Zero()) {
		return false
	}
	if !(s.Silent == false) {
		return false
	}
	if !(s.Background == false) {
		return false
	}
	if !(s.ClearDraft == false) {
		return false
	}
	if !(s.Noforwards == false) {
		return false
	}
	if !(s.Peer == nil) {
		return false
	}
	if !(s.ReplyToMsgID == 0) {
		return false
	}
	if !(s.Media == nil) {
		return false
	}
	if !(s.Message == "") {
		return false
	}
	if !(s.RandomID == 0) {
		return false
	}
	if !(s.ReplyMarkup == nil) {
		return false
	}
	if !(s.Entities == nil) {
		return false
	}
	if !(s.ScheduleDate == 0) {
		return false
	}
	if !(s.SendAs == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *MessagesSendMediaRequest) String() string {
	if s == nil {
		return "MessagesSendMediaRequest(nil)"
	}
	type Alias MessagesSendMediaRequest
	return fmt.Sprintf("MessagesSendMediaRequest%+v", Alias(*s))
}

// FillFrom fills MessagesSendMediaRequest from given interface.
func (s *MessagesSendMediaRequest) FillFrom(from interface {
	GetSilent() (value bool)
	GetBackground() (value bool)
	GetClearDraft() (value bool)
	GetNoforwards() (value bool)
	GetPeer() (value InputPeerClass)
	GetReplyToMsgID() (value int, ok bool)
	GetMedia() (value InputMediaClass)
	GetMessage() (value string)
	GetRandomID() (value int64)
	GetReplyMarkup() (value ReplyMarkupClass, ok bool)
	GetEntities() (value []MessageEntityClass, ok bool)
	GetScheduleDate() (value int, ok bool)
	GetSendAs() (value InputPeerClass, ok bool)
}) {
	s.Silent = from.GetSilent()
	s.Background = from.GetBackground()
	s.ClearDraft = from.GetClearDraft()
	s.Noforwards = from.GetNoforwards()
	s.Peer = from.GetPeer()
	if val, ok := from.GetReplyToMsgID(); ok {
		s.ReplyToMsgID = val
	}

	s.Media = from.GetMedia()
	s.Message = from.GetMessage()
	s.RandomID = from.GetRandomID()
	if val, ok := from.GetReplyMarkup(); ok {
		s.ReplyMarkup = val
	}

	if val, ok := from.GetEntities(); ok {
		s.Entities = val
	}

	if val, ok := from.GetScheduleDate(); ok {
		s.ScheduleDate = val
	}

	if val, ok := from.GetSendAs(); ok {
		s.SendAs = val
	}

}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesSendMediaRequest) TypeID() uint32 {
	return MessagesSendMediaRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesSendMediaRequest) TypeName() string {
	return "messages.sendMedia"
}

// TypeInfo returns info about TL type.
func (s *MessagesSendMediaRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.sendMedia",
		ID:   MessagesSendMediaRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Silent",
			SchemaName: "silent",
			Null:       !s.Flags.Has(5),
		},
		{
			Name:       "Background",
			SchemaName: "background",
			Null:       !s.Flags.Has(6),
		},
		{
			Name:       "ClearDraft",
			SchemaName: "clear_draft",
			Null:       !s.Flags.Has(7),
		},
		{
			Name:       "Noforwards",
			SchemaName: "noforwards",
			Null:       !s.Flags.Has(14),
		},
		{
			Name:       "Peer",
			SchemaName: "peer",
		},
		{
			Name:       "ReplyToMsgID",
			SchemaName: "reply_to_msg_id",
			Null:       !s.Flags.Has(0),
		},
		{
			Name:       "Media",
			SchemaName: "media",
		},
		{
			Name:       "Message",
			SchemaName: "message",
		},
		{
			Name:       "RandomID",
			SchemaName: "random_id",
		},
		{
			Name:       "ReplyMarkup",
			SchemaName: "reply_markup",
			Null:       !s.Flags.Has(2),
		},
		{
			Name:       "Entities",
			SchemaName: "entities",
			Null:       !s.Flags.Has(3),
		},
		{
			Name:       "ScheduleDate",
			SchemaName: "schedule_date",
			Null:       !s.Flags.Has(10),
		},
		{
			Name:       "SendAs",
			SchemaName: "send_as",
			Null:       !s.Flags.Has(13),
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (s *MessagesSendMediaRequest) SetFlags() {
	if !(s.Silent == false) {
		s.Flags.Set(5)
	}
	if !(s.Background == false) {
		s.Flags.Set(6)
	}
	if !(s.ClearDraft == false) {
		s.Flags.Set(7)
	}
	if !(s.Noforwards == false) {
		s.Flags.Set(14)
	}
	if !(s.ReplyToMsgID == 0) {
		s.Flags.Set(0)
	}
	if !(s.ReplyMarkup == nil) {
		s.Flags.Set(2)
	}
	if !(s.Entities == nil) {
		s.Flags.Set(3)
	}
	if !(s.ScheduleDate == 0) {
		s.Flags.Set(10)
	}
	if !(s.SendAs == nil) {
		s.Flags.Set(13)
	}
}

// Encode implements bin.Encoder.
func (s *MessagesSendMediaRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode messages.sendMedia#e25ff8e0 as nil")
	}
	b.PutID(MessagesSendMediaRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *MessagesSendMediaRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode messages.sendMedia#e25ff8e0 as nil")
	}
	s.SetFlags()
	if err := s.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.sendMedia#e25ff8e0: field flags: %w", err)
	}
	if s.Peer == nil {
		return fmt.Errorf("unable to encode messages.sendMedia#e25ff8e0: field peer is nil")
	}
	if err := s.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.sendMedia#e25ff8e0: field peer: %w", err)
	}
	if s.Flags.Has(0) {
		b.PutInt(s.ReplyToMsgID)
	}
	if s.Media == nil {
		return fmt.Errorf("unable to encode messages.sendMedia#e25ff8e0: field media is nil")
	}
	if err := s.Media.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.sendMedia#e25ff8e0: field media: %w", err)
	}
	b.PutString(s.Message)
	b.PutLong(s.RandomID)
	if s.Flags.Has(2) {
		if s.ReplyMarkup == nil {
			return fmt.Errorf("unable to encode messages.sendMedia#e25ff8e0: field reply_markup is nil")
		}
		if err := s.ReplyMarkup.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.sendMedia#e25ff8e0: field reply_markup: %w", err)
		}
	}
	if s.Flags.Has(3) {
		b.PutVectorHeader(len(s.Entities))
		for idx, v := range s.Entities {
			if v == nil {
				return fmt.Errorf("unable to encode messages.sendMedia#e25ff8e0: field entities element with index %d is nil", idx)
			}
			if err := v.Encode(b); err != nil {
				return fmt.Errorf("unable to encode messages.sendMedia#e25ff8e0: field entities element with index %d: %w", idx, err)
			}
		}
	}
	if s.Flags.Has(10) {
		b.PutInt(s.ScheduleDate)
	}
	if s.Flags.Has(13) {
		if s.SendAs == nil {
			return fmt.Errorf("unable to encode messages.sendMedia#e25ff8e0: field send_as is nil")
		}
		if err := s.SendAs.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.sendMedia#e25ff8e0: field send_as: %w", err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *MessagesSendMediaRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode messages.sendMedia#e25ff8e0 to nil")
	}
	if err := b.ConsumeID(MessagesSendMediaRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.sendMedia#e25ff8e0: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *MessagesSendMediaRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode messages.sendMedia#e25ff8e0 to nil")
	}
	{
		if err := s.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messages.sendMedia#e25ff8e0: field flags: %w", err)
		}
	}
	s.Silent = s.Flags.Has(5)
	s.Background = s.Flags.Has(6)
	s.ClearDraft = s.Flags.Has(7)
	s.Noforwards = s.Flags.Has(14)
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.sendMedia#e25ff8e0: field peer: %w", err)
		}
		s.Peer = value
	}
	if s.Flags.Has(0) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.sendMedia#e25ff8e0: field reply_to_msg_id: %w", err)
		}
		s.ReplyToMsgID = value
	}
	{
		value, err := DecodeInputMedia(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.sendMedia#e25ff8e0: field media: %w", err)
		}
		s.Media = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode messages.sendMedia#e25ff8e0: field message: %w", err)
		}
		s.Message = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode messages.sendMedia#e25ff8e0: field random_id: %w", err)
		}
		s.RandomID = value
	}
	if s.Flags.Has(2) {
		value, err := DecodeReplyMarkup(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.sendMedia#e25ff8e0: field reply_markup: %w", err)
		}
		s.ReplyMarkup = value
	}
	if s.Flags.Has(3) {
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.sendMedia#e25ff8e0: field entities: %w", err)
		}

		if headerLen > 0 {
			s.Entities = make([]MessageEntityClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeMessageEntity(b)
			if err != nil {
				return fmt.Errorf("unable to decode messages.sendMedia#e25ff8e0: field entities: %w", err)
			}
			s.Entities = append(s.Entities, value)
		}
	}
	if s.Flags.Has(10) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.sendMedia#e25ff8e0: field schedule_date: %w", err)
		}
		s.ScheduleDate = value
	}
	if s.Flags.Has(13) {
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.sendMedia#e25ff8e0: field send_as: %w", err)
		}
		s.SendAs = value
	}
	return nil
}

// SetSilent sets value of Silent conditional field.
func (s *MessagesSendMediaRequest) SetSilent(value bool) {
	if value {
		s.Flags.Set(5)
		s.Silent = true
	} else {
		s.Flags.Unset(5)
		s.Silent = false
	}
}

// GetSilent returns value of Silent conditional field.
func (s *MessagesSendMediaRequest) GetSilent() (value bool) {
	if s == nil {
		return
	}
	return s.Flags.Has(5)
}

// SetBackground sets value of Background conditional field.
func (s *MessagesSendMediaRequest) SetBackground(value bool) {
	if value {
		s.Flags.Set(6)
		s.Background = true
	} else {
		s.Flags.Unset(6)
		s.Background = false
	}
}

// GetBackground returns value of Background conditional field.
func (s *MessagesSendMediaRequest) GetBackground() (value bool) {
	if s == nil {
		return
	}
	return s.Flags.Has(6)
}

// SetClearDraft sets value of ClearDraft conditional field.
func (s *MessagesSendMediaRequest) SetClearDraft(value bool) {
	if value {
		s.Flags.Set(7)
		s.ClearDraft = true
	} else {
		s.Flags.Unset(7)
		s.ClearDraft = false
	}
}

// GetClearDraft returns value of ClearDraft conditional field.
func (s *MessagesSendMediaRequest) GetClearDraft() (value bool) {
	if s == nil {
		return
	}
	return s.Flags.Has(7)
}

// SetNoforwards sets value of Noforwards conditional field.
func (s *MessagesSendMediaRequest) SetNoforwards(value bool) {
	if value {
		s.Flags.Set(14)
		s.Noforwards = true
	} else {
		s.Flags.Unset(14)
		s.Noforwards = false
	}
}

// GetNoforwards returns value of Noforwards conditional field.
func (s *MessagesSendMediaRequest) GetNoforwards() (value bool) {
	if s == nil {
		return
	}
	return s.Flags.Has(14)
}

// GetPeer returns value of Peer field.
func (s *MessagesSendMediaRequest) GetPeer() (value InputPeerClass) {
	if s == nil {
		return
	}
	return s.Peer
}

// SetReplyToMsgID sets value of ReplyToMsgID conditional field.
func (s *MessagesSendMediaRequest) SetReplyToMsgID(value int) {
	s.Flags.Set(0)
	s.ReplyToMsgID = value
}

// GetReplyToMsgID returns value of ReplyToMsgID conditional field and
// boolean which is true if field was set.
func (s *MessagesSendMediaRequest) GetReplyToMsgID() (value int, ok bool) {
	if s == nil {
		return
	}
	if !s.Flags.Has(0) {
		return value, false
	}
	return s.ReplyToMsgID, true
}

// GetMedia returns value of Media field.
func (s *MessagesSendMediaRequest) GetMedia() (value InputMediaClass) {
	if s == nil {
		return
	}
	return s.Media
}

// GetMessage returns value of Message field.
func (s *MessagesSendMediaRequest) GetMessage() (value string) {
	if s == nil {
		return
	}
	return s.Message
}

// GetRandomID returns value of RandomID field.
func (s *MessagesSendMediaRequest) GetRandomID() (value int64) {
	if s == nil {
		return
	}
	return s.RandomID
}

// SetReplyMarkup sets value of ReplyMarkup conditional field.
func (s *MessagesSendMediaRequest) SetReplyMarkup(value ReplyMarkupClass) {
	s.Flags.Set(2)
	s.ReplyMarkup = value
}

// GetReplyMarkup returns value of ReplyMarkup conditional field and
// boolean which is true if field was set.
func (s *MessagesSendMediaRequest) GetReplyMarkup() (value ReplyMarkupClass, ok bool) {
	if s == nil {
		return
	}
	if !s.Flags.Has(2) {
		return value, false
	}
	return s.ReplyMarkup, true
}

// SetEntities sets value of Entities conditional field.
func (s *MessagesSendMediaRequest) SetEntities(value []MessageEntityClass) {
	s.Flags.Set(3)
	s.Entities = value
}

// GetEntities returns value of Entities conditional field and
// boolean which is true if field was set.
func (s *MessagesSendMediaRequest) GetEntities() (value []MessageEntityClass, ok bool) {
	if s == nil {
		return
	}
	if !s.Flags.Has(3) {
		return value, false
	}
	return s.Entities, true
}

// SetScheduleDate sets value of ScheduleDate conditional field.
func (s *MessagesSendMediaRequest) SetScheduleDate(value int) {
	s.Flags.Set(10)
	s.ScheduleDate = value
}

// GetScheduleDate returns value of ScheduleDate conditional field and
// boolean which is true if field was set.
func (s *MessagesSendMediaRequest) GetScheduleDate() (value int, ok bool) {
	if s == nil {
		return
	}
	if !s.Flags.Has(10) {
		return value, false
	}
	return s.ScheduleDate, true
}

// SetSendAs sets value of SendAs conditional field.
func (s *MessagesSendMediaRequest) SetSendAs(value InputPeerClass) {
	s.Flags.Set(13)
	s.SendAs = value
}

// GetSendAs returns value of SendAs conditional field and
// boolean which is true if field was set.
func (s *MessagesSendMediaRequest) GetSendAs() (value InputPeerClass, ok bool) {
	if s == nil {
		return
	}
	if !s.Flags.Has(13) {
		return value, false
	}
	return s.SendAs, true
}

// MapEntities returns field Entities wrapped in MessageEntityClassArray helper.
func (s *MessagesSendMediaRequest) MapEntities() (value MessageEntityClassArray, ok bool) {
	if !s.Flags.Has(3) {
		return value, false
	}
	return MessageEntityClassArray(s.Entities), true
}

// MessagesSendMedia invokes method messages.sendMedia#e25ff8e0 returning error if any.
// Send a media
//
// Possible errors:
//
//	400 BOT_PAYMENTS_DISABLED: Please enable bot payments in botfather before calling this method.
//	400 BROADCAST_PUBLIC_VOTERS_FORBIDDEN: You can't forward polls with public voters.
//	400 BUTTON_DATA_INVALID: The data of one or more of the buttons you provided is invalid.
//	400 BUTTON_TYPE_INVALID: The type of one or more of the buttons you provided is invalid.
//	400 BUTTON_URL_INVALID: Button URL invalid.
//	400 CHANNEL_INVALID: The provided channel is invalid.
//	400 CHANNEL_PRIVATE: You haven't joined this channel/supergroup.
//	400 CHAT_ADMIN_REQUIRED: You must be an admin in this chat to do this.
//	400 CHAT_FORWARDS_RESTRICTED: You can't forward messages from a protected chat.
//	400 CHAT_RESTRICTED: You can't send messages in this chat, you were restricted.
//	403 CHAT_SEND_GIFS_FORBIDDEN: You can't send gifs in this chat.
//	403 CHAT_SEND_MEDIA_FORBIDDEN: You can't send media in this chat.
//	403 CHAT_SEND_POLL_FORBIDDEN: You can't send polls in this chat.
//	403 CHAT_SEND_STICKERS_FORBIDDEN: You can't send stickers in this chat.
//	403 CHAT_WRITE_FORBIDDEN: You can't write in this chat.
//	400 CURRENCY_TOTAL_AMOUNT_INVALID: The total amount of all prices is invalid.
//	400 EMOTICON_INVALID: The specified emoji is invalid.
//	400 EXTERNAL_URL_INVALID: External URL invalid.
//	400 FILE_PARTS_INVALID: The number of file parts is invalid.
//	400 FILE_PART_LENGTH_INVALID: The length of a file part is invalid.
//	400 FILE_REFERENCE_EMPTY: An empty file reference was specified.
//	400 FILE_REFERENCE_EXPIRED: File reference expired, it must be refetched as described in the documentation.
//	400 GAME_BOT_INVALID: Bots can't send another bot's game.
//	400 IMAGE_PROCESS_FAILED: Failure while processing image.
//	400 INPUT_USER_DEACTIVATED: The specified user was deleted.
//	400 MD5_CHECKSUM_INVALID: The MD5 checksums do not match.
//	400 MEDIA_CAPTION_TOO_LONG: The caption is too long.
//	400 MEDIA_EMPTY: The provided media object is invalid.
//	400 MEDIA_INVALID: Media invalid.
//	400 MSG_ID_INVALID: Invalid message ID provided.
//	400 PAYMENT_PROVIDER_INVALID: The specified payment provider is invalid.
//	400 PEER_ID_INVALID: The provided peer id is invalid.
//	400 PHOTO_EXT_INVALID: The extension of the photo is invalid.
//	400 PHOTO_INVALID_DIMENSIONS: The photo dimensions are invalid.
//	400 PHOTO_SAVE_FILE_INVALID: Internal issues, try again later.
//	400 POLL_ANSWERS_INVALID: Invalid poll answers were provided.
//	400 POLL_ANSWER_INVALID: One of the poll answers is not acceptable.
//	400 POLL_OPTION_DUPLICATE: Duplicate poll options provided.
//	400 POLL_OPTION_INVALID: Invalid poll option provided.
//	400 POLL_QUESTION_INVALID: One of the poll questions is not acceptable.
//	400 QUIZ_CORRECT_ANSWERS_EMPTY: No correct quiz answer was specified.
//	400 QUIZ_CORRECT_ANSWERS_TOO_MUCH: You specified too many correct answers in a quiz, quizzes can only have one right answer!
//	400 QUIZ_CORRECT_ANSWER_INVALID: An invalid value was provided to the correct_answers field.
//	400 QUIZ_MULTIPLE_INVALID: Quizzes can't have the multiple_choice flag set!
//	500 RANDOM_ID_DUPLICATE: You provided a random ID that was already used.
//	400 REPLY_MARKUP_BUY_EMPTY: Reply markup for buy button empty.
//	400 REPLY_MARKUP_INVALID: The provided reply markup is invalid.
//	400 SCHEDULE_BOT_NOT_ALLOWED: Bots cannot schedule messages.
//	400 SCHEDULE_DATE_TOO_LATE: You can't schedule a message this far in the future.
//	400 SCHEDULE_TOO_MUCH: There are too many scheduled messages.
//	400 SEND_AS_PEER_INVALID: You can't send messages as the specified peer.
//	420 SLOWMODE_WAIT_%d: Slowmode is enabled in this chat: wait %d seconds before sending another message to this chat.
//	400 TTL_MEDIA_INVALID: Invalid media Time To Live was provided.
//	400 USER_BANNED_IN_CHANNEL: You're banned from sending messages in supergroups/channels.
//	403 USER_IS_BLOCKED: You were blocked by this user.
//	400 USER_IS_BOT: Bots can't send messages to other bots.
//	400 VIDEO_CONTENT_TYPE_INVALID: The video's content type is invalid.
//	400 WEBPAGE_CURL_FAILED: Failure while fetching the webpage with cURL.
//	400 WEBPAGE_MEDIA_EMPTY: Webpage media empty.
//	400 YOU_BLOCKED_USER: You blocked this user.
//
// See https://core.telegram.org/method/messages.sendMedia for reference.
// Can be used by bots.
func (c *Client) MessagesSendMedia(ctx context.Context, request *MessagesSendMediaRequest) (UpdatesClass, error) {
	var result UpdatesBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Updates, nil
}
