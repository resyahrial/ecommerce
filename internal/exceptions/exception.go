package exceptions

import (
	"reflect"

	"github.com/resyahrial/go-commerce/pkg/gexception"
)

type E struct{}

func init() {
	AddAllErrors(E{})
}

func Handle(e interface{}) error {
	if e == nil {
		return nil
	}

	if reflect.TypeOf(e) == gexception.ExceptionType {
		return e.(error)
	}

	return gexception.BaseInternalServerError
}

func AddAllErrors(e E) {
	g := e
	methodFinder := reflect.TypeOf(&g)
	for i := 0; i < methodFinder.NumMethod(); i++ {
		method := methodFinder.Method(i)
		method.Func.Call([]reflect.Value{reflect.ValueOf(&g)})
	}
}
