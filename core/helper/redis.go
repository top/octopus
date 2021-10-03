package helper

import (
	"errors"
	"time"

	"github.com/go-redis/redis"
)

// Redis Redis
type Redis struct {
	rc *redis.Client
}

// NewRedis NewRedis
func NewRedis(addr, password string, db int) *Redis {
	return &Redis{
		rc: redis.NewClient(&redis.Options{
			Addr: addr,
			OnConnect: func(c *redis.Conn) error {
				_, err := c.Ping().Result()
				return err
			},
			Password: password,
			DB:       db,
		}),
	}
}

func (r *Redis) Get(key string) (string, error) {
	if val, err := r.rc.Get(key).Result(); err == nil {
		return val, nil
	}
	return "", errors.New("key not found")
}

func (r *Redis) SetExp(key string, value interface{}, expire time.Duration) error {
	if _, err := r.rc.Set(key, value, expire).Result(); err != nil {
		return err
	}
	return nil
}

// Poll Poll
func (s *Redis) Poll(ch string) string {
	if val, err := s.rc.LPop(ch).Result(); err == nil {
		return val
	}
	return ""
}

// RandomKey RandomKey
func (s *Redis) RandomKey() string {
	if val, err := s.rc.RandomKey().Result(); err == nil {
		return val
	}
	return ""
}

// Incr Incr
func (s *Redis) Incr(value string) error {
	if _, err := s.rc.Incr(value).Result(); err != nil {
		return err
	}
	return nil
}

// RPush RPush
func (s *Redis) RPush(channel string, item interface{}) error {
	if str, err := NewJson().Marshal(item); err == nil {
		if _, err := s.rc.RPush(channel, string(str)).Result(); err != nil {
			return err
		}
	}
	return nil
}

// LPush LPush
func (s *Redis) LPush(channel string, item interface{}) error {
	if str, err := NewJson().Marshal(item); err == nil {
		if _, err := s.rc.LPush(channel, string(str)).Result(); err != nil {
			return err
		}
	}
	return nil
}

// Count Count
func (s *Redis) Count(list string) int {
	n, err := s.rc.LLen(list).Result()
	if err != nil {
		return 0
	}
	return int(n)
}

// Remove Remove
func (s *Redis) Remove(item string) bool {
	if _, err := s.rc.Del(item).Result(); err != nil {
		return false
	}
	return true
}

func (s *Redis) Close() error {
	return s.rc.Close()
}
