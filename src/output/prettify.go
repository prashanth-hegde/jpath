package output

import (
	"bytes"
	jsonE "encoding/json"
	"github.com/TylerBrock/colorjson"
	"jpath/common"
)

func Prettify(json [][]byte, indent int) []byte {
	joined := append(bytes.Join(json, []byte(",")), byte(']'))
	outJson := append([]byte("["), joined...)
	f := colorjson.NewFormatter()
	f.Indent = indent

	var unmarshalInterface []interface{}
	er := jsonE.Unmarshal(outJson, &unmarshalInterface)
	if er != nil {
		common.ExitWithError(common.MarshalError)
	}
	marshalledOutput, er := f.Marshal(unmarshalInterface)
	if er != nil {
		common.ExitWithError(common.MarshalError)
	}

	return marshalledOutput
}
