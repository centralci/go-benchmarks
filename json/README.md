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
- **Operations**: Marshal (Go → JSON) and Unmarshal (JSON → Go)
- **Libraries**:
    - Go standard library (encoding/json)
    - Go experimental JSON v2 (GOEXPERIMENT=jsonv2)
    - ByteDance Sonic (standard configuration)
    - ByteDance Sonic (Fastest configuration)

## Results Summary

### Standard JSON Library (encoding/json)

| Operation | Size  | Ops/sec | Time (ns/op) | Throughput (MB/s) | Memory (B/op) | Allocations |
|-----------|-------|---------|--------------|-------------------|---------------|-------------|
| Marshal   | 1MB   | 1,959   | 610,353      | 882.83            | 541,784       | 2           |
| Marshal   | 10MB  | 193     | 6,165,694    | 874.80            | 5,398,654     | 2           |
| Marshal   | 100MB | 18      | 62,439,028   | 864.20            | 61,418,554    | 4           |
| Unmarshal | 1MB   | 379     | 3,125,512    | 172.40            | 748,480       | 7,010       |
| Unmarshal | 10MB  | 38      | 31,286,201   | 172.40            | 8,231,273     | 69,959      |
| Unmarshal | 100MB | 4       | 314,478,188  | 171.58            | 105,354,036   | 699,480     |

### Experimental JSON v2 (GOEXPERIMENT=jsonv2)

| Operation | Size  | Ops/sec | Time (ns/op) | Throughput (MB/s) | Memory (B/op) | Allocations |
|-----------|-------|---------|--------------|-------------------|---------------|-------------|
| Marshal   | 1MB   | 1,927   | 624,070      | 863.65            | 542,046       | 3           |
| Marshal   | 10MB  | 188     | 6,300,620    | 856.24            | 5,539,678     | 3           |
| Marshal   | 100MB | 18      | 65,392,123   | 825.02            | 88,911,580    | 9           |
| Unmarshal | 1MB   | 776     | 1,505,468    | 358.01            | 743,571       | 6,533       |
| Unmarshal | 10MB  | 78      | 15,071,988   | 357.94            | 8,198,108     | 65,356      |
| Unmarshal | 100MB | 7       | 154,660,244  | 348.83            | 104,950,678   | 653,108     |

### Sonic Standard

| Operation | Size  | Ops/sec | Time (ns/op) | Throughput (MB/s) | Memory (B/op) | Allocations |
|-----------|-------|---------|--------------|-------------------|---------------|-------------|
| Marshal   | 1MB   | 2,562   | 449,775      | 1,198.33          | 540,738       | 3           |
| Marshal   | 10MB  | 228     | 5,158,108    | 1,045.90          | 23,216,768    | 25          |
| Marshal   | 100MB | 22      | 50,555,526   | 1,067.15          | 206,166,227   | 33          |
| Unmarshal | 1MB   | 2,384   | 482,393      | 1,117.30          | 665,167       | 508         |
| Unmarshal | 10MB  | 237     | 4,979,802    | 1,083.35          | 33,151,878    | 5,026       |
| Unmarshal | 100MB | 22      | 49,441,763   | 1,091.18          | 483,697,166   | 50,245      |

### Sonic Fastest

| Operation | Size  | Ops/sec | Time (ns/op) | Throughput (MB/s) | Memory (B/op) | Allocations |
|-----------|-------|---------|--------------|-------------------|---------------|-------------|
| Marshal   | 1MB   | 2,572   | 447,438      | 1,204.59          | 540,738       | 3           |
| Marshal   | 10MB  | 230     | 5,151,831    | 1,047.17          | 23,217,372    | 25          |
| Marshal   | 100MB | 20      | 50,988,033   | 1,058.09          | 206,166,730   | 33          |
| Unmarshal | 1MB   | 2,395   | 480,562      | 1,121.56          | 666,044       | 508         |
| Unmarshal | 10MB  | 238     | 5,030,564    | 1,072.42          | 33,183,377    | 5,026       |
| Unmarshal | 100MB | 24      | 49,123,267   | 1,098.26          | 483,484,825   | 50,245      |

## Performance Comparisons

### Marshal Performance (Relative to Standard JSON)

| Size  | JSON v2 | Sonic  | Sonic Fastest |
|-------|---------|--------|---------------|
| 1MB   | -2.2%   | +35.7% | +36.4%        |
| 10MB  | -2.1%   | +19.6% | +19.7%        |
| 100MB | -4.5%   | +23.5% | +22.5%        |

### Unmarshal Performance (Relative to Standard JSON)

| Size  | JSON v2 | Sonic   | Sonic Fastest |
|-------|---------|---------|---------------|
| 1MB   | +107.6% | +548.0% | +550.3%       |
| 10MB  | +107.6% | +528.4% | +522.0%       |
| 100MB | +103.3% | +536.0% | +540.2%       |

## Key Findings

### 1. JSON v2 Performance Characteristics

#### Marshaling

- **Performance**: JSON v2 is slightly slower than the current standard library (2-5% slower)
- **Memory**: Uses comparable memory to standard JSON, with 1 additional allocation
- **Scaling**: Degrades more noticeably with larger datasets (100MB shows 4.5% slower performance)

#### Unmarshaling

- **Performance**: JSON v2 is ~2x faster than standard JSON, a significant improvement
- **Memory**: Uses slightly less memory with fewer allocations (6-7% reduction)
- **Consistency**: Maintains ~358 MB/s throughput for small data, ~349 MB/s for large data

### 2. Sonic Performance Advantages

#### Against Standard JSON

- **Marshal**: 20-36% faster depending on size
- **Unmarshal**: 5.2-5.5x faster across all sizes

#### Against JSON v2

- **Marshal**: 22-40% faster
- **Unmarshal**: 2.6-3.2x faster

### 3. Memory Trade-offs

| Library       | Marshal Memory Usage                          | Unmarshal Memory Usage           |
|---------------|-----------------------------------------------|----------------------------------|
| Standard JSON | Baseline                                      | Baseline                         |
| JSON v2       | Similar to standard (+0.1% to +45% for 100MB) | Slightly less (-0.4% to -0.6%)   |
| Sonic         | 1MB: Similar; 10MB: 4.3x; 100MB: 3.4x         | 1MB: -11%; 10MB: 4x; 100MB: 4.6x |

### 4. Allocation Patterns

- **Standard JSON**: Minimal allocations for marshal (2-4), massive for unmarshal (7K-699K)
- **JSON v2**: Similar to standard but slightly fewer unmarshal allocations (6.5K-653K)
- **Sonic**: More marshal allocations (3-33) but dramatically fewer unmarshal allocations (508-50K)

## Analysis by Use Case

### 1. Conservative Upgrade Path (JSON v2)

**Pros:**

- 2x faster unmarshaling with minimal code changes
- Slightly reduced memory usage for unmarshaling
- Maintains similar memory profile to standard library
- Official Go experimental feature

**Cons:**

- Slightly slower marshaling (2-5%)
- Still significantly slower than Sonic
- Experimental status may change

**Best for:** Teams wanting moderate improvements without external dependencies

### 2. Maximum Performance (Sonic)

**Pros:**

- 20-36% faster marshaling
- 5.2-5.5x faster unmarshaling
- Consistent high throughput across all sizes
- Production-proven at ByteDance scale

**Cons:**

- External dependency
- Significantly higher memory usage for large datasets
- More complex memory allocation patterns

**Best for:** High-throughput services where performance is critical

### 3. Current Standard Library

**Pros:**

- No dependencies or experimental flags
- Predictable, well-understood behavior
- Minimal memory footprint for marshaling

**Cons:**

- Slowest performance across all operations
- Massive allocation count for unmarshaling

**Best for:** Simple applications or memory-constrained environments

## Recommendations

1. **For API Services**: Use Sonic for maximum throughput, especially if handling large request volumes

2. **For Conservative Teams**: Consider JSON v2 as it offers significant unmarshal improvements with minimal risk

3. **For Memory-Constrained Systems**: Stick with standard JSON or carefully evaluate JSON v2

4. **For Mixed Workloads**:
    - Heavy unmarshaling: Sonic provides 5x improvement
    - Heavy marshaling: Sonic provides 20-36% improvement
    - Balanced: JSON v2 offers a reasonable middle ground

5. **Migration Strategy**:
    - Start with JSON v2 for easy 2x unmarshal gains
    - Move to Sonic for services requiring maximum performance
    - Keep standard JSON for stable, low-traffic services

## Conclusion

The experimental JSON v2 provides a solid middle ground with 2x faster unmarshaling while maintaining compatibility and
similar memory characteristics to the standard library. However, Sonic remains the clear performance leader with 5x
faster unmarshaling and significant marshaling improvements, making it the best choice for performance-critical
applications despite its higher memory usage. The JSON v2 experiment shows Go's commitment to improving JSON
performance, but third-party solutions like Sonic still offer superior performance for teams willing to take on external
dependencies.