package gctx

import (
	"context"
	"errors"
)

var key *bool

type CtxData struct {
	SampleData string
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
		if incomingData.SampleData != "" {
			data.SampleData = incomingData.SampleData
		}
		return NewCtxWithData(ctx, data)
	}

	return NewCtxWithData(ctx, incomingData)
}

func GetSampleData(ctx context.Context) (string, error) {
	data, ok := GetDataFromCtx(ctx)
	if !ok {
		return "", errors.New("error when get sample data")
	}
	return data.SampleData, nil
}
