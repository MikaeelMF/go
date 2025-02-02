package weatherforecastserver

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"testing"
	"time"
)

var port = 9898
var addr = "127.0.0.1:" + fmt.Sprintf("%d", port)

func TestForecastServer(t *testing.T) {
	for {
		l, err := net.Listen("tcp", addr)
		if err == nil {
			_ = l.Close()
			break
		}
		port -= 1
		addr = "127.0.0.1:" + fmt.Sprintf("%d", port)
	}
	testCases := []struct {
		method   string
		path     string
		response string
		code     int
	}{
		{http.MethodGet, fmt.Sprintf("http://%s/", addr), "Please use '/city' to ask for weather status.", http.StatusBadRequest},
		{http.MethodGet, fmt.Sprintf("http://%s/City", addr), "Please use '/city' to ask for weather status.", http.StatusBadRequest},
		{http.MethodGet, fmt.Sprintf("http://%s/city", addr), "Goblinocus will have a nice sunny day tomorrow!", http.StatusOK},
		{http.MethodPut, fmt.Sprintf("http://%s/city", addr), "", http.StatusMethodNotAllowed},
	}
	srv := ForecastServer(addr, 30*time.Second)
	go func() {
		err := srv.ListenAndServe()
		if err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()
	time.Sleep(3 * time.Second)
	for i, c := range testCases {
		client := http.Client{}
		req, err := http.NewRequest(c.method, c.path, http.NoBody)
		if err != nil {
			t.Fatalf("could not make the request. error: %s", err)
		}
		resp, err := client.Do(req)
		if err != nil {
			t.Fatalf("could not make the request. error: %s", err)
		}
		if c.code != resp.StatusCode {
			t.Fatalf("expected response code %d, received %d. test case number: %d", c.code, resp.StatusCode, i+1)
		}
		b, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Fatal(err)
		}
		_ = resp.Body.Close()
		if string(b) != c.response {
			t.Fatalf("expected response: %s, received %s. test case number: %d", c.response, string(b), i+1)
		}
	}
	srv.Close()
}
