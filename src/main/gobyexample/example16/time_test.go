package main

import (
	"fmt"
	"time"
)

func Test(t *testing.T) {

	p := fmt.Println

	now := time.Now()
	p(now)

	then := time.Date(2018, 10, 1, 23, 59, 59, 651387237, time.UTC)
	p(then)
	p(then.Year())
	p(then.Month())
	p(then.Day())
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Nanosecond())
	p(then.Location())
	p(then.Weekday())
	p(then.Before(now))
	p(then.After(now))
	p(then.Equal(now))

	// Sub 返回一个Duration表示两个时间点的间隔时间
	diff := now.Sub(then)
	p(diff)

	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Nanoseconds())

	p(then.Add(diff))
	p(then.Add(-diff))

}
