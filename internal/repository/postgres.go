package repository

import (
	"en_train/internal/config"
	"en_train/pkg/helpers"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"

	_ "github.com/jackc/pgx/stdlib"
)

const (
	irregularVerbs = "irregular_verbs"
)

func NewPostgresDB(cfg config.DBConfig) (*sqlx.DB, error) {
	uri := helpers.PgURI(cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.DBName)
	logrus.Info(uri)
	db, err := sqlx.Open("pgx", uri)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
