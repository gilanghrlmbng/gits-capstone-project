package entity

type Pembayaran struct {
	Id                string  `gorm:"type:varchar(50);primaryKey" json:"id" form:"id"`
	Order             []Order `gorm:"foreignKey:id_pembayaran;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"id_order,omitempty" form:"id_order"`
	Jumlah_pembayaran int64   `gorm:"not null" json:"jumlah_pembayaran" form:"jumlah_pembayaran"`
	Jenis             string  `gorm:"type:varchar(50);not null" json:"jenis" form:"jenis"`
}

func (Pembayaran) TableName() string {
	return "pembayaran"
}
