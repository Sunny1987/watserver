package Database

import (
	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"webserver/resultsapp"
)

type DBBundle struct {
	logger *log.Logger
	db     *gorm.DB
}

func NewDBBundle(logger *log.Logger) DBService {
	return &DBBundle{logger: logger}
}

func (dBundle *DBBundle) InitDB() {
	dsn := resultsapp.GetEnvValueFor("DSN")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		dBundle.logger.Fatalf("db not initialized %v", err)
	}
	dBundle.db = db
}

func (dBundle *DBBundle) CreateResult(url string) uuid.UUID {
	scan := Scan{Url: url}
	tx := dBundle.db.Create(&scan)
	if tx.Error != nil {
		dBundle.logger.Printf("creating record failed %v", tx.Error)
	}
	dBundle.logger.Printf("Record updated successfully RowsAffected=%v", tx.RowsAffected)
	return scan.ID
}

func (dBundle *DBBundle) UpdateResults(id uuid.UUID, result []resultsapp.FinalResponse) {
	var scan Scan
	if err := dBundle.db.First(&scan, id).Error; err != nil {
		dBundle.logger.Fatalf("updating record failed %v", err)
	}
	scan.Result = result
	dBundle.db.Save(&scan)
}
