package output

import (
	"bytes"
	jsonE "encoding/json"
	"github.com/TylerBrock/colorjson"
	parser "github.com/buger/jsonparser"
	"jpath/common"
)

func Prettify(json [][]byte, indent int) []byte {
	outJson := common.WrapIntoArray(json)
	f := colorjson.NewFormatter()
	f.Indent = indent

	_, t, _, e := parser.Get(json[0])
	if e != nil || t != parser.Object {
		// output is not an object, salvage it by treating it as an array of strings
		outJson = append(append([]byte("[\""), bytes.Join(json, []byte(`","`))...), []byte("\"]")...)
	}

	var unmarshalInterface []interface{}
	HandleError(jsonE.Unmarshal(outJson, &unmarshalInterface))
	marshalledOut, _ := f.Marshal(unmarshalInterface)
	return marshalledOut
}

func HandleError(e error) {
	if e != nil {
		common.ExitWithError(common.UnmarshalError)
	}
}
