package logo

import (
	"bytes"
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDebugNoTraceId(t *testing.T) {
	Init(true)

	buf := &bytes.Buffer{}
	d.SetOutput(buf)

	Debug(context.Background(), "hello")
	assert.Equal(t, "[D] hello\n", buf.String())
}

func TestDebugWithTraceId(t *testing.T) {
	Init(true)

	buf := &bytes.Buffer{}
	d.SetOutput(buf)
	ctx := context.WithValue(context.Background(), TraceIdKey, "x123")

	Debug(ctx, "hello")
	assert.Equal(t, "[D] [trace: x123] hello\n", buf.String())
}
