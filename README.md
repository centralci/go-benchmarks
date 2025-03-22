# Go Benchmarks

This repository contains a collection of Go benchmarks using the standard Go testing package for evaluating and
comparing the performance of various Go implementations and patterns used at [CentralCI](https://centralci.com/).
The benchmarks focus on real-world scenarios and provide insights for making informed implementation choices in your Go
applications.

## Benchmarks

### Concurrent Maps

Compares different concurrent map implementations in Go:

- Custom `RWMutexMap` using sync.RWMutex
- Standard library's `sync.Map`

The benchmarks evaluate performance across different access patterns:

- Read-Heavy Workload: 90% reads, 10% writes
- Write-Heavy Workload: 50% reads, 50% writes
- Mixed Operations: 70% reads, 20% writes, 10% deletes

## Compression

Compares different compression implementations in Go:

- zstd compression - klauspost vs datadog (Cgo)

### More Benchmarks Coming Soon

This repository will be expanded with additional benchmarks for:

- Concurrency patterns
- Memory allocation strategies
- JSON encoding/decoding
- HTTP server implementations
- And more...

## Requirements

- Go 1.24 or later

## Running the Benchmarks

Clone this repository:

```bash
git clone https://github.com/centralci/go-benchmarks.git
cd go-benchmarks
```

### Running All Benchmarks

```bash
go test ./... -bench=.
```

### Running Specific Benchmark Groups

```bash
# Run concurrent map benchmarks
go test ./maps -bench=.

# Run a specific benchmark
go test ./maps -bench=BenchmarkRWMutexMapReadHeavy
```

### Common Benchmark Flags

Go's benchmark tool supports various flags to customize execution:

```bash
# Run with multiple CPU counts to compare scaling
go test ./maps -bench=. -cpu=1,2,4,8,16

# Include memory allocation statistics
go test ./maps -bench=. -benchmem

# Run each benchmark multiple times for more stable results
go test ./maps -bench=. -count=5

# Generate CPU profiles
go test ./maps -bench=. -cpuprofile=cpu.prof

# Generate memory profiles
go test ./maps -bench=. -memprofile=mem.prof

# Run benchmarks for longer to get more stable results
go test ./maps -bench=. -benchtime=5s
```

## Viewing Benchmark Results

### Sample Output

```
goos: linux
goarch: amd64
pkg: github.com/yourusername/go-benchmarks/maps
cpu: Intel(R) Core(TM) i7-9700K CPU @ 3.60GHz
BenchmarkRWMutexMapReadHeavy-8          10000000               119 ns/op             0 B/op          0 allocs/op
BenchmarkSyncMapReadHeavy-8             20000000                59.5 ns/op           0 B/op          0 allocs/op
BenchmarkRWMutexMapWriteHeavy-8          3000000               496 ns/op             0 B/op          0 allocs/op
BenchmarkSyncMapWriteHeavy-8             1000000              1015 ns/op            78 B/op          1 allocs/op
BenchmarkRWMutexMapMixedOps-8            5000000               295 ns/op             0 B/op          0 allocs/op
BenchmarkSyncMapMixedOps-8               2000000               604 ns/op            48 B/op          0 allocs/op
PASS
ok      github.com/yourusername/go-benchmarks/maps  10.211s
```

### Using Benchstat

For more detailed statistical analysis of benchmark results, use the `benchstat` tool:

```bash
# Install benchstat
go install golang.org/x/perf/cmd/benchstat@latest

# Compare two benchmark runs
go test ./maps -bench=. -count=5 > before.txt
# Make changes
go test ./maps -bench=. -count=5 > after.txt
benchstat before.txt after.txt
```

## Repository Structure

The repository is organized by benchmark categories:

```
go-benchmarks/
  ├── maps/                  # Map implementation benchmarks
  │   ├── concurrent_test.go # Concurrent map benchmarks
  │   └── README.md          # Specific documentation for map benchmarks
  ├── README.md              # Main repository documentation
  └── LICENSE                # License file
```

## Contributing

Contributions are welcome! If you'd like to add a new benchmark or improve an existing one:

1. Fork the repository
2. Create a new branch for your changes
3. Add your benchmark in its own package or directory with a descriptive name
4. Include a README.md within the benchmark directory explaining its purpose
5. Submit a pull request

Please ensure your benchmark includes:

- Clear documentation about what's being tested and why
- Proper benchmark functions with descriptive names
- Comparable implementations when testing alternatives
- Cleanup of any resources used during benchmarking

## License

[Apache 2.0 License](LICENSE)