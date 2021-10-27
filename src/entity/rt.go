package entity

type Rt struct {
	Id            string       `gorm:"type:varchar(50);primaryKey" json:"id"`
	PengurusRT    []PengurusRT `gorm:"foreignKey:id_rt;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"pengurusRT"`
	NamaRt        string       `gorm:"type:varchar(50);not null" json:"nama_rt"`
	NamaRw        string       `gorm:"type:varchar(50);not null" json:"nama_rw"`
	Kelurahan     string       `gorm:"type:varchar(50);not null" json:"kelurahan"`
	Kecamatan     string       `gorm:"type:varchar(50);not null" json:"kecamatan"`
	Kota          string       `gorm:"type:varchar(50);not null" json:"kota"`
	Provinsi      string       `gorm:"type:varchar(50);not null" json:"provinsi"`
	Biaya_bulanan string       `gorm:"type:varchar(50);not null" json:"biaya_bulanan"`
}

func (Rt) TableName() string {
	return "rt"
}
