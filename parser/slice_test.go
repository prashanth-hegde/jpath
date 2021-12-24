package parser

import (
	"github.com/prashanth-hegde/jpath/common"
	"reflect"
	"strings"
	"testing"
)

func TestSlice(t *testing.T) {
	testJson, _ := common.Tokenize([]byte(`[{"one",1},{"two",2},{"three",3},{"four",4},{"five",5},{"six",6}]`))
	testData := []struct {
		name     string
		input    string
		expected []string
	}{
		{"base", `[0:1]`, []string{`{"one",1}`}},
		{"empty", `[1:1]`, []string{}},
		{"last", `[-1:]`, []string{`{"six",6}`}},
		{"first", `[:-5]`, []string{`{"one",1}`}},
		{"first2", `[:1]`, []string{`{"one",1}`}},
		{"reversed", `[1:0]`, []string{`{"one",1}`}},
		{"all", `[:]`, []string{`{"one",1}`, `{"two",2}`, `{"three",3}`, `{"four",4}`, `{"five",5}`, `{"six",6}`}},
	}
	for _, testcase := range testData {
		output, _ := Slice(testcase.input, testJson)
		var expB = make([][]byte, 0)
		for _, exp := range testcase.expected {
			expB = append(expB, []byte(exp))
		}
		if !reflect.DeepEqual(output, expB) {
			t.Errorf("%s --> failed \n===\nexpected: %s\n===\nactual: %s\n", testcase.name, expB, output)
		}
	}

	// error cases
	errorData := []struct {
		name     string
		input    string
		expected common.ErrorCode
	}{
		{"overflow", `[0:10]`, common.NumberError},
		{"overflow", `[10:]`, common.NumberError},
	}
	for _, testcase := range errorData {
		_, e := Slice(testcase.input, testJson)
		if e == nil || !strings.ContainsAny(e.Error(), testcase.expected.GetMsg()) {
			t.Errorf("%s --> failed \n===\nexpected: %s\n===\nactual: %s\n",
				testcase.name, testcase.expected.GetMsg(), e.Error())
		}
	}
}
