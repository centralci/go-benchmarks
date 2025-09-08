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

### Go 1.25 Results

```
goos: darwin
goarch: arm64
pkg: go-benchmarks/maps
cpu: Apple M1 Pro
BenchmarkRWMutexMapReadHeavy        	82041930	        14.99 ns/op
BenchmarkRWMutexMapReadHeavy-2      	52191877	        23.66 ns/op
BenchmarkRWMutexMapReadHeavy-4      	39310750	        30.40 ns/op
BenchmarkRWMutexMapReadHeavy-8      	35395923	        32.37 ns/op
BenchmarkRWMutexMapReadHeavy-16     	34912178	        34.06 ns/op
BenchmarkSyncMapReadHeavy           	54095376	        22.03 ns/op
BenchmarkSyncMapReadHeavy-2         	78385047	        13.38 ns/op
BenchmarkSyncMapReadHeavy-4         	129736099	         9.138 ns/op
BenchmarkSyncMapReadHeavy-8         	100000000	        12.03 ns/op
BenchmarkSyncMapReadHeavy-16        	100000000	        10.31 ns/op
BenchmarkRWMutexMapWriteHeavy       	77759660	        15.21 ns/op
BenchmarkRWMutexMapWriteHeavy-2     	30364306	        39.91 ns/op
BenchmarkRWMutexMapWriteHeavy-4     	18328888	        61.03 ns/op
BenchmarkRWMutexMapWriteHeavy-8     	14709321	        80.12 ns/op
BenchmarkRWMutexMapWriteHeavy-16    	13350471	        83.20 ns/op
BenchmarkSyncMapWriteHeavy          	28567545	        42.84 ns/op
BenchmarkSyncMapWriteHeavy-2        	42715202	        33.16 ns/op
BenchmarkSyncMapWriteHeavy-4        	46165862	        23.99 ns/op
BenchmarkSyncMapWriteHeavy-8        	46027197	        27.70 ns/op
BenchmarkSyncMapWriteHeavy-16       	45115820	        26.67 ns/op
BenchmarkRWMutexMapMixedOps         	75721124	        15.86 ns/op
BenchmarkRWMutexMapMixedOps-2       	35453616	        36.48 ns/op
BenchmarkRWMutexMapMixedOps-4       	25388496	        49.68 ns/op
BenchmarkRWMutexMapMixedOps-8       	21837636	        48.34 ns/op
BenchmarkRWMutexMapMixedOps-16      	25616006	        46.82 ns/op
BenchmarkSyncMapMixedOps            	39693368	        29.13 ns/op
BenchmarkSyncMapMixedOps-2          	66883729	        18.31 ns/op
BenchmarkSyncMapMixedOps-4          	89313126	        13.95 ns/op
BenchmarkSyncMapMixedOps-8          	83577343	        14.16 ns/op
BenchmarkSyncMapMixedOps-16         	85875049	        14.33 ns/op
```

## Analysis

### Single-Goroutine Performance

With a single goroutine (no contention):

- **Read-Heavy**: RWMutexMap (14.99 ns/op) outperforms sync.Map (22.03 ns/op) by ~32%
- **Write-Heavy**: RWMutexMap (15.21 ns/op) outperforms sync.Map (42.84 ns/op) by ~65%
- **Mixed Operations**: RWMutexMap (15.86 ns/op) outperforms sync.Map (29.13 ns/op) by ~46%

This shows that for single-threaded access with no contention, the simpler RWMutexMap implementation has less overhead.

### Multi-Goroutine Scaling

As concurrency increases:

1. **RWMutexMap**:
    - Performance degrades significantly with increased goroutine count
    - At 16 goroutines, Read-Heavy: 34.06 ns/op (2.3x slower than single goroutine)
    - At 16 goroutines, Write-Heavy: 83.20 ns/op (5.5x slower than single goroutine)
    - At 16 goroutines, Mixed Ops: 46.82 ns/op (3x slower than single goroutine)

2. **sync.Map**:
    - Performance improves with increased concurrency for most workloads
    - At 4 goroutines, Read-Heavy: 9.138 ns/op (2.4x faster than single goroutine)
    - At 4 goroutines, Write-Heavy: 23.99 ns/op (1.8x faster than single goroutine)
    - At 4 goroutines, Mixed Ops: 13.95 ns/op (2.1x faster than single goroutine)
    - Performance remains relatively stable from 4-16 goroutines

### Workload-Specific Performance

1. **Read-Heavy (90% reads)**:
    - sync.Map scales extremely well, reaching 9.138 ns/op with 4 goroutines
    - sync.Map is up to 3.3x faster than RWMutexMap at high concurrency (16 goroutines)
    - Best sync.Map performance at 4 goroutines, slight degradation at higher concurrency

2. **Write-Heavy (50% writes)**:
    - Both implementations slow down with heavy write contention
    - sync.Map performs significantly better at scale (3.1x faster at 16 goroutines)
    - sync.Map shows best performance at 4 goroutines (23.99 ns/op)

3. **Mixed Operations**:
    - sync.Map handles mixed loads efficiently at scale
    - At 4+ goroutines, sync.Map is 3.3-3.5x faster than RWMutexMap
    - sync.Map performance plateaus around 13-14 ns/op from 4-16 goroutines

## Conclusions

1. **For single-threaded use cases**: RWMutexMap is simpler and faster by 32-65%

2. **For concurrent access patterns**: sync.Map is the clear winner, especially as concurrency increases:
    - Performs better with 2+ goroutines in all workloads
    - Excels at read-heavy workloads (common in many applications)
    - Handles contention much more gracefully than mutex-based approaches
    - Shows consistent 3x+ performance advantage at moderate to high concurrency

3. **Scaling characteristics**:
    - RWMutexMap performance degrades predictably as contention increases
    - sync.Map actually improves in performance up to 4 goroutines before stabilizing
    - The crossover point where sync.Map becomes faster is at 2 goroutines for all workloads

4. **Optimal concurrency levels**:
    - sync.Map performs best at 4 goroutines for most workloads
    - Performance remains relatively stable from 4-16 goroutines
    - This suggests sync.Map has excellent internal optimization for moderate concurrency

## Recommendations

1. **Use the standard map with a mutex when**:
    - Access is primarily single-threaded
    - The performance profile is predictable
    - Simplicity is valued over concurrent scaling
    - You have complete control over access patterns

2. **Use sync.Map when**:
    - There's any concurrent access from 2+ goroutines
    - Read operations dominate the workload
    - You need consistent performance as load increases
    - The key set is relatively stable (keys added once, read many times)
    - You want automatic optimization for concurrent access patterns

3. **Consider other options when**:
    - You need ordered iteration (neither implementation provides this)
    - Specialized needs like expiring entries, custom eviction, etc.
    - You need deterministic performance guarantees

## Running the Benchmarks

```bash
# Run all map benchmarks
go test -bench=.

# Compare scaling across different CPU counts
go test -bench=. -cpu=1,2,4,8,16

# Include memory allocation statistics
go test -bench=. -benchmem

# Run specific workload pattern
go test -bench=ReadHeavy

# Run with specific concurrency level
go test -bench=. -cpu=4

# Save results to file
go test -bench=. -cpu=1,2,4,8,16 | tee results.txt
```