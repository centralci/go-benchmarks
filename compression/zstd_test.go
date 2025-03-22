package zstd_benchmark

import (
	"bytes"
	"testing"

	datadog "github.com/DataDog/zstd"
	"github.com/brianvoe/gofakeit/v7"
	klauspost "github.com/klauspost/compress/zstd"
)

// Test data sizes
const (
	smallSize  = 1 << 20   // 1 MB
	mediumSize = 10 << 20  // 10 MB
	largeSize  = 100 << 20 // 100 MB
)

// generateTestData creates sample data of specified type and size using gofakeit
func generateTestData(size int, dataType string) []byte {
	// Initialize gofakeit with a fixed seed for reproducible results
	gofakeit.Seed(42)

	data := make([]byte, 0, size)

	switch dataType {
	case "random":
		// Generate random bytes until we reach the desired size
		for len(data) < size {
			data = append(data, byte(gofakeit.IntRange(0, 255)))
		}
	case "text":
		// Generate realistic text data using gofakeit
		for len(data) < size {
			var chunk []byte
			switch gofakeit.IntRange(0, 3) {
			case 0:
				// Generate a paragraph
				chunk = []byte(gofakeit.Paragraph(1, 5, 20, " "))
			case 1:
				// Generate a sentence
				chunk = []byte(gofakeit.Sentence(gofakeit.IntRange(3, 10)))
			case 2:
				// Generate JSON-like data
				jsonData, _ := gofakeit.JSON(nil)
				chunk = jsonData
			case 3:
				// Generate a word
				chunk = []byte(gofakeit.Word())
			}

			// Add a space or newline between chunks
			if len(data) > 0 {
				if gofakeit.Bool() {
					chunk = append([]byte{' '}, chunk...)
				} else {
					chunk = append([]byte{'\n'}, chunk...)
				}
			}

			// Add the chunk to our data
			data = append(data, chunk...)
		}
	case "binary":
		// Create binary data mixing patterns and random bytes
		for len(data) < size {
			if len(data)%1024 < 512 {
				// Pattern-based section
				patternLength := gofakeit.IntRange(64, 256)
				pattern := make([]byte, patternLength)
				for i := range pattern {
					pattern[i] = byte(i % 256)
				}

				// Repeat pattern until we fill a section
				sectionSize := gofakeit.IntRange(512, 4096)
				for i := 0; i < sectionSize && len(data) < size; i++ {
					data = append(data, pattern[i%len(pattern)])
				}
			} else {
				// Random binary section
				sectionSize := gofakeit.IntRange(512, 4096)
				for i := 0; i < sectionSize && len(data) < size; i++ {
					data = append(data, byte(gofakeit.IntRange(0, 255)))
				}
			}
		}
	}

	// Trim to exact size
	if len(data) > size {
		data = data[:size]
	}

	return data
}

// Benchmark Klauspost zstd at different levels
func BenchmarkKlauspostCompress1MB_Level1(b *testing.B) {
	benchmarkKlauspostCompress(b, smallSize, "text", 1)
}

func BenchmarkKlauspostCompress1MB_Level3(b *testing.B) {
	benchmarkKlauspostCompress(b, smallSize, "text", 3)
}

func BenchmarkKlauspostCompress1MB_Level7(b *testing.B) {
	benchmarkKlauspostCompress(b, smallSize, "text", 7)
}

func BenchmarkKlauspostCompress10MB_Level1(b *testing.B) {
	benchmarkKlauspostCompress(b, mediumSize, "text", 1)
}

func BenchmarkKlauspostCompress10MB_Level3(b *testing.B) {
	benchmarkKlauspostCompress(b, mediumSize, "text", 3)
}

func BenchmarkKlauspostCompress10MB_Level7(b *testing.B) {
	benchmarkKlauspostCompress(b, mediumSize, "text", 7)
}

func BenchmarkKlauspostCompress100MB_Level1(b *testing.B) {
	benchmarkKlauspostCompress(b, largeSize, "text", 1)
}

func BenchmarkKlauspostCompress100MB_Level3(b *testing.B) {
	benchmarkKlauspostCompress(b, largeSize, "text", 3)
}

func BenchmarkKlauspostCompress100MB_Level7(b *testing.B) {
	benchmarkKlauspostCompress(b, largeSize, "text", 7)
}

func BenchmarkKlauspostCompressBinary1MB_Level1(b *testing.B) {
	benchmarkKlauspostCompress(b, smallSize, "binary", 1)
}

func BenchmarkKlauspostCompressBinary1MB_Level3(b *testing.B) {
	benchmarkKlauspostCompress(b, smallSize, "binary", 3)
}

func BenchmarkKlauspostCompressBinary1MB_Level7(b *testing.B) {
	benchmarkKlauspostCompress(b, smallSize, "binary", 7)
}

func BenchmarkKlauspostCompressBinary10MB_Level1(b *testing.B) {
	benchmarkKlauspostCompress(b, mediumSize, "binary", 1)
}

func BenchmarkKlauspostCompressBinary10MB_Level3(b *testing.B) {
	benchmarkKlauspostCompress(b, mediumSize, "binary", 3)
}

func BenchmarkKlauspostCompressBinary10MB_Level7(b *testing.B) {
	benchmarkKlauspostCompress(b, mediumSize, "binary", 7)
}

// Benchmark Klauspost zstd decompression
func BenchmarkKlauspostDecompress1MB_Level1(b *testing.B) {
	benchmarkKlauspostDecompress(b, smallSize, "text", 1)
}

func BenchmarkKlauspostDecompress1MB_Level3(b *testing.B) {
	benchmarkKlauspostDecompress(b, smallSize, "text", 3)
}

func BenchmarkKlauspostDecompress1MB_Level7(b *testing.B) {
	benchmarkKlauspostDecompress(b, smallSize, "text", 7)
}

func BenchmarkKlauspostDecompress10MB_Level1(b *testing.B) {
	benchmarkKlauspostDecompress(b, mediumSize, "text", 1)
}

func BenchmarkKlauspostDecompress10MB_Level3(b *testing.B) {
	benchmarkKlauspostDecompress(b, mediumSize, "text", 3)
}

func BenchmarkKlauspostDecompress10MB_Level7(b *testing.B) {
	benchmarkKlauspostDecompress(b, mediumSize, "text", 7)
}

func BenchmarkKlauspostDecompress100MB_Level1(b *testing.B) {
	benchmarkKlauspostDecompress(b, largeSize, "text", 1)
}

func BenchmarkKlauspostDecompress100MB_Level3(b *testing.B) {
	benchmarkKlauspostDecompress(b, largeSize, "text", 3)
}

func BenchmarkKlauspostDecompress100MB_Level7(b *testing.B) {
	benchmarkKlauspostDecompress(b, largeSize, "text", 7)
}

func BenchmarkKlauspostDecompressBinary1MB_Level1(b *testing.B) {
	benchmarkKlauspostDecompress(b, smallSize, "binary", 1)
}

func BenchmarkKlauspostDecompressBinary1MB_Level3(b *testing.B) {
	benchmarkKlauspostDecompress(b, smallSize, "binary", 3)
}

func BenchmarkKlauspostDecompressBinary1MB_Level7(b *testing.B) {
	benchmarkKlauspostDecompress(b, smallSize, "binary", 7)
}

func BenchmarkKlauspostDecompressBinary10MB_Level1(b *testing.B) {
	benchmarkKlauspostDecompress(b, mediumSize, "binary", 1)
}

func BenchmarkKlauspostDecompressBinary10MB_Level3(b *testing.B) {
	benchmarkKlauspostDecompress(b, mediumSize, "binary", 3)
}

func BenchmarkKlauspostDecompressBinary10MB_Level7(b *testing.B) {
	benchmarkKlauspostDecompress(b, mediumSize, "binary", 7)
}

// Benchmark DataDog zstd at different levels
func BenchmarkDataDogCompress1MB_Level1(b *testing.B) {
	benchmarkDataDogCompress(b, smallSize, "text", 1)
}

func BenchmarkDataDogCompress1MB_Level3(b *testing.B) {
	benchmarkDataDogCompress(b, smallSize, "text", 3)
}

func BenchmarkDataDogCompress1MB_Level7(b *testing.B) {
	benchmarkDataDogCompress(b, smallSize, "text", 7)
}

func BenchmarkDataDogCompress10MB_Level1(b *testing.B) {
	benchmarkDataDogCompress(b, mediumSize, "text", 1)
}

func BenchmarkDataDogCompress10MB_Level3(b *testing.B) {
	benchmarkDataDogCompress(b, mediumSize, "text", 3)
}

func BenchmarkDataDogCompress10MB_Level7(b *testing.B) {
	benchmarkDataDogCompress(b, mediumSize, "text", 7)
}

func BenchmarkDataDogCompress100MB_Level1(b *testing.B) {
	benchmarkDataDogCompress(b, largeSize, "text", 1)
}

func BenchmarkDataDogCompress100MB_Level3(b *testing.B) {
	benchmarkDataDogCompress(b, largeSize, "text", 3)
}

func BenchmarkDataDogCompress100MB_Level7(b *testing.B) {
	benchmarkDataDogCompress(b, largeSize, "text", 7)
}

func BenchmarkDataDogCompressBinary1MB_Level1(b *testing.B) {
	benchmarkDataDogCompress(b, smallSize, "binary", 1)
}

func BenchmarkDataDogCompressBinary1MB_Level3(b *testing.B) {
	benchmarkDataDogCompress(b, smallSize, "binary", 3)
}

func BenchmarkDataDogCompressBinary1MB_Level7(b *testing.B) {
	benchmarkDataDogCompress(b, smallSize, "binary", 7)
}

func BenchmarkDataDogCompressBinary10MB_Level1(b *testing.B) {
	benchmarkDataDogCompress(b, mediumSize, "binary", 1)
}

func BenchmarkDataDogCompressBinary10MB_Level3(b *testing.B) {
	benchmarkDataDogCompress(b, mediumSize, "binary", 3)
}

func BenchmarkDataDogCompressBinary10MB_Level7(b *testing.B) {
	benchmarkDataDogCompress(b, mediumSize, "binary", 7)
}

// Benchmark DataDog zstd decompression
func BenchmarkDataDogDecompress1MB_Level1(b *testing.B) {
	benchmarkDataDogDecompress(b, smallSize, "text", 1)
}

func BenchmarkDataDogDecompress1MB_Level3(b *testing.B) {
	benchmarkDataDogDecompress(b, smallSize, "text", 3)
}

func BenchmarkDataDogDecompress1MB_Level7(b *testing.B) {
	benchmarkDataDogDecompress(b, smallSize, "text", 7)
}

func BenchmarkDataDogDecompress10MB_Level1(b *testing.B) {
	benchmarkDataDogDecompress(b, mediumSize, "text", 1)
}

func BenchmarkDataDogDecompress10MB_Level3(b *testing.B) {
	benchmarkDataDogDecompress(b, mediumSize, "text", 3)
}

func BenchmarkDataDogDecompress10MB_Level7(b *testing.B) {
	benchmarkDataDogDecompress(b, mediumSize, "text", 7)
}

func BenchmarkDataDogDecompress100MB_Level1(b *testing.B) {
	benchmarkDataDogDecompress(b, largeSize, "text", 1)
}

func BenchmarkDataDogDecompress100MB_Level3(b *testing.B) {
	benchmarkDataDogDecompress(b, largeSize, "text", 3)
}

func BenchmarkDataDogDecompress100MB_Level7(b *testing.B) {
	benchmarkDataDogDecompress(b, largeSize, "text", 7)
}

func BenchmarkDataDogDecompressBinary1MB_Level1(b *testing.B) {
	benchmarkDataDogDecompress(b, smallSize, "binary", 1)
}

func BenchmarkDataDogDecompressBinary1MB_Level3(b *testing.B) {
	benchmarkDataDogDecompress(b, smallSize, "binary", 3)
}

func BenchmarkDataDogDecompressBinary1MB_Level7(b *testing.B) {
	benchmarkDataDogDecompress(b, smallSize, "binary", 7)
}

func BenchmarkDataDogDecompressBinary10MB_Level1(b *testing.B) {
	benchmarkDataDogDecompress(b, mediumSize, "binary", 1)
}

func BenchmarkDataDogDecompressBinary10MB_Level3(b *testing.B) {
	benchmarkDataDogDecompress(b, mediumSize, "binary", 3)
}

func BenchmarkDataDogDecompressBinary10MB_Level7(b *testing.B) {
	benchmarkDataDogDecompress(b, mediumSize, "binary", 7)
}

// Helper function for benchmarking Klauspost compression
func benchmarkKlauspostCompress(b *testing.B, size int, dataType string, level int) {
	data := generateTestData(size, dataType)

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
func benchmarkKlauspostDecompress(b *testing.B, size int, dataType string, level int) {
	data := generateTestData(size, dataType)

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
func benchmarkDataDogCompress(b *testing.B, size int, dataType string, level int) {
	data := generateTestData(size, dataType)

	b.ResetTimer()
	b.SetBytes(int64(size))
	for i := 0; i < b.N; i++ {
		_, _ = datadog.CompressLevel(nil, data, level)
	}
}

// Helper function for benchmarking DataDog decompression
func benchmarkDataDogDecompress(b *testing.B, size int, dataType string, level int) {
	data := generateTestData(size, dataType)
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

// BenchmarkCompressionRatio measures and reports compression ratios
func BenchmarkCompressionRatio(b *testing.B) {
	// This doesn't actually benchmark speed but uses the benchmark framework to report metrics
	b.Run("Klauspost-1MB-Level1", func(b *testing.B) {
		measureCompressionRatio(b, smallSize, "text", 1, "klauspost")
	})
	b.Run("Klauspost-1MB-Level3", func(b *testing.B) {
		measureCompressionRatio(b, smallSize, "text", 3, "klauspost")
	})
	b.Run("Klauspost-1MB-Level7", func(b *testing.B) {
		measureCompressionRatio(b, smallSize, "text", 7, "klauspost")
	})
	b.Run("Klauspost-10MB-Level1", func(b *testing.B) {
		measureCompressionRatio(b, mediumSize, "text", 1, "klauspost")
	})
	b.Run("Klauspost-10MB-Level3", func(b *testing.B) {
		measureCompressionRatio(b, mediumSize, "text", 3, "klauspost")
	})
	b.Run("Klauspost-10MB-Level7", func(b *testing.B) {
		measureCompressionRatio(b, mediumSize, "text", 7, "klauspost")
	})

	b.Run("Klauspost-100MB-Level1", func(b *testing.B) {
		measureCompressionRatio(b, largeSize, "text", 1, "klauspost")
	})
	b.Run("Klauspost-100MB-Level3", func(b *testing.B) {
		measureCompressionRatio(b, largeSize, "text", 3, "klauspost")
	})
	b.Run("Klauspost-100MB-Level7", func(b *testing.B) {
		measureCompressionRatio(b, largeSize, "text", 7, "klauspost")
	})

	b.Run("DataDog-1MB-Level1", func(b *testing.B) {
		measureCompressionRatio(b, smallSize, "text", 1, "datadog")
	})
	b.Run("DataDog-1MB-Level3", func(b *testing.B) {
		measureCompressionRatio(b, smallSize, "text", 3, "datadog")
	})
	b.Run("DataDog-1MB-Level7", func(b *testing.B) {
		measureCompressionRatio(b, smallSize, "text", 7, "datadog")
	})
	b.Run("DataDog-10MB-Level1", func(b *testing.B) {
		measureCompressionRatio(b, mediumSize, "text", 1, "datadog")
	})
	b.Run("DataDog-10MB-Level3", func(b *testing.B) {
		measureCompressionRatio(b, mediumSize, "text", 3, "datadog")
	})
	b.Run("DataDog-10MB-Level7", func(b *testing.B) {
		measureCompressionRatio(b, mediumSize, "text", 7, "datadog")
	})

	b.Run("DataDog-100MB-Level1", func(b *testing.B) {
		measureCompressionRatio(b, largeSize, "text", 1, "datadog")
	})
	b.Run("DataDog-100MB-Level3", func(b *testing.B) {
		measureCompressionRatio(b, largeSize, "text", 3, "datadog")
	})
	b.Run("DataDog-100MB-Level7", func(b *testing.B) {
		measureCompressionRatio(b, largeSize, "text", 7, "datadog")
	})

	// Binary data
	b.Run("Klauspost-Binary-1MB-Level1", func(b *testing.B) {
		measureCompressionRatio(b, smallSize, "binary", 1, "klauspost")
	})
	b.Run("Klauspost-Binary-1MB-Level7", func(b *testing.B) {
		measureCompressionRatio(b, smallSize, "binary", 7, "klauspost")
	})
	b.Run("DataDog-Binary-1MB-Level1", func(b *testing.B) {
		measureCompressionRatio(b, smallSize, "binary", 1, "datadog")
	})
	b.Run("DataDog-Binary-1MB-Level7", func(b *testing.B) {
		measureCompressionRatio(b, smallSize, "binary", 7, "datadog")
	})
}

func measureCompressionRatio(b *testing.B, size int, dataType string, level int, implementation string) {
	b.Helper()
	data := generateTestData(size, dataType)
	var compressed []byte

	if implementation == "klauspost" {
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
