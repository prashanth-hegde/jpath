package common

import (
	"reflect"
	"testing"
)

func TestTokenizeB(t *testing.T) {
	testData := []struct {
		name     string
		input    string
		expected []string
	}{
		{"base", `{"hello": "world"}`, []string{`{"hello": "world"}`}},
		{"array", `[{"hello": "world"}, {"hello": "world"}]`, []string{`{"hello": "world"}`, `{"hello": "world"}`}},
		{"array of array", `[[{"hello": "world"}, {"hello": "world"}]]`, []string{`{"hello": "world"}`, `{"hello": "world"}`}},
	}
	for _, testcase := range testData {
		output, _ := Tokenize([]byte(testcase.input))
		var expB [][]byte
		for _, exp := range testcase.expected {
			expB = append(expB, []byte(exp))
		}
		if !reflect.DeepEqual(output, expB) {
			t.Errorf("%s --> failed", testcase.name)
		}
	}
}
