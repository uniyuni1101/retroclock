package engine

import (
	"time"

	"github.com/uniyuni1101/retroclock/config"
)

type Renderer interface {
	Display(t time.Time)
}

type Ticker interface {
	Tick() time.Time
}

type TickRate int

func (t TickRate) Interval() time.Duration {
	return time.Second / time.Duration(t)
}

type TickController struct {
	c        chan time.Time
	TickRate TickRate
	tickTime time.Time
	done     chan struct{}
}

func NewTickController(tick TickRate) *TickController {
	ticker := &TickController{
		c:        make(chan time.Time),
		TickRate: tick,
		tickTime: time.Now(),
	}
	ticker.start()
	return ticker
}

func (t *TickController) start() {
	go func() {
	Loop:
		for {
			select {
			case <-t.done:
				break Loop
			default:
				time.Sleep(t.when())
				t.tickTime = t.tickTime.Add(t.TickRate.Interval())
				t.c <- t.tickTime
			}
		}
	}()
}

func (t *TickController) when() time.Duration {
	target := t.tickTime.Add(t.TickRate.Interval())
	return time.Until(target)
}

func (t *TickController) Tick() time.Time {
	tick := <-t.c
	return tick
}

type Engine struct {
	Config config.Config
	Render Renderer
	Ticker Ticker
}

func (e *Engine) Tick(t time.Time) {
	e.Render.Display(t)
}

func (e *Engine) Run() {
	for {
		t := e.Ticker.Tick()
		e.Tick(t)
	}
}
