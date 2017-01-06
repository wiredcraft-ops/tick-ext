package client

import (
	"encoding/base64"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestBasicAuth(t *testing.T) {

	oldHtpasswdPath := htpasswdPath
	htpasswdPath = "../testdata/.htpasswd"

	tests := []struct {
		credential         string
		expectedStatusCode int
		expectedResult     bool
	}{
		{"test:test", http.StatusOK, true},
		{"t:", http.StatusUnauthorized, false},
		{"test:eRa787kqS", http.StatusUnauthorized, false},
	}

	for i, test := range tests {

		req, err := http.NewRequest("GET", "/", nil)
		rec := httptest.NewRecorder()
		if err != nil {
			t.Fatalf("F! test %d fatal %v", i, err)
		}
		auth := "Basic " + base64.StdEncoding.EncodeToString([]byte(test.credential))
		req.Header.Set("Authorization", auth)
		if result := basicAuth(rec, req); test.expectedResult != result {
			t.Logf("I! rec %v", rec)
			t.Errorf("E! test %d got %t, expected %t", i, result, test.expectedResult)
		}
	}

	htpasswdPath = oldHtpasswdPath
}
