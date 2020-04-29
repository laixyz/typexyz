package typexyz

import (
	"database/sql"
	"database/sql/driver"
	"time"
)

type DateTime struct {
	Time time.Time
}

func NewDateTime() Date {
	var newdate Date
	newdate.Time = time.Now()
	return newdate
}

func (t DateTime) IsNull() bool {
	return t.Time.Unix() == 0
}

func (t DateTime) String() string {
	if t.Time.Unix() == 0 {
		return DateTimeNULL
	}
	return t.Time.Format(DateTimeFormat)
}

// Value 用于mysql int类型字段存取unix_timestamp值
func (t *DateTime) Value() (driver.Value, error) {
	return t.Time.Unix(), nil
}

// Scan 用于mysql int类型字段存取unix_timestamp值
func (t *DateTime) Scan(src interface{}) error {
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
