package reflection

import (
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	City string
	Age  int
}

func TestWalk(t *testing.T) {
	cases := []struct {
		Name  string
		Input any
		Want  []string
	}{
		{
			Name: "struct with one string field",
			Input: struct {
				Name string
			}{
				Name: "Juan",
			},
			Want: []string{"Juan"},
		},
		{
			Name: "struct with two string fields",
			Input: struct {
				Name string
				City string
			}{
				Name: "Juan",
				City: "Baguio",
			},
			Want: []string{"Juan", "Baguio"},
		},
		{
			Name: "struct with non string field",
			Input: struct {
				Name string
				Age  int
			}{
				Name: "Juan",
				Age:  29,
			},
			Want: []string{"Juan"},
		},
		{
			Name: "nested fields",
			Input: Person{
				Name: "Juan",
				Profile: Profile{
					City: "Baguio",
					Age:  29,
				},
			},
			Want: []string{"Juan", "Baguio"},
		},
		{
			Name: "pointers to things",
			Input: &Person{
				Name: "Juan",
				Profile: Profile{
					City: "Baguio",
					Age:  29,
				},
			},
			Want: []string{"Juan", "Baguio"},
		},
		{
			Name: "slices",
			Input: []Person{
				{
					Name: "Juan",
					Profile: Profile{
						City: "Baguio",
						Age:  29,
					},
				},
				{
					Name: "Maria",
					Profile: Profile{
						City: "Manila",
						Age:  30,
					},
				},
			},
			Want: []string{"Juan", "Baguio", "Maria", "Manila"},
		},
		{
			Name: "arrays",
			Input: [2]Profile{
				{
					City: "Baguio",
					Age:  29,
				},
				{
					City: "Manila",
					Age:  30,
				},
			},
			Want: []string{"Baguio", "Manila"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			got := make([]string, 0)
			Walk(test.Input, func(input string) {
				got = append(got, input)
			})
		})
	}

	t.Run("with maps", func(t *testing.T) {
		aMap := map[string]string{
			"Dog": "Ruff",
			"Cat": "Meow",
		}

		got := make([]string, 0)
		Walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Ruff")
		assertContains(t, got, "Meow")
	})

	t.Run("with channels", func(t *testing.T) {
		profileChan := make(chan Profile)

		go func() {
			profileChan <- Profile{
				City: "Baguio",
				Age:  29,
			}
			profileChan <- Profile{
				City: "Manila",
				Age:  30,
			}
			close(profileChan)
		}()

		got := make([]string, 0)
		want := []string{"Baguio", "Manila"}

		Walk(profileChan, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Fatalf("got %v, want %v", got, want)
		}
	})

	t.Run("with function", func(t *testing.T) {
		function := func() (Profile, Profile) {
			return Profile{
					City: "Baguio",
					Age:  29,
				}, Profile{
					City: "Manila",
					Age:  30,
				}
		}

		got := make([]string, 0)
		want := []string{"Baguio", "Manila"}

		Walk(function, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Fatalf("got %v, want %v", got, want)
		}
	})
}

func assertContains(t *testing.T, haystack []string, needle string) {
	t.Helper()

	contains := false
	for _, v := range haystack {
		if v == needle {
			contains = true
			break
		}
	}

	if !contains {
		t.Fatalf("expected %v to contain %q but it didn't", haystack, needle)
	}
}
