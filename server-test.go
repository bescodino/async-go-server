package server

import (
	"net/http"
	"testing"
)

func testServer(t *testing.T) {
	server := Init("127.0.0.1:0")
	if err := server.Start(); err != nil {
		t.Fatal(err)
	}
	url := "http://" + server.Addr
	if _, err := http.Get(url); err != nil {
		t.Fatal(err)
	}
	server.Stop()
	if _, err := http.Get(url); err == nil {
		t.Fatal("error expected")
	}
}
