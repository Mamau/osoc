package post

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"osoc/internal/entity"
	"osoc/pkg/log"
)

const MaxPosts = 1000

type Cache struct {
	connection *redis.Client
	log        log.Logger
}

func NewCacheRepository(rc *redis.Client, l log.Logger) *Cache {
	return &Cache{
		connection: rc,
		log:        l,
	}
}

func (c *Cache) Save(ctx context.Context, userID int, post entity.Post) error {
	data, err := json.Marshal(post)
	if err != nil {
		return err
	}
	key := storageKey(userID)

	if c.connection.LLen(ctx, key).Val() >= MaxPosts {
		c.connection.RPop(ctx, key)
	}

	if err := c.connection.LPush(ctx, key, data).Err(); err != nil {
		return err
	}

	return nil
}

func (c *Cache) GetFeeds(ctx context.Context, userID int) ([]entity.Post, error) {
	result, err := c.connection.LRange(ctx, storageKey(userID), 0, MaxPosts).Result()
	if err != nil {
		return nil, err
	}

	var posts []entity.Post
	for _, v := range result {
		var post entity.Post
		if err := json.Unmarshal([]byte(v), &post); err != nil {
			c.log.Err(err).Msg("error while get post by key")
			continue
		}

		posts = append(posts, post)
	}

	return posts, nil
}

func storageKey(userID int) string {
	return fmt.Sprintf("user:%d", userID)
}
