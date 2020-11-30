package utils

import (
	"encoding/json"
	"io/ioutil"
)

func JsonFromFile(filename string, ptr interface{}) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, ptr)
}
