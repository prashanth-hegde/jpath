package parser

import (
	"bytes"
	parser "github.com/buger/jsonparser"
	"jpath/common"
	"regexp"
	"strings"
)

// Get navigates a dot separated list of fields within a document
// input is the fields and list of tokenized json docs to parse
func Get(path string, json [][]byte) [][]byte {
	var tokens [][]byte
	fields := strings.Split(path, ".")
	for _, item := range json {
		v, t, _, e := parser.Get(item, fields[0])
		if e != nil {
			// ignore the error and move on
		} else if t == parser.Array {
			tokens = append(tokens, common.Tokenize(v)...)
		} else {
			tokens = append(tokens, v)
		}
	}
	if len(fields) > 1 {
		return Get(strings.Join(fields[1:], "."), tokens)
	} else {
		return tokens
	}
}

func Filter(path string, json [][]byte) [][]byte {
	filterexp := regexp.MustCompile(FilterRegex)
	var filtered [][]byte
	// don't worry about the nested for loops
	// the outermost and innermost are guaranteed to have a size of 1
	// fixme: find a more elegant way to handle the loops :sad_panda:
	for _, line := range filterexp.FindAllStringSubmatch(path, -1) {
		// example: results[name.first = Benjamin ]
		//          <--1--> <---2----> 3 <--4--->
		descend := line[1] // capture group 1
		field := line[2]   // capture group 2
		operator := line[3]
		value := line[4]

		if len(descend) > 0 {
			json = Get(descend, json)
		}
		for _, doc := range json {
			currDoc := common.Tokenize(doc)
			for _, intermediate := range Get(field, currDoc) {
				switch operator {
				case "=":
					if bytes.Compare(intermediate, []byte(value)) == 0 {
						filtered = append(filtered, doc)
					}
				case "!=":
					if bytes.Compare(intermediate, []byte(value)) != 0 {
						filtered = append(filtered, doc)
					}
				default:
					// todo: pending
				}
			}
		}

		// todo: implementation pending
	}
	return filtered
}
