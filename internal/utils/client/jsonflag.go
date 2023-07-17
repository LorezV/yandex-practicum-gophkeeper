package client

import (
	"encoding/json"
)

type JSONFlag struct {
	Target interface{}
}

// String is used both by fmt.Print and by Cobra in help text
// https://stackoverflow.com/questions/70285369/how-can-i-provide-json-array-as-argument-to-cobra-cli
func (f *JSONFlag) String() string {
	b, err := json.Marshal(f.Target)
	if err != nil {
		return "failed to marshal object"
	}

	return string(b)
}

// Set must have pointer receiver so it doesn't change the value of a copy
func (f *JSONFlag) Set(v string) error {
	return json.Unmarshal([]byte(v), f.Target)
}

// Type is only used in help text
func (f *JSONFlag) Type() string {
	return "json"
}
