package config

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func (cfg *Config) ReadJSON(req *http.Request) error {
	data := ReadValue{
		B: []byte(""),
		D: nil,
	}
	if body, err := ioutil.ReadAll(req.Body); err == nil {
		if err = json.Unmarshal([]byte(body), &data.D); err == nil {
			data.B = []byte(body)
			cfg.DataChan <- data
		} else {
			return err
		}
	} else {
		return err
	}
	return nil
}
