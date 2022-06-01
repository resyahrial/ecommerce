package rest

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/resyahrial/go-commerce/internal/exceptions"
	"github.com/resyahrial/go-commerce/pkg/gctx"
	"github.com/resyahrial/go-commerce/pkg/gtrace"
)

type RoleMiddleware struct{}

func NewRoleMiddleware() *RoleMiddleware {
	return &RoleMiddleware{}
}

func (m *RoleMiddleware) Wrap(nextHandler httprouter.Handle, role string) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		var err error
		var actor gctx.Actor

		newCtx, span := gtrace.Start(r.Context())
		defer gtrace.End(span, err)

		if actor, err = gctx.GetActor(newCtx); err != nil {
			panic(err)
		}

		if role != "" && !actor.Is(role) {
			panic(exceptions.AuthForbiddenRole)
		}

		nextHandler(w, r.WithContext(newCtx), params)
	}
}
