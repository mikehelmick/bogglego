package trie

import (
	"testing"
)

func TestNew(t *testing.T) {
	d := New()
	if len(d.roots) != 0 {
		t.Errorf("New trie wasn't empty.")
	}
}

func TestAddEmptyWord(t *testing.T) {
	d := New()
	err := d.AddWord("")
	if err == nil {
		t.Errorf("Expected error condition.")
	} else {
		if "Word to add can't be empty." != err.Error() {
			t.Errorf("Wrong error message.")
		}
	}
}

func TestAddWordInvalid(t *testing.T) {
	d := New()
	err := d.AddWord("word!")
	if err == nil {
		t.Errorf("Expected error condition.")
	}
}

func TestAddWord(t *testing.T) {
	d := New()
	err := d.AddWord("word")
	if err != nil {
		t.Errorf("Unexpected error: %s", err.Error())
	}
}

func TestIsPrefixEmptyWord(t *testing.T) {
	d := New()
	if d.IsPrefix("") {
		t.Errorf("Empty string was reported as a prefix, and shouldn't be.")
	}
}

func TestIsPrefix(t *testing.T) {
	d := New()
	d.AddWord("word")
	words := [3]string{"w", "wo", "wor"}
	for _, w := range words {
		if !d.IsPrefix(w) {
			t.Errorf("'%s' shold be a prefix", w)
		}
	}
	if d.IsPrefix("word") {
		t.Error("'word' shouldn't be a prefix")
	}
}

func TestIsWord(t *testing.T) {
	d := New()
	d.AddWord("word")
	words := [3]string{"w", "wo", "wor"}
	for _, w := range words {
		if d.IsWord(w) {
			t.Errorf("'%s' sholdn't be a word", w)
		}
	}
	if d.IsWord("pizza") {
		t.Error("'pizza' should be a word")
	}
}
