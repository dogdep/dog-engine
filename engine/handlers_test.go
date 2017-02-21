package engine

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"io/ioutil"
	"os"
)

func TestDeployEndpointDoesNotAllowGet(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/deploy", nil)
	if err != nil {
		t.Fatal(err)
	}

	r := httptest.NewRecorder()
	handler := http.HandlerFunc(PostDeployHandler)

	handler.ServeHTTP(r, req)

	if status := r.Code; status != http.StatusMethodNotAllowed {
		t.Errorf("wrong status code: got %v wanted %v", status, http.StatusMethodNotAllowed)
	}
}

func TestAcceptsArchiveUpload(t *testing.T) {
	file, err := ioutil.TempFile(os.TempDir(), "prefix")
	defer os.Remove(file.Name())

	req, err := http.NewRequest("POST", "/api/deploy", nil)
	if err != nil {
		t.Fatal(err)
	}

	r := httptest.NewRecorder()
	handler := http.HandlerFunc(PostDeployHandler)

	handler.ServeHTTP(r, req)

	if status := r.Code; status != http.StatusOK {
		t.Errorf("wrong status code: got %v wanted %v", status, http.StatusOK)
	}
}
