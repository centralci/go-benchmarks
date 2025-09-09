# Benchmark Analysis: Sonic vs Standard JSON vs JSON v2 in Go

This analysis compares the performance of ByteDance's Sonic JSON parser with Go's standard library encoding/json package
and the experimental JSON v2 implementation across different data sizes and operations.

## Test Environment

- **OS**: Darwin (macOS)
- **Architecture**: arm64
- **CPU**: Apple M1 Pro
- **Go Version**: 1.25
- **Sonic Version**: 1.14.1
- **Data Sizes**: 1MB (250 rows), 10MB (2,500 rows), 100MB (25,000 rows)
- **Operations**:
    - Marshal (Go → JSON bytes)
    - Unmarshal (JSON → Go)
    - ToString (Go → JSON string)
- **Libraries**:
    - Go standard library (encoding/json)
    - Go experimental JSON v2 (GOEXPERIMENT=jsonv2)
    - ByteDance Sonic (standard configuration)
    - ByteDance Sonic (Fastest configuration)

## Results Summary

### Standard JSON Library (encoding/json)

| Operation | Size  | Ops/sec | Time (ns/op) | Throughput (MB/s) | Memory (B/op) | Allocations |
|-----------|-------|---------|--------------|-------------------|---------------|-------------|
| Marshal   | 1MB   | 1,942   | 616,655      | 873.72            | 541,788       | 2           |
| Marshal   | 10MB  | 193     | 6,183,486    | 872.50            | 5,572,563     | 2           |
| Marshal   | 100MB | 19      | 61,751,412   | 873.41            | 61,001,653    | 4           |
| ToString  | 1MB   | 1,870   | 629,225      | 856.27            | 1,090,406     | 3           |
| ToString  | 10MB  | 163     | 6,674,942    | 808.26            | 11,827,045    | 5           |
| ToString  | 100MB | 14      | 73,166,387   | 737.15            | 174,991,918   | 23          |
| Unmarshal | 1MB   | 379     | 3,122,794    | 172.53            | 748,576       | 7,010       |
| Unmarshal | 10MB  | 38      | 31,535,310   | 171.08            | 8,234,969     | 69,958      |
| Unmarshal | 100MB | 4       | 316,455,167  | 170.43            | 105,325,484   | 699,505     |

### Experimental JSON v2 (GOEXPERIMENT=jsonv2)

| Operation | Size  | Ops/sec | Time (ns/op) | Throughput (MB/s) | Memory (B/op) | Allocations |
|-----------|-------|---------|--------------|-------------------|---------------|-------------|
| Marshal   | 1MB   | 1,881   | 647,621      | 833.57            | 542,074       | 3           |
| Marshal   | 10MB  | 183     | 6,477,885    | 832.38            | 5,398,637     | 3           |
| Marshal   | 100MB | 16      | 65,046,706   | 829.46            | 73,625,289    | 6           |
| ToString  | 1MB   | 1,814   | 674,643      | 800.18            | 1,093,988     | 4           |
| ToString  | 10MB  | 168     | 7,201,848    | 748.70            | 12,848,307    | 7           |
| ToString  | 100MB | 13      | 80,582,388   | 669.55            | 374,146,762   | 51          |
| Unmarshal | 1MB   | 675     | 1,538,335    | 350.92            | 744,818       | 6,548       |
| Unmarshal | 10MB  | 78      | 15,746,964   | 342.42            | 8,194,827     | 65,320      |
| Unmarshal | 100MB | 7       | 159,448,661  | 338.38            | 104,961,381   | 653,396     |

### Sonic Standard

| Operation | Size  | Ops/sec | Time (ns/op) | Throughput (MB/s) | Memory (B/op) | Allocations |
|-----------|-------|---------|--------------|-------------------|---------------|-------------|
| Marshal   | 1MB   | 2,628   | 457,356      | 1,180.35          | 541,042       | 3           |
| Marshal   | 10MB  | 207     | 5,333,467    | 1,010.98          | 18,681,198    | 23          |
| Marshal   | 100MB | 21      | 52,044,865   | 1,036.68          | 184,980,011   | 32          |
| ToString  | 1MB   | 2,628   | 469,392      | 1,150.08          | 541,879       | 3           |
| ToString  | 10MB  | 222     | 5,292,330    | 1,018.84          | 18,680,538    | 23          |
| ToString  | 100MB | 22      | 52,408,968   | 1,029.48          | 184,978,892   | 31          |
| Unmarshal | 1MB   | 2,358   | 500,648      | 1,078.28          | 672,278       | 508         |
| Unmarshal | 10MB  | 223     | 5,488,916    | 982.35            | 33,404,348    | 5,032       |
| Unmarshal | 100MB | 22      | 49,587,256   | 1,088.06          | 481,182,032   | 50,238      |

### Sonic Fastest

| Operation | Size  | Ops/sec | Time (ns/op) | Throughput (MB/s) | Memory (B/op) | Allocations |
|-----------|-------|---------|--------------|-------------------|---------------|-------------|
| Marshal   | 1MB   | 2,638   | 462,730      | 1,166.64          | 540,829       | 3           |
| Marshal   | 10MB  | 226     | 5,166,940    | 1,043.56          | 18,681,146    | 23          |
| Marshal   | 100MB | 21      | 53,160,296   | 1,014.93          | 184,980,078   | 32          |
| ToString  | 1MB   | 2,485   | 521,106      | 1,035.95          | 542,885       | 3           |
| ToString  | 10MB  | 228     | 5,297,964    | 1,017.76          | 18,681,761    | 23          |
| ToString  | 100MB | 22      | 52,228,426   | 1,033.03          | 184,972,561   | 31          |
| Unmarshal | 1MB   | 2,392   | 500,321      | 1,078.98          | 667,791       | 508         |
| Unmarshal | 10MB  | 218     | 5,327,794    | 1,012.06          | 33,238,565    | 5,032       |
| Unmarshal | 100MB | 22      | 49,629,062   | 1,087.14          | 483,730,181   | 50,238      |

## Performance Comparisons

### Marshal Performance (Relative to Standard JSON)

| Size  | JSON v2 | Sonic  | Sonic Fastest |
|-------|---------|--------|---------------|
| 1MB   | -4.8%   | +34.9% | +33.3%        |
| 10MB  | -4.5%   | +15.9% | +19.7%        |
| 100MB | -5.1%   | +18.7% | +16.2%        |

### ToString Performance (Relative to Standard JSON)

| Size  | JSON v2 | Sonic  | Sonic Fastest |
|-------|---------|--------|---------------|
| 1MB   | -6.7%   | +34.0% | +20.7%        |
| 10MB  | -7.4%   | +26.1% | +26.0%        |
| 100MB | -9.2%   | +39.5% | +40.0%        |

### Unmarshal Performance (Relative to Standard JSON)

| Size  | JSON v2 | Sonic   | Sonic Fastest |
|-------|---------|---------|---------------|
| 1MB   | +103.0% | +523.7% | +524.0%       |
| 10MB  | +100.2% | +474.7% | +492.1%       |
| 100MB | +98.5%  | +538.3% | +537.8%       |

## Key Findings

### 1. JSON v2 String Output Performance

**Surprising Result**: JSON v2 performs **worse** than standard JSON for string output:

- **1MB**: 6.7% slower
- **10MB**: 7.4% slower
- **100MB**: 9.2% slower

**Memory Impact**: JSON v2 uses significantly more memory for string output:

- **100MB**: 374MB vs 175MB (2.14x more than standard JSON)
- More allocations: 51 vs 23 for 100MB

This suggests JSON v2's experimental implementation hasn't optimized the string conversion path yet.

### 2. Sonic's String Output Dominance

**Performance Gains over Standard JSON**:

- **1MB**: 34-34% faster
- **10MB**: 26% faster
- **100MB**: 39-40% faster

**Performance Gains over JSON v2**:

- **1MB**: 43-44% faster
- **10MB**: 36% faster
- **100MB**: 53-54% faster

Sonic's `MarshalToString()` consistently outperforms both standard and v2 implementations.

### 3. Marshal Performance Rankings

**Winners by Size**:

- **1MB**: Sonic Standard (1,180 MB/s)
- **10MB**: Sonic Fastest (1,044 MB/s)
- **100MB**: Sonic Standard (1,037 MB/s)

**JSON v2 Marshal**: Consistently 4-5% slower than standard JSON, suggesting regression in the experimental version.

### 4. Unmarshal Performance Analysis

**JSON v2 Improvements**:

- Delivers on promise: ~2x faster than standard JSON
- 350 MB/s vs 172 MB/s throughput
- 6.5K allocations vs 7K (slight improvement)

**Sonic Dominance**:

- 5.2x faster than standard JSON
- 2.6x faster than JSON v2
- 1,078-1,088 MB/s throughput

### 5. Memory Usage Patterns

| Operation    | Library  | 1MB Memory | 10MB Memory   | 100MB Memory |
|--------------|----------|------------|---------------|--------------|
| **Marshal**  |          |            |               |              |
|              | Standard | 542KB      | 5.4MB         | 61MB         |
|              | JSON v2  | 542KB      | 5.4MB         | 74MB (+21%)  |
|              | Sonic    | 541KB      | 18.7MB (3.5x) | 185MB (3x)   |
| **ToString** |          |            |               |              |
|              | Standard | 1.1MB      | 11.8MB        | 175MB        |
|              | JSON v2  | 1.1MB      | 12.8MB        | 374MB (2.1x) |
|              | Sonic    | 542KB      | 18.7MB        | 185MB        |

## Analysis by Use Case

### 1. String Output for APIs

**Clear Winner**: Sonic

- 34-40% faster than standard JSON
- 43-54% faster than JSON v2
- Direct string generation without conversion overhead

**Avoid**: JSON v2 for string output (slower and uses more memory)

### 2. Conservative Upgrade Path

**Recommendation**: JSON v2 for **unmarshaling only**

- 2x faster unmarshaling as promised
- Marshal performance regression makes it unsuitable for write-heavy workloads
- String output performance is particularly poor

### 3. Maximum Performance

**Recommendation**: Sonic (Standard or Fastest)

- 5x faster unmarshaling
- 15-35% faster marshaling
- 26-40% faster string output

### 4. Memory-Constrained Systems

**Recommendation**: Standard JSON

- Lowest memory footprint
- JSON v2 surprisingly uses more memory in some cases
- Sonic uses 3-4x memory for large payloads

## Recommendations by Workload

### Read-Heavy Services (80%+ Unmarshal)

1. **Sonic**: 5x improvement dominates
2. **JSON v2**: Acceptable 2x improvement if avoiding dependencies
3. **Standard**: Only if memory is critical

### Write-Heavy Services (80%+ Marshal)

1. **Sonic**: 15-35% improvement
2. **Standard JSON**: Reliable, low memory
3. **Avoid JSON v2**: 5% regression

### API Services (String Output)

1. **Sonic MarshalToString()**: 34-40% faster
2. **Standard JSON**: Acceptable performance
3. **Avoid JSON v2**: 6-9% slower with 2x memory

### Balanced Workloads

1. **Sonic**: Best overall performance
2. **JSON v2**: Only if unmarshal matters more
3. **Standard**: For stability and memory efficiency

## Code Examples

### Standard JSON (Avoid for APIs)

```go
data, _ := json.Marshal(obj)
str := string(data) // Extra allocation and copy
// 1MB: 1.1MB memory used
```

### JSON v2 (Worse for strings)

```go
// With GOEXPERIMENT=jsonv2
data, _ := json.Marshal(obj)
str := string(data) // Even worse performance
// 100MB: 374MB memory used (!)
```

### Sonic (Optimal)

```go
str, _ := sonic.MarshalToString(obj) // Direct generation
// Consistent memory usage, 40% faster
```

## Conclusion (Sep 2025)

The experimental JSON v2 shows mixed results:

- ✅ **Delivers** on unmarshal performance (2x faster)
- ❌ **Regresses** on marshal performance (5% slower)
- ❌ **Significantly worse** for string output (9% slower, 2x memory)

**Sonic remains the clear performance leader** across all operations:

- 5x faster unmarshaling
- 15-35% faster marshaling
- 34-40% faster string output
- Consistent performance across all data sizes

**JSON v2's experimental status is evident** - while unmarshal improvements are solid, the marshal and string output
regressions suggest it's not ready for production use in write-heavy or API scenarios.

For production systems, the choice is between:

1. **Sonic** for maximum performance (if you can accept external dependencies)
2. **Standard JSON** for stability and memory efficiency
3. **JSON v2** only for read-heavy workloads where unmarshal performance is critical

The string output benchmarks particularly highlight Sonic's advantage for modern API services, where JSON responses are
common and the 40% performance improvement can significantly impact latency and throughput.