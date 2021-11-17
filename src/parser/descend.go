package parser

import (
	parser "github.com/buger/jsonparser"
	"jpath-go/common"
	"regexp"
)

func Get(path string, json [][]byte) [][]byte {
	var tokens [][]byte
	for _, item := range json {
		v, t, _, e := parser.Get(item, path)
		if e != nil {
			// ignore the error and move on
		} else if t == parser.Array {
			tokens = append(tokens, common.Tokenize(v)...)
		} else {
			tokens = append(tokens, v)
		}
	}
	return tokens
}

func Filter(path string, json [][]byte) [][]byte {
	filterexp := regexp.MustCompile(`^(\w+)?\[([\w.]+)(=|!=)([^]]+)]$`)
	var tokens [][]byte
	for _, line := range filterexp.FindAllStringSubmatch(path, -1) {
		descend := line[1] // capture group 1
		//field := line[2]   // capture group 2
		//operator := line[3]
		//value := line[4]   // capture group 3

		if len(descend) > 0 {
			json = Get(descend, json)
		}
		//intermediate := ProcessExpression(field, json)
		// todo: implementation pending
	}
	return tokens
}
