package Sql

import "GoClean/GoClean.Domain/Entities/Sql"

type ISqlSampleEntityRepository interface {
	ISqlBaseRepository[Sql.SqlSampleEntity]
}
