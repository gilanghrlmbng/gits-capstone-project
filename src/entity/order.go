package entity

import (
	"net/http"
	"src/utils"
	"time"

	"gorm.io/gorm"
)

type Order struct {
	Id           string          `gorm:"type:varchar(50);primaryKey" json:"id" form:"id"`
	IdWarga      string          `gorm:"type:varchar(50);not null" json:"id_warga" form:"id_warga"`
	IdPembayaran string          `gorm:"type:varchar(50);not null" json:"id_pembayaran" form:"id_pembayaran"`
	ItemOrder    []ItemOrder     `gorm:"foreignKey:id_order;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"item_order,omitempty" form:"item_order"`
	Harga_total  int64           `gorm:"not null" json:"harga_total" form:"harga_total"`
	Status       string          `gorm:"type:varchar(200);not null" json:"status" form:"status"`
	CreatedAt    time.Time       `gorm:"type:timestamptz;not null" json:"created_at"`
	UpdatedAt    time.Time       `gorm:"type:timestamptz;" json:"updated_at"`
	DeletedAt    *gorm.DeletedAt `json:"deleted_at,omitempty"`
}

func (Order) TableName() string {
	return "orders"
}

func (ord Order) ValidateCreate() utils.Error {
	if ord.Harga_total == 0 {
		return utils.Error{
			Code:    http.StatusBadRequest,
			Message: "Harga Total tidak boleh 0",
		}
	}
	return utils.Error{}
}
