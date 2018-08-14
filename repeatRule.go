package icalendar

// RepeatRule represents a repeating rule
type RepeatRule struct {
	//FREQ=YEARLY;INTERVAL=1;BYMONTH=2;BYMONTHDAY=12
	Frequency  Frequency
	Interval   int64
	BySecond   int64
	ByMinute   int64
	ByHour     int64
	ByDay      int64
	ByMonth    int64
	ByMonthDay int64
	ByYearDay  int64
	ByWeekNo   int64
	BySetPos   int64
}

func (r *RepeatRule) String() (out string) {
	var buf []byte
	buf = append(buf, "FREQ="+r.Frequency+";"...)
	buf = appendInt64(buf, "INTERVAL", r.Interval, ";")
	buf = appendInt64(buf, "BYSECOND", r.Interval, ";")
	buf = appendInt64(buf, "BYMINUTE", r.Interval, ";")
	buf = appendInt64(buf, "BYHOUR", r.Interval, ";")
	buf = appendInt64(buf, "BYDAY", r.Interval, ";")
	buf = appendInt64(buf, "BYMONTH", r.Interval, ";")
	buf = appendInt64(buf, "BYMONTHDAY", r.Interval, ";")
	buf = appendInt64(buf, "BYYEARDAY", r.Interval, ";")
	buf = appendInt64(buf, "BYWEEKNO", r.Interval, ";")
	buf = appendInt64(buf, "BYSETPOS", r.Interval, ";")
	return string(buf)
}
