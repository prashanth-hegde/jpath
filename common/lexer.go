package common

import (
	"bytes"
	parser "github.com/buger/jsonparser"
	"github.com/pkg/errors"
)

func Tokenize(json []byte) ([][]byte, error) {
	if len(json) == 0 {
		return nil, nil
	}
	var tokens [][]byte
	v, t, _, e := parser.Get(json)
	if e != nil {
		return nil, errors.Wrap(InvalidJson.Error(), "error during tokenization")
	}
	if t == parser.Array {
		tokens = append(tokens, extractElementsFromArray(v)...)
	} else {
		tokens = append(tokens, v)
	}

	return tokens, nil
}

func extractElementsFromArray(json []byte) [][]byte {
	var tokens [][]byte
	_, er := parser.ArrayEach(json, func(value []byte, dataType parser.ValueType, offset int, err error) {
		v, typ, _, _ := parser.Get(value)
		if typ == parser.Array {
			tokens = append(tokens, extractElementsFromArray(v)...)
		} else {
			tokens = append(tokens, v)
		}
	})
	if er != nil {
		ExitWithError(InvalidJson)
	}
	return tokens
}

func WrapIntoArray(json [][]byte) []byte {
	return append(append([]byte("["), bytes.Join(json, []byte(","))...), []byte("]")...)
}
