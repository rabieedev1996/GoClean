package Sql

import (
	GoClean_Domain "GoClean/GoClean.Domain"
	"fmt"
	"gorm.io/gorm"
	"reflect"
	"strings"
)

type SqlBaseRepository[entity any] struct {
	Context *gorm.DB
}

func NewSqlBaseRepository[T any](conn GoClean_Domain.GormConn) *SqlBaseRepository[T] {
	return &SqlBaseRepository[T]{
		Context: conn.DB,
	}
}

func (r SqlBaseRepository[entity]) Init(guid string) {
	//databaseConnection := InjectionConfig.PrepareObject[Config.DatabaseConnection](reflect.TypeOf(Config.DatabaseConnection{}), guid)
	//r.Context = databaseConnection.Context
}

func (r SqlBaseRepository[entity]) GetById(id any) entity {
	titleArray := strings.Split(fmt.Sprintf("%T", *new(entity)), ".")
	var model entity
	r.Context.Table(titleArray[len(titleArray)-1]).Where("id = ?", id).First(&model)
	return model
}
func (r SqlBaseRepository[entity]) GetAll() []entity {
	titleArray := strings.Split(fmt.Sprintf("%T", *new(entity)), ".")
	model := []entity{}

	hasSoftDelete := false
	var temp entity
	var softDeleteFieldName string
	entityType := reflect.TypeOf(temp)
	for i := 0; i < entityType.NumField(); i++ {
		if entityType.Field(i).Name == "Is_Deleted" {
			hasSoftDelete = true
			softDeleteFieldName = "Is_Deleted"
			break
		}
		if entityType.Field(i).Name == "Is_Delete" {
			hasSoftDelete = true
			softDeleteFieldName = "Is_Delete"
			break
		}
	}
	if hasSoftDelete {
		r.Context.Table(titleArray[len(titleArray)-1]).Where(softDeleteFieldName + "=false").Find(&model)
	} else {
		r.Context.Table(titleArray[len(titleArray)-1]).Find(model)
	}
	return model
}
func (r SqlBaseRepository[entity]) Create(model *entity) *entity {
	titleArray := strings.Split(fmt.Sprintf("%T", *new(entity)), ".")
	r.Context.Table(titleArray[len(titleArray)-1]).Create(model)
	return model
}

func (r SqlBaseRepository[entity]) Update(model *entity, id any) bool {
	titleArray := strings.Split(fmt.Sprintf("%T", *new(entity)), ".")
	r.Context.Table(titleArray[len(titleArray)-1]).Save(model)
	return true
}

func (r SqlBaseRepository[entity]) Delete(model *entity) bool {
	titleArray := strings.Split(fmt.Sprintf("%T", *new(entity)), ".")
	r.Context.Table(titleArray[len(titleArray)-1]).Delete(model)
	return true
}

func (r SqlBaseRepository[entity]) GetContext() *gorm.DB {
	return r.Context
}
