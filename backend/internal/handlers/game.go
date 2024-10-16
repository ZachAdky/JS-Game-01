package handlers

import (
	"encoding/json"
	"net/http"
)

type GameState struct {
	Players []string `json:"players"`
	Score   int      `json:"score"`
}

func GameHandler(w http.ResponseWriter, r *http.Request) {
	state := GameState{
		Players: []string{"Player1"},
		Score:   100,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(state)
}
