package compression

import (
	"bytes"
	"testing"

	datadog "github.com/DataDog/zstd"
	klauspost "github.com/klauspost/compress/zstd"
)

// Benchmark Klauspost zstd at different levels
func BenchmarkKlauspostZstdCompress1MB_Level1(b *testing.B) {
	benchmarkKlauspostZstdCompress(b, SmallSize, TextData, 1)
}

func BenchmarkKlauspostZstdCompress1MB_Level3(b *testing.B) {
	benchmarkKlauspostZstdCompress(b, SmallSize, TextData, 3)
}

func BenchmarkKlauspostZstdCompress1MB_Level7(b *testing.B) {
	benchmarkKlauspostZstdCompress(b, SmallSize, TextData, 7)
}

func BenchmarkKlauspostZstdCompress10MB_Level1(b *testing.B) {
	benchmarkKlauspostZstdCompress(b, MediumSize, TextData, 1)
}

func BenchmarkKlauspostZstdCompress10MB_Level3(b *testing.B) {
	benchmarkKlauspostZstdCompress(b, MediumSize, TextData, 3)
}

func BenchmarkKlauspostZstdCompress10MB_Level7(b *testing.B) {
	benchmarkKlauspostZstdCompress(b, MediumSize, TextData, 7)
}

func BenchmarkKlauspostZstdCompress100MB_Level1(b *testing.B) {
	benchmarkKlauspostZstdCompress(b, LargeSize, TextData, 1)
}

func BenchmarkKlauspostZstdCompress100MB_Level3(b *testing.B) {
	benchmarkKlauspostZstdCompress(b, LargeSize, TextData, 3)
}

func BenchmarkKlauspostZstdCompress100MB_Level7(b *testing.B) {
	benchmarkKlauspostZstdCompress(b, LargeSize, TextData, 7)
}

func BenchmarkKlauspostZstdCompressBinary1MB_Level1(b *testing.B) {
	benchmarkKlauspostZstdCompress(b, SmallSize, BinaryData, 1)
}

func BenchmarkKlauspostZstdCompressBinary1MB_Level3(b *testing.B) {
	benchmarkKlauspostZstdCompress(b, SmallSize, BinaryData, 3)
}

func BenchmarkKlauspostZstdCompressBinary1MB_Level7(b *testing.B) {
	benchmarkKlauspostZstdCompress(b, SmallSize, BinaryData, 7)
}

func BenchmarkKlauspostZstdCompressBinary10MB_Level1(b *testing.B) {
	benchmarkKlauspostZstdCompress(b, MediumSize, BinaryData, 1)
}

func BenchmarkKlauspostZstdCompressBinary10MB_Level3(b *testing.B) {
	benchmarkKlauspostZstdCompress(b, MediumSize, BinaryData, 3)
}

func BenchmarkKlauspostZstdCompressBinary10MB_Level7(b *testing.B) {
	benchmarkKlauspostZstdCompress(b, MediumSize, BinaryData, 7)
}

// Benchmark Klauspost zstd decompression
func BenchmarkKlauspostZstdDecompress1MB_Level1(b *testing.B) {
	benchmarkKlauspostZstdDecompress(b, SmallSize, TextData, 1)
}

func BenchmarkKlauspostZstdDecompress1MB_Level3(b *testing.B) {
	benchmarkKlauspostZstdDecompress(b, SmallSize, TextData, 3)
}

func BenchmarkKlauspostZstdDecompress1MB_Level7(b *testing.B) {
	benchmarkKlauspostZstdDecompress(b, SmallSize, TextData, 7)
}

func BenchmarkKlauspostZstdDecompress10MB_Level1(b *testing.B) {
	benchmarkKlauspostZstdDecompress(b, MediumSize, TextData, 1)
}

func BenchmarkKlauspostZstdDecompress10MB_Level3(b *testing.B) {
	benchmarkKlauspostZstdDecompress(b, MediumSize, TextData, 3)
}

func BenchmarkKlauspostZstdDecompress10MB_Level7(b *testing.B) {
	benchmarkKlauspostZstdDecompress(b, MediumSize, TextData, 7)
}

func BenchmarkKlauspostZstdDecompress100MB_Level1(b *testing.B) {
	benchmarkKlauspostZstdDecompress(b, LargeSize, TextData, 1)
}

func BenchmarkKlauspostZstdDecompress100MB_Level3(b *testing.B) {
	benchmarkKlauspostZstdDecompress(b, LargeSize, TextData, 3)
}

func BenchmarkKlauspostZstdDecompress100MB_Level7(b *testing.B) {
	benchmarkKlauspostZstdDecompress(b, LargeSize, TextData, 7)
}

func BenchmarkKlauspostZstdDecompressBinary1MB_Level1(b *testing.B) {
	benchmarkKlauspostZstdDecompress(b, SmallSize, BinaryData, 1)
}

func BenchmarkKlauspostZstdDecompressBinary1MB_Level3(b *testing.B) {
	benchmarkKlauspostZstdDecompress(b, SmallSize, BinaryData, 3)
}

func BenchmarkKlauspostZstdDecompressBinary1MB_Level7(b *testing.B) {
	benchmarkKlauspostZstdDecompress(b, SmallSize, BinaryData, 7)
}

func BenchmarkKlauspostZstdDecompressBinary10MB_Level1(b *testing.B) {
	benchmarkKlauspostZstdDecompress(b, MediumSize, BinaryData, 1)
}

func BenchmarkKlauspostZstdDecompressBinary10MB_Level3(b *testing.B) {
	benchmarkKlauspostZstdDecompress(b, MediumSize, BinaryData, 3)
}

func BenchmarkKlauspostZstdDecompressBinary10MB_Level7(b *testing.B) {
	benchmarkKlauspostZstdDecompress(b, MediumSize, BinaryData, 7)
}

// Benchmark DataDog zstd at different levels
func BenchmarkDataDogZstdCompress1MB_Level1(b *testing.B) {
	benchmarkDataDogZstdCompress(b, SmallSize, TextData, 1)
}

func BenchmarkDataDogZstdCompress1MB_Level3(b *testing.B) {
	benchmarkDataDogZstdCompress(b, SmallSize, TextData, 3)
}

func BenchmarkDataDogZstdCompress1MB_Level7(b *testing.B) {
	benchmarkDataDogZstdCompress(b, SmallSize, TextData, 7)
}

func BenchmarkDataDogZstdCompress10MB_Level1(b *testing.B) {
	benchmarkDataDogZstdCompress(b, MediumSize, TextData, 1)
}

func BenchmarkDataDogZstdCompress10MB_Level3(b *testing.B) {
	benchmarkDataDogZstdCompress(b, MediumSize, TextData, 3)
}

func BenchmarkDataDogZstdCompress10MB_Level7(b *testing.B) {
	benchmarkDataDogZstdCompress(b, MediumSize, TextData, 7)
}

func BenchmarkDataDogZstdCompress100MB_Level1(b *testing.B) {
	benchmarkDataDogZstdCompress(b, LargeSize, TextData, 1)
}

func BenchmarkDataDogZstdCompress100MB_Level3(b *testing.B) {
	benchmarkDataDogZstdCompress(b, LargeSize, TextData, 3)
}

func BenchmarkDataDogZstdCompress100MB_Level7(b *testing.B) {
	benchmarkDataDogZstdCompress(b, LargeSize, TextData, 7)
}

func BenchmarkDataDogZstdCompressBinary1MB_Level1(b *testing.B) {
	benchmarkDataDogZstdCompress(b, SmallSize, BinaryData, 1)
}

func BenchmarkDataDogZstdCompressBinary1MB_Level3(b *testing.B) {
	benchmarkDataDogZstdCompress(b, SmallSize, BinaryData, 3)
}

func BenchmarkDataDogZstdCompressBinary1MB_Level7(b *testing.B) {
	benchmarkDataDogZstdCompress(b, SmallSize, BinaryData, 7)
}

func BenchmarkDataDogZstdCompressBinary10MB_Level1(b *testing.B) {
	benchmarkDataDogZstdCompress(b, MediumSize, BinaryData, 1)
}

func BenchmarkDataDogZstdCompressBinary10MB_Level3(b *testing.B) {
	benchmarkDataDogZstdCompress(b, MediumSize, BinaryData, 3)
}

func BenchmarkDataDogZstdCompressBinary10MB_Level7(b *testing.B) {
	benchmarkDataDogZstdCompress(b, MediumSize, BinaryData, 7)
}

// Benchmark DataDog zstd decompression
func BenchmarkDataDogZstdDecompress1MB_Level1(b *testing.B) {
	benchmarkDataDogZstdDecompress(b, SmallSize, TextData, 1)
}

func BenchmarkDataDogZstdDecompress1MB_Level3(b *testing.B) {
	benchmarkDataDogZstdDecompress(b, SmallSize, TextData, 3)
}

func BenchmarkDataDogZstdDecompress1MB_Level7(b *testing.B) {
	benchmarkDataDogZstdDecompress(b, SmallSize, TextData, 7)
}

func BenchmarkDataDogZstdDecompress10MB_Level1(b *testing.B) {
	benchmarkDataDogZstdDecompress(b, MediumSize, TextData, 1)
}

func BenchmarkDataDogZstdDecompress10MB_Level3(b *testing.B) {
	benchmarkDataDogZstdDecompress(b, MediumSize, TextData, 3)
}

func BenchmarkDataDogZstdDecompress10MB_Level7(b *testing.B) {
	benchmarkDataDogZstdDecompress(b, MediumSize, TextData, 7)
}

func BenchmarkDataDogZstdDecompress100MB_Level1(b *testing.B) {
	benchmarkDataDogZstdDecompress(b, LargeSize, TextData, 1)
}

func BenchmarkDataDogZstdDecompress100MB_Level3(b *testing.B) {
	benchmarkDataDogZstdDecompress(b, LargeSize, TextData, 3)
}

func BenchmarkDataDogZstdDecompress100MB_Level7(b *testing.B) {
	benchmarkDataDogZstdDecompress(b, LargeSize, TextData, 7)
}

func BenchmarkDataDogZstdDecompressBinary1MB_Level1(b *testing.B) {
	benchmarkDataDogZstdDecompress(b, SmallSize, BinaryData, 1)
}

func BenchmarkDataDogZstdDecompressBinary1MB_Level3(b *testing.B) {
	benchmarkDataDogZstdDecompress(b, SmallSize, BinaryData, 3)
}

func BenchmarkDataDogZstdDecompressBinary1MB_Level7(b *testing.B) {
	benchmarkDataDogZstdDecompress(b, SmallSize, BinaryData, 7)
}

func BenchmarkDataDogZstdDecompressBinary10MB_Level1(b *testing.B) {
	benchmarkDataDogZstdDecompress(b, MediumSize, BinaryData, 1)
}

func BenchmarkDataDogZstdDecompressBinary10MB_Level3(b *testing.B) {
	benchmarkDataDogZstdDecompress(b, MediumSize, BinaryData, 3)
}

func BenchmarkDataDogZstdDecompressBinary10MB_Level7(b *testing.B) {
	benchmarkDataDogZstdDecompress(b, MediumSize, BinaryData, 7)
}

// Helper function for benchmarking Klauspost compression
func benchmarkKlauspostZstdCompress(b *testing.B, size int, dataType string, level int) {
	data := GenerateTestData(size, dataType)

	// Map integer levels to klauspost EncoderLevel based on Klauspost's own documentation:
	// SpeedFastest: roughly equivalent to zstd level 1
	// SpeedDefault: roughly equivalent to zstd level 3
	// SpeedBetterCompression: roughly equivalent to zstd level 7-8
	// SpeedBestCompression: best available compression option
	var encoderLevel klauspost.EncoderLevel
	switch level {
	case 1:
		encoderLevel = klauspost.SpeedFastest // "roughly equivalent to the fastest Zstandard mode"
	case 3:
		encoderLevel = klauspost.SpeedDefault // "roughly equivalent to the default Zstandard mode (level 3)"
	case 7:
		encoderLevel = klauspost.SpeedBetterCompression // "about zstd level 7-8"
	case 19, 20:
		encoderLevel = klauspost.SpeedBestCompression // best compression
	default:
		encoderLevel = klauspost.SpeedDefault
	}

	enc, err := klauspost.NewWriter(nil, klauspost.WithEncoderLevel(encoderLevel))
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	b.SetBytes(int64(size))
	for i := 0; i < b.N; i++ {
		_ = enc.EncodeAll(data, nil)
	}
}

// Helper function for benchmarking Klauspost decompression
func benchmarkKlauspostZstdDecompress(b *testing.B, size int, dataType string, level int) {
	data := GenerateTestData(size, dataType)

	// Map integer levels to klauspost EncoderLevel based on Klauspost's own documentation:
	// SpeedFastest: roughly equivalent to zstd level 1
	// SpeedDefault: roughly equivalent to zstd level 3
	// SpeedBetterCompression: roughly equivalent to zstd level 7-8
	// SpeedBestCompression: best available compression option
	var encoderLevel klauspost.EncoderLevel
	switch level {
	case 1:
		encoderLevel = klauspost.SpeedFastest // "roughly equivalent to the fastest Zstandard mode"
	case 3:
		encoderLevel = klauspost.SpeedDefault // "roughly equivalent to the default Zstandard mode (level 3)"
	case 7:
		encoderLevel = klauspost.SpeedBetterCompression // "about zstd level 7-8"
	case 19, 20:
		encoderLevel = klauspost.SpeedBestCompression // best compression
	default:
		encoderLevel = klauspost.SpeedDefault
	}

	enc, err := klauspost.NewWriter(nil, klauspost.WithEncoderLevel(encoderLevel))
	if err != nil {
		b.Fatal(err)
	}

	compressed := enc.EncodeAll(data, nil)

	dec, err := klauspost.NewReader(nil)
	if err != nil {
		b.Fatal(err)
	}

	// Verify decompression
	decompressed, err := dec.DecodeAll(compressed, nil)
	if err != nil {
		b.Fatal(err)
	}
	if !bytes.Equal(data, decompressed) {
		b.Fatal("Decompressed data does not match original")
	}

	b.ResetTimer()
	b.SetBytes(int64(size))
	for i := 0; i < b.N; i++ {
		_, err := dec.DecodeAll(compressed, nil)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// Helper function for benchmarking DataDog compression
func benchmarkDataDogZstdCompress(b *testing.B, size int, dataType string, level int) {
	data := GenerateTestData(size, dataType)

	b.ResetTimer()
	b.SetBytes(int64(size))
	for i := 0; i < b.N; i++ {
		_, _ = datadog.CompressLevel(nil, data, level)
	}
}

// Helper function for benchmarking DataDog decompression
func benchmarkDataDogZstdDecompress(b *testing.B, size int, dataType string, level int) {
	data := GenerateTestData(size, dataType)
	compressed, _ := datadog.CompressLevel(nil, data, level)

	// Verify decompression
	decompressed, err := datadog.Decompress(nil, compressed)
	if err != nil {
		b.Fatal(err)
	}
	if !bytes.Equal(data, decompressed) {
		b.Fatal("Decompressed data does not match original")
	}

	b.ResetTimer()
	b.SetBytes(int64(size))
	for i := 0; i < b.N; i++ {
		_, err := datadog.Decompress(nil, compressed)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkZstdCompressionRatio measures and reports compression ratios
func BenchmarkZstdCompressionRatio(b *testing.B) {
	// This doesn't actually benchmark speed but uses the benchmark framework to report metrics
	b.Run("Klauspost-Zstd-1MB-Level1", func(b *testing.B) {
		measureZstdCompressionRatio(b, SmallSize, TextData, 1, "klauspost")
	})
	b.Run("Klauspost-Zstd-1MB-Level3", func(b *testing.B) {
		measureZstdCompressionRatio(b, SmallSize, TextData, 3, "klauspost")
	})
	b.Run("Klauspost-Zstd-1MB-Level7", func(b *testing.B) {
		measureZstdCompressionRatio(b, SmallSize, TextData, 7, "klauspost")
	})
	b.Run("Klauspost-Zstd-10MB-Level1", func(b *testing.B) {
		measureZstdCompressionRatio(b, MediumSize, TextData, 1, "klauspost")
	})
	b.Run("Klauspost-Zstd-10MB-Level3", func(b *testing.B) {
		measureZstdCompressionRatio(b, MediumSize, TextData, 3, "klauspost")
	})
	b.Run("Klauspost-Zstd-10MB-Level7", func(b *testing.B) {
		measureZstdCompressionRatio(b, MediumSize, TextData, 7, "klauspost")
	})

	b.Run("Klauspost-Zstd-100MB-Level1", func(b *testing.B) {
		measureZstdCompressionRatio(b, LargeSize, TextData, 1, "klauspost")
	})
	b.Run("Klauspost-Zstd-100MB-Level3", func(b *testing.B) {
		measureZstdCompressionRatio(b, LargeSize, TextData, 3, "klauspost")
	})
	b.Run("Klauspost-Zstd-100MB-Level7", func(b *testing.B) {
		measureZstdCompressionRatio(b, LargeSize, TextData, 7, "klauspost")
	})

	b.Run("DataDog-Zstd-1MB-Level1", func(b *testing.B) {
		measureZstdCompressionRatio(b, SmallSize, TextData, 1, "datadog")
	})
	b.Run("DataDog-Zstd-1MB-Level3", func(b *testing.B) {
		measureZstdCompressionRatio(b, SmallSize, TextData, 3, "datadog")
	})
	b.Run("DataDog-Zstd-1MB-Level7", func(b *testing.B) {
		measureZstdCompressionRatio(b, SmallSize, TextData, 7, "datadog")
	})
	b.Run("DataDog-Zstd-10MB-Level1", func(b *testing.B) {
		measureZstdCompressionRatio(b, MediumSize, TextData, 1, "datadog")
	})
	b.Run("DataDog-Zstd-10MB-Level3", func(b *testing.B) {
		measureZstdCompressionRatio(b, MediumSize, TextData, 3, "datadog")
	})
	b.Run("DataDog-Zstd-10MB-Level7", func(b *testing.B) {
		measureZstdCompressionRatio(b, MediumSize, TextData, 7, "datadog")
	})

	b.Run("DataDog-Zstd-100MB-Level1", func(b *testing.B) {
		measureZstdCompressionRatio(b, LargeSize, TextData, 1, "datadog")
	})
	b.Run("DataDog-Zstd-100MB-Level3", func(b *testing.B) {
		measureZstdCompressionRatio(b, LargeSize, TextData, 3, "datadog")
	})
	b.Run("DataDog-Zstd-100MB-Level7", func(b *testing.B) {
		measureZstdCompressionRatio(b, LargeSize, TextData, 7, "datadog")
	})

	// Binary data
	b.Run("Klauspost-Zstd-Binary-1MB-Level1", func(b *testing.B) {
		measureZstdCompressionRatio(b, SmallSize, BinaryData, 1, "klauspost")
	})
	b.Run("Klauspost-Zstd-Binary-1MB-Level7", func(b *testing.B) {
		measureZstdCompressionRatio(b, SmallSize, BinaryData, 7, "klauspost")
	})
	b.Run("DataDog-Zstd-Binary-1MB-Level1", func(b *testing.B) {
		measureZstdCompressionRatio(b, SmallSize, BinaryData, 1, "datadog")
	})
	b.Run("DataDog-Zstd-Binary-1MB-Level7", func(b *testing.B) {
		measureZstdCompressionRatio(b, SmallSize, BinaryData, 7, "datadog")
	})

	// Random data (should compress poorly)
	b.Run("Klauspost-Zstd-Random-1MB-Level7", func(b *testing.B) {
		measureZstdCompressionRatio(b, SmallSize, RandomData, 7, "klauspost")
	})
	b.Run("DataDog-Zstd-Random-1MB-Level7", func(b *testing.B) {
		measureZstdCompressionRatio(b, SmallSize, RandomData, 7, "datadog")
	})
}

func measureZstdCompressionRatio(b *testing.B, size int, dataType string, level int, implementation string) {
	b.Helper()
	data := GenerateTestData(size, dataType)
	var compressed []byte

	if implementation == "klauspost" {
		// Map integer levels to klauspost EncoderLevel based on Klauspost's own documentation
		var encoderLevel klauspost.EncoderLevel
		switch level {
		case 1:
			encoderLevel = klauspost.SpeedFastest
		case 3:
			encoderLevel = klauspost.SpeedDefault
		case 7:
			encoderLevel = klauspost.SpeedBetterCompression
		default:
			encoderLevel = klauspost.SpeedDefault
		}

		enc, err := klauspost.NewWriter(nil, klauspost.WithEncoderLevel(encoderLevel))
		if err != nil {
			b.Fatal(err)
		}
		compressed = enc.EncodeAll(data, nil)
	} else {
		var err error
		compressed, err = datadog.CompressLevel(nil, data, level)
		if err != nil {
			b.Fatal(err)
		}
	}

	ratio := float64(len(data)) / float64(len(compressed))
	b.ReportMetric(ratio, "ratio")
	b.ReportMetric(0, "ns/op") // This benchmark doesn't measure time
}
