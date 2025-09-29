package useragentrt_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mi-wada/useragentrt"
)

func Test(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if got, want := r.Header.Get("User-Agent"), "MyAgent/1.0"; got != want {
			t.Errorf("User-Agent got %v, want %v", got, want)
		}
		w.WriteHeader(http.StatusOK)
	}))
	defer srv.Close()

	client := &http.Client{}
	client.Transport = useragentrt.New(client.Transport, "MyAgent/1.0")
	resp, _ := client.Get(srv.URL)
	resp.Body.Close()
}
