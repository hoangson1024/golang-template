package infra

import (
	"github.com/go-pg/pg"
	"github.com/sirupsen/logrus"
)

func ProvidePostgreSQL(cfg *AppConfig) (db *pg.DB, cleanup func(), err error) {
	db = pg.Connect(&pg.Options{
		Addr:     cfg.Postgres.Address,
		Database: cfg.Postgres.Database,
		User:     cfg.Postgres.Username,
		Password: cfg.Postgres.Password,
	})

	if _, err = db.ExecOne("SELECT 1"); err != nil {
		logrus.WithError(err).Errorf("Cannot ping to db instance...")
		db.Close()
		return nil, nil, err
	}

	return db, func() {
		logrus.Info("Cleanup db...")
		db.Close()
	}, nil
}

type DBSlave struct {
	DB *pg.DB
}

func ProvidePostgreSQLSlave(cfg *AppConfig) (*DBSlave, func(), error) {
	db := pg.Connect(&pg.Options{
		Addr:     cfg.Postgres.SlaveAddress,
		Database: cfg.Postgres.Database,
		User:     cfg.Postgres.Username,
		Password: cfg.Postgres.Password,
	})

	if _, err := db.ExecOne("SELECT 1"); err != nil {
		logrus.WithError(err).Errorf("Cannot ping to slave instance...")
		db.Close()
		return nil, nil, err
	}

	return &DBSlave{DB: db}, func() {
		db.Close()
	}, nil
}
