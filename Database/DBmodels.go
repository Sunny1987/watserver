package Database

import (
	"github.com/google/uuid"
)

type Scan struct {
	ID     uuid.UUID `gorm:"primary_key;unique;type:uuid;column:id;default:uuid_generate_v4()" json:"id"`
	Url    string    `json:"url"`
	Result string    `json:"result"`
}
