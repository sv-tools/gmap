# gmap

Simple Generic wrapper for [sync.Map](https://pkg.go.dev/sync#Map) 

The repository is archived because the feature is fully implemented and will be replaced by sync v2 package in the future.

## Usage

```shell
go get github.com/sv-tools/gmap
```

```go
	m := gmap.New[string, int]()
	m.Store("foo", 42)
	value, _ := m.Load("foo")
	fmt.Printf("value = (%T) %v", value, value)
	// Output: value = (int) 42
```

## Benchmarks

```shell
% go test -bench=. -benchmem ./...
goos: darwin
goarch: arm64
pkg: github.com/sv-tools/gmap
BenchmarkSyncMap-8       6721870               153.4 ns/op            40 B/op          3 allocs/op
BenchmarkMap-8           7920562               154.2 ns/op            41 B/op          3 allocs/op
PASS
```
