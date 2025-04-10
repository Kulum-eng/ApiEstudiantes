package useCaseJutsu

import (
	"API-HEXAGONAL/src/jutsu/domain"
	"API-HEXAGONAL/src/jutsu/domain/entities"
)

type CreateJutsu struct {
	db domain.IJutsu
}

func NewCreateJutsu(db domain.IJutsu) *CreateJutsu {
	return &CreateJutsu{db: db}
}

func (create *CreateJutsu) Run(jutsu entities.Jutsu) (entities.Jutsu, error) {
	// Guardar el Jutsu en la base de datos
	err := create.db.SaveJutsu(jutsu.Name, jutsu.JutsuType, jutsu.Nature, jutsu.DifficultyLevel, jutsu.CreatedBy)
	if err != nil {
		return entities.Jutsu{}, err
	}

	// Retornar el Jutsu creado
	return jutsu, nil
}
