package cache

import (
	"CRUD-SQL/model"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

// SessionStore defines the interface for session management operations
type SessionStore interface {
	SetSession(token string, user *model.UserRegister) error
	GetSession(token string) (*model.UserRegister, error)
	DeleteSession(token string) error
}

// RedisSessionStore implements SessionStore using Redis for session storage
type RedisSessionStore struct {
	client *redis.Client
	prefix string
	ttl    time.Duration
}

// NewRedisSessionStore creates a new RedisSessionStore instance
func NewRedisSessionStore(addr, password string, db int, prefix string, ttl time.Duration) (*RedisSessionStore, error) {
	client := redis.NewClient(&redis.Options{
		Addr:        addr,
		Password:    password,
		DB:          db,
		DialTimeout: 10 * time.Second, // Set a reasonable dial timeout
	})

	err := client.Ping(context.Background()).Err()
	if err != nil {
		return nil, fmt.Errorf("error connecting to Redis: %w", err)
	}

	log.Println("Connected to Redis:", addr)

	return &RedisSessionStore{
		client: client,
		prefix: prefix,
		ttl:    ttl,
	}, nil
}

func (s *RedisSessionStore) SetSession(token string, user *model.UserRegister) error {
	userJson, err := json.Marshal(user)
	if err != nil {
		log.Println("Error marshaling user data:", err)
		return fmt.Errorf("error marshaling user data: %w", err)
	}

	key := s.generateKey(token)
	log.Println("Setting session key:", key)

	err = s.client.Set(context.Background(), key, userJson, s.ttl).Err()
	if err != nil {
		log.Println("Error setting session:", err)
		return fmt.Errorf("error setting session: %w", err)
	}

	log.Println("Session set successfully")
	return nil
}

func (s *RedisSessionStore) GetSession(token string) (*model.UserRegister, error) {
	key := s.generateKey(token)
	// log.Println("Getting session key:", key)

	val, err := s.client.Get(context.Background(), key).Result()

	if err != nil {
		log.Println("Error getting session:", err)
		if errors.Is(err, redis.Nil) {
			log.Println("Session key not found")
			return nil, nil
		}
		return nil, fmt.Errorf("error getting session: %w", err)
	}

	// log.Println("Session value:", val)

	var user model.UserRegister
	err = json.Unmarshal([]byte(val), &user)
	if err != nil {
		log.Println("Error unmarshaling user data:", err)
		return nil, fmt.Errorf("error unmarshaling user data: %w", err)
	}

	log.Println("Session retrieved successfully")
	return &user, nil
}

// DeleteSession removes the session data associated with the token
func (s *RedisSessionStore) DeleteSession(token string) error {
	key := s.generateKey(token)

	err := s.client.Del(context.Background(), key).Err()
	if err != nil && !errors.Is(err, redis.Nil) { // Handle missing key gracefully
		return fmt.Errorf("error deleting session: %w", err)
	}

	return nil
}

// generateKey creates a session key with the configured prefix
func (s *RedisSessionStore) generateKey(token string) string {
	return fmt.Sprintf("%s:%s", s.prefix, token)
}
