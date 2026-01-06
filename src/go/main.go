package main

import (
	"demo-parsing-binding/packages/dtos"
	parser_singleton "demo-parsing-binding/packages/singleton"
	"fmt"
	"syscall/js"
)

func main() {
	c := make(chan struct{}, 0)

	js.Global().Set("CreateParserInstance", js.FuncOf(CreateParserInstance))
	js.Global().Set("GetGameState", js.FuncOf(GetGameState))
	fmt.Println("WASM Go Initialized")

	<-c
}

// #region public

func CreateParserInstance(_ js.Value, args []js.Value) interface{} {
	parser_singleton.Init(args[0])
	return nil
}

func GetGameState(this js.Value, args []js.Value) interface{} {
	getGameState(args[0])
	return nil
}

// #region private

func getGameState(callback js.Value) {
	gameState := parser_singleton.GetInstance().GameState()
	dto := dtos.NewGameStateDto(gameState)

	callback.Invoke(js.ValueOf(dto))
}
