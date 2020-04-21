package sgtime

import (
	"fmt"
	"testing"
	"time"
)

func Test_NextUpdate(t *testing.T) {
	now := New()
	nowsec := GetTotalSecond(now)
	nextSec := GetNextDay(nowsec, 0, 0, 0)
	fmt.Println("1,now is:", nowsec, ",format:", time.Unix(nowsec, 0))
	fmt.Println("1next day is :", nextSec, ",format:", time.Unix(nextSec, 0))

	nowsec = 1587571299
	nextSec = GetNextDay(nowsec, 0, 0, 0)
	fmt.Println("2,now is:", nowsec, ",format:", time.Unix(nowsec, 0))
	fmt.Println("2next day is :", nextSec, ",format:", time.Unix(nextSec, 0))

	nowsec = GetTotalSecond(now)
	nextSec = GetNextDay(nowsec, 2, 0, 0)
	fmt.Println("3,now is:", nowsec, ",format:", time.Unix(nowsec, 0))
	fmt.Println("3next day is :", nextSec, ",format:", time.Unix(nextSec, 0))

	nextSec = GetNextDay(nowsec, 1, 30, 59)
	fmt.Println("4,now is:", nowsec, ",format:", time.Unix(nowsec, 0))
	fmt.Println("4next day is :", nextSec, ",format:", time.Unix(nextSec, 0))

	nextSec = GetNextDay(nowsec, 23, 59, 59)
	fmt.Println("5,now is:", nowsec, ",format:", time.Unix(nowsec, 0))
	fmt.Println("5next day is :", nextSec, ",format:", time.Unix(nextSec, 0))

	t.Log("test time ok")
}

func Test_NextMonthUpdate(t *testing.T) {
	now := New()
	nowsec := GetTotalSecond(now)
	nextSec := GetNextMonth(nowsec, 1, 0, 0, 0)
	fmt.Println("1,now is:", nowsec, ",format:", time.Unix(nowsec, 0))
	fmt.Println("1next month is :", nextSec, ",format:", time.Unix(nextSec, 0))
	fmt.Println("=======end===\n ")
	nowsec = 1577881500
	nextSec = GetNextMonth(nowsec, 1, 0, 0, 0)
	fmt.Println("2,now is:", nowsec, ",format:", time.Unix(nowsec, 0))
	fmt.Println("2next month is :", nextSec, ",format:", time.Unix(nextSec, 0))
	fmt.Println("=======end===\n ")

	nowsec = 1577805900
	nextSec = GetNextMonth(nowsec, 1, 0, 0, 0)
	fmt.Println("3,now is:", nowsec, ",format:", time.Unix(nowsec, 0))
	fmt.Println("3next month is :", nextSec, ",format:", time.Unix(nextSec, 0))
	fmt.Println("=======end===\n ")

	nowsec = 1609428300
	nextSec = GetNextMonth(nowsec, 1, 0, 0, 0)
	fmt.Println("4,now is:", nowsec, ",format:", time.Unix(nowsec, 0))
	fmt.Println("4next month is :", nextSec, ",format:", time.Unix(nextSec, 0))
	fmt.Println("=======end===\n ")

	GetTotalSecond(now)
	nextSec = GetNextMonth(nowsec, 15, 23, 59, 59)
	fmt.Println("5,now is:", nowsec, ",format:", time.Unix(nowsec, 0))
	fmt.Println("5next month is :", nextSec, ",format:", time.Unix(nextSec, 0))
	fmt.Println("=======end===\n ")

	t.Log("test time ok")
}

func Test_NextWeek(t *testing.T) {
	now := New()
	nowsec := GetTotalSecond(now)
	nowsec = 1587905820
	nextSec := GetNextWeekDay(nowsec, SG_WEEKDAY_MONDAY, 0, 0, 0)
	fmt.Println("1,now is:", nowsec, ",format:", time.Unix(nowsec, 0))
	fmt.Println("1next week is :", nextSec, ",format:", time.Unix(nextSec, 0))
	fmt.Println("=======end===\n ")

	nextSec = GetNextWeekDay(nowsec, SG_WEEKDAY_TUESDAY, 1, 0, 0)
	fmt.Println("2,now is:", nowsec, ",format:", time.Unix(nowsec, 0))
	fmt.Println("2next week is :", nextSec, ",format:", time.Unix(nextSec, 0))
	fmt.Println("=======end===\n ")

	nextSec = GetNextWeekDay(nowsec, SG_WEEKDAY_WEDNESDAY, 0, 0, 2)
	fmt.Println("3,now is:", nowsec, ",format:", time.Unix(nowsec, 0))
	fmt.Println("3next week is :", nextSec, ",format:", time.Unix(nextSec, 0))
	fmt.Println("=======end===\n ")

	nextSec = GetNextWeekDay(nowsec, SG_WEEKDAY_THURSDAY, 0, 3, 0)
	fmt.Println("4,now is:", nowsec, ",format:", time.Unix(nowsec, 0))
	fmt.Println("4next week is :", nextSec, ",format:", time.Unix(nextSec, 0))
	fmt.Println("=======end===\n ")

	nextSec = GetNextWeekDay(nowsec, SG_WEEKDAY_FRIDAY, 4, 0, 0)
	fmt.Println("5,now is:", nowsec, ",format:", time.Unix(nowsec, 0))
	fmt.Println("5next week is :", nextSec, ",format:", time.Unix(nextSec, 0))
	fmt.Println("=======end===\n ")

	nextSec = GetNextWeekDay(nowsec, SG_WEEKDAY_SATURDAY, 23, 0, 0)
	fmt.Println("6,now is:", nowsec, ",format:", time.Unix(nowsec, 0))
	fmt.Println("6next week is :", nextSec, ",format:", time.Unix(nextSec, 0))
	fmt.Println("=======end===\n ")

	nextSec = GetNextWeekDay(nowsec, SG_WEEKDAY_SUNDAY, 20, 58, 30)
	fmt.Println("7,now is:", nowsec, ",format:", time.Unix(nowsec, 0))
	fmt.Println("7next week is :", nextSec, ",format:", time.Unix(nextSec, 0))
	fmt.Println("=======end===\n ")

	t.Log("test time ok")
}
