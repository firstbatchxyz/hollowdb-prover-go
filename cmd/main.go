package main

import (
	hollowprover "hollow-prover/pkg"
	"log"
	"os"
)

const wasmFName = "../res/circuit.wasm"
const zkeyFName = "../res/proverkey.zkey"

const proofFName = "./out/proof.json"
const publicInputsFName = "./out/public.json"

func readCircuitFiles() (wasm []byte, zkey []byte) {
	// Prover key
	zkeyBytes, err := os.ReadFile(zkeyFName)
	if err != nil {
		log.Fatal("Failed to read prover key.\n", err)
	}
	// WASM file of a Circuit, created by Circom.
	wasmBytes, err := os.ReadFile(wasmFName)
	if err != nil {
		log.Fatal("Failed to read WASM file.\n", err)
	}

	return wasmBytes, zkeyBytes
}

func exportFullproof(proof string, publicInputs string) {
	log.Println("Exporting proof and public signals.")
	if err := os.WriteFile(proofFName, []byte(proof), 0644); err != nil {
		log.Fatal("Could not write proof to file.\n", err)
	}
	if err := os.WriteFile(publicInputsFName, []byte(publicInputs), 0644); err != nil {
		log.Fatal("Could not write public signals to file.\n", err)
	}
}

func main() {
	wasm, zkey := readCircuitFiles()
	proof, publicInputs := hollowprover.FullProve(wasm, zkey, "901231230202", "3279874327432432781189", "9811872342347234789723")
	exportFullproof(proof, publicInputs)

}
