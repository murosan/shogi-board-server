package usi

import (
	"github.com/pkg/errors"
	"strconv"
	"strings"

	"github.com/murosan/shogi-board-server/app/domain/model/shogi"
	"github.com/murosan/shogi-board-server/app/domain/model/usi"
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

// ParseInfo generates engine.Info parsing from given string, and returns it
func ParseInfo(s string) (*usi.Info, int, error) {
	// should not pass 'info string'
	if strings.HasPrefix(s, "info string") {
		return nil, 0, errors.New("'info string' was given")
	}

	a := strings.Split(s, " ")
	r := &usi.Info{Values: make(map[string]int)}
	mpv := 0

	nan := "given value was not a number. value = "

	i := 0
	for i < len(a) {
		switch strings.TrimSpace(a[i]) {
		case depth:
			i++
			n, err := strconv.Atoi(a[i])
			if err != nil {
				return nil, 0, errors.Wrap(err, nan+a[i])
			}
			r.Values[depth] = n

		case selDepth:
			i++
			n, err := strconv.Atoi(a[i])
			if err != nil {
				return nil, 0, errors.Wrap(err, nan+a[i])
			}
			r.Values[selDepth] = n

		case time:
			i++
			n, err := strconv.Atoi(a[i])
			if err != nil {
				return nil, 0, errors.Wrap(err, nan+a[i])
			}
			r.Values[time] = n

		case nodes:
			i++
			n, err := strconv.Atoi(a[i])
			if err != nil {
				return nil, 0, errors.Wrap(err, nan+a[i])
			}
			r.Values[nodes] = n

		case hashFull:
			i++
			n, err := strconv.Atoi(a[i])
			if err != nil {
				return nil, 0, errors.Wrap(err, nan+a[i])
			}
			r.Values[hashFull] = n

		case nps:
			i++
			n, err := strconv.Atoi(a[i])
			if err != nil {
				return nil, 0, errors.Wrap(err, nan+a[i])
			}
			r.Values[nps] = n

		case score:
			if a[i+1] == "cp" || a[i+1] == "mate" {
				n, err := strconv.Atoi(a[i+2])
				if err != nil {
					return nil, 0, errors.Wrap(err, nan+a[i+2])
				}
				r.Score = n
			}
			i += 2

		case multiPv:
			i++
			n, err := strconv.Atoi(a[i])
			if err != nil {
				return nil, 0, errors.Wrap(err, nan+a[i])
			}
			mpv = n

		case pv:
			m := make([]*shogi.Move, len(a[i+1:]))
			for j, v := range a[i+1:] {
				mv, err := ParseMove(v)
				if err != nil {
					return nil, 0, errors.Wrap(err, "error at ParseMove")
				}
				m[j] = mv
			}
			r.Moves = m
			i += len(a) // force to end this loop, because pv must be in ending of info.
		}
		i++
	}

	return r, mpv, nil
}
