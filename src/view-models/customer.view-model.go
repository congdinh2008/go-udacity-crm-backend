package viewmodels

import "github.com/google/uuid"

type CustomerViewModel struct {
	ID        uuid.UUID
	Name      string
	Role      string
	Email     string
	Phone     string
	Contacted bool
}
