package log

import (
	"context"
	"gorm.io/gorm/logger"
	"time"
)

type traceRecorder struct {
	logger.Interface
	BeginAt      time.Time
	SQL          string
	RowsAffected int64
	Err          error
}

func (t traceRecorder) New() *traceRecorder {
	return &traceRecorder{Interface: t.Interface, BeginAt: time.Now()}
}

func (t *traceRecorder) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	t.BeginAt = begin
	t.SQL, t.RowsAffected = fc()
	t.Err = err
}
