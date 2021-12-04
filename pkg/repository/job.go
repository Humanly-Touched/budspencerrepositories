package repository

import (
	"github.com/Humanly-Touched/budspencerrepositories/pkg/models"
)

// Repository interface with the contract to be implemented by
// the concrete db access layer.
type Job interface {
	Get(uuid string) (models.Job, error)
	GetAll() ([]models.Job, error)
	Create(message string) (models.Job, error)
	Update(job models.Job) (models.Job, error)
	Upsert(job models.Job) (models.Job, error)
	Delete(uuid string) error
}
