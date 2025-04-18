# Benchmark Analysis: Sonic vs Standard JSON in Go

This analysis compares the performance of ByteDance's Sonic JSON parser with Go's standard library encoding/json package
across different data sizes and operations.

## Test Environment

- **OS**: Darwin (macOS)
- **Architecture**: arm64
- **CPU**: Apple M1 Pro
- **Data Sizes**: 1MB (250 rows), 10MB (2,500 rows), 100MB (25,000 rows)
- **Operations**: Marshal (Go → JSON) and Unmarshal (JSON → Go)
- **Libraries**:
    - Go standard library (encoding/json)
    - ByteDance Sonic (standard configuration)
    - ByteDance Sonic (Fastest configuration)

## Results Summary

| Operation | Size  | Library       | Ops/sec | Time (ns/op) | Throughput (MB/s) | Memory (B/op) | Allocations |
|-----------|-------|---------------|---------|--------------|-------------------|---------------|-------------|
| Marshal   | 1MB   | Standard      | 1,996   | 597,638      | 904.27            | 543,877       | 2           |
| Marshal   | 1MB   | Sonic         | 2,294   | 493,572      | 1,094.93          | 541,242       | 3           |
| Marshal   | 1MB   | Sonic Fastest | 2,286   | 495,058      | 1,091.65          | 541,751       | 3           |
| Unmarshal | 1MB   | Standard      | 397     | 2,979,115    | 181.41            | 749,896       | 7,016       |
| Unmarshal | 1MB   | Sonic         | 2,426   | 469,300      | 1,151.56          | 666,953       | 510         |
| Unmarshal | 1MB   | Sonic Fastest | 2,371   | 471,210      | 1,146.89          | 666,993       | 510         |
| Marshal   | 10MB  | Standard      | 198     | 5,996,535    | 899.05            | 5,398,652     | 2           |
| Marshal   | 10MB  | Sonic         | 230     | 5,174,333    | 1,041.91          | 17,099,370    | 15          |
| Marshal   | 10MB  | Sonic Fastest | 229     | 5,180,989    | 1,040.57          | 17,099,376    | 15          |
| Unmarshal | 10MB  | Standard      | 38      | 29,921,751   | 180.18            | 8,230,885     | 69,957      |
| Unmarshal | 10MB  | Sonic         | 243     | 4,882,403    | 1,104.21          | 33,166,332    | 5,026       |
| Unmarshal | 10MB  | Sonic Fastest | 243     | 4,931,819    | 1,093.14          | 33,135,479    | 5,026       |
| Marshal   | 100MB | Standard      | 18      | 62,312,974   | 865.71            | 61,410,324    | 4           |
| Marshal   | 100MB | Sonic         | 21      | 51,223,653   | 1,053.12          | 153,584,899   | 23          |
| Marshal   | 100MB | Sonic Fastest | 20      | 51,227,715   | 1,053.04          | 153,585,234   | 23          |
| Unmarshal | 100MB | Standard      | 4       | 300,303,167  | 179.63            | 105,348,484   | 699,498     |
| Unmarshal | 100MB | Sonic         | 24      | 47,519,102   | 1,135.22          | 483,452,097   | 50,246      |
| Unmarshal | 100MB | Sonic Fastest | 24      | 47,506,182   | 1,135.53          | 483,452,105   | 50,246      |

## Key Findings

### 1. Performance Comparison

#### Marshaling (Go → JSON)

- **Speed**:
    - Sonic is consistently faster than the standard library, with improvements of:
        - 1MB: ~15% faster
        - 10MB: ~16% faster
        - 100MB: ~18% faster
    - Sonic Fastest configuration shows nearly identical performance to regular Sonic:
        - 1MB: Almost identical (difference <0.5%)
        - 10MB: Almost identical (difference <0.5%)
        - 100MB: Almost identical (difference <0.5%)

- **Throughput**: Both Sonic variants achieve higher throughput across all data sizes:
    - 1MB: ~1,090-1,094 MB/s vs 904 MB/s (standard)
    - 10MB: ~1,040 MB/s vs 899 MB/s (standard)
    - 100MB: ~1,053 MB/s vs 865 MB/s (standard)

- **Memory Usage**: Sonic configurations use comparable memory for small data but significantly more for larger
  datasets:
    - 1MB: Similar (~540KB for all)
    - 10MB: Sonic variants use ~3.2x more memory
    - 100MB: Sonic variants use ~2.5x more memory

#### Unmarshaling (JSON → Go)

- **Speed**: Both Sonic variants dramatically outperform the standard library:
    - 1MB: ~6.3x faster
    - 10MB: ~6.1x faster
    - 100MB: ~6.3x faster
    - Sonic Fastest configuration is nearly identical to regular Sonic:
        - 1MB: Regular Sonic is ~0.4% faster
        - 10MB: Regular Sonic is ~1% faster
        - 100MB: Nearly identical (difference <0.1%)

- **Throughput**: Both Sonic variants achieve much higher throughput:
    - 1MB: ~1,146-1,151 MB/s vs 181 MB/s (standard)
    - 10MB: ~1,093-1,104 MB/s vs 180 MB/s (standard)
    - 100MB: ~1,135 MB/s vs 179 MB/s (standard)

- **Memory Usage**: Sonic variants use more memory but with far fewer allocations:
    - 1MB: Sonic uses ~11% less memory with 7,016 vs 510 allocations
    - 10MB: Sonic uses ~4x more memory but 14x fewer allocations
    - 100MB: Sonic uses ~4.6x more memory but 14x fewer allocations

### 2. Scaling Characteristics

- **Standard Library**: Shows consistent throughput across sizes:
    - Marshal: ~865-904 MB/s regardless of size
    - Unmarshal: ~179-181 MB/s regardless of size

- **Sonic**: Also maintains consistent throughput across sizes:
    - Marshal: ~1,041-1,094 MB/s
    - Unmarshal: ~1,104-1,151 MB/s

- **Sonic Fastest**: Performs nearly identically to regular Sonic:
    - Marshal: ~1,040-1,091 MB/s
    - Unmarshal: ~1,093-1,146 MB/s

- **Memory Scaling**: All implementations show roughly linear memory scaling with data size, but Sonic's memory usage
  grows faster with size.

### 3. Memory vs Speed Tradeoff

- Sonic makes a clear tradeoff: it uses more memory (especially for large datasets) to achieve significantly better
  performance.
- For unmarshaling, Sonic uses more total memory but makes far fewer allocations, which contributes to its speed
  advantage.

## Conclusions

1. **For Marshaling**: Sonic offers moderate speed improvements (15-18%) over the standard library with increased memory
   usage. The advantage is consistent but not dramatic.

2. **For Unmarshaling**: Sonic offers exceptional performance improvements (~6.3x faster) compared to the standard
   library. This comes at the cost of higher memory usage, but with far fewer allocations.

3. **Sonic vs Sonic Fastest**: The "Fastest" configuration of Sonic shows nearly identical performance to standard
   Sonic, with differences typically less than 1%. This suggests that for Go applications, the default Sonic
   configuration is already well-optimized.

4. **Size Considerations**:
    - For small data (1MB): Sonic variants are clearly superior with minimal memory overhead.
    - For medium data (10MB): Sonic's performance advantage outweighs its higher memory usage in most scenarios.
    - For large data (100MB): Sonic maintains its performance edge but requires significant additional memory.

5. **Use Case Recommendations**:
    - **API Servers**: Sonic is ideal for high-throughput JSON parsing in API servers where performance is critical.
    - **Memory-Constrained Environments**: The standard library may be preferable where memory is severely limited.
    - **Unmarshal-Heavy Workloads**: Applications that process large amounts of incoming JSON will benefit dramatically
      from Sonic.

6. **Overall**: ByteDance's Sonic JSON parser delivers on its promise of high-performance JSON processing. Its advantage
   is most pronounced for unmarshaling operations, where it achieves dramatic speedups through more efficient parsing
   algorithms and memory management strategies. The "Fastest" configuration provides little additional benefit in Go
   applications over the standard Sonic configuration.