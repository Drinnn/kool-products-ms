package data

import "testing"

func TestChecksValidation(t *testing.T) {
	p := &Product{
		Name:  "any_product",
		Price: 10,
		SKU:   "aaa-bbb-ccc",
	}

	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}
