package hollowprover

import (
	"errors"
	"math/big"
	"os"
)

const bn254PrimeStr = "21888242871839275222246405745257275088548364400416034343698204186575808495617"

// A Groth16-based zero-knowledge prover utility to be used with HollowDB.
// Use NewProver function to create it.
type Prover struct {
	wasmBytes  []byte
	zkeyBytes  []byte
	bn254prime *big.Int
}

// Creates a Groth16-based zero-knowledge prover utility to be used with HollowDB.
//
// You will need to provide paths to a WASM circuit, and a prover key.
// You can find these files at: https://github.com/firstbatchxyz/hollowdb-prover/tree/master/circuits/groth16
//
// It is up to you to decide where to place them for your application.
// For example, in a web-app you may place under the `public` directory.
func NewProver(wasmPath string, proverKeyPath string) (*Prover, error) {
	wasmBytes, err := os.ReadFile(wasmPath)
	if err != nil {
		return nil, err
	}

	pkeyBytes, err := os.ReadFile(proverKeyPath)
	if err != nil {
		return nil, err
	}

	bn254Prime, ok := new(big.Int).SetString(bn254PrimeStr, 10)
	if !ok {
		return nil, errors.New("could not prepare BN254 prime")
	}

	return &Prover{wasmBytes, pkeyBytes, bn254Prime}, nil
}

// Generates a proof, returns (proof, publicSignals).
//
// Current value and next value can be anything, they will be hashed-to-group and then ProveHashed
// will be called to generate the actual proof.
func (prover *Prover) Prove(preimage *big.Int, curValue any, nextValue any) (string, string, error) {
	curValueHash, err := HashToGroup(curValue)
	if err != nil {
		return "", "", err
	}
	nextValueHash, err := HashToGroup(nextValue)
	if err != nil {
		return "", "", err
	}
	return prover.ProveHashed(preimage, curValueHash, nextValueHash)
}

// Generates a proof, returns (proof, publicSignals).
//
// Inputs are assumed to be hashed-to-group.
func (prover *Prover) ProveHashed(preimage *big.Int, curValueHash *big.Int, nextValueHash *big.Int) (string, string, error) {
	InputTooLargeErr := errors.New("input larger than BN254 order")
	if preimage.Cmp(prover.bn254prime) != -1 {
		return "", "", InputTooLargeErr
	}
	if curValueHash.Cmp(prover.bn254prime) != -1 {
		return "", "", InputTooLargeErr
	}
	if nextValueHash.Cmp(prover.bn254prime) != -1 {
		return "", "", InputTooLargeErr
	}

	input := map[string]interface{}{
		"preimage":      preimage,
		"curValueHash":  curValueHash,
		"nextValueHash": nextValueHash,
	}

	wtnsBytes, err := computeWitness(prover.wasmBytes, input)
	if err != nil {
		return "", "", err
	}

	proof, publicInputs, err := generateProof(wtnsBytes, prover.zkeyBytes)
	if err != nil {
		return "", "", err
	}

	return proof, publicInputs, nil
}
