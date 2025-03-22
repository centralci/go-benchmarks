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

| Level | Implementation | Text (MB/s) | Binary (MB/s) | Notes                              |
|-------|----------------|-------------|---------------|------------------------------------|
| 1     | Klauspost      | 198-205     | 1544-1607     | Consistent across data sizes       |
| 1     | DataDog        | 463-466     | 2927-3226     | ~2x faster than Klauspost          |
| 3     | Klauspost      | 165-185     | 1042-1162     |                                    |
| 3     | DataDog        | 265-280     | 1551-1677     | Significantly faster               |
| 7     | Klauspost      | 101-140     | 455-545       | Better at text compression         |
| 7     | DataDog        | 85-86       | 652-725       | Slower for text, faster for binary |

### Decompression Speed

| Level | Implementation | Text (MB/s) | Binary (MB/s) | Notes                      |
|-------|----------------|-------------|---------------|----------------------------|
| 1-7   | Klauspost      | 472-604     | 3320-3795     | Speed increases with level |
| 1-7   | DataDog        | 1184-1358   | 5374-16322    | 2-4x faster than Klauspost |

### Compression Ratio

| Level | Implementation | Text Ratio | Binary Ratio | Notes                          |
|-------|----------------|------------|--------------|--------------------------------|
| 1     | Klauspost      | 2.87-2.94  | 2.12         | Better than DataDog at level 1 |
| 1     | DataDog        | 2.76-2.80  | 2.14         |                                |
| 3     | Klauspost      | 3.13-3.20  | -            | Similar to DataDog             |
| 3     | DataDog        | 3.10-3.24  | -            |                                |
| 7     | Klauspost      | 3.24-3.53  | 2.15         |                                |
| 7     | DataDog        | 3.46-3.57  | 2.15         | Slightly better compression    |

## Key Insights

1. **Implementation Tradeoffs**:
    - DataDog (CGo): Excels at all compression levels for binary data and all decompression tasks
    - Klauspost (Pure Go): Better at level 7 text compression

2. **Performance Patterns**:
    - Binary data compresses/decompresses significantly faster than text for both implementations
    - Decompression speed increases with compression level for Klauspost
    - Both implementations show excellent scaling across data sizes

3. **Level Selection Impact**:
    - Level 1 → 7 compression speed cost:
        * Klauspost: ~50% slower for text, ~70% slower for binary
        * DataDog: ~5x slower for text, ~77% slower for binary
    - Level 1 → 7 compression ratio gain:
        * Klauspost: ~13-20% improvement
        * DataDog: ~25-28% improvement

4. **Implementation Differences**:
    - DataDog's CGo approach delivers exceptional binary data processing and decompression
    - Klauspost's pure Go implementation shows more balanced level scaling for text data
    - DataDog's decompression of binary data is remarkably fast (up to 16 GB/s)

## Recommendations

| Use Case                                 | Recommended Implementation     |
|------------------------------------------|--------------------------------|
| Decompression-heavy workloads            | DataDog (2-4x faster)          |
| Binary data processing                   | DataDog (faster at all levels) |
| Level 1 compression                      | DataDog (2x faster)            |
| Level 7 text compression                 | Klauspost (1.2x faster)        |
| Pure Go requirement                      | Klauspost (only option)        |
| Read-heavy with tight memory constraints | DataDog                        |
| Best compression ratio at high levels    | DataDog (slightly better)      |

## Conclusion

Both implementations offer excellent performance with different strengths:

- **DataDog/zstd** leverages the C implementation for exceptional decompression speed (up to 16 GB/s for binary data)
  and faster compression of binary data at all levels. It achieves slightly better compression ratios at high levels but
  performance drops significantly for text compression at level 7.

- **Klauspost/compress/zstd** provides solid performance in a pure Go implementation without CGo dependencies. It excels
  particularly at high-level text compression and shows more consistent scaling across compression levels for text data.

The choice depends primarily on your data type (text vs binary), compression level requirements, decompression needs,
and whether CGo is acceptable in your environment.