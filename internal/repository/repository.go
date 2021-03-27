package repository

import (
	"en_train/internal/domain"
	"github.com/jmoiron/sqlx"
)

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		IrregularVerbs: NewIrregularVerbs(db),
	}
}

type Repository struct {
	IrregularVerbs
}

type IrregularVerbs interface{
	GetAll() ([]domain.IrregularVerbs, error)
	GetIrregularVerbById(id int64) (domain.IrregularVerbs, error)
	GetRandomVerb() (domain.IrregularVerbs, error)
}
