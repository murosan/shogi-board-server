package parse

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/murosan/shogi-board-server/app/domain/entity/shogi"
	"github.com/murosan/shogi-board-server/app/domain/entity/usi"
)

// Move generates shogi.Move parsing from given string, and returns it
func Move(s string) (*shogi.Move, error) {
	a := strings.Split(strings.TrimSpace(s), "")

	if len(a) < 4 {
		return nil, errors.New("insufficient length. input = " + s)
	}

	// is from captured.
	if strings.Contains(s, "*") {
		piece, err := Piece(usi.Piece(a[0]))
		if err != nil {
			msg := "failed to parse captured piece on Move. input = " + a[0] + ": %w"
			return nil, fmt.Errorf(msg, err)
		}

		src := &shogi.Point{Row: -1, Column: -1}

		row, err := parseRow(a[3])
		if err != nil {
			return nil, fmt.Errorf("failed to parse row. input = %s: %w", a[3], err)
		}

		col, err := parseColumn(a[2])
		if err != nil {
			return nil, fmt.Errorf("failed to parse column. input = %s: %w", a[2], err)
		}

		dst := &shogi.Point{Row: row, Column: col}
		return &shogi.Move{
			Source:     src,
			Dest:       dst,
			PieceID:    piece,
			IsPromoted: false,
		}, nil
	}

	srow, err := parseRow(a[1])
	if err != nil {
		return nil, fmt.Errorf("failed to parse row. input = %s: %w", a[1], err)
	}

	scol, err := parseColumn(a[0])
	if err != nil {
		return nil, fmt.Errorf("failed to parse column. input = %s: %w", a[0], err)
	}

	drow, err := parseRow(a[3])
	if err != nil {
		return nil, fmt.Errorf("failed to parse row. input = %s: %w", a[3], err)
	}

	dcol, err := parseColumn(a[2])
	if err != nil {
		return nil, fmt.Errorf("failed to parse column. input = %s: %w", a[2], err)
	}

	src := &shogi.Point{Row: srow, Column: scol}
	dst := &shogi.Point{Row: drow, Column: dcol}
	prm := len(a) == 5 && a[4] == "+"

	return &shogi.Move{
		Source:     src,
		Dest:       dst,
		PieceID:    0,
		IsPromoted: prm,
	}, nil
}

func parseColumn(s string) (int, error) {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("is not a number. input = %s: %w", s, err)
	}
	if i < 1 || i > 9 {
		return 0, errors.New("invalid column number. input = " + s)
	}

	// decrease 1. because given string is in 1-9,
	// but our column should be in 0-8
	return i - 1, nil
}

func parseRow(s string) (int, error) {
	r := []rune(s)[0]
	if r < 97 || r > 105 {
		return 0, errors.New("invalid row number. input = " + s)
	}
	return int(r - 97), nil
}
