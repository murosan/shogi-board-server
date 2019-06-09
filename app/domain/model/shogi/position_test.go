package shogi

import (
	"bytes"
	"github.com/pkg/errors"
	"strings"
	"testing"
)

var (
	errEmpty = errors.New("")
)

func TestPosition_ToUSI(t *testing.T) {
	cases := []struct {
		in   *Position
		want []byte
		err  error
	}{
		{
			&Position{
				Pos: [][]int{
					{-2, -3, -4, -5, -8, 0, -4, -3, -2},
					{0, 0, 0, 0, 0, 0, -5, -6, 0},
					{-1, 0, -1, -1, -1, -1, 0, 0, -1},
					{0, 0, 0, 0, 0, 0, 7, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0, 0},
					{0, -7, 1, 0, 0, 0, 0, 0, 0},
					{1, 0, 0, 1, 1, 1, 1, 0, 1},
					{0, 6, 5, 0, 0, 0, 0, 0, 0},
					{2, 3, 4, 0, 8, 5, 4, 3, 2},
				},
				Cap0:      []int{3, 0, 0, 0, 0, 0, 0},
				Cap1:      []int{2, 0, 0, 0, 0, 0, 0},
				Turn:      -1,
				MoveCount: 100,
			},
			[]byte("position sfen lnsgk1snl/6gb1/p1pppp2p/6R2/9/1rP6/P2PPPP1P/1BG6/LNS1KGSNL w 3P2p 100"),
			nil,
		},
		{
			&Position{
				Pos: [][]int{
					{-2, -3, -4, -5, -8, -5, -4, -3, -2},
					{0, -7, 0, 0, 0, 0, 0, -6, 0},
					{-1, -1, -1, -1, -1, -1, -1, -1, -1},
					{0, 0, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0, 0},
					{1, 1, 1, 1, 1, 1, 1, 1, 1},
					{0, 6, 0, 0, 0, 0, 0, 7, 0},
					{2, 3, 4, 5, 8, 5, 4, 3, 2},
				},
				Cap0:      []int{0, 0, 0, 0, 0, 0, 0},
				Cap1:      []int{0, 0, 0, 0, 0, 0, 0},
				Turn:      1,
				MoveCount: 1,
			},
			[]byte("position sfen lnsgkgsnl/1r5b1/ppppppppp/9/9/9/PPPPPPPPP/1B5R1/LNSGKGSNL b - 1"),
			nil,
		},
		{
			&Position{
				Pos:       [][]int{},
				Cap0:      []int{0, 0, 0, 0, 0, 0, 0},
				Cap1:      []int{0, 0, 0, 0, 0, 0, 0},
				Turn:      1,
				MoveCount: 1,
			},
			nil,
			errEmpty,
		},
		{
			&Position{
				Pos: [][]int{
					{-2, -15, -4, -5, -8, -5, -4, -3, -2},
					{0, -7, 0, 0, 0, 0, 0, -6, 0},
					{-1, -1, -1, -1, -1, -1, -1, -1, -1},
					{0, 0, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0, 0},
					{1, 1, 1, 1, 1, 1, 1, 1, 1},
					{0, 6, 0, 0, 0, 0, 0, 7, 0},
					{2, 3, 4, 5, 8, 5, 4, 3, 2},
				},
				Cap0:      []int{0, 0, 0, 0, 0, 0, 0},
				Cap1:      []int{0, 0, 0, 0, 0, 0, 0},
				Turn:      1,
				MoveCount: 1,
			},
			nil,
			errEmpty,
		},
		{
			&Position{
				Pos: [][]int{
					{-2, -3, -4, -5, -8, -5, -4, -3, -2},
					{0, -7, 0, 0, 0, 0, 0, -6, 0},
					{-1, -1, -1, -1, -1, -1, -1, -1, -1},
					{0, 0, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0, 0},
					{1, 1, 1, 1, 1, 1, 1, 1, 1},
					{0, 6, 0, 0, 0, 0, 0, 7, 0},
					{2, 3, 4, 5, 8, 5, 4, 3, 2},
				},
				Cap0:      []int{0, 0, 0, 0, 0, 0, 0},
				Cap1:      []int{0, 0, 0, 0, 0, 0, 0},
				Turn:      0,
				MoveCount: 1,
			},
			nil,
			errEmpty,
		},
		{
			&Position{
				Pos: [][]int{
					{-2, -3, -4, -5, -8, -5, -4, -3, -2},
					{0, -7, 0, 0, 0, 0, 0, -6, 0},
					{-1, -1, -1, -1, -1, -1, -1, -1, -1},
					{0, 0, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0, 0},
					{1, 1, 1, 1, 1, 1, 1, 1, 1},
					{0, 6, 0, 0, 0, 0, 0, 7, 0},
					{2, 3, 4, 5, 8, 5, 4, 3, 2},
				},
				Cap0:      []int{0, 0, 0, 0, 0, 0, 0, 1},
				Cap1:      []int{0, 0, 0, 0, 0, 0, 0},
				Turn:      1,
				MoveCount: 1,
			},
			nil,
			errEmpty,
		},
		{
			&Position{
				Pos: [][]int{
					{-2, -3, -4, -5, -8, -5, -4, -3, -2},
					{0, -7, 0, 0, 0, 0, 0, -6, 0},
					{-1, -1, -1, -1, -1, -1, -1, -1, -1},
					{0, 0, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0, 0},
					{1, 1, 1, 1, 1, 1, 1, 1, 1},
					{0, 6, 0, 0, 0, 0, 0, 7, 0},
					{2, 3, 4, 5, 8, 5, 4, 3, 2},
				},
				Cap0:      []int{0, 0, 0, 0, 0, 0, 0},
				Cap1:      []int{0, 0, 0, 0, 0, 0, 0, 1},
				Turn:      1,
				MoveCount: 1,
			},
			nil,
			errEmpty,
		},
	}

	for i, c := range cases {
		b, err := c.in.ToUSI()

		if (c.err == nil) != (err == nil) {
			msg := "The types of two errors does not match"
			positionToUSIErrorPrintHelper(t, i, msg, c.in, err, c.err)
			return
		}

		if (c.err != nil) && (err != nil) {
			if !strings.Contains(err.Error(), c.err.Error()) {
				msg := "Two errors are not nil as expected, but the error message is not correct."
				positionToUSIErrorPrintHelper(t, i, msg, c.in, err, c.err)
			}
			return
		}

		if !bytes.Equal(b, c.want) {
			t.Errorf(`
Index:    %d
Expected: %s
Actual:   %s
`, i, string(c.want), string(b))
		}
	}
}

func positionToUSIErrorPrintHelper(
	t *testing.T,
	i int,
	msg string,
	in *Position,
	expected, actual interface{},
) {
	t.Helper()
	t.Errorf(`[Position.ToUSI] %s
Index:    %d
Input:    %v
Expected: %v
Actual:   %v
`, msg, i, in, expected, actual)
}
