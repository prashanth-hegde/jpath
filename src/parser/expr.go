package parser

import (
	"jpath/common"
	"regexp"
)

// fixme: need a smarter way to hold parsed regexes. This is very ugly
const (
	KeyRegex       string = `^[a-zA-Z0-9_-]+$`
	FilterRegex           = `^(\w+)?\[([\w.]+)(=|!=)([^]]+)]$`
	SelectionRegex        = `^{[\w,.-]+}$`
	SliceRegex            = `^\[(\d+)?:(\d+)?\]$`
	CountRegex            = `^#$`
)

// parseExpression Parses the expressions and makes workable tokens out of the expression
// The basic idea is to keep track of dot separators and paranthesis
// Dots are allowed within square and curly brackets. Not allowed outside of them
func parseExpression(expr string) []string {
	var tokens []string
	var paranstack []bool

	var token = ""
	for _, char := range expr {
		switch char {
		case ' ', '"', '\'':
			continue
		case '.':
			if token != "" && len(paranstack) == 0 {
				tokens = append(tokens, token)
				token = ""
			} else if len(paranstack) > 0 {
				token = token + string(char)
			}
		case '[', '{':
			paranstack = append(paranstack, true)
			token = token + string(char)
		case ']', '}':
			if len(paranstack) == 0 {
				common.ExitWithError(common.InvalidExpr)
			}
			paranstack = paranstack[:len(paranstack)-1]
			token = token + string(char)
		default:
			token = token + string(char)
		}
	}
	// if after the loop we still have open paranthesis, that's an error
	if len(token) > 0 && len(paranstack) > 0 {
		common.ExitWithError(common.InvalidExpr)
	} else if len(token) > 0 && len(paranstack) == 0 {
		tokens = append(tokens, token)
	}
	return tokens
}

// ProcessExpression processes a given json array with the matching expressions
// This is the entry point for the jpath parser. This expects a json byte array
// tokenized input already.
func ProcessExpression(expr string, json [][]byte) [][]byte {
	// fixme: the regexes must be centralized somehow. This is a bad place to put it
	keyReg := regexp.MustCompile(KeyRegex)
	filterReg := regexp.MustCompile(FilterRegex)
	selectionReg := regexp.MustCompile(SelectionRegex)
	for _, exp := range parseExpression(expr) {
		if keyReg.MatchString(exp) {
			json = Get(exp, json)
		} else if filterReg.MatchString(exp) {
			json = Filter(exp, json)
		} else if selectionReg.MatchString(exp) {
			// todo: implementation pending
		} else {
			common.ExitWithError(common.InvalidExpr)
		}
	}
	return json
}
