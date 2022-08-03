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

// InputPhoneContact represents TL type `inputPhoneContact#f392b7f4`.
// Phone contact. The client_id is just an arbitrary contact ID: it should be set, for
// example, to an incremental number when using contacts.importContacts¹, in order to
// retry importing only the contacts that weren't imported successfully.
//
// Links:
//  1. https://core.telegram.org/method/contacts.importContacts
//
// See https://core.telegram.org/constructor/inputPhoneContact for reference.
type InputPhoneContact struct {
	// User identifier on the client
	ClientID int64
	// Phone number
	Phone string
	// Contact's first name
	FirstName string
	// Contact's last name
	LastName string
}

// InputPhoneContactTypeID is TL type id of InputPhoneContact.
const InputPhoneContactTypeID = 0xf392b7f4

// Ensuring interfaces in compile-time for InputPhoneContact.
var (
	_ bin.Encoder     = &InputPhoneContact{}
	_ bin.Decoder     = &InputPhoneContact{}
	_ bin.BareEncoder = &InputPhoneContact{}
	_ bin.BareDecoder = &InputPhoneContact{}
)

func (i *InputPhoneContact) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.ClientID == 0) {
		return false
	}
	if !(i.Phone == "") {
		return false
	}
	if !(i.FirstName == "") {
		return false
	}
	if !(i.LastName == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputPhoneContact) String() string {
	if i == nil {
		return "InputPhoneContact(nil)"
	}
	type Alias InputPhoneContact
	return fmt.Sprintf("InputPhoneContact%+v", Alias(*i))
}

// FillFrom fills InputPhoneContact from given interface.
func (i *InputPhoneContact) FillFrom(from interface {
	GetClientID() (value int64)
	GetPhone() (value string)
	GetFirstName() (value string)
	GetLastName() (value string)
}) {
	i.ClientID = from.GetClientID()
	i.Phone = from.GetPhone()
	i.FirstName = from.GetFirstName()
	i.LastName = from.GetLastName()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InputPhoneContact) TypeID() uint32 {
	return InputPhoneContactTypeID
}

// TypeName returns name of type in TL schema.
func (*InputPhoneContact) TypeName() string {
	return "inputPhoneContact"
}

// TypeInfo returns info about TL type.
func (i *InputPhoneContact) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inputPhoneContact",
		ID:   InputPhoneContactTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ClientID",
			SchemaName: "client_id",
		},
		{
			Name:       "Phone",
			SchemaName: "phone",
		},
		{
			Name:       "FirstName",
			SchemaName: "first_name",
		},
		{
			Name:       "LastName",
			SchemaName: "last_name",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (i *InputPhoneContact) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputPhoneContact#f392b7f4 as nil")
	}
	b.PutID(InputPhoneContactTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InputPhoneContact) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputPhoneContact#f392b7f4 as nil")
	}
	b.PutLong(i.ClientID)
	b.PutString(i.Phone)
	b.PutString(i.FirstName)
	b.PutString(i.LastName)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputPhoneContact) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputPhoneContact#f392b7f4 to nil")
	}
	if err := b.ConsumeID(InputPhoneContactTypeID); err != nil {
		return fmt.Errorf("unable to decode inputPhoneContact#f392b7f4: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InputPhoneContact) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputPhoneContact#f392b7f4 to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode inputPhoneContact#f392b7f4: field client_id: %w", err)
		}
		i.ClientID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode inputPhoneContact#f392b7f4: field phone: %w", err)
		}
		i.Phone = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode inputPhoneContact#f392b7f4: field first_name: %w", err)
		}
		i.FirstName = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode inputPhoneContact#f392b7f4: field last_name: %w", err)
		}
		i.LastName = value
	}
	return nil
}

// GetClientID returns value of ClientID field.
func (i *InputPhoneContact) GetClientID() (value int64) {
	if i == nil {
		return
	}
	return i.ClientID
}

// GetPhone returns value of Phone field.
func (i *InputPhoneContact) GetPhone() (value string) {
	if i == nil {
		return
	}
	return i.Phone
}

// GetFirstName returns value of FirstName field.
func (i *InputPhoneContact) GetFirstName() (value string) {
	if i == nil {
		return
	}
	return i.FirstName
}

// GetLastName returns value of LastName field.
func (i *InputPhoneContact) GetLastName() (value string) {
	if i == nil {
		return
	}
	return i.LastName
}
