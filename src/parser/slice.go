package parser

import (
	"github.com/pkg/errors"
	"jpath/common"
	"regexp"
	"strconv"
)

func Slice(expr string, json [][]byte) ([][]byte, error) {
	sliceRe := regexp.MustCompile(SliceRegex)
	jsonLength := int64(len(json))
	lower := int64(0)
	upper := jsonLength
	for _, line := range sliceRe.FindAllStringSubmatch(expr, -1) {
		var e error
		lower, e = strconv.ParseInt(line[1], 10, 64)
		if e != nil {
			lower = 0
			//return nil, errors.Wrapf(e, "error converting to int: %s", line[1])
		}
		upper, e = strconv.ParseInt(line[2], 10, 64)
		if e != nil {
			upper = jsonLength
			//return nil, errors.Wrapf(e, "error converting to int: %s", line[2])
		}
	}

	// negative slice numbers
	for lower < 0 {
		lower = jsonLength + lower
	}
	for upper < 0 {
		upper = jsonLength + upper
	}

	// out of range slices
	if lower > jsonLength {
		return nil, errors.Wrapf(common.NumberError.Error(), "index %d is longer than array length", lower)
	}
	if upper > jsonLength {
		return nil, errors.Wrapf(common.NumberError.Error(), "index %d is longer than array length", upper)
	}
	if lower > upper {
		// if the order is swapped, reverse the order and proceed
		temp := lower
		lower = upper
		upper = temp
	}

	return json[lower:upper], nil
}
