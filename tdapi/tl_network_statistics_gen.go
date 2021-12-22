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

// NetworkStatistics represents TL type `networkStatistics#d86acb3c`.
type NetworkStatistics struct {
	// Point in time (Unix timestamp) from which the statistics are collected
	SinceDate int32
	// Network statistics entries
	Entries []NetworkStatisticsEntryClass
}

// NetworkStatisticsTypeID is TL type id of NetworkStatistics.
const NetworkStatisticsTypeID = 0xd86acb3c

// Ensuring interfaces in compile-time for NetworkStatistics.
var (
	_ bin.Encoder     = &NetworkStatistics{}
	_ bin.Decoder     = &NetworkStatistics{}
	_ bin.BareEncoder = &NetworkStatistics{}
	_ bin.BareDecoder = &NetworkStatistics{}
)

func (n *NetworkStatistics) Zero() bool {
	if n == nil {
		return true
	}
	if !(n.SinceDate == 0) {
		return false
	}
	if !(n.Entries == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (n *NetworkStatistics) String() string {
	if n == nil {
		return "NetworkStatistics(nil)"
	}
	type Alias NetworkStatistics
	return fmt.Sprintf("NetworkStatistics%+v", Alias(*n))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*NetworkStatistics) TypeID() uint32 {
	return NetworkStatisticsTypeID
}

// TypeName returns name of type in TL schema.
func (*NetworkStatistics) TypeName() string {
	return "networkStatistics"
}

// TypeInfo returns info about TL type.
func (n *NetworkStatistics) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "networkStatistics",
		ID:   NetworkStatisticsTypeID,
	}
	if n == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "SinceDate",
			SchemaName: "since_date",
		},
		{
			Name:       "Entries",
			SchemaName: "entries",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (n *NetworkStatistics) Encode(b *bin.Buffer) error {
	if n == nil {
		return fmt.Errorf("can't encode networkStatistics#d86acb3c as nil")
	}
	b.PutID(NetworkStatisticsTypeID)
	return n.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (n *NetworkStatistics) EncodeBare(b *bin.Buffer) error {
	if n == nil {
		return fmt.Errorf("can't encode networkStatistics#d86acb3c as nil")
	}
	b.PutInt32(n.SinceDate)
	b.PutInt(len(n.Entries))
	for idx, v := range n.Entries {
		if v == nil {
			return fmt.Errorf("unable to encode networkStatistics#d86acb3c: field entries element with index %d is nil", idx)
		}
		if err := v.EncodeBare(b); err != nil {
			return fmt.Errorf("unable to encode bare networkStatistics#d86acb3c: field entries element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (n *NetworkStatistics) Decode(b *bin.Buffer) error {
	if n == nil {
		return fmt.Errorf("can't decode networkStatistics#d86acb3c to nil")
	}
	if err := b.ConsumeID(NetworkStatisticsTypeID); err != nil {
		return fmt.Errorf("unable to decode networkStatistics#d86acb3c: %w", err)
	}
	return n.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (n *NetworkStatistics) DecodeBare(b *bin.Buffer) error {
	if n == nil {
		return fmt.Errorf("can't decode networkStatistics#d86acb3c to nil")
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode networkStatistics#d86acb3c: field since_date: %w", err)
		}
		n.SinceDate = value
	}
	{
		headerLen, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode networkStatistics#d86acb3c: field entries: %w", err)
		}

		if headerLen > 0 {
			n.Entries = make([]NetworkStatisticsEntryClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeNetworkStatisticsEntry(b)
			if err != nil {
				return fmt.Errorf("unable to decode networkStatistics#d86acb3c: field entries: %w", err)
			}
			n.Entries = append(n.Entries, value)
		}
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (n *NetworkStatistics) EncodeTDLibJSON(b tdjson.Encoder) error {
	if n == nil {
		return fmt.Errorf("can't encode networkStatistics#d86acb3c as nil")
	}
	b.ObjStart()
	b.PutID("networkStatistics")
	b.FieldStart("since_date")
	b.PutInt32(n.SinceDate)
	b.FieldStart("entries")
	b.ArrStart()
	for idx, v := range n.Entries {
		if v == nil {
			return fmt.Errorf("unable to encode networkStatistics#d86acb3c: field entries element with index %d is nil", idx)
		}
		if err := v.EncodeTDLibJSON(b); err != nil {
			return fmt.Errorf("unable to encode networkStatistics#d86acb3c: field entries element with index %d: %w", idx, err)
		}
	}
	b.ArrEnd()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (n *NetworkStatistics) DecodeTDLibJSON(b tdjson.Decoder) error {
	if n == nil {
		return fmt.Errorf("can't decode networkStatistics#d86acb3c to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("networkStatistics"); err != nil {
				return fmt.Errorf("unable to decode networkStatistics#d86acb3c: %w", err)
			}
		case "since_date":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode networkStatistics#d86acb3c: field since_date: %w", err)
			}
			n.SinceDate = value
		case "entries":
			if err := b.Arr(func(b tdjson.Decoder) error {
				value, err := DecodeTDLibJSONNetworkStatisticsEntry(b)
				if err != nil {
					return fmt.Errorf("unable to decode networkStatistics#d86acb3c: field entries: %w", err)
				}
				n.Entries = append(n.Entries, value)
				return nil
			}); err != nil {
				return fmt.Errorf("unable to decode networkStatistics#d86acb3c: field entries: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetSinceDate returns value of SinceDate field.
func (n *NetworkStatistics) GetSinceDate() (value int32) {
	if n == nil {
		return
	}
	return n.SinceDate
}

// GetEntries returns value of Entries field.
func (n *NetworkStatistics) GetEntries() (value []NetworkStatisticsEntryClass) {
	if n == nil {
		return
	}
	return n.Entries
}
