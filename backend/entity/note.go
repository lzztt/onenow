package entity

import (
	"time"

	"github.com/google/uuid"
)

type Note struct {
	Id             uuid.UUID
	Body           string
	CreateTime     time.Time
	LastUpdateTime time.Time
}
