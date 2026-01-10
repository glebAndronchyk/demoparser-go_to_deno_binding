package gamestate

import (
	"syscall/js"

	"github.com/markus-wa/demoinfocs-golang/v5/pkg/demoinfocs/sendtables"
)

type EntityStateBindings struct {
	OnDestroyFunc js.Func
}

func NewEntityStateBindings(entity sendtables.Entity) *EntityStateBindings {
	bindings := &EntityStateBindings{
		OnDestroyFunc: js.FuncOf(func(_ js.Value, args []js.Value) any {
			entity.OnDestroy(func() {
				callback := args[0]
				callback.Invoke()
			})
			return nil
		}),
	}

	return bindings
}

func (b *EntityStateBindings) ToJS() map[string]interface{} {
	return map[string]interface{}{
		"onDestroy": b.OnDestroyFunc,
		"release": js.FuncOf(func(_ js.Value, _ []js.Value) any {
			b.Release()
			return nil
		}),
	}
}

func (b *EntityStateBindings) Release() {
	b.OnDestroyFunc.Release()
}
