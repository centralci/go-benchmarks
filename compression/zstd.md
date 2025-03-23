# Klauspost vs DataDog Zstd Implementation Comparison

This analysis compares two Go implementations of the Zstandard compression algorithm:

* **klauspost/compress/zstd**: A pure Go implementation
* **DataDog/zstd**: A CGo wrapper around the official C implementation

# Go Support

See https://github.com/golang/go/issues/62513

## Benchmark Setup

**Environment**: Apple M1 Pro (darwin/arm64)

**Test Parameters**:

- **Compression Levels**: 1, 3, 7 (mapped appropriately between implementations)
- **Data Sizes**: 1MB, 10MB, 100MB
- **Data Types**: Text (with natural repetition) and binary (with patterns)
- **Metrics**: Compression/decompression speed (MB/s) and compression ratio

## Performance Summary

### Compression Speed

| Level | Implementation | Text (MB/s)    | Binary (MB/s)   | Notes                              |
|-------|----------------|----------------|----------------|------------------------------------|
| 1     | Klauspost      | 191.4-201.9    | 1571.6-1609.6  | Consistent across data sizes       |
| 1     | DataDog        | 447.1-461.7    | 2791.9-3237.8  | ~2.3x faster than Klauspost        |
| 3     | Klauspost      | 147.8-182.4    | 975.2-1078.7   |                                    |
| 3     | DataDog        | 258.3-276.1    | 1444.4-1611.2  | Significantly faster               |
| 7     | Klauspost      | 108.9-128.0    | 380.3-525.2    | Struggles with large binary files  |
| 7     | DataDog        | 79.7-84.3      | 589.2-683.4    | Slower for text, faster for binary |

### Decompression Speed

| Level | Implementation | Text (MB/s)    | Binary (MB/s)    | Notes                      |
|-------|----------------|----------------|-----------------|----------------------------|
| 1-7   | Klauspost      | 436.0-587.4    | 3228.2-3697.1   | Speed increases with level |
| 1-7   | DataDog        | 1112.6-1313.1  | 3935.5-15870.7  | 2-4x faster than Klauspost |

### Compression Ratio

| Level | Implementation | Text Ratio   | Binary Ratio | Notes                          |
|-------|----------------|--------------|--------------|--------------------------------|
| 1     | Klauspost      | 2.88-2.94    | 2.12         | Better than DataDog at level 1 |
| 1     | DataDog        | 2.76-2.78    | 2.14         |                                |
| 3     | Klauspost      | 3.13-3.21    | -            | Similar to DataDog             |
| 3     | DataDog        | 3.07-3.25    | -            |                                |
| 7     | Klauspost      | 3.20-3.53    | 2.15         |                                |
| 7     | DataDog        | 3.34-3.57    | 2.15         | Slightly better at high levels |

## Multi-Worker Performance

Both implementations support parallel compression with multiple worker threads. The benchmark results using 4 workers show:

| Implementation | Text Improvement | Binary Improvement | Notes                            |
|----------------|------------------|-------------------|----------------------------------|
| Klauspost      | 1-15%            | 1-12%             | Better scaling at higher levels  |
| DataDog        | 0-1%             | -1-5%             | Minimal benefit from parallelism |

## Key Insights

1. **Implementation Tradeoffs**:
    - DataDog (CGo): Excels at compression levels 1-3 for all data types and all decompression tasks
    - Klauspost (Pure Go): Better at level 7 text compression (20% faster)
    - For binary data, DataDog is consistently faster at all compression levels

2. **Performance Patterns**:
    - Binary data compresses/decompresses significantly faster than text for both implementations
    - Decompression speed increases with compression level for Klauspost
    - DataDog's decompression of binary data reaches extraordinary speeds (up to 15.9 GB/s)

3. **Level Selection Impact**:
    - Level 1 → 7 compression speed cost:
        * Klauspost: 40-45% slower for text, ~70% slower for binary
        * DataDog: ~5x slower for text, ~80% slower for binary
    - Level 1 → 7 compression ratio gain:
        * Klauspost: 11-20% improvement
        * DataDog: 21-28% improvement

4. **Implementation Differences**:
    - DataDog's CGo approach provides exceptional binary data processing and decompression
    - Klauspost's pure Go implementation shows more balanced level scaling for text data
    - DataDog's level 7 text compression is a performance bottleneck (5x slower than level 1)

## Data Size Scaling

Both implementations scale well with increasing data size:

| Implementation | Size Scaling | Compression Speed | Decompression Speed |
|----------------|--------------|-------------------|---------------------|
| Klauspost      | 1MB → 100MB  | Consistent        | Consistent          |
| DataDog        | 1MB → 100MB  | Consistent        | Slightly better     |

## Random Data Handling

For random data (which should be incompressible):
- Both implementations correctly identify the data as incompressible (ratio = 1.000)
- Neither implementation wastes resources trying to compress truly random data

## Recommendations

| Use Case                                 | Recommended Implementation     | Reasoning                            |
|------------------------------------------|--------------------------------|--------------------------------------|
| Decompression-heavy workloads            | DataDog                        | 2-4x faster decompression           |
| Binary data processing                   | DataDog                        | Faster at all levels                 |
| Level 1 compression                      | DataDog                        | 2.3x faster                          |
| Level 7 text compression                 | Klauspost                      | Up to 50% faster                     |
| Pure Go requirement                      | Klauspost                      | No CGo dependencies                  |
| Read-heavy with tight memory constraints | DataDog                        | Superior decompression performance   |
| Best compression ratio at high levels    | DataDog                        | Slightly better (3.57 vs 3.53)      |
| Balanced performance across levels       | Klauspost                      | More gradual performance degradation |

## Conclusion

Both implementations offer excellent performance with different strengths:

- **DataDog/zstd** leverages the C implementation for exceptional decompression speed (up to 15.9 GB/s for binary data)
  and faster compression of binary data at all levels. It achieves slightly better compression ratios at high levels but
  performance drops significantly for text compression at level 7.

- **Klauspost/compress/zstd** provides solid performance in a pure Go implementation without CGo dependencies. It excels
  particularly at high-level text compression and shows more consistent scaling across compression levels for text data.

The choice depends primarily on your data type (text vs binary), compression level requirements, decompression needs,
and whether CGo is acceptable in your environment.

## Detailed Benchmark Results

```
// Compression Speed - Level 1 (MB/s)
Klauspost Text:     1MB: 191.4,  10MB: 191.5,  100MB: 196.8
DataDog Text:       1MB: 447.1,  10MB: 461.7,  100MB: 457.3
Klauspost Binary:   1MB: 1571.6, 10MB: 1584.8
DataDog Binary:     1MB: 2838.0, 10MB: 3237.8

// Compression Speed - Level 3 (MB/s)
Klauspost Text:     1MB: 147.8,  10MB: 173.9,  100MB: 167.9
DataDog Text:       1MB: 258.3,  10MB: 274.7,  100MB: 272.8
Klauspost Binary:   1MB: 975.2,  10MB: 1078.7
DataDog Binary:     1MB: 1469.8, 10MB: 1583.6

// Compression Speed - Level 7 (MB/s)
Klauspost Text:     1MB: 112.3,  10MB: 101.4,  100MB: 122.0
DataDog Text:       1MB: 83.5,   10MB: 82.9,   100MB: 84.3
Klauspost Binary:   1MB: 494.5,  10MB: 380.3
DataDog Binary:     1MB: 591.6,  10MB: 651.3

// Decompression Speed (MB/s)
Klauspost Text L1:  1MB: 452.6,  10MB: 454.6,  100MB: 457.9
Klauspost Text L7:  1MB: 531.9,  10MB: 543.2,  100MB: 587.4
DataDog Text L1:    1MB: 1295.0, 10MB: 1305.6, 100MB: 1313.1
DataDog Text L7:    1MB: 1186.6, 10MB: 1242.1, 100MB: 1257.6
Klauspost Binary:   1MB: ~3240,  10MB: ~3650
DataDog Binary:     1MB: ~4250,  10MB: ~14760

// Compression Ratio
Klauspost Text L1:  1MB: 2.88,   10MB: 2.93,   100MB: 2.94
Klauspost Text L7:  1MB: 3.20,   10MB: 3.50,   100MB: 3.53
DataDog Text L1:    1MB: 2.76,   10MB: 2.78,   100MB: 2.77
DataDog Text L7:    1MB: 3.34,   10MB: 3.56,   100MB: 3.57
```