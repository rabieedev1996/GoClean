package GoClean_Common

import (
	ptime "github.com/yaa110/go-persian-calendar"
	"time"
)

type JalaliDateTool struct {
	Time ptime.Time
}

func NewJalaliDateTool(time ptime.Time) *JalaliDateTool {
	return &JalaliDateTool{
		Time: time,
	}

}

func (dt JalaliDateTool) Format(format string) string {
	return dt.Time.Format(format)
}
func (dt *JalaliDateTool) AddDay(days int) {
	dt.Time = dt.Time.AddDate(0, 0, days)
}

func (dt *JalaliDateTool) AddMonth(month int) {
	dt.Time = dt.Time.AddDate(0, month, 0)
}
func (dt *JalaliDateTool) AddYear(years int) {
	dt.Time = dt.Time.AddDate(years, 0, 0)
}

func (dt *JalaliDateTool) AddMinute(minutes int) {
	dt.Time = dt.Time.Add(time.Duration(minutes) * time.Minute)
}

func (dt *JalaliDateTool) AddSecond(seconds int) {
	dt.Time = dt.Time.Add(time.Duration(seconds) * time.Second)
}
func (dt *JalaliDateTool) ToMiladiDateToll() *MiladiDateTool {
	return NewMiladiDateTool(dt.Time.Time())
}
