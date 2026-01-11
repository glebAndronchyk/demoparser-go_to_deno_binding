package events

import (
	"syscall/js"

	"github.com/markus-wa/demoinfocs-golang/v5/pkg/demoinfocs"
	"github.com/markus-wa/demoinfocs-golang/v5/pkg/demoinfocs/events"
	dispatch "github.com/markus-wa/godispatch"
)

type EventBridge struct {
	parser    demoinfocs.Parser
	eventsMap map[int]dispatch.HandlerIdentifier
	nextEventId int
}

func NewEventBridge(parser demoinfocs.Parser) *EventBridge {
	return &EventBridge{parser: parser, eventsMap: make(map[int]dispatch.HandlerIdentifier), nextEventId: 0}
}

func (eventBridge *EventBridge) RegisterEventByName(name string, onExecute js.Value) int {
	switch name {
	case "frame-done":
		handle := eventBridge.parser.RegisterEventHandler(func(e events.FrameDone) {
			onExecute.Invoke(js.ValueOf(NewFrameDone(e)))
		})
		return eventBridge.registerHandleByNumericId(handle)
		// case "kill":
		// 	handlerId := eventBridge.parser.RegisterEventHandler(func (e events.Kill) {
		// 		onExecute.Invoke(js.ValueOf(NewFrameDone(e)))
		// 	})

		// 	return int(*handlerId)
	default:
		return -1
	}
}

func (eventBridge EventBridge) UnregisterEventByHandleId(handleId int) {
	dispatcherId := eventBridge.eventsMap[handleId]
	if dispatcherId != nil {
		eventBridge.parser.UnregisterEventHandler(dispatcherId)
		delete(eventBridge.eventsMap, handleId)
	}
}

func (eventBridge EventBridge) registerHandleByNumericId(handle dispatch.HandlerIdentifier) int {
	eventBridge.nextEventId++
	id := eventBridge.nextEventId
	eventBridge.eventsMap[id] = handle;

	return id;
}
