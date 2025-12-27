# Benchmark Analysis: Sonic vs Standard JSON vs JSON v2 in Go

This analysis compares the performance of ByteDance's Sonic JSON parser with Go's standard library encoding/json package
and the experimental JSON v2 implementation across different data sizes and operations.

## Test Environment

- **OS**: Darwin (macOS)
- **Architecture**: arm64
- **CPU**: Apple M1 Pro
- **Go Version**: 1.25.5
- **Sonic Version**: 1.14.2
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
| Marshal   | 1MB   | 1,742   | 670,502      | 804.69            | 540,702       | 2           |
| Marshal   | 10MB  | 176     | 6,751,089    | 798.19            | 5,485,836     | 2           |
| Marshal   | 100MB | 16      | 68,165,651   | 791.59            | 62,350,649    | 4           |
| ToString  | 1MB   | 1,706   | 690,641      | 781.23            | 1,085,084     | 3           |
| ToString  | 10MB  | 169     | 7,178,527    | 750.66            | 12,271,396    | 6           |
| ToString  | 100MB | 14      | 72,933,759   | 739.84            | 155,862,842   | 18          |
| Unmarshal | 1MB   | 384     | 3,074,064    | 175.52            | 749,400       | 7,010       |
| Unmarshal | 10MB  | 38      | 30,691,465   | 175.57            | 8,228,387     | 69,966      |
| Unmarshal | 100MB | 4       | 309,007,531  | 174.62            | 105,355,204   | 699,501     |

### Experimental JSON v2 (GOEXPERIMENT=jsonv2)

| Operation | Size  | Ops/sec | Time (ns/op) | Throughput (MB/s) | Memory (B/op) | Allocations |
|-----------|-------|---------|--------------|-------------------|---------------|-------------|
| Marshal   | 1MB   | 1,848   | 636,421      | 846.19            | 542,098       | 3           |
| Marshal   | 10MB  | 180     | 6,456,167    | 835.12            | 5,545,863     | 3           |
| Marshal   | 100MB | 16      | 64,463,430   | 836.85            | 73,617,097    | 6           |
| ToString  | 1MB   | 1,755   | 679,104      | 793.00            | 1,094,411     | 4           |
| ToString  | 10MB  | 172     | 6,884,947    | 783.11            | 13,108,331    | 7           |
| ToString  | 100MB | 14      | 76,833,440   | 702.12            | 287,693,708   | 36          |
| Unmarshal | 1MB   | 760     | 1,566,304    | 343.82            | 744,584       | 6,546       |
| Unmarshal | 10MB  | 79      | 15,712,219   | 343.15            | 8,195,323     | 65,328      |
| Unmarshal | 100MB | 7       | 157,554,792  | 342.40            | 104,949,756   | 653,557     |

### Sonic Standard

| Operation | Size  | Ops/sec | Time (ns/op) | Throughput (MB/s) | Memory (B/op) | Allocations |
|-----------|-------|---------|--------------|-------------------|---------------|-------------|
| Marshal   | 1MB   | 2,635   | 457,725      | 1,176.54          | 540,882       | 3           |
| Marshal   | 10MB  | 229     | 5,247,813    | 1,027.41          | 19,897,049    | 23          |
| Marshal   | 100MB | 21      | 54,648,510   | 987.15            | 240,259,552   | 33          |
| ToString  | 1MB   | 2,592   | 491,180      | 1,096.40          | 541,744       | 3           |
| ToString  | 10MB  | 226     | 5,292,139    | 1,018.81          | 19,897,678    | 23          |
| ToString  | 100MB | 21      | 52,799,875   | 1,021.71          | 240,260,227   | 34          |
| Unmarshal | 1MB   | 2,281   | 499,042      | 1,079.13          | 667,927       | 508         |
| Unmarshal | 10MB  | 224     | 5,288,388    | 1,019.53          | 33,193,374    | 5,029       |
| Unmarshal | 100MB | 22      | 50,617,670   | 1,065.76          | 483,672,264   | 50,233      |

### Sonic Fastest

| Operation | Size  | Ops/sec | Time (ns/op) | Throughput (MB/s) | Memory (B/op) | Allocations |
|-----------|-------|---------|--------------|-------------------|---------------|-------------|
| Marshal   | 1MB   | 2,628   | 457,412      | 1,177.34          | 541,518       | 3           |
| Marshal   | 10MB  | 228     | 5,236,938    | 1,029.54          | 19,897,662    | 23          |
| Marshal   | 100MB | 21      | 54,744,040   | 985.43            | 240,266,192   | 33          |
| ToString  | 1MB   | 2,596   | 465,585      | 1,156.68          | 542,493       | 3           |
| ToString  | 10MB  | 229     | 5,296,731    | 1,017.92          | 19,898,278    | 23          |
| ToString  | 100MB | 21      | 53,468,861   | 1,008.93          | 240,260,227   | 34          |
| Unmarshal | 1MB   | 2,392   | 492,846      | 1,092.70          | 673,062       | 508         |
| Unmarshal | 10MB  | 214     | 5,280,203    | 1,021.11          | 33,194,938    | 5,029       |
| Unmarshal | 100MB | 24      | 50,689,012   | 1,064.26          | 483,459,945   | 50,233      |

## Performance Comparisons

### Marshal Performance (Relative to Standard JSON)

| Size  | JSON v2 | Sonic  | Sonic Fastest |
|-------|---------|--------|---------------|
| 1MB   | +5.1%   | +46.5% | +46.6%        |
| 10MB  | +4.6%   | +28.7% | +28.9%        |
| 100MB | +5.7%   | +24.7% | +24.5%        |

### ToString Performance (Relative to Standard JSON)

| Size  | JSON v2 | Sonic  | Sonic Fastest |
|-------|---------|--------|---------------|
| 1MB   | +1.5%   | +40.4% | +48.3%        |
| 10MB  | +4.3%   | +35.6% | +35.5%        |
| 100MB | -5.1%   | +38.1% | +36.4%        |

### Unmarshal Performance (Relative to Standard JSON)

| Size  | JSON v2  | Sonic   | Sonic Fastest |
|-------|----------|---------|---------------|
| 1MB   | +96.2%   | +515.9% | +523.7%       |
| 10MB  | +95.3%   | +480.3% | +481.4%       |
| 100MB | +96.1%   | +510.4% | +509.8%       |

## Key Findings

### 1. JSON v2 Delivers on Unmarshal Promise

**Performance Gains over Standard JSON**:

- **1MB**: 96% faster (1.96x)
- **10MB**: 95% faster (1.95x)
- **100MB**: 96% faster (1.96x)

JSON v2 consistently delivers ~2x faster unmarshaling as promised, with slightly reduced allocations (6,546 vs 7,010 for 1MB).

### 2. JSON v2 Marshal: Modest Improvement

**Performance Gains over Standard JSON**:

- **1MB**: 5.1% faster
- **10MB**: 4.6% faster
- **100MB**: 5.7% faster

Marshal performance improved compared to previous experimental versions that showed regressions.

### 3. JSON v2 String Output: Mixed Results

**Performance vs Standard JSON**:

- **1MB**: 1.5% faster
- **10MB**: 4.3% faster
- **100MB**: 5.1% **slower** (702 MB/s vs 740 MB/s)

**Memory Impact at 100MB**: JSON v2 uses 288MB vs 156MB (1.8x more). The 100MB string output remains a weak point.

### 4. Sonic's Continued Dominance

**Sonic vs JSON v2**:

| Operation | 1MB     | 10MB    | 100MB   |
|-----------|---------|---------|---------|
| Marshal   | +39%    | +23%    | +18%    |
| ToString  | +38%    | +30%    | +45%    |
| Unmarshal | +214%   | +197%   | +211%   |

Sonic remains **3x faster** than JSON v2 for unmarshaling, even after JSON v2's 2x improvement.

### 5. Memory Usage Patterns

| Operation    | Library  | 1MB Memory | 10MB Memory   | 100MB Memory   |
|--------------|----------|------------|---------------|----------------|
| **Marshal**  |          |            |               |                |
|              | Standard | 541KB      | 5.5MB         | 62MB           |
|              | JSON v2  | 542KB      | 5.5MB         | 74MB (+18%)    |
|              | Sonic    | 541KB      | 19.9MB (3.6x) | 240MB (3.9x)   |
| **ToString** |          |            |               |                |
|              | Standard | 1.1MB      | 12.3MB        | 156MB          |
|              | JSON v2  | 1.1MB      | 13.1MB        | 288MB (1.8x)   |
|              | Sonic    | 542KB      | 19.9MB        | 240MB          |
| **Unmarshal**|          |            |               |                |
|              | Standard | 749KB      | 8.2MB         | 105MB          |
|              | JSON v2  | 745KB      | 8.2MB         | 105MB          |
|              | Sonic    | 668KB      | 33.2MB (4x)   | 484MB (4.6x)   |

## Analysis by Use Case

### 1. String Output for APIs

**Clear Winner**: Sonic

- 38-48% faster than standard JSON
- 30-45% faster than JSON v2
- Direct string generation without conversion overhead

**Avoid**: JSON v2 for large payloads (100MB: slower and uses 1.8x memory)

### 2. Conservative Upgrade Path

**Recommendation**: JSON v2

- 2x faster unmarshaling (as promised)
- 5% faster marshaling
- No external dependencies
- Drop-in replacement with `GOEXPERIMENT=jsonv2`

**Caveat**: Avoid for large string outputs (100MB+)

### 3. Maximum Performance

**Recommendation**: Sonic

- 5-6x faster unmarshaling (3x faster than JSON v2)
- 25-47% faster marshaling
- 35-48% faster string output

### 4. Memory-Constrained Systems

**Recommendation**: Standard JSON or JSON v2

- Lowest memory footprint
- JSON v2 comparable to standard for marshal/unmarshal
- Sonic uses 3-5x memory for large payloads

## Recommendations by Workload

### Read-Heavy Services (80%+ Unmarshal)

1. **Sonic**: 6x improvement over standard, 3x over JSON v2
2. **JSON v2**: Solid 2x improvement with no dependencies
3. **Standard**: Only if memory is critical

### Write-Heavy Services (80%+ Marshal)

1. **Sonic**: 25-47% improvement
2. **JSON v2**: 5% improvement, good conservative choice
3. **Standard JSON**: Reliable baseline

### API Services (String Output)

1. **Sonic MarshalString()**: 35-48% faster
2. **Standard JSON**: Acceptable performance
3. **JSON v2**: Avoid for large payloads

### Balanced Workloads

1. **Sonic**: Best overall performance
2. **JSON v2**: Good middle ground with 2x unmarshal boost
3. **Standard**: For stability and memory efficiency

## Code Examples

### Standard JSON

```go
data, _ := json.Marshal(obj)
str := string(data) // Extra allocation and copy
// 1MB: 1.1MB memory, 781 MB/s
```

### JSON v2 (Optimal for no-dependency upgrade)

```go
// Build with: GOEXPERIMENT=jsonv2 go build
data, _ := json.Marshal(obj)  // Same API, 2x faster unmarshal
str := string(data)
// 1MB unmarshal: 344 MB/s (vs 176 MB/s standard)
```

### Sonic (Optimal for performance)

```go
str, _ := sonic.MarshalString(obj) // Direct generation
// 1MB: 542KB memory, 1,096 MB/s (40% faster than standard)

var result []Data
sonic.Unmarshal(jsonBytes, &result)
// 1MB: 1,079 MB/s (6x faster than standard, 3x faster than JSON v2)
```

## Conclusion (Dec 2025)

**JSON v2** has matured significantly:

- ✅ **Delivers** on 2x faster unmarshal promise
- ✅ **Improved** marshal performance (5% faster)
- ⚠️ **Mixed** string output (slower at 100MB with higher memory)

**Sonic** remains the performance leader:

- ✅ **6x faster** unmarshaling vs standard (3x faster than JSON v2)
- ✅ **25-47% faster** marshaling
- ✅ **35-48% faster** string output
- ⚠️ **3-5x higher** memory usage for large payloads

### Decision Matrix

| Priority | Recommendation |
|----------|----------------|
| Maximum performance | Sonic |
| No external dependencies | JSON v2 |
| Memory efficiency | Standard JSON |
| Read-heavy workloads | Sonic > JSON v2 > Standard |
| Write-heavy workloads | Sonic > JSON v2 ≈ Standard |
| Large string outputs | Sonic > Standard > JSON v2 |

For production systems:

1. **Sonic** for maximum performance (if you can accept external dependencies and higher memory)
2. **JSON v2** for a solid 2x unmarshal boost with zero code changes
3. **Standard JSON** for stability, memory efficiency, and proven reliability

The benchmarks demonstrate that JSON v2 is now a viable upgrade path for read-heavy services, while Sonic remains unmatched for applications where JSON performance is critical.