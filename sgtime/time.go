package sgtime

import (
	"fmt"
	"time"
)

func init() {
	globalTimeLocation = time.FixedZone("CST", 8*3600)
}

var globalTimeLocation *time.Location

const (
	ONE_DAY_SECOND  = 60 * 60 * 24
	ONE_HOUR_SECOND = 60 * 60
)

const (
	FORMAT_TIME_NORMAL          = "2006-01-02 15:04:05"
	FORMAT_TIME_LOG             = "2006/01/02 15:04:05.000"
	FORMAT_TIME_RFC_3339_SIMPLE = "2006-01-02T15:04:05"
	FORMAT_TIME_YEAR            = "2006"
	FORMAT_TIME_MONTH           = "01"
	FORMAT_TIME_DAY             = "02"
	FORMAT_TIME_HOUR            = "15"
	FORMAT_TIME_MINUTE          = "04"
	FORMAT_TIME_SECOND          = "05"
	FORMAT_TIME_YEAR_MONTH      = "200601"
	FORMAT_TIME_YEAR_MONTH_DAY  = "20060102"
)

const (
	SG_WEEKDAY_MONDAY    int = 0
	SG_WEEKDAY_TUESDAY   int = 1
	SG_WEEKDAY_WEDNESDAY int = 2
	SG_WEEKDAY_THURSDAY  int = 3
	SG_WEEKDAY_FRIDAY    int = 4
	SG_WEEKDAY_SATURDAY  int = 5
	SG_WEEKDAY_SUNDAY    int = 6
)

type DateTime = time.Time

type DateTimeDuration = time.Duration

func New() *DateTime {
	now := time.Now().In(globalTimeLocation)
	return &now
}

func TransfromTimeToDateTime(t time.Time) *DateTime {
	return &t
}

func TransformDateTimeToTime(t *DateTime) time.Time {
	return *t
}

func InitTimeLocation(l string) {
	tmp, err := time.LoadLocation(l)
	if err != nil {
		return
	}
	globalTimeLocation = tmp
}

func GetTotalSecond(dateTime *DateTime) int64 {
	timestamp := dateTime.Unix()
	//_, offset := dateTime.Zone()
	//timestamp += int64(offset)
	return timestamp //时区
}

func GetTotalDay(dateTime *DateTime) int64 {
	return GetTotalSecond(dateTime) / ONE_DAY_SECOND
}

func NormalString(dateTime *DateTime) string {
	str_time := dateTime.Format(FORMAT_TIME_NORMAL)
	return str_time
}

func LogString(dateTime *DateTime) string {
	str_time := dateTime.Format(FORMAT_TIME_LOG)
	return str_time
}

func YearString(dateTime *DateTime) string {
	str_time := dateTime.Format(FORMAT_TIME_YEAR)
	return str_time
}

func MonthString(dateTime *DateTime) string {
	str_time := dateTime.Format(FORMAT_TIME_MONTH)
	return str_time
}

func DayString(dateTime *DateTime) string {
	str_time := dateTime.Format(FORMAT_TIME_DAY)
	return str_time
}

func HourString(dateTime *DateTime) string {
	str_time := dateTime.Format(FORMAT_TIME_HOUR)
	return str_time
}

func MinuteString(dateTime *DateTime) string {
	str_time := dateTime.Format(FORMAT_TIME_MINUTE)
	return str_time
}

func SecondString(dateTime *DateTime) string {
	str_time := dateTime.Format(FORMAT_TIME_SECOND)
	return str_time
}

func YMDString(dateTime *DateTime) string {
	return dateTime.Format(FORMAT_TIME_YEAR_MONTH_DAY)
}

func YMString(dateTime *DateTime) string {
	return dateTime.Format(FORMAT_TIME_YEAR_MONTH)
}

func ParseInLocation(format string, timestr string) (DateTime, error) {
	dt, err := time.ParseInLocation(format, timestr, globalTimeLocation)
	return dt, err
}

func GetTimeZone() *time.Location {
	return globalTimeLocation
}

func GetNextDay(now int64, hour int, min int, sec int) int64 {
	if 0 == now {
		nowDt := New()
		now = GetTotalSecond(nowDt)
	}
	curDt := time.Unix(now, 0)
	curZero := time.Date(curDt.Year(), curDt.Month(), curDt.Day(), 0, 0, 0, 0, curDt.Location())
	return GetTotalSecond(TransfromTimeToDateTime(curZero)) + int64(ONE_DAY_SECOND+hour*ONE_HOUR_SECOND+min*60+sec)
}

func GetNextMonth(now int64, day int, hour int, min int, sec int) int64 {
	if 0 == now {
		nowDt := New()
		now = GetTotalSecond(nowDt)
	}
	curDt := time.Unix(now, 0)
	curZero := time.Date(curDt.Year(), curDt.Month(), 1, 0, 0, 0, 0, curDt.Location())
	fmt.Println("curZero", curZero)
	curZero = curZero.AddDate(0, 1, day-1)
	fmt.Println("curZero2", curZero)
	return GetTotalSecond(TransfromTimeToDateTime(curZero)) + int64(hour*ONE_HOUR_SECOND+min*60+sec)
}

func GetNextWeek(now int64, weekday int, hour int, min int, sec int) int64 {
	if 0 == now {
		nowDt := New()
		now = GetTotalSecond(nowDt)
	}
	return 0
}

/*
time.Weekday类型可以做运算，强制转int,会得到偏差数。
默认是 Sunday 开始到 Saturday 算 0,1,2,3,4,5,6

所以只有Monday减去Sunday的时候是正数，特殊处理下就可以了。
*/
func GetWeekDay(now int64) int {
	if 0 == now {
		nowDt := New()
		now = GetTotalSecond(nowDt)
	}
	curDt := time.Unix(now, 0)
	offset := int(curDt.Weekday())
	if offset == 0 {
		offset = SG_WEEKDAY_SUNDAY
	} else {
		offset -= 1
	}
	return offset
}

func GetNextWeekDay(now int64, weekday int, hour int, min int, sec int) int64 {
	if 0 == now {
		nowDt := New()
		now = GetTotalSecond(nowDt)
	}
	curWeekDay := GetWeekDay(now)
	timeInNextWeek := true
	if curWeekDay < weekday {
		//同一周
		timeInNextWeek = false
	} else if curWeekDay == weekday {
		//同一天
		nowTime := time.Unix(now, 0)
		if nowTime.Hour() < hour {
			timeInNextWeek = false
		} else if hour == nowTime.Hour() {
			if nowTime.Minute() < min {
				timeInNextWeek = false
			} else if min == nowTime.Minute() {
				if nowTime.Second() < sec {
					timeInNextWeek = false
				}
			}
		}
	}
	todayZero := GetNextDay(now, 0, 0, 0) - ONE_DAY_SECOND
	mondyZero := todayZero - int64(curWeekDay*ONE_DAY_SECOND)
	if timeInNextWeek {
		return mondyZero + 7*ONE_DAY_SECOND + int64(weekday*ONE_DAY_SECOND) + int64(hour*ONE_HOUR_SECOND+min*60+sec)
	} else {
		return mondyZero + int64(weekday*ONE_DAY_SECOND) + int64(hour*ONE_HOUR_SECOND+min*60+sec)
	}
}
