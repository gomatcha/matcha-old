package text

import (
	"fmt"
	"runtime"
	"sync"
	"unicode/utf8"

	"github.com/gogo/protobuf/proto"
	"github.com/overcyn/mochi/layout"
	"github.com/overcyn/mochi/pb"
	"github.com/overcyn/mochibridge"
	"golang.org/x/text/unicode/norm"
)

type Position struct {
	id   int64
	text *Text
}

// -1 if the position has been removed
func (p *Position) Index() int {
	p.text.positionMu.Lock()
	defer p.text.positionMu.Unlock()
	return p.text.positions[p.id]
}

type position struct {
	id    int64
	index int
}

type Text struct {
	bytes         []byte
	isRune        []bool
	isGlyph       []bool
	runeCount     int
	glyphCount    int
	positions     map[int64]int
	positionMaxId int64
	positionMu    *sync.Mutex
	//
	str   string
	style *Style
}

func New(b []byte) *Text {
	t := &Text{}
	t.bytes = b
	t.positions = map[int64]int{}
	t.normalize()
	return t
}

func (t *Text) Size(min layout.Point, max layout.Point) layout.Point {
	pbFunc := &pb.SizeFunc{
		Text:    t.EncodeProtobuf(),
		MinSize: min.EncodeProtobuf(),
		MaxSize: max.EncodeProtobuf(),
	}
	data, err := proto.Marshal(pbFunc)
	if err != nil {
		return layout.Pt(0, 0)
	}

	pointData := mochibridge.Root().Call("sizeForAttributedString:", mochibridge.Bytes(data)).ToInterface().([]byte)
	pbpoint := &pb.Point{}
	err = proto.Unmarshal(pointData, pbpoint)
	if err != nil {
		fmt.Println("size decode error", err)
		return layout.Pt(0, 0)
	}
	return layout.Pt(pbpoint.X, pbpoint.Y)
}

func (t *Text) EncodeProtobuf() *pb.Text {
	return &pb.Text{
		Text:  t.str,
		Style: t.style.EncodeProtobuf(),
	}
}

func (t *Text) ByteAt(byteIdx int) byte {
	return t.bytes[byteIdx]
}

func (t *Text) RuneAt(byteIdx int) rune {
	// Start at the position and look backwards until we find the start of the rune
	var runeStart int = -1
	for i := byteIdx; i >= 0; i -= 1 {
		isRune := t.isRune[i]
		if isRune {
			runeStart = i
			break
		}
	}

	if runeStart == -1 {
		panic("RuneAt: Couldn't find rune start")
	}

	bytes := []byte{t.bytes[runeStart]}
	// Add bytes until next rune
	for i := runeStart + 1; i < len(bytes); i++ {
		if t.isRune[i] {
			break
		}
		bytes = append(bytes, t.bytes[i])
	}
	return []rune(string(bytes))[0]
}

func (t *Text) GlyphAt(byteIdx int) string {
	// Start at the position and look backwards until we find the start of the glyph
	var glyphStart int = -1
	for i := byteIdx; i >= 0; i -= 1 {
		isGlyph := t.isGlyph[i]
		if isGlyph {
			glyphStart = i
			break
		}
	}

	if glyphStart == -1 {
		panic("GlyphAt: Couldn't find glyph start")
	}

	bytes := []byte{t.bytes[glyphStart]}
	// Add bytes until next glyph
	for i := glyphStart + 1; i < len(bytes); i++ {
		if t.isGlyph[i] {
			break
		}
		bytes = append(bytes, t.bytes[i])
	}
	return string(bytes)
}

func (t *Text) ByteIndex(byteIdx int) int {
	return 0
}

func (t *Text) RuneIndex(runeIdx int) int {
	return 0
}

func (t *Text) GlyphIndex(glyphIdx int) int {
	return 0
}

func (t *Text) ByteNextIndex(byteIdx int) int {
	return byteIdx + 1
}

func (t *Text) RuneNextIndex(byteIdx int) int {
	return 0
}

func (t *Text) GlyphNextIndex(byteIdx int) int {
	return 0
}

func (t *Text) BytePrevIndex(byteIdx int) int {
	return byteIdx - 1
}

func (t *Text) RunePrevIndex(byteIdx int) int {
	return 0
}

func (t *Text) GlyphPrevIndex(byteIdx int) int {
	return 0
}

func (t *Text) ByteCount() int {
	return len(t.str)
}

func (t *Text) RuneCount() int {
	return t.runeCount
}

func (t *Text) GlyphCount() int {
	return t.glyphCount
}

func (t *Text) ReplaceRange(minByteIdx, maxByteIdx int, new string) {
}

func (t *Text) Position(byteIdx int) *Position {
	t.positionMu.Lock()
	defer t.positionMu.Unlock()

	t.positionMaxId += 1
	t.positions[t.positionMaxId] = byteIdx

	p := &Position{
		id:   t.positionMaxId,
		text: t,
	}
	runtime.SetFinalizer(p, func(final *Position) {
		text := final.text
		text.positionMu.Lock()
		defer text.positionMu.Unlock()
		delete(text.positions, final.id)
	})
	return p
}

func (t *Text) normalize() {
	runeCount := 0
	glyphCount := 0
	isRune := make([]bool, 0, len(t.bytes))
	isGlyph := make([]bool, 0, len(t.bytes))
	bytes := make([]byte, 0, len(t.bytes))

	var iter norm.Iter
	iter.Init(norm.NFD, t.bytes)
	for !iter.Done() {
		glyph := iter.Next()
		rc := utf8.RuneCount(glyph)
		bytes = append(bytes, glyph...)

		for i := range glyph {
			isGlyph = append(isGlyph, i == 0)
		}
		for i := 0; i < rc; i++ {
			isRune = append(isGlyph, i == 0)
		}

		runeCount += rc
		glyphCount += 1
	}
	t.glyphCount = glyphCount
	t.runeCount = runeCount
	t.isGlyph = isGlyph
	t.isRune = isRune
	t.bytes = bytes
}

func (t *Text) String() string {
	if t != nil {
		return t.str
	}
	return ""
}

func (t *Text) SetString(text string) {
	t.str = text
}

func (t *Text) Style() *Style {
	if t.style == nil {
		t.style = &Style{}
	}
	return t.style
}

func (t *Text) SetStyle(f *Style) {
	t.style = f
}
