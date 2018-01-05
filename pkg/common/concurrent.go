package common

import "time"

func DoEvery(d time.Duration, f func()) {

	for {
		f()
		timer := time.NewTimer(d)
		<-timer.C
	}
}
