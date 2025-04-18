package benchmark

import (
	"bytes"
	"io"
	"math/rand"
	"strconv"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	goyaml "github.com/goccy/go-yaml"
	yamlv2 "gopkg.in/yaml.v2"
	yamlv3 "gopkg.in/yaml.v3"
)

// Sample data structures for benchmarking
type SimpleStruct struct {
	Name        string  `yaml:"name"`
	Value       int     `yaml:"value"`
	Description string  `yaml:"description"`
	Enabled     bool    `yaml:"enabled"`
	Score       float64 `yaml:"score"`
}

type ComplexStruct struct {
	ID          string                 `yaml:"id"`
	Name        string                 `yaml:"name"`
	Description string                 `yaml:"description"`
	Active      bool                   `yaml:"active"`
	Tags        []string               `yaml:"tags"`
	Metadata    map[string]interface{} `yaml:"metadata"`
	Items       []Item                 `yaml:"items"`
	Categories  []Category             `yaml:"categories"`
	Config      Config                 `yaml:"config"`
}

type Item struct {
	ID          string            `yaml:"id"`
	Name        string            `yaml:"name"`
	Description string            `yaml:"description"`
	Price       float64           `yaml:"price"`
	Quantity    int               `yaml:"quantity"`
	Properties  map[string]string `yaml:"properties"`
}

type Category struct {
	ID       string   `yaml:"id"`
	Name     string   `yaml:"name"`
	ParentID string   `yaml:"parent_id"`
	Path     []string `yaml:"path"`
	Priority int      `yaml:"priority"`
}

type Config struct {
	Timeout          int      `yaml:"timeout"`
	RetryCount       int      `yaml:"retry_count"`
	Debug            bool     `yaml:"debug"`
	LogLevel         string   `yaml:"log_level"`
	AllowedDomains   []string `yaml:"allowed_domains"`
	RateLimitPerHour int      `yaml:"rate_limit_per_hour"`
}

func init() {
	// Initialize the faker
	gofakeit.Seed(0)
}

func createSimpleData() SimpleStruct {
	return SimpleStruct{
		Name:        gofakeit.ProductName(),
		Value:       gofakeit.Number(1, 10000),
		Description: gofakeit.Sentence(5),
		Enabled:     gofakeit.Bool(),
		Score:       gofakeit.Float64Range(0.0, 100.0),
	}
}

func createComplexData() ComplexStruct {
	// Create random number of items
	numItems := gofakeit.Number(3, 15)
	items := make([]Item, numItems)
	for i := 0; i < numItems; i++ {
		// Create random properties
		numProps := gofakeit.Number(2, 8)
		props := make(map[string]string)
		for j := 0; j < numProps; j++ {
			props[gofakeit.Word()] = gofakeit.Sentence(2)
		}

		items[i] = Item{
			ID:          gofakeit.UUID(),
			Name:        gofakeit.ProductName(),
			Description: gofakeit.Paragraph(1, 3, 5, "."),
			Price:       gofakeit.Price(0.99, 999.99),
			Quantity:    gofakeit.Number(1, 100),
			Properties:  props,
		}
	}

	// Create random number of categories
	numCategories := gofakeit.Number(2, 10)
	categories := make([]Category, numCategories)
	for i := 0; i < numCategories; i++ {
		pathDepth := gofakeit.Number(1, 5)
		path := make([]string, pathDepth)
		for j := 0; j < pathDepth; j++ {
			path[j] = gofakeit.Word()
		}

		categories[i] = Category{
			ID:       gofakeit.UUID(),
			Name:     gofakeit.BuzzWord(),
			ParentID: gofakeit.UUID(),
			Path:     path,
			Priority: gofakeit.Number(1, 100),
		}
	}

	// Create random metadata
	numMetadata := gofakeit.Number(5, 20)
	metadata := make(map[string]interface{})
	for i := 0; i < numMetadata; i++ {
		key := gofakeit.Word()
		switch gofakeit.Number(0, 4) {
		case 0:
			metadata[key] = gofakeit.Bool()
		case 1:
			metadata[key] = gofakeit.Number(1, 1000)
		case 2:
			metadata[key] = gofakeit.Float64()
		case 3:
			metadata[key] = gofakeit.Sentence(3)
		case 4:
			// Nested map
			nestedMap := make(map[string]interface{})
			nestedKeys := gofakeit.Number(2, 5)
			for j := 0; j < nestedKeys; j++ {
				nestedMap[gofakeit.Word()] = gofakeit.Sentence(2)
			}
			metadata[key] = nestedMap
		}
	}

	// Create a list of random tags
	numTags := gofakeit.Number(3, 15)
	tags := make([]string, numTags)
	for i := 0; i < numTags; i++ {
		tags[i] = gofakeit.HackerNoun()
	}

	return ComplexStruct{
		ID:          gofakeit.UUID(),
		Name:        gofakeit.AppName(),
		Description: gofakeit.Paragraph(2, 4, 6, "."),
		Active:      gofakeit.Bool(),
		Tags:        tags,
		Metadata:    metadata,
		Items:       items,
		Categories:  categories,
		Config: Config{
			Timeout:          gofakeit.Number(1000, 30000),
			RetryCount:       gofakeit.Number(0, 10),
			Debug:            gofakeit.Bool(),
			LogLevel:         gofakeit.RandomString([]string{"debug", "info", "warn", "error", "fatal"}),
			AllowedDomains:   []string{gofakeit.DomainName(), gofakeit.DomainName(), gofakeit.DomainName()},
			RateLimitPerHour: gofakeit.Number(100, 10000),
		},
	}
}

// Generate a more realistic template with variable numbers of resources and jobs
func createTemplateData() string {
	templateBuilder := `
# This is a sample CI/CD configuration file
version: ` + strconv.Itoa(gofakeit.Number(1, 5)) + `
env:
  GIT_AUTHOR_NAME: "` + gofakeit.Name() + `"
  GIT_AUTHOR_EMAIL: "` + gofakeit.Email() + `"
  API_KEY: ((secrets.api_key))
  DEBUG: ` + strconv.FormatBool(gofakeit.Bool()) + `

resources:
`

	// Add random number of resources
	numResources := gofakeit.Number(3, 10)
	for i := 0; i < numResources; i++ {
		resourceType := gofakeit.RandomString([]string{"git", "s3", "registry-image", "time", "webhook"})

		templateBuilder += `- name: ` + gofakeit.AppName() + `-` + strconv.Itoa(i) + `
  type: ` + resourceType + `
  source:
`
		switch resourceType {
		case "git":
			templateBuilder += `    uri: git@github.com:` + gofakeit.Username() + `/` + gofakeit.AppName() + `.git
    branch: ((` + gofakeit.Word() + `.branch))
    private_key: ((secrets.` + gofakeit.Word() + `_key))
`
		case "s3":
			templateBuilder += `    bucket: ((` + gofakeit.Word() + `.bucket))
    access_key_id: ((secrets.aws_access_key))
    secret_access_key: ((secrets.aws_secret_key))
    region_name: ` + gofakeit.RandomString([]string{"us-east-1", "us-west-2", "eu-west-1", "ap-southeast-1"}) + `
`
		case "registry-image":
			templateBuilder += `    repository: ` + gofakeit.Username() + `/` + gofakeit.AppName() + `
    tag: ((` + gofakeit.Word() + `.version))
    username: ((secrets.registry_username))
    password: ((secrets.registry_password))
`
		}
	}

	templateBuilder += `
jobs:
`
	// Add random number of jobs
	numJobs := gofakeit.Number(2, 8)
	for i := 0; i < numJobs; i++ {
		jobName := gofakeit.JobTitle() + "-service-" + strconv.Itoa(i)
		templateBuilder += `- name: ` + jobName + `
  plan:
`
		// Add random number of steps
		numSteps := gofakeit.Number(2, 6)
		for j := 0; j < numSteps; j++ {
			stepType := gofakeit.RandomString([]string{"get", "put", "task"})

			switch stepType {
			case "get":
				templateBuilder += `  - get: ` + gofakeit.AppName() + `-` + strconv.Itoa(rand.Intn(numResources)) + `
    trigger: ` + strconv.FormatBool(gofakeit.Bool()) + `
`
			case "put":
				templateBuilder += `  - put: ` + gofakeit.AppName() + `-` + strconv.Itoa(rand.Intn(numResources)) + `
    params:
      file: ` + gofakeit.FileExtension() + `
`
			case "task":
				templateBuilder += `  - task: ` + gofakeit.HackerVerb() + `-` + gofakeit.HackerNoun() + `
    file: ((` + gofakeit.Word() + `.path))/tasks/` + gofakeit.Word() + `.yml
    params:
      ENV: ((` + gofakeit.Word() + `.env))
      DEBUG: ((` + gofakeit.Word() + `.debug))
      TIMEOUT: ` + strconv.Itoa(gofakeit.Number(30, 600)) + `
`
			}
		}
	}

	return templateBuilder
}

// Benchmarks for YAML v2

func BenchmarkYAMLv2MarshalSimple(b *testing.B) {
	data := createSimpleData()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := yamlv2.Marshal(data)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkYAMLv2UnmarshalSimple(b *testing.B) {
	data := createSimpleData()
	bytes, _ := yamlv2.Marshal(data)
	var result SimpleStruct
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		err := yamlv2.Unmarshal(bytes, &result)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkYAMLv2MarshalComplex(b *testing.B) {
	data := createComplexData()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := yamlv2.Marshal(data)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkYAMLv2UnmarshalComplex(b *testing.B) {
	data := createComplexData()
	bytes, _ := yamlv2.Marshal(data)
	var result ComplexStruct
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		err := yamlv2.Unmarshal(bytes, &result)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkYAMLv2UnmarshalTemplate(b *testing.B) {
	templateStr := createTemplateData()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var result interface{}
		err := yamlv2.Unmarshal([]byte(templateStr), &result)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// Benchmarks for YAML v3

func BenchmarkYAMLv3MarshalSimple(b *testing.B) {
	data := createSimpleData()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := yamlv3.Marshal(data)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// Benchmarks for go-yaml (goccy)

func BenchmarkGoYAMLMarshalSimple(b *testing.B) {
	data := createSimpleData()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := goyaml.Marshal(data)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkYAMLv3UnmarshalSimple(b *testing.B) {
	data := createSimpleData()
	bytes, _ := yamlv3.Marshal(data)
	var result SimpleStruct
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		err := yamlv3.Unmarshal(bytes, &result)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkGoYAMLUnmarshalSimple(b *testing.B) {
	data := createSimpleData()
	bytes, _ := goyaml.Marshal(data)
	var result SimpleStruct
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		err := goyaml.Unmarshal(bytes, &result)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkYAMLv3MarshalComplex(b *testing.B) {
	data := createComplexData()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := yamlv3.Marshal(data)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkGoYAMLMarshalComplex(b *testing.B) {
	data := createComplexData()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := goyaml.Marshal(data)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkYAMLv3UnmarshalComplex(b *testing.B) {
	data := createComplexData()
	bytes, _ := yamlv3.Marshal(data)
	var result ComplexStruct
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		err := yamlv3.Unmarshal(bytes, &result)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkGoYAMLUnmarshalComplex(b *testing.B) {
	data := createComplexData()
	bytes, _ := goyaml.Marshal(data)
	var result ComplexStruct
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		err := goyaml.Unmarshal(bytes, &result)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkYAMLv3UnmarshalTemplate(b *testing.B) {
	templateStr := createTemplateData()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var result interface{}
		err := yamlv3.Unmarshal([]byte(templateStr), &result)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkGoYAMLUnmarshalTemplate(b *testing.B) {
	templateStr := createTemplateData()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var result interface{}
		err := goyaml.Unmarshal([]byte(templateStr), &result)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// Mixed benchmarks (more like real-world use cases)

func BenchmarkMixedTemplateYAMLv2(b *testing.B) {
	templateStr := createTemplateData()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var result interface{}
		// First unmarshal the template
		err := yamlv2.Unmarshal([]byte(templateStr), &result)
		if err != nil {
			b.Fatal(err)
		}

		// Then marshal it back (simulating template processing)
		_, err = yamlv2.Marshal(result)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkMixedTemplateGoYAML(b *testing.B) {
	templateStr := createTemplateData()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var result interface{}
		// First unmarshal the template
		err := goyaml.Unmarshal([]byte(templateStr), &result)
		if err != nil {
			b.Fatal(err)
		}

		// Then marshal it back (simulating template processing)
		_, err = goyaml.Marshal(result)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkMixedTemplateYAMLv3(b *testing.B) {
	templateStr := createTemplateData()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var result interface{}
		// First unmarshal the template
		err := yamlv3.Unmarshal([]byte(templateStr), &result)
		if err != nil {
			b.Fatal(err)
		}

		// Then marshal it back (simulating template processing)
		_, err = yamlv3.Marshal(result)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// Random data benchmarks to simulate real-world scenarios

func BenchmarkRandomDataYAMLv2(b *testing.B) {
	b.Run("SimpleVaried", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			// Create new random data for each iteration
			data := createSimpleData()
			bytes, err := yamlv2.Marshal(data)
			if err != nil {
				b.Fatal(err)
			}

			var result SimpleStruct
			err = yamlv2.Unmarshal(bytes, &result)
			if err != nil {
				b.Fatal(err)
			}
		}
	})

	b.Run("ComplexVaried", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			// Create new random data for each iteration
			data := createComplexData()
			bytes, err := yamlv2.Marshal(data)
			if err != nil {
				b.Fatal(err)
			}

			var result interface{}
			err = yamlv2.Unmarshal(bytes, &result)
			if err != nil {
				b.Fatal(err)
			}
		}
	})

	b.Run("TemplateVaried", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			// Create new random template for each iteration
			templateStr := createTemplateData()

			var result interface{}
			err := yamlv2.Unmarshal([]byte(templateStr), &result)
			if err != nil {
				b.Fatal(err)
			}
		}
	})
}

func BenchmarkRandomDataYAMLv3(b *testing.B) {
	b.Run("SimpleVaried", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			// Create new random data for each iteration
			data := createSimpleData()
			bytes, err := yamlv3.Marshal(data)
			if err != nil {
				b.Fatal(err)
			}

			var result SimpleStruct
			err = yamlv3.Unmarshal(bytes, &result)
			if err != nil {
				b.Fatal(err)
			}
		}
	})

	b.Run("ComplexVaried", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			// Create new random data for each iteration
			data := createComplexData()
			bytes, err := yamlv3.Marshal(data)
			if err != nil {
				b.Fatal(err)
			}

			var result interface{}
			err = yamlv3.Unmarshal(bytes, &result)
			if err != nil {
				b.Fatal(err)
			}
		}
	})

	b.Run("TemplateVaried", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			// Create new random template for each iteration
			templateStr := createTemplateData()

			var result interface{}
			err := yamlv3.Unmarshal([]byte(templateStr), &result)
			if err != nil {
				b.Fatal(err)
			}
		}
	})
}

func BenchmarkRandomDataGoYAML(b *testing.B) {
	b.Run("SimpleVaried", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			// Create new random data for each iteration
			data := createSimpleData()
			bytes, err := goyaml.Marshal(data)
			if err != nil {
				b.Fatal(err)
			}

			var result SimpleStruct
			err = goyaml.Unmarshal(bytes, &result)
			if err != nil {
				b.Fatal(err)
			}
		}
	})

	b.Run("ComplexVaried", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			// Create new random data for each iteration
			data := createComplexData()
			bytes, err := goyaml.Marshal(data)
			if err != nil {
				b.Fatal(err)
			}

			var result interface{}
			err = goyaml.Unmarshal(bytes, &result)
			if err != nil {
				b.Fatal(err)
			}
		}
	})

	b.Run("TemplateVaried", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			// Create new random template for each iteration
			templateStr := createTemplateData()

			var result interface{}
			err := goyaml.Unmarshal([]byte(templateStr), &result)
			if err != nil {
				b.Fatal(err)
			}
		}
	})
}

// Bulk operation benchmarks

func BenchmarkBulkOperationsYAMLv2(b *testing.B) {
	// Create a batch of random data objects
	batchSize := 50
	simpleDataBatch := make([]SimpleStruct, batchSize)
	complexDataBatch := make([]ComplexStruct, batchSize)

	for i := 0; i < batchSize; i++ {
		simpleDataBatch[i] = createSimpleData()
		complexDataBatch[i] = createComplexData()
	}

	b.Run("BulkMarshalSimple", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			for j := 0; j < batchSize; j++ {
				_, err := yamlv2.Marshal(simpleDataBatch[j])
				if err != nil {
					b.Fatal(err)
				}
			}
		}
	})

	b.Run("BulkMarshalComplex", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			for j := 0; j < batchSize; j++ {
				_, err := yamlv2.Marshal(complexDataBatch[j])
				if err != nil {
					b.Fatal(err)
				}
			}
		}
	})
}

func BenchmarkBulkOperationsYAMLv3(b *testing.B) {
	// Create a batch of random data objects
	batchSize := 50
	simpleDataBatch := make([]SimpleStruct, batchSize)
	complexDataBatch := make([]ComplexStruct, batchSize)

	for i := 0; i < batchSize; i++ {
		simpleDataBatch[i] = createSimpleData()
		complexDataBatch[i] = createComplexData()
	}

	b.Run("BulkMarshalSimple", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			for j := 0; j < batchSize; j++ {
				_, err := yamlv3.Marshal(simpleDataBatch[j])
				if err != nil {
					b.Fatal(err)
				}
			}
		}
	})

	b.Run("BulkMarshalComplex", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			for j := 0; j < batchSize; j++ {
				_, err := yamlv3.Marshal(complexDataBatch[j])
				if err != nil {
					b.Fatal(err)
				}
			}
		}
	})
}

func BenchmarkBulkOperationsGoYAML(b *testing.B) {
	// Create a batch of random data objects
	batchSize := 50
	simpleDataBatch := make([]SimpleStruct, batchSize)
	complexDataBatch := make([]ComplexStruct, batchSize)

	for i := 0; i < batchSize; i++ {
		simpleDataBatch[i] = createSimpleData()
		complexDataBatch[i] = createComplexData()
	}

	b.Run("BulkMarshalSimple", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			for j := 0; j < batchSize; j++ {
				_, err := goyaml.Marshal(simpleDataBatch[j])
				if err != nil {
					b.Fatal(err)
				}
			}
		}
	})

	b.Run("BulkMarshalComplex", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			for j := 0; j < batchSize; j++ {
				_, err := goyaml.Marshal(complexDataBatch[j])
				if err != nil {
					b.Fatal(err)
				}
			}
		}
	})
}

// Additional benchmarks specific to go-yaml features

func BenchmarkGoYAMLStreamProcessing(b *testing.B) {
	// Testing stream decode capability - a feature specific to go-yaml
	templateStr := createTemplateData()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Create a bytes.Reader from the string to properly use the decoder
		reader := bytes.NewReader([]byte(templateStr))
		decoder := goyaml.NewDecoder(reader)
		for {
			var v interface{}
			if err := decoder.Decode(&v); err != nil {
				if err == io.EOF {
					break
				}
				b.Fatal(err)
			}
		}
	}
}
