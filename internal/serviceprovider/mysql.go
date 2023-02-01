package serviceprovider

import (
	"osoc/internal/config"

	"osoc/pkg/mysql"
)

func NewSlaveMysql() (*mysql.SlaveMysql, func(), error) {
	db, err := mysql.Open(
		mysql.Host("osoc_node1"),
		mysql.Port(3308),
		mysql.User("root"),
		mysql.Password("root"),
		mysql.DBName("osoc"),
		mysql.ParseTime(true),
		mysql.MaxIdleConnections(25),
		mysql.VerificationRequired(false),
	)
	if err != nil {
		return nil, nil, err
	}

	closeDB := func() {
		_ = db.Close()
	}

	return &mysql.SlaveMysql{DB: db}, closeDB, nil
}
func NewMysql(conf config.Mysql) (*mysql.DB, func(), error) {
	db, err := mysql.Open(
		mysql.Host(conf.Host),
		mysql.Port(conf.Port),
		mysql.User(conf.User),
		mysql.Password(conf.Password),
		mysql.DBName(conf.DbName),
		mysql.ParseTime(conf.ParseTime),
		mysql.MaxIdleConnections(25),
		mysql.VerificationRequired(false),
	)
	if err != nil {
		return nil, nil, err
	}

	closeDB := func() {
		_ = db.Close()
	}

	return db, closeDB, nil
}
