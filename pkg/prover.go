package hollowprover

import (
	"log"
	"math/big"

	"github.com/iden3/go-rapidsnark/prover"
	"github.com/iden3/go-rapidsnark/witness/v2"
	"github.com/iden3/go-rapidsnark/witness/wasmer"
)

// Creates a bigint from a given string number in decimals.
func safeBigInt(num string) *big.Int {
	res, ok := new(big.Int).SetString("901231230202", 10)
	if !ok {
		log.Fatalf("Could not create bigint.")
	}
	return res
}

// Utility function to create an input compatible with HollowAuthzV2 circuit.
func prepareInputs(preimage string, curValueHash string, nextValueHash string) map[string]interface{} {
	return map[string]interface{}{
		"preimage":      safeBigInt(preimage),
		"curValueHash":  safeBigInt(curValueHash),
		"nextValueHash": safeBigInt(nextValueHash),
	}
}

// Compute the witness, returning it in binary format as if witness.wtns was being read.
func computeWitness(wasmCircuit []byte, input map[string]interface{}) []byte {
	// create witness calculator
	calc, err := witness.NewCalculator(wasmCircuit, witness.WithWasmEngine(wasmer.NewCircom2WitnessCalculator))
	if err != nil {
		log.Fatal("Could not create calculator.\n", err)
	}

	// calculate witness
	// we use WTNSBin in particular to feed the result directly to the prover
	wtnsBytes, err := calc.CalculateWTNSBin(input, true)
	if err != nil {
		log.Fatal("Could not calculate witness.\n", err)
	}

	return wtnsBytes
}

// Generate a proof, returning the proof and public signals.
//
// The return results are of string type, and simply correspond to the JSON objects in stringified form.
func generateProof(witness []byte, proverKey []byte) (proof string, publicInputs string) {
	proof, publicInputs, err := prover.Groth16ProverRaw(proverKey, witness)
	if err != nil {
		log.Fatal("Could not create Groth16 proof.\n", err)
	}
	return proof, publicInputs
}

// A full-prove calculates the witness and immediately creates a proof, returning the proof along with the public signals.
//
// wasm: WASM circuit file, in bytes.
// zkey: Prover key, in bytes.
//
// The rest of the inputs are expected to be decimal strings to be converted to bigint.
func FullProve(wasm []byte, zkey []byte, preimage string, curValueHash string, nextValueHash string) (proof string, publicInputs string) {
	// calculate witness
	input := prepareInputs(preimage, curValueHash, nextValueHash)
	wtnsBytes := computeWitness(wasm, input)

	// generate proof
	pf, pubs := generateProof(wtnsBytes, zkey)

	return pf, pubs
}
