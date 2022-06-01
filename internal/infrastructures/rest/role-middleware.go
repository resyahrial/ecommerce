package rest

import (
	"net/http"

	"github.com/resyahrial/go-commerce/internal/exceptions"
	"github.com/resyahrial/go-commerce/pkg/gctx"
	"github.com/resyahrial/go-commerce/pkg/gtrace"
)

type RoleMiddleware struct {
	nextHandler http.Handler
	role        string
}

func NewRoleMiddleware(nextHandler http.Handler, role string) *RoleMiddleware {
	return &RoleMiddleware{nextHandler: nextHandler, role: role}
}

func (m *RoleMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var err error
	var actor gctx.Actor

	newCtx, span := gtrace.Start(r.Context())
	defer gtrace.End(span, err)

	if actor, err = gctx.GetActor(newCtx); err != nil {
		panic(err)
	}

	if !actor.Is(m.role) {
		panic(exceptions.AuthForbiddenRole)
	}

	m.nextHandler.ServeHTTP(w, r.WithContext(newCtx))
}
