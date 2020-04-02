package convert

import (
	"errors"
	"fmt"
	"strings"

	"github.com/murosan/shogi-board-server/app/domain/entity/shogi"
	"github.com/murosan/shogi-board-server/app/domain/entity/usi"
)

// Position converts shogi.Position to usi-position command bytes.
func Position(p *shogi.Position) ([]byte, error) {
	// for safety
	if len(p.Pos) != 9 {
		return nil, errors.New("length of position.Pos is not 9")
	}
	if len(p.Cap0) != 7 || len(p.Cap1) != 7 {
		return nil, errors.New("length of position.Cap* is not 7")
	}

	rows := make([]string, 9)

	for i, r := range p.Pos {
		usir, err := rowToUSI(r)
		if err != nil {
			return nil, err
		}
		rows[i] = usir
	}

	s := []byte("position sfen " + strings.Join(rows, "/"))
	if p.Turn == shogi.Sente {
		s = append(s, []byte(" "+usi.Sente+" ")...)
	} else if p.Turn == shogi.Gote {
		s = append(s, []byte(" "+usi.Gote+" ")...)
	} else {
		return nil, errors.New("unknown turn number. Turn = " + fmt.Sprint(p.Turn))
	}

	caps := []byte("")

	for i, c := range p.Cap0 {
		if c != 0 {
			p, err := Piece(shogi.Piece(i + 1))
			if err != nil {
				return nil, err
			}
			caps = append(caps, []byte(fmt.Sprint(c)+p.String())...)
		}
	}

	for i, c := range p.Cap1 {
		if c != 0 {
			p, err := Piece(shogi.Piece(-i - 1))
			if err != nil {
				return nil, err
			}
			caps = append(caps, []byte(fmt.Sprint(c)+p.String())...)
		}
	}

	if len(caps) == 0 {
		s = append(s, []byte("-")...)
	} else {
		s = append(s, caps...)
	}

	return append(s, []byte(" "+fmt.Sprint(p.MoveCount))...), nil
}

func rowToUSI(row []int) (s string, err error) {
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
		if id == shogi.Empty.ToInt() {
			emp++
			continue
		}

		appendCount()

		p, err := Piece(shogi.Piece(id))
		if err != nil {
			return "", fmt.Errorf("failed to convert piece to usi: %w", err)
		}

		s += p.String()
	}

	appendCount()

	return
}
