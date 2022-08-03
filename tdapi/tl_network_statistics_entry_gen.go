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

// NetworkStatisticsEntryFile represents TL type `networkStatisticsEntryFile#b3b8f62`.
type NetworkStatisticsEntryFile struct {
	// Type of the file the data is part of; pass null if the data isn't related to files
	FileType FileTypeClass
	// Type of the network the data was sent through. Call setNetworkType to maintain the
	// actual network type
	NetworkType NetworkTypeClass
	// Total number of bytes sent
	SentBytes int64
	// Total number of bytes received
	ReceivedBytes int64
}

// NetworkStatisticsEntryFileTypeID is TL type id of NetworkStatisticsEntryFile.
const NetworkStatisticsEntryFileTypeID = 0xb3b8f62

// construct implements constructor of NetworkStatisticsEntryClass.
func (n NetworkStatisticsEntryFile) construct() NetworkStatisticsEntryClass { return &n }

// Ensuring interfaces in compile-time for NetworkStatisticsEntryFile.
var (
	_ bin.Encoder     = &NetworkStatisticsEntryFile{}
	_ bin.Decoder     = &NetworkStatisticsEntryFile{}
	_ bin.BareEncoder = &NetworkStatisticsEntryFile{}
	_ bin.BareDecoder = &NetworkStatisticsEntryFile{}

	_ NetworkStatisticsEntryClass = &NetworkStatisticsEntryFile{}
)

func (n *NetworkStatisticsEntryFile) Zero() bool {
	if n == nil {
		return true
	}
	if !(n.FileType == nil) {
		return false
	}
	if !(n.NetworkType == nil) {
		return false
	}
	if !(n.SentBytes == 0) {
		return false
	}
	if !(n.ReceivedBytes == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (n *NetworkStatisticsEntryFile) String() string {
	if n == nil {
		return "NetworkStatisticsEntryFile(nil)"
	}
	type Alias NetworkStatisticsEntryFile
	return fmt.Sprintf("NetworkStatisticsEntryFile%+v", Alias(*n))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*NetworkStatisticsEntryFile) TypeID() uint32 {
	return NetworkStatisticsEntryFileTypeID
}

// TypeName returns name of type in TL schema.
func (*NetworkStatisticsEntryFile) TypeName() string {
	return "networkStatisticsEntryFile"
}

// TypeInfo returns info about TL type.
func (n *NetworkStatisticsEntryFile) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "networkStatisticsEntryFile",
		ID:   NetworkStatisticsEntryFileTypeID,
	}
	if n == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "FileType",
			SchemaName: "file_type",
		},
		{
			Name:       "NetworkType",
			SchemaName: "network_type",
		},
		{
			Name:       "SentBytes",
			SchemaName: "sent_bytes",
		},
		{
			Name:       "ReceivedBytes",
			SchemaName: "received_bytes",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (n *NetworkStatisticsEntryFile) Encode(b *bin.Buffer) error {
	if n == nil {
		return fmt.Errorf("can't encode networkStatisticsEntryFile#b3b8f62 as nil")
	}
	b.PutID(NetworkStatisticsEntryFileTypeID)
	return n.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (n *NetworkStatisticsEntryFile) EncodeBare(b *bin.Buffer) error {
	if n == nil {
		return fmt.Errorf("can't encode networkStatisticsEntryFile#b3b8f62 as nil")
	}
	if n.FileType == nil {
		return fmt.Errorf("unable to encode networkStatisticsEntryFile#b3b8f62: field file_type is nil")
	}
	if err := n.FileType.Encode(b); err != nil {
		return fmt.Errorf("unable to encode networkStatisticsEntryFile#b3b8f62: field file_type: %w", err)
	}
	if n.NetworkType == nil {
		return fmt.Errorf("unable to encode networkStatisticsEntryFile#b3b8f62: field network_type is nil")
	}
	if err := n.NetworkType.Encode(b); err != nil {
		return fmt.Errorf("unable to encode networkStatisticsEntryFile#b3b8f62: field network_type: %w", err)
	}
	b.PutInt53(n.SentBytes)
	b.PutInt53(n.ReceivedBytes)
	return nil
}

// Decode implements bin.Decoder.
func (n *NetworkStatisticsEntryFile) Decode(b *bin.Buffer) error {
	if n == nil {
		return fmt.Errorf("can't decode networkStatisticsEntryFile#b3b8f62 to nil")
	}
	if err := b.ConsumeID(NetworkStatisticsEntryFileTypeID); err != nil {
		return fmt.Errorf("unable to decode networkStatisticsEntryFile#b3b8f62: %w", err)
	}
	return n.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (n *NetworkStatisticsEntryFile) DecodeBare(b *bin.Buffer) error {
	if n == nil {
		return fmt.Errorf("can't decode networkStatisticsEntryFile#b3b8f62 to nil")
	}
	{
		value, err := DecodeFileType(b)
		if err != nil {
			return fmt.Errorf("unable to decode networkStatisticsEntryFile#b3b8f62: field file_type: %w", err)
		}
		n.FileType = value
	}
	{
		value, err := DecodeNetworkType(b)
		if err != nil {
			return fmt.Errorf("unable to decode networkStatisticsEntryFile#b3b8f62: field network_type: %w", err)
		}
		n.NetworkType = value
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode networkStatisticsEntryFile#b3b8f62: field sent_bytes: %w", err)
		}
		n.SentBytes = value
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode networkStatisticsEntryFile#b3b8f62: field received_bytes: %w", err)
		}
		n.ReceivedBytes = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (n *NetworkStatisticsEntryFile) EncodeTDLibJSON(b tdjson.Encoder) error {
	if n == nil {
		return fmt.Errorf("can't encode networkStatisticsEntryFile#b3b8f62 as nil")
	}
	b.ObjStart()
	b.PutID("networkStatisticsEntryFile")
	b.Comma()
	b.FieldStart("file_type")
	if n.FileType == nil {
		return fmt.Errorf("unable to encode networkStatisticsEntryFile#b3b8f62: field file_type is nil")
	}
	if err := n.FileType.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode networkStatisticsEntryFile#b3b8f62: field file_type: %w", err)
	}
	b.Comma()
	b.FieldStart("network_type")
	if n.NetworkType == nil {
		return fmt.Errorf("unable to encode networkStatisticsEntryFile#b3b8f62: field network_type is nil")
	}
	if err := n.NetworkType.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode networkStatisticsEntryFile#b3b8f62: field network_type: %w", err)
	}
	b.Comma()
	b.FieldStart("sent_bytes")
	b.PutInt53(n.SentBytes)
	b.Comma()
	b.FieldStart("received_bytes")
	b.PutInt53(n.ReceivedBytes)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (n *NetworkStatisticsEntryFile) DecodeTDLibJSON(b tdjson.Decoder) error {
	if n == nil {
		return fmt.Errorf("can't decode networkStatisticsEntryFile#b3b8f62 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("networkStatisticsEntryFile"); err != nil {
				return fmt.Errorf("unable to decode networkStatisticsEntryFile#b3b8f62: %w", err)
			}
		case "file_type":
			value, err := DecodeTDLibJSONFileType(b)
			if err != nil {
				return fmt.Errorf("unable to decode networkStatisticsEntryFile#b3b8f62: field file_type: %w", err)
			}
			n.FileType = value
		case "network_type":
			value, err := DecodeTDLibJSONNetworkType(b)
			if err != nil {
				return fmt.Errorf("unable to decode networkStatisticsEntryFile#b3b8f62: field network_type: %w", err)
			}
			n.NetworkType = value
		case "sent_bytes":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode networkStatisticsEntryFile#b3b8f62: field sent_bytes: %w", err)
			}
			n.SentBytes = value
		case "received_bytes":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode networkStatisticsEntryFile#b3b8f62: field received_bytes: %w", err)
			}
			n.ReceivedBytes = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetFileType returns value of FileType field.
func (n *NetworkStatisticsEntryFile) GetFileType() (value FileTypeClass) {
	if n == nil {
		return
	}
	return n.FileType
}

// GetNetworkType returns value of NetworkType field.
func (n *NetworkStatisticsEntryFile) GetNetworkType() (value NetworkTypeClass) {
	if n == nil {
		return
	}
	return n.NetworkType
}

// GetSentBytes returns value of SentBytes field.
func (n *NetworkStatisticsEntryFile) GetSentBytes() (value int64) {
	if n == nil {
		return
	}
	return n.SentBytes
}

// GetReceivedBytes returns value of ReceivedBytes field.
func (n *NetworkStatisticsEntryFile) GetReceivedBytes() (value int64) {
	if n == nil {
		return
	}
	return n.ReceivedBytes
}

// NetworkStatisticsEntryCall represents TL type `networkStatisticsEntryCall#2bedbbad`.
type NetworkStatisticsEntryCall struct {
	// Type of the network the data was sent through. Call setNetworkType to maintain the
	// actual network type
	NetworkType NetworkTypeClass
	// Total number of bytes sent
	SentBytes int64
	// Total number of bytes received
	ReceivedBytes int64
	// Total call duration, in seconds
	Duration float64
}

// NetworkStatisticsEntryCallTypeID is TL type id of NetworkStatisticsEntryCall.
const NetworkStatisticsEntryCallTypeID = 0x2bedbbad

// construct implements constructor of NetworkStatisticsEntryClass.
func (n NetworkStatisticsEntryCall) construct() NetworkStatisticsEntryClass { return &n }

// Ensuring interfaces in compile-time for NetworkStatisticsEntryCall.
var (
	_ bin.Encoder     = &NetworkStatisticsEntryCall{}
	_ bin.Decoder     = &NetworkStatisticsEntryCall{}
	_ bin.BareEncoder = &NetworkStatisticsEntryCall{}
	_ bin.BareDecoder = &NetworkStatisticsEntryCall{}

	_ NetworkStatisticsEntryClass = &NetworkStatisticsEntryCall{}
)

func (n *NetworkStatisticsEntryCall) Zero() bool {
	if n == nil {
		return true
	}
	if !(n.NetworkType == nil) {
		return false
	}
	if !(n.SentBytes == 0) {
		return false
	}
	if !(n.ReceivedBytes == 0) {
		return false
	}
	if !(n.Duration == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (n *NetworkStatisticsEntryCall) String() string {
	if n == nil {
		return "NetworkStatisticsEntryCall(nil)"
	}
	type Alias NetworkStatisticsEntryCall
	return fmt.Sprintf("NetworkStatisticsEntryCall%+v", Alias(*n))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*NetworkStatisticsEntryCall) TypeID() uint32 {
	return NetworkStatisticsEntryCallTypeID
}

// TypeName returns name of type in TL schema.
func (*NetworkStatisticsEntryCall) TypeName() string {
	return "networkStatisticsEntryCall"
}

// TypeInfo returns info about TL type.
func (n *NetworkStatisticsEntryCall) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "networkStatisticsEntryCall",
		ID:   NetworkStatisticsEntryCallTypeID,
	}
	if n == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "NetworkType",
			SchemaName: "network_type",
		},
		{
			Name:       "SentBytes",
			SchemaName: "sent_bytes",
		},
		{
			Name:       "ReceivedBytes",
			SchemaName: "received_bytes",
		},
		{
			Name:       "Duration",
			SchemaName: "duration",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (n *NetworkStatisticsEntryCall) Encode(b *bin.Buffer) error {
	if n == nil {
		return fmt.Errorf("can't encode networkStatisticsEntryCall#2bedbbad as nil")
	}
	b.PutID(NetworkStatisticsEntryCallTypeID)
	return n.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (n *NetworkStatisticsEntryCall) EncodeBare(b *bin.Buffer) error {
	if n == nil {
		return fmt.Errorf("can't encode networkStatisticsEntryCall#2bedbbad as nil")
	}
	if n.NetworkType == nil {
		return fmt.Errorf("unable to encode networkStatisticsEntryCall#2bedbbad: field network_type is nil")
	}
	if err := n.NetworkType.Encode(b); err != nil {
		return fmt.Errorf("unable to encode networkStatisticsEntryCall#2bedbbad: field network_type: %w", err)
	}
	b.PutInt53(n.SentBytes)
	b.PutInt53(n.ReceivedBytes)
	b.PutDouble(n.Duration)
	return nil
}

// Decode implements bin.Decoder.
func (n *NetworkStatisticsEntryCall) Decode(b *bin.Buffer) error {
	if n == nil {
		return fmt.Errorf("can't decode networkStatisticsEntryCall#2bedbbad to nil")
	}
	if err := b.ConsumeID(NetworkStatisticsEntryCallTypeID); err != nil {
		return fmt.Errorf("unable to decode networkStatisticsEntryCall#2bedbbad: %w", err)
	}
	return n.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (n *NetworkStatisticsEntryCall) DecodeBare(b *bin.Buffer) error {
	if n == nil {
		return fmt.Errorf("can't decode networkStatisticsEntryCall#2bedbbad to nil")
	}
	{
		value, err := DecodeNetworkType(b)
		if err != nil {
			return fmt.Errorf("unable to decode networkStatisticsEntryCall#2bedbbad: field network_type: %w", err)
		}
		n.NetworkType = value
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode networkStatisticsEntryCall#2bedbbad: field sent_bytes: %w", err)
		}
		n.SentBytes = value
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode networkStatisticsEntryCall#2bedbbad: field received_bytes: %w", err)
		}
		n.ReceivedBytes = value
	}
	{
		value, err := b.Double()
		if err != nil {
			return fmt.Errorf("unable to decode networkStatisticsEntryCall#2bedbbad: field duration: %w", err)
		}
		n.Duration = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (n *NetworkStatisticsEntryCall) EncodeTDLibJSON(b tdjson.Encoder) error {
	if n == nil {
		return fmt.Errorf("can't encode networkStatisticsEntryCall#2bedbbad as nil")
	}
	b.ObjStart()
	b.PutID("networkStatisticsEntryCall")
	b.Comma()
	b.FieldStart("network_type")
	if n.NetworkType == nil {
		return fmt.Errorf("unable to encode networkStatisticsEntryCall#2bedbbad: field network_type is nil")
	}
	if err := n.NetworkType.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode networkStatisticsEntryCall#2bedbbad: field network_type: %w", err)
	}
	b.Comma()
	b.FieldStart("sent_bytes")
	b.PutInt53(n.SentBytes)
	b.Comma()
	b.FieldStart("received_bytes")
	b.PutInt53(n.ReceivedBytes)
	b.Comma()
	b.FieldStart("duration")
	b.PutDouble(n.Duration)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (n *NetworkStatisticsEntryCall) DecodeTDLibJSON(b tdjson.Decoder) error {
	if n == nil {
		return fmt.Errorf("can't decode networkStatisticsEntryCall#2bedbbad to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("networkStatisticsEntryCall"); err != nil {
				return fmt.Errorf("unable to decode networkStatisticsEntryCall#2bedbbad: %w", err)
			}
		case "network_type":
			value, err := DecodeTDLibJSONNetworkType(b)
			if err != nil {
				return fmt.Errorf("unable to decode networkStatisticsEntryCall#2bedbbad: field network_type: %w", err)
			}
			n.NetworkType = value
		case "sent_bytes":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode networkStatisticsEntryCall#2bedbbad: field sent_bytes: %w", err)
			}
			n.SentBytes = value
		case "received_bytes":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode networkStatisticsEntryCall#2bedbbad: field received_bytes: %w", err)
			}
			n.ReceivedBytes = value
		case "duration":
			value, err := b.Double()
			if err != nil {
				return fmt.Errorf("unable to decode networkStatisticsEntryCall#2bedbbad: field duration: %w", err)
			}
			n.Duration = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetNetworkType returns value of NetworkType field.
func (n *NetworkStatisticsEntryCall) GetNetworkType() (value NetworkTypeClass) {
	if n == nil {
		return
	}
	return n.NetworkType
}

// GetSentBytes returns value of SentBytes field.
func (n *NetworkStatisticsEntryCall) GetSentBytes() (value int64) {
	if n == nil {
		return
	}
	return n.SentBytes
}

// GetReceivedBytes returns value of ReceivedBytes field.
func (n *NetworkStatisticsEntryCall) GetReceivedBytes() (value int64) {
	if n == nil {
		return
	}
	return n.ReceivedBytes
}

// GetDuration returns value of Duration field.
func (n *NetworkStatisticsEntryCall) GetDuration() (value float64) {
	if n == nil {
		return
	}
	return n.Duration
}

// NetworkStatisticsEntryClassName is schema name of NetworkStatisticsEntryClass.
const NetworkStatisticsEntryClassName = "NetworkStatisticsEntry"

// NetworkStatisticsEntryClass represents NetworkStatisticsEntry generic type.
//
// Example:
//
//	g, err := tdapi.DecodeNetworkStatisticsEntry(buf)
//	if err != nil {
//	    panic(err)
//	}
//	switch v := g.(type) {
//	case *tdapi.NetworkStatisticsEntryFile: // networkStatisticsEntryFile#b3b8f62
//	case *tdapi.NetworkStatisticsEntryCall: // networkStatisticsEntryCall#2bedbbad
//	default: panic(v)
//	}
type NetworkStatisticsEntryClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() NetworkStatisticsEntryClass

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

	// Type of the network the data was sent through. Call setNetworkType to maintain the
	// actual network type
	GetNetworkType() (value NetworkTypeClass)
	// Total number of bytes sent
	GetSentBytes() (value int64)
	// Total number of bytes received
	GetReceivedBytes() (value int64)
}

// DecodeNetworkStatisticsEntry implements binary de-serialization for NetworkStatisticsEntryClass.
func DecodeNetworkStatisticsEntry(buf *bin.Buffer) (NetworkStatisticsEntryClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case NetworkStatisticsEntryFileTypeID:
		// Decoding networkStatisticsEntryFile#b3b8f62.
		v := NetworkStatisticsEntryFile{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode NetworkStatisticsEntryClass: %w", err)
		}
		return &v, nil
	case NetworkStatisticsEntryCallTypeID:
		// Decoding networkStatisticsEntryCall#2bedbbad.
		v := NetworkStatisticsEntryCall{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode NetworkStatisticsEntryClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode NetworkStatisticsEntryClass: %w", bin.NewUnexpectedID(id))
	}
}

// DecodeTDLibJSONNetworkStatisticsEntry implements binary de-serialization for NetworkStatisticsEntryClass.
func DecodeTDLibJSONNetworkStatisticsEntry(buf tdjson.Decoder) (NetworkStatisticsEntryClass, error) {
	id, err := buf.FindTypeID()
	if err != nil {
		return nil, err
	}
	switch id {
	case "networkStatisticsEntryFile":
		// Decoding networkStatisticsEntryFile#b3b8f62.
		v := NetworkStatisticsEntryFile{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode NetworkStatisticsEntryClass: %w", err)
		}
		return &v, nil
	case "networkStatisticsEntryCall":
		// Decoding networkStatisticsEntryCall#2bedbbad.
		v := NetworkStatisticsEntryCall{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode NetworkStatisticsEntryClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode NetworkStatisticsEntryClass: %w", tdjson.NewUnexpectedID(id))
	}
}

// NetworkStatisticsEntry boxes the NetworkStatisticsEntryClass providing a helper.
type NetworkStatisticsEntryBox struct {
	NetworkStatisticsEntry NetworkStatisticsEntryClass
}

// Decode implements bin.Decoder for NetworkStatisticsEntryBox.
func (b *NetworkStatisticsEntryBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode NetworkStatisticsEntryBox to nil")
	}
	v, err := DecodeNetworkStatisticsEntry(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.NetworkStatisticsEntry = v
	return nil
}

// Encode implements bin.Encode for NetworkStatisticsEntryBox.
func (b *NetworkStatisticsEntryBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.NetworkStatisticsEntry == nil {
		return fmt.Errorf("unable to encode NetworkStatisticsEntryClass as nil")
	}
	return b.NetworkStatisticsEntry.Encode(buf)
}

// DecodeTDLibJSON implements bin.Decoder for NetworkStatisticsEntryBox.
func (b *NetworkStatisticsEntryBox) DecodeTDLibJSON(buf tdjson.Decoder) error {
	if b == nil {
		return fmt.Errorf("unable to decode NetworkStatisticsEntryBox to nil")
	}
	v, err := DecodeTDLibJSONNetworkStatisticsEntry(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.NetworkStatisticsEntry = v
	return nil
}

// EncodeTDLibJSON implements bin.Encode for NetworkStatisticsEntryBox.
func (b *NetworkStatisticsEntryBox) EncodeTDLibJSON(buf tdjson.Encoder) error {
	if b == nil || b.NetworkStatisticsEntry == nil {
		return fmt.Errorf("unable to encode NetworkStatisticsEntryClass as nil")
	}
	return b.NetworkStatisticsEntry.EncodeTDLibJSON(buf)
}
