package db

import (
	"math/rand"
	"src/entity"
	"time"

	"github.com/oklog/ulid/v2"
	"gorm.io/gorm"
)

func SeedDompetRT(db *gorm.DB, listRT []string) []string {
	var idDompet []string
	for _, val := range listRT {
		entropy1 := ulid.Monotonic(rand.New(rand.NewSource(time.Now().UnixNano())), 0)
		Id := ulid.MustNew(ulid.Timestamp(time.Now()), entropy1).String()

		d := entity.DompetRT{
			Id:     Id,
			IdRT:   val,
			Jumlah: 500000,
		}
		db.Create(&d)

		idDompet = append(idDompet, Id)
	}

	return idDompet
}
