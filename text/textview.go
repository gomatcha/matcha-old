package text

import (
	"github.com/overcyn/mochi"
	"image/color"
)

type Alignment int

const (
	AlignmentLeft Alignment = iota
	AlignmentRight
	AlignmentCenter
	AlignmentJustified
)

type Decoration int

const (
	DecorationNone Decoration = iota
	DecorationStrikethrough
	DecorationUnderline
	DecorationOverline
)

type DecorationStyle int

const (
	DecorationStyleSingle DecorationStyle = iota
	DecorationStyleDouble
	DecorationStyleThick
	DecorationStyleDotted
	DecorationStyleDashed
)

type Font struct {
	Family string
	Face   string
	Size   float64
	Weight int // 100 - 900 or maybe -1 to 1
}

type TextWrap int

const (
	TextWrapNone TextWrap = iota
	TextWrapWord
	TextWrapCharacter
)

type Truncation int

const (
	TruncationNone Truncation = iota
	TruncationStart
	TruncationMiddle
	TruncationEnd
)

type AttributeKey int

const (
	AttributeKeyAlignment AttributeKey = iota
	AttributeKeyDecoration
	AttributeKeyDecorationColor
	AttributeKeyDecorationStyle
	AttributeKeyFont
	AttributeKeyHyphenation
	AttributeKeyLineHeightMultiple
	AttributeKeyMaxLines
	AttributeKeyTextColor
	AttributeKeyTextWrap
	AttributeKeyTruncation
	AttributeKeyTruncationString
)

type Format struct {
	attributes map[AttributeKey]interface{}
}

func (f *Format) Attributes() map[AttributeKey]interface{} {
	return f.attributes
}

func (f *Format) Del(k AttributeKey) {
	delete(f.attributes, k)
}

func (f *Format) Get(k AttributeKey) interface{} {
	v, ok := f.attributes[k]
	if ok {
		return v
	}
	switch k {
	case AttributeKeyAlignment:
		return AlignmentLeft
	case AttributeKeyDecoration:
		return DecorationNone
	case AttributeKeyDecorationColor:
		return color.Gray{0}
	case AttributeKeyDecorationStyle:
		return DecorationStyleSingle
	case AttributeKeyFont:
		return nil // TODO(KD): what should the default font be?
	case AttributeKeyHyphenation:
		return 0
	case AttributeKeyLineHeightMultiple:
		return 1
	case AttributeKeyMaxLines:
		return 0
	case AttributeKeyTextColor:
		return color.Gray{0}
	case AttributeKeyTextWrap:
		return TextWrapWord
	case AttributeKeyTruncation:
		return TruncationNone
	case AttributeKeyTruncationString:
		return "…"
	}
}

func (f *Format) Set(k AttributeKey, v interface{}) {
	if f.attributes == nil {
		f.attributes = map[AttributeKey]attribute{}
	}
	switch k {
	case AttributeKeyAlignment:
		f.attributes[k] = v.(Alignment)
	case AttributeKeyDecoration:
		f.attributes[k] = v.(Decoration)
	case AttributeKeyDecorationColor:
		f.attributes[k] = v.(color.Color)
	case AttributeKeyDecorationStyle:
		f.attributes[k] = v.(DecorationStyle)
	case AttributeKeyFont:
		f.attributes[k] = v.(Font)
	case AttributeKeyHyphenation:
		f.attributes[k] = v.(Hyphenation)
	case AttributeKeyLineHeightMultiple:
		f.attributes[k] = v.(int)
	case AttributeKeyMaxLines:
		f.attributes[k] = v.(int)
	case AttributeKeyTextColor:
		f.attributes[k] = v.(color.Color)
	case AttributeKeyTextWrap:
		f.attributes[k] = v.(TextWrap)
	case AttributeKeyTruncation:
		f.attributes[k] = v.(Truncation)
	case AttributeKeyTruncationString:
		f.attributes[k] = v.(string)
	}
}

func (f *Format) Alignment() Alignment {
	return f.Get(AttributeKeyAlignment).(Alignment)
}

func (f *Format) SetAlignment(v Alignment) {
	f.Set(AttributeKeyAlignment, v)
}

func (f *Format) Decoration() Decoration {
	return f.Get(AttributeKeyDecoration).(Decoration)
}

func (f *Format) SetDecoration(v Decoration) {
	f.Set(AttributeKeyDecoration, v)
}

func (f *Format) DecorationColor() color.Color {
	return f.Get(AttributeKeyDecorationColor).(color.Color)
}

func (f *Format) SetDecorationColor(v color.Color) {
	f.Set(AttributeKeyDecorationColor, v)
}

func (f *Format) DecorationStyle() DecorationStyle {
	return f.Get(AttributeKeyDecorationStyle).(DecorationStyle)
}

func (f *Format) SetDecorationStyle(v DecorationStyle) {
	f.Set(AttributeKeyDecorationStyle, v)
}

func (f *Format) Font() Font {
	return f.Get(AttributeKeyFont).(Font)
}

func (f *Format) SetFont(v Font) {
	f.Set(AttributeKeyFont, v)
}

func (f *Format) Hyphenation() float64 {
	return f.Get(AttributeKeyHyphenation).(float64)
}

func (f *Format) SetHyphenation(v float64) {
	f.Set(AttributeKeyHyphenation, v)
}

func (f *Format) LineHeightMultiple() float64 {
	return f.Get(AttributeKeyLineHeightMultiple).(float64)
}

func (f *Format) SetLineHeightMultiple(v float64) {
	f.Set(AttributeKeyLineHeightMultiple, v)
}

func (f *Format) MaxLines() int {
	return f.Get(AttributeKeyMaxLines).(int)
}

func (f *Format) SetMaxLines(v int) {
	f.Set(AttributeKeyMaxLines, v)
}

func (f *Format) TextColor() color.Color {
	return f.Get(AttributeKeyTextColor).(color.Color)
}

func (f *Format) SetTextColor(v color.Color) {
	f.Set(AttributeKeyTextColor, v)
}

func (f *Format) TextWrap() TextWrap {
	return f.Get(AttributeKeyTextWrap).(TextWrap)
}

func (f *Format) SetTextWrap(v TextWrap) {
	f.Set(AttributeKeyTextWrap, v)
}

func (f *Format) Truncation() Truncation {
	return f.Get(AttributeKeyTruncation).(Truncation)
}

func (f *Format) SetTruncation(v Truncation) {
	f.Set(AttributeKeyTruncation, v)
}

func (f *Format) TruncationString() string {
	return f.Get(AttributeKeyTruncationString).(string)
}

func (f *Format) SetTruncationString(v string) {
	f.Set(AttributeKeyTruncationString, v)
}

type FormattedText struct {
	str    string
	format Format
}

func (ts *FormattedText) String() string {
	return ts.str
}

func (ts *FormattedText) setString(text string) {
	ts.str = text
}

func (ts *FormattedText) FormatAt(idx int) *Format {
	return ts.format
}

func (ts *FormattedText) SetFormatAt(f *Format, start, end int) {
	ts.format = f
}

type TextView struct {
	Text          string
	Format        *Format
	FormattedText *FormattedText
	PaintOptions  PaintOptions
}

func NewTextView(p interface{}) *TextView {
	v, ok := p.(*TextView)
	if !ok {
		v = &TextView{}
		v.Format = Format{}
	}
	return v
}

func (v *TextView) Update(p *Node) *Node {
	n := NewNode()
	n.PaintOptions = v.PaintOptions
	n.Bridge.Name = "github.com/overcyn/mochi TextView"
	n.Bridge.State = &TextViewState{
		Text:          v.Text,
		Format:        v.Format,
		FormattedText: v.FormattedText,
	}
	return n
}

func (v *TextView) NeedsUpdate() {
	// ??
}

type TextViewState struct {
	Text          string
	Format        *Format
	FormattedText *FormattedText
}
