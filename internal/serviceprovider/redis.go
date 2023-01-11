package serviceprovider

import (
	"context"
	"time"

	"osoc/internal/config"

	r "github.com/go-redis/redis/v8"
	"osoc/pkg/log"
	"osoc/pkg/redis"
)

func NewRedis(conf config.Redis, logger log.Logger) (*r.Client, func(), error) {
	rc := redis.New(
		redis.Username(conf.Username),
		redis.Password(conf.Password),
		redis.DB(conf.DB),
		redis.Host(conf.Host),
		redis.Port(conf.Port),
		redis.Logger(logger),
	)

	cleanup := func() {
		if err := rc.Close(); err != nil {
			logger.Info().Msg(err.Error())
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	statusCmd := rc.Ping(ctx)
	if err := statusCmd.Err(); err != nil {
		return rc, cleanup, err
	}

	return rc, cleanup, nil
}
