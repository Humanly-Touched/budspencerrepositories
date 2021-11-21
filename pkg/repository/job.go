package repositoriy

import (
	"github.com/profe-ajedrez/budspencerrepositories/pkg/models"
)

// Repository interface with the contract to be implemented by
// the concrete db access layer.
type Job interface {
	Get(uuid string) (models.Job, error)
	GetAll() ([]models.Job, error)
	Create(job models.Job) (models.Job, error)
	Update(job models.Job) (models.Job, error)
	Upsert(job models.Job) (models.Job, error)
	Delete(uuid string) error
}
