package gexception

import (
	"fmt"
	"reflect"

	log "github.com/sirupsen/logrus"
)

type E struct{}
type ExceptionMap map[string]Exception

var Expections ExceptionMap
var ExceptionType reflect.Type

func init() {
	Expections = make(ExceptionMap)
	ExceptionType = reflect.TypeOf(&Exception{})
	AddBaseExceptions(E{})
}

type Exception struct {
	// Description to help other engineers translates code to
	// more understandable term for end-users
	Description string `json:"desc"`
	Module      string `json:"module"`
	Code        string `json:"message"`
	HttpStatus  int    `json:"status"`
}

func (e *Exception) Error() string {
	return fmt.Sprintf("%s-%s", e.Module, e.Description)
}

func (e *Exception) New(errDesc string) *Exception {
	e.Description = errDesc
	return e
}

func AddBaseExceptions(e interface{}) {
	g := e
	methodFinder := reflect.TypeOf(&g)
	for i := 0; i < methodFinder.NumMethod(); i++ {
		method := methodFinder.Method(i)
		method.Func.Call([]reflect.Value{reflect.ValueOf(&g)})
	}
}

// RegisterException is shortcut so it's easier to add new err
func RegisterException(moduleName string, es ...*Exception) {
	for _, e := range es {
		e.Module = moduleName
		key := e.Error()
		if _, result := Expections[key]; result {
			log.Error(e)
			log.Error(Expections[key])
			panic("duplicate error definition: " + key)
		}
		Expections[key] = *e
	}
}
