package Database

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
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

func (dBundle *DBBundle) InitDB(dsn string) error {

	//create connection
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		dBundle.logger.Fatalf("db not initialized %v", err)
		return err
	}
	//migrate schema
	if err = db.AutoMigrate(&Scan{}); err != nil {
		dBundle.logger.Fatalf("failed to migrate %v", err)
		return err
	}
	dBundle.db = db
	return nil
}

func (dBundle *DBBundle) CreateResult(url string) (uuid.UUID, error) {
	scan := Scan{Url: url}
	tx := dBundle.db.Table("scans").Create(&scan)
	if tx.Error != nil {
		dBundle.logger.Printf("creating record failed %v", tx.Error)
		return uuid.UUID{}, tx.Error
	}
	dBundle.logger.Printf("Record updated successfully RowsAffected=%v", tx.RowsAffected)
	return scan.ID, nil
}

func (dBundle *DBBundle) UpdateResults(id uuid.UUID, result []resultsapp.FinalResponse) error {
	var scan Scan
	if err := dBundle.db.Table("scans").First(&scan, id).Error; err != nil {
		dBundle.logger.Printf("updating record failed %v", err)
		return err
	}
	jsondata, err := json.Marshal(scan)
	if err != nil {
		dBundle.logger.Printf("json marshalling failed %v", err)
		return err
	}
	//smallData, err := compressData(jsondata)
	//if err != nil {
	//	dBundle.logger.Printf("compression failed %v", err)
	//}
	//if !utf8.ValidString(string(jsondata)) {
	//	dBundle.logger.Println("invalid UTF8 encoding")
	//	return errors.New("invalid UTF8 encoding")
	//}
	scan.Result = string(jsondata)
	tx := dBundle.db.Table("scans").Save(&scan)
	if tx.Error != nil {
		dBundle.logger.Printf("updating record failed %v", tx.Error)
		return tx.Error
	}
	dBundle.logger.Printf("Record updated successfully RowsAffected=%v", tx.RowsAffected)
	return nil
}

func compressData(data []byte) ([]byte, error) {
	var compressed bytes.Buffer
	gz := gzip.NewWriter(&compressed)
	_, err := gz.Write(data)
	if err != nil {
		return nil, err
	}
	gz.Close()
	return compressed.Bytes(), nil
}
