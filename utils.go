package hollowprover

import (
	"encoding/json"
	"fmt"
	"math/big"

	// proof & witness calculations
	rapidsnarkProver "github.com/iden3/go-rapidsnark/prover"
	"github.com/iden3/go-rapidsnark/witness/v2"
	"github.com/iden3/go-rapidsnark/witness/wasmer"

	// hashing
	"github.com/iden3/go-iden3-crypto/poseidon"
	"golang.org/x/crypto/ripemd160"
)

// Compute the witness, returning it in binary format as if witness.wtns was being read.
func computeWitness(wasmCircuit []byte, input map[string]interface{}) ([]byte, error) {
	// create witness calculator
	calc, err := witness.NewCalculator(wasmCircuit, witness.WithWasmEngine(wasmer.NewCircom2WitnessCalculator))
	if err != nil {
		return nil, err
	}

	// calculate witness
	// we use WTNSBin in particular to feed the result directly to the prover
	wtnsBytes, err := calc.CalculateWTNSBin(input, true)
	if err != nil {
		return nil, err
	}

	return wtnsBytes, nil
}

// Generate a proof, returning the proof and public signals.
//
// The return results are of string type, and simply correspond to the JSON objects in stringified form.
func generateProof(witness []byte, proverKey []byte) (string, string, error) {
	proof, publicInputs, err := rapidsnarkProver.Groth16ProverRaw(proverKey, witness)
	if err != nil {
		return "", "", err
	}
	return proof, publicInputs, nil
}

// Compute the key that is the Poseidon hash of some preimage.
//
// The returned key is a string in hexadecimal format with 0x prefix.
func ComputeKey(preimage *big.Int) (string, error) {
	key, err := poseidon.Hash([]*big.Int{preimage})
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("0x%x", key), nil
}

// Given an input, stringifies and then hashes it and make sure the result is circuit-friendly for
// BN254 (see https://docs.circom.io/background/background/#signals-of-a-circuit).
//
// Uses Ripemd160 for the hash where 160-bit output is guaranteed to be
// circuit-friendly (i.e. within the order of the curve's scalar field).
//
// If a given value is nil, it will NOT be hashed but instead mapped to 0.
func HashToGroup(input any) (*big.Int, error) {
	if input == nil {
		// nil values are mapped to 0
		return big.NewInt(0), nil
	} else {
		// other values are first "stringified"
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

// A full-prove calculates the witness and immediately creates a proof, returning the proof along with the public signals.
//
// wasm: WASM circuit file, in bytes.
// zkey: Prover key, in bytes.
//
// The rest of the inputs are expected to be decimal strings to be converted to bigint.
// CAN THIS BE COMPILED TO BE USED BY MOBILE?
//
// func FullProve(wasm []byte, zkey []byte, preimage string, curValueHash string, nextValueHash string) (proof string, publicInputs string) {
// 	// calculate witness
// 	input := prepareInputs(preimage, curValueHash, nextValueHash)
// 	wtnsBytes := computeWitness(wasm, input)

// 	// generate proof
// 	pf, pubs := generateProof(wtnsBytes, zkey)

// 	return pf, pubs
// }
