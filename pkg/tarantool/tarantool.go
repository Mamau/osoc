package tarantool

import (
	"fmt"
	"github.com/tarantool/go-tarantool"
)

type Connection struct {
	*tarantool.Connection
}

func NewConnection(opts ...Option) (*Connection, error) {
	o := &options{
		host:     "localhost",
		port:     3301,
		user:     "admin",
		password: "admin",
	}
	for _, opt := range opts {
		opt(o)
	}

	conn, err := tarantool.Connect(fmt.Sprintf("%s:%d", o.host, o.port), tarantool.Opts{
		User: o.user,
		Pass: o.password,
	})

	if err != nil {
		return nil, err
	}

	return &Connection{
		conn,
	}, nil
}
