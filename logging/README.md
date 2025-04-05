# Go Logging Library Benchmarks

This package contains benchmarks comparing the performance of different logging libraries in Go:

1. **log/slog**: Go's built-in structured logging library from the standard library
2. **lager**: The logging library from Cloudfoundry (code.cloudfoundry.org/lager/v3)

## Benchmark Results

The benchmarks were run on an Apple M1 Pro (ARM64) with a variety of logging scenarios:

- **Simple Logging**: Basic string messages
- **Structured Logging**: Adding key-value fields
- **JSON Output**: Formatted as JSON
- **Context Handling**: Using context or session data
- **Concurrency**: Performance under parallel execution
- **Variable Payload Size**: Small, medium, and large log messages
- **Filtering**: Performance when logs are filtered out
- **Complex Data**: Handling nested structures
- **Multiple Destinations**: Writing to multiple outputs
- **Error Handling**: Logging errors with and without stack traces

### Results Summary

```
goos: darwin
goarch: arm64
pkg: go-benchmarks/logging
cpu: Apple M1 Pro
BenchmarkSlogSimpleMessage-10         	 2280758	       501.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkLagerSimpleMessage-10        	 1863362	       651.9 ns/op	     488 B/op	      10 allocs/op
BenchmarkSlogWithFields-10            	 1610775	       737.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkLagerWithFields-10           	  870069	      1308 ns/op	    1417 B/op	      21 allocs/op
BenchmarkSlogJSON-10                  	 1832324	       646.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkLagerJSON-10                 	  832414	      1312 ns/op	    1417 B/op	      21 allocs/op
BenchmarkSlogWithContext-10           	 1903300	       629.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkLagerWithContext-10          	  726060	      1494 ns/op	    1561 B/op	      24 allocs/op
BenchmarkSlogDeepContext-10           	 1816792	       653.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkLagerDeepContext-10          	  494845	      2208 ns/op	    2418 B/op	      33 allocs/op
BenchmarkSlogParallel-10              	 4795230	       248.3 ns/op	       8 B/op	       0 allocs/op
BenchmarkLagerParallel-10             	 2119204	       596.2 ns/op	    1235 B/op	      18 allocs/op
BenchmarkSlogHighConcurrency-10       	 5239953	       228.8 ns/op	      16 B/op	       1 allocs/op
BenchmarkLagerHighConcurrency-10      	 1610395	       733.8 ns/op	    1355 B/op	      21 allocs/op
BenchmarkSlogSmallPayload-10          	 2581264	       452.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkLagerSmallPayload-10         	 1889598	       635.1 ns/op	     472 B/op	      10 allocs/op
BenchmarkSlogMediumPayload-10         	  754803	      1560 ns/op	       0 B/op	       0 allocs/op
BenchmarkLagerMediumPayload-10        	 1314696	       907.8 ns/op	     985 B/op	      10 allocs/op
BenchmarkSlogLargePayload-10          	  108510	     11037 ns/op	       0 B/op	       0 allocs/op
BenchmarkLagerLargePayload-10         	  359397	      3164 ns/op	    5670 B/op	      10 allocs/op
BenchmarkSlogFilteredOut-10           	90318374	        12.63 ns/op	       8 B/op	       1 allocs/op
BenchmarkLagerFilteredOut-10          	 2450464	       482.7 ns/op	     472 B/op	       8 allocs/op
BenchmarkSlogComplexStructures-10     	  124852	      9192 ns/op	    3491 B/op	      66 allocs/op
BenchmarkLagerComplexStructures-10    	  189454	      6111 ns/op	    5462 B/op	      99 allocs/op
BenchmarkSlogMultiDestination-10      	 1498198	       795.5 ns/op	     724 B/op	       0 allocs/op
BenchmarkLagerMultiDestination-10     	  474717	      2209 ns/op	    2992 B/op	      32 allocs/op
BenchmarkSlogError-10                 	 1650428	       720.7 ns/op	       3 B/op	       1 allocs/op
BenchmarkLagerError-10                	  919860	      1225 ns/op	    1321 B/op	      20 allocs/op
BenchmarkSlogErrorWithStack-10        	  161526	      6841 ns/op	    4824 B/op	       4 allocs/op
BenchmarkLagerErrorWithStack-10       	  243735	      4464 ns/op	    6786 B/op	      21 allocs/op
BenchmarkSlogCustomTimeFormat-10      	 1345260	       892.9 ns/op	      72 B/op	       4 allocs/op
BenchmarkLagerCustomTimeFormat-10     	  813915	      1321 ns/op	    1081 B/op	      19 allocs/op
BenchmarkSlogWithRedaction-10         	 1549944	       777.6 ns/op	       8 B/op	       0 allocs/op
BenchmarkLagerWithRedaction-10        	  650437	      1679 ns/op	    1671 B/op	      24 allocs/op
BenchmarkSlogInfoWithFile-10          	  516919	      2195 ns/op	       8 B/op	       0 allocs/op
```

## Analysis

### Basic Logging Performance

- **Simple Messages**: `slog` (501.3 ns/op) outperforms `lager` (651.9 ns/op) by ~23%
- **Structured Fields**: `slog` (737.9 ns/op) outperforms `lager` (1308 ns/op) by ~44%
- **JSON Output**: `slog` (646.6 ns/op) outperforms `lager` (1312 ns/op) by ~51%

In all basic logging scenarios, `slog` is significantly faster and makes virtually no memory allocations compared to `lager`.

### Memory Efficiency

The memory allocation differences are striking:

- `slog` often requires zero allocations where `lager` needs 10-30 allocations
- Simple message: 0 B/op vs 488 B/op
- Structured fields: 0 B/op vs 1417 B/op
- With context: 0 B/op vs 1561 B/op

This translates to much lower garbage collection pressure with `slog`.

### Context and Nesting

As complexity increases with deeper context nesting:

1. **slog**:
    - Basic context: 629.0 ns/op
    - Deep context: 653.3 ns/op (only 4% slower)

2. **lager**:
    - Basic context: 1494 ns/op
    - Deep context: 2208 ns/op (48% slower)

This shows that `slog` scales much better with context complexity.

### Concurrency Performance

`slog` shows exceptional performance under concurrent load:

- **Parallel execution**: `slog` (248.3 ns/op) vs `lager` (596.2 ns/op) - 2.4x faster
- **High concurrency**: `slog` (228.8 ns/op) vs `lager` (733.8 ns/op) - 3.2x faster

Memory efficiency in concurrent scenarios is even more pronounced:
- `slog` with high concurrency: 16 B/op with 1 allocation
- `lager` with high concurrency: 1355 B/op with 21 allocations

### Payload Size Impact

Interestingly, as payload size increases, the performance difference shifts:

- **Small payloads**: `slog` is 40% faster (452.0 ns/op vs 635.1 ns/op)
- **Medium payloads**: `lager` is 42% faster (907.8 ns/op vs 1560 ns/op)
- **Large payloads**: `lager` is significantly faster (3164 ns/op vs 11037 ns/op)

This suggests that `lager` has more efficient large string handling, while `slog` excels with smaller payloads.

### Filtering Performance

When logs are filtered out by log level:

- `slog`: 12.63 ns/op
- `lager`: 482.7 ns/op

This is a 38x performance difference, showing that `slog` has extremely efficient level filtering.

### Complex Operations

In some complex scenarios, `lager` shows better performance:

- **Complex structures**: `lager` (6111 ns/op) is 33% faster than `slog` (9192 ns/op)
- **Error with stack trace**: `lager` (4464 ns/op) is 35% faster than `slog` (6841 ns/op)

However, for multiple destination logging, `slog` (795.5 ns/op) significantly outperforms `lager` (2209 ns/op).

## Conclusions

1. **Overall Performance**: `slog` generally outperforms `lager` by a significant margin in most scenarios (1.5-3x faster).

2. **Memory Efficiency**: `slog` is remarkably memory efficient, often requiring no allocations where `lager` requires many.

3. **Scaling Characteristics**:
    - `slog` scales much better with increased context complexity
    - `slog` excels in concurrent environments
    - `slog` has extremely efficient log level filtering

4. **Areas where `lager` performs better**:
    - Large payload handling
    - Certain complex data structures
    - Error logging with stack traces

## Recommendations

1. **Use `log/slog` when**:
    - Memory efficiency is important
    - High concurrency is expected
    - Log filtering is common (production environments)
    - Most logs contain small to medium-sized messages
    - Context or structured logging is heavily used

2. **Consider `lager` when**:
    - Your application primarily logs very large messages
    - You work with extremely complex nested data structures
    - You're already using it and don't have performance issues
    - You need its specific features (particularly with test helpers)

3. **Migration considerations**:
    - Moving from `lager` to `slog` should yield significant performance benefits in most cases
    - Watch for large payload performance when migrating

## Running the Benchmarks

```bash
# Run all logging benchmarks
go test -bench=.

# Run with memory allocation statistics
go test -bench=. -benchmem

# Run a specific benchmark
go test -bench=BenchmarkSlogJSON

# Run shorter benchmarks (skips file I/O and sustained tests)
go test -bench=. -short
```