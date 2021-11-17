package parser

import (
	"jpath-go/common"
	"regexp"
)

// fixme: need a smarter way to hold parsed regexes. This is very ugly
type ExprRegex string

const (
	KEY_REGEX    ExprRegex = `^[a-zA-Z0-9_-]+$`
	FILTER_REGEX           = `^(\w+)?\[([\w.]+)(=|!=)([^]]+)]$`
)

// ParseExpression Parses the expressions and makes workable tokens out of the expression
// The basic idea is to keep track of dot separators and paranthesis
// Dots are allowed within square and curly brackets. Not allowed outside of them
func ParseExpression(expr string) []string {
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

func ProcessExpression(expr string, json [][]byte) [][]byte {
	var filtered [][]byte
	// fixme: the regexes must be centralized somehow. This is a bad place to put it
	keyReg := regexp.MustCompile(`^[a-zA-Z0-9_-]+$`)
	filterReg := regexp.MustCompile(`^(\w+)?\[([\w.]+)(=|!=)([^]]+)]$`)
	for _, exp := range ParseExpression(expr) {
		if keyReg.MatchString(exp) {
			json = Get(exp, json)
		} else if filterReg.MatchString(exp) {
			json = append(filtered, Filter(exp, json)...)
			// todo: implementation pending
		} else {
			common.ExitWithError(common.InvalidExpr)
		}
	}

	return json
}
