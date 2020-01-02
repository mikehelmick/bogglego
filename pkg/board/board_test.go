package board

import (
	"testing"
	"unicode/utf8"
)

func TestHeight(t *testing.T) {
	b := New()
	if height := b.Height(); height != HEIGHT {
		t.Errorf("Height() incorrect, want %v got %v", HEIGHT, height)
	}
}

func TestWidth(t *testing.T) {
	b := New()
	if width := b.Width(); width != WIDTH {
		t.Errorf("Width() incorrect, want %v got %v", WIDTH, width)
	}
}

func TestGetAtXInvalid(t *testing.T) {
	b := New()
	expected := "X value of -1 is out of bounds must be 0 =< x < 4"
	if _, err := b.GetAt(-1, 0); err != nil {
		if err.Error() != expected {
			t.Errorf("Error message incorrect, got '%s' - want '%s'", err.Error(), expected)
		}
	}
}

func TestGetAtYInvalid(t *testing.T) {
	b := New()
	expected := "Y value of 5 is out of bounds must be 0 =< y < 4"
	if _, err := b.GetAt(0, 5); err != nil {
		if err.Error() != expected {
			t.Errorf("Error message incorrect, got '%s' - want '%s'", err.Error(), expected)
		}
	}
}

func TestGetAt(t *testing.T) {
	b := New()
	for y := 0; y < HEIGHT; y++ {
		for x := 0; x < WIDTH; x++ {
			if val, err := b.GetAt(y, x); err != nil {
				t.Errorf("Error returned when position was valid: %s", err.Error())
			} else {
				// Board is random, just check this is a character.
				if len(val) != 1 {
					t.Errorf("Expected a single character, got '%s' as position %d,%d", val, x, y)
				}
				rune, _ := utf8.DecodeRuneInString(val[0:])
				if !(rune >= 'a' && rune <= 'z') {
					t.Errorf("Rune is not in expected range of lowercase 'a' -- 'z', %v, decoded=%v", val, rune)
				}
			}
		}
	}
}
