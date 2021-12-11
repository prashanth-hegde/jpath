package output

import (
	"fmt"
	"jpath/common"
	"jpath/input"
	"jpath/parser"
	"strings"
	"testing"
)

func TestPrintOutput(t *testing.T) {
	json := `
    [{"name": "Abraham Lincoln", "quote": "Whatever you are, be a good one"},
    {"name": "Oscar Wilde", "quote": "I can resist everything except temptation"}]
    `
	testData := []struct {
		name    string
		unwrap  bool
		compact bool
	}{
		{"unwrapped-expand", true, false},
		{"unwrapped-compact", true, true},
		{"wrapped-expand", false, false},
		{"wrapped-compact", false, true},
	}
	jsonb, _ := input.ParseInputJson(strings.TrimSpace(json))
	tokenized, _ := common.Tokenize(jsonb)
	for _, testcase := range testData {
		parsed, _ := parser.ProcessExpression(".", tokenized)
		common.Conf.Compress = testcase.compact
		common.Conf.Unwrap = testcase.unwrap
		fmt.Printf("%s -->\n", testcase.name)
		_ = PrintOutput(parsed)
		fmt.Printf("===\n\n")
	}
}

func TestColoredUnwrappedOutput(t *testing.T) {
	testData := []struct {
		name  string
		input []string
		err   bool
	}{
		{"base", []string{`{"one":1, "two":2}`, `{"one":1, "two":2}`}, false},
		{"base2", []string{`{"one": "Uno", "two": "Dos"}`}, false},
		{"string", []string{`one`, `two`}, false},
		{"array", []string{`[{"one":1, "two":2},{"one":1, "two":2}]`}, false},

		// error cases
		{"unmarshalObject", []string{`{"one":1, "two":2`, `{"one":1, "two":2}`}, true},
		{"unmarshalArray", []string{`[{"one":1, "two":2,{"one":1, "two":2}]`}, true},
	}

	for _, testcase := range testData {
		var json [][]byte
		for _, i := range testcase.input {
			json = append(json, []byte(i))
		}
		actual, e := ColoredUnwrappedOutput(json)
		if e != nil && !testcase.err {
			t.Errorf("%s --> not expecting error but got - %s\n", testcase.name, e.Error())
			continue
		} else if e != nil && testcase.err {
			fmt.Printf("error expected: %s\n", e.Error())
			continue
		}
		fmt.Printf("%s -->\n", testcase.name)
		for i := 0; i < len(json); i++ {
			fmt.Printf("%s\n", actual[i])
		}
		fmt.Printf("===\n\n")
	}
}
