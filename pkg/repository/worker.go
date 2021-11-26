package repository

import (
	"github.com/Humanly-Touched/budspencerrepositories/pkg/models"
)

// Repository interface with the contract to be implemented by
// the concrete db access layer.
type Worker interface {
	Get(uuid string) (models.Worker, error)
	GetAll() ([]models.Worker, error)
	Create(job models.Worker) (models.Worker, error)
	Update(job models.Worker) (models.Worker, error)
	Upsert(job models.Worker) (models.Worker, error)
	Delete(uuid string) error
}
