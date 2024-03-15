package models

import "github.com/gofrs/uuid"

type User struct {
	ID uuid.UUID `json:"id"`
}
