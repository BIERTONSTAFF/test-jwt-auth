package models

import (
	"github.com/google/uuid"
)

type RefreshToken struct {
	ID     int
	UserID uuid.UUID
	Token  string
	Valid  bool
}
