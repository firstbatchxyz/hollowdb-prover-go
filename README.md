<p align="center">
  <img src="https://raw.githubusercontent.com/firstbatchxyz/hollowdb/master/logo.svg" alt="logo" width="142">
</p>

<p align="center">
  <h1 align="center">
    HollowDB Prover
  </h1>
  <p align="center">
    <i>Proof generator package for HollowDB.</i>
  </p>
</p>

<p align="center">
    <a href="https://opensource.org/licenses/MIT" target="_blank">
        <img alt="License: MIT" src="https://img.shields.io/badge/license-MIT-yellow.svg">
    </a>
    <a href="https://docs.hollowdb.xyz/zero-knowledge-proofs/hollowdb-prover" target="_blank">
        <img alt="Docs" src="https://img.shields.io/badge/docs-hollowdb-3884FF.svg?logo=gitbook">
    </a>
    <a href="https://github.com/firstbatchxyz/hollowdb" target="_blank">
        <img alt="GitHub: HollowDB" src="https://img.shields.io/badge/github-hollowdb-5C3EFE?logo=github">
    </a>
    <a href="https://discord.gg/2wuU9ym6fq" target="_blank">
        <img alt="Discord" src="https://dcbadge.vercel.app/api/server/2wuU9ym6fq?style=flat">
    </a>
</p>

## Usage

We use Go RapidSnark to create witnesses and generate proofs. As for the WASM execution engine, it currently uses [RapidSnark Wasmer](https://github.com/iden3/go-rapidsnark/tree/main/witness/wasmer). To execute Go code, simply:

```sh
go run ./cmd/main.go
```

We do not have verifier implemented, as in our scenario we do that in the contract-level with SnarkJS.

## Testing

Running the Go tests will generate a proof and public signals under `out` folder. These can be verified using SnarkJS, if you have it installed globally.

```sh
# run tests & generate proofs
go test ./... -test.v
# or
yarn test

# verify Groth16 proof
snarkjs g16v ./circuits/verification_key.json ./out/public.json ./out/proof.json
# or
yarn verify

```
