package Sql

import (
	Rasa_Common "GoClean/GoClean.Common"
	"strconv"
	"time"
)

type SqlBaseEntity struct {
	Id int64 `gorm:" primaryKey;column: Id"`
	///for MSSQL
	//CreatedOn    time.Time `gorm:" column: CreatedOn;type:datetime"`
	///for PGSQL
	CreatedOn    time.Time `gorm:" column: CreatedOn;"`
	CreatedOnFa  int64     `gorm:" column: CreatedOnFa"`
	Time         string    `gorm:" column: Time"`
	IsDeleted    bool      `gorm:" column: IsDeleted"`
	FullDateTime int64     `gorm:" column: FullDateTime"`
}

func NewSqlBaseEntity() *SqlBaseEntity {
	now := time.Now()

	miladi := *Rasa_Common.NewMiladiDateTool(now)
	jalali := *miladi.ToJalaliDateToll()

	fullDateTime, _ := strconv.ParseInt(miladi.Format("yyyyMMddHHmmss"), 10, 64)
	//dateTime, _ := strconv.ParseInt(miladi.Format("yyyyMMdd"), 10, 64)
	pdateTime, _ := strconv.ParseInt(jalali.Format("yyyyMMdd"), 10, 64)

	return &SqlBaseEntity{
		CreatedOn:    miladi.Time,
		CreatedOnFa:  pdateTime,
		Time:         miladi.Format("HH:mm:ss"),
		IsDeleted:    false,
		FullDateTime: fullDateTime,
	}
}
