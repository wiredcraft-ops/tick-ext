package license

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Wiredcraft/tick-ext/common"
)

func TestValidate(t *testing.T) {

	tests := []struct {
		input    string
		expected bool
	}{
		{"valideKey", true},
		{"invalidKey", false},
		{"ranlQax1N", false},
	}

	// setup test server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		msg := common.Msg{}
		msgInJson := []byte("")
		l := new(License)
		err := json.NewDecoder(r.Body).Decode(l)
		if err != nil {
			msg.Success = false
			msg.Message = err.Error()
		} else {
			switch l.Key {
			case "valideKey":
				msg.Success = true
				msg.Message = "valid key"
			case "invalidKey":
				msg.Success = false
				msg.Message = "invalid key"
			default:
				msg.Success = false
				msg.Message = "invalid key"
			}
		}

		msgInJson, _ = json.Marshal(msg)
		fmt.Fprintf(w, "%s", msgInJson)

	}))
	defer ts.Close()

	for i, test := range tests {
		ret, err := Validate(test.input, ts.URL)
		if err != nil {
			t.Errorf("E! test %d error %v", i, err)
		}
		if test.expected != ret {
			t.Errorf("E! test %d got %t expected %t", i, ret, test.expected)
		}
	}
}
