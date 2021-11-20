package common

import (
	"fmt"
	"os"
)

type ErrorCode int

const (
	Success ErrorCode = iota
	InvalidJson
	InvalidExpr
	FileError
	UnknownDataType
	UnmarshalError
	UnprintableData
)

func ExitWithError(code ErrorCode) {
	_, _ = fmt.Fprintf(os.Stderr, "%s\n", ErrorMessage[code])
	os.Exit(int(code))
}

var ErrorMessage = map[ErrorCode]string{
	Success:         "",
	InvalidJson:     "invalid json, aborting",
	InvalidExpr:     "invalid expression, aborting",
	FileError:       "file not json or unreadable, aborting",
	UnknownDataType: "unknown data type in json object",
	UnmarshalError:  "error while unmarshalling json into map",
	UnprintableData: "data not printable",
}

func (e ErrorCode) GetMsg() string {
	return ErrorMessage[e]
}

func (e ErrorCode) ExitWithMessage() {
	_, _ = fmt.Fprintf(os.Stderr, "%s\n", e.GetMsg())
	os.Exit(int(e))
}
