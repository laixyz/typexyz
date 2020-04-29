package typexyz

import (
	"database/sql"
	"database/sql/driver"
	"time"
)

type Date struct {
	Time time.Time
}

func NewDate(params ...int) Date{
	var t time.Time
	var year,month,day int
	if len(params)==3{
		year = params[0]
		month = params[1]
		day=params[2]
		t=time.Date(year,time.Month(month),day,0,0,0,0,t.Location())
	} else {
		t=time.Now()
		t=time.Date(t.Year(),t.Month(),t.Day(),0,0,0,0,t.Location())
	}
	var newdate Date
	newdate.Time=t
	return newdate
}

func (t Date) IsNull() bool {
	return t.Time.Unix() == 0
}

func (t Date) String() string {
	if t.Time.Unix() == 0 {
		return DateNULL
	}
	return t.Time.Format(DateFormat)
}

// Value 用于mysql int类型字段存取unix_timestamp值
func (t *Date) Value() (driver.Value, error) {
	return t.Time.Unix(), nil
}

// Scan 用于mysql int类型字段存取unix_timestamp值
func (t *Date) Scan(src interface{}) error {
	if src == nil {
		t.Time = time.Unix(0, 0)
		return nil
	}
	var nullInt64 sql.NullInt64
	err := nullInt64.Scan(src)
	if err != nil {
		return err
	}
	t.Time = time.Unix(nullInt64.Int64, 0)
	return nil
}
