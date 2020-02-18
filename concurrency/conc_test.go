package concurrency

import (
	"reflect"
	"testing"
)

func mockWebsiteChecker(url string) bool {
	if url == "davidtest" {
		return false
	}
	return true
}

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"http://www.thewire.co.uk",
		"http://www.theguardian.co.uk",
		"davidtest",
	}

	want := map[string]bool{
		"http://www.thewire.co.uk":     true,
		"http://www.theguardian.co.uk": true,
		"davidtest":                    false,
	}

	got := CheckWebsites(mockWebsiteChecker, websites)

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("Wanteds %v, got %v", want, got)
	}
}
