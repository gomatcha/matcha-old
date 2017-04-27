package text

import "testing"

func TestXxx(t *testing.T) {
	text := New([]byte("cafe\u0301"))
	if text.ByteAt(3) != byte('e') {
		t.Error()
	}
	if text.RuneAt(4) != rune('e') {
		t.Error("Invalid rune", text.RuneAt(4))
	}
	// if text.GlyphAt(4) != "\u0301" {
	// 	t.Error("Incorrect glyph")
	// }
}
