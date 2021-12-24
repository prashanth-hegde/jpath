package parser

import (
	"bytes"
	"fmt"
	"github.com/prashanth-hegde/jpath/common"
	"strings"
)

func Select(path string, json [][]byte) ([][]byte, error) {
	selectRe := common.SelectionReg
	//var selected = make([][]byte, 0)
	var selected [][]byte
	for _, line := range selectRe.FindAllStringSubmatch(path, -1) {
		fields := strings.Split(line[1], ",")
		for _, doc := range json {
			var jsonArray [][]byte
			jsonArray = append(jsonArray, doc)
			var keys [][]byte
			var values [][]byte
			for _, field := range fields {
				result, e := Get(field, jsonArray, false)
				if e != nil {
					return nil, e
				}
				if len(result) > 0 {
					fieldName := field[strings.LastIndex(field, ".")+1:]
					keys = append(keys, []byte(fieldName))
					values = append(values, result[0])
				}
			}
			wrappedObj := constructObject(keys, values)
			tokenizedOutput, e := common.Tokenize(wrappedObj)
			if e != nil {
				return nil, e
			}
			selected = append(selected, tokenizedOutput...)
		}
	}
	return selected, nil
}

func constructObject(keys [][]byte, values [][]byte) []byte {
	if len(values) == 0 || len(keys) != len(values) {
		fmt.Printf("error while parsing, bug in the program :(\n")
		return []byte("")
	}
	var kvp [][]byte
	for i, v := range values {
		var valueWrap []byte
		if v[0] != '{' && v[0] != '[' {
			// wrap the value in double quotes making it "value"
			valueWrap = append(append([]byte("\""), v...), byte('"'))
		} else {
			valueWrap = v
		}
		// format into "key":"value"
		obj := append(append(append([]byte(`"`), keys[i]...), []byte(`":`)...), valueWrap...)
		kvp = append(kvp, obj)
	}
	// wrap the list of "key":"value" from above in a curly bracket, separating the list by comma
	return append(append([]byte("{"), bytes.Join(kvp, []byte(","))...), byte('}'))
}
