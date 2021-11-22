package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"jpath/common"
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

	// mandatory variables
	var expr string
	var json string

	// root command parser
	var rootCmd = &cobra.Command{
		Use:   "jpath <expression> <json>",
		Short: "analyzer for json data",
		Long: `An easy to use json filter to analyze json documents
                Complete documentation is available at https://gitlab.com/encyclopaedia/jpath/-/blob/main/readme.md`,
		Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			switch len(args) {
			case 1:
				expr = strings.TrimSpace(args[0])
				json = ""
			case 2:
				expr = strings.TrimSpace(args[0])
				json = strings.TrimSpace(args[1])
			default:
			}
		},
	}

	// table output
	var table bool
	rootCmd.Flags().BoolVarP(&table, "table", "t", false, "print output as table")

	// parse input args
	if err := rootCmd.Execute(); err != nil {
		//_, _ = fmt.Fprintf(os.Stderr, "\n%s\n\n%s\n", err.Error(), rootCmd.UsageString())
		os.Exit(1)
	} else if expr == "" && json == "" {
		_, _ = fmt.Fprintf(os.Stderr, "\n%s\n\n%s\n", "no expression or json document provided", rootCmd.UsageString())
		os.Exit(1)
	}

	// parse the input json
	jsonb, err := input.ParseInputJson(json)
	if err != nil {
		common.ExitWithError(common.InvalidJson)
	}
	// parse the expression
	parsedOutput, err := parser.ProcessExpression(expr, jsonb)
	if err != nil {
		if strings.Contains(err.Error(), common.InvalidExpr.GetMsg()) {
			common.ExitWithError(common.InvalidExpr)
		} else {
			// fixme: handle more error types here
			common.ExitWithError(common.Success)
		}
	} else if parsedOutput == nil {
		os.Exit(int(common.Success))
	}

	// process output
	// fixme: the error handling below is sort of wonky. need more elegant handling
	if table {
		err = output.PrintJsonTable(parsedOutput)
		if err == nil {
			os.Exit(int(common.Success))
		}
	}
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "\n%s, printing as json\n", err.Error())
	}
	marshal := output.Prettify(parsedOutput, 2)
	fmt.Printf("%s\n", marshal)
}
