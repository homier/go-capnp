// Code generated by capnpc-go. DO NOT EDIT.

package stream

import (
	capnp "github.com/homier/go-capnp/v3"
	text "github.com/homier/go-capnp/v3/encoding/text"
	schemas "github.com/homier/go-capnp/v3/schemas"
)

type StreamResult capnp.Struct

// StreamResult_TypeID is the unique identifier for the type StreamResult.
const StreamResult_TypeID = 0x995f9a3377c0b16e

func NewStreamResult(s *capnp.Segment) (StreamResult, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return StreamResult(st), err
}

func NewRootStreamResult(s *capnp.Segment) (StreamResult, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return StreamResult(st), err
}

func ReadRootStreamResult(msg *capnp.Message) (StreamResult, error) {
	root, err := msg.Root()
	return StreamResult(root.Struct()), err
}

func (s StreamResult) String() string {
	str, _ := text.Marshal(0x995f9a3377c0b16e, capnp.Struct(s))
	return str
}

func (s StreamResult) EncodeAsPtr(seg *capnp.Segment) capnp.Ptr {
	return capnp.Struct(s).EncodeAsPtr(seg)
}

func (StreamResult) DecodeFromPtr(p capnp.Ptr) StreamResult {
	return StreamResult(capnp.Struct{}.DecodeFromPtr(p))
}

func (s StreamResult) ToPtr() capnp.Ptr {
	return capnp.Struct(s).ToPtr()
}
func (s StreamResult) IsValid() bool {
	return capnp.Struct(s).IsValid()
}

func (s StreamResult) Message() *capnp.Message {
	return capnp.Struct(s).Message()
}

func (s StreamResult) Segment() *capnp.Segment {
	return capnp.Struct(s).Segment()
}

// StreamResult_List is a list of StreamResult.
type StreamResult_List = capnp.StructList[StreamResult]

// NewStreamResult creates a new list of StreamResult.
func NewStreamResult_List(s *capnp.Segment, sz int32) (StreamResult_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return capnp.StructList[StreamResult](l), err
}

// StreamResult_Future is a wrapper for a StreamResult promised by a client call.
type StreamResult_Future struct{ *capnp.Future }

func (f StreamResult_Future) Struct() (StreamResult, error) {
	p, err := f.Future.Ptr()
	return StreamResult(p.Struct()), err
}

const schema_86c366a91393f3f8 = "x\xda\x12\x88p`1\xe4\xcdgb`\x0a\x94ae" +
	"\xfb\x9f\xb7\xf1@\xb9\xf1\xac\xf8\x99\x0c\x82\xbc\x8c\xff\x7f" +
	"|\x9e,\xbc2\xedp\x1b\x03\x0b;\x03\x83\xb0,\xe3" +
	"%aMFv\x06\xd7\xff\xc5%E\xa9\x89\xb9z\xc9" +
	"L\x89\x05y\x05V\xc1`^Pjqi\x0ecI" +
	"\x00## \x00\x00\xff\xff\x13|\x19\xd3"

func RegisterSchema(reg *schemas.Registry) {
	reg.Register(&schemas.Schema{
		String: schema_86c366a91393f3f8,
		Nodes: []uint64{
			0x995f9a3377c0b16e,
		},
		Compressed: true,
	})
}
