package store

import (
	"fmt"
	"sync"

	"github.com/murosan/shogi-board-server/app/domain/entity/engine"
	"github.com/murosan/shogi-board-server/app/domain/infrastructure"
)

// EngineStore is an in-memory store for shogi engines.
type EngineStore interface {
	Insert(*engine.Engine, infrastructure.Connector) error
	Delete(engine.ID) error
	Find(engine.ID) (*engine.Engine, infrastructure.Connector, bool)
	FindAllKeys() []engine.ID
	Exists(engine.ID) bool
}

func NewEngineStore() EngineStore {
	return &engineStore{
		engines:    make(map[engine.ID]*engine.Engine),
		connectors: make(map[engine.ID]infrastructure.Connector),
	}
}

type engineStore struct {
	sync.RWMutex
	engines    map[engine.ID]*engine.Engine
	connectors map[engine.ID]infrastructure.Connector
}

func (s *engineStore) Insert(e *engine.Engine, conn infrastructure.Connector) error {
	id := e.GetID()
	if s.Exists(id) {
		return fmt.Errorf("already exists. id=%s", id)
	}

	s.Lock()
	defer s.Unlock()
	s.engines[id] = e
	s.connectors[id] = conn
	return nil
}

func (s *engineStore) Delete(id engine.ID) error {
	if !s.Exists(id) {
		return fmt.Errorf("no such key. id=%s", id)
	}
	s.Lock()
	defer s.Unlock()
	delete(s.engines, id)
	delete(s.connectors, id)
	return nil
}

func (s *engineStore) Find(id engine.ID) (*engine.Engine, infrastructure.Connector, bool) {
	s.RLock()
	defer s.RUnlock()

	e, ok1 := s.engines[id]
	c, ok2 := s.connectors[id]
	return e, c, ok1 && ok2
}

func (s *engineStore) FindAllKeys() []engine.ID {
	s.RLock()
	defer s.RUnlock()

	v := make([]engine.ID, len(s.engines))

	i := 0
	for id := range s.engines {
		v[i] = id
		i++
	}

	return v
}

func (s *engineStore) Exists(id engine.ID) bool {
	s.RLock()
	defer s.RUnlock()
	_, ok1 := s.engines[id]
	_, ok2 := s.connectors[id]
	return ok1 && ok2
}
