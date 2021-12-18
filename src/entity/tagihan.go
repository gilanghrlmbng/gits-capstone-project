package entity

import (
	"net/http"
	"src/utils"
	"time"

	"gorm.io/gorm"
)

type Tagihan struct {
	Id         string          `gorm:"type:varchar(50);primaryKey" json:"id" form:"id"`
	IdKeluarga string          `gorm:"type:varchar(50);not null" json:"id_keluarga" form:"id_keluarga"`
	IdRT       string          `gorm:"type:varchar(50);not null" json:"id_rt" form:"id_rt"`
	Nama       string          `gorm:"type:varchar(50);not null" json:"nama" form:"nama"`
	Detail     string          `gorm:"not null" json:"detail" form:"detail"`
	Jumlah     int64           `gorm:"not null" json:"jumlah" form:"jumlah"`
	Terbayar   bool            `gorm:"default:false" json:"terbayar" form:"terbayar"`
	CreatedAt  time.Time       `gorm:"type:timestamptz;not null" json:"created_at"`
	UpdatedAt  time.Time       `gorm:"type:timestamptz;" json:"updated_at"`
	DeletedAt  *gorm.DeletedAt `json:"deleted_at,omitempty"`
}

func (Tagihan) TableName() string {
	return "tagihan"
}

func (t Tagihan) ValidateCreate() utils.Error {
	if t.Nama == "" {
		return utils.Error{
			Code:    http.StatusBadRequest,
			Message: "Nama tagigan tidak boleh kosong",
		}
	}
	if t.Detail == "" {
		return utils.Error{
			Code:    http.StatusBadRequest,
			Message: "Detail tagihan tidak boleh kosong",
		}
	}
	if t.Jumlah == 0 {
		return utils.Error{
			Code:    http.StatusBadRequest,
			Message: "Jumlah tagihan tidak boleh kosong",
		}
	}

	return utils.Error{}
}
