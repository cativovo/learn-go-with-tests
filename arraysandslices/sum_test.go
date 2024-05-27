package arraysandslices

import "testing"

func TestSum(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	got := Sum(nums)
	want := 15

	if got != want {
		t.Fatalf("got %v, want %v, nums %v", got, want, nums)
	}
}
