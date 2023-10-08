package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type InstanceInfo struct {
	InstanceInfoId uint `gorm:"primarykey"`
	Hostname       string
}

func main() {
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&InstanceInfo{})
	if err != nil {
		panic(err)
	}
}
