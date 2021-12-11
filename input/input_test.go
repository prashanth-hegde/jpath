package input

import (
	"jpath/common"
	"reflect"
	"testing"
)

func TestParseInputJson(t *testing.T) {
	testData := []struct {
		name     string
		input    string
		expected []string
	}{
		{"base", `{"hello": "world"}`, []string{`{"hello": "world"}`}},
		// todo: add stdin tests, I don't know how yet
	}
	for _, testcase := range testData {
		parsed, _ := ParseInputJson(testcase.input)
		output, _ := common.Tokenize(parsed)
		var expB [][]byte
		for _, exp := range testcase.expected {
			expB = append(expB, []byte(exp))
		}
		if !reflect.DeepEqual(output, expB) {
			t.Errorf("%s --> failed", testcase.name)
		}
	}
}
