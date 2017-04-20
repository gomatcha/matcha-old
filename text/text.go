package text

import (
	_ "fmt"
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
	AttributeKeyStrikethroughStyle
	AttributeKeyStrikethroughColor
	AttributeKeyUnderlineStyle
	AttributeKeyUnderlineColor
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
	case AttributeKeyStrikethroughStyle:
		return StrikethroughStyleNone
	case AttributeKeyStrikethroughColor:
		return color.Gray{0}
	case AttributeKeyUnderlineStyle:
		return UnderlineStyleNone
	case AttributeKeyUnderlineColor:
		return color.Gray{0}
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
	return nil
}

func (f *Format) Set(k AttributeKey, v interface{}) {
	if f.attributes == nil {
		f.attributes = map[AttributeKey]interface{}{}
	}
	switch k {
	case AttributeKeyAlignment:
		f.attributes[k] = v.(Alignment)
	case AttributeKeyStrikethroughStyle:
		f.attributes[k] = v.(StrikethroughStyle)
	case AttributeKeyStrikethroughColor:
		f.attributes[k] = v.(color.Color)
	case AttributeKeyUnderlineStyle:
		f.attributes[k] = v.(UnderlineStyle)
	case AttributeKeyUnderlineColor:
		f.attributes[k] = v.(color.Color)
	case AttributeKeyFont:
		f.attributes[k] = v.(Font)
	case AttributeKeyHyphenation:
		f.attributes[k] = v.(float64)
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

func (f *Format) StrikethroughStyle() StrikethroughStyle {
	return f.Get(AttributeKeyStrikethroughStyle).(StrikethroughStyle)
}

func (f *Format) SetStrikethroughStyle(v StrikethroughStyle) {
	f.Set(AttributeKeyStrikethroughStyle, v)
}

func (f *Format) StrikethroughColor() color.Color {
	return f.Get(AttributeKeyStrikethroughColor).(color.Color)
}

func (f *Format) SetStrikethroughColor(v color.Color) {
	f.Set(AttributeKeyStrikethroughColor, v)
}

func (f *Format) UnderlineStyle() UnderlineStyle {
	return f.Get(AttributeKeyUnderlineStyle).(UnderlineStyle)
}

func (f *Format) SetUnderlineStyle(v UnderlineStyle) {
	f.Set(AttributeKeyUnderlineStyle, v)
}

func (f *Format) UnderlineColor() color.Color {
	return f.Get(AttributeKeyUnderlineColor).(color.Color)
}

func (f *Format) SetUnderlineColor(v color.Color) {
	f.Set(AttributeKeyUnderlineColor, v)
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
	format *Format
}

func (ts *FormattedText) String() string {
	return ts.str
}

func (ts *FormattedText) setString(text string) {
	ts.str = text
}

func (ts *FormattedText) Format() *Format {
	if ts.format == nil {
		ts.format = &Format{}
	}
	return ts.format
}

func (ts *FormattedText) SetFormat(f *Format) {
	ts.format = f
}

type textViewLayouter struct {
	formattedText *FormattedText
}

func (l *textViewLayouter) Layout(ctx *mochi.LayoutContext) (mochi.Guide, map[interface{}]mochi.Guide) {
	bridge.Root()
	// bridge.Root().Call("sizeForAttributedString:minSize:maxSize:", bridge.Interface(l.formattedText), bridge.Interface(minSize), bridge.Interface(maxSize))
	// val := bridge.Call("Size", bridge.WithGo(blah));

	g := mochi.Guide{Frame: mochi.Rt(0, 0, ctx.MinSize.X, ctx.MinSize.Y)}
	return g, nil
}

type TextView struct {
	marker        mochi.Marker
	Text          string
	Format        *Format
	FormattedText *FormattedText
	PaintOptions  mochi.PaintOptions
}

func NewTextView(p interface{}) *TextView {
	v, ok := p.(*TextView)
	if !ok {
		v = &TextView{}
		v.Format = &Format{}
	}
	return v
}

func (v *TextView) Mount(m mochi.Marker) {
	v.marker = m
}

func (v *TextView) Update(p *mochi.Node) *mochi.Node {
	ft := v.FormattedText
	if ft == nil {
		ft = &FormattedText{
			str:    v.Text,
			format: v.Format,
		}
	}

	n := mochi.NewNode()
	// n.Layouter = &textViewLayouter{formattedText: ft}
	n.PaintOptions = v.PaintOptions
	n.Bridge.Name = "github.com/overcyn/mochi TextView"
	n.Bridge.State = &TextViewState{
		FormattedText: ft,
	}
	return n
}

func (v *TextView) Unmount() {
	v.marker = nil
}

type TextViewState struct {
	FormattedText *FormattedText
}
