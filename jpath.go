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

// jpathAppRunner the main usage of jPath application
// If jpath is used as a library in other applications, this is how it is
func jpathAppRunner(parsedJson []byte) {
	// tokenize the json doc
	tokenized, err := common.Tokenize(parsedJson)
	if err != nil {
		common.ExitWithMessage("error: " + err.Error())
	}
	// parse the expression
	parsedOutput, err := parser.ProcessExpression(common.Conf.Expr, tokenized)
	if err != nil {
		common.ExitWithMessage("error: " + err.Error())
	}
	// print the output
	err = output.PrintOutput(parsedOutput)
	if err != nil {
		common.ExitWithMessage(err.Error())
	}
}

func streamOutput(outChannel <-chan []byte) {
    for {
        doc := <-outChannel
        if len(doc) > 0 {
            jpathAppRunner(doc)
            common.Conf.Wg.Done()
        }
    }
}

func main() {
	// mandatory variables
	var json string

	// root command parser
	var rootCmd = &cobra.Command{
		Use:   "jpath <expression> <json>",
		Short: "analyzer for json data",
		Long: `An easy to use json filter to analyze json documents
                Complete documentation is available at https://github.com/prashanth-hegde/jpath`,
		Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			switch len(args) {
			case 1:
				common.Conf.Expr = strings.TrimSpace(args[0])
				json = ""
			case 2:
				common.Conf.Expr = strings.TrimSpace(args[0])
				json = strings.TrimSpace(args[1])
			default:
			}
		},
	}

	rootCmd.Flags().BoolVarP(&common.Conf.Table, "table", "t", false, "print output as table")
	rootCmd.Flags().BoolVarP(&common.Conf.Unwrap, "unwrap", "u", false, "unwrap the output from array")
	rootCmd.Flags().BoolVarP(&common.Conf.Compress, "compress", "c", false, "compress the output")

	// parse input args
	if err := rootCmd.Execute(); err != nil {
		//_, _ = fmt.Fprintf(os.Stderr, "\n%s\n\n%s\n", err.Error(), rootCmd.UsageString())
		os.Exit(1)
	} else if common.Conf.Expr == "" && json == "" {
		_, _ = fmt.Fprintf(os.Stderr, "\n%s\n\n%s\n", "no expression or json document provided", rootCmd.UsageString())
		os.Exit(1)
	}

	// if unwrap option is selected, create an output channel
	if common.Conf.Unwrap {
		go streamOutput(common.Conf.Channel)
	}

	// parse the input json
	jsonb, err := input.ParseInputJson(json)
	if err != nil {
		common.ExitWithMessage(err.Error())
	}
	jpathAppRunner(jsonb)
}
