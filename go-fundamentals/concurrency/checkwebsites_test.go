package concurrency

import (
	"reflect"
	"testing"
	"time"
)

func mockWebsiteChecker(url string) bool {
	return url != "https://www.invalid.com"
}

func slowStubWebsiteChecker(_ string) bool {
	time.Sleep(time.Millisecond * 20)
	return true
}

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"https://www.google.com",
		"https://www.yahoo.com",
		"https://www.invalid.com",
		"https://www.gmail.com",
	}

	want := map[string]bool{
		"https://www.google.com":  true,
		"https://www.yahoo.com":   true,
		"https://www.invalid.com": false,
		"https://www.gmail.com":   true,
	}
	got := CheckWebsites(mockWebsiteChecker, websites)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CheckWebsites(slowStubWebsiteChecker, urls)
	}
}
