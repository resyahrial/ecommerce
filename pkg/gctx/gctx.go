package gctx

import (
	"context"
	"errors"

	"github.com/segmentio/ksuid"
)

var key *bool

type CtxData struct {
	Actor
}

type Actor struct {
	ID   ksuid.KSUID `json:"id"`
	Role string      `json:"role"`
}

func (a Actor) Is(role string) bool {
	return a.Role == role
}

// NewCtxWithData returns a new Context that carries personalized data.
func NewCtxWithData(ctx context.Context, data CtxData) context.Context {
	return context.WithValue(ctx, key, data)
}

// GetDataFromCtx returns the data value stored in ctx, if any.
func GetDataFromCtx(ctx context.Context) (CtxData, bool) {
	u, ok := ctx.Value(key).(CtxData)
	return u, ok
}

// SetDataAndGetNewCtx
// - Update the Span if the function that is being passed the new Ctx
// is the child of the span's function.
func SetDataAndGetNewCtx(ctx context.Context, incomingData CtxData) context.Context {
	data, ok := GetDataFromCtx(ctx)
	if ok {
		if !incomingData.Actor.ID.IsNil() {
			data.Actor = incomingData.Actor
		}
		return NewCtxWithData(ctx, data)
	}

	return NewCtxWithData(ctx, incomingData)
}

func GetActor(ctx context.Context) (Actor, error) {
	data, ok := GetDataFromCtx(ctx)
	if !ok {
		return Actor{}, errors.New("error when get actor")
	}
	return data.Actor, nil
}
