package parser

import (
	"bytes"
	"fmt"
	parser "github.com/buger/jsonparser"
	"github.com/pkg/errors"
	"github.com/prashanth-hegde/jpath/common"
	"regexp"
	"strconv"
	"strings"
)

// Get navigates a dot separated list of fields within a document
// input is the fields and list of tokenized json docs to parse
// Tokenize param if true, treats array as a list of items (for further breakdown)
// if false, treats array as just a string, and does not break down the array
func Get(path string, json [][]byte, tokenize bool) ([][]byte, error) {
	//var tokens = make([][]byte, 0)
	var tokens [][]byte
	fields := strings.Split(path, ".")
	for _, item := range json {
		v, t, _, e := parser.Get(item, fields[0])
		if e != nil {
			// ignore the error and move on
		} else if t == parser.Array && tokenize {
			tokenizedJson, e := common.Tokenize(v)
			if e != nil {
				return nil, errors.Wrapf(e, "error while tokenizing get from field %s", fields[0])
			}
			tokens = append(tokens, tokenizedJson...)
		} else {
			tokens = append(tokens, v)
		}
	}
	if len(fields) > 1 {
		return Get(strings.Join(fields[1:], "."), tokens, tokenize)
	} else {
		return tokens, nil
	}
}

func Filter(path string, json [][]byte) ([][]byte, error) {
	filterRe := common.Matcher.FilterReg
	var e error
	//var filtered = make([][]byte, 0)
	var filtered [][]byte
	// don't worry about the nested for loops
	// the outermost and innermost are guaranteed to have a size of 1
	for _, line := range filterRe.FindAllStringSubmatch(path, -1) {
		// example: results[name.first = Benjamin ]
		//          <--1--> <---2----> 3 <--4--->
		descend := line[1] // capture group 1
		field := line[2]   // capture group 2
		operator := line[3]
		value := line[4]

		if len(descend) > 0 {
			json, e = Get(descend, json, true)
			if e != nil {
				return nil, e
			}
		}
		for _, doc := range json {
			currDoc, e := common.Tokenize(doc)
			if e != nil {
				return nil, e
			}
			values, e := Get(field, currDoc, true)
			if e != nil {
				return nil, e
			}
			for _, intermediate := range values {
				switch operator {
				case "=":
					if bytes.Compare(intermediate, []byte(value)) == 0 {
						filtered = append(filtered, doc)
					}
				case "!=":
					if bytes.Compare(intermediate, []byte(value)) != 0 {
						filtered = append(filtered, doc)
					}
				case "<", ">", "<=", ">=":
					lhs, e := strconv.ParseFloat(string(intermediate), 64)
					if e != nil {
						return nil, errors.Wrapf(e, "error while parsing lhs %s", intermediate)
					}
					rhs, e := strconv.ParseFloat(value, 64)
					if e != nil {
						return nil, errors.Wrapf(e, "error while parsing rhs %s", value)
					}
					if operator == "<" && lhs < rhs ||
						operator == "<=" && lhs <= rhs ||
						operator == ">" && lhs > rhs ||
						operator == ">=" && lhs >= rhs {
						filtered = append(filtered, doc)
					}
				case "~":
					regex := regexp.MustCompile(value)
					matched := regex.Find(intermediate)
					if len(matched) > 0 {
						filtered = append(filtered, doc)
					}
				default:
					// todo: pending
				}
			}
		}
		// todo: implementation pending
	}
	return filtered, nil
}

func Select(path string, json [][]byte) ([][]byte, error) {
	selectRe := common.Matcher.SelectionReg
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
