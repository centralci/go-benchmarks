# Benchmark Analysis: Sonic vs Standard JSON vs JSON v2 in Go

This analysis compares the performance of ByteDance's Sonic JSON parser with Go's standard library encoding/json package
and the experimental JSON v2 implementation across different data sizes and operations.

## Test Environment

- **OS**: Darwin (macOS)
- **Architecture**: arm64
- **CPU**: Apple M1 Pro
- **Go Version**: 1.25.5
- **Sonic Version**: 1.15.0
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
| Marshal   | 1MB   | 1,780   | 667,528      | 807.45            | 541,887       | 2           |
| Marshal   | 10MB  | 176     | 6,728,575    | 801.34            | 5,493,998     | 2           |
| Marshal   | 100MB | 15      | 68,274,561   | 790.18            | 62,901,857    | 5           |
| ToString  | 1MB   | 1,732   | 695,526      | 774.95            | 1,091,116     | 3           |
| ToString  | 10MB  | 164     | 7,261,260    | 742.55            | 12,025,640    | 5           |
| ToString  | 100MB | 15      | 72,712,061   | 741.96            | 161,599,650   | 20          |
| Unmarshal | 1MB   | 386     | 3,079,793    | 175.01            | 748,056       | 7,014       |
| Unmarshal | 10MB  | 38      | 30,626,546   | 176.05            | 8,231,410     | 69,972      |
| Unmarshal | 100MB | 4       | 310,443,416  | 173.78            | 105,351,148   | 699,505     |

### Experimental JSON v2 (GOEXPERIMENT=jsonv2)

| Operation | Size  | Ops/sec | Time (ns/op) | Throughput (MB/s) | Memory (B/op) | Allocations |
|-----------|-------|---------|--------------|-------------------|---------------|-------------|
| Marshal   | 1MB   | 1,910   | 623,239      | 864.63            | 542,054       | 3           |
| Marshal   | 10MB  | 187     | 6,592,623    | 818.30            | 5,540,380     | 3           |
| Marshal   | 100MB | 18      | 63,508,757   | 849.51            | 71,431,914    | 6           |
| ToString  | 1MB   | 1,854   | 641,048      | 840.61            | 1,092,348     | 4           |
| ToString  | 10MB  | 176     | 6,594,401    | 818.08            | 12,453,866    | 6           |
| ToString  | 100MB | 13      | 78,203,849   | 689.88            | 374,125,598   | 50          |
| Unmarshal | 1MB   | 762     | 1,571,959    | 342.80            | 744,664       | 6,538       |
| Unmarshal | 10MB  | 79      | 14,886,932   | 362.38            | 8,199,274     | 65,404      |
| Unmarshal | 100MB | 7       | 150,372,988  | 358.78            | 104,966,089   | 653,033     |

### Sonic Standard

| Operation | Size  | Ops/sec | Time (ns/op) | Throughput (MB/s) | Memory (B/op) | Allocations |
|-----------|-------|---------|--------------|-------------------|---------------|-------------|
| Marshal   | 1MB   | 2,732   | 440,140      | 1,224.60          | 540,979       | 3           |
| Marshal   | 10MB  | 231     | 5,135,821    | 1,049.85          | 23,053,051    | 23          |
| Marshal   | 100MB | 22      | 50,008,390   | 1,078.80          | 167,918,289   | 31          |
| ToString  | 1MB   | 2,715   | 441,499      | 1,220.84          | 540,878       | 3           |
| ToString  | 10MB  | 228     | 5,161,551    | 1,044.62          | 23,052,453    | 23          |
| ToString  | 100MB | 22      | 49,123,117   | 1,098.24          | 167,918,930   | 32          |
| Unmarshal | 1MB   | 2,404   | 472,962      | 1,139.62          | 669,491       | 506         |
| Unmarshal | 10MB  | 231     | 5,055,148    | 1,066.61          | 33,167,886    | 5,032       |
| Unmarshal | 100MB | 22      | 48,787,597   | 1,105.80          | 483,713,380   | 50,244      |

### Sonic Fastest

| Operation | Size  | Ops/sec | Time (ns/op) | Throughput (MB/s) | Memory (B/op) | Allocations |
|-----------|-------|---------|--------------|-------------------|---------------|-------------|
| Marshal   | 1MB   | 2,691   | 438,936      | 1,227.97          | 541,753       | 3           |
| Marshal   | 10MB  | 231     | 5,141,631    | 1,048.67          | 23,053,647    | 23          |
| Marshal   | 100MB | 22      | 50,343,438   | 1,071.62          | 167,911,954   | 31          |
| ToString  | 1MB   | 2,719   | 439,195      | 1,227.24          | 541,589       | 3           |
| ToString  | 10MB  | 229     | 5,161,919    | 1,044.55          | 23,053,059    | 23          |
| ToString  | 100MB | 22      | 49,684,955   | 1,085.83          | 167,912,604   | 32          |
| Unmarshal | 1MB   | 2,492   | 474,916      | 1,134.93          | 667,653       | 506         |
| Unmarshal | 10MB  | 237     | 5,068,177    | 1,063.87          | 33,135,430    | 5,032       |
| Unmarshal | 100MB | 24      | 48,612,246   | 1,109.79          | 481,165,611   | 50,244      |

## Performance Comparisons

### Marshal Performance (Relative to Standard JSON)

| Size  | JSON v2 | Sonic  | Sonic Fastest |
|-------|---------|--------|---------------|
| 1MB   | +7.1%   | +51.7% | +52.1%        |
| 10MB  | +2.1%   | +31.0% | +30.9%        |
| 100MB | +7.5%   | +36.5% | +35.6%        |

### ToString Performance (Relative to Standard JSON)

| Size  | JSON v2 | Sonic  | Sonic Fastest |
|-------|---------|--------|---------------|
| 1MB   | +8.5%   | +57.5% | +58.4%        |
| 10MB  | +10.1%  | +40.7% | +40.7%        |
| 100MB | -7.0%   | +48.0% | +46.3%        |

### Unmarshal Performance (Relative to Standard JSON)

| Size  | JSON v2  | Sonic   | Sonic Fastest |
|-------|----------|---------|---------------|
| 1MB   | +95.9%   | +551.2% | +548.5%       |
| 10MB  | +105.7%  | +505.8% | +504.3%       |
| 100MB | +106.5%  | +536.4% | +538.7%       |

## Key Findings

### 1. JSON v2 Delivers on Unmarshal Promise

**Performance Gains over Standard JSON**:

- **1MB**: 96% faster (1.96x)
- **10MB**: 106% faster (2.06x)
- **100MB**: 107% faster (2.07x)

JSON v2 consistently delivers ~2x faster unmarshaling as promised, with slightly reduced allocations (6,538 vs 7,014 for 1MB).

### 2. JSON v2 Marshal: Modest Improvement

**Performance Gains over Standard JSON**:

- **1MB**: 7.1% faster
- **10MB**: 2.1% faster
- **100MB**: 7.5% faster

Marshal performance improved compared to previous experimental versions that showed regressions.

### 3. JSON v2 String Output: Mixed Results

**Performance vs Standard JSON**:

- **1MB**: 8.5% faster
- **10MB**: 10.1% faster
- **100MB**: 7.0% **slower** (690 MB/s vs 742 MB/s)

**Memory Impact at 100MB**: JSON v2 uses 374MB vs 162MB (2.3x more). The 100MB string output remains a weak point.

### 4. Sonic's Continued Dominance

**Sonic vs JSON v2**:

| Operation | 1MB     | 10MB    | 100MB   |
|-----------|---------|---------|---------|
| Marshal   | +42%    | +28%    | +27%    |
| ToString  | +45%    | +28%    | +59%    |
| Unmarshal | +233%   | +194%   | +208%   |

Sonic remains **3x faster** than JSON v2 for unmarshaling, even after JSON v2's 2x improvement.

### 5. Memory Usage Patterns

| Operation    | Library  | 1MB Memory | 10MB Memory   | 100MB Memory   |
|--------------|----------|------------|---------------|----------------|
| **Marshal**  |          |            |               |                |
|              | Standard | 542KB      | 5.5MB         | 63MB           |
|              | JSON v2  | 542KB      | 5.5MB         | 71MB (+13%)    |
|              | Sonic    | 541KB      | 23MB (4.2x)   | 168MB (2.7x)   |
| **ToString** |          |            |               |                |
|              | Standard | 1.1MB      | 12MB          | 162MB          |
|              | JSON v2  | 1.1MB      | 12.5MB        | 374MB (2.3x)   |
|              | Sonic    | 541KB      | 23MB          | 168MB          |
| **Unmarshal**|          |            |               |                |
|              | Standard | 748KB      | 8.2MB         | 105MB          |
|              | JSON v2  | 745KB      | 8.2MB         | 105MB          |
|              | Sonic    | 669KB      | 33MB (4x)     | 484MB (4.6x)   |

## Analysis by Use Case

### 1. String Output for APIs

**Clear Winner**: Sonic

- 40-58% faster than standard JSON
- 28-59% faster than JSON v2
- Direct string generation without conversion overhead

**Avoid**: JSON v2 for large payloads (100MB: slower and uses 2.3x memory)

### 2. Conservative Upgrade Path

**Recommendation**: JSON v2

- 2x faster unmarshaling (as promised)
- 2-8% faster marshaling
- No external dependencies
- Drop-in replacement with `GOEXPERIMENT=jsonv2`

**Caveat**: Avoid for large string outputs (100MB+)

### 3. Maximum Performance

**Recommendation**: Sonic

- 6x faster unmarshaling (3x faster than JSON v2)
- 31-52% faster marshaling
- 40-58% faster string output

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

1. **Sonic**: 31-52% improvement
2. **JSON v2**: 2-8% improvement, good conservative choice
3. **Standard JSON**: Reliable baseline

### API Services (String Output)

1. **Sonic MarshalString()**: 40-58% faster
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
// 1MB: 1.1MB memory, 775 MB/s
```

### JSON v2 (Optimal for no-dependency upgrade)

```go
// Build with: GOEXPERIMENT=jsonv2 go build
data, _ := json.Marshal(obj)  // Same API, 2x faster unmarshal
str := string(data)
// 1MB unmarshal: 343 MB/s (vs 175 MB/s standard)
```

### Sonic (Optimal for performance)

```go
str, _ := sonic.MarshalString(obj) // Direct generation
// 1MB: 541KB memory, 1,221 MB/s (58% faster than standard)

var result []Data
sonic.Unmarshal(jsonBytes, &result)
// 1MB: 1,140 MB/s (6.5x faster than standard, 3.3x faster than JSON v2)
```

## Conclusion (Jan 2026)

**JSON v2** has matured significantly:

- ✅ **Delivers** on 2x faster unmarshal promise
- ✅ **Improved** marshal performance (2-8% faster)
- ⚠️ **Mixed** string output (slower at 100MB with 2.3x higher memory)

**Sonic** remains the performance leader:

- ✅ **6x faster** unmarshaling vs standard (3x faster than JSON v2)
- ✅ **31-52% faster** marshaling
- ✅ **40-58% faster** string output
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