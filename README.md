# Go Witness Calculator & Prover

We use Go RapidSnark to create witnesses and generate proofs. As for the WASM execution engine, it currently uses [RapidSnark Wasmer](github.com/iden3/go-rapidsnark/witness/wasmer).

## Usage

You need to have prepared the prover key and WASM circuit. You can use Circomkit for that in the working directory:

```sh
npx circomkit setup hollow-authz-v2
```

To execute Go code, simply:

```sh
go run ./cmd/main.go
```

We do not have verifier implemented, as in our scenario we do that in the contract-level with SnarkJS.
