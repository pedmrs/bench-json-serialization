# Simple Golang JSON Benchmark

This repository contains benchmarks to evaluate the performance, memory usage, and consistency of three popular JSON libraries in Golang:

- `encoding/json` (standard library)
- `github.com/goccy/go-json`
- `github.com/json-iterator/go`

The benchmarks focus on both serialization (marshaling) and deserialization (unmarshaling) of a struct, with additional tests for handling large datasets, parallel execution and invalid JSON.

## Running the Benchmarks

### Requirements
Ensure you have Go installed (version 1.18 or higher recommended).

Clone the repository and navigate to the project directory:

```bash
git clone https://github.com/pedmrs/bench-json-serialization.git
cd bench-json-serialization
```

## Running the Benchmarks

To run all benchmarks with memory allocation information:

```bash
go test -bench=. -benchtime=3s -benchmem
```

This will output benchmark results for all JSON libraries under different conditions (standard usage, large data, parallel processing, error handling).

## Profiling (Optional)

To gather detailed CPU or memory profiling data:

### CPU Profiling

```bash
go test -bench=. -benchtime=3s -benchmem -cpuprofile=cpu.prof
go tool pprof cpu.prof
```
This will open an interactive profiling session, where you can explore hotspots in the libraries CPU usage.

### Memory Profiling

```bash
go test -bench=. -benchtime=3s -benchmem -memprofile=mem.prof
go tool pprof mem.prof
```

This will open a session where you can analyze memory allocations during benchmarking.