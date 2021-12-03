package input

import (
	"bufio"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"io"
	"jpath/common"
	"os"
	"time"
)

// ParseInputJson reads the input and makes a json out of it
func ParseInputJson(json string) ([][]byte, error) {
	start := time.Now()
	var parsedb []byte
	if json == "" {
		p, e := readMultiDocumentArray(os.Stdin)
		if e != nil {
			return nil, e
		}
		parsedb = p
	} else if json[0] == '{' || json[0] == '[' {
		log.Debugln("valid json, parsing it")
		parsedb = []byte(json)
	} else {
		log.Debugf("could be a file, checking %s\n", json)
		file, e := os.Open(json)
		defer closeFile(file)
		if e != nil {
			return nil, errors.Wrapf(e, "error while opening file %s\n", json)
		}
		p, e := readMultiDocumentArray(file)
		if e != nil {
			return nil, e
		}
		parsedb = p
	}

	log.Infof("time taken to process = %s", time.Since(start))
	return common.Tokenize(parsedb)
}

// readMultiDocumentArray - get the json documents from stdin
// One limitation with the underlying library is that it parses fully
// formed json, and ignores the rest. For example, consider this input:
//
// {"name": "Abraham Lincoln": "quote": "Whatever you are, be a good one"}
// {"name": "Oscar Wilde": "quote": "I can resist everything except temptation"}
//
// Notice the above input has a new line separator, and are two independent json objects
// The jsonparser library only takes the first one if we provide this as an input
// So we need to be a little proactive and wrap these objects into a json array
// There are multiple approaches for doing this
// 1. separate by new lines --> not a great assumption to make that input json will always be compressed
// 2. parse the input manually, wrap the input objects into an array and tokenize them --> more code, but reliable
func readMultiDocumentArray(file *os.File) ([]byte, error) {
	reader := bufio.NewReader(file) // file reader
	documents := make([][]byte, 0)  // output documents
	var parenthesis []bool          // parenthesis stack
	var document []byte             // current document

	for {
		// read each byte
		input, err := reader.ReadByte()
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			return nil, errors.Wrap(err, "error while reading stdin")
		}

		// trim whitespace if it's not part of the json
		if len(parenthesis) == 0 && (input == ' ' || input == '\n' || input == '\t') {
			continue
		} else {
			document = append(document, input)
		}

		// keep track of parenthesis
		switch input {
		case '[', '{':
			parenthesis = append(parenthesis, true)
		case ']', '}':
			if len(parenthesis) == 0 {
				return nil, common.InvalidJson.Error()
			}
			parenthesis = parenthesis[1:]
		default:
		}

		// append finalized documents
		if len(parenthesis) == 0 && len(document) > 0 {
			// if the last parenthesis is closed, add this to the array and clear the document
			documents = append(documents, document)
			document = nil
		}
	}
	// if parenthesis is not closed after reading the whole stdin, that's an error
	if len(parenthesis) > 0 {
		return nil, common.InvalidJson.Error()
	}

	return common.WrapIntoArray(documents), nil
}

func closeFile(f *os.File) {
	err := f.Close()
	if err != nil {
		common.ExitWithError(common.FileError)
	}
}
