package compression

import (
	"bytes"
	"compress/gzip"
	"io"
	"testing"

	klauspost "github.com/klauspost/compress/gzip"
)

// Benchmark Klauspost gzip at different levels
func BenchmarkKlauspostGzipCompress1MB_Level1(b *testing.B) {
	benchmarkKlauspostGzipCompress(b, SmallSize, TextData, 1)
}

func BenchmarkKlauspostGzipCompress1MB_Level3(b *testing.B) {
	benchmarkKlauspostGzipCompress(b, SmallSize, TextData, 3)
}

func BenchmarkKlauspostGzipCompress1MB_Level9(b *testing.B) {
	benchmarkKlauspostGzipCompress(b, SmallSize, TextData, 9)
}

func BenchmarkKlauspostGzipCompress10MB_Level1(b *testing.B) {
	benchmarkKlauspostGzipCompress(b, MediumSize, TextData, 1)
}

func BenchmarkKlauspostGzipCompress10MB_Level3(b *testing.B) {
	benchmarkKlauspostGzipCompress(b, MediumSize, TextData, 3)
}

func BenchmarkKlauspostGzipCompress10MB_Level9(b *testing.B) {
	benchmarkKlauspostGzipCompress(b, MediumSize, TextData, 9)
}

func BenchmarkKlauspostGzipCompress100MB_Level1(b *testing.B) {
	benchmarkKlauspostGzipCompress(b, LargeSize, TextData, 1)
}

func BenchmarkKlauspostGzipCompress100MB_Level3(b *testing.B) {
	benchmarkKlauspostGzipCompress(b, LargeSize, TextData, 3)
}

func BenchmarkKlauspostGzipCompress100MB_Level9(b *testing.B) {
	benchmarkKlauspostGzipCompress(b, LargeSize, TextData, 9)
}

func BenchmarkKlauspostGzipCompressBinary1MB_Level1(b *testing.B) {
	benchmarkKlauspostGzipCompress(b, SmallSize, BinaryData, 1)
}

// Added Level 3 binary compression
func BenchmarkKlauspostGzipCompressBinary1MB_Level3(b *testing.B) {
	benchmarkKlauspostGzipCompress(b, SmallSize, BinaryData, 3)
}

func BenchmarkKlauspostGzipCompressBinary1MB_Level9(b *testing.B) {
	benchmarkKlauspostGzipCompress(b, SmallSize, BinaryData, 9)
}

func BenchmarkKlauspostGzipCompressBinary10MB_Level1(b *testing.B) {
	benchmarkKlauspostGzipCompress(b, MediumSize, BinaryData, 1)
}

// Added Level 3 binary compression
func BenchmarkKlauspostGzipCompressBinary10MB_Level3(b *testing.B) {
	benchmarkKlauspostGzipCompress(b, MediumSize, BinaryData, 3)
}

func BenchmarkKlauspostGzipCompressBinary10MB_Level9(b *testing.B) {
	benchmarkKlauspostGzipCompress(b, MediumSize, BinaryData, 9)
}

// Benchmark Klauspost gzip decompression
func BenchmarkKlauspostGzipDecompress1MB_Level1(b *testing.B) {
	benchmarkKlauspostGzipDecompress(b, SmallSize, TextData, 1)
}

func BenchmarkKlauspostGzipDecompress1MB_Level3(b *testing.B) {
	benchmarkKlauspostGzipDecompress(b, SmallSize, TextData, 3)
}

func BenchmarkKlauspostGzipDecompress1MB_Level9(b *testing.B) {
	benchmarkKlauspostGzipDecompress(b, SmallSize, TextData, 9)
}

func BenchmarkKlauspostGzipDecompress10MB_Level1(b *testing.B) {
	benchmarkKlauspostGzipDecompress(b, MediumSize, TextData, 1)
}

func BenchmarkKlauspostGzipDecompress10MB_Level3(b *testing.B) {
	benchmarkKlauspostGzipDecompress(b, MediumSize, TextData, 3)
}

func BenchmarkKlauspostGzipDecompress10MB_Level9(b *testing.B) {
	benchmarkKlauspostGzipDecompress(b, MediumSize, TextData, 9)
}

func BenchmarkKlauspostGzipDecompress100MB_Level1(b *testing.B) {
	benchmarkKlauspostGzipDecompress(b, LargeSize, TextData, 1)
}

func BenchmarkKlauspostGzipDecompress100MB_Level3(b *testing.B) {
	benchmarkKlauspostGzipDecompress(b, LargeSize, TextData, 3)
}

func BenchmarkKlauspostGzipDecompress100MB_Level9(b *testing.B) {
	benchmarkKlauspostGzipDecompress(b, LargeSize, TextData, 9)
}

func BenchmarkKlauspostGzipDecompressBinary1MB_Level1(b *testing.B) {
	benchmarkKlauspostGzipDecompress(b, SmallSize, BinaryData, 1)
}

// Added Level 3 binary decompression
func BenchmarkKlauspostGzipDecompressBinary1MB_Level3(b *testing.B) {
	benchmarkKlauspostGzipDecompress(b, SmallSize, BinaryData, 3)
}

func BenchmarkKlauspostGzipDecompressBinary1MB_Level9(b *testing.B) {
	benchmarkKlauspostGzipDecompress(b, SmallSize, BinaryData, 9)
}

func BenchmarkKlauspostGzipDecompressBinary10MB_Level1(b *testing.B) {
	benchmarkKlauspostGzipDecompress(b, MediumSize, BinaryData, 1)
}

// Added Level 3 binary decompression
func BenchmarkKlauspostGzipDecompressBinary10MB_Level3(b *testing.B) {
	benchmarkKlauspostGzipDecompress(b, MediumSize, BinaryData, 3)
}

func BenchmarkKlauspostGzipDecompressBinary10MB_Level9(b *testing.B) {
	benchmarkKlauspostGzipDecompress(b, MediumSize, BinaryData, 9)
}

// Benchmark standard library gzip at different levels
func BenchmarkStdlibGzipCompress1MB_Level1(b *testing.B) {
	benchmarkStdlibGzipCompress(b, SmallSize, TextData, 1)
}

func BenchmarkStdlibGzipCompress1MB_Level3(b *testing.B) {
	benchmarkStdlibGzipCompress(b, SmallSize, TextData, 3)
}

func BenchmarkStdlibGzipCompress1MB_Level9(b *testing.B) {
	benchmarkStdlibGzipCompress(b, SmallSize, TextData, 9)
}

func BenchmarkStdlibGzipCompress10MB_Level1(b *testing.B) {
	benchmarkStdlibGzipCompress(b, MediumSize, TextData, 1)
}

func BenchmarkStdlibGzipCompress10MB_Level3(b *testing.B) {
	benchmarkStdlibGzipCompress(b, MediumSize, TextData, 3)
}

func BenchmarkStdlibGzipCompress10MB_Level9(b *testing.B) {
	benchmarkStdlibGzipCompress(b, MediumSize, TextData, 9)
}

func BenchmarkStdlibGzipCompress100MB_Level1(b *testing.B) {
	benchmarkStdlibGzipCompress(b, LargeSize, TextData, 1)
}

func BenchmarkStdlibGzipCompress100MB_Level3(b *testing.B) {
	benchmarkStdlibGzipCompress(b, LargeSize, TextData, 3)
}

func BenchmarkStdlibGzipCompress100MB_Level9(b *testing.B) {
	benchmarkStdlibGzipCompress(b, LargeSize, TextData, 9)
}

func BenchmarkStdlibGzipCompressBinary1MB_Level1(b *testing.B) {
	benchmarkStdlibGzipCompress(b, SmallSize, BinaryData, 1)
}

// Added Level 3 binary compression
func BenchmarkStdlibGzipCompressBinary1MB_Level3(b *testing.B) {
	benchmarkStdlibGzipCompress(b, SmallSize, BinaryData, 3)
}

func BenchmarkStdlibGzipCompressBinary1MB_Level9(b *testing.B) {
	benchmarkStdlibGzipCompress(b, SmallSize, BinaryData, 9)
}

func BenchmarkStdlibGzipCompressBinary10MB_Level1(b *testing.B) {
	benchmarkStdlibGzipCompress(b, MediumSize, BinaryData, 1)
}

// Added Level 3 binary compression
func BenchmarkStdlibGzipCompressBinary10MB_Level3(b *testing.B) {
	benchmarkStdlibGzipCompress(b, MediumSize, BinaryData, 3)
}

func BenchmarkStdlibGzipCompressBinary10MB_Level9(b *testing.B) {
	benchmarkStdlibGzipCompress(b, MediumSize, BinaryData, 9)
}

// Benchmark standard library gzip decompression
func BenchmarkStdlibGzipDecompress1MB_Level1(b *testing.B) {
	benchmarkStdlibGzipDecompress(b, SmallSize, TextData, 1)
}

func BenchmarkStdlibGzipDecompress1MB_Level3(b *testing.B) {
	benchmarkStdlibGzipDecompress(b, SmallSize, TextData, 3)
}

func BenchmarkStdlibGzipDecompress1MB_Level9(b *testing.B) {
	benchmarkStdlibGzipDecompress(b, SmallSize, TextData, 9)
}

func BenchmarkStdlibGzipDecompress10MB_Level1(b *testing.B) {
	benchmarkStdlibGzipDecompress(b, MediumSize, TextData, 1)
}

func BenchmarkStdlibGzipDecompress10MB_Level3(b *testing.B) {
	benchmarkStdlibGzipDecompress(b, MediumSize, TextData, 3)
}

func BenchmarkStdlibGzipDecompress10MB_Level9(b *testing.B) {
	benchmarkStdlibGzipDecompress(b, MediumSize, TextData, 9)
}

func BenchmarkStdlibGzipDecompress100MB_Level1(b *testing.B) {
	benchmarkStdlibGzipDecompress(b, LargeSize, TextData, 1)
}

func BenchmarkStdlibGzipDecompress100MB_Level3(b *testing.B) {
	benchmarkStdlibGzipDecompress(b, LargeSize, TextData, 3)
}

func BenchmarkStdlibGzipDecompress100MB_Level9(b *testing.B) {
	benchmarkStdlibGzipDecompress(b, LargeSize, TextData, 9)
}

func BenchmarkStdlibGzipDecompressBinary1MB_Level1(b *testing.B) {
	benchmarkStdlibGzipDecompress(b, SmallSize, BinaryData, 1)
}

// Added Level 3 binary decompression
func BenchmarkStdlibGzipDecompressBinary1MB_Level3(b *testing.B) {
	benchmarkStdlibGzipDecompress(b, SmallSize, BinaryData, 3)
}

func BenchmarkStdlibGzipDecompressBinary1MB_Level9(b *testing.B) {
	benchmarkStdlibGzipDecompress(b, SmallSize, BinaryData, 9)
}

func BenchmarkStdlibGzipDecompressBinary10MB_Level1(b *testing.B) {
	benchmarkStdlibGzipDecompress(b, MediumSize, BinaryData, 1)
}

// Added Level 3 binary decompression
func BenchmarkStdlibGzipDecompressBinary10MB_Level3(b *testing.B) {
	benchmarkStdlibGzipDecompress(b, MediumSize, BinaryData, 3)
}

func BenchmarkStdlibGzipDecompressBinary10MB_Level9(b *testing.B) {
	benchmarkStdlibGzipDecompress(b, MediumSize, BinaryData, 9)
}

// Helper function for benchmarking Klauspost gzip compression
func benchmarkKlauspostGzipCompress(b *testing.B, size int, dataType string, level int) {
	data := GenerateTestData(size, dataType)

	b.ResetTimer()
	b.SetBytes(int64(size))
	for i := 0; i < b.N; i++ {
		var buf bytes.Buffer
		w, err := klauspost.NewWriterLevel(&buf, level)
		if err != nil {
			b.Fatal(err)
		}
		_, err = w.Write(data)
		if err != nil {
			b.Fatal(err)
		}
		err = w.Close()
		if err != nil {
			b.Fatal(err)
		}
	}
}

// Helper function for benchmarking Klauspost gzip decompression
func benchmarkKlauspostGzipDecompress(b *testing.B, size int, dataType string, level int) {
	data := GenerateTestData(size, dataType)

	// Compress data
	var buf bytes.Buffer
	w, err := klauspost.NewWriterLevel(&buf, level)
	if err != nil {
		b.Fatal(err)
	}
	_, err = w.Write(data)
	if err != nil {
		b.Fatal(err)
	}
	err = w.Close()
	if err != nil {
		b.Fatal(err)
	}
	compressed := buf.Bytes()

	// Verify decompression works
	r, err := klauspost.NewReader(bytes.NewBuffer(compressed))
	if err != nil {
		b.Fatal(err)
	}
	decompressed, err := io.ReadAll(r)
	if err != nil {
		b.Fatal(err)
	}
	if !bytes.Equal(data, decompressed) {
		b.Fatal("Decompressed data does not match original")
	}

	b.ResetTimer()
	b.SetBytes(int64(size))
	for i := 0; i < b.N; i++ {
		r, err := klauspost.NewReader(bytes.NewBuffer(compressed))
		if err != nil {
			b.Fatal(err)
		}
		_, err = io.ReadAll(r)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// Helper function for benchmarking standard library gzip compression
func benchmarkStdlibGzipCompress(b *testing.B, size int, dataType string, level int) {
	data := GenerateTestData(size, dataType)

	b.ResetTimer()
	b.SetBytes(int64(size))
	for i := 0; i < b.N; i++ {
		var buf bytes.Buffer
		w, err := gzip.NewWriterLevel(&buf, level)
		if err != nil {
			b.Fatal(err)
		}
		_, err = w.Write(data)
		if err != nil {
			b.Fatal(err)
		}
		err = w.Close()
		if err != nil {
			b.Fatal(err)
		}
	}
}

// Helper function for benchmarking standard library gzip decompression
func benchmarkStdlibGzipDecompress(b *testing.B, size int, dataType string, level int) {
	data := GenerateTestData(size, dataType)

	// Compress data
	var buf bytes.Buffer
	w, err := gzip.NewWriterLevel(&buf, level)
	if err != nil {
		b.Fatal(err)
	}
	_, err = w.Write(data)
	if err != nil {
		b.Fatal(err)
	}
	err = w.Close()
	if err != nil {
		b.Fatal(err)
	}
	compressed := buf.Bytes()

	// Verify decompression works
	r, err := gzip.NewReader(bytes.NewBuffer(compressed))
	if err != nil {
		b.Fatal(err)
	}
	decompressed, err := io.ReadAll(r)
	if err != nil {
		b.Fatal(err)
	}
	if !bytes.Equal(data, decompressed) {
		b.Fatal("Decompressed data does not match original")
	}

	b.ResetTimer()
	b.SetBytes(int64(size))
	for i := 0; i < b.N; i++ {
		r, err := gzip.NewReader(bytes.NewBuffer(compressed))
		if err != nil {
			b.Fatal(err)
		}
		_, err = io.ReadAll(r)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkGzipCompressionRatio measures and reports compression ratios
func BenchmarkGzipCompressionRatio(b *testing.B) {
	// This doesn't actually benchmark speed but uses the benchmark framework to report metrics
	b.Run("Klauspost-Gzip-1MB-Level1", func(b *testing.B) {
		measureGzipCompressionRatio(b, SmallSize, TextData, 1, "klauspost")
	})
	b.Run("Klauspost-Gzip-1MB-Level3", func(b *testing.B) {
		measureGzipCompressionRatio(b, SmallSize, TextData, 3, "klauspost")
	})
	b.Run("Klauspost-Gzip-1MB-Level9", func(b *testing.B) {
		measureGzipCompressionRatio(b, SmallSize, TextData, 9, "klauspost")
	})
	b.Run("Klauspost-Gzip-10MB-Level1", func(b *testing.B) {
		measureGzipCompressionRatio(b, MediumSize, TextData, 1, "klauspost")
	})
	b.Run("Klauspost-Gzip-10MB-Level3", func(b *testing.B) {
		measureGzipCompressionRatio(b, MediumSize, TextData, 3, "klauspost")
	})
	b.Run("Klauspost-Gzip-10MB-Level9", func(b *testing.B) {
		measureGzipCompressionRatio(b, MediumSize, TextData, 9, "klauspost")
	})

	b.Run("Klauspost-Gzip-100MB-Level1", func(b *testing.B) {
		measureGzipCompressionRatio(b, LargeSize, TextData, 1, "klauspost")
	})
	b.Run("Klauspost-Gzip-100MB-Level3", func(b *testing.B) {
		measureGzipCompressionRatio(b, LargeSize, TextData, 3, "klauspost")
	})
	b.Run("Klauspost-Gzip-100MB-Level9", func(b *testing.B) {
		measureGzipCompressionRatio(b, LargeSize, TextData, 9, "klauspost")
	})

	b.Run("Stdlib-Gzip-1MB-Level1", func(b *testing.B) {
		measureGzipCompressionRatio(b, SmallSize, TextData, 1, "stdlib")
	})
	b.Run("Stdlib-Gzip-1MB-Level3", func(b *testing.B) {
		measureGzipCompressionRatio(b, SmallSize, TextData, 3, "stdlib")
	})
	b.Run("Stdlib-Gzip-1MB-Level9", func(b *testing.B) {
		measureGzipCompressionRatio(b, SmallSize, TextData, 9, "stdlib")
	})
	b.Run("Stdlib-Gzip-10MB-Level1", func(b *testing.B) {
		measureGzipCompressionRatio(b, MediumSize, TextData, 1, "stdlib")
	})
	b.Run("Stdlib-Gzip-10MB-Level3", func(b *testing.B) {
		measureGzipCompressionRatio(b, MediumSize, TextData, 3, "stdlib")
	})
	b.Run("Stdlib-Gzip-10MB-Level9", func(b *testing.B) {
		measureGzipCompressionRatio(b, MediumSize, TextData, 9, "stdlib")
	})

	b.Run("Stdlib-Gzip-100MB-Level1", func(b *testing.B) {
		measureGzipCompressionRatio(b, LargeSize, TextData, 1, "stdlib")
	})
	b.Run("Stdlib-Gzip-100MB-Level3", func(b *testing.B) {
		measureGzipCompressionRatio(b, LargeSize, TextData, 3, "stdlib")
	})
	b.Run("Stdlib-Gzip-100MB-Level9", func(b *testing.B) {
		measureGzipCompressionRatio(b, LargeSize, TextData, 9, "stdlib")
	})

	// Binary data
	b.Run("Klauspost-Gzip-Binary-1MB-Level1", func(b *testing.B) {
		measureGzipCompressionRatio(b, SmallSize, BinaryData, 1, "klauspost")
	})
	// Added Level 3 binary compression ratio test
	b.Run("Klauspost-Gzip-Binary-1MB-Level3", func(b *testing.B) {
		measureGzipCompressionRatio(b, SmallSize, BinaryData, 3, "klauspost")
	})
	b.Run("Klauspost-Gzip-Binary-1MB-Level9", func(b *testing.B) {
		measureGzipCompressionRatio(b, SmallSize, BinaryData, 9, "klauspost")
	})
	// Added Level 3 binary compression ratio tests for 10MB
	b.Run("Klauspost-Gzip-Binary-10MB-Level1", func(b *testing.B) {
		measureGzipCompressionRatio(b, MediumSize, BinaryData, 1, "klauspost")
	})
	b.Run("Klauspost-Gzip-Binary-10MB-Level3", func(b *testing.B) {
		measureGzipCompressionRatio(b, MediumSize, BinaryData, 3, "klauspost")
	})
	b.Run("Klauspost-Gzip-Binary-10MB-Level9", func(b *testing.B) {
		measureGzipCompressionRatio(b, MediumSize, BinaryData, 9, "klauspost")
	})

	b.Run("Stdlib-Gzip-Binary-1MB-Level1", func(b *testing.B) {
		measureGzipCompressionRatio(b, SmallSize, BinaryData, 1, "stdlib")
	})
	// Added Level 3 binary compression ratio test
	b.Run("Stdlib-Gzip-Binary-1MB-Level3", func(b *testing.B) {
		measureGzipCompressionRatio(b, SmallSize, BinaryData, 3, "stdlib")
	})
	b.Run("Stdlib-Gzip-Binary-1MB-Level9", func(b *testing.B) {
		measureGzipCompressionRatio(b, SmallSize, BinaryData, 9, "stdlib")
	})
	// Added Level 3 binary compression ratio tests for 10MB
	b.Run("Stdlib-Gzip-Binary-10MB-Level1", func(b *testing.B) {
		measureGzipCompressionRatio(b, MediumSize, BinaryData, 1, "stdlib")
	})
	b.Run("Stdlib-Gzip-Binary-10MB-Level3", func(b *testing.B) {
		measureGzipCompressionRatio(b, MediumSize, BinaryData, 3, "stdlib")
	})
	b.Run("Stdlib-Gzip-Binary-10MB-Level9", func(b *testing.B) {
		measureGzipCompressionRatio(b, MediumSize, BinaryData, 9, "stdlib")
	})

	// Random data (should compress poorly)
	b.Run("Klauspost-Gzip-Random-1MB-Level9", func(b *testing.B) {
		measureGzipCompressionRatio(b, SmallSize, RandomData, 9, "klauspost")
	})
	b.Run("Stdlib-Gzip-Random-1MB-Level9", func(b *testing.B) {
		measureGzipCompressionRatio(b, SmallSize, RandomData, 9, "stdlib")
	})
}

func measureGzipCompressionRatio(b *testing.B, size int, dataType string, level int, implementation string) {
	b.Helper()
	data := GenerateTestData(size, dataType)
	var compressed []byte

	if implementation == "klauspost" {
		var buf bytes.Buffer
		w, err := klauspost.NewWriterLevel(&buf, level)
		if err != nil {
			b.Fatal(err)
		}
		_, err = w.Write(data)
		if err != nil {
			b.Fatal(err)
		}
		err = w.Close()
		if err != nil {
			b.Fatal(err)
		}
		compressed = buf.Bytes()
	} else {
		var buf bytes.Buffer
		w, err := gzip.NewWriterLevel(&buf, level)
		if err != nil {
			b.Fatal(err)
		}
		_, err = w.Write(data)
		if err != nil {
			b.Fatal(err)
		}
		err = w.Close()
		if err != nil {
			b.Fatal(err)
		}
		compressed = buf.Bytes()
	}

	ratio := float64(len(data)) / float64(len(compressed))
	b.ReportMetric(ratio, "ratio")
	b.ReportMetric(0, "ns/op") // This benchmark doesn't measure time
}
