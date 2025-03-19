package parse

import (
	"errors"
	"strconv"
	"strings"

	"github.com/murosan/shogi-board-server/app/domain/entity/shogi"
	"github.com/murosan/shogi-board-server/app/domain/entity/usi"
	"golang.org/x/xerrors"
)

const (
	depth    = "depth"
	selDepth = "seldepth"
	time     = "time"
	nodes    = "nodes"
	hashFull = "hashfull"
	nps      = "nps"
	score    = "score"
	pv       = "pv"
	multiPv  = "multipv"
)

// Info generates engine.Info parsing from given string, and returns it
func Info(s string) (*usi.Info, int, error) {
	// should not pass 'info string'
	if strings.HasPrefix(s, "info string") {
		return nil, 0, xerrors.New("'info string' was given")
	}

	a := strings.Split(s, " ")
	r := &usi.Info{Values: make(map[string]int)}
	mpv := 0

	for len(a) > 0 {
		tpe := strings.TrimSpace(a[0])

		a = a[1:]
		if len(a) == 0 {
			return nil, 0, xerrors.New("wrong info format")
		}

		switch tpe {
		case depth:
			n, err := strconv.Atoi(a[0])
			if err != nil {
				return nil, 0, xerrors.Errorf(": %w", err)
			}
			a = a[1:]
			r.Values[depth] = n

		case selDepth:
			n, err := strconv.Atoi(a[0])
			if err != nil {
				return nil, 0, xerrors.Errorf(": %w", err)
			}
			a = a[1:]
			r.Values[selDepth] = n

		case time:
			n, err := strconv.Atoi(a[0])
			if err != nil {
				return nil, 0, xerrors.Errorf(": %w", err)
			}
			a = a[1:]
			r.Values[time] = n

		case nodes:
			n, err := strconv.Atoi(a[0])
			if err != nil {
				return nil, 0, xerrors.Errorf(": %w", err)
			}
			a = a[1:]
			r.Values[nodes] = n

		case hashFull:
			n, err := strconv.Atoi(a[0])
			if err != nil {
				return nil, 0, xerrors.Errorf(": %w", err)
			}
			a = a[1:]
			r.Values[hashFull] = n

		case nps:
			n, err := strconv.Atoi(a[0])
			if err != nil {
				return nil, 0, xerrors.Errorf(": %w", err)
			}
			a = a[1:]
			r.Values[nps] = n

		case score:
			if len(a) < 2 {
				return nil, 0, xerrors.New("wrong score format")
			}
			if a[0] == "cp" || a[0] == "mate" {
				n, err := strconv.Atoi(a[1])
				if err != nil {
					return nil, 0, xerrors.Errorf(": %w", err)
				}
				r.Score = n
			}
			a = a[2:]

		case multiPv:
			n, err := strconv.Atoi(a[0])
			if err != nil {
				return nil, 0, xerrors.Errorf(": %w", err)
			}
			a = a[1:]
			mpv = n

		case pv:
			m := make([]*shogi.Move, 0, len(a))
			for j, v := range a {
				mv, err := Move(v)
				if err != nil {
					if errors.Is(err, ErrCustomUSIFormat) {
						continue
					}
					return nil, 0, xerrors.Errorf("pv(j=%v,v=%s): %w", j, v, err)
				}
				m = append(m, mv)
			}
			r.Moves = m

			// exit loop, because pv must be in last part of info.
			a = nil
		}
	}

	return r, mpv, nil
}
