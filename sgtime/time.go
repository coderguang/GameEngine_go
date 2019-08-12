package sgtime

import "time"

const (
	ONE_DAY_SECOND  = 60 * 60 * 24
	ONE_HOUR_SECOND = 60 * 60
)

const (
	FORMAT_TIME_NORMAL = "2006-01-02 15:04:05"
	FORMAT_TIME_LOG    = "2006/01/02 15:04:05.000"
	FORMAT_TIME_YEAR   = "2006"
	FORMAT_TIME_MONTH  = "01"
	FORMAT_TIME_DAY    = "02"
	FORMAT_TIME_HOUR   = "15"
	FORMAT_TIME_MINUTE = "04"
	FORMAT_TIME_SECOND = "05"
)

type DateTime struct {
	dt time.Time
}

type DateTimeDuration struct {
	duration time.Duration
}

func New() *DateTime {
	dateTime := new(DateTime)
	dateTime.dt = time.Now()
	return dateTime
}

func (datetime *DateTime) Location() *time.Location {
	return datetime.dt.Location()
}

func Date(year int, month time.Month, day, hour, min, sec, nsec int, loc *time.Location) *DateTime {
	dateTime := new(DateTime)
	dateTime.dt = time.Date(year, month, day, hour, min, sec, nsec, loc)
	return dateTime
}

func (dateTime *DateTime) Year() int {
	return dateTime.dt.Year()
}

func (dateTime *DateTime) Month() time.Month {
	return dateTime.dt.Month()
}

func (dateTime *DateTime) Day() int {
	return dateTime.dt.Day()
}

func (dateTime *DateTime) Hour() int {
	return dateTime.dt.Hour()
}

func (dateTime *DateTime) Minute() int {
	return dateTime.dt.Minute()
}

func (dateTime *DateTime) Second() int {
	return dateTime.dt.Second()
}

func (dateTime *DateTime) GetTotalSecond() int64 {
	return dateTime.dt.Unix() + 8*ONE_HOUR_SECOND*1000 //时区
}

func (dateTime *DateTime) GetTotalDay() int64 {
	return dateTime.GetTotalSecond() / ONE_DAY_SECOND
}

func (dateTime *DateTime) NormalString() string {
	str_time := dateTime.dt.Format(FORMAT_TIME_NORMAL)
	return str_time
}

func (dateTime *DateTime) LogString() string {
	str_time := dateTime.dt.Format(FORMAT_TIME_LOG)
	return str_time
}

func (dateTime *DateTime) YearString() string {
	str_time := dateTime.dt.Format(FORMAT_TIME_YEAR)
	return str_time
}

func (dateTime *DateTime) MonthString() string {
	str_time := dateTime.dt.Format(FORMAT_TIME_MONTH)
	return str_time
}

func (dateTime *DateTime) DayString() string {
	str_time := dateTime.dt.Format(FORMAT_TIME_DAY)
	return str_time
}

func (dateTime *DateTime) HourString() string {
	str_time := dateTime.dt.Format(FORMAT_TIME_HOUR)
	return str_time
}

func (dateTime *DateTime) MinuteString() string {
	str_time := dateTime.dt.Format(FORMAT_TIME_MINUTE)
	return str_time
}

func (dateTime *DateTime) SecondString() string {
	str_time := dateTime.dt.Format(FORMAT_TIME_SECOND)
	return str_time
}

func (dateTime *DateTime) Before(other *DateTime) bool {
	return dateTime.dt.Before(other.dt)
}

func (dateTime *DateTime) Sub(other *DateTime) *DateTimeDuration {
	dtDuration := new(DateTimeDuration)
	dtDuration.duration = dateTime.dt.Sub(other.dt)
	return dtDuration
}

func (dateTime *DateTime) Add(seconds int) *DateTime {
	tmp := dateTime
	tmp.dt = tmp.dt.Add(time.Duration(seconds) * time.Second)
	return tmp
}

func (datetime *DateTime) AddDate(year int, month int, day int) {
	datetime.dt = datetime.dt.AddDate(year, month, day)
}

func (dateTime *DateTime) Parse(timestr string, format string) {
	dateTime.dt, _ = time.Parse(format, timestr)
}
