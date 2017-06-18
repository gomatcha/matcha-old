// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/overcyn/matcha/pb/layout/layout.proto

/*
Package layout is a generated protocol buffer package.

It is generated from these files:
	github.com/overcyn/matcha/pb/layout/layout.proto

It has these top-level messages:
	Point
	Rect
	Insets
	Guide
*/
package layout

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Point struct {
	X float64 `protobuf:"fixed64,1,opt,name=x" json:"x,omitempty"`
	Y float64 `protobuf:"fixed64,2,opt,name=y" json:"y,omitempty"`
}

func (m *Point) Reset()                    { *m = Point{} }
func (m *Point) String() string            { return proto.CompactTextString(m) }
func (*Point) ProtoMessage()               {}
func (*Point) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Point) GetX() float64 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *Point) GetY() float64 {
	if m != nil {
		return m.Y
	}
	return 0
}

type Rect struct {
	Min *Point `protobuf:"bytes,1,opt,name=min" json:"min,omitempty"`
	Max *Point `protobuf:"bytes,2,opt,name=max" json:"max,omitempty"`
}

func (m *Rect) Reset()                    { *m = Rect{} }
func (m *Rect) String() string            { return proto.CompactTextString(m) }
func (*Rect) ProtoMessage()               {}
func (*Rect) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Rect) GetMin() *Point {
	if m != nil {
		return m.Min
	}
	return nil
}

func (m *Rect) GetMax() *Point {
	if m != nil {
		return m.Max
	}
	return nil
}

type Insets struct {
	Top    float64 `protobuf:"fixed64,1,opt,name=top" json:"top,omitempty"`
	Left   float64 `protobuf:"fixed64,2,opt,name=left" json:"left,omitempty"`
	Bottom float64 `protobuf:"fixed64,3,opt,name=bottom" json:"bottom,omitempty"`
	Right  float64 `protobuf:"fixed64,4,opt,name=right" json:"right,omitempty"`
}

func (m *Insets) Reset()                    { *m = Insets{} }
func (m *Insets) String() string            { return proto.CompactTextString(m) }
func (*Insets) ProtoMessage()               {}
func (*Insets) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Insets) GetTop() float64 {
	if m != nil {
		return m.Top
	}
	return 0
}

func (m *Insets) GetLeft() float64 {
	if m != nil {
		return m.Left
	}
	return 0
}

func (m *Insets) GetBottom() float64 {
	if m != nil {
		return m.Bottom
	}
	return 0
}

func (m *Insets) GetRight() float64 {
	if m != nil {
		return m.Right
	}
	return 0
}

type Guide struct {
	Frame  *Rect   `protobuf:"bytes,1,opt,name=frame" json:"frame,omitempty"`
	Insets *Insets `protobuf:"bytes,2,opt,name=insets" json:"insets,omitempty"`
	ZIndex int64   `protobuf:"varint,3,opt,name=zIndex" json:"zIndex,omitempty"`
}

func (m *Guide) Reset()                    { *m = Guide{} }
func (m *Guide) String() string            { return proto.CompactTextString(m) }
func (*Guide) ProtoMessage()               {}
func (*Guide) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Guide) GetFrame() *Rect {
	if m != nil {
		return m.Frame
	}
	return nil
}

func (m *Guide) GetInsets() *Insets {
	if m != nil {
		return m.Insets
	}
	return nil
}

func (m *Guide) GetZIndex() int64 {
	if m != nil {
		return m.ZIndex
	}
	return 0
}

func init() {
	proto.RegisterType((*Point)(nil), "matcha.layout.Point")
	proto.RegisterType((*Rect)(nil), "matcha.layout.Rect")
	proto.RegisterType((*Insets)(nil), "matcha.layout.Insets")
	proto.RegisterType((*Guide)(nil), "matcha.layout.Guide")
}

func init() { proto.RegisterFile("github.com/overcyn/matcha/pb/layout/layout.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 280 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xcd, 0x4a, 0xc3, 0x40,
	0x10, 0xc7, 0x49, 0xf3, 0x71, 0x18, 0x3f, 0x90, 0xb1, 0x4a, 0xbc, 0x49, 0x04, 0xd1, 0x83, 0x89,
	0xe8, 0x1b, 0xf4, 0x22, 0x05, 0x85, 0xb2, 0x07, 0x0f, 0xe2, 0x65, 0x93, 0x6e, 0x9b, 0x85, 0x26,
	0x1b, 0xd2, 0x89, 0x24, 0x3e, 0x8e, 0x4f, 0x2a, 0x3b, 0xbb, 0x17, 0x0b, 0x9e, 0x32, 0xff, 0xc9,
	0x6f, 0x66, 0x7f, 0xcb, 0xc2, 0xe3, 0x56, 0x53, 0x3d, 0x94, 0x79, 0x65, 0x9a, 0xc2, 0x7c, 0xa9,
	0xbe, 0x9a, 0xda, 0xa2, 0x91, 0x54, 0xd5, 0xb2, 0xe8, 0xca, 0x62, 0x27, 0x27, 0x33, 0x90, 0xff,
	0xe4, 0x5d, 0x6f, 0xc8, 0xe0, 0x89, 0xfb, 0x9d, 0xbb, 0x66, 0x76, 0x03, 0xf1, 0xca, 0xe8, 0x96,
	0xf0, 0x18, 0x82, 0x31, 0x0d, 0xae, 0x83, 0xbb, 0x40, 0x04, 0xa3, 0x4d, 0x53, 0x3a, 0x73, 0x69,
	0xca, 0xde, 0x21, 0x12, 0xaa, 0x22, 0xbc, 0x85, 0xb0, 0xd1, 0x2d, 0x53, 0x47, 0x4f, 0xf3, 0xfc,
	0xcf, 0xa6, 0x9c, 0xd7, 0x08, 0x0b, 0x30, 0x27, 0x47, 0x9e, 0xff, 0x9f, 0x93, 0x63, 0xf6, 0x09,
	0xc9, 0xb2, 0xdd, 0x2b, 0xda, 0xe3, 0x19, 0x84, 0x64, 0x3a, 0x7f, 0xbe, 0x2d, 0x11, 0x21, 0xda,
	0xa9, 0x0d, 0x79, 0x09, 0xae, 0xf1, 0x12, 0x92, 0xd2, 0x10, 0x99, 0x26, 0x0d, 0xb9, 0xeb, 0x13,
	0xce, 0x21, 0xee, 0xf5, 0xb6, 0xa6, 0x34, 0xe2, 0xb6, 0x0b, 0xd9, 0x04, 0xf1, 0xcb, 0xa0, 0xd7,
	0x0a, 0xef, 0x21, 0xde, 0xf4, 0xb2, 0x51, 0x5e, 0xfc, 0xfc, 0x40, 0xc8, 0x5e, 0x4d, 0x38, 0x02,
	0x1f, 0x20, 0xd1, 0x6c, 0xe4, 0xe5, 0x2f, 0x0e, 0x58, 0xa7, 0x2b, 0x3c, 0x64, 0x85, 0xbe, 0x97,
	0xed, 0x5a, 0x8d, 0x2c, 0x14, 0x0a, 0x9f, 0x16, 0x57, 0x1f, 0x89, 0x1b, 0xf8, 0x99, 0x9d, 0xbe,
	0xf1, 0x82, 0x57, 0x8e, 0xab, 0x45, 0x99, 0xf0, 0x33, 0x3c, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff,
	0x00, 0x39, 0x05, 0x70, 0xba, 0x01, 0x00, 0x00,
}