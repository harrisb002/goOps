package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// File extension must end in _test to be considered test file!

func TestSegmentingSucceed(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(handleSegment))
	defer testServer.Close()

	testClient := testServer.Client()

	resp, err := testClient.Get(testServer.URL)
	if err != nil {
		t.Error(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, "Segmenting!", string(body))
	assert.Equal(t, 200, resp.StatusCode)

}

func TestSegmentingFail(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(handleSegment))
	defer testServer.Close()

	testClient := testServer.Client()

	body := strings.NewReader("my body")

	resp, err := testClient.Post(testServer.URL, "application/json", body)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, 405, resp.StatusCode)

}

func TestHealthSucceed(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(handleHealth))
	defer testServer.Close()

	testClient := testServer.Client()

	resp, err := testClient.Get(testServer.URL)
	if err != nil {
		t.Error(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, "Server Status Ok", string(body))
	assert.Equal(t, 200, resp.StatusCode)

}

func TestHealthFail(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(handleHealth))
	defer testServer.Close()

	testClient := testServer.Client()

	body := strings.NewReader("my body")

	resp, err := testClient.Post(testServer.URL, "application/json", body)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, 405, resp.StatusCode)

}
