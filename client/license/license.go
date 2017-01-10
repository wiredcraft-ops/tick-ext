package license

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/Wiredcraft/tick-ext/common"
)

type License struct {
	// ip + company name
	Key string `json:"key"`
}

func Validate(key, apiUrl string) (bool, error) {

	data := License{key}
	dataInJson, err := json.Marshal(data)
	if err != nil {
		return false, err
	}
	reader := bytes.NewReader(dataInJson)
	defaultClient := &http.Client{}

	req, err := http.NewRequest("POST", apiUrl, reader)
	if err != nil {
		return false, err
	}
	resp, err := defaultClient.Do(req)
	if err != nil {
		return false, err
	}

	ret := new(common.Msg)
	err = json.NewDecoder(resp.Body).Decode(ret)
	if err != nil {
		return false, err
	}

	if ret.Success {
		return true, nil
	}

	return false, nil
}
