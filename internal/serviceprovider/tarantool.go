package serviceprovider

import (
	"osoc/internal/config"
	"osoc/pkg/tarantool"
)

func NewTarantool(conf config.Tarantool) (*tarantool.Connection, func(), error) {
	db, err := tarantool.NewConnection(
		tarantool.User(conf.User),
		tarantool.Password(conf.Password),
		tarantool.Host(conf.Host),
		tarantool.Port(conf.Port),
	)

	if err != nil {
		return nil, nil, err
	}

	closeDB := func() {
		_ = db.Close()
	}

	return db, closeDB, nil
}
