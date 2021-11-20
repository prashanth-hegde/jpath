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
		Use:   "jpath",
		Short: "analyzer for json data",
		Long: `An easy to use json filter to analyze json documents
                Complete documentation is available at https://gitlab.com/encyclopaedia/jpath/-/blob/main/readme.md`,
		Run: func(cmd *cobra.Command, args []string) {
			expr = strings.TrimSpace(args[0])
			json = strings.TrimSpace(args[1])
		},
	}

	// table output
	var table bool
	rootCmd.Flags().BoolVarP(&table, "table", "t", false, "print output in table format")

	// parse input args
	if err := rootCmd.Execute(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "\n%s\n\n%s\n", err.Error(), rootCmd.UsageString())
		os.Exit(1)
	} else if expr == "" && json == "" {
		_, _ = fmt.Fprintf(os.Stderr, "\n%s\n\n%s\n", "no expression or json document provided", rootCmd.UsageString())
		os.Exit(1)
	}

	// parse the input
	jsonb := input.ParseInputJson(json)
	parsedOutput, err := parser.ProcessExpression(expr, jsonb)
	if err != nil {
		if strings.Contains(err.Error(), common.InvalidExpr.GetMsg()) {
			_, _ = fmt.Fprintf(os.Stderr, "\n%s\n", common.InvalidExpr.GetMsg())
			os.Exit(int(common.InvalidExpr))
		}
	} else if parsedOutput == nil {
		os.Exit(int(common.Success))
	}

	// process output
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
