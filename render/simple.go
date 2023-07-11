package render

import (
	"fmt"
	"io"
	"time"
)

func NewSimple(w io.Writer) *Simple {
	return &Simple{w}
}

type Simple struct {
	w io.Writer
}

func (s *Simple) Display(t time.Time) {
	fmt.Fprintf(s.w, "\r%d/%d/%d %d-%d-%d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
}
