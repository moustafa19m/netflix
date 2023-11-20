package stringset

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSet(t *testing.T) {
	t.Run("TestSet Positive", func(t *testing.T) {
		s := NewSet("a", "b", "c", "a", "b", "c")
		if s.Size() != 3 {
			t.Errorf("expected size of 3, got %d", s.Size())
		}
		if !s.Contains("a") {
			t.Errorf("expected set to contain 'a'")
		}
		if !s.Contains("b") {
			t.Errorf("expected set to contain 'b'")
		}
		if !s.Contains("c") {
			t.Errorf("expected set to contain 'c'")
		}
		if s.Contains("d") {
			t.Errorf("expected set to not contain 'd'")
		}
	})

	t.Run("TestSet Negative", func(t *testing.T) {
		s := NewSet("a", "a", "a", "a", "a", "a")
		if s.Size() == 7 {
			t.Errorf("expected size of 1, got %d", s.Size())
		}
		if s.Contains("b") {
			t.Errorf("expected set to contain only 'a' found 'b'")
		}
	})

	t.Run("TestSet Elements", func(t *testing.T) {
		s := NewSet("a", "b", "c", "d", "e", "f", "g")
		k := NewSet("g", "a", "e", "f", "b", "c", "d")
		if s.Size() != k.Size() {
			t.Errorf("expected size of 3, got %d", s.Size())
		}

		require.EqualValues(t, s.Elements(), k.Elements())
	})
}
