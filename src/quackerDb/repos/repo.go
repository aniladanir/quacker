package repos

import (
	"github.com/aniladanir/quacker/quackerDb/models"
)

type Repo interface {
	Add(model models.Model) error
	Get(id int) (models.Model, error)
	Update(model models.Model) error
	Remove(id int) error
}
