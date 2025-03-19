package parse

import (
	"bytes"
	"errors"
	"strings"

	"github.com/murosan/shogi-board-server/app/domain/entity/shogi"
	"github.com/murosan/shogi-board-server/app/domain/entity/usi"
	"golang.org/x/xerrors"
)

// ErrCustomUSIFormat means the input is a non-official usi format.
var ErrCustomUSIFormat = errors.New("custom usi format error")

// Move generates shogi.Move parsing from given string, and returns it
func Move(s string) (*shogi.Move, error) {
	s = strings.TrimSpace(s)

	// https://yaneuraou.yaneu.com/2017/06/16/拡張usiプロトコル-読み筋出力について/
	// ↑の一覧以外にも rep_win が来ることを確認したので rep_ でチェック
	if s == "win" || s == "resign" || strings.HasPrefix(s, "rep_") {
		return nil, ErrCustomUSIFormat
	}

	a := []byte(s)
	if len(a) < 4 {
		return nil, errors.New("insufficient length. input = " + s)
	}

	// is from captured.
	if bytes.IndexByte(a, '*') >= 0 {
		piece, err := Piece(usi.Piece(a[0]))
		if err != nil {
			msg := "failed to parse captured piece on Move. input = " + string(a[0]) + ": %w"
			return nil, xerrors.Errorf(msg, err)
		}

		src := &shogi.Point{Row: -1, Column: -1}

		row, err := parseRow(a[3])
		if err != nil {
			return nil, xerrors.Errorf(": %w", a[3], err)
		}

		col, err := parseColumn(a[2])
		if err != nil {
			return nil, xerrors.Errorf(": %w", a[2], err)
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
		return nil, xerrors.Errorf(": %w", a[1], err)
	}

	scol, err := parseColumn(a[0])
	if err != nil {
		return nil, xerrors.Errorf(": %w", a[0], err)
	}

	drow, err := parseRow(a[3])
	if err != nil {
		return nil, xerrors.Errorf(": %w", a[3], err)
	}

	dcol, err := parseColumn(a[2])
	if err != nil {
		return nil, xerrors.Errorf(": %w", a[2], err)
	}

	src := &shogi.Point{Row: srow, Column: scol}
	dst := &shogi.Point{Row: drow, Column: dcol}
	prm := len(a) == 5 && a[4] == '+'

	return &shogi.Move{
		Source:     src,
		Dest:       dst,
		PieceID:    0,
		IsPromoted: prm,
	}, nil
}

func parseColumn(b byte) (int, error) {
	if b < '1' || b > '9' {
		return 0, errors.New("invalid column number. input = " + string(b))
	}
	return int(b - '1'), nil
}

func parseRow(b byte) (int, error) {
	if b < 'a' || b > 'i' {
		return 0, errors.New("invalid row number. input = " + string(b))
	}
	return int(b - 'a'), nil
}
