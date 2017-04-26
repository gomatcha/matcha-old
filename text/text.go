package text

import (
	"github.com/overcyn/mochi"
	"image/color"
	"mochi/bridge"
)

type Alignment int

const (
	AlignmentLeft Alignment = iota
	AlignmentRight
	AlignmentCenter
	AlignmentJustified
)

type StrikethroughStyle int

const (
	StrikethroughStyleNone StrikethroughStyle = iota
	StrikethroughStyleSingle
	StrikethroughStyleDouble
	StrikethroughStyleThick
	StrikethroughStyleDotted
	StrikethroughStyleDashed
)

type UnderlineStyle int

const (
	UnderlineStyleNone UnderlineStyle = iota
	UnderlineStyleSingle
	UnderlineStyleDouble
	UnderlineStyleThick
	UnderlineStyleDotted
	UnderlineStyleDashed
)

// TODO(KD): Rethink how to do this.
type Font struct {
	Family string
	Face   string
	Size   float64
}

type Wrap int

const (
	WrapNone Wrap = iota
	WrapWord
	WrapCharacter
)

type Truncation int

const (
	TruncationNone Truncation = iota
	TruncationStart
	TruncationMiddle
	TruncationEnd
)

type FormatKey int

const (
	FormatKeyAlignment FormatKey = iota
	FormatKeyStrikethroughStyle
	FormatKeyStrikethroughColor
	FormatKeyUnderlineStyle
	FormatKeyUnderlineColor
	FormatKeyFont
	FormatKeyHyphenation
	FormatKeyLineHeightMultiple
	FormatKeyMaxLines
	FormatKeyTextColor
	FormatKeyWrap
	FormatKeyTruncation
	FormatKeyTruncationString
)

type Format struct {
	attributes map[FormatKey]interface{}
}

func (f *Format) Map() map[FormatKey]interface{} {
	return f.attributes
}

func (f *Format) Del(k FormatKey) {
	delete(f.attributes, k)
}

func (f *Format) Get(k FormatKey) interface{} {
	v, ok := f.attributes[k]
	if ok {
		return v
	}
	switch k {
	case FormatKeyAlignment:
		return AlignmentLeft
	case FormatKeyStrikethroughStyle:
		return StrikethroughStyleNone
	case FormatKeyStrikethroughColor:
		return color.Gray{0}
	case FormatKeyUnderlineStyle:
		return UnderlineStyleNone
	case FormatKeyUnderlineColor:
		return color.Gray{0}
	case FormatKeyFont:
		return nil // TODO(KD): what should the default font be?
	case FormatKeyHyphenation:
		return 0
	case FormatKeyLineHeightMultiple:
		return 1
	case FormatKeyMaxLines:
		return 0
	case FormatKeyTextColor:
		return color.Gray{0}
	case FormatKeyWrap:
		return WrapWord
	case FormatKeyTruncation:
		return TruncationNone
	case FormatKeyTruncationString:
		return "…"
	}
	return nil
}

func (f *Format) Set(k FormatKey, v interface{}) {
	if f.attributes == nil {
		f.attributes = map[FormatKey]interface{}{}
	}
	f.attributes[k] = v
}

func (f *Format) Alignment() Alignment {
	return f.Get(FormatKeyAlignment).(Alignment)
}

func (f *Format) SetAlignment(v Alignment) {
	f.Set(FormatKeyAlignment, v)
}

func (f *Format) StrikethroughStyle() StrikethroughStyle {
	return f.Get(FormatKeyStrikethroughStyle).(StrikethroughStyle)
}

func (f *Format) SetStrikethroughStyle(v StrikethroughStyle) {
	f.Set(FormatKeyStrikethroughStyle, v)
}

func (f *Format) StrikethroughColor() color.Color {
	return f.Get(FormatKeyStrikethroughColor).(color.Color)
}

func (f *Format) SetStrikethroughColor(v color.Color) {
	f.Set(FormatKeyStrikethroughColor, v)
}

func (f *Format) UnderlineStyle() UnderlineStyle {
	return f.Get(FormatKeyUnderlineStyle).(UnderlineStyle)
}

func (f *Format) SetUnderlineStyle(v UnderlineStyle) {
	f.Set(FormatKeyUnderlineStyle, v)
}

func (f *Format) UnderlineColor() color.Color {
	return f.Get(FormatKeyUnderlineColor).(color.Color)
}

func (f *Format) SetUnderlineColor(v color.Color) {
	f.Set(FormatKeyUnderlineColor, v)
}

func (f *Format) Font() Font {
	return f.Get(FormatKeyFont).(Font)
}

func (f *Format) SetFont(v Font) {
	f.Set(FormatKeyFont, v)
}

func (f *Format) Hyphenation() float64 {
	return f.Get(FormatKeyHyphenation).(float64)
}

func (f *Format) SetHyphenation(v float64) {
	f.Set(FormatKeyHyphenation, v)
}

func (f *Format) LineHeightMultiple() float64 {
	return f.Get(FormatKeyLineHeightMultiple).(float64)
}

func (f *Format) SetLineHeightMultiple(v float64) {
	f.Set(FormatKeyLineHeightMultiple, v)
}

func (f *Format) MaxLines() int {
	return f.Get(FormatKeyMaxLines).(int)
}

func (f *Format) SetMaxLines(v int) {
	f.Set(FormatKeyMaxLines, v)
}

func (f *Format) TextColor() color.Color {
	return f.Get(FormatKeyTextColor).(color.Color)
}

func (f *Format) SetTextColor(v color.Color) {
	f.Set(FormatKeyTextColor, v)
}

func (f *Format) Wrap() Wrap {
	return f.Get(FormatKeyWrap).(Wrap)
}

func (f *Format) SetWrap(v Wrap) {
	f.Set(FormatKeyWrap, v)
}

func (f *Format) Truncation() Truncation {
	return f.Get(FormatKeyTruncation).(Truncation)
}

func (f *Format) SetTruncation(v Truncation) {
	f.Set(FormatKeyTruncation, v)
}

func (f *Format) TruncationString() string {
	return f.Get(FormatKeyTruncationString).(string)
}

func (f *Format) SetTruncationString(v string) {
	f.Set(FormatKeyTruncationString, v)
}

type Text struct {
	str    string
	format *Format
}

func (ts *Text) String() string {
	return ts.str
}

func (ts *Text) SetString(text string) {
	ts.str = text
}

func (ts *Text) Format() *Format {
	if ts.format == nil {
		ts.format = &Format{}
	}
	return ts.format
}

func (ts *Text) SetFormat(f *Format) {
	ts.format = f
}

func (ts *Text) Size(max mochi.Point) mochi.Point {
	return bridge.Root().Call("sizeForAttributedString:minSize:maxSize:", bridge.Interface(ts), nil, bridge.Interface(max)).ToInterface().(mochi.Point)
}
