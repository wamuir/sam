package entity

import (
	_ "embed"
	"encoding/json"
	"testing"
)

//go:embed testdata/120406462.PUBLIC.json
var x []byte

func TestUnmarshalJSON(t *testing.T) {

	r := new(Response)
	if err := json.Unmarshal(x, r); err != nil {
		t.Fatal(err)
	}
}
