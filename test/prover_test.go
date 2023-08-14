package hollowdbprover_test

import (
	"math/big"
	"os"
	"testing"

	hollowprover "hollowdb-prover"
)

const wasmPath = "../circuits/hollow-authz.wasm"
const pkeyPath = "../circuits/prover-key.zkey"

func TestProver(t *testing.T) {

	t.Log("Creating prover.")
	prover, err := hollowprover.Prover(wasmPath, pkeyPath)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("Computing preimage from secret.")
	secret := "my lovely secret key"
	preimage, err := hollowprover.HashToGroup(secret)
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

func BenchmarkProver(b *testing.B) {
	b.StopTimer()
	prover, err := hollowprover.Prover(wasmPath, pkeyPath)
	if err != nil {
		b.Fatal(err)
	}
	preimage := big.NewInt(12345)

	for i := 0; i < b.N; i++ {
		prev := b.Elapsed()
		b.StartTimer()
		prover.Prove(preimage, struct {
			Foo int `json:"foo"`
		}{123}, struct {
			Foo int `json:"foo"`
		}{456})
		b.StopTimer()
		b.Log(b.Elapsed() - prev)
	}
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
