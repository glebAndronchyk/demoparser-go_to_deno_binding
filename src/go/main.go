package main

import (
	"bytes"
	// "encoding/json"
	"fmt"
	"syscall/js"

	Parser "github.com/markus-wa/demoinfocs-golang/v5/pkg/demoinfocs"
)

func main() {
	c := make(chan struct{}, 0)

	REGISTER_API();
	fmt.Println("WASM Go Initialized");

	<-c;
}

func REGISTER_API() {
	js.Global().Set("get_game_state", js.FuncOf(EXPOSE_get_game_state))
}

func uint8ArrayToBytes(value js.Value) []byte {
	s := make([]byte, value.Get("byteLength").Int())
	js.CopyBytesToGo(s, value)
	return s
}

func EXPOSE_get_game_state(this js.Value, args []js.Value) interface {} {
	PRIVATE_get_game_state(args[0], args[1]);
	return nil;
}

func PRIVATE_get_game_state(data js.Value, callback js.Value) {
	b := bytes.NewBuffer(uint8ArrayToBytes(data));
	parser := Parser.NewParser(b);
	err := parser.ParseToEnd();

	if err != nil {
		fmt.Println("Error marshaling JSON:", err);
		return;
	}

	callback.Invoke(parser.GameState().IngameTick());
}
