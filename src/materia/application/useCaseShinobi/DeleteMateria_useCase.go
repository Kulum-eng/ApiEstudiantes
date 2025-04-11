package useCaseShinobi

import (
	"API-HEXAGONAL/src/materia/domain"
)

type DeleteShinobi struct {
	db domain.IShinobi
}

func NewDeleteShinobi(db domain.IShinobi) *DeleteShinobi {
	return &DeleteShinobi{db: db}
}

func (delete *DeleteShinobi) Run(id int32) (int64, error) {
	return delete.db.DeleteShinobi(id)
}
