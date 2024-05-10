package reader

import (
	"testing"
)

type TestPerson struct {
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
}

func TestReadJSON(t *testing.T) {

	// Prepare a test person:
	p := TestPerson{}

	// Read JSON file:
	err := ReadJSON(&p, "reader_test.json")

	if err != nil {
		t.Errorf(err.Error())
	}

	// Validate first name:
	if p.FirstName != "Kasper" {
		t.Errorf("Unexpected value of FirstName property")
	}

	// Validate last name:
	if p.LastName != "Tidemann" {
		t.Errorf("Unexpected value of LastName property")
	}
}
