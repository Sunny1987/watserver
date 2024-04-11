package Database

import (
	"github.com/google/uuid"
	"webserver/resultsapp"
)

type DBService interface {
	InitDB(dns string) error
	CreateResult(url string) (uuid.UUID, error)
	UpdateResults(id uuid.UUID, result []resultsapp.FinalResponse) error
}
