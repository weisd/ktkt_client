package quan

import (
	"fmt"
	"github.com/hprose/hprose-go/hprose"
)

type LogFilter struct {
}

func (LogFilter) InputFilter(data []byte, context hprose.Context) []byte {
	fmt.Println(string(data))
	return data
}

func (LogFilter) OutputFilter(data []byte, context hprose.Context) []byte {
	fmt.Println(string(data))
	return data
}
