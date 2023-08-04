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

We use [Go RapidSnark](https://github.com/iden3/go-rapidsnark) to create witnesses and generate proofs. As for the WASM execution engine, it currently uses [RapidSnark Wasmer](https://github.com/iden3/go-rapidsnark/tree/main/witness/wasmer). To create a prover:

```go
prover, err := hollowprover.NewProver(wasmPath, pkeyPath)
```

The `prove` function accepts any type for the current value and next value, where the inputs will be stringified and then hashed. The resulting string should match that of `JSON.stringify` in JavaScript. To this extend, user must be aware of the following in regard of Go's JSON marshalling:

- Maps have their keys sorted lexicographically
- Structs keys are marshalled in the order defined in the struct

The effect of this is that if you have the following object `{b: 1, a: 2}` in a Map then Go will stringify it as `{"a":2,"b":1}` but JavaScript will have it as `{"b":1,"a":2}`, resulting in different hashes! We suggest using struct for the inputs to the prover, but if you really have to use maps you can perhaps hash them elsewhere with more care on their keys, and then use `proveHashed` function.

## Testing

Running the Go tests will generate a proof and public signals under `out` folder, which can be verified using SnarkJS. You can run all tests with:

```sh
yarn test
```

which will run Go tests, and then run SnarkJS to verify the proofs. To verify generate proofs you can also type `yarn verify`. To run Go tests without SnarkJS, you can do:

```sh
go test ./test/*.go -test.v
```
