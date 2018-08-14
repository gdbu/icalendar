package icalendar

// RepeatRule represents a repeating rule
type RepeatRule struct {
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
	// Set frequency
	buf = append(buf, "FREQ="+r.Frequency+";"...)
	// Set interval
	buf = appendInt64(buf, "INTERVAL", r.Interval, ";")
	// Set by second
	buf = appendInt64(buf, "BYSECOND", r.Interval, ";")
	// Set by minute
	buf = appendInt64(buf, "BYMINUTE", r.Interval, ";")
	// Set by hour
	buf = appendInt64(buf, "BYHOUR", r.Interval, ";")
	// Set by day
	buf = appendInt64(buf, "BYDAY", r.Interval, ";")
	// Set by month
	buf = appendInt64(buf, "BYMONTH", r.Interval, ";")
	// Set by month-day
	buf = appendInt64(buf, "BYMONTHDAY", r.Interval, ";")
	// Set by year-day
	buf = appendInt64(buf, "BYYEARDAY", r.Interval, ";")
	// Set by weekend
	buf = appendInt64(buf, "BYWEEKNO", r.Interval, ";")
	// Set by set position
	buf = appendInt64(buf, "BYSETPOS", r.Interval, ";")
	// Convert the byteslice buffer to a string and return
	return string(buf)
}
