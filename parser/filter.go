package parser

import (
	"bytes"
	"github.com/pkg/errors"
	"github.com/prashanth-hegde/jpath/common"
	"regexp"
	"strconv"
)

func Filter(path string, json [][]byte) ([][]byte, error) {
	filterRe := common.FilterReg
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
