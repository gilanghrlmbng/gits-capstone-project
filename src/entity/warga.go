package entity

type Warga struct {
	Id         string `gorm:"type:varchar(50);primaryKey" json:"id"`
	IdKeluarga string `gorm:"type:varchar(50);not null" json:"id_keluarga"`
	Nama       string `gorm:"type:varchar(50);not null" json:"nama"`
	Alamat     string `gorm:"type:varchar(50);not null" json:"alamat"`
}

func (Warga) TableName() string {
	return "warga"
}
