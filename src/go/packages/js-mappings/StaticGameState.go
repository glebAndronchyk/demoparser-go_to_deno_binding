package js_mappings

import "github.com/markus-wa/demoinfocs-golang/v5/pkg/demoinfocs"

func NewGameStateDto(gs demoinfocs.GameState) map[string]interface{} {
	result := map[string]interface{}{
		"ingameTick":         gs.IngameTick(),
		"gamePhase":          int(gs.GamePhase()),
		"overtimeCount":      gs.OvertimeCount(),
		"isMatchStarted":     gs.IsMatchStarted(),
		"totalRoundsPlayed":  gs.TotalRoundsPlayed(),
		"isWarmupPeriod":     gs.IsWarmupPeriod(),
		"isFreezetimePeriod": gs.IsFreezetimePeriod(),
	}

	return result
}
