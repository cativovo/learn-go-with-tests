package counter

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("incrementing the counter 3 times leaves it at 3", func(t *testing.T) {
		counter := NewCounter()

		counter.Increment()
		counter.Increment()
		counter.Increment()

		assertCounter(t, counter, 3)
	})

	t.Run("it runs safely concurrently", func(t *testing.T) {
		counter := NewCounter()
		wantedCount := 1000

		var wg sync.WaitGroup
		wg.Add(wantedCount)

		for i := 0; i < wantedCount; i++ {
			go func() {
				defer wg.Done()
				counter.Increment()
			}()
		}

		wg.Wait()

		assertCounter(t, counter, wantedCount)
	})
}

func assertCounter(t *testing.T, got *Counter, want int) {
	t.Helper()

	if v := got.Value(); v != want {
		t.Errorf("got %d, want %d", v, want)
	}
}
