package goraven

import (
	"testing"
)

func TestPow10Int(t *testing.T) {
	i := pow10Int(int64(5))
	if i != int64(100000) {
		t.Fatalf("Expected 100000, got '%d'\n", i)
	}
}

func TestGetFloat64(t *testing.T) {
	f, err := getFloat64("0x001738", "0x00000000", "0x000003e8", "0x05")
	if err != nil {
		t.Fatalf("Convert error: %s\n", err)
	}
	if f != float64(5.944) {
		t.Fatalf("Expected 5.944, got '%f'\n", f)
	}

	f, err = getFloat64("0x001738", "0x00000001", "0x000003e8", "0x05")
	if err != nil {
		t.Fatalf("Convert error: %s\n", err)
	}
	if f != float64(5.944) {
		t.Fatalf("Expected 5.944, got '%f'\n", f)
	}
}

func TestFormat(t *testing.T) {
	f, err := getFloat64("0x000000003ee29330", "0x00000001", "0x00002710", "0x05")
	if err != nil {
		t.Fatalf("Convert error: %s\n", err)
	}
	if f != float64(5503.6208) {
		t.Fatalf("Expected 5503.944, got '%f'\n", f)
	}
}
