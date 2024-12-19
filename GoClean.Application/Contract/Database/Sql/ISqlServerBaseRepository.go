package Sql

import "gorm.io/gorm"

type ISqlBaseRepository[entity any] interface {
	GetById(id any) entity
	GetAll() []entity
	Create(model *entity) *entity
	Update(model *entity, id any) bool
	Delete(model *entity) bool

	GetContext() *gorm.DB
}
