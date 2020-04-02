package store

import (
	"sync"

	"github.com/murosan/shogi-board-server/app/domain/entity/engine"
	"github.com/murosan/shogi-board-server/app/domain/entity/usi"
)

// EngineInfoStore is a in memory store
// for information given from shogi engines.
type EngineInfoStore interface {
	FindAll(engine.ID) usi.Result
	Upsert(engine.ID, int, *usi.Info)
	DeleteAll(engine.ID)
}

func NewEngineInfoStore() EngineInfoStore {
	return &engineInfoStore{
		m: make(map[engine.ID]usi.Result),
	}
}

type engineInfoStore struct {
	sync.RWMutex
	m map[engine.ID]usi.Result
}

func (repo *engineInfoStore) FindAll(id engine.ID) usi.Result {
	repo.RLock()
	a := repo.m[id]
	repo.RUnlock()

	m := make(usi.Result)
	for i, info := range a {
		m[i] = info
	}

	return m
}

func (repo *engineInfoStore) Upsert(id engine.ID, index int, info *usi.Info) {
	repo.Lock()
	defer repo.Unlock()

	m, ok := repo.m[id]
	if !ok {
		m = make(usi.Result)
		repo.m[id] = m
	}
	m[index] = info
}

func (repo *engineInfoStore) DeleteAll(id engine.ID) {
	repo.Lock()
	delete(repo.m, id)
	repo.Unlock()
}
