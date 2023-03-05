package posts

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
	"osoc/internal/config"
	"osoc/internal/entity"
	"osoc/internal/usecase/userinfo"
	"osoc/pkg/broker/rabbit/consumer"
	"osoc/pkg/log"
	"time"
)

const PopularUserID = 1

type Consumer struct {
	logger   log.Logger
	userRepo userinfo.UserRepo
	done     chan struct{}
	rc       consumer.RMQConsumer
	conf     config.Rabbit
}

func NewConsumer(l log.Logger, rc consumer.RMQConsumer, conf config.Rabbit, ur userinfo.UserRepo) *Consumer {
	return &Consumer{
		logger:   l,
		rc:       rc,
		conf:     conf,
		userRepo: ur,
	}
}

func (c *Consumer) Run() {
	defer close(c.done)
	c.rc.MsgHandler = func(queue string, msg amqp.Delivery, err error) {
		if err != nil {
			c.logger.Err(err).Msg("error from rabbit")
			return
		}
		var p entity.Post
		if err := json.Unmarshal(msg.Body, &p); err != nil {
			c.logger.Err(err).Msg("error while unmarshal post")
			return
		}

		// считаем, что друзья подписаны на посты
		c.sendNotificationToFriend(p.UserID)
	}

	c.rc.Consume(fmt.Sprintf("%s.%d", c.conf.PostChannel, PopularUserID))
}

func (c *Consumer) Terminate(ctx context.Context) error {
	c.logger.Info().Msg("terminating observe post consumer")

	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-c.done:
		return nil
	}
}

func (c *Consumer) sendNotificationToFriend(userID int) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	friends, err := c.userRepo.GetFriends(ctx, userID)
	if err != nil {
		c.logger.Err(err).Msg("error while get friends in consumer posts")
		return
	}

	for _, v := range friends {
		c.logger.Info().Msgf("пользователь #%d создал пост и друг %s %s получил уведомление", userID, v.FirstName, v.LastName)
	}
}
