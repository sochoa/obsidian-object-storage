package object

import (
	"net/http/httptest"
	"testing"
)

func FormatObjectURI(b string, o string) string {
	return fmt.Sprintf("/api/v1/%s/%s", b, o)
}

func FailOnErr(t *testing.T, e error) {
	if err != nil {
		t.Fatal(err)
	}
}

func TestCrud(t *testing.T) {
	b := "test-bucket"
	o := "test-object"
	req, err := http.NewRequest("GET", FormatObjectURI(b, o))
	FailOnErr(t, e)
	reqRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(GetObject)
	handler.ServeHTTP(reqRecorder, req)
}
