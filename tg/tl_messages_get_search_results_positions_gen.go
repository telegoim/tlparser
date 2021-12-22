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

// MessagesGetSearchResultsPositionsRequest represents TL type `messages.getSearchResultsPositions#6e9583a3`.
//
// See https://core.telegram.org/method/messages.getSearchResultsPositions for reference.
type MessagesGetSearchResultsPositionsRequest struct {
	// Peer field of MessagesGetSearchResultsPositionsRequest.
	Peer InputPeerClass
	// Filter field of MessagesGetSearchResultsPositionsRequest.
	Filter MessagesFilterClass
	// OffsetID field of MessagesGetSearchResultsPositionsRequest.
	OffsetID int
	// Limit field of MessagesGetSearchResultsPositionsRequest.
	Limit int
}

// MessagesGetSearchResultsPositionsRequestTypeID is TL type id of MessagesGetSearchResultsPositionsRequest.
const MessagesGetSearchResultsPositionsRequestTypeID = 0x6e9583a3

// Ensuring interfaces in compile-time for MessagesGetSearchResultsPositionsRequest.
var (
	_ bin.Encoder     = &MessagesGetSearchResultsPositionsRequest{}
	_ bin.Decoder     = &MessagesGetSearchResultsPositionsRequest{}
	_ bin.BareEncoder = &MessagesGetSearchResultsPositionsRequest{}
	_ bin.BareDecoder = &MessagesGetSearchResultsPositionsRequest{}
)

func (g *MessagesGetSearchResultsPositionsRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Peer == nil) {
		return false
	}
	if !(g.Filter == nil) {
		return false
	}
	if !(g.OffsetID == 0) {
		return false
	}
	if !(g.Limit == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *MessagesGetSearchResultsPositionsRequest) String() string {
	if g == nil {
		return "MessagesGetSearchResultsPositionsRequest(nil)"
	}
	type Alias MessagesGetSearchResultsPositionsRequest
	return fmt.Sprintf("MessagesGetSearchResultsPositionsRequest%+v", Alias(*g))
}

// FillFrom fills MessagesGetSearchResultsPositionsRequest from given interface.
func (g *MessagesGetSearchResultsPositionsRequest) FillFrom(from interface {
	GetPeer() (value InputPeerClass)
	GetFilter() (value MessagesFilterClass)
	GetOffsetID() (value int)
	GetLimit() (value int)
}) {
	g.Peer = from.GetPeer()
	g.Filter = from.GetFilter()
	g.OffsetID = from.GetOffsetID()
	g.Limit = from.GetLimit()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesGetSearchResultsPositionsRequest) TypeID() uint32 {
	return MessagesGetSearchResultsPositionsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesGetSearchResultsPositionsRequest) TypeName() string {
	return "messages.getSearchResultsPositions"
}

// TypeInfo returns info about TL type.
func (g *MessagesGetSearchResultsPositionsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.getSearchResultsPositions",
		ID:   MessagesGetSearchResultsPositionsRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Peer",
			SchemaName: "peer",
		},
		{
			Name:       "Filter",
			SchemaName: "filter",
		},
		{
			Name:       "OffsetID",
			SchemaName: "offset_id",
		},
		{
			Name:       "Limit",
			SchemaName: "limit",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *MessagesGetSearchResultsPositionsRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getSearchResultsPositions#6e9583a3 as nil")
	}
	b.PutID(MessagesGetSearchResultsPositionsRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *MessagesGetSearchResultsPositionsRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getSearchResultsPositions#6e9583a3 as nil")
	}
	if g.Peer == nil {
		return fmt.Errorf("unable to encode messages.getSearchResultsPositions#6e9583a3: field peer is nil")
	}
	if err := g.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.getSearchResultsPositions#6e9583a3: field peer: %w", err)
	}
	if g.Filter == nil {
		return fmt.Errorf("unable to encode messages.getSearchResultsPositions#6e9583a3: field filter is nil")
	}
	if err := g.Filter.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.getSearchResultsPositions#6e9583a3: field filter: %w", err)
	}
	b.PutInt(g.OffsetID)
	b.PutInt(g.Limit)
	return nil
}

// Decode implements bin.Decoder.
func (g *MessagesGetSearchResultsPositionsRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getSearchResultsPositions#6e9583a3 to nil")
	}
	if err := b.ConsumeID(MessagesGetSearchResultsPositionsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.getSearchResultsPositions#6e9583a3: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *MessagesGetSearchResultsPositionsRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getSearchResultsPositions#6e9583a3 to nil")
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.getSearchResultsPositions#6e9583a3: field peer: %w", err)
		}
		g.Peer = value
	}
	{
		value, err := DecodeMessagesFilter(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.getSearchResultsPositions#6e9583a3: field filter: %w", err)
		}
		g.Filter = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getSearchResultsPositions#6e9583a3: field offset_id: %w", err)
		}
		g.OffsetID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getSearchResultsPositions#6e9583a3: field limit: %w", err)
		}
		g.Limit = value
	}
	return nil
}

// GetPeer returns value of Peer field.
func (g *MessagesGetSearchResultsPositionsRequest) GetPeer() (value InputPeerClass) {
	if g == nil {
		return
	}
	return g.Peer
}

// GetFilter returns value of Filter field.
func (g *MessagesGetSearchResultsPositionsRequest) GetFilter() (value MessagesFilterClass) {
	if g == nil {
		return
	}
	return g.Filter
}

// GetOffsetID returns value of OffsetID field.
func (g *MessagesGetSearchResultsPositionsRequest) GetOffsetID() (value int) {
	if g == nil {
		return
	}
	return g.OffsetID
}

// GetLimit returns value of Limit field.
func (g *MessagesGetSearchResultsPositionsRequest) GetLimit() (value int) {
	if g == nil {
		return
	}
	return g.Limit
}

// MessagesGetSearchResultsPositions invokes method messages.getSearchResultsPositions#6e9583a3 returning error if any.
//
// See https://core.telegram.org/method/messages.getSearchResultsPositions for reference.
func (c *Client) MessagesGetSearchResultsPositions(ctx context.Context, request *MessagesGetSearchResultsPositionsRequest) (*MessagesSearchResultsPositions, error) {
	var result MessagesSearchResultsPositions

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
