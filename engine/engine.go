package engine

import "time"

type Renderer interface {
	Render(t time.Time)
}

const (
	DefaultDelayMS = 50 * time.Millisecond
)

type Engine struct {
	DelayMS time.Duration
	Render  Renderer
}

func (e *Engine) Tick(t time.Time) {
	e.Render.Render(t)
}
