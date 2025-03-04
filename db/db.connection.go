package db

import (
	"log"
	"os"

	"app.myriadflow.com/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	db, err := Connect()
	if err != nil {
		log.Println(err)
		log.Fatal(err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		log.Println(err)
		log.Fatal(err)
	}
	sqlDB.SetMaxIdleConns(50)
	sqlDB.SetMaxOpenConns(200)

	err = sqlDB.Ping()
	if err != nil {
		log.Println(err)
		log.Fatal(err)
	}
	log.Println("Successfully connected to the database!")

	err = InitMigration(db)
	if err != nil {
		log.Println(err)
	}

	err = db.AutoMigrate(
		&models.User{},
		&models.Brand{},
		&models.Collection{},
		&models.Phygital{},
		&models.WebXR{},
		&models.Avatar{},
		&models.Variant{},
		&models.FanToken{},
		&models.ChainType{},
		&models.NftEntries{},
		&models.Profile{},
		&models.CartItem{},
		&models.OTPStore{},
		&models.OTPData{},
		&models.MainnetFanToken{},
		&models.DelegateMintFanTokenRequest{},
		&models.Elevate{},
		&models.Address{},
		&models.Agent{},
	)

	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	} else {
		log.Println("Database migrated successfully")
		if err := UniqueConstraints(db); err != nil {
			log.Printf("failed to create unique constraints: %v", err)
		} else {
			log.Println("Unique constraints created successfully")
		}
	}
}

func Connect() (*gorm.DB, error) {
	// dsn := "host=satao.db.elephantsql.com user=atvcirnc password=fFa7mq_RGHrfMQ1tvfsNanUhjF96sbCk dbname=atvcirnc port=5432 sslmode=disable"
	dsn := "host=" + os.Getenv("DB_HOST") + " port=" + os.Getenv("DB_PORT") + " user=" + os.Getenv("DB_USERNAME") + " password=" + os.Getenv("DB_PASSWORD") + " dbname=" + os.Getenv("DB_NAME") + " sslmode=" + os.Getenv("DB_SSL_MODE") + ""

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v\n", err)
		return nil, err
	} else {
		log.Println("Connected to the database")
	}

	return DB, nil
}
func InitMigration(db *gorm.DB) error {
	return db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp"`).Error
}

func UniqueConstraints(db *gorm.DB) error {
	uniqueSql := `ALTER TABLE brands ADD CONSTRAINT unique_slug_name UNIQUE (slug_name);`
	return db.Exec(uniqueSql).Error
}
