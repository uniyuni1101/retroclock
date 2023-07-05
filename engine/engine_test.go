package engine_test

import (
	"testing"
	"time"

	"github.com/uniyuni1101/retroclock/engine"
)

type SpyRender struct {
	callRender int
}

func (r *SpyRender) Render(t time.Time) {
	r.callRender++
}

func TestEngine(t *testing.T) {
	spyRender := &SpyRender{}
	engine := &engine.Engine{
		DelayMS:  engine.DefaultDelayMS,
		Renderer: spyRender,
	}

	want := 1
	engine.Tick(time.Date(2023, 1, 1, 1, 23, 45, 0, time.UTC))

	if spyRender.callRender != want {
		t.Errorf("call tick want: %d, got: %d", want, spyRender.callRender)
	}
}
