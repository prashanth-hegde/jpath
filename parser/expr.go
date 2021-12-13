package parser

import (
	"fmt"
	"github.com/pkg/errors"
	"jpath/common"
)

// parseExpression Parses the expressions and makes workable tokens out of the expression
// The basic idea is to keep track of dot separators and parenthesis
// Dots are allowed within square and curly brackets. Not allowed outside them
func parseExpression(expr string) ([]string, error) {
	var tokens []string
	var parenStack []bool

	var token = ""
	for _, char := range expr {
		switch char {
		case ' ', '"', '\'':
			continue
		case '.':
			if token != "" && len(parenStack) == 0 {
				// example results.name
				tokens = append(tokens, token)
				token = ""
			} else if len(parenStack) > 0 {
				// example results[name.first=Adam]
				token = token + string(char)
			}
		case '[', '{':
			parenStack = append(parenStack, true)
			token = token + string(char)
		case ']', '}':
			if len(parenStack) == 0 {
				return nil, common.InvalidExpr.Error()
			}
			parenStack = parenStack[:len(parenStack)-1]
			token = token + string(char)
		default:
			token = token + string(char)
		}
	}
	// if after the loop we still have open parenthesis, that's an error
	if len(token) > 0 && len(parenStack) > 0 {
		return nil, common.InvalidExpr.Error()
	} else if len(token) > 0 && len(parenStack) == 0 {
		tokens = append(tokens, token)
	}
	return tokens, nil
}

// ProcessExpression processes a given json array with the matching expressions
// This is the entry point for the jpath parser. This expects a json byte array
// tokenized input already.
func ProcessExpression(expr string, json [][]byte) ([][]byte, error) {
	parsedExpr, e := parseExpression(expr)
	if e != nil {
		return nil, errors.Wrap(e, "error while parsing expression")
	}
	for _, exp := range parsedExpr {
		if common.Matcher.KeyReg.MatchString(exp) {
			json, e = Get(exp, json, true)
		} else if common.Matcher.FilterReg.MatchString(exp) {
			json, e = Filter(exp, json)
		} else if common.Matcher.SelectionReg.MatchString(exp) {
			json, e = Select(exp, json)
		} else if common.Matcher.CountReg.MatchString(exp) {
			count := make([][]byte, 1)
			count[0] = []byte(fmt.Sprintf("%d", len(json)))
			json = count
		} else if common.Matcher.SliceReg.MatchString(exp) {
			json, e = Slice(exp, json)
		} else {
			return nil, common.InvalidExpr.Error()
		}
	}
	return json, e
}
