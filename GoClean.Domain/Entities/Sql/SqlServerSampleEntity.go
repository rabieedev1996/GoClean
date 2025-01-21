package Sql

type SqlSampleEntity struct {
	*SqlBaseEntity
}

func NewSqlSampleEntity() SqlSampleEntity {
	return SqlSampleEntity{
		SqlBaseEntity: NewSqlBaseEntity(),
	}
}
func (SqlSampleEntity) TableName() string {
	return "SqlSampleEntity" // نام مورد نظر شما
}
