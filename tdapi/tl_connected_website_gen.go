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

// ConnectedWebsite represents TL type `connectedWebsite#324ea36d`.
type ConnectedWebsite struct {
	// Website identifier
	ID int64
	// The domain name of the website
	DomainName string
	// User identifier of a bot linked with the website
	BotUserID int64
	// The version of a browser used to log in
	Browser string
	// Operating system the browser is running on
	Platform string
	// Point in time (Unix timestamp) when the user was logged in
	LogInDate int32
	// Point in time (Unix timestamp) when obtained authorization was last used
	LastActiveDate int32
	// IP address from which the user was logged in, in human-readable format
	IP string
	// Human-readable description of a country and a region, from which the user was logged
	// in, based on the IP address
	Location string
}

// ConnectedWebsiteTypeID is TL type id of ConnectedWebsite.
const ConnectedWebsiteTypeID = 0x324ea36d

// Ensuring interfaces in compile-time for ConnectedWebsite.
var (
	_ bin.Encoder     = &ConnectedWebsite{}
	_ bin.Decoder     = &ConnectedWebsite{}
	_ bin.BareEncoder = &ConnectedWebsite{}
	_ bin.BareDecoder = &ConnectedWebsite{}
)

func (c *ConnectedWebsite) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.ID == 0) {
		return false
	}
	if !(c.DomainName == "") {
		return false
	}
	if !(c.BotUserID == 0) {
		return false
	}
	if !(c.Browser == "") {
		return false
	}
	if !(c.Platform == "") {
		return false
	}
	if !(c.LogInDate == 0) {
		return false
	}
	if !(c.LastActiveDate == 0) {
		return false
	}
	if !(c.IP == "") {
		return false
	}
	if !(c.Location == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *ConnectedWebsite) String() string {
	if c == nil {
		return "ConnectedWebsite(nil)"
	}
	type Alias ConnectedWebsite
	return fmt.Sprintf("ConnectedWebsite%+v", Alias(*c))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ConnectedWebsite) TypeID() uint32 {
	return ConnectedWebsiteTypeID
}

// TypeName returns name of type in TL schema.
func (*ConnectedWebsite) TypeName() string {
	return "connectedWebsite"
}

// TypeInfo returns info about TL type.
func (c *ConnectedWebsite) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "connectedWebsite",
		ID:   ConnectedWebsiteTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ID",
			SchemaName: "id",
		},
		{
			Name:       "DomainName",
			SchemaName: "domain_name",
		},
		{
			Name:       "BotUserID",
			SchemaName: "bot_user_id",
		},
		{
			Name:       "Browser",
			SchemaName: "browser",
		},
		{
			Name:       "Platform",
			SchemaName: "platform",
		},
		{
			Name:       "LogInDate",
			SchemaName: "log_in_date",
		},
		{
			Name:       "LastActiveDate",
			SchemaName: "last_active_date",
		},
		{
			Name:       "IP",
			SchemaName: "ip",
		},
		{
			Name:       "Location",
			SchemaName: "location",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *ConnectedWebsite) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode connectedWebsite#324ea36d as nil")
	}
	b.PutID(ConnectedWebsiteTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *ConnectedWebsite) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode connectedWebsite#324ea36d as nil")
	}
	b.PutLong(c.ID)
	b.PutString(c.DomainName)
	b.PutInt53(c.BotUserID)
	b.PutString(c.Browser)
	b.PutString(c.Platform)
	b.PutInt32(c.LogInDate)
	b.PutInt32(c.LastActiveDate)
	b.PutString(c.IP)
	b.PutString(c.Location)
	return nil
}

// Decode implements bin.Decoder.
func (c *ConnectedWebsite) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode connectedWebsite#324ea36d to nil")
	}
	if err := b.ConsumeID(ConnectedWebsiteTypeID); err != nil {
		return fmt.Errorf("unable to decode connectedWebsite#324ea36d: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *ConnectedWebsite) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode connectedWebsite#324ea36d to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode connectedWebsite#324ea36d: field id: %w", err)
		}
		c.ID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode connectedWebsite#324ea36d: field domain_name: %w", err)
		}
		c.DomainName = value
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode connectedWebsite#324ea36d: field bot_user_id: %w", err)
		}
		c.BotUserID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode connectedWebsite#324ea36d: field browser: %w", err)
		}
		c.Browser = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode connectedWebsite#324ea36d: field platform: %w", err)
		}
		c.Platform = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode connectedWebsite#324ea36d: field log_in_date: %w", err)
		}
		c.LogInDate = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode connectedWebsite#324ea36d: field last_active_date: %w", err)
		}
		c.LastActiveDate = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode connectedWebsite#324ea36d: field ip: %w", err)
		}
		c.IP = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode connectedWebsite#324ea36d: field location: %w", err)
		}
		c.Location = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (c *ConnectedWebsite) EncodeTDLibJSON(b tdjson.Encoder) error {
	if c == nil {
		return fmt.Errorf("can't encode connectedWebsite#324ea36d as nil")
	}
	b.ObjStart()
	b.PutID("connectedWebsite")
	b.FieldStart("id")
	b.PutLong(c.ID)
	b.FieldStart("domain_name")
	b.PutString(c.DomainName)
	b.FieldStart("bot_user_id")
	b.PutInt53(c.BotUserID)
	b.FieldStart("browser")
	b.PutString(c.Browser)
	b.FieldStart("platform")
	b.PutString(c.Platform)
	b.FieldStart("log_in_date")
	b.PutInt32(c.LogInDate)
	b.FieldStart("last_active_date")
	b.PutInt32(c.LastActiveDate)
	b.FieldStart("ip")
	b.PutString(c.IP)
	b.FieldStart("location")
	b.PutString(c.Location)
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (c *ConnectedWebsite) DecodeTDLibJSON(b tdjson.Decoder) error {
	if c == nil {
		return fmt.Errorf("can't decode connectedWebsite#324ea36d to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("connectedWebsite"); err != nil {
				return fmt.Errorf("unable to decode connectedWebsite#324ea36d: %w", err)
			}
		case "id":
			value, err := b.Long()
			if err != nil {
				return fmt.Errorf("unable to decode connectedWebsite#324ea36d: field id: %w", err)
			}
			c.ID = value
		case "domain_name":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode connectedWebsite#324ea36d: field domain_name: %w", err)
			}
			c.DomainName = value
		case "bot_user_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode connectedWebsite#324ea36d: field bot_user_id: %w", err)
			}
			c.BotUserID = value
		case "browser":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode connectedWebsite#324ea36d: field browser: %w", err)
			}
			c.Browser = value
		case "platform":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode connectedWebsite#324ea36d: field platform: %w", err)
			}
			c.Platform = value
		case "log_in_date":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode connectedWebsite#324ea36d: field log_in_date: %w", err)
			}
			c.LogInDate = value
		case "last_active_date":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode connectedWebsite#324ea36d: field last_active_date: %w", err)
			}
			c.LastActiveDate = value
		case "ip":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode connectedWebsite#324ea36d: field ip: %w", err)
			}
			c.IP = value
		case "location":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode connectedWebsite#324ea36d: field location: %w", err)
			}
			c.Location = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetID returns value of ID field.
func (c *ConnectedWebsite) GetID() (value int64) {
	if c == nil {
		return
	}
	return c.ID
}

// GetDomainName returns value of DomainName field.
func (c *ConnectedWebsite) GetDomainName() (value string) {
	if c == nil {
		return
	}
	return c.DomainName
}

// GetBotUserID returns value of BotUserID field.
func (c *ConnectedWebsite) GetBotUserID() (value int64) {
	if c == nil {
		return
	}
	return c.BotUserID
}

// GetBrowser returns value of Browser field.
func (c *ConnectedWebsite) GetBrowser() (value string) {
	if c == nil {
		return
	}
	return c.Browser
}

// GetPlatform returns value of Platform field.
func (c *ConnectedWebsite) GetPlatform() (value string) {
	if c == nil {
		return
	}
	return c.Platform
}

// GetLogInDate returns value of LogInDate field.
func (c *ConnectedWebsite) GetLogInDate() (value int32) {
	if c == nil {
		return
	}
	return c.LogInDate
}

// GetLastActiveDate returns value of LastActiveDate field.
func (c *ConnectedWebsite) GetLastActiveDate() (value int32) {
	if c == nil {
		return
	}
	return c.LastActiveDate
}

// GetIP returns value of IP field.
func (c *ConnectedWebsite) GetIP() (value string) {
	if c == nil {
		return
	}
	return c.IP
}

// GetLocation returns value of Location field.
func (c *ConnectedWebsite) GetLocation() (value string) {
	if c == nil {
		return
	}
	return c.Location
}
