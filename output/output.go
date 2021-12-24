package output

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/prashanth-hegde/jpath/common"
)

// PrintOutput prints output that is non-tabular format
// Depending on the input flags, it opens up a matrix of output options
// example json: [{"one":"01", "two":"02"},{"three":"03", "four":"04"}]
// -------------------+-------------------------------+------------------------------------------------------|
//                    |  --unwrap=true                |           --unwrap=false                             |
// -------------------|-------------------------------|------------------------------------------------------|
// --compact=true     | {"one":"01","two":"02"}       | [{"one":"01","two":"02"},{"three":"03","four":"04"}] |
//                    | {"three":"03","four":"04"}    |                                                      |
// -------------------|-------------------------------|------------------------------------------------------|
// --compact=false    | {                             | [                                                    |
//                    |   "one": "01",                |   {                                                  |
//                    |   "two": "02"                 |     "one": "01",                                     |
//                    | }                             |     "two": "02"                                      |
//                    | {                             |   },                                                 |
//                    |   "three": "03",              |   {                                                  |
//                    |   "four": "04"                |     "three": "03",                                   |
//                    | }                             |     "four": "04"                                     |
//                    |                               |   }                                                  |
//                    |                               | ]                                                    |
// -------------------+-------------------------------+------------------------------------------------------|
func PrintOutput(parsed [][]byte) error {
	if common.Conf.Table {
		e := PrintJsonTable(parsed)
		return errors.Wrap(e, common.UnprintableTable.GetMsg())
	} else if common.Conf.Compress && common.Conf.Unwrap {
		for i := 0; i < len(parsed); i++ {
			fmt.Printf("%s\n", compact(parsed[i]))
		}
	} else if common.Conf.Compress && !common.Conf.Unwrap {
		wrapped := common.WrapIntoArray(parsed)
		fmt.Printf("%s\n", compact(wrapped))
	} else if !common.Conf.Compress && common.Conf.Unwrap {
		op, e := ColoredUnwrappedOutput(parsed)
		if e != nil {
			return errors.Wrap(e, "error while printing uncompressed unwrapped output")
		}
		for i := 0; i < len(op); i++ {
			fmt.Printf("%s\n", op[i])
		}
	} else {
		// todo: error handling
		out, e := ColoredOutput(parsed)
		if e != nil {
			return errors.Wrap(e, "error while printing uncompressed wrapped output")
		}
		fmt.Printf("%s\n", out)
	}
	return nil
}
