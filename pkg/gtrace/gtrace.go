package gtrace

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	"github.com/resyahrial/go-commerce/pkg/inspect"
	log "github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

var ServiceName string = "UndefinedService"
var argsType reflect.Type

func init() {
	var temp map[string]interface{}
	argsType = reflect.TypeOf(temp)
}

// Start creates a span and a context.Context containing the newly-created span.
// optionaly, you can directly set attributes to the span (but don't log password/key)
func Start(ctx context.Context, args ...interface{}) (context.Context, trace.Span) {
	funName, _ := inspect.GetParentFuncProps()

	newCtx, span := otel.Tracer(ServiceName).Start(ctx, funName)

	traceID := span.SpanContext().TraceID().String()
	log.WithFields(log.Fields{
		"trace_id": traceID,
	}).Trace(funName)

	for i, v := range args {
		if reflect.TypeOf(v) == argsType {
			setAttrFromMap(span, v.(map[string]interface{}))
		} else {
			setAttr(span, i, v)
		}
	}

	return newCtx, span
}

func End(span trace.Span, err error) {
	Error(span, err)
	span.End()
}

func Error(span trace.Span, err error) {
	if err == nil {
		return
	}
	log.Error(err.Error())

	if span == nil {
		return
	}
	span.SetAttributes(attribute.String("error.msg", err.Error()))
}

func setAttrFromMap(span trace.Span, m map[string]interface{}) {
	for k, v := range m {
		setAttr(span, k, v)
	}
}

func setAttr(span trace.Span, k, v interface{}) {
	kl := strings.ToLower(fmt.Sprintf("%+v", k))
	if kl == "password" || kl == "key" {
		v = "?"
	}
	key := fmt.Sprintf("args.%s", kl)
	value := fmt.Sprintf("%+v", v)
	log.Trace(key, ": ", value)
	span.SetAttributes(attribute.String(key, value))

}
