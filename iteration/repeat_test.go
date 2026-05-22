package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("test repeat with count of 5", func(t *testing.T) {
		repeated := Repeat("a", 5)
		expected := "aaaaa"

		assertSame(t, expected, repeated)
	})

	t.Run("test repeat with a count of 12", func(t *testing.T) {
		repeated := Repeat("a", 12)
		expected := "aaaaaaaaaaaa"

		assertSame(t, expected, repeated)
	})
}

func assertSame(t *testing.T, expected string, actual string) {
	if expected != actual {
		t.Errorf("expected %s, got %s", expected, actual)
	}
}

func ExampleRepeat() {
	repeated := Repeat("a", 5)
	fmt.Println(repeated)
	// Output: aaaaa
}

func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat("a", 5)
	}
}
