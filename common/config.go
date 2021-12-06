package common

type JPathConf struct {
	Expr    string
	Unwrap  bool
	Table   bool
	Compact bool
	Indent  int
	Channel chan []byte
}

var Conf *JPathConf = &JPathConf{
	Expr:    ".",
	Unwrap:  false,
	Table:   false,
	Compact: false,
	Indent:  2,
}
