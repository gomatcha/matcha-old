package text

import "image/color"

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

type StyleKey int

const (
	StyleKeyAlignment StyleKey = iota
	StyleKeyStrikethroughStyle
	StyleKeyStrikethroughColor
	StyleKeyUnderlineStyle
	StyleKeyUnderlineColor
	StyleKeyFont
	StyleKeyHyphenation
	StyleKeyLineHeightMultiple
	StyleKeyMaxLines
	StyleKeyTextColor
	StyleKeyWrap
	StyleKeyTruncation
	StyleKeyTruncationString
)

type Style struct {
	attributes map[StyleKey]interface{}
	cleared    map[StyleKey]bool
}

func (f *Style) Map() map[StyleKey]interface{} {
	return f.attributes
}

func (f *Style) Clear(k StyleKey) {
	if f.cleared == nil || f.attributes == nil {
		f.attributes = map[StyleKey]interface{}{}
		f.cleared = map[StyleKey]bool{}
	}

	delete(f.attributes, k)
	f.cleared[k] = true
}

func (f *Style) Get(k StyleKey) interface{} {
	v, ok := f.attributes[k]
	if ok {
		return v
	}
	switch k {
	case StyleKeyAlignment:
		return AlignmentLeft
	case StyleKeyStrikethroughStyle:
		return StrikethroughStyleNone
	case StyleKeyStrikethroughColor:
		return color.Gray{0}
	case StyleKeyUnderlineStyle:
		return UnderlineStyleNone
	case StyleKeyUnderlineColor:
		return color.Gray{0}
	case StyleKeyFont:
		return nil // TODO(KD): what should the default font be?
	case StyleKeyHyphenation:
		return 0
	case StyleKeyLineHeightMultiple:
		return 1
	case StyleKeyMaxLines:
		return 0
	case StyleKeyTextColor:
		return color.Gray{0}
	case StyleKeyWrap:
		return WrapWord
	case StyleKeyTruncation:
		return TruncationNone
	case StyleKeyTruncationString:
		return "…"
	}
	return nil
}

func (f *Style) Set(k StyleKey, v interface{}) {
	if f.cleared == nil || f.attributes == nil {
		f.attributes = map[StyleKey]interface{}{}
		f.cleared = map[StyleKey]bool{}
	}

	f.attributes[k] = v
	delete(f.cleared, k)
}

func (f *Style) Copy() *Style {
	c := &Style{
		attributes: map[StyleKey]interface{}{},
		cleared:    map[StyleKey]bool{},
	}
	for k, v := range f.attributes {
		c.attributes[k] = v
	}
	for k, v := range f.cleared {
		c.attributes[k] = v
	}
	return c
}

func (f *Style) Update(u *Style) {
	for k, v := range u.attributes {
		f.attributes[k] = v
	}
	for k := range u.cleared {
		delete(f.attributes, k)
	}
}

func (f *Style) Alignment() Alignment {
	return f.Get(StyleKeyAlignment).(Alignment)
}

func (f *Style) SetAlignment(v Alignment) {
	f.Set(StyleKeyAlignment, v)
}

func (f *Style) DeleteAlignment() {
	f.Clear(StyleKeyAlignment)
}

func (f *Style) StrikethroughStyle() StrikethroughStyle {
	return f.Get(StyleKeyStrikethroughStyle).(StrikethroughStyle)
}

func (f *Style) SetStrikethroughStyle(v StrikethroughStyle) {
	f.Set(StyleKeyStrikethroughStyle, v)
}

func (f *Style) ClearStrikethroughStyle() {
	f.Clear(StyleKeyStrikethroughStyle)
}

func (f *Style) StrikethroughColor() color.Color {
	return f.Get(StyleKeyStrikethroughColor).(color.Color)
}

func (f *Style) SetStrikethroughColor(v color.Color) {
	f.Set(StyleKeyStrikethroughColor, v)
}

func (f *Style) ClearStrikethroughColor() {
	f.Clear(StyleKeyStrikethroughColor)
}

func (f *Style) UnderlineStyle() UnderlineStyle {
	return f.Get(StyleKeyUnderlineStyle).(UnderlineStyle)
}

func (f *Style) SetUnderlineStyle(v UnderlineStyle) {
	f.Set(StyleKeyUnderlineStyle, v)
}

func (f *Style) ClearUnderlineStyle() {
	f.Clear(StyleKeyUnderlineStyle)
}

func (f *Style) UnderlineColor() color.Color {
	return f.Get(StyleKeyUnderlineColor).(color.Color)
}

func (f *Style) SetUnderlineColor(v color.Color) {
	f.Set(StyleKeyUnderlineColor, v)
}

func (f *Style) ClearUnderlineColor() {
	f.Clear(StyleKeyUnderlineColor)
}

func (f *Style) Font() Font {
	return f.Get(StyleKeyFont).(Font)
}

func (f *Style) SetFont(v Font) {
	f.Set(StyleKeyFont, v)
}

func (f *Style) ClearFont() {
	f.Clear(StyleKeyFont)
}

func (f *Style) Hyphenation() float64 {
	return f.Get(StyleKeyHyphenation).(float64)
}

func (f *Style) SetHyphenation(v float64) {
	f.Set(StyleKeyHyphenation, v)
}

func (f *Style) ClearHyphenation() {
	f.Clear(StyleKeyHyphenation)
}

func (f *Style) LineHeightMultiple() float64 {
	return f.Get(StyleKeyLineHeightMultiple).(float64)
}

func (f *Style) SetLineHeightMultiple(v float64) {
	f.Set(StyleKeyLineHeightMultiple, v)
}

func (f *Style) ClearLineHeightMultiple() {
	f.Clear(StyleKeyLineHeightMultiple)
}

func (f *Style) MaxLines() int {
	return f.Get(StyleKeyMaxLines).(int)
}

func (f *Style) SetMaxLines(v int) {
	f.Set(StyleKeyMaxLines, v)
}

func (f *Style) ClearMaxLines() {
	f.Clear(StyleKeyMaxLines)
}

func (f *Style) TextColor() color.Color {
	return f.Get(StyleKeyTextColor).(color.Color)
}

func (f *Style) SetTextColor(v color.Color) {
	f.Set(StyleKeyTextColor, v)
}

func (f *Style) ClearTextColor() {
	f.Clear(StyleKeyTextColor)
}

func (f *Style) Wrap() Wrap {
	return f.Get(StyleKeyWrap).(Wrap)
}

func (f *Style) SetWrap(v Wrap) {
	f.Set(StyleKeyWrap, v)
}

func (f *Style) ClearWrap() {
	f.Clear(StyleKeyWrap)
}

func (f *Style) Truncation() Truncation {
	return f.Get(StyleKeyTruncation).(Truncation)
}

func (f *Style) SetTruncation(v Truncation) {
	f.Set(StyleKeyTruncation, v)
}

func (f *Style) ClearTruncation() {
	f.Clear(StyleKeyTruncation)
}

func (f *Style) TruncationString() string {
	return f.Get(StyleKeyTruncationString).(string)
}

func (f *Style) SetTruncationString(v string) {
	f.Set(StyleKeyTruncationString, v)
}

func (f *Style) ClearTruncationString() {
	f.Clear(StyleKeyTruncationString)
}
