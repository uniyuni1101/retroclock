package engine

import (
	"time"

	"github.com/uniyuni1101/retroclock/config"
)

type Renderer interface {
	Render(t time.Time)
}

type TickPerSecond int

func (t TickPerSecond) Interval() time.Duration {
	return time.Second / time.Duration(t)
}

type Ticker struct {
	c             chan time.Time
	tickPerSecond TickPerSecond
	tickTime      time.Time
	done          chan struct{}
}

func NewTicker(tick TickPerSecond) *Ticker {
	ticker := &Ticker{
		c:             make(chan time.Time),
		tickPerSecond: tick,
		tickTime:      time.Now(),
	}
	ticker.start()
	return ticker
}

func (t *Ticker) start() {
	go func() {
	Loop:
		for {
			select {
			case <-t.done:
				break Loop
			default:
				time.Sleep(t.when())
				t.tickTime = t.tickTime.Add(t.tickPerSecond.Interval())
				t.c <- t.tickTime
			}
		}
	}()
}

func (t *Ticker) when() time.Duration {
	target := t.tickTime.Add(t.tickPerSecond.Interval())
	return time.Until(target)
}

func (t *Ticker) Tick() time.Time {
	tick := <-t.c
	return tick
}

type Engine struct {
	Config config.Config
	Render Renderer
}

func (e *Engine) Tick(t time.Time) {
	e.Render.Render(t)
}
