package repo

import (
	"github.com/aniladanir/quacker/quackerDb/model"
)

type Repo interface {
	Add(model model.Model) error
	Get(id int) (model.Model, error)
	Update(id int, model model.Model) error
	Remove(id int) (model.Model, error)
}
