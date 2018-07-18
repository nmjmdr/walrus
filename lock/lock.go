package lock

import (
  "github.com/go-redis/redis"
  "github.com/satori/go.uuid"
  "time"
  "errors"
)

type lockExp struct {
  Lock func(id string, expiry time.Duration) (bool, error)
  Unlock func(id string) error
  IsLocked func(id string) (bool, error)
}

func NewLockExp(client *redis.Client) *lockExp {
  l := &lockExp{}
  randomValue := uuid.NewV4().String()
  l.Lock = func (id string, expiry time.Duration) (bool, error) {
    isSet, err := client.SetNX(id, randomValue, expiry).Result()
    return isSet, err
  }
  l.Unlock = func (id string) error {
    val, err := client.Get(id).Result()
    if err != nil && err != redis.Nil {
      return err
    }
    // check if it is this instance of LockExp that owns the lock
    if val ==  randomValue {
      return client.Del(id).Err()
    } else {
      return errors.New("Cannot unlock, not the owner of lock")
    }
  }
  l.IsLocked = func (id string) (bool, error) {
    _, err := client.Get(id).Result()
    if err == redis.Nil {
      return false, nil
    }
    return true, err

  }
  return l
}
