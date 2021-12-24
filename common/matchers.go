package common

import (
	"regexp"
)

const (
	keyRegex       string = `^[a-zA-Z0-9_-]+$`
	filterRegex           = `^(\w+)?\[([\w.]+)(=|!=|<=?|>=?|~)(.*)]$`
	selectionRegex        = `^{([\w,.-]+)}$`
	sliceRegex            = `^(\w+)?\[(-?\d+)?:(-?\d+)?\]$`
	countRegex            = `^#$`
)

var (
	KeyReg       = regexp.MustCompile(keyRegex)
	FilterReg    = regexp.MustCompile(filterRegex)
	SelectionReg = regexp.MustCompile(selectionRegex)
	CountReg     = regexp.MustCompile(countRegex)
	SliceReg     = regexp.MustCompile(sliceRegex)
)
