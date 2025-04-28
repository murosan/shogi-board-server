package shogi

// Position represents positions of pieces.
type Position struct {
	// positions of pieces on the board
	Pos [][]int `json:"pos"`

	// captures of the first player
	// each element represents the number of pieces.
	// [Fu, Kyou, Kei, Gin, Kei, Kin, Kaku, Hisha]
	Cap0 []int `json:"cap0"`

	// captures of the second player
	// each element represents the number of pieces.
	// [Fu, Kyou, Kei, Gin, Kei, Kin, Kaku, Hisha]
	Cap1 []int `json:"cap1"`

	Turn Turn `json:"turn"`

	// MoveCount is the count of moves.
	// The count of initial positions is 0.
	MoveCount int `json:"moveCount"`
}
