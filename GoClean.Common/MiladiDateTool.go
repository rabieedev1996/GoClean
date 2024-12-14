package GoClean_Common

import (
	ptime2 "github.com/yaa110/go-persian-calendar"
	"strings"
	"time"
)

type MiladiDateTool struct {
	Time time.Time
}

func NewMiladiDateTool(time time.Time) *MiladiDateTool {
	return &MiladiDateTool{
		Time: time,
	}

}

func (dt MiladiDateTool) Format(format string) string {
	replacements := []struct {
		from string
		to   string
	}{
		{"yyyy", "2006"},
		{"yy", "06"},
		{"MM", "01"},
		{"M", "1"},
		{"dd", "02"},
		{"d", "2"},
		{"HH", "15"},
		{"H", "3"},
		{"mm", "04"},
		{"m", "4"},
		{"ss", "05"},
		{"s", "5"},
		{"tt", "PM"},
	}

	for _, r := range replacements {
		format = strings.ReplaceAll(format, r.from, r.to)
	}

	return dt.Time.Format(format)
}
func (dt *MiladiDateTool) AddDay(days int) {
	dt.Time = dt.Time.AddDate(0, 0, days)
}

func (dt *MiladiDateTool) AddMonth(month int) {
	dt.Time = dt.Time.AddDate(0, month, 0)
}
func (dt *MiladiDateTool) AddYear(years int) {
	dt.Time = dt.Time.AddDate(years, 0, 0)
}

func (dt *MiladiDateTool) AddMinute(minutes int) {
	dt.Time = dt.Time.Add(time.Duration(minutes) * time.Minute)
}

func (dt *MiladiDateTool) AddSecond(seconds int) {
	dt.Time = dt.Time.Add(time.Duration(seconds) * time.Second)
}
func (dt *MiladiDateTool) ToJalaliDateToll() *JalaliDateTool {
	ptime := ptime2.New(dt.Time)
	return NewJalaliDateTool(ptime)
}
