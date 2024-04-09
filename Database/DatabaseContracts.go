package Database

import (
	"github.com/google/uuid"
	"webserver/resultsapp"
)

type DBService interface {
	InitDB()
	CreateResult(url string) uuid.UUID
	UpdateResults(id uuid.UUID, result []resultsapp.FinalResponse)
}
