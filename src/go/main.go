package main

import (
	js_mappings "demo-parsing-binding/packages/js-mappings"
	parser_singleton "demo-parsing-binding/packages/singleton"
	"fmt"
	"syscall/js"
)

func main() {
	c := make(chan struct{}, 0)

	js.Global().Set("CreateParserInstance", js.FuncOf(CreateParserInstance))
	js.Global().Set("GetStaticGameState", js.FuncOf(GetStaticGameState))
	js.Global().Set("GetEntityState", js.FuncOf(GetEntityState))
	fmt.Println("WASM Go Initialized")

	<-c
}

// #region public

func CreateParserInstance(_ js.Value, args []js.Value) interface{} {
	parser_singleton.Init(args[0])
	return nil
}

func GetStaticGameState(this js.Value, args []js.Value) interface{} {
	getStaticGameState(args[0])
	return nil
}

func GetEntityState(this js.Value, args []js.Value) interface{} {
	getEntityState(args[0], args[1])
	return nil
}

// #region private

func getEntityState(callback js.Value, handle js.Value) {
	gameState := parser_singleton.GetInstance().GameState()
	entity := gameState.EntityByHandle(uint64(handle.Int()))

	mappings := js_mappings.NewEntityStateBindings(entity)

	if entity == nil {
		callback.Invoke(js.Null())
	} else {
		callback.Invoke(js.ValueOf(mappings.ToJS()))
	}
}

func getStaticGameState(callback js.Value) {
	gameState := parser_singleton.GetInstance().GameState()
	dto := js_mappings.NewGameStateDto(gameState)

	callback.Invoke(js.ValueOf(dto))
}
