package main

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestItCanReadContent(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hi, there")
	}))
	defer ts.Close()
	buf := new(bytes.Buffer)
	readContent(ts.URL, buf)

	if !strings.ContainsAny(buf.String(), "Hi, there") {
		t.Error("Failed to get the content")
	}
}

func TestItCanAppendProtocol(t *testing.T) {
	const url = "www.bbc.co.uk"
	actualURL := appendProtocol(url)

	if !strings.ContainsAny(actualURL, "http://") {
		t.Error("Could append the protocol")
	}
}
