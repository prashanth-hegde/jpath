package output

import (
	jsonE "encoding/json"
	"fmt"
	parser "github.com/buger/jsonparser"
	"github.com/olekukonko/tablewriter"
	"github.com/pkg/errors"
	"jpath/common"
	"os"
)

// KeyProp defines the structure for individual columns in the table
// obtained by parsing the json document
type KeyProp struct {
	Key          string
	Type         parser.ValueType
	MaxLen       int
	PercentWidth int
}

// determineDataTypes tries to determine the data types for the given document
// Here are the steps of operation
// 01. peek at the first row
// 02. unmarshal the fields in the first row
// 03. get the contents of all fields in first row
// 04. if the field contents are object, ignore the field and move on
func determineDataTypes(jsonRow []byte) ([]KeyProp, error) {
	_, typ, _, e := parser.Get(jsonRow)
	if e != nil {
		return nil, errors.Wrap(e, fmt.Sprintf("%s - %s\n", common.UnknownDataType.GetMsg(), jsonRow))
	}
	var keyProps []KeyProp

	switch typ {
	case parser.Object:
		var fieldMap map[string]interface{}
		e = jsonE.Unmarshal(jsonRow, &fieldMap)
		if e != nil {
			return nil, errors.Wrap(e, common.UnmarshalError.GetMsg())
		}
		for k := range fieldMap {
			val, t, _, _ := parser.Get(jsonRow, k)
			if t != parser.Object && t != parser.Array {
				keyProps = append(keyProps, KeyProp{
					Key:    k,
					Type:   t,
					MaxLen: len(val),
				})
			}
		}
	default:

	}

	if keyProps == nil || len(keyProps) == 0 {
		return nil, errors.New(common.UnprintableData.GetMsg())
	}
	return keyProps, nil
}

func PrintJsonTable(json [][]byte) error {
	if len(json) == 0 {
		return nil
	}
	keyProps, e := determineDataTypes(json[0])
	if e != nil {
		return errors.Wrap(e, common.UnknownDataType.GetMsg())
	}
	printableJson := make([][]string, len(json))

	// determine headers
	var headers []string
	for _, h := range keyProps {
		headers = append(headers, h.Key)
	}

	// build the data table
	for i, doc := range json {
		printableJson[i] = make([]string, len(keyProps))
		for j, k := range keyProps {
			v, _, _, err := parser.Get(doc, k.Key)
			if err != nil {
				// when one field in one json document is not like others
				return errors.Wrap(err, fmt.Sprintf("%s for document[%d] %s\n",
					common.UnknownDataType.GetMsg(), i, k.Key))
			}
			printableJson[i][j] = string(v)
		}
	}

	fmt.Println("")
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(headers)
	table.AppendBulk(printableJson)
	// markdown style table
	table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	table.SetCenterSeparator("|")
	table.Render() // Send output
	fmt.Println("")

	return nil
}
