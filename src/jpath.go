package main

import (
	"flag"
	"jpath-go/common"
	"jpath-go/parser"
	"strings"
)

func main() {
	// todo: flags and features
	// 1. -v for verbose logging
	// 2. -i for indented output (-i 0) for compressed output
	// 3. -z for specifying timezones for timestamps
	// 4. pretty print output
	// 5. colorized output
	// 6. composition

	flag.Parse()
	expr := strings.TrimSpace(flag.Arg(0))
	json := strings.TrimSpace(flag.Arg(1))

	jsonb := common.Tokenize([]byte(json))
	parser.ProcessExpression(expr, jsonb)
}
