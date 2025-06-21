package middleware

import (
	"bufio"
	"bytes"
	"errors"
	"github.com/cynxees/hermes-user/internal/helper"
	"io"
	"log"
	"net"
	"net/http"
	"time"
)

// LoggingMiddleware is a middleware that logs details of the HTTP request and response.
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Log the incoming request details
		startTime := time.Now()
		logger.Printf("Incoming Request: %s %s\n", r.Method, r.URL.Path)
		logger.Printf("Headers: %v\n", r.Header)
		logger.Printf("Remote Addr: %s\n", r.RemoteAddr)
		logger.Printf("Client IP: %s\n", helper.GetClientIP(r))

		// Read the body of the request (if possible)
		var requestBody string
		if r.Body != nil {
			bodyBytes, err := io.ReadAll(r.Body)
			if err != nil {
				logger.Printf("Error reading request body: %v\n", err)
			} else {
				requestBody = string(bodyBytes)
			}
			// Re-assign the body to allow the next handler to read it
			r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		}

		if requestBody != "" {
			logger.Printf("Request Body: %s\n", requestBody)
		}

		// Create a ResponseWriter wrapper to capture the status code and response body
		lrw := &LoggingResponseWriter{ResponseWriter: w, statusCode: http.StatusOK, body: &bytes.Buffer{}}
		next.ServeHTTP(lrw, r)

		// Log the response details
		logger.Printf("Response Status: %d\n", lrw.statusCode)
		logger.Printf("Response Body: %s\n", lrw.body.String())
		logger.Printf("Request processed in %s\n", time.Since(startTime))
	})
}

// LoggingResponseWriter wraps the standard http.ResponseWriter to capture status code and response body
type LoggingResponseWriter struct {
	http.ResponseWriter
	body       *bytes.Buffer
	statusCode int
}

func (lrw *LoggingResponseWriter) WriteHeader(statusCode int) {
	lrw.statusCode = statusCode
	lrw.ResponseWriter.WriteHeader(statusCode)
}

func (lrw *LoggingResponseWriter) Write(p []byte) (n int, err error) {
	lrw.body.Write(p)
	return lrw.ResponseWriter.Write(p)
}

func (lrw *LoggingResponseWriter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	hj, ok := lrw.ResponseWriter.(http.Hijacker)
	if !ok {
		return nil, nil, errors.New("underlying ResponseWriter does not implement http.Hijacker")
	}
	return hj.Hijack()
}
