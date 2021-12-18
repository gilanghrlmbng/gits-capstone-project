package entity

import (
	"time"
)

type ForgetPasswordPengurus struct {
	Id         string    `gorm:"type:varchar(50);primaryKey" json:"id" form:"id"`
	IdPengurus string    `gorm:"type:varchar(50)" json:"id_pengurus" form:"id_pengurus"`
	Kode       string    `gorm:"type:varchar(20)" json:"kode" form:"kode"`
	CreatedAt  time.Time `gorm:"type:timestamptz;not null" json:"created_at"`
	UpdatedAt  time.Time `gorm:"type:timestamptz" json:"updated_at"`
}

func (ForgetPasswordPengurus) TableName() string {
	return "forget_password_pengurus"
}
