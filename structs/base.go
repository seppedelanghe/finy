package structs

import (
	"time"

	"github.com/google/uuid"
)

type BaseModel struct {
	Id string

	CreatedAt int64
	UpdatedAt int64
	DeletedAt int64
}

func (m *BaseModel) New() {
	m.Id = uuid.NewString()
	
	now := time.Now().UTC().Unix()
	m.CreatedAt = now
	m.UpdatedAt= now
	m.DeletedAt = -1
}
