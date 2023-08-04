package hollowprover

import (
	"os"

	"github.com/iden3/go-rapidsnark/prover"
	"github.com/iden3/go-rapidsnark/witness/v2"
	"github.com/iden3/go-rapidsnark/witness/wasmer"
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
	proof, publicInputs, err := prover.Groth16ProverRaw(proverKey, witness)
	if err != nil {
		return "", "", err
	}
	return proof, publicInputs, nil
}

func exportFullproof(proof string, publicInputs string) error {
	if err := os.WriteFile("../out/proof.json", []byte(proof), 0644); err != nil {
		return err
	}
	if err := os.WriteFile("../out/public.json", []byte(publicInputs), 0644); err != nil {
		return err
	}
	return nil
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
