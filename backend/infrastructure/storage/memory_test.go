package storage

import (
	"testing"
)

func TestMemoryStorage(t *testing.T) {
	store := NewMemoryStorage()

	store.Save("a", "b")

	if data, _ := store.Fetch("a"); data != "b" {
		t.Errorf("storage returned invalid data: got %v want %v", data, "b")
	}

	store.Delete("a")

	if _, err := store.Fetch("a"); err.Error() != "key not found" {
		t.Error("Invalid error message")
	}
}
