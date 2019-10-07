package sgtime

import (
	"time"
)

func init() {
	globalTimeLocation, _ = time.LoadLocation("PRC")
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
	_, offset := dateTime.Zone()
	timestamp += int64(offset)
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
