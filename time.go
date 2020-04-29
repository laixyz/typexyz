package typexyz

import "time"

// GetToday 获取今天的起始时间
func GetToday() time.Time {
	t := time.Now()
	year, month, day := t.Date()
	today := time.Date(year, month, day, 0, 0, 0, 0, t.Location())
	return today
}

// TSWK 获取本周的起始时间
func TSWK() time.Time {
	t := time.Now()
	year, month, day := t.Date()
	var w int = int(t.Weekday())
	if w == 0 {
		w = 6
	} else {
		w = w - 1
	}
	return  time.Date(year, month, day-w, 0, 0, 0, 0, t.Location())
}

// LastDay 获取昨天起始与结束的unix_timestamp值
func LastDay() (time.Time, time.Time) {
	t := time.Now()
	year, month, day := t.Date()
	begin := time.Date(year, month, day-1, 0, 0, 0, 0, t.Location())
	end := time.Date(year, month, day, 0, 0, -1, 0, t.Location())
	return begin,end
}

// LastWeek 获取上一周起始与结束的unix_timestamp值
func LastWeek() (time.Time, time.Time) {
	t := time.Now()
	year, month, day := t.Date()
	var w int = int(t.Weekday())
	if w == 0 {
		w = 6
	} else {
		w = w - 1
	}
	begin := time.Date(year, month, day-w-7, 0, 0, 0, 0, t.Location())
	end := time.Date(year, month, day-w, 0, 0, -1, 0, t.Location())
	return begin,end
}

// GetDate 获取指定参数的时间unix_timestamp值
func GetDate(year, month, day, hour, min, sec, nsec int) time.Time {
	t := time.Now()
	return time.Date(year, time.Month(month), day, hour, min, sec, nsec, t.Location())
}

// ThisMonth 获取本月的起始时间unix_timestamp值
func ThisMonth() time.Time {
	t := time.Now()
	year, month, _ := t.Date()
	return time.Date(year, month, 1, 0, 0, 0, 0, t.Location())
}

// ThisYear 获取本月的起始时间unix_timestamp值
func ThisYear() time.Time {
	t := time.Now()
	year, _, _ := t.Date()
	return time.Date(year, 1, 1, 0, 0, 0, 0, t.Location())
}

