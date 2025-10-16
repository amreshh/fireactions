package stringid

import (
	"strconv"
	"testing"
)

func TestNew(t *testing.T) {
	id := New()
	if len(id) != StringIDLength*2 {
		t.Errorf("Expected ID length %d, got %d", StringIDLength*2, len(id))
	}

	_, err := strconv.ParseInt(id, 10, 64)
	if err == nil {
		t.Errorf("Expected ID to not be a valid integer, but it was")
	}
}

func BenchmarkNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		New()
	}
}
