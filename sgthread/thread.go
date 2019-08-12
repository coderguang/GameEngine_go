package sgthread

import (
	"time"
)

func SleepBySecond(times int) {
	time.Sleep(time.Duration(times) * time.Second)
}

func SleepByMillSecond(times int) {
	time.Sleep(time.Duration(times) * time.Millisecond)

}
