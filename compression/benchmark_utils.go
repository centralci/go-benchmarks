package compression

import (
	"github.com/brianvoe/gofakeit/v7"
)

// Test data sizes
const (
	SmallSize  = 1 << 20   // 1 MB
	MediumSize = 10 << 20  // 10 MB
	LargeSize  = 100 << 20 // 100 MB
)

// Data types for test generation
const (
	RandomData = "random"
	TextData   = "text"
	BinaryData = "binary"
)

// GenerateTestData creates sample data of specified type and size using gofakeit
func GenerateTestData(size int, dataType string) []byte {
	// Initialize gofakeit with a fixed seed for reproducible results
	gofakeit.Seed(42)

	data := make([]byte, 0, size)

	switch dataType {
	case RandomData:
		// Generate random bytes until we reach the desired size
		for len(data) < size {
			data = append(data, byte(gofakeit.IntRange(0, 255)))
		}
	case TextData:
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
	case BinaryData:
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
