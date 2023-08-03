package hollowprover

import (
	"log"
	"os"
	"testing"
)

func exportFullproof(proof string, publicInputs string) {
	log.Println("Exporting proof and public signals.")
	if err := os.WriteFile("../out/proof.json", []byte(proof), 0644); err != nil {
		log.Fatal("Could not write proof to file.\n", err)
	}
	if err := os.WriteFile("../out/public.json", []byte(publicInputs), 0644); err != nil {
		log.Fatal("Could not write public signals to file.\n", err)
	}
}

func TestProver(t *testing.T) {
	const wasmPath = "../circuits/hollow-authz.wasm"
	const pkeyPath = "../circuits/prover_key.zkey"

	prover, err := NewProver(wasmPath, pkeyPath)
	if err != nil {
		t.Error(err)
	}

	proof, publicSignals, err := prover.Prove("901231230202", "3279874327432432781189", "9811872342347234789723")
	if err != nil {
		t.Error(err)
	}

	exportFullproof(proof, publicSignals)
}
