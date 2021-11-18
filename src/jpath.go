package main

import (
	"flag"
	"fmt"
	"jpath/input"
	"jpath/output"
	"jpath/parser"
	"os"
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
	// 7. tabular output

	flag.Parse()
	expr := strings.TrimSpace(flag.Arg(0))
	json := strings.TrimSpace(flag.Arg(1))

	// validation and print help
	if expr == "" {
		fmt.Println(input.PrintHelp())
		os.Exit(0)
	}

	jsonb := input.ParseInputJson(json)

	parsedOutput := parser.ProcessExpression(expr, jsonb)
	marshal := output.Prettify(parsedOutput, 2)
	fmt.Printf("%s\n", marshal)
}
