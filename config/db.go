package config

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"project-kp/model"
)

var DB *gorm.DB

func InitDb() {
	// Configure DB Connection
	var err error
	dsn := fmt.Sprintf("host=ec2-52-204-72-14.compute-1.amazonaws.com user=mgyaxhlhqhaudn password=7099e16fa30948963bf9643a2a0a8fa321eb66704559eba65e4a71263a54ee8d dbname=d1sq6oe1fk26aq port=5432 sslmode=require TimeZone=Asia/Jakarta")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		panic("config.InitDb, Failed to connect database")
	}

	autoMigrate()
}

func autoMigrate() {
	// Auto Migrate Schema
	err := DB.AutoMigrate(
		model.Uom{},
	)

	if err != nil {
		log.Fatalf("config.autoMigrate, error auto migrate: %v", err)
	}
}
