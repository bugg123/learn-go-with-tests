package iteration

import (
	"fmt"
	"testing"
)

func assertEquals(repeated, expected string, t *testing.T) {
	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}
func TestRepeat(t *testing.T) {
	t.Run("run a 5 times", func(t *testing.T) {
		repeated := Repeat("a", 5)
		expected := "aaaaa"
		assertEquals(repeated, expected, t)
	})

	t.Run("run asd 10 times", func(t *testing.T) {
		repeated := Repeat("asd", 10)
		expected := "asdasdasdasdasdasdasdasdasdasd"
		assertEquals(repeated, expected, t)
	})
}

func ExampleRepeat() {
	repeat := Repeat("a", 5)
	fmt.Println(repeat)
	// Output: aaaaa
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
