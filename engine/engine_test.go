package engine_test

import (
	"testing"
	"time"

	"github.com/uniyuni1101/retroclock/config"
	"github.com/uniyuni1101/retroclock/engine"
)

type SpyRender struct {
	callRender int
	timeStack  []time.Time
}

func (r *SpyRender) Render(t time.Time) {
	r.callRender++
}

type StubTicker struct {
	beginTime time.Time
	delta     time.Duration
	tickCount int
}

func (s *StubTicker) Tick() time.Time {
	s.tickCount++
	return s.beginTime.Add(s.delta * time.Duration(s.tickCount))
}

func TestTicker_CollectOfTickIntervals(t *testing.T) {
	cases := map[string]struct {
		tick          engine.TickRate
		count         int
		totalDuration time.Duration
	}{
		"duration of 4 counts at 20 tick per second":  {20, 5, time.Second / 4},
		"duration of 4 counts at 60 tick per second":  {40, 5, time.Second / 8},
		"duration of 4 counts at 120 tick per second": {80, 5, time.Second / 16},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {

			ticker := engine.NewTicker(c.tick)

			t1 := time.Now()
			times := []time.Time{}
			for i := 0; i < c.count; i++ {
				times = append(times, ticker.Tick())
			}
			t2 := time.Now()
			got := t2.Sub(t1)

			assertIntervalWIthinTolerance(t, times, c.tick.Interval(), c.tick.Interval()/20)
			assertDurationWithinTolerance(t, got, c.totalDuration, 5*time.Millisecond)
		})
	}
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
			Config: config.DefaultConfig,
			Render: spyRender,
		}

		for _, time := range c.tickTimes {
			engine.Tick(time)
		}

		if spyRender.callRender != c.want {
			t.Errorf("call tick want: %d, got: %d", c.want, spyRender.callRender)
		}
	}
}

func assertDurationWithinTolerance(t *testing.T, got, want, tolerance time.Duration) {
	t.Helper()

	if got < want-tolerance || want+tolerance < got {
		t.Errorf("got %s, want acceptable duration is %s±%s", got, want, tolerance)
	}
}

func assertIntervalWIthinTolerance(t *testing.T, got []time.Time, delta, tolerance time.Duration) {
	t.Helper()

	for i := range got[:len(got)-1] {
		gotDelta := got[i+1].Sub(got[i]).Abs()

		if gotDelta < delta-tolerance || delta+tolerance < gotDelta {
			t.Errorf("got deltas %v, acceptable delta is %s±%s", got, delta, tolerance)
		}
	}
}
