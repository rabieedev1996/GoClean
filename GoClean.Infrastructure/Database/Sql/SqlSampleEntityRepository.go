package Sql

import "GoClean/GoClean.Domain/Entities/Sql"

type SqlSampleEntityRepository struct {
	*SqlBaseRepository[Sql.SqlSampleEntity]
}
