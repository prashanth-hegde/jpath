package common

import (
	"fmt"
	"os"
)

type ErrorCode int8

const (
	Success ErrorCode = iota
	InvalidJson
	InvalidExpr
	FileError
	MarshalError
)

func ExitWithMessage(msg string) {
	_, _ = fmt.Fprintf(os.Stderr, "%s\n", msg)
	os.Exit(2)
}

func ExitWithError(code ErrorCode) {
	switch code {
	case InvalidJson:
		ExitWithMessage("invalid json, aborting")
	case InvalidExpr:
		ExitWithMessage("invalid expression, aborting")
	case FileError:
		ExitWithMessage("file not json or unreadable, aborting")
	case MarshalError:
		ExitWithMessage("error while rendering output, must be a bug, sorry :(")
	default:
		// no op
	}
}
