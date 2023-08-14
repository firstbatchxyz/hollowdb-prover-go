package hollowdbprover_test

import (
	"math/big"
	"testing"

	hollowprover "hollowdb-prover"
)

func TestComputeKey(t *testing.T) {
	preimage, _ := new(big.Int).SetString("123456789", 10)
	expectedKey := "0xfb849f7cf35865c838cef48782e803b2c38263e2f467799c87eff168eb4d897"
	key, err := hollowprover.ComputeKey(preimage)
	if err != nil {
		t.Fatal(err)
	}

	if key != expectedKey {
		t.Fail()
	}
}

func TestHashToGroup(t *testing.T) {
	// these test cases were prepared outside in JS
	// we are checking to see if the results match for Go
	cases := []struct {
		input    any
		expected string
	}{
		{"quick brown fox jumpes over the lazy dog", "1428051172494059108075242303790827279360348377618"},
		{"hi there", "225037454736096360469008883958883233963769495287"},
		// struct example
		{struct {
			Foo int    `json:"foo"`
			Bar bool   `json:"bar"`
			Baz string `json:"baz"`
		}{
			123,
			true,
			"zab",
		}, "456108647815456389709004505861143737447371420350"},
		// map example
		{
			map[string]any{
				"aaa": "aaa",
				"bbb": "bbb",
			}, "1250384956827933784767489576791202493093638690875",
		},
	}

	for _, test := range cases {
		hash, err := hollowprover.HashToGroup(test.input)
		if err != nil {
			t.Fatal(err)
		}
		expected, _ := new(big.Int).SetString(test.expected, 10)

		if hash.Cmp(expected) != 0 {
			t.Error("Expected hashes to match.")
		}

	}

}
