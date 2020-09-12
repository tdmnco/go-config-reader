package reader

import (
	"encoding/json"
	"io/ioutil"
)

// ReadJSON reads a configuration file in JSON
func ReadJSON(config interface{}, filename string) error {
	f, err := ioutil.ReadFile(filename)

	if err != nil {
		return err
	}

	err = json.Unmarshal([]byte(f), &config)

	if err != nil {
		return err
	}

	return nil
}
