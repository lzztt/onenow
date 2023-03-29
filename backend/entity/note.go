package entity

import (
	"time"
)

type Note struct {
	Id             int64
	Body           string
	CreateTime     time.Time
	LastUpdateTime time.Time
}
