package cache

import (
	"CRUD-SQL/model"
	"encoding/json"

	"github.com/go-redis/redis"
)

type RedisSessionStore struct {
	client *redis.Client
}

func NewRedisSessionStore(addr string) *RedisSessionStore {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return &RedisSessionStore{client: client}
}

func (s *RedisSessionStore) SetSession(token string, user *model.User) error {
	err := s.client.Set(token, user, 0).Err()
	return err
}

func (s *RedisSessionStore) GetSession(token string) (*model.User, error) {
	val, err := s.client.Get(token).Result()
	if err != nil {
		return nil, err
	}
	user := &model.User{}
	err = json.Unmarshal([]byte(val), user)
	return user, err
}

func (s *RedisSessionStore) DeleteSession(token string) error {
	err := s.client.Del(token).Err()
	return err
}
