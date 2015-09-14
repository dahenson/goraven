package goraven

import (
	"testing"
)

func TestGetFloat64(t *testing.T) {
	f, err := getFloat64("0x001738", "0x00000000", "0x000003e8")
	if err != nil {
		t.Fatalf("Convert error: %s", err)
	}
	if f != float64(5.944) {
		t.Fatalf("Expected 5.944, got '%f'", f)
	}
}
