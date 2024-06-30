package test

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_testHandler(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(testHandler))
	res, err := http.Get(server.URL)
	if err != nil  {
		t.Error(err)
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("expected status code %d, got %d", http.StatusOK, res.StatusCode)

	}

}
