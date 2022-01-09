package db

import (
	"math/rand"
	"src/entity"
	"time"

	"github.com/oklog/ulid/v2"
	"gorm.io/gorm"
)

func SeedKeranjang(db *gorm.DB, idWarga []string) []string {
	var idKeranjang []string
	for _, val := range idWarga {
		entropy1 := ulid.Monotonic(rand.New(rand.NewSource(time.Now().UnixNano())), 0)
		Id := ulid.MustNew(ulid.Timestamp(time.Now()), entropy1).String()

		k := entity.Keranjang{
			Id:      Id,
			IdWarga: val,
		}
		db.Create(&k)

		idKeranjang = append(idKeranjang, Id)
	}

	return idKeranjang
}
