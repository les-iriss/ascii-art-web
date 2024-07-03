package controller

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetPage(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(GetRequest))
	defer server.Close()
	res, err := http.Get(server.URL)
	if err != nil {
		t.Error(err)
	}
	defer res.Body.Close()

	// Check the response status code
	if res.StatusCode != http.StatusOK {
		t.Errorf("expected status code %d, got %d", http.StatusOK, res.StatusCode)
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("failed to read response body: %v", err)
	}
	fmt.Println(string(body))
	// Example assertion for the response body content
	expectedContent := "no such" // Adjust this based on actual expected content
	if !strings.Contains(string(body), expectedContent) {
		t.Errorf("expected response body to contain %q", expectedContent)
	}
}
