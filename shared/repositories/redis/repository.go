package redis_repo

import (
	"encoding/json"
	"time"
)

type RedisData struct {
	T       time.Time `json:"t"`
	Content string    `json:"content"`
}

func (repo *RedisRepo) Set(key string, content string) error {
	redisData := RedisData{
		T:       time.Now(),
		Content: content,
	}
	b, err := json.Marshal(redisData)
	if err != nil {
		return err
	}
	err = repo.client.Set(repo.ctx, key, b, 0).Err()
	if err != nil {
		return err
	}
	return nil
}

func (repo *RedisRepo) Get(key string) (string, error) {
	strData, err := repo.client.Get(repo.ctx, key).Result()
	if err != nil {
		return "", err
	}
	var redisData RedisData
	err = json.Unmarshal([]byte(strData), &redisData)
	if err != nil {
		return "", err
	}

	if isRefreshNeeded(redisData.T, repo.refreshTime) {
		_, err = repo.client.Del(repo.ctx, key).Result()
		if err != nil {
			return "", err
		}
		return "", nil
	}
	return redisData.Content, nil
}

func isRefreshNeeded(before time.Time, authorizedDuration time.Duration) bool {
	realDuration := time.Since(before)
	if realDuration > authorizedDuration {
		return false
	}
	return true
}
