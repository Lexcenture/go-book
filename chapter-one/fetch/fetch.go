package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {

		buf := new(bytes.Buffer)
		readContent(url, buf)
	}
}

func readContent(url string, buf *bytes.Buffer) {

	resp, err := http.Get(appendProtocol(url))
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		os.Exit(1)
	}
	_, err = io.Copy(buf, resp.Body)

	resp.Body.Close()

	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
		os.Exit(1)
	}
	fmt.Printf("%s\nStatus: %s\n", buf.String(), resp.Status)
}

func appendProtocol(url string) string {
	if strings.HasPrefix(url, "http://") {
		return url
	}
	return "http://" + url
}
