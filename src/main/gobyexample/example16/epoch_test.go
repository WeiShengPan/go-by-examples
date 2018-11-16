package main

import (
	"fmt"
	"testing"
	"time"
)

// 时间戳
func TestEpoch(t *testing.T) {

	p := fmt.Println

	now := time.Now()
	p(now)

	secs := now.Unix()
	nanos := now.UnixNano()
	// 没有now.UnixMillis
	millis := nanos / 1000000
	p(secs)
	p(nanos)
	p(millis)

	p(time.Unix(secs, 0))
	p(time.Unix(0, nanos))
}
