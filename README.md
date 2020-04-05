# json-benchmarks ![Go](https://github.com/raschmitt/json-benchmarks/workflows/Go/badge.svg?branch=master)

Repository dedicated to benchmark the encoding/decoding of jsons in go with the following packages:

- [encoding/json](https://golang.org/pkg/encoding/json/)
- [easyjson](https://godoc.org/github.com/mailru/easyjson)

## Tests

### Assertion

To run assertion tests run:

`go test -v`

### Assertion

To run benchmark tests run:

`go test -bench=.`

## References

- [How to Test in Go](https://semaphoreci.com/community/tutorials/how-to-test-in-go)
- [How to write benchmarks in Go](https://dave.cheney.net/2013/06/30/how-to-write-benchmarks-in-go)
