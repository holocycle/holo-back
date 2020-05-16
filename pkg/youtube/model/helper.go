package model

import (
	"strconv"
	"strings"
)

type JSONInt64 struct {
	Value int64
}

func (i *JSONInt64) UnmarshalJSON(json []byte) error {
	s := string(json)                       // "\"123\""
	s = strings.ReplaceAll(s, "\"", "")     // "123"
	val, err := strconv.ParseInt(s, 10, 64) // 123
	if err != nil {
		return err
	}
	i.Value = val
	return nil
}
