package useragentrt_test

import (
	"net/http"

	"github.com/mi-wada/useragentrt"
)

func Example() {
	client := &http.Client{}
	client.Transport = useragentrt.New(http.DefaultTransport, "MyAgent/1.0")
	resp, _ := client.Get("https://example.com")
	resp.Body.Close()
}
