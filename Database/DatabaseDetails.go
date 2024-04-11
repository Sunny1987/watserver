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

// DBBundle is the object for DB functions
type DBBundle struct {
	logger *log.Logger
	db     *gorm.DB
}

// NewDBBundle is the constructor for DBBundle
func NewDBBundle(logger *log.Logger) DBService {
	return &DBBundle{logger: logger}
}

// InitDB will initiate DB connection and add schema for scans
func (dBundle *DBBundle) InitDB(dsn string) error {

	//create connection
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		dBundle.logger.Fatalf("db not initialized %v", err)
		return err
	}
	//migrate scan schema
	if err = db.AutoMigrate(&Scan{}); err != nil {
		dBundle.logger.Fatalf("failed to migrate %v", err)
		return err
	}
	dBundle.db = db
	return nil
}

// CreateResultForScan will create scans records
func (dBundle *DBBundle) CreateResultForScan(url string) (uuid.UUID, error) {
	scan := Scan{Url: url}
	tx := dBundle.db.Table("scans").Create(&scan)
	if tx.Error != nil {
		dBundle.logger.Printf("creating record failed %v", tx.Error)
		return uuid.UUID{}, tx.Error
	}
	dBundle.logger.Printf("Record updated successfully RowsAffected=%v", tx.RowsAffected)
	return scan.ID, nil
}

// UpdateResultsForScan will update the scan records
func (dBundle *DBBundle) UpdateResultsForScan(id uuid.UUID, result []resultsapp.FinalResponse) error {
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

// GetRecordsForScan will fetch all records from scan
func (dBundle *DBBundle) GetRecordsForScan() (string, error) {
	var scans []Scan
	if err := dBundle.db.Table("scans").Find(&scans).Error; err != nil {
		dBundle.logger.Printf("getting record failed %v", err)
		return "", err
	}
	resp, err := json.MarshalIndent(scans, "", " ")
	if err != nil {
		dBundle.logger.Printf("json marshalling failed %v", err)
		return "", err
	}
	return string(resp), nil
}

func (dBundle *DBBundle) GetRecordForId(id string) (string, error) {
	var scan Scan
	parseId, err := uuid.Parse(id)
	if err != nil {
		dBundle.logger.Printf("parsing record failed %v", err)
		return "", err
	}

	if err := dBundle.db.Table("scans").First(&scan, parseId).Error; err != nil {
		dBundle.logger.Printf("getting record failed %v", err)
		return "", err
	}
	resp, err := json.MarshalIndent(scan, "", " ")
	if err != nil {
		dBundle.logger.Printf("json marshalling failed %v", err)
		return "", err
	}
	return string(resp), nil
}
