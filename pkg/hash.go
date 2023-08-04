package hollowprover

import (
	"encoding/json"
	"math/big"

	"golang.org/x/crypto/ripemd160"
)

// Hashes a given input to a 160-bit bigint using Ripemd160
func HashToGroup(input any) (*big.Int, error) {
	if input == nil {
		// nil values are mapped to 0
		return big.NewInt(0), nil
	} else {
		jsonBytes, err := json.Marshal(input)
		if err != nil {
			return nil, err
		}

		hash := ripemd160.New()
		if _, err := hash.Write(jsonBytes); err != nil {
			return nil, err
		}
		digest := hash.Sum([]byte{})

		return new(big.Int).SetBytes(digest), nil
	}
}
