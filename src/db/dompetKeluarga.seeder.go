package db

import (
	"fmt"
	"math/rand"
	"src/entity"
	"time"

	"github.com/oklog/ulid/v2"
	"gorm.io/gorm"
)

func SeedDompetKeluarga(db *gorm.DB, listIdKeluarga []string) []string {
	var idDompet []string
	for idx, val := range listIdKeluarga {
		fmt.Println(idx)
		entropy1 := ulid.Monotonic(rand.New(rand.NewSource(time.Now().UnixNano())), 0)
		Id := ulid.MustNew(ulid.Timestamp(time.Now()), entropy1).String()

		d := entity.DompetKeluarga{
			Id:         Id,
			IdKeluarga: val,
			Jumlah:     500000,
		}
		db.Create(&d)

		idDompet = append(idDompet, Id)
	}

	return idDompet
}
