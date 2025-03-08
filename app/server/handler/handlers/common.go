package handlers

import (
	"github.com/murosan/shogi-board-server/app/domain/entity/engine"
	"github.com/murosan/shogi-board-server/app/domain/framework"
	"github.com/murosan/shogi-board-server/app/server/handler"
)

// QueryKeys is a set of uri query keys.
var QueryKeys = struct {
	EngineID,
	EngineIDAlias string
}{
	EngineID:      "engine",
	EngineIDAlias: "key", // for backward compatibility
}

// WithEngineID executes block with an engine.ID if the engine name is specified,
// otherwise returns BAD_REQUEST error.
func WithEngineID(ctx *handler.Context, block func(engine.ID) error) error {
	id, err := GetEngineID(ctx)
	if err != nil {
		return err
	}
	return block(id)
}

// GetEngineID gets the engine.ID from the uri query and returns it,
// otherwise returns BAD_REQUEST error.
func GetEngineID(ctx *handler.Context) (engine.ID, error) {
	if v := ctx.GetQuery(QueryKeys.EngineID); v != "" {
		return engine.ID(v), nil
	}

	if v := ctx.GetQuery(QueryKeys.EngineIDAlias); v != "" {
		return engine.ID(v), nil
	}

	errMsg := "please specify engine id in query parameter"
	return "", framework.ErrBadRequest.With(errMsg)
}
