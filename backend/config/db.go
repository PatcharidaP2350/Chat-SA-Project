package config

import (
	"SA-67-SongThor-SUT/entity"
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func DB() *gorm.DB {
	return db
}

func ConnectionDB() {
	database, err := gorm.Open(sqlite.Open("sa.db?cache=shared"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("connected database")
	db = database
}

func SetupDatabase() {

	db.AutoMigrate(
		&entity.Member{},
		&entity.RoomChat{},
		&entity.Seller{},
		&entity.Message{},
	)

	
	Member := &entity.Member{
		FirstName: "Software",
		LastName:  "Analysis",
		Email:     "sa@gmail.com",
		Username: "Jibjib",
		Password: "12345",
		PhoneNumber: "021313343",
		Address: "Mittrphap Road Korat",
		PicProfile: "ksdkgsgs",
		
	}
	db.FirstOrCreate(Member, &entity.Member{
		Email: "sa@gmail.com",
	})

}
