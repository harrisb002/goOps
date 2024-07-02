package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// File extension must end in _test to be considered test file!
func TestSegment(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(handleSegment))
	defer testServer.Close()

	testClient := testServer.Client()

	resp, err := testClient.Get(testServer.URL)
	if err != nil {
		t.Fatalf("Failed to send GET request: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Expected status OK, got %v", resp.Status)
	}
}
