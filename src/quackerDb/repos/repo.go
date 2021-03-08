package repos

import (
	"github.com/aniladanir/quacker/quackerDb/model"
)

type Repo interface {
	Add(model model.Model) error
	Get(id int) (model.Model, error)
	Update(model model.Model) error
	Remove(id int) error
}
