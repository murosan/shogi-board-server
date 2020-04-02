package parse

import (
	"reflect"
	"strings"
	"testing"

	"github.com/murosan/shogi-board-server/app/domain/entity/shogi"
	"github.com/murosan/shogi-board-server/app/domain/entity/usi"
)

func TestInfo(t *testing.T) {
	cases := []struct {
		in   string
		want *usi.Info
		mpv  int
		err  error
	}{
		{
			"info time 1141 depth 3 seldepth 3 nodes 135125 score cp -1521 pv 3a3b L*4h 4c4d",
			&usi.Info{
				Values: map[string]int{
					time:     1141,
					depth:    3,
					selDepth: 3,
					nodes:    135125,
				},
				Score: -1521,
				Moves: []*shogi.Move{
					{
						Source:     &shogi.Point{Row: 0, Column: 2},
						Dest:       &shogi.Point{Row: 1, Column: 2},
						PieceID:    0,
						IsPromoted: false,
					},
					{
						Source:     &shogi.Point{Row: -1, Column: -1},
						Dest:       &shogi.Point{Row: 7, Column: 3},
						PieceID:    2,
						IsPromoted: false,
					},
					{
						Source:     &shogi.Point{Row: 2, Column: 3},
						Dest:       &shogi.Point{Row: 3, Column: 3},
						PieceID:    0,
						IsPromoted: false,
					},
				},
			},
			0,
			nil,
		},
		{
			"info nodes 120000 nps 116391 hashfull 104",
			&usi.Info{
				Values: map[string]int{
					nodes:    120000,
					nps:      116391,
					hashFull: 104,
				},
				Score: 0,
				Moves: []*shogi.Move{},
			},
			0,
			nil,
		},
		{
			"info score cp 156 multipv 1 pv P*5h 4g5g 5h5g 8b8f",
			&usi.Info{
				Values: make(map[string]int),
				Score:  156,
				Moves: []*shogi.Move{
					{
						Source:     &shogi.Point{Row: -1, Column: -1},
						Dest:       &shogi.Point{Row: 7, Column: 4},
						PieceID:    1,
						IsPromoted: false,
					},
					{
						Source:     &shogi.Point{Row: 6, Column: 3},
						Dest:       &shogi.Point{Row: 6, Column: 4},
						PieceID:    0,
						IsPromoted: false,
					},
					{
						Source:     &shogi.Point{Row: 7, Column: 4},
						Dest:       &shogi.Point{Row: 6, Column: 4},
						PieceID:    0,
						IsPromoted: false,
					},
					{
						Source:     &shogi.Point{Row: 1, Column: 7},
						Dest:       &shogi.Point{Row: 5, Column: 7},
						PieceID:    0,
						IsPromoted: false,
					},
				},
			},
			1,
			nil,
		},
		{
			"info score cp -99 multipv 2 pv 2d4d 3c4e 8h5e N*7f",
			&usi.Info{
				Values: make(map[string]int),
				Score:  -99,
				Moves: []*shogi.Move{
					{
						Source:     &shogi.Point{Row: 3, Column: 1},
						Dest:       &shogi.Point{Row: 3, Column: 3},
						PieceID:    0,
						IsPromoted: false,
					},
					{
						Source:     &shogi.Point{Row: 2, Column: 2},
						Dest:       &shogi.Point{Row: 4, Column: 3},
						PieceID:    0,
						IsPromoted: false,
					},
					{
						Source:     &shogi.Point{Row: 7, Column: 7},
						Dest:       &shogi.Point{Row: 4, Column: 4},
						PieceID:    0,
						IsPromoted: false,
					},
					{
						Source:     &shogi.Point{Row: -1, Column: -1},
						Dest:       &shogi.Point{Row: 5, Column: 6},
						PieceID:    3,
						IsPromoted: false,
					},
				},
			},
			2,
			nil,
		},
		{
			"info score cp -157 multipv 3 pv 5g5f 4g4f 4e3c+ 4c3c",
			&usi.Info{
				Values: make(map[string]int),
				Score:  -157,
				Moves: []*shogi.Move{
					{
						Source:     &shogi.Point{Row: 6, Column: 4},
						Dest:       &shogi.Point{Row: 5, Column: 4},
						PieceID:    0,
						IsPromoted: false,
					},
					{
						Source:     &shogi.Point{Row: 6, Column: 3},
						Dest:       &shogi.Point{Row: 5, Column: 3},
						PieceID:    0,
						IsPromoted: false,
					},
					{
						Source:     &shogi.Point{Row: 4, Column: 3},
						Dest:       &shogi.Point{Row: 2, Column: 2},
						PieceID:    0,
						IsPromoted: true,
					},

					{
						Source:     &shogi.Point{Row: 2, Column: 3},
						Dest:       &shogi.Point{Row: 2, Column: 2},
						PieceID:    0,
						IsPromoted: false,
					},
				},
			},
			3,
			nil,
		},
		{
			"info score cp -157 str multipv 3 lalala... pv 5g5f 4g4f 4e3c+ 4c3c",
			&usi.Info{
				Values: make(map[string]int),
				Score:  -157,
				Moves: []*shogi.Move{
					{
						Source:     &shogi.Point{Row: 6, Column: 4},
						Dest:       &shogi.Point{Row: 5, Column: 4},
						PieceID:    0,
						IsPromoted: false,
					},
					{
						Source:     &shogi.Point{Row: 6, Column: 3},
						Dest:       &shogi.Point{Row: 5, Column: 3},
						PieceID:    0,
						IsPromoted: false,
					},
					{
						Source:     &shogi.Point{Row: 4, Column: 3},
						Dest:       &shogi.Point{Row: 2, Column: 2},
						PieceID:    0,
						IsPromoted: true,
					},

					{
						Source:     &shogi.Point{Row: 2, Column: 3},
						Dest:       &shogi.Point{Row: 2, Column: 2},
						PieceID:    0,
						IsPromoted: false,
					},
				},
			},
			3,
			nil,
		},
		{
			"info score cp -225 multipv 4 pv 5g6h 8b8f P*8g 8f5f",
			&usi.Info{
				Values: make(map[string]int),
				Score:  -225,
				Moves: []*shogi.Move{
					{
						Source:     &shogi.Point{Row: 6, Column: 4},
						Dest:       &shogi.Point{Row: 7, Column: 5},
						PieceID:    0,
						IsPromoted: false,
					},
					{
						Source:     &shogi.Point{Row: 1, Column: 7},
						Dest:       &shogi.Point{Row: 5, Column: 7},
						PieceID:    0,
						IsPromoted: false,
					},
					{
						Source:     &shogi.Point{Row: -1, Column: -1},
						Dest:       &shogi.Point{Row: 6, Column: 7},
						PieceID:    1,
						IsPromoted: false,
					},
					{
						Source:     &shogi.Point{Row: 5, Column: 7},
						Dest:       &shogi.Point{Row: 5, Column: 4},
						PieceID:    0,
						IsPromoted: false,
					},
				},
			},
			4,
			nil,
		},
		{
			"info score cp aaa multipv 4 pv 5g6h 8b8f P*8g 8f5f",
			&usi.Info{},
			0,
			errEmp,
		},
		{
			"info score cp 4 multipv 4 pv 5g6h 8b8f P*8g 8f5z",
			&usi.Info{},
			0,
			errEmp,
		},
		{"info string test", nil, 0, errEmp},
		{"info nodes a nps 116391 hashfull 104", nil, 0, errEmp},
		{"info nodes 12000 nps a hashfull 104", nil, 0, errEmp},
		{"info nodes 12000 nps 116391 hashfull a", nil, 0, errEmp},
		{"info seldepth a", nil, 0, errEmp},
		{"info time a", nil, 0, errEmp},
		{"info depth a", nil, 0, errEmp},
		{"info score cp 4 multipv a", nil, 0, errEmp},
	}

	for i, c := range cases {
		infoHelper(t, i, c.in, c.want, c.mpv, c.err)
	}
}

func infoHelper(
	t *testing.T,
	i int,
	in string,
	want *usi.Info,
	mpv int,
	err error,
) {
	t.Helper()
	msg := ""
	res, mpvKey, e := Info(in)

	if (e == nil) != (err == nil) {
		msg = "The types of two errors does not match"
		infoErrorPrintHelper(t, i, msg, in, err, e)
		return
	}

	if (e != nil) && (err != nil) {
		if !strings.Contains(e.Error(), err.Error()) {
			msg = "Two errors are not nil as expected, but the error message is not correct."
			infoErrorPrintHelper(t, i, msg, in, err, e)
		}
		return
	}

	if (res == nil) != (want == nil) {
		msg = "The types of two *usi.Info does not match"
		infoErrorPrintHelper(t, i, msg, in, err, e)
		return
	}

	if res == nil {
		// won't come here
		t.Errorf("should not come here, because err is nil")
		return
	}

	if mpvKey != mpv {
		msg = "The multipv index value was not as expected."
		infoErrorPrintHelper(t, i, msg, in, mpv, mpvKey)
	}

	valuesEquals := reflect.DeepEqual(want.Values, res.Values)
	scoreEquals := want.Score == res.Score
	movesEquals := func() bool {
		if len(want.Moves) != len(res.Moves) {
			return false
		}
		for j := range want.Moves {
			if !reflect.DeepEqual(want.Moves[j], res.Moves[j]) {
				return false
			}
		}
		return true
	}()

	if !valuesEquals || !scoreEquals || !movesEquals {
		msg = "The value was not as expected."
		infoErrorPrintHelper(t, i, msg, in, want, res)
	}
}

func infoErrorPrintHelper(
	t *testing.T,
	i int,
	msg,
	in string,
	expected,
	actual interface{},
) {
	t.Helper()
	t.Errorf(`[Info] %s
Index:    %d
Input:    %s
Expected: %v
Actual:   %v
`, msg, i, in, expected, actual)
}
