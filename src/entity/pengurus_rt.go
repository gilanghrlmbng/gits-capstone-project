package entity

import (
	"net/http"
	"net/mail"
	"src/utils"
	"time"

	"gorm.io/gorm"
)

type PengurusRT struct {
	Id        string          `gorm:"type:varchar(50);primaryKey" json:"id" form:"id"`
	IdRT      string          `gorm:"type:varchar(50);not null" json:"id_rt" form:"id_rt"`
	Nama      string          `gorm:"type:varchar(50);not null" json:"nama" form:"nama"`
	Email     string          `gorm:"type:varchar(120);not null" json:"email" form:"email"`
	Password  string          `gorm:"type:varchar(100);not null" json:"password" form:"password"`
	CreatedAt time.Time       `gorm:"type:timestamptz;not null" json:"created_at"`
	UpdatedAt time.Time       `gorm:"type:timestamptz;" json:"updated_at"`
	DeletedAt *gorm.DeletedAt `json:"deleted_at,omitempty"`
}

func (PengurusRT) TableName() string {
	return "pengurus_rt"
}

func (prt PengurusRT) ValidateCreate() utils.Error {
	if prt.Nama == "" {
		return utils.Error{
			Code:    http.StatusBadRequest,
			Message: "Nama tidak boleh kosong",
		}
	}
	if _, err := mail.ParseAddress(prt.Email); err != nil {
		return utils.Error{
			Code:    http.StatusBadRequest,
			Message: "Email tidak valid",
		}
	}
	if prt.Password == "" {
		return utils.Error{
			Code:    http.StatusBadRequest,
			Message: "Password tidak boleh kosong",
		}
	}
	return utils.Error{}
}
