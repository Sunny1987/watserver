package Database

import (
	"github.com/google/uuid"
	"webserver/resultsapp"
)

type DBService interface {
	InitDB(dns string) error
	CreateResultForScan(url string) (uuid.UUID, error)
	UpdateResultsForScan(id uuid.UUID, result []resultsapp.FinalResponse) error
	GetRecordsForScan() (string, error)
	GetRecordForId(id string) (string, error)
}
