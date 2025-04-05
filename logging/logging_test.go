package logging

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"log/slog"
	"os"
	"runtime"
	"strings"
	"sync"
	"testing"
	"time"

	"code.cloudfoundry.org/lager/v3"
	"code.cloudfoundry.org/lager/v3/lagerctx"
)

// Custom types for more realistic logging scenarios
type RequestInfo struct {
	Method     string
	Path       string
	Duration   time.Duration
	StatusCode int
	UserAgent  string
	IPAddress  string
}

type UserInfo struct {
	ID        string
	Name      string
	Email     string
	Role      string
	Teams     []string
	Metadata  map[string]interface{}
	CreatedAt time.Time
}

// Size variants for testing with different payload sizes
var (
	smallString  = "small log message"
	mediumString = strings.Repeat("medium sized log message with some additional context ", 5)
	largeString  = strings.Repeat("this is a much larger log message that contains significantly more text to test logging performance with larger payloads ", 20)
)

// setupSlogText creates a new slog logger with text output
func setupSlogText(out io.Writer) *slog.Logger {
	handler := slog.NewTextHandler(out, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	})
	return slog.New(handler)
}

// setupSlogJSON creates a new slog logger with JSON output
func setupSlogJSON(out io.Writer) *slog.Logger {
	handler := slog.NewJSONHandler(out, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	})
	return slog.New(handler)
}

// setupLager creates a new lager logger with the given output
func setupLager(out io.Writer) lager.Logger {
	logger := lager.NewLogger("benchmark-logger")
	logger.RegisterSink(lager.NewWriterSink(out, lager.INFO))
	return logger
}

// setupLagerJSON creates a new lager logger with JSON formatting
func setupLagerJSON(out io.Writer) lager.Logger {
	logger := lager.NewLogger("benchmark-logger")
	sink := lager.NewWriterSink(out, lager.INFO)
	logger.RegisterSink(sink)
	return logger
}

// setupLagerWithRedaction creates a lager logger with a custom redacting sink
func setupLagerWithRedaction(out io.Writer) lager.Logger {
	logger := lager.NewLogger("benchmark-logger")
	// Create custom redacting sink since SetRedactor isn't a standard method
	redactingSink := &customRedactingSink{
		out:      out,
		minLevel: lager.INFO,
	}
	logger.RegisterSink(redactingSink)
	return logger
}

// customRedactingSink implements a lager.Sink with redaction capability
type customRedactingSink struct {
	out      io.Writer
	minLevel lager.LogLevel
}

func (s *customRedactingSink) Log(log lager.LogFormat) {
	if log.LogLevel < s.minLevel {
		return
	}

	// Convert to JSON
	b, err := json.Marshal(log)
	if err != nil {
		return
	}

	// Apply redaction
	redacted := bytes.Replace(b, []byte("secret"), []byte("\"[REDACTED]\""), -1)

	// Write to output
	s.out.Write(redacted)
	s.out.Write([]byte("\n"))
}

// setupSlogWithRedaction creates a slog logger with redaction
func setupSlogWithRedaction(out io.Writer) *slog.Logger {
	// Create a custom handler that redacts sensitive information
	baseHandler := slog.NewTextHandler(out, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	})

	redactingHandler := &redactingHandler{
		handler: baseHandler,
	}

	return slog.New(redactingHandler)
}

// redactingHandler is a custom slog.Handler that redacts sensitive information
type redactingHandler struct {
	handler slog.Handler
}

func (h *redactingHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return h.handler.Enabled(ctx, level)
}

func (h *redactingHandler) Handle(ctx context.Context, r slog.Record) error {
	// Clone the record so we can modify it
	r2 := r.Clone()

	// Add pre-redaction handling here if needed

	return h.handler.Handle(ctx, r2)
}

func (h *redactingHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	// Redact attributes if needed
	redactedAttrs := make([]slog.Attr, len(attrs))
	for i, attr := range attrs {
		if attr.Value.Kind() == slog.KindString && attr.Value.String() == "secret" {
			redactedAttrs[i] = slog.String(attr.Key, "[REDACTED]")
		} else {
			redactedAttrs[i] = attr
		}
	}

	return &redactingHandler{
		handler: h.handler.WithAttrs(redactedAttrs),
	}
}

func (h *redactingHandler) WithGroup(name string) slog.Handler {
	return &redactingHandler{
		handler: h.handler.WithGroup(name),
	}
}

// multiWriter is a simple multi-destination writer
type multiWriter struct {
	writers []io.Writer
}

func newMultiWriter(writers ...io.Writer) io.Writer {
	return &multiWriter{writers: writers}
}

func (w *multiWriter) Write(p []byte) (n int, err error) {
	for _, writer := range w.writers {
		n, err = writer.Write(p)
		if err != nil {
			return n, err
		}
	}
	return len(p), nil
}

// setupSlogMulti creates a slog logger that writes to multiple destinations
func setupSlogMulti() *slog.Logger {
	buffer1 := &bytes.Buffer{}
	buffer2 := &bytes.Buffer{}
	multi := newMultiWriter(buffer1, buffer2, io.Discard)

	return setupSlogText(multi)
}

// setupLagerMulti creates a lager logger that writes to multiple destinations
func setupLagerMulti() lager.Logger {
	logger := lager.NewLogger("benchmark-logger")

	buffer1 := &bytes.Buffer{}
	buffer2 := &bytes.Buffer{}

	sink1 := lager.NewWriterSink(buffer1, lager.INFO)
	sink2 := lager.NewWriterSink(buffer2, lager.INFO)
	sink3 := lager.NewWriterSink(io.Discard, lager.INFO)

	logger.RegisterSink(sink1)
	logger.RegisterSink(sink2)
	logger.RegisterSink(sink3)

	return logger
}

// Basic benchmarks

// BenchmarkSlogSimpleMessage tests slog with a simple string message
func BenchmarkSlogSimpleMessage(b *testing.B) {
	out := io.Discard
	logger := setupSlogText(out)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.Info("this is a simple log message")
	}
}

// BenchmarkLagerSimpleMessage tests lager with a simple string message
func BenchmarkLagerSimpleMessage(b *testing.B) {
	out := io.Discard
	logger := setupLager(out)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.Info("this is a simple log message", lager.Data{})
	}
}

// Structured logging benchmarks

// BenchmarkSlogWithFields tests slog with structured field data
func BenchmarkSlogWithFields(b *testing.B) {
	out := io.Discard
	logger := setupSlogText(out)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.Info("log message with fields",
			"requestID", "req-123",
			"user", "john",
			"action", "login",
			"statusCode", 200,
		)
	}
}

// BenchmarkLagerWithFields tests lager with structured field data
func BenchmarkLagerWithFields(b *testing.B) {
	out := io.Discard
	logger := setupLager(out)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.Info("log message with fields", lager.Data{
			"requestID":  "req-123",
			"user":       "john",
			"action":     "login",
			"statusCode": 200,
		})
	}
}

// JSON output benchmarks

// BenchmarkSlogJSON tests slog with JSON output
func BenchmarkSlogJSON(b *testing.B) {
	out := io.Discard
	logger := setupSlogJSON(out)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.Info("log message with JSON fields",
			"requestID", "req-123",
			"user", "john",
			"action", "login",
			"statusCode", 200,
		)
	}
}

// BenchmarkLagerJSON tests lager with JSON output
func BenchmarkLagerJSON(b *testing.B) {
	out := io.Discard
	logger := setupLagerJSON(out)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.Info("log message with JSON fields", lager.Data{
			"requestID":  "req-123",
			"user":       "john",
			"action":     "login",
			"statusCode": 200,
		})
	}
}

// Context and inheritance benchmarks

// BenchmarkSlogWithContext tests slog with context
func BenchmarkSlogWithContext(b *testing.B) {
	out := io.Discard
	logger := setupSlogText(out)

	// Create a context with values
	ctx := context.Background()
	logger = logger.With("requestID", "req-123", "user", "john")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.InfoContext(ctx, "log message with context",
			"action", "view",
			"resource", "dashboard",
		)
	}
}

// BenchmarkLagerWithContext tests lager with session (lager's context equivalent)
func BenchmarkLagerWithContext(b *testing.B) {
	out := io.Discard
	logger := setupLager(out)

	// Create a context with lager session
	ctx := context.Background()
	sessionLogger := logger.Session("request-session", lager.Data{
		"requestID": "req-123",
		"user":      "john",
	})
	ctx = lagerctx.NewContext(ctx, sessionLogger)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		contextLogger := lagerctx.FromContext(ctx)
		contextLogger.Info("log message with context", lager.Data{
			"action":   "view",
			"resource": "dashboard",
		})
	}
}

// Deep Context/Nesting benchmarks

// BenchmarkSlogDeepContext tests slog with deeply nested context and groups
func BenchmarkSlogDeepContext(b *testing.B) {
	out := io.Discard
	baseLogger := setupSlogText(out)

	// Create a deeply nested logger with groups
	logger := baseLogger.
		With("app", "benchmark").
		With("env", "test").
		WithGroup("request").
		With("id", "req-123").
		With("method", "GET").
		WithGroup("user").
		With("id", "user-456").
		With("role", "admin")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.Info("deeply nested log",
			"action", "view",
			"resource", "dashboard",
		)
	}
}

// BenchmarkLagerDeepContext tests lager with deeply nested contexts
func BenchmarkLagerDeepContext(b *testing.B) {
	out := io.Discard
	baseLogger := setupLager(out)

	// Create deeply nested sessions
	appLogger := baseLogger.Session("app", lager.Data{"name": "benchmark", "env": "test"})
	requestLogger := appLogger.Session("request", lager.Data{"id": "req-123", "method": "GET"})
	userLogger := requestLogger.Session("user", lager.Data{"id": "user-456", "role": "admin"})

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		userLogger.Info("deeply nested log", lager.Data{
			"action":   "view",
			"resource": "dashboard",
		})
	}
}

// Parallel benchmarks

// BenchmarkSlogParallel tests slog with parallel goroutines
func BenchmarkSlogParallel(b *testing.B) {
	out := io.Discard
	logger := setupSlogText(out)

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		counter := 0
		for pb.Next() {
			counter++
			logger.Info("parallel log message",
				"counter", counter,
				"goroutine", "worker",
			)
		}
	})
}

// BenchmarkLagerParallel tests lager with parallel goroutines
func BenchmarkLagerParallel(b *testing.B) {
	out := io.Discard
	logger := setupLager(out)

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		counter := 0
		for pb.Next() {
			counter++
			logger.Info("parallel log message", lager.Data{
				"counter":   counter,
				"goroutine": "worker",
			})
		}
	})
}

// High Concurrency benchmarks

// BenchmarkSlogHighConcurrency tests slog under high concurrency
func BenchmarkSlogHighConcurrency(b *testing.B) {
	out := io.Discard
	logger := setupSlogText(out)

	// Number of goroutines to use (adjust based on available cores)
	concurrency := runtime.GOMAXPROCS(0) * 4

	var wg sync.WaitGroup
	messagePerGoroutine := b.N / concurrency

	b.ResetTimer()

	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			for j := 0; j < messagePerGoroutine; j++ {
				logger.Info("high concurrency message",
					"goroutine", id,
					"counter", j,
					"timestamp", time.Now().UnixNano(),
				)
			}
		}(i)
	}

	wg.Wait()
}

// BenchmarkLagerHighConcurrency tests lager under high concurrency
func BenchmarkLagerHighConcurrency(b *testing.B) {
	out := io.Discard
	logger := setupLager(out)

	// Number of goroutines to use
	concurrency := runtime.GOMAXPROCS(0) * 4

	var wg sync.WaitGroup
	messagePerGoroutine := b.N / concurrency

	b.ResetTimer()

	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			for j := 0; j < messagePerGoroutine; j++ {
				logger.Info("high concurrency message", lager.Data{
					"goroutine": id,
					"counter":   j,
					"timestamp": time.Now().UnixNano(),
				})
			}
		}(i)
	}

	wg.Wait()
}

// Variable payload size benchmarks

// BenchmarkSlogSmallPayload tests slog with a small payload
func BenchmarkSlogSmallPayload(b *testing.B) {
	out := io.Discard
	logger := setupSlogText(out)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.Info(smallString)
	}
}

// BenchmarkLagerSmallPayload tests lager with a small payload
func BenchmarkLagerSmallPayload(b *testing.B) {
	out := io.Discard
	logger := setupLager(out)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.Info(smallString, lager.Data{})
	}
}

// BenchmarkSlogMediumPayload tests slog with a medium payload
func BenchmarkSlogMediumPayload(b *testing.B) {
	out := io.Discard
	logger := setupSlogText(out)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.Info(mediumString)
	}
}

// BenchmarkLagerMediumPayload tests lager with a medium payload
func BenchmarkLagerMediumPayload(b *testing.B) {
	out := io.Discard
	logger := setupLager(out)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.Info(mediumString, lager.Data{})
	}
}

// BenchmarkSlogLargePayload tests slog with a large payload
func BenchmarkSlogLargePayload(b *testing.B) {
	out := io.Discard
	logger := setupSlogText(out)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.Info(largeString)
	}
}

// BenchmarkLagerLargePayload tests lager with a large payload
func BenchmarkLagerLargePayload(b *testing.B) {
	out := io.Discard
	logger := setupLager(out)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.Info(largeString, lager.Data{})
	}
}

// Filtered out (no-op) logging

// BenchmarkSlogFilteredOut tests slog when log level filtering prevents output
func BenchmarkSlogFilteredOut(b *testing.B) {
	out := io.Discard
	// Set log level to ERROR, so INFO logs will be filtered out
	handler := slog.NewTextHandler(out, &slog.HandlerOptions{
		Level: slog.LevelError,
	})
	logger := slog.New(handler)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// This should be a no-op due to level filtering
		logger.Info("this message should be filtered out",
			"requestID", "req-123",
			"counter", i,
		)
	}
}

// BenchmarkLagerFilteredOut tests lager when log level filtering prevents output
func BenchmarkLagerFilteredOut(b *testing.B) {
	out := io.Discard
	logger := lager.NewLogger("benchmark-logger")
	// Only register an ERROR level sink, so INFO logs will be filtered out
	logger.RegisterSink(lager.NewWriterSink(out, lager.ERROR))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// This should be a no-op due to level filtering
		logger.Info("this message should be filtered out", lager.Data{
			"requestID": "req-123",
			"counter":   i,
		})
	}
}

// Complex data structure benchmarks

// BenchmarkSlogComplexMessage tests slog with complex nested data
func BenchmarkSlogComplexStructures(b *testing.B) {
	out := io.Discard
	logger := setupSlogText(out)

	// Prepare complex data structures
	requestInfo := RequestInfo{
		Method:     "POST",
		Path:       "/api/v1/users",
		Duration:   153 * time.Millisecond,
		StatusCode: 201,
		UserAgent:  "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36",
		IPAddress:  "192.168.1.1",
	}

	userInfo := UserInfo{
		ID:        "user-123",
		Name:      "Jane Doe",
		Email:     "jane.doe@example.com",
		Role:      "admin",
		Teams:     []string{"engineering", "security", "platform"},
		CreatedAt: time.Now().Add(-24 * time.Hour * 30),
		Metadata: map[string]interface{}{
			"lastLogin":     time.Now().Add(-24 * time.Hour),
			"loginCount":    42,
			"preferences":   map[string]interface{}{"theme": "dark", "notifications": true},
			"securityLevel": "high",
			"tags":          []string{"vip", "early-adopter", "beta-tester"},
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.Info("complex structured log",
			"requestID", "req-789",
			"timestamp", time.Now(),
			"request", requestInfo,
			"user", userInfo,
			"sessionData", map[string]interface{}{
				"duration":      time.Duration(i) * time.Minute,
				"actions":       []string{"login", "view-dashboard", "update-profile"},
				"authenticated": true,
				"mfaVerified":   true,
				"deviceInfo": map[string]string{
					"type":     "desktop",
					"os":       "macos",
					"browser":  "chrome",
					"language": "en-US",
				},
			},
		)
	}
}

// BenchmarkLagerComplexStructures tests lager with complex nested data
func BenchmarkLagerComplexStructures(b *testing.B) {
	out := io.Discard
	logger := setupLager(out)

	// Prepare user data
	userData := lager.Data{
		"id":        "user-123",
		"name":      "Jane Doe",
		"email":     "jane.doe@example.com",
		"role":      "admin",
		"teams":     []string{"engineering", "security", "platform"},
		"createdAt": time.Now().Add(-24 * time.Hour * 30).String(),
		"metadata": lager.Data{
			"lastLogin":     time.Now().Add(-24 * time.Hour).String(),
			"loginCount":    42,
			"preferences":   lager.Data{"theme": "dark", "notifications": true},
			"securityLevel": "high",
			"tags":          []string{"vip", "early-adopter", "beta-tester"},
		},
	}

	// Prepare request data
	requestData := lager.Data{
		"method":     "POST",
		"path":       "/api/v1/users",
		"duration":   "153ms",
		"statusCode": 201,
		"userAgent":  "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36",
		"ipAddress":  "192.168.1.1",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.Info("complex structured log", lager.Data{
			"requestID": "req-789",
			"timestamp": time.Now().String(),
			"request":   requestData,
			"user":      userData,
			"sessionData": lager.Data{
				"duration":      (time.Duration(i) * time.Minute).String(),
				"actions":       []string{"login", "view-dashboard", "update-profile"},
				"authenticated": true,
				"mfaVerified":   true,
				"deviceInfo": lager.Data{
					"type":     "desktop",
					"os":       "macos",
					"browser":  "chrome",
					"language": "en-US",
				},
			},
		})
	}
}

// Multiple output destinations

// BenchmarkSlogMultiDestination tests slog with multiple output destinations
func BenchmarkSlogMultiDestination(b *testing.B) {
	logger := setupSlogMulti()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.Info("multi-destination log message",
			"requestID", "req-123",
			"counter", i,
		)
	}
}

// BenchmarkLagerMultiDestination tests lager with multiple output destinations
func BenchmarkLagerMultiDestination(b *testing.B) {
	logger := setupLagerMulti()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.Info("multi-destination log message", lager.Data{
			"requestID": "req-123",
			"counter":   i,
		})
	}
}

// Error logging benchmarks

// BenchmarkSlogError tests slog error logging
func BenchmarkSlogError(b *testing.B) {
	out := io.Discard
	logger := setupSlogText(out)
	err := io.EOF

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.Error("error occurred",
			"error", err,
			"operation", "read",
			"retryCount", 3,
		)
	}
}

// BenchmarkLagerError tests lager error logging
func BenchmarkLagerError(b *testing.B) {
	out := io.Discard
	logger := setupLager(out)
	err := io.EOF

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.Error("error occurred", err, lager.Data{
			"operation":  "read",
			"retryCount": 3,
		})
	}
}

// Error logging with stack traces

// BenchmarkSlogErrorWithStack tests slog error logging with stack traces
func BenchmarkSlogErrorWithStack(b *testing.B) {
	out := io.Discard
	logger := setupSlogText(out)

	// Create an error with a stack trace
	err := makeErrorWithStack()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.Error("error with stack trace",
			"error", err,
			"stack", getStackTrace(),
		)
	}
}

// Helper to create an error with a stack trace
func makeErrorWithStack() error {
	return &customError{
		err:   io.EOF,
		stack: getStackTrace(),
	}
}

type customError struct {
	err   error
	stack string
}

func (e *customError) Error() string {
	return e.err.Error()
}

func getStackTrace() string {
	buf := make([]byte, 4096)
	n := runtime.Stack(buf, false)
	return string(buf[:n])
}

// BenchmarkLagerErrorWithStack tests lager error logging with stack traces
func BenchmarkLagerErrorWithStack(b *testing.B) {
	out := io.Discard
	logger := setupLager(out)

	// Create an error with a stack trace
	err := makeErrorWithStack()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.Error("error with stack trace", err, lager.Data{
			"stack": getStackTrace(),
		})
	}
}

// Custom time formatting

// BenchmarkSlogCustomTimeFormat tests slog with custom time formatting
func BenchmarkSlogCustomTimeFormat(b *testing.B) {
	var buf bytes.Buffer

	// Create slog with custom time format
	opts := &slog.HandlerOptions{
		Level: slog.LevelInfo,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.TimeKey {
				if t, ok := a.Value.Any().(time.Time); ok {
					return slog.String(slog.TimeKey, t.Format(time.RFC3339Nano))
				}
			}
			return a
		},
	}

	handler := slog.NewTextHandler(&buf, opts)
	logger := slog.New(handler)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf.Reset()
		logger.Info("log with custom time format",
			"requestID", "req-123",
			"counter", i,
		)
	}
}

// BenchmarkLagerCustomTimeFormat tests lager with custom time formatting
func BenchmarkLagerCustomTimeFormat(b *testing.B) {
	var buf bytes.Buffer
	logger := lager.NewLogger("benchmark-logger")

	// Create a custom sink that formats time
	customSink := lager.NewWriterSink(&buf, lager.INFO)
	logger.RegisterSink(customSink)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf.Reset()
		logger.Info("log with custom time format", lager.Data{
			"requestID": "req-123",
			"counter":   i,
			"timestamp": time.Now().Format(time.RFC3339Nano),
		})
	}
}

// Sensitive data redaction

// BenchmarkSlogWithRedaction tests slog with redaction of sensitive data
func BenchmarkSlogWithRedaction(b *testing.B) {
	out := io.Discard
	logger := setupSlogWithRedaction(out)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.Info("log with sensitive data",
			"username", "john",
			"password", "secret",
			"apiKey", "secret",
			"counter", i,
		)
	}
}

// BenchmarkLagerWithRedaction tests lager with redaction of sensitive data
func BenchmarkLagerWithRedaction(b *testing.B) {
	out := io.Discard
	logger := setupLagerWithRedaction(out)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.Info("log with sensitive data", lager.Data{
			"username": "john",
			"password": "secret",
			"apiKey":   "secret",
			"counter":  i,
		})
	}
}

// File I/O benchmarks

// BenchmarkSlogInfoWithFile tests slog writing to a file
func BenchmarkSlogInfoWithFile(b *testing.B) {
	// Skip during short tests
	if testing.Short() {
		b.Skip("skipping file I/O test in short mode")
	}

	f, err := os.CreateTemp("", "slog-benchmark-*.log")
	if err != nil {
		b.Fatal(err)
	}
	defer os.Remove(f.Name())
	defer f.Close()

	logger := setupSlogText(f)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.Info("file log message", "iteration", i)
	}
}
