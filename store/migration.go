package store

import (
	"en_train/config"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/github"
	"github.com/sirupsen/logrus"

)

func Migrate(cfg config.Config) {
	m, err := migrate.New("file://"+cfg.MigrationScriptName,cfg.DbDsn)
	if err != nil {
		fmt.Errorf("")
	}
	cl := customLog{}
	m.Log = cl
}

func (c customLog) Printf(format string, v ...interface{}) {
	logrus.Printf(format,v)
}

func (c customLog) Verbose() bool{
	return true
}

type customLog struct {}


