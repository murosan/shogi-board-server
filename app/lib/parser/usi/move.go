package usi

import (
	"github.com/pkg/errors"
	"strconv"
	"strings"

	"github.com/murosan/shogi-board-server/app/domain/model/shogi"
)

// ParseMove generates shogi.Move parsing from given string, and returns it
func ParseMove(s string) (*shogi.Move, error) {
	a := strings.Split(strings.TrimSpace(s), "")

	if len(a) < 4 {
		return nil, errors.New("insufficient length. input = " + s)
	}

	// is from captured.
	if strings.Contains(s, "*") {
		piece, err := ParsePiece(a[0])
		if err != nil {
			msg := "failed to parse captured piece on ParseMove. input = " + a[0]
			return nil, errors.Wrap(err, msg)
		}

		src := &shogi.Point{Row: -1, Column: -1}

		row, err := parseRow(a[3])
		if err != nil {
			return nil, errors.Wrap(err, "failed to parse row. input = "+a[3])
		}

		col, err := parseColumn(a[2])
		if err != nil {
			return nil, errors.Wrap(err, "failed to parse column. input = "+a[2])
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
		return nil, errors.Wrap(err, "failed to parse row. input = "+a[1])
	}

	scol, err := parseColumn(a[0])
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse column. input = "+a[0])
	}

	drow, err := parseRow(a[3])
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse row. input = "+a[3])
	}

	dcol, err := parseColumn(a[2])
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse column. input = "+a[2])
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
		return 0, errors.Wrap(err, "is not a number. input = "+s)
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
