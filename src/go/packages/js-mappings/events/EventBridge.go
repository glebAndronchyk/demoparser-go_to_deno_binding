package events

import (
	"syscall/js"

	"github.com/markus-wa/demoinfocs-golang/v5/pkg/demoinfocs"
	"github.com/markus-wa/demoinfocs-golang/v5/pkg/demoinfocs/events"
	dispatch "github.com/markus-wa/godispatch"
)

type EventBridge struct {
	parser demoinfocs.Parser
}

func NewEventBridge(parser demoinfocs.Parser) *EventBridge {
	return &EventBridge{ parser: parser }
}

func (eventBridge *EventBridge) RegisterEventByName(name string, onExecute js.Value, registerDispose js.Value) {
	switch name {
		case "frame-done":
			handlerId := eventBridge.parser.RegisterEventHandler(func (e events.FrameDone) {
				onExecute.Invoke(js.ValueOf(NewFrameDone(e)))
			})
			registerDispose.Invoke(createJsDisposer(handlerId, eventBridge.parser))
		// case "kill":
		// 	handlerId := eventBridge.parser.RegisterEventHandler(func (e events.Kill) {
		// 		onExecute.Invoke(js.ValueOf(NewFrameDone(e)))
		// 	})
			
		// 	registerDispose.Invoke(createJsDisposer(handlerId, eventBridge.parser))
		}
}

func createJsDisposer(handlerId dispatch.HandlerIdentifier, parser demoinfocs.Parser) js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) any {
		parser.UnregisterEventHandler(handlerId)
		return nil
	})
}
