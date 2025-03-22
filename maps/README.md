# Concurrent Maps Benchmarks

This package contains benchmarks comparing the performance of different concurrent map implementations in Go:

1. **RWMutexMap**: A custom map implementation protected by a read-write mutex (`sync.RWMutex`)
2. **sync.Map**: Go's built-in concurrent map from the standard library

## Benchmark Results

The benchmarks were run on an Apple M1 Pro (ARM64) with the following workloads:

- **Read-Heavy**: 90% reads, 10% writes
- **Write-Heavy**: 50% reads, 50% writes
- **Mixed Operations**: 70% reads, 20% writes, 10% deletes

Each benchmark was run with varying levels of concurrency (1, 2, 4, 8, and 16 goroutines).

### Results Summary

```
goos: darwin
goarch: arm64
pkg: go-benchmarks/maps
cpu: Apple M1 Pro
BenchmarkRWMutexMapReadHeavy        	84669350	        14.01 ns/op
BenchmarkRWMutexMapReadHeavy-2      	51196630	        22.65 ns/op
BenchmarkRWMutexMapReadHeavy-4      	36205647	        33.11 ns/op
BenchmarkRWMutexMapReadHeavy-8      	34268498	        35.05 ns/op
BenchmarkRWMutexMapReadHeavy-16     	33892998	        36.37 ns/op
BenchmarkSyncMapReadHeavy           	57972414	        21.49 ns/op
BenchmarkSyncMapReadHeavy-2         	81328593	        13.59 ns/op
BenchmarkSyncMapReadHeavy-4         	147637832	         7.966 ns/op
BenchmarkSyncMapReadHeavy-8         	100000000	        10.22 ns/op
BenchmarkSyncMapReadHeavy-16        	100000000	        10.79 ns/op
BenchmarkRWMutexMapWriteHeavy       	79413225	        14.68 ns/op
BenchmarkRWMutexMapWriteHeavy-2     	28473128	        41.14 ns/op
BenchmarkRWMutexMapWriteHeavy-4     	18950098	        62.62 ns/op
BenchmarkRWMutexMapWriteHeavy-8     	15257396	        80.64 ns/op
BenchmarkRWMutexMapWriteHeavy-16    	13607871	        88.31 ns/op
BenchmarkSyncMapWriteHeavy          	28798761	        40.63 ns/op
BenchmarkSyncMapWriteHeavy-2        	42907136	        26.37 ns/op
BenchmarkSyncMapWriteHeavy-4        	69762429	        17.56 ns/op
BenchmarkSyncMapWriteHeavy-8        	46850445	        25.11 ns/op
BenchmarkSyncMapWriteHeavy-16       	42911804	        25.74 ns/op
BenchmarkRWMutexMapMixedOps         	77070255	        15.12 ns/op
BenchmarkRWMutexMapMixedOps-2       	37058673	        31.21 ns/op
BenchmarkRWMutexMapMixedOps-4       	25773841	        46.89 ns/op
BenchmarkRWMutexMapMixedOps-8       	24309376	        48.60 ns/op
BenchmarkRWMutexMapMixedOps-16      	23761552	        49.18 ns/op
BenchmarkSyncMapMixedOps            	42591766	        25.57 ns/op
BenchmarkSyncMapMixedOps-2          	67241329	        16.67 ns/op
BenchmarkSyncMapMixedOps-4          	108342964	        10.31 ns/op
BenchmarkSyncMapMixedOps-8          	79787233	        14.54 ns/op
BenchmarkSyncMapMixedOps-16         	83495446	        13.93 ns/op
```

## Analysis

### Single-Goroutine Performance

With a single goroutine (no contention):

- **Read-Heavy**: RWMutexMap (14.01 ns/op) outperforms sync.Map (21.49 ns/op) by ~35%
- **Write-Heavy**: RWMutexMap (14.68 ns/op) outperforms sync.Map (40.63 ns/op) by ~64%
- **Mixed Operations**: RWMutexMap (15.12 ns/op) outperforms sync.Map (25.57 ns/op) by ~41%

This shows that for single-threaded access with no contention, the simpler RWMutexMap implementation has less overhead.

### Multi-Goroutine Scaling

As concurrency increases:

1. **RWMutexMap**:
    - Performance degrades significantly with increased goroutine count
    - At 16 goroutines, Read-Heavy: 36.37 ns/op (2.6x slower than single goroutine)
    - At 16 goroutines, Write-Heavy: 88.31 ns/op (6x slower than single goroutine)
    - At 16 goroutines, Mixed Ops: 49.18 ns/op (3.3x slower than single goroutine)

2. **sync.Map**:
    - Performance improves with increased concurrency for most workloads
    - At 4 goroutines, Read-Heavy: 7.966 ns/op (2.7x faster than single goroutine)
    - At 4 goroutines, Write-Heavy: 17.56 ns/op (2.3x faster than single goroutine)
    - At 4 goroutines, Mixed Ops: 10.31 ns/op (2.5x faster than single goroutine)

### Workload-Specific Performance

1. **Read-Heavy (90% reads)**:
    - sync.Map scales extremely well, reaching 7.966 ns/op with 4 goroutines
    - sync.Map is up to 4.2x faster than RWMutexMap at high concurrency

2. **Write-Heavy (50% writes)**:
    - Both implementations slow down with heavy write contention
    - sync.Map performs significantly better at scale (3.4x faster at 16 goroutines)

3. **Mixed Operations**:
    - sync.Map handles mixed loads efficiently at scale
    - At 4+ goroutines, sync.Map is 3.5-4.5x faster than RWMutexMap

## Conclusions

1. **For single-threaded use cases**: RWMutexMap is simpler and faster by 35-64%

2. **For concurrent access patterns**: sync.Map is the clear winner, especially as concurrency increases:
    - Performs better with 2+ goroutines in all workloads
    - Excels at read-heavy workloads (common in many applications)
    - Handles contention much more gracefully than mutex-based approaches

3. **Scaling characteristics**:
    - RWMutexMap performance degrades predictably as contention increases
    - sync.Map actually improves in performance up to 4 goroutines before slightly degrading

## Recommendations

1. **Use the standard map with a mutex when**:
    - Access is primarily single-threaded
    - The performance profile is predictable
    - Simplicity is valued over concurrent scaling

2. **Use sync.Map when**:
    - There's significant concurrent access from multiple goroutines
    - Read operations dominate the workload
    - You need consistent performance as load increases
    - The key set is relatively stable (keys added once, read many times)

3. **Consider other options when**:
    - You need ordered iteration (neither implementation provides this)
    - Specialized needs like expiring entries, custom eviction, etc.

## Running the Benchmarks

```bash
# Run all map benchmarks
go test -bench=.

# Compare scaling across different CPU counts
go test -bench=. -cpu=1,2,4,8,16

# Include memory allocation statistics
go test -bench=. -benchmem
```