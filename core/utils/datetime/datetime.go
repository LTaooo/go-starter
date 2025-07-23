package datetime

import (
	"time"
)

type Datetime struct {
	Time time.Time
}

func FromTimestamp(timestamp int64) *Datetime {
	return &Datetime{
		Time: time.Unix(timestamp, 0),
	}
}

func FromDatetime(datetime string) (*Datetime, error) {
	t, err := time.Parse("2006-01-02 15:04:05", datetime)
	if err != nil {
		return nil, err
	}
	return &Datetime{
		Time: t,
	}, nil
}

func FromNow() *Datetime {
	return &Datetime{
		Time: time.Now(),
	}
}

func (d *Datetime) Format(layout string) string {
	return d.Time.Format(layout)
}

func (d *Datetime) Datetime() string {
	return d.Format("2006-01-02 15:04:05")
}

func (d *Datetime) Timestamp() int64 {
	return d.Time.Unix()
}

func (d *Datetime) Milisecond() int64 {
	return d.Time.UnixMilli()
}
