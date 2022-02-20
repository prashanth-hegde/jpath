package input

import (
	"bufio"
	"github.com/pkg/errors"
	"github.com/prashanth-hegde/jpath/common"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

// ParseInputJson reads the input and makes a json out of it
func ParseInputJson(json string) ([]byte, error) {
	var parsedb []byte
	if json == "" {
		p, e := readMultiDocumentArray(os.Stdin)
		if e != nil {
			return nil, e
		}
		parsedb = p
	} else if json[0] == '{' || json[0] == '[' {
		parsedb = []byte(json)
	} else if len(json) > 4 && json[0:4] == "http" {
		resp, e := httpGet(json, common.Conf.Headers)
		if e != nil {
			return nil, errors.Wrapf(e, "error making http request\n")
		}
		parsedb = resp
	} else {
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

	return parsedb, nil
}

// readMultiDocumentArray - get the json documents from stdin
// One limitation with the underlying library is that it parses fully
// formed json, and ignores the rest. For example, consider this input:
//
// {"name": "Abraham Lincoln", "quote": "Whatever you are, be a good one"}
// {"name": "Oscar Wilde", "quote": "I can resist everything except temptation"}
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
			// when an individual doc is read:
			if common.Conf.Unwrap {
				common.Conf.Channel <- document
				common.Conf.Wg.Add(1)
			} else {
				// if not streaming input, add it to array to process later (example: table out)
				documents = append(documents, document)
			}
			// ensure the document is cleared
			document = nil
		}
	}
	// if there are any pending operations, wait to finish them before proceeding
	common.Conf.Wg.Wait()

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

var client = &http.Client{
	CheckRedirect: func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	},
	Timeout: 10 * time.Second,
}

func httpGet(url string, headers []string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, errors.Wrapf(err, "err while creating http request %s", url)
	}
	for _, h := range headers {
		tokens := strings.SplitN(h, ":", 2)
		req.Header.Add(strings.TrimSpace(tokens[0]), strings.TrimSpace(tokens[1]))
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.Wrapf(err, "err while reading %s", url)
	}
	if resp.StatusCode >= 300 {
		return nil, errors.Errorf("non-success status code %d\n url %s\n", resp.StatusCode, url)
	}
	defer func() { _ = resp.Body.Close() }()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("error reading response body: %s", err.Error())
	}

	return body, nil
}
