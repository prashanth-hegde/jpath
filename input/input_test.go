package input

import (
	"github.com/prashanth-hegde/jpath/common"
	"github.com/prashanth-hegde/jpath/parser"
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

	httpTests := []struct {
		name  string
		input string
		error bool
	}{
		{"base", "https://randomuser.me/api/?results=10", false},
		{"nonexistent", "https://nonexistent.com/", true},
		{"nonjson", "https://example.com/", false},
	}
	for _, testcase := range httpTests {
		parsed, e := ParseInputJson(testcase.input)
		if testcase.error != (e != nil) {
			// testcase.error XOR (e == nil)
			t.Errorf("test error: %s\n", testcase.name)
		}
		tokenized, e := common.Tokenize(parsed)
		out, e := parser.ProcessExpression(".#", tokenized)
		if e != nil && len(out) <= 0 {
			t.Errorf("test error: %s: Expected non-zero, got zero", testcase.name)
		}
	}
}
