package entity

import (
	"time"

	"github.com/google/uuid"
)

type ItemEntity struct {
	Id uuid.UUID
	Name string
	JanCode string
	CreatedAt time.Time
	Updatedat time.Time
}