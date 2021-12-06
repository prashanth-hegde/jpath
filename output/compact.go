package output

func compact(parsed []byte) []byte {
	out := make([]byte, 0)
	parenthesis := make([]bool, 0)
	closedQuote := true

	for _, b := range parsed {
		switch b {
		case '[', '{':
			parenthesis = append(parenthesis, true)
		case ']', '}':
			parenthesis = parenthesis[1:]
		case '"':
			closedQuote = !closedQuote
		}

		// trim whitespace if it's not part of the key or value
		if closedQuote && (b == ' ' || b == '\n' || b == '\t') {
			continue
		} else {
			out = append(out, b)
		}
	}
	return out
}
