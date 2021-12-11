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

type MatcherExp struct {
	KeyReg       *regexp.Regexp
	FilterReg    *regexp.Regexp
	SelectionReg *regexp.Regexp
	CountReg     *regexp.Regexp
	SliceReg     *regexp.Regexp
}

var Matcher = &MatcherExp{
	KeyReg:       regexp.MustCompile(keyRegex),
	FilterReg:    regexp.MustCompile(filterRegex),
	SelectionReg: regexp.MustCompile(selectionRegex),
	CountReg:     regexp.MustCompile(countRegex),
	SliceReg:     regexp.MustCompile(sliceRegex),
}
