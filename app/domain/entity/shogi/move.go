package shogi

// Move represents one move of the game.
type Move struct {
	// A place where the piece was moving
	Source *Point `json:"source"`

	// A place where the piece moved to
	Dest *Point `json:"dest"`

	PieceID Piece `json:"pieceId"`

	// true: promoted
	// false: not promoted or cannot promote
	IsPromoted bool `json:"isPromoted"`
}

// Point represents the location of the board.
// If the piece is captured one, the Point is Point{-1, -1}
type Point struct {
	Row    int `json:"row"`
	Column int `json:"column"`
}
