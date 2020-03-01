package store

import (
	"fmt"

	redis "github.com/go-redis/redis/v7"
)

type Store interface {
	CreateEntry(key string, value string) error
	RetrieveEntry(key string) (string, error)
	DeleteEntry(key string) (int, error)
}

type redisStore struct{}

func newClient() (*redis.Client, error) {
	var client *redis.Client
	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	fmt.Println("Redis health check - ")
	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	// Output: PONG <nil>
	if err != nil {
		return nil, err
	}

	return client, nil
}

// NewStore returns a new Redis Store
func NewStore() Store {
	return new(redisStore)
}

// CreateEntry adds a new entry into DB
func (r *redisStore) CreateEntry(key string, value string) error {
	client, err := newClient()
	// connection issue
	if err != nil {
		return err
	}
	err = client.Set(key, value, 0).Err()
	if err != nil {
		return err
	}

	return nil
}

// DeleteEntry deletes an exisiting Entry
func (r *redisStore) DeleteEntry(key string) (int, error) {
	client, err := newClient()
	// connection issue
	if err != nil {
		return -1, err
	}
	delCount := client.Del(key)
	return int(delCount.Val()), nil
}

// RetrieveEntry gets an Entry if exists
func (r *redisStore) RetrieveEntry(key string) (string, error) {
	client, err := newClient()
	// connection issue
	if err != nil {
		return "", err
	}
	val, err := client.Get(key).Result()
	return val, err
}
