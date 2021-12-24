package parser

import (
	parser "github.com/buger/jsonparser"
	"github.com/pkg/errors"
	"github.com/prashanth-hegde/jpath/common"
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
