package entity

import (
	"time"
)

type ForgetPasswordWarga struct {
	Id        string    `gorm:"type:varchar(50);primaryKey" json:"id" form:"id"`
	IdWarga   string    `gorm:"type:varchar(50)" json:"id_warga" form:"id_warga"`
	Kode      string    `gorm:"type:varchar(20)" json:"kode" form:"kode"`
	CreatedAt time.Time `gorm:"type:timestamptz;not null" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:timestamptz" json:"updated_at"`
}

func (ForgetPasswordWarga) TableName() string {
	return "forget_password_warga"
}
