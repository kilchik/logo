package logo

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

var (
	d *log.Logger
	i *log.Logger
	e *log.Logger
)

const TraceIdKey = "trace_id"

func Init(debugOn bool)  {
	var debugWriter io.Writer
	if debugOn {
		debugWriter = os.Stderr
	} else {
		debugWriter = ioutil.Discard
	}
	d = log.New(debugWriter, "[D] ", 0)
	i = log.New(os.Stdout, "[I] ", log.Ldate | log.Ltime)
	e = log.New(os.Stderr, "[E] ", log.Ldate | log.Ltime)
}

func Info(ctx context.Context, format string, v ...interface{}) {
	i.Printf(getTraceIdPrefix(ctx) + format, v...)
}

func Error(ctx context.Context, format string, v ...interface{}) {
	e.Printf(getTraceIdPrefix(ctx) + format, v...)
}

func Fatal(ctx context.Context, format string, v ...interface{}) {
	e.Fatalf(getTraceIdPrefix(ctx) + format, v...)
}

func Debug(ctx context.Context, format string, v ...interface{}) {
	d.Printf(getTraceIdPrefix(ctx) + format, v...)
}

func getTraceIdPrefix(ctx context.Context) string {
	traceId, ok := ctx.Value(TraceIdKey).(string)
	if ok && traceId != "" {
		return fmt.Sprintf("[trace: %s] ", traceId)
	}
	return ""
}
