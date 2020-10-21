package main

import (
	"strconv"
	"strings"
)

func main() {
	zhToUnicode([]byte("\\u5931\\u8d25"))
}

func zhToUnicode(raw []byte) ([]byte, error) {
	q := strconv.Quote(string(raw))
	r := strings.Replace(q, `\\u`, `\u`, -1)
	str, err := strconv.Unquote(r)
	if err != nil {
		return nil, err
	}
	return []byte(str), nil
}
