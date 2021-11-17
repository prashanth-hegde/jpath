package input

import (
	"bufio"
	"bytes"
	log "github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"jpath-go/common"
	"os"
)

func ParseInputJson(json string) [][]byte {
	var parsedb []byte
	if json == "" {
		log.Debugln("json input not provided, fetching from stdin")
		reader := bufio.NewReader(os.Stdin)
		var out []byte
		for {
			ru, _, err := reader.ReadRune()
			if err != nil && err == io.EOF {
				break
			}
			out = append(out, byte(ru))
		}
		parsedb = out
	} else if json[0] == '{' || json[0] == '[' {
		log.Debugln("valid json, parsing it")
		parsedb = []byte(json)
	} else {
		log.Debugf("could be a file, checking %s\n", json)
		data, err := ioutil.ReadFile(json)
		if err != nil {
			common.ExitWithError(common.FileError)
		}
		parsedb = bytes.TrimSpace(data)
	}

	return common.Tokenize(parsedb)
}
