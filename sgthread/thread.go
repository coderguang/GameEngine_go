package sgthread

import (
	"os"
	"time"
)

func SleepBySecond(times int) {
	time.Sleep(time.Duration(times) * time.Second)
}

func SleepByMillSecond(times int) {
	time.Sleep(time.Duration(times) * time.Millisecond)

}

func DelayExit(delaytime int) {
	SleepBySecond(delaytime)
	os.Exit(1)
}
