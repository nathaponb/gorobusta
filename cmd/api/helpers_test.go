package main

import (
	"testing"
)

func TestHasher(t *testing.T) {

	testApp := Config{}

	var table = []struct {
		s    string
		want string
	}{
		{s: "12345", want: "5994471abb01112afcc18159f6cc74b4f511b99806da59b3caf5a9c173cacfc5"},
		{s: "1@2d3R4f", want: "76c59542742880da8133b74b2a1a535a84bd3c76b40431a3c59c1239b9aed3b9"},
		{s: "1D2da45ef3ew4FF", want: "e6853609c83a23d80c46e330d48d25c1824f674a70bbbbd04433e262d94e25f7"},
	}

	for _, row := range table {

		hashed := testApp.hasher(row.s)

		if row.want != hashed {
			t.Errorf("hashed unmatchs from want field of %s", row.s)
		}
	}
}
