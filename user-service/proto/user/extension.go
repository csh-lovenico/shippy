package user

import (
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
)

func (m *User) BeforeCreate(scope *gorm.Scope) error {
	myUuid := uuid.NewV4()
	return scope.SetColumn("Id", myUuid.String())
}
