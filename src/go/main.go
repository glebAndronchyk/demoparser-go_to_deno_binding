package main

import (
	"bytes"
	js_mappings_events "demo-parsing-binding/packages/js-mappings/events"
	js_mappings_game_state "demo-parsing-binding/packages/js-mappings/game-state"
	"fmt"
	"sync"
	"syscall/js"

	"github.com/markus-wa/demoinfocs-golang/v5/pkg/demoinfocs"
)

var (
	eventBridge *js_mappings_events.EventBridge
	parser      demoinfocs.Parser
	once        sync.Once
)

func main() {
	c := make(chan struct{}, 0)

	jsApiContainer := js.Global().Get("__PARSER_API__");

	jsApiContainer.Set("Create", js.FuncOf(Create))
	jsApiContainer.Set("GetStaticGameState", js.FuncOf(GetStaticGameState))
	jsApiContainer.Set("GetEntityState", js.FuncOf(GetEntityState))
	jsApiContainer.Set("ParseToEnd", js.FuncOf(ParseToEnd))
	jsApiContainer.Set("ParseNextFrame", js.FuncOf(ParseNextFrame))
	jsApiContainer.Set("Close", js.FuncOf(Close))
	jsApiContainer.Set("RegisterEvent", js.FuncOf(RegisterEvent))
	jsApiContainer.Set("UnregisterEvent", js.FuncOf(UnregisterEvent))

	fmt.Println("WASM Go Initialized")

	<-c
}

// #region public

func RegisterEvent(_ js.Value, args []js.Value) interface{} {
	config := args[1]
	onExecute := config.Get("onExecute")

	return registerEvent(args[0].String(), onExecute)
}

func UnregisterEvent(_ js.Value, args []js.Value) interface{} {
	unregisterEvent(args[0].Int())
	return nil
}

func Create(_ js.Value, args []js.Value) interface{} {
	create(args[0])
	return nil
}

func ParseNextFrame(_ js.Value, args []js.Value) interface{} {
	parser.ParseNextFrame()
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
	return getStaticGameState()
}

func GetEntityState(this js.Value, args []js.Value) interface{} {
	return getEntityState(args[0])
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

func registerEvent(eventName string, onExecute js.Value) js.Value {
	return js.ValueOf(eventBridge.RegisterEventByName(eventName, onExecute))
}

func unregisterEvent(handleId int) {
	eventBridge.UnregisterEventByHandleId(handleId);
}

func getEntityState(handle js.Value) js.Value {
	gameState := parser.GameState()
	entity := gameState.EntityByHandle(uint64(handle.Int()))

	mappings := js_mappings_game_state.NewEntityStateBindings(entity)

	if entity == nil {
		return js.Null()
	} else {
		return js.ValueOf(mappings.ToJS())
	}
}

func getStaticGameState() js.Value {
	gameState := parser.GameState()
	dto := js_mappings_game_state.NewGameStateDto(gameState)

	return js.ValueOf(dto)
}

func uint8ArrayToBytes(value js.Value) []byte {
	s := make([]byte, value.Get("byteLength").Int())
	js.CopyBytesToGo(s, value)
	return s
}
