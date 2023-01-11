# Go package for Mamoru Sniffer

## Usage

See `tests/sniffer_test.go`

## Test

```shell
make test
```

## Bench

```shell
make bench
```

## Regenerate bindings

First, install `c-for-go`:
```shell
go install github.com/xlab/c-for-go@latest
```
Generate the bindings:
```shell
make generate-go-bindings
```
