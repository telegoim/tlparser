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

// MessagesHideAllChatJoinRequestsRequest represents TL type `messages.hideAllChatJoinRequests#e085f4ea`.
//
// See https://core.telegram.org/method/messages.hideAllChatJoinRequests for reference.
type MessagesHideAllChatJoinRequestsRequest struct {
	// Flags field of MessagesHideAllChatJoinRequestsRequest.
	Flags bin.Fields
	// Approved field of MessagesHideAllChatJoinRequestsRequest.
	Approved bool
	// Peer field of MessagesHideAllChatJoinRequestsRequest.
	Peer InputPeerClass
	// Link field of MessagesHideAllChatJoinRequestsRequest.
	//
	// Use SetLink and GetLink helpers.
	Link string
}

// MessagesHideAllChatJoinRequestsRequestTypeID is TL type id of MessagesHideAllChatJoinRequestsRequest.
const MessagesHideAllChatJoinRequestsRequestTypeID = 0xe085f4ea

// Ensuring interfaces in compile-time for MessagesHideAllChatJoinRequestsRequest.
var (
	_ bin.Encoder     = &MessagesHideAllChatJoinRequestsRequest{}
	_ bin.Decoder     = &MessagesHideAllChatJoinRequestsRequest{}
	_ bin.BareEncoder = &MessagesHideAllChatJoinRequestsRequest{}
	_ bin.BareDecoder = &MessagesHideAllChatJoinRequestsRequest{}
)

func (h *MessagesHideAllChatJoinRequestsRequest) Zero() bool {
	if h == nil {
		return true
	}
	if !(h.Flags.Zero()) {
		return false
	}
	if !(h.Approved == false) {
		return false
	}
	if !(h.Peer == nil) {
		return false
	}
	if !(h.Link == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (h *MessagesHideAllChatJoinRequestsRequest) String() string {
	if h == nil {
		return "MessagesHideAllChatJoinRequestsRequest(nil)"
	}
	type Alias MessagesHideAllChatJoinRequestsRequest
	return fmt.Sprintf("MessagesHideAllChatJoinRequestsRequest%+v", Alias(*h))
}

// FillFrom fills MessagesHideAllChatJoinRequestsRequest from given interface.
func (h *MessagesHideAllChatJoinRequestsRequest) FillFrom(from interface {
	GetApproved() (value bool)
	GetPeer() (value InputPeerClass)
	GetLink() (value string, ok bool)
}) {
	h.Approved = from.GetApproved()
	h.Peer = from.GetPeer()
	if val, ok := from.GetLink(); ok {
		h.Link = val
	}

}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesHideAllChatJoinRequestsRequest) TypeID() uint32 {
	return MessagesHideAllChatJoinRequestsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesHideAllChatJoinRequestsRequest) TypeName() string {
	return "messages.hideAllChatJoinRequests"
}

// TypeInfo returns info about TL type.
func (h *MessagesHideAllChatJoinRequestsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.hideAllChatJoinRequests",
		ID:   MessagesHideAllChatJoinRequestsRequestTypeID,
	}
	if h == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Approved",
			SchemaName: "approved",
			Null:       !h.Flags.Has(0),
		},
		{
			Name:       "Peer",
			SchemaName: "peer",
		},
		{
			Name:       "Link",
			SchemaName: "link",
			Null:       !h.Flags.Has(1),
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (h *MessagesHideAllChatJoinRequestsRequest) SetFlags() {
	if !(h.Approved == false) {
		h.Flags.Set(0)
	}
	if !(h.Link == "") {
		h.Flags.Set(1)
	}
}

// Encode implements bin.Encoder.
func (h *MessagesHideAllChatJoinRequestsRequest) Encode(b *bin.Buffer) error {
	if h == nil {
		return fmt.Errorf("can't encode messages.hideAllChatJoinRequests#e085f4ea as nil")
	}
	b.PutID(MessagesHideAllChatJoinRequestsRequestTypeID)
	return h.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (h *MessagesHideAllChatJoinRequestsRequest) EncodeBare(b *bin.Buffer) error {
	if h == nil {
		return fmt.Errorf("can't encode messages.hideAllChatJoinRequests#e085f4ea as nil")
	}
	h.SetFlags()
	if err := h.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.hideAllChatJoinRequests#e085f4ea: field flags: %w", err)
	}
	if h.Peer == nil {
		return fmt.Errorf("unable to encode messages.hideAllChatJoinRequests#e085f4ea: field peer is nil")
	}
	if err := h.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.hideAllChatJoinRequests#e085f4ea: field peer: %w", err)
	}
	if h.Flags.Has(1) {
		b.PutString(h.Link)
	}
	return nil
}

// Decode implements bin.Decoder.
func (h *MessagesHideAllChatJoinRequestsRequest) Decode(b *bin.Buffer) error {
	if h == nil {
		return fmt.Errorf("can't decode messages.hideAllChatJoinRequests#e085f4ea to nil")
	}
	if err := b.ConsumeID(MessagesHideAllChatJoinRequestsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.hideAllChatJoinRequests#e085f4ea: %w", err)
	}
	return h.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (h *MessagesHideAllChatJoinRequestsRequest) DecodeBare(b *bin.Buffer) error {
	if h == nil {
		return fmt.Errorf("can't decode messages.hideAllChatJoinRequests#e085f4ea to nil")
	}
	{
		if err := h.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messages.hideAllChatJoinRequests#e085f4ea: field flags: %w", err)
		}
	}
	h.Approved = h.Flags.Has(0)
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.hideAllChatJoinRequests#e085f4ea: field peer: %w", err)
		}
		h.Peer = value
	}
	if h.Flags.Has(1) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode messages.hideAllChatJoinRequests#e085f4ea: field link: %w", err)
		}
		h.Link = value
	}
	return nil
}

// SetApproved sets value of Approved conditional field.
func (h *MessagesHideAllChatJoinRequestsRequest) SetApproved(value bool) {
	if value {
		h.Flags.Set(0)
		h.Approved = true
	} else {
		h.Flags.Unset(0)
		h.Approved = false
	}
}

// GetApproved returns value of Approved conditional field.
func (h *MessagesHideAllChatJoinRequestsRequest) GetApproved() (value bool) {
	if h == nil {
		return
	}
	return h.Flags.Has(0)
}

// GetPeer returns value of Peer field.
func (h *MessagesHideAllChatJoinRequestsRequest) GetPeer() (value InputPeerClass) {
	if h == nil {
		return
	}
	return h.Peer
}

// SetLink sets value of Link conditional field.
func (h *MessagesHideAllChatJoinRequestsRequest) SetLink(value string) {
	h.Flags.Set(1)
	h.Link = value
}

// GetLink returns value of Link conditional field and
// boolean which is true if field was set.
func (h *MessagesHideAllChatJoinRequestsRequest) GetLink() (value string, ok bool) {
	if h == nil {
		return
	}
	if !h.Flags.Has(1) {
		return value, false
	}
	return h.Link, true
}

// MessagesHideAllChatJoinRequests invokes method messages.hideAllChatJoinRequests#e085f4ea returning error if any.
//
// See https://core.telegram.org/method/messages.hideAllChatJoinRequests for reference.
func (c *Client) MessagesHideAllChatJoinRequests(ctx context.Context, request *MessagesHideAllChatJoinRequestsRequest) (UpdatesClass, error) {
	var result UpdatesBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Updates, nil
}
