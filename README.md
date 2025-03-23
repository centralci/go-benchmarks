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

The compression benchmarks compare the performance of different compression libraries available in Go across various
scenarios to help you make informed decisions about which implementation to use for specific workloads.

### ZSTD Compression Benchmarks

Compares different ZSTD implementations in Go:

- [Klauspost's native Go implementation](https://github.com/klauspost/compress) (github.com/klauspost/compress/zstd)
- [DataDog's CGO wrapper](https://github.com/DataDog/zstd) (github.com/DataDog/zstd)

The benchmarks evaluate:

- **Compression Speed**: At multiple compression levels (1, 3, and 7)
- **Decompression Speed**: For data compressed at different levels
- **Compression Ratio**: Measuring the effectiveness of compression
- **Data Types**: Testing different content types (text, binary, random) to simulate real-world data
- **Data Sizes**: Small (1MB), Medium (10MB), and Large (100MB) payloads

### GZIP Compression Benchmarks

Compares different GZIP implementations in Go:

- [Klauspost's optimized implementation](https://github.com/klauspost/compress) (github.com/klauspost/compress/gzip)
- Go standard library (compress/gzip)

The benchmarks evaluate:

- **Compression Speed**: At standard gzip levels (1, 3, and 9)
- **Decompression Speed**: For data compressed at different levels
- **Compression Ratio**: Measuring the effectiveness of compression
- **Data Types**: Testing different content types (text, binary, random)
- **Data Sizes**: Small (1MB), Medium (10MB), and Large (100MB) payloads

Running specific compression benchmarks:

```bash
# Run all compression benchmarks
go test ./compression -bench=.

# Run only ZSTD benchmarks
go test ./compression -bench=Zstd

# Run only GZIP benchmarks
go test ./compression -bench=Gzip

# Run only compression performance tests (excluding decompression)
go test ./compression -bench=Compress

# Run only decompression performance tests
go test ./compression -bench=Decompress

# Run only compression ratio tests
go test ./compression -bench=Ratio

# Test specific data sizes
go test ./compression -bench=1MB    # Small size (1MB)
go test ./compression -bench=10MB   # Medium size (10MB)
go test ./compression -bench=100MB  # Large size (100MB)

# Test specific data types
go test ./compression -bench=Text
go test ./compression -bench=Binary
go test ./compression -bench=Random

# Run tests by compression level
go test ./compression -bench=Level1
go test ./compression -bench=Level9
```

### More Benchmarks Coming Soon

This repository will be expanded with additional benchmarks for:

- Concurrency patterns
- Memory allocation strategies
- JSON encoding/decoding
- HTTP server implementations
- And more...

## Requirements

- Go 1.24 or later
- For DataDog's ZSTD implementation: C compiler (CGO required)

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

## Using Benchstat

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

### Compression Benchmark Insights

For compression benchmarks, you should consider:

1. **Throughput (MB/s)**: Higher is better, indicates how fast the library can process data
2. **Compression Ratio**: Higher is better, indicates how effectively the library reduces data size
3. **Memory Usage (B/op and allocs/op)**: Lower is better, indicates resource efficiency

Different workloads may prioritize different metrics:

- **Speed-critical applications**: Focus on throughput (MB/s)
- **Network or storage-constrained applications**: Focus on compression ratio
- **Memory-constrained environments**: Focus on B/op and allocs/op

Also consider:

- CGO vs pure Go implementations (deployment considerations)
- Different data types affect compression performance differently
- Higher compression levels trade speed for better ratios

## Repository Structure

The repository is organized by benchmark categories:

```
go-benchmarks/
  ├── maps/                      # Map implementation benchmarks
  │   ├── concurrent_test.go     # Concurrent map benchmarks
  │   └── README.md              # Specific documentation for map benchmarks
  ├── compression/               # Compression benchmarks
  │   ├── benchmark_utils.go     # Shared utilities for compression tests
  │   ├── zstd_test.go           # ZSTD compression benchmarks
  │   ├── gzip_test.go           # GZIP compression benchmarks 
  │   └── README.md              # Documentation for compression benchmarks
  ├── README.md                  # Main repository documentation
  └── LICENSE                    # License file
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