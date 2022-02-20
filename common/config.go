package common

import (
	"sync"
)

type JPathConf struct {
	Expr     string
	Unwrap   bool
	Table    bool
	Compress bool
	Indent   int
	Channel  chan []byte
	Wg       sync.WaitGroup
	Headers  []string
}

var Conf = &JPathConf{
	Expr:     ".",
	Unwrap:   false,
	Table:    false,
	Compress: false,
	Indent:   2,
	Channel:  make(chan []byte, 1000),
	Headers:  nil,
}
