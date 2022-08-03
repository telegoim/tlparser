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

// MessagesReportRequest represents TL type `messages.report#8953ab4e`.
// Report a message in a chat for violation of telegram's Terms of Service
//
// See https://core.telegram.org/method/messages.report for reference.
type MessagesReportRequest struct {
	// Peer
	Peer InputPeerClass
	// IDs of messages to report
	ID []int
	// Why are these messages being reported
	Reason ReportReasonClass
	// Comment for report moderation
	Message string
}

// MessagesReportRequestTypeID is TL type id of MessagesReportRequest.
const MessagesReportRequestTypeID = 0x8953ab4e

// Ensuring interfaces in compile-time for MessagesReportRequest.
var (
	_ bin.Encoder     = &MessagesReportRequest{}
	_ bin.Decoder     = &MessagesReportRequest{}
	_ bin.BareEncoder = &MessagesReportRequest{}
	_ bin.BareDecoder = &MessagesReportRequest{}
)

func (r *MessagesReportRequest) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.Peer == nil) {
		return false
	}
	if !(r.ID == nil) {
		return false
	}
	if !(r.Reason == nil) {
		return false
	}
	if !(r.Message == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *MessagesReportRequest) String() string {
	if r == nil {
		return "MessagesReportRequest(nil)"
	}
	type Alias MessagesReportRequest
	return fmt.Sprintf("MessagesReportRequest%+v", Alias(*r))
}

// FillFrom fills MessagesReportRequest from given interface.
func (r *MessagesReportRequest) FillFrom(from interface {
	GetPeer() (value InputPeerClass)
	GetID() (value []int)
	GetReason() (value ReportReasonClass)
	GetMessage() (value string)
}) {
	r.Peer = from.GetPeer()
	r.ID = from.GetID()
	r.Reason = from.GetReason()
	r.Message = from.GetMessage()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesReportRequest) TypeID() uint32 {
	return MessagesReportRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesReportRequest) TypeName() string {
	return "messages.report"
}

// TypeInfo returns info about TL type.
func (r *MessagesReportRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.report",
		ID:   MessagesReportRequestTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Peer",
			SchemaName: "peer",
		},
		{
			Name:       "ID",
			SchemaName: "id",
		},
		{
			Name:       "Reason",
			SchemaName: "reason",
		},
		{
			Name:       "Message",
			SchemaName: "message",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (r *MessagesReportRequest) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode messages.report#8953ab4e as nil")
	}
	b.PutID(MessagesReportRequestTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *MessagesReportRequest) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode messages.report#8953ab4e as nil")
	}
	if r.Peer == nil {
		return fmt.Errorf("unable to encode messages.report#8953ab4e: field peer is nil")
	}
	if err := r.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.report#8953ab4e: field peer: %w", err)
	}
	b.PutVectorHeader(len(r.ID))
	for _, v := range r.ID {
		b.PutInt(v)
	}
	if r.Reason == nil {
		return fmt.Errorf("unable to encode messages.report#8953ab4e: field reason is nil")
	}
	if err := r.Reason.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.report#8953ab4e: field reason: %w", err)
	}
	b.PutString(r.Message)
	return nil
}

// Decode implements bin.Decoder.
func (r *MessagesReportRequest) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode messages.report#8953ab4e to nil")
	}
	if err := b.ConsumeID(MessagesReportRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.report#8953ab4e: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *MessagesReportRequest) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode messages.report#8953ab4e to nil")
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.report#8953ab4e: field peer: %w", err)
		}
		r.Peer = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.report#8953ab4e: field id: %w", err)
		}

		if headerLen > 0 {
			r.ID = make([]int, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Int()
			if err != nil {
				return fmt.Errorf("unable to decode messages.report#8953ab4e: field id: %w", err)
			}
			r.ID = append(r.ID, value)
		}
	}
	{
		value, err := DecodeReportReason(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.report#8953ab4e: field reason: %w", err)
		}
		r.Reason = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode messages.report#8953ab4e: field message: %w", err)
		}
		r.Message = value
	}
	return nil
}

// GetPeer returns value of Peer field.
func (r *MessagesReportRequest) GetPeer() (value InputPeerClass) {
	if r == nil {
		return
	}
	return r.Peer
}

// GetID returns value of ID field.
func (r *MessagesReportRequest) GetID() (value []int) {
	if r == nil {
		return
	}
	return r.ID
}

// GetReason returns value of Reason field.
func (r *MessagesReportRequest) GetReason() (value ReportReasonClass) {
	if r == nil {
		return
	}
	return r.Reason
}

// GetMessage returns value of Message field.
func (r *MessagesReportRequest) GetMessage() (value string) {
	if r == nil {
		return
	}
	return r.Message
}

// MessagesReport invokes method messages.report#8953ab4e returning error if any.
// Report a message in a chat for violation of telegram's Terms of Service
//
// Possible errors:
//
//	400 CHANNEL_PRIVATE: You haven't joined this channel/supergroup.
//	400 PEER_ID_INVALID: The provided peer id is invalid.
//
// See https://core.telegram.org/method/messages.report for reference.
func (c *Client) MessagesReport(ctx context.Context, request *MessagesReportRequest) (bool, error) {
	var result BoolBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
