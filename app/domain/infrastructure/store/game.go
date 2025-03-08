package store

import (
	"sync"

	"github.com/murosan/shogi-board-server/app/domain/entity/engine"
	"github.com/murosan/shogi-board-server/app/domain/entity/shogi"
)

// GameStore is an in-memory store that holds game state.
type GameStore interface {
	FindPosition(engine.ID) (*shogi.Position, bool)
	UpsertPosition(engine.ID, *shogi.Position)
	DeletePosition(engine.ID)
}

func NewGameStore() GameStore {
	return &gameStore{
		pos: make(map[engine.ID]*shogi.Position),
	}
}

type gameStore struct {
	sync.RWMutex
	pos map[engine.ID]*shogi.Position
}

func (s *gameStore) FindPosition(id engine.ID) (*shogi.Position, bool) {
	s.RLock()
	pos, ok := s.pos[id]
	s.RUnlock()
	return pos, ok
}

func (s *gameStore) UpsertPosition(id engine.ID, pos *shogi.Position) {
	s.Lock()
	s.pos[id] = pos
	s.Unlock()
}

func (s *gameStore) DeletePosition(id engine.ID) {
	s.Lock()
	delete(s.pos, id)
	s.Unlock()
}
