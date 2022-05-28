package gtrace

import (
	"context"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	gorm_logger "gorm.io/gorm/logger"
)

// ktrace.LogAndTracer is package to help set the log
type LogAndTracer struct {
	LogLevel gorm_logger.LogLevel

	Title      string
	StringName string
	CountName  string

	formattedStringName  string
	formattedCountName   string
	formattedElapsedName string
}

func NewLogAndTracer(l LogAndTracer) LogAndTracer {
	l.formattedStringName = fmt.Sprintf("%s.%s", l.Title, l.StringName)
	l.formattedCountName = fmt.Sprintf("%s.%s", l.Title, l.CountName)
	l.formattedElapsedName = fmt.Sprintf("%s.elapsed", l.Title)
	return l
}

func (l *LogAndTracer) LogMode(level gorm_logger.LogLevel) gorm_logger.Interface {
	newlogger := *l
	newlogger.LogLevel = level
	return &newlogger
}

func (l LogAndTracer) Info(ctx context.Context, msg string, data ...interface{}) {
	if isCanPrint(log.GetLevel(), gorm_logger.Info) {
		print(ctx, msg, data)
	}
}

func (l LogAndTracer) Warn(ctx context.Context, msg string, data ...interface{}) {
	if isCanPrint(log.GetLevel(), gorm_logger.Warn) {
		print(ctx, msg, data)
	}

	span := trace.SpanFromContext(ctx)

	span.SetAttributes(attribute.String(fmt.Sprintf("%s.warn", l.Title), msg))
}

func (l LogAndTracer) Error(ctx context.Context, msg string, data ...interface{}) {
	if isCanPrint(log.GetLevel(), gorm_logger.Error) {
		print(ctx, msg, data)
	}

	span := trace.SpanFromContext(ctx)
	span.SetAttributes(attribute.String(fmt.Sprintf("%s.error", l.Title), msg))
}

func (l LogAndTracer) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	elapsed := time.Since(begin)

	span := trace.SpanFromContext(ctx)
	traceID := span.SpanContext().TraceID().String()

	if err != nil {
		log.WithFields(log.Fields{"trace_id": traceID}).Error(err)
	}

	str, count := fc()
	span.SetAttributes(
		attribute.String(l.formattedStringName, str),
		attribute.Int64(l.formattedCountName, count),
		attribute.String(l.formattedElapsedName, elapsed.String()),
	)

	if err != nil {
		span.SetAttributes(attribute.String(fmt.Sprintf("%s.error", l.Title), err.Error()))
	}

	log.WithFields(log.Fields{
		"trace_id": traceID,
	}).Debugf("%dms | %d %s(s) | %s\n", elapsed.Milliseconds(), count, l.CountName, str)
}

func print(ctx context.Context, msg string, data ...interface{}) {
	span := trace.SpanFromContext(ctx)
	traceID := span.SpanContext().TraceID().String()

	fmt.Println(msg, traceID)
	for _, v := range data {
		fmt.Println(v)
	}
}

func isCanPrint(one log.Level, two gorm_logger.LogLevel) bool {
	return int(one) >= int(two)+2
}
