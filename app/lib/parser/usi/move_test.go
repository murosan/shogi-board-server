package usi

import (
	"reflect"
	"strings"
	"testing"

	"github.com/murosan/shogi-board-server/app/domain/model/shogi"
)

func TestParseMove(t *testing.T) {
	cases := []struct {
		in   string
		want *shogi.Move
		err  error
	}{
		{"7g7f",
			&shogi.Move{
				Source:     &shogi.Point{Row: 6, Column: 6},
				Dest:       &shogi.Point{Row: 5, Column: 6},
				PieceID:    0,
				IsPromoted: false,
			},
			nil,
		},
		{"8h2b+",
			&shogi.Move{
				Source:     &shogi.Point{Row: 7, Column: 7},
				Dest:       &shogi.Point{Row: 1, Column: 1},
				PieceID:    0,
				IsPromoted: true,
			},
			nil},
		{"G*5b",
			&shogi.Move{
				Source:     &shogi.Point{Row: -1, Column: -1},
				Dest:       &shogi.Point{Row: 1, Column: 4},
				PieceID:    5,
				IsPromoted: false,
			},
			nil,
		},
		{
			"s*5b",
			&shogi.Move{
				Source:     &shogi.Point{Row: -1, Column: -1},
				Dest:       &shogi.Point{Row: 1, Column: 4},
				PieceID:    -4,
				IsPromoted: false,
			},
			nil,
		},
		{"", &shogi.Move{}, errEmp},
		{"7g7z", &shogi.Move{}, errEmp},
		{"7gaf", &shogi.Move{}, errEmp},
		{"7g7$", &shogi.Move{}, errEmp},
		{"0g7a", &shogi.Move{}, errEmp},
		{"1x7a", &shogi.Move{}, errEmp},
		{"G*vb", &shogi.Move{}, errEmp},
		{"G*4z", &shogi.Move{}, errEmp},
		{"A*7a", &shogi.Move{}, errEmp},
	}

	for i, c := range cases {
		moveHelper(t, i, c.in, c.want, c.err)
	}
}

func moveHelper(t *testing.T, i int, in string, want *shogi.Move, err error) {
	t.Helper()
	res, e := ParseMove(in)
	msg := ""

	if (e == nil) != (err == nil) {
		msg = "The types of two errors does not match"
		moveErrorPrintHelper(t, i, msg, in, want, res)
		return
	}

	if (e != nil) && (err != nil) {
		if !strings.Contains(e.Error(), err.Error()) {
			msg = "Two errors are not nil as expected, but the error message is not correct."
			moveErrorPrintHelper(t, i, msg, in, want, res)
		}
		return
	}

	if (res == nil) != (want == nil) {
		msg = "The types of two *usi.Info does not match"
		moveErrorPrintHelper(t, i, msg, in, want, res)
		return
	}

	if res == nil {
		// won't come here
		t.Errorf("should not come here, because err is nil")
		return
	}

	if !reflect.DeepEqual(res, want) {
		msg = "The value was not as expected."
		moveErrorPrintHelper(t, i, msg, in, want, res)
	}
}

func moveErrorPrintHelper(
	t *testing.T,
	i int,
	msg, in string,
	expected, actual interface{},
) {
	t.Helper()
	t.Errorf(`[ParseMove] %s
Index:    %d
Input:    %s
Expected: %v
Actual:   %v
`, msg, i, in, expected, actual)
}
