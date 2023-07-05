package engine_test

import (
	"testing"
	"time"

	"github.com/uniyuni1101/retroclock/engine"
)

type SpyRender struct {
	callRender int
	timeStack  []time.Time
}

func (r *SpyRender) Render(t time.Time) {
	r.callRender++
}

func TestEngine_RenderingPerTick(t *testing.T) {
	cases := []struct {
		tickTimes []time.Time
		want      int
	}{
		{
			tickTimes: []time.Time{
				time.Now(),
				time.Now(),
				time.Now(),
			},
			want: 3,
		},
		{
			tickTimes: []time.Time{
				time.Now(),
				time.Now(),
				time.Now(),
				time.Now(),
				time.Now(),
			},
			want: 5,
		},
	}

	for _, c := range cases {
		spyRender := &SpyRender{}
		engine := &engine.Engine{
			DelayMS: engine.DefaultDelayMS,
			Render:  spyRender,
		}

		for _, time := range c.tickTimes {
			engine.Tick(time)
		}

		if spyRender.callRender != c.want {
			t.Errorf("call tick want: %d, got: %d", c.want, spyRender.callRender)
		}
	}
}
