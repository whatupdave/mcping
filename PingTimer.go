package mcping

import (
	"time"
)

type PingTimer struct {
	start   uint64 //Start time in ms
	end     uint64 //End time in ms
	latency uint64 //Latency time in ms
}

func getMS() uint64 {
	return uint64(time.Now().UnixNano() / int64(time.Millisecond))
}

func (t *PingTimer) Start() {
	t.start = getMS()
}

func (t *PingTimer) End() (latency uint64) {
	t.end = getMS()
	t.latency = t.end - t.start
	latency = t.latency
	return
}
