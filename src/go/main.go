package main

import (
	js_mappings_events "demo-parsing-binding/packages/js-mappings/events"
	js_mappings_game_state "demo-parsing-binding/packages/js-mappings/game-state"
	"fmt"
	"syscall/js"
	"bytes"
	"sync"

	"github.com/markus-wa/demoinfocs-golang/v5/pkg/demoinfocs"
)

var (
	eventBridge *js_mappings_events.EventBridge
	parser demoinfocs.Parser
	once     sync.Once
)

func main() {
	c := make(chan struct{}, 0)

	js.Global().Set("Create", js.FuncOf(Create))
	js.Global().Set("GetStaticGameState", js.FuncOf(GetStaticGameState))
	js.Global().Set("GetEntityState", js.FuncOf(GetEntityState))
	js.Global().Set("ParseToEnd", js.FuncOf(ParseToEnd))
	js.Global().Set("Close", js.FuncOf(Close))
	js.Global().Set("RegisterEvent", js.FuncOf(RegisterEvent))

	fmt.Println("WASM Go Initialized")

	<-c
}

// #region public

func RegisterEvent(_ js.Value, args []js.Value) interface{} {
	config := args[1];
	onExecute := config.Get("onExecute")
	onDispose := config.Get("onDispose")

	registerEvent(args[0].String(), onExecute, onDispose)
	return nil
}

func Create(_ js.Value, args []js.Value) interface{} {
	create(args[0])
	return nil
}

func ParseToEnd(_ js.Value, _ []js.Value) interface{} {
	parser.ParseToEnd()
	return nil
}

func Close(_ js.Value, _ []js.Value) interface{} {
	parser.Close()
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

func create(data js.Value) {
	once.Do(func() {
		b := bytes.NewBuffer(uint8ArrayToBytes(data))
		p := demoinfocs.NewParser(b)
		parser = p
		eventBridge = js_mappings_events.NewEventBridge(parser)
	})
}

func registerEvent(eventName string, onExecute js.Value, onDispose js.Value) {
	eventBridge.RegisterEventByName(eventName, onExecute, onDispose)
}

func getEntityState(callback js.Value, handle js.Value) {
	gameState := parser.GameState()
	entity := gameState.EntityByHandle(uint64(handle.Int()))

	mappings := js_mappings_game_state.NewEntityStateBindings(entity)

	if entity == nil {
		callback.Invoke(js.Null())
	} else {
		callback.Invoke(js.ValueOf(mappings.ToJS()))
	}
}

func getStaticGameState(callback js.Value) {
	gameState := parser.GameState()
	dto := js_mappings_game_state.NewGameStateDto(gameState)

	callback.Invoke(js.ValueOf(dto))
}

func uint8ArrayToBytes(value js.Value) []byte {
	s := make([]byte, value.Get("byteLength").Int())
	js.CopyBytesToGo(s, value)
	return s
}
