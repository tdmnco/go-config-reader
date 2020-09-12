package reader

import (
	"testing"
)

type person struct {
	Firstname string `json:"firstname,omitempty"`
	Lastname  string `json:"lastname,omitempty"`
}

func TestReadJSON(t *testing.T) {
	p := person{}

	err := ReadJSON(&p, "config.json")

	if err != nil {
		t.Errorf(err.Error())
	}

	if p.Firstname != "Kasper" {
		t.Errorf("Unexpected value of Firstname property")
	}

	if p.Lastname != "Tidemann" {
		t.Errorf("Unexpected value of Lastname property")
	}
}
