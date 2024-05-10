package reader

import (
	"encoding/json"
	"os"
)

// ReadJSON reads a configuration file in JSON.
func ReadJSON(config interface{}, filename string) error {

	// Read the JSON file:
	f, err := os.ReadFile(filename)

	if err != nil {
		return err
	}

	// Unmarshal JSON data:
	err = json.Unmarshal([]byte(f), &config)

	if err != nil {
		return err
	}

	// Return nil on success:
	return nil

}
