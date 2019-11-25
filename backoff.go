package notif

import (
	"time"
)

type Backoff interface {
	next(retry int) time.Duration
}

type ConstantBackoff struct {
	BackoffInterval int64
}

func (cb ConstantBackoff) next(retry int) time.Duration {
	return time.Duration(cb.BackoffInterval) * time.Millisecond
}

type ExponentBackoff struct {
	InitialInterval  int64
	ExponentInterval int64
}

func (eb ExponentBackoff) next(retry int) time.Duration {
	return (time.Duration(eb.InitialInterval) * time.Millisecond) + (time.Duration(eb.ExponentInterval*int64(retry)) * time.Millisecond)
}
