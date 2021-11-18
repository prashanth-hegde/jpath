package output

import (
	jsonE "encoding/json"
	"github.com/TylerBrock/colorjson"
	"jpath/common"
)

func Prettify(json [][]byte, indent int) []byte {
	outJson := common.WrapIntoArray(json)
	f := colorjson.NewFormatter()
	f.Indent = indent
	var unmarshalInterface []interface{}
	handleError(jsonE.Unmarshal(outJson, &unmarshalInterface))
	marshalledOut, _ := f.Marshal(unmarshalInterface)
	return marshalledOut
}

func handleError(e error) {
	if e != nil {
		common.ExitWithError(common.MarshalError)
	}
}
