# WASM

## Goal

Illustrate how to use Go in WebAssembly.

## Usage

### WebAssembly

```shell
GOOS=js GOARCH=wasm go build -o site/main.wasm
```

```shell
cp "$(go env GOROOT)/lib/wasm/wasm_exec.js" site/
```

```shell
go run ./cmd/server/main.go
```

Open [site](http://localhost:8080/) in your browser.

#### WebAssembly System Interface

```shell
GOOS=wasip1 GOARCH=wasm go build -o wasi.wasm
```

```shell
wasmtime wasi.wasm
```
