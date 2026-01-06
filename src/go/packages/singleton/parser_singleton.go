package singleton

import (
	"bytes"
	"sync"
	"syscall/js"

	"github.com/markus-wa/demoinfocs-golang/v5/pkg/demoinfocs"
)

var (
	instance *demoinfocs.Parser
	once     sync.Once
)

func uint8ArrayToBytes(value js.Value) []byte {
	s := make([]byte, value.Get("byteLength").Int())
	js.CopyBytesToGo(s, value)
	return s
}

func Init(data js.Value) {
	once.Do(func() {
		b := bytes.NewBuffer(uint8ArrayToBytes(data))
		p := demoinfocs.NewParser(b)
		instance = &p
	})
}

func GetInstance() demoinfocs.Parser {
	if instance == nil {
		panic("instance not initialized")
	}

	return *instance
}
