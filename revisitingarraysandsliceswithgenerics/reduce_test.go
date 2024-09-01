package revisitingarraysandsliceswithgenerics

import "testing"

func TestReduce(t *testing.T) {
	t.Run("multiplication of all elements", func(t *testing.T) {
		s := []int{1, 2, 3}
		cb := func(acc, v int) int {
			return acc * v
		}
		initialValue := 1

		got := Reduce(s, cb, initialValue)
		want := 6
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})

	t.Run("concatenate strings", func(t *testing.T) {
		s := []string{"a", "b", "c"}
		cb := func(acc, v string) string {
			return acc + v
		}
		initialValue := ""

		got := Reduce(s, cb, initialValue)
		want := "abc"
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
