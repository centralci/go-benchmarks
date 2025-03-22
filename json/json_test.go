package json

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/bytedance/sonic"
)

// Initialize the faker once
var faker = gofakeit.New(0)

// Set up sonic with fastest configuration
var sonicFastest = sonic.ConfigFastest

// A more reliable approach to generate test data
type TestData struct {
	ID          int      `json:"id"`
	FirstName   string   `json:"first_name"`
	LastName    string   `json:"last_name"`
	Email       string   `json:"email"`
	Phone       string   `json:"phone"`
	Company     string   `json:"company"`
	JobTitle    string   `json:"job_title"`
	Address     Address  `json:"address"`
	CreatedAt   string   `json:"created_at"`
	UpdatedAt   string   `json:"updated_at"`
	Description string   `json:"description"`
	Password    string   `json:"password"`
	IPAddress   string   `json:"ip_address"`
	UserAgent   string   `json:"user_agent"`
	Tags        []string `json:"tags"`
	Status      string   `json:"status"`
	Metadata    Metadata `json:"metadata"`
}

type Address struct {
	Street     string  `json:"street"`
	City       string  `json:"city"`
	State      string  `json:"state"`
	PostalCode string  `json:"postal_code"`
	Country    string  `json:"country"`
	Latitude   float64 `json:"latitude"`
	Longitude  float64 `json:"longitude"`
}

type Metadata struct {
	Views     int    `json:"views"`
	Likes     int    `json:"likes"`
	Favorites int    `json:"favorites"`
	LastLogin string `json:"last_login"`
	IsPremium bool   `json:"is_premium"`
}

// Statuses to choose from
var statuses = []string{"active", "pending", "inactive", "deleted"}

// Generate test data manually
func generateTestData(count int) []TestData {
	data := make([]TestData, count)

	for i := 0; i < count; i++ {
		tags := make([]string, 5)
		for j := 0; j < 5; j++ {
			tags[j] = faker.Word()
		}

		data[i] = TestData{
			ID:        i + 1,
			FirstName: faker.FirstName(),
			LastName:  faker.LastName(),
			Email:     faker.Email(),
			Phone:     faker.Phone(),
			Company:   faker.Company(),
			JobTitle:  faker.JobTitle(),
			Address: Address{
				Street:     faker.Street(),
				City:       faker.City(),
				State:      faker.State(),
				PostalCode: faker.Zip(),
				Country:    faker.Country(),
				Latitude:   faker.Latitude(),
				Longitude:  faker.Longitude(),
			},
			CreatedAt:   faker.Date().Format("2006-01-02"),
			UpdatedAt:   faker.Date().Format("2006-01-02"),
			Description: faker.Paragraph(3, 5, 15, " "),
			Password:    faker.Password(true, true, true, true, false, 16),
			IPAddress:   faker.IPv4Address(),
			UserAgent:   faker.UserAgent(),
			Tags:        tags,
			Status:      statuses[faker.Number(0, len(statuses)-1)],
			Metadata: Metadata{
				Views:     faker.Number(100, 10000),
				Likes:     faker.Number(10, 1000),
				Favorites: faker.Number(0, 500),
				LastLogin: faker.Date().Format("2006-01-02"),
				IsPremium: faker.Bool(),
			},
		}
	}

	return data
}

// Generate our test data at different sizes
func generateJSON(sizeMB int) []byte {
	var rowCount int

	// Scale the number of rows based on target size
	switch sizeMB {
	case 1:
		rowCount = 250
	case 10:
		rowCount = 2500
	case 100:
		rowCount = 25000
	default:
		rowCount = 100
	}

	data := generateTestData(rowCount)
	jsonData, err := json.Marshal(data)
	if err != nil {
		panic(fmt.Sprintf("Failed to generate test data: %v", err))
	}

	// Print actual size for debugging
	sizeMB_actual := float64(len(jsonData)) / (1024 * 1024)
	fmt.Printf("Generated %d rows, %.2f MB\n", rowCount, sizeMB_actual)

	// If we're way off on size, adjust and try again
	if sizeMB_actual < float64(sizeMB)*0.5 && sizeMB > 1 {
		return generateJSON(sizeMB) // Try again with automatic adjustment
	}

	return jsonData
}

// Generate our test data once for each size
var (
	jsonData1MB   []byte
	jsonData10MB  []byte
	jsonData100MB []byte
)

func init() {
	// We'll generate data in init() to ensure it's available when tests run
	fmt.Println("Generating 1MB test data...")
	jsonData1MB = generateJSON(1)

	fmt.Println("Generating 10MB test data...")
	jsonData10MB = generateJSON(10)

	fmt.Println("Generating 100MB test data...")
	jsonData100MB = generateJSON(100)
}

// 1MB benchmarks
func BenchmarkStdJSONMarshal1MB(b *testing.B) {
	// First unmarshal the data to a Go structure
	var data []TestData
	err := json.Unmarshal(jsonData1MB, &data)
	if err != nil {
		b.Fatalf("Failed to unmarshal test data: %v", err)
	}

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_, err := json.Marshal(data)
		if err != nil {
			b.Fatal(err)
		}
	}

	b.SetBytes(int64(len(jsonData1MB)))
}

func BenchmarkSonicMarshal1MB(b *testing.B) {
	// First unmarshal the data to a Go structure
	var data []TestData
	err := json.Unmarshal(jsonData1MB, &data)
	if err != nil {
		b.Fatalf("Failed to unmarshal test data: %v", err)
	}

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_, err := sonic.Marshal(data)
		if err != nil {
			b.Fatal(err)
		}
	}

	b.SetBytes(int64(len(jsonData1MB)))
}

func BenchmarkSonicFastestMarshal1MB(b *testing.B) {
	// First unmarshal the data to a Go structure
	var data []TestData
	err := json.Unmarshal(jsonData1MB, &data)
	if err != nil {
		b.Fatalf("Failed to unmarshal test data: %v", err)
	}

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_, err := sonicFastest.Marshal(data)
		if err != nil {
			b.Fatal(err)
		}
	}

	b.SetBytes(int64(len(jsonData1MB)))
}

func BenchmarkStdJSONUnmarshal1MB(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var result []TestData
		err := json.Unmarshal(jsonData1MB, &result)
		if err != nil {
			b.Fatal(err)
		}
	}

	b.SetBytes(int64(len(jsonData1MB)))
}

func BenchmarkSonicUnmarshal1MB(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var result []TestData
		err := sonic.Unmarshal(jsonData1MB, &result)
		if err != nil {
			b.Fatal(err)
		}
	}

	b.SetBytes(int64(len(jsonData1MB)))
}

func BenchmarkSonicFastestUnmarshal1MB(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var result []TestData
		err := sonicFastest.Unmarshal(jsonData1MB, &result)
		if err != nil {
			b.Fatal(err)
		}
	}

	b.SetBytes(int64(len(jsonData1MB)))
}

// 10MB benchmarks
func BenchmarkStdJSONMarshal10MB(b *testing.B) {
	// First unmarshal the data to a Go structure
	var data []TestData
	err := json.Unmarshal(jsonData10MB, &data)
	if err != nil {
		b.Fatalf("Failed to unmarshal test data: %v", err)
	}

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_, err := json.Marshal(data)
		if err != nil {
			b.Fatal(err)
		}
	}

	b.SetBytes(int64(len(jsonData10MB)))
}

func BenchmarkSonicMarshal10MB(b *testing.B) {
	// First unmarshal the data to a Go structure
	var data []TestData
	err := json.Unmarshal(jsonData10MB, &data)
	if err != nil {
		b.Fatalf("Failed to unmarshal test data: %v", err)
	}

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_, err := sonic.Marshal(data)
		if err != nil {
			b.Fatal(err)
		}
	}

	b.SetBytes(int64(len(jsonData10MB)))
}

func BenchmarkSonicFastestMarshal10MB(b *testing.B) {
	// First unmarshal the data to a Go structure
	var data []TestData
	err := json.Unmarshal(jsonData10MB, &data)
	if err != nil {
		b.Fatalf("Failed to unmarshal test data: %v", err)
	}

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_, err := sonicFastest.Marshal(data)
		if err != nil {
			b.Fatal(err)
		}
	}

	b.SetBytes(int64(len(jsonData10MB)))
}

func BenchmarkStdJSONUnmarshal10MB(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var result []TestData
		err := json.Unmarshal(jsonData10MB, &result)
		if err != nil {
			b.Fatal(err)
		}
	}

	b.SetBytes(int64(len(jsonData10MB)))
}

func BenchmarkSonicUnmarshal10MB(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var result []TestData
		err := sonic.Unmarshal(jsonData10MB, &result)
		if err != nil {
			b.Fatal(err)
		}
	}

	b.SetBytes(int64(len(jsonData10MB)))
}

func BenchmarkSonicFastestUnmarshal10MB(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var result []TestData
		err := sonicFastest.Unmarshal(jsonData10MB, &result)
		if err != nil {
			b.Fatal(err)
		}
	}

	b.SetBytes(int64(len(jsonData10MB)))
}

// 100MB benchmarks
func BenchmarkStdJSONMarshal100MB(b *testing.B) {
	// First unmarshal the data to a Go structure
	var data []TestData
	err := json.Unmarshal(jsonData100MB, &data)
	if err != nil {
		b.Fatalf("Failed to unmarshal test data: %v", err)
	}

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_, err := json.Marshal(data)
		if err != nil {
			b.Fatal(err)
		}
	}

	b.SetBytes(int64(len(jsonData100MB)))
}

func BenchmarkSonicMarshal100MB(b *testing.B) {
	// First unmarshal the data to a Go structure
	var data []TestData
	err := json.Unmarshal(jsonData100MB, &data)
	if err != nil {
		b.Fatalf("Failed to unmarshal test data: %v", err)
	}

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_, err := sonic.Marshal(data)
		if err != nil {
			b.Fatal(err)
		}
	}

	b.SetBytes(int64(len(jsonData100MB)))
}

func BenchmarkSonicFastestMarshal100MB(b *testing.B) {
	// First unmarshal the data to a Go structure
	var data []TestData
	err := json.Unmarshal(jsonData100MB, &data)
	if err != nil {
		b.Fatalf("Failed to unmarshal test data: %v", err)
	}

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_, err := sonicFastest.Marshal(data)
		if err != nil {
			b.Fatal(err)
		}
	}

	b.SetBytes(int64(len(jsonData100MB)))
}

func BenchmarkStdJSONUnmarshal100MB(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var result []TestData
		err := json.Unmarshal(jsonData100MB, &result)
		if err != nil {
			b.Fatal(err)
		}
	}

	b.SetBytes(int64(len(jsonData100MB)))
}

func BenchmarkSonicUnmarshal100MB(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var result []TestData
		err := sonic.Unmarshal(jsonData100MB, &result)
		if err != nil {
			b.Fatal(err)
		}
	}

	b.SetBytes(int64(len(jsonData100MB)))
}

func BenchmarkSonicFastestUnmarshal100MB(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var result []TestData
		err := sonicFastest.Unmarshal(jsonData100MB, &result)
		if err != nil {
			b.Fatal(err)
		}
	}

	b.SetBytes(int64(len(jsonData100MB)))
}
