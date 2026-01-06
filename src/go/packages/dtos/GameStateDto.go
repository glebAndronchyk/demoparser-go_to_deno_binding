package dtos

import "github.com/markus-wa/demoinfocs-golang/v5/pkg/demoinfocs"

func NewGameStateDto(gs demoinfocs.GameState) map[string]interface{} {
	result := map[string]interface{}{
		"ingameTick": gs.IngameTick(),
		"gamePhase":  int(gs.GamePhase()),
	}

	return result
}
