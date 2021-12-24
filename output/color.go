package output

import (
	"bytes"
	jsonE "encoding/json"
	"fmt"
	"github.com/TylerBrock/colorjson"
	parser "github.com/buger/jsonparser"
	"github.com/pkg/errors"
	"github.com/prashanth-hegde/jpath/common"
)

func ColoredOutput(json [][]byte) ([]byte, error) {
	if len(json) == 0 {
		return make([]byte, 0), nil
	}
	outJson := common.WrapIntoArray(json)
	f := colorjson.NewFormatter()
	f.Indent = common.Conf.Indent

	_, t, _, e := parser.Get(json[0])
	if e != nil || t != parser.Object {
		// output is not an object, salvage it by treating it as an array of strings
		outJson = append(append([]byte("[\""), bytes.Join(json, []byte(`","`))...), []byte("\"]")...)
	}

	var unmarshalInterface []interface{}
	e = jsonE.Unmarshal(outJson, &unmarshalInterface)
	if e != nil {
		return nil, errors.Wrap(e, common.UnmarshalError.GetMsg())
	}
	marshalledOut, e := f.Marshal(unmarshalInterface)
	if e != nil {
		return nil, errors.Wrap(e, common.UnmarshalError.GetMsg())
	}
	return marshalledOut, nil
}

func ColoredUnwrappedOutput(json [][]byte) ([][]byte, error) {
	if len(json) == 0 {
		return make([][]byte, 0), nil
	}
	f := colorjson.NewFormatter()
	f.Indent = common.Conf.Indent
	var marshalledOut [][]byte // output

	_, t, _, _ := parser.Get(json[0])
	for _, doc := range json {
		switch t {
		case parser.Object:
			var unmarshalInterface map[string]interface{}
			er := jsonE.Unmarshal(doc, &unmarshalInterface)
			if er != nil {
				return nil, errors.Wrap(er, "error while unmarshalling object")
			}
			m, er := f.Marshal(unmarshalInterface)
			if er != nil {
				return nil, errors.Wrap(er, "error while marshalling object")
			}
			marshalledOut = append(marshalledOut, m)
		case parser.Array:
			var unmarshalInterface []interface{}
			er := jsonE.Unmarshal(doc, &unmarshalInterface)
			if er != nil {
				return nil, errors.Wrap(er, "error while unmarshalling object")
			}
			m, er := f.Marshal(unmarshalInterface)
			if er != nil {
				return nil, errors.Wrap(er, "error while marshalling object")
			}
			marshalledOut = append(marshalledOut, m)
		default:
			marshalledOut = append(marshalledOut, []byte(fmt.Sprintf("%s", doc)))
		}
	}
	return marshalledOut, nil
}
