# Benchmark Analysis: YAML Libraries in Go

This analysis compares the performance of three popular YAML libraries for Go:

- yaml.v2 (gopkg.in/yaml.v2)
- yaml.v3 (gopkg.in/yaml.v3)
- go-yaml (github.com/goccy/go-yaml)

## Test Environment

- **OS**: Darwin (macOS)
- **Architecture**: arm64
- **CPU**: Apple M1 Pro
- **Data Types**:
    - Simple (basic struct with primitive types)
    - Complex (nested structs with maps, arrays, and varied data types)
    - Template (CI/CD pipeline configuration template)
- **Operations**: Marshal (Go → YAML) and Unmarshal (YAML → Go)

## Results Summary

| Operation | Data Type | Library | Speed (ns/op) | Relative Performance |
|-----------|-----------|---------|---------------|----------------------|
| Marshal   | Simple    | yaml.v2 | 3,845         | Fastest (1.0x)       |
| Marshal   | Simple    | yaml.v3 | 4,653         | 1.2x slower          |
| Marshal   | Simple    | go-yaml | 8,746         | 2.3x slower          |
| Unmarshal | Simple    | yaml.v2 | 5,562         | Fastest (1.0x)       |
| Unmarshal | Simple    | yaml.v3 | 7,949         | 1.4x slower          |
| Unmarshal | Simple    | go-yaml | 12,307        | 2.2x slower          |
| Marshal   | Complex   | yaml.v2 | 165,577       | Fastest (1.0x)       |
| Marshal   | Complex   | yaml.v3 | 224,269       | 1.4x slower          |
| Marshal   | Complex   | go-yaml | 468,760       | 2.8x slower          |
| Unmarshal | Complex   | yaml.v3 | 144,317       | Fastest (1.0x)       |
| Unmarshal | Complex   | yaml.v2 | 151,490       | 1.05x slower         |
| Unmarshal | Complex   | go-yaml | 539,876       | 3.7x slower          |
| Unmarshal | Template  | yaml.v3 | 145,784       | Fastest (1.0x)       |
| Unmarshal | Template  | yaml.v2 | 169,573       | 1.2x slower          |
| Unmarshal | Template  | go-yaml | 179,797       | 1.2x slower          |

## Key Findings

### 1. Performance by Library

#### yaml.v2

- **Strengths**: Fastest for simple data marshaling, competitive across most operations
- **Performance**:
    - Simple data: Excellent
    - Complex data: Very good for marshaling, good for unmarshaling
    - Template data: Good

#### yaml.v3

- **Strengths**: Best complex data unmarshaling, best template unmarshaling
- **Performance**:
    - Simple data: Good (slightly slower than yaml.v2)
    - Complex data: Very good for unmarshaling, good for marshaling
    - Template data: Excellent

#### go-yaml (goccy)

- **Strengths**: Offers unique features like stream processing
- **Performance**:
    - Simple data: Significantly slower than the alternatives
    - Complex data: Significantly slower than the alternatives
    - Template data: Comparable to yaml.v2 but slower than yaml.v3

### 2. Performance by Operation Type

#### Marshaling (Go → YAML)

- yaml.v2 is consistently the fastest for marshaling operations
- yaml.v3 is typically 1.2x-1.4x slower than yaml.v2
- go-yaml is significantly slower (2.3x-2.8x) than yaml.v2

#### Unmarshaling (YAML → Go)

- For simple data: yaml.v2 is fastest
- For complex data and templates: yaml.v3 has an edge
- go-yaml is consistently the slowest, particularly with complex data

### 3. Mixed and Real-World Operations

#### Mixed Template Processing (Unmarshal then Marshal)

| Library | Speed (ns/op) | Relative Performance |
|---------|---------------|----------------------|
| yaml.v2 | 238,367       | Fastest (1.0x)       |
| yaml.v3 | 299,954       | 1.3x slower          |
| go-yaml | 505,169       | 2.1x slower          |

#### Random Data Processing

| Operation | Data Type | Library | Speed (ns/op) | Relative Performance |
|-----------|-----------|---------|---------------|----------------------|
| Mixed     | Simple    | yaml.v2 | 12,763        | Fastest (1.0x)       |
| Mixed     | Simple    | yaml.v3 | 15,644        | 1.2x slower          |
| Mixed     | Simple    | go-yaml | 25,013        | 2.0x slower          |
| Mixed     | Complex   | yaml.v2 | 484,533       | Fastest (1.0x)       |
| Mixed     | Complex   | yaml.v3 | 602,455       | 1.2x slower          |
| Mixed     | Complex   | go-yaml | 971,051       | 2.0x slower          |
| Mixed     | Template  | yaml.v2 | 157,994       | Fastest (1.0x)       |
| Mixed     | Template  | yaml.v3 | 189,908       | 1.2x slower          |
| Mixed     | Template  | go-yaml | 321,343       | 2.0x slower          |

#### Bulk Operations (50 items)

| Operation | Data Type | Library | Speed (ns/op) | Relative Performance |
|-----------|-----------|---------|---------------|----------------------|
| Marshal   | Simple    | yaml.v2 | 238,134       | Fastest (1.0x)       |
| Marshal   | Simple    | yaml.v3 | 266,939       | 1.1x slower          |
| Marshal   | Simple    | go-yaml | 515,107       | 2.2x slower          |
| Marshal   | Complex   | yaml.v2 | 10,546,540    | Fastest (1.0x)       |
| Marshal   | Complex   | yaml.v3 | 12,084,435    | 1.1x slower          |
| Marshal   | Complex   | go-yaml | 22,626,796    | 2.1x slower          |

### 4. Specialized Features

#### Stream Processing (go-yaml specific)

- go-yaml offers stream processing capabilities not available in other libraries
- Performance: 317,366 ns/op

## Conclusions

1. **For General Use**: yaml.v2 offers the best overall performance and is an excellent default choice for most Go YAML
   processing needs.

2. **For Complex Unmarshaling**: yaml.v3 shows superior performance for unmarshaling complex data structures and
   templates, making it ideal for configuration file processing and similar use cases.

3. **For Advanced Features**: go-yaml provides additional features like stream processing but at a significant
   performance cost (generally 2-3x slower than the alternatives).

4. **Library Selection Considerations**:
    - **Speed-Critical Applications**: Use yaml.v2 for marshaling and yaml.v3 for complex unmarshaling
    - **Feature Requirements**: If stream processing or other advanced features are needed, the performance cost of
      go-yaml may be justified
    - **Maintainability**: yaml.v3 is the newer and actively maintained version of the canonical Go YAML package

5. **Data Complexity Impact**:
    - Performance differences between libraries increase with data complexity
    - For very simple data, the absolute performance differences may be negligible
    - For complex data or bulk operations, choosing the right library can have a substantial impact

6. **Operation Type Impact**:
    - Marshal operations show less variation between libraries than unmarshal operations
    - The libraries show specialized performance characteristics for different data types and operations

Overall, yaml.v2 and yaml.v3 both offer excellent performance, with specific strengths depending on the task. The
go-yaml library from goccy prioritizes features over raw performance and may be appropriate when those features are
required.