package Database

import (
	"github.com/google/uuid"
	"webserver/resultsapp"
)

type Scan struct {
	ID     uuid.UUID                  `gorm:"primary_key;unique;type:uuid;column:id;default:uuid_generate_v4()" json:"id"`
	Url    string                     `json:"url"`
	Result []resultsapp.FinalResponse `json:"result"`
}
