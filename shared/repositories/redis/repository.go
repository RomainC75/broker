package redis_repo

import (
	"time"
)

type RedisData struct {
	T       time.Time `json:"t"`
	Content string    `json:"content"`
}

func (repo *RedisRepo) Set(key string, value string) error {
	err := repo.client.Set(repo.ctx, key, value, repo.refreshTime).Err()
	if err != nil {
		return err
	}
	return nil
}

func (repo *RedisRepo) Get(key string) (string, error) {
	strData, err := repo.client.Get(repo.ctx, key).Result()
	// fmt.Println("RESULT ROOT: ", strData)
	if err != nil {
		return "", err
	}
	return strData, nil
}

func isRefreshNeeded(before time.Time, authorizedDuration time.Duration) bool {
	realDuration := time.Since(before)
	if realDuration > authorizedDuration {
		return false
	}
	return true
}
