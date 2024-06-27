package controller

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetPage(t *testing.T) {
	data := Data{
		Text: "",
		Banner: "standard",
		Result: "",
	}
	// Create a new HTTP request
	req := httptest.NewRequest(http.MethodGet, "/", nil)

	// Create a new HTTP response recorder
	w := httptest.NewRecorder()
	getPage(w, req, &data)
	res := w.Result()
	defer res.Body.Close()

	// Check the response status code
	if w.Code != http.StatusOK {
		t.Errorf("expected status code %d, got %d", http.StatusOK, w.Code)
	}


	// TODO: Add more assertions to validate the response body or other aspects of the test
}

// func Test_getPage(t *testing.T) {
// 	type args struct {
// 		w    http.ResponseWriter
// 		r    *http.Request
// 		data *Data
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			getPage(tt.args.w, tt.args.r, tt.args.data)
// 		})
// 	}
// }
