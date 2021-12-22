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

// TestCallVectorIntObjectRequest represents TL type `testCallVectorIntObject#ca57f472`.
type TestCallVectorIntObjectRequest struct {
	// Vector of objects to return
	X []TestInt
}

// TestCallVectorIntObjectRequestTypeID is TL type id of TestCallVectorIntObjectRequest.
const TestCallVectorIntObjectRequestTypeID = 0xca57f472

// Ensuring interfaces in compile-time for TestCallVectorIntObjectRequest.
var (
	_ bin.Encoder     = &TestCallVectorIntObjectRequest{}
	_ bin.Decoder     = &TestCallVectorIntObjectRequest{}
	_ bin.BareEncoder = &TestCallVectorIntObjectRequest{}
	_ bin.BareDecoder = &TestCallVectorIntObjectRequest{}
)

func (t *TestCallVectorIntObjectRequest) Zero() bool {
	if t == nil {
		return true
	}
	if !(t.X == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (t *TestCallVectorIntObjectRequest) String() string {
	if t == nil {
		return "TestCallVectorIntObjectRequest(nil)"
	}
	type Alias TestCallVectorIntObjectRequest
	return fmt.Sprintf("TestCallVectorIntObjectRequest%+v", Alias(*t))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*TestCallVectorIntObjectRequest) TypeID() uint32 {
	return TestCallVectorIntObjectRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*TestCallVectorIntObjectRequest) TypeName() string {
	return "testCallVectorIntObject"
}

// TypeInfo returns info about TL type.
func (t *TestCallVectorIntObjectRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "testCallVectorIntObject",
		ID:   TestCallVectorIntObjectRequestTypeID,
	}
	if t == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "X",
			SchemaName: "x",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (t *TestCallVectorIntObjectRequest) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode testCallVectorIntObject#ca57f472 as nil")
	}
	b.PutID(TestCallVectorIntObjectRequestTypeID)
	return t.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (t *TestCallVectorIntObjectRequest) EncodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode testCallVectorIntObject#ca57f472 as nil")
	}
	b.PutInt(len(t.X))
	for idx, v := range t.X {
		if err := v.EncodeBare(b); err != nil {
			return fmt.Errorf("unable to encode bare testCallVectorIntObject#ca57f472: field x element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (t *TestCallVectorIntObjectRequest) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode testCallVectorIntObject#ca57f472 to nil")
	}
	if err := b.ConsumeID(TestCallVectorIntObjectRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode testCallVectorIntObject#ca57f472: %w", err)
	}
	return t.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (t *TestCallVectorIntObjectRequest) DecodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode testCallVectorIntObject#ca57f472 to nil")
	}
	{
		headerLen, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode testCallVectorIntObject#ca57f472: field x: %w", err)
		}

		if headerLen > 0 {
			t.X = make([]TestInt, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value TestInt
			if err := value.DecodeBare(b); err != nil {
				return fmt.Errorf("unable to decode bare testCallVectorIntObject#ca57f472: field x: %w", err)
			}
			t.X = append(t.X, value)
		}
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (t *TestCallVectorIntObjectRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if t == nil {
		return fmt.Errorf("can't encode testCallVectorIntObject#ca57f472 as nil")
	}
	b.ObjStart()
	b.PutID("testCallVectorIntObject")
	b.FieldStart("x")
	b.ArrStart()
	for idx, v := range t.X {
		if err := v.EncodeTDLibJSON(b); err != nil {
			return fmt.Errorf("unable to encode testCallVectorIntObject#ca57f472: field x element with index %d: %w", idx, err)
		}
	}
	b.ArrEnd()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (t *TestCallVectorIntObjectRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if t == nil {
		return fmt.Errorf("can't decode testCallVectorIntObject#ca57f472 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("testCallVectorIntObject"); err != nil {
				return fmt.Errorf("unable to decode testCallVectorIntObject#ca57f472: %w", err)
			}
		case "x":
			if err := b.Arr(func(b tdjson.Decoder) error {
				var value TestInt
				if err := value.DecodeTDLibJSON(b); err != nil {
					return fmt.Errorf("unable to decode testCallVectorIntObject#ca57f472: field x: %w", err)
				}
				t.X = append(t.X, value)
				return nil
			}); err != nil {
				return fmt.Errorf("unable to decode testCallVectorIntObject#ca57f472: field x: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetX returns value of X field.
func (t *TestCallVectorIntObjectRequest) GetX() (value []TestInt) {
	if t == nil {
		return
	}
	return t.X
}

// TestCallVectorIntObject invokes method testCallVectorIntObject#ca57f472 returning error if any.
func (c *Client) TestCallVectorIntObject(ctx context.Context, x []TestInt) (*TestVectorIntObject, error) {
	var result TestVectorIntObject

	request := &TestCallVectorIntObjectRequest{
		X: x,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
