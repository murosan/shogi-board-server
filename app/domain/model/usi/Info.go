package usi

import "github.com/murosan/shogi-board-server/app/domain/model/shogi"

// Info represents the output from the shogi engine.
// We drop the info string, for now.
type Info struct {
	// depth, seldepth, time, nodes, nps, hashfull
	Values map[string]int `json:"values"`

	Score int `json:"score"`

	Moves []*shogi.Move `json:"moves"`
}
