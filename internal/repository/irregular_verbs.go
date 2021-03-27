package repository

import (
	"en_train/internal/domain"
	"github.com/jmoiron/sqlx"
)

type IrregularVerbsPostgres struct {
	db *sqlx.DB
}

func NewIrregularVerbs(db *sqlx.DB) *IrregularVerbsPostgres {
	return &IrregularVerbsPostgres{db: db}
}

func (i *IrregularVerbsPostgres) GetAll() ([]domain.IrregularVerbs, error) {
	iv := make([]domain.IrregularVerbs, 0, 200)

	query := "select * from irregular_verbs"
	err := i.db.Select(&iv, query)
	if err != nil {
		return nil, err
	}
	return iv, nil
}

func (i *IrregularVerbsPostgres) GetIrregularVerbById(id int64) (domain.IrregularVerbs, error) {
	iv := domain.IrregularVerbs{}
	query := "select * from irregular_verbs where id = $1"
	err := i.db.Get(&iv, query, id)
	if err != nil {
		return iv, err
	}
	return iv, err
}

func (i *IrregularVerbsPostgres) GetRandomVerb() (domain.IrregularVerbs, error) {
	iv := domain.IrregularVerbs{}
	query := "select * from irregular_verbs iv where iv.id = " +
		"(select rand from floor(random()*(select count(id) from irregular_verbs)) as rand)"
	err := i.db.Get(&iv,query)
	if err != nil {
		return iv, err
	}
	return iv, nil
}
