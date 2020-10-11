package domain

import "context"

type Resident struct {
	ID string `json:"id"`
	AuthType string `json:"type"`

}

type ResidentRepository interface {
	GetByID(ctx context.Context, id string) (Resident, error)
	DeleteByID(ctx context.Context, id string) (Resident, error)
}