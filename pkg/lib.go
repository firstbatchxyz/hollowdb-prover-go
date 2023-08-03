package hollowprover

import (
	"errors"
	"os"
)

type Prover struct {
	wasmBytes []byte
	zkeyBytes []byte
}

// Creates a prover object.
func NewProver(wasmPath string, proverKeyPath string) (*Prover, error) {
	wasmBytes, err := os.ReadFile(wasmPath)
	if err != nil {
		return nil, err
	}

	pkeyBytes, err := os.ReadFile(proverKeyPath)
	if err != nil {
		return nil, err
	}

	return &Prover{wasmBytes, pkeyBytes}, nil
}

// Generates a proof.
//
// Returns (proof, publicSignals) and an error if any.
func (prover *Prover) Prove(preimage string, curValueHash string, nextValueHash string) (string, string, error) {
	input, err := prepareInputs(preimage, curValueHash, nextValueHash)
	if err != nil {
		return "", "", err
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

func HashToGroup() error {
	// must act like JSON.stringify and use ripemd160 to BigInt
	return errors.New("not implemented")
}

func ComputeKey() error {
	// https://github.com/iden3/go-iden3-crypto has Poseidon for BN254
	return errors.New("not implemented")
}
