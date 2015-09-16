package goraven

import (
	"testing"
)

func TestGetPrice(t *testing.T) {
	p := &PriceCluster{
		Price:          "0x00000008",
		TrailingDigits: "0x02",
	}

	price, err := p.GetPrice()
	if err != nil {
		t.Fatalf("Convert error: %s", err)
	}
	if price != 0.08 {
		t.Fatalf("Expected 10.00, got '%f'", price)
	}

	p.Price = "0x000003e8"
	price, err = p.GetPrice()
	if err != nil {
		t.Fatalf("Convert error: %s", err)
	}
	if price != 10.00 {
		t.Fatalf("Expected 10.00, got '%f'", price)
	}

}
