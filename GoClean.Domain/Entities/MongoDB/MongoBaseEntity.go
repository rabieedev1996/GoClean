package MongoDB

import (
	Rasa_Common "GoClean/GoClean.Common"
	"strconv"
	"time"
)

type MongoBaseEntity struct {
	CreatedOn    int64
	CreatedOnFa  int64
	Time         string
	IsDeleted    bool
	FullDateTime int64
}

func NewMongoBaseEntity() *MongoBaseEntity {
	now := time.Now()

	miladi := *Rasa_Common.NewMiladiDateTool(now)
	jalali := *miladi.ToJalaliDateToll()

	fullDateTime, _ := strconv.ParseInt(miladi.Format("yyyyMMddHHmmss"), 10, 64)
	dateTime, _ := strconv.ParseInt(miladi.Format("yyyyMMdd"), 10, 64)
	pdateTime, _ := strconv.ParseInt(jalali.Format("yyyyMMdd"), 10, 64)

	return &MongoBaseEntity{
		CreatedOn:    dateTime,
		CreatedOnFa:  pdateTime,
		Time:         miladi.Format("HH:mm:ss"),
		IsDeleted:    false,
		FullDateTime: fullDateTime,
	}
}
