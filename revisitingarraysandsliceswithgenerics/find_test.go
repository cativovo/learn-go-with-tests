package revisitingarraysandsliceswithgenerics

import (
	"strings"
	"testing"
)

func TestFind(t *testing.T) {
	t.Run("find first even number", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		firstEvenNumber, found := Find(numbers, func(x int) bool {
			return x%2 == 0
		})

		if !found {
			t.Error("got false, want true")
		}

		want := 2
		if firstEvenNumber != want {
			t.Errorf("got %d, want %d", firstEvenNumber, want)
		}
	})

	type Person struct {
		Name string
	}

	t.Run("Find the best programmer", func(t *testing.T) {
		people := []Person{
			{Name: "Kent Beck"},
			{Name: "Martin Fowler"},
			{Name: "Chris James"},
		}

		king, found := Find(people, func(p Person) bool {
			return strings.Contains(p.Name, "Chris")
		})

		if !found {
			t.Error("got false, want true")
		}

		want := Person{
			Name: "Chris James",
		}
		if king != want {
			t.Errorf("got %q, want %q", king, want)
		}
	})
}
