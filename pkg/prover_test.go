package hollowprover

import (
	"testing"
)

func TestProver(t *testing.T) {
	const wasmPath = "../circuits/hollow-authz.wasm"
	const pkeyPath = "../circuits/prover-key.zkey"

	t.Log("Creating prover.")
	prover, err := NewProver(wasmPath, pkeyPath)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("Computing preimage from secret.")
	secret := "my lovely secret key"
	preimage, err := HashToGroup(secret)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("Generating proof.")
	proof, publicSignals, err := prover.Prove(preimage, map[string]interface{}{
		"foo":     123,
		"hollow":  true,
		"awesome": "yes",
	}, map[string]interface{}{
		"foo":     123789789,
		"hollow":  false,
		"awesome": "yes",
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Log("Exporting proof and public signals.")
	if err := exportFullproof(proof, publicSignals); err != nil {
		t.Fatal(err)
	}
}
