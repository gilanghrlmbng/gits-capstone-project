package db

import (
	"math/rand"
	"src/entity"
	"src/utils"
	"time"

	"github.com/oklog/ulid/v2"
	"gorm.io/gorm"
)

func SeedWarga(db *gorm.DB, listKeluarga []string) []string {
	// Data 1
	entropy1 := ulid.Monotonic(rand.New(rand.NewSource(time.Now().UnixNano())), 0)
	Id1 := ulid.MustNew(ulid.Timestamp(time.Now()), entropy1).String()
	data1 := entity.Warga{
		Id:          Id1,
		IdKeluarga:  listKeluarga[0],
		Nama:        "Agustina",
		Email:       "warga1@gmail.com",
		Gender:      "perempuan",
		NoHandphone: "08123123123123",
		Gambar: "https://dummyimage.com/500x500/eee/fff&text=A",
		Password:    "aaaaaa",
		CreatedAt:   time.Now(),
	}
	data1.Password = utils.HashPassword(data1.Password, data1.Id)

	db.Create(&data1)

	// Data 2
	entropy2 := ulid.Monotonic(rand.New(rand.NewSource(time.Now().UnixNano())), 0)
	Id2 := ulid.MustNew(ulid.Timestamp(time.Now()), entropy2).String()
	data2 := entity.Warga{
		Id:          Id2,
		IdKeluarga:  listKeluarga[0],
		Nama:        "Ronals",
		Email:       "warga2@gmail.com",
		Gender:      "laki-laki",
		NoHandphone: "08123123123123",
		Gambar: "https://dummyimage.com/500x500/eee/fff&text=A",
		Password:    "PasswordWargaSaya",
		CreatedAt:   time.Now(),
	}
	data2.Password = utils.HashPassword(data2.Password, data2.Id)

	db.Create(&data2)

	// Data 3
	entropy3 := ulid.Monotonic(rand.New(rand.NewSource(time.Now().UnixNano())), 0)
	Id3 := ulid.MustNew(ulid.Timestamp(time.Now()), entropy3).String()
	data3 := entity.Warga{
		Id:          Id3,
		IdKeluarga:  listKeluarga[1],
		Nama:        "mamat",
		Email:       "warga3@gmail.com",
		Gender:      "laki-laki",
		NoHandphone: "08123123123123",
		Gambar: "https://dummyimage.com/500x500/eee/fff&text=M",
		Password:    "PasswordWargaSaya",
		CreatedAt:   time.Now(),
	}
	data3.Password = utils.HashPassword(data3.Password, data3.Id)

	db.Create(&data3)

	return []string{Id1, Id2, Id3}
}
