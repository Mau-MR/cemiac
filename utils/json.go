package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// ToJSON serializes the given interface into a string based JSON format
func ToJSON(i interface{}, w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(i)
}

// FromJSON deserializes the object from JSON string
// in an io.Reader to the given interface
func FromJSON(i interface{}, r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(i)
}

func ParseRequest(i interface{}, r io.Reader, rw http.ResponseWriter) error{
	err := FromJSON(i,r)
	if err !=nil {
		ToJSON(GenericError{
			Message: fmt.Sprintf("Bad  format for type: %T", i),
		},rw)
	}
	return err
}
