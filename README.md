# useragentrt

[![Godoc Reference](https://godoc.org/github.com/mi-wada/useragentrt?status.svg)](http://godoc.org/github.com/mi-wada/useragentrt)

A Go package that provides a http.RoundTripper which sets a User-Agent header.

## Usage

```go
package main

import (
	"net/http"

	"github.com/mi-wada/useragentrt"
)

func main() {
	client := &http.Client{}
	client.Transport = useragentrt.New(client.Transport, "MyAgent/1.0")
	resp, _ := client.Get("https://example.com")
	resp.Body.Close()
}
```

## Install

```shell
go get github.com/mi-wada/useragentrt@latest
```
