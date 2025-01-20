package Sql

type SqlSampleEntity struct {
	*SqlBaseEntity
}

func NewSqlSampleEntity() SqlSampleEntity {
	return SqlSampleEntity{
		SqlBaseEntity: NewSqlBaseEntity(),
	}
}
