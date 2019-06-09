package shogi

import (
	"fmt"
	"github.com/pkg/errors"
	"strings"
)

// Position represents positions of pieces.
type Position struct {
	// positions of pieces on the board
	Pos [][]int `json:"pos"`

	// captures of the first player
	// each elements represents the number of piece.
	// [Fu, Kyou, Kei, Gin, Kei, Kin, Kaku, Hisha]
	Cap0 []int `json:"cap0"`

	// captures of the second player
	// each elements represents the number of piece.
	// [Fu, Kyou, Kei, Gin, Kei, Kin, Kaku, Hisha]
	Cap1 []int `json:"cap1"`

	Turn Turn `json:"turn"`

	// MoveCount is the count of moves.
	// The count of initial positions is 0.
	MoveCount int `json:"moveCount"`
}

// ToUSI converts Position to USI position command.
func (p *Position) ToUSI() ([]byte, error) {
	// for safety
	if len(p.Pos) != 9 {
		return nil, errors.New("length of position.Pos is not 9")
	}

	rows := make([]string, 9)

	for i, r := range p.Pos {
		usir, err := p.rowToUSI(r)
		if err != nil {
			return nil, err
		}
		rows[i] = usir
	}

	s := []byte("position sfen " + strings.Join(rows, "/"))
	if p.Turn == Sente {
		s = append(s, []byte(" b ")...)
	} else if p.Turn == Gote {
		s = append(s, []byte(" w ")...)
	} else {
		return nil, errors.New("unknown turn number. Turn = " + fmt.Sprint(p.Turn))
	}

	caps := []byte("")

	for i, c := range p.Cap0 {
		if c != 0 {
			p, err := PieceToUSI(i + 1)
			if err != nil {
				return nil, err
			}
			caps = append(caps, []byte(fmt.Sprint(c)+p)...)
		}
	}

	for i, c := range p.Cap1 {
		if c != 0 {
			p, err := PieceToUSI(-i - 1)
			if err != nil {
				return nil, err
			}
			caps = append(caps, []byte(fmt.Sprint(c)+p)...)
		}
	}

	if len(caps) == 0 {
		s = append(s, []byte("-")...)
	} else {
		s = append(s, caps...)
	}

	return append(s, []byte(" "+fmt.Sprint(p.MoveCount))...), nil
}

func (p *Position) rowToUSI(row []int) (s string, err error) {
	emp := 0

	// if the count of empty cells is not 0,
	// then append the count to result.
	appendCount := func() {
		if emp != 0 {
			s += fmt.Sprint(emp)
			emp = 0
		}
	}

	for _, id := range row {
		// if the piece is empty cell, then count up.
		if id == Empty {
			emp++
			continue
		}

		appendCount()

		p, err := PieceToUSI(id)
		if err != nil {
			return "", errors.Wrap(err, "failed to convert piece to USIPiece")
		}

		s += p
	}

	appendCount()

	return
}
