package GoClean_Domain

import "gorm.io/gorm"

type GormConn struct {
	DB *gorm.DB
}
