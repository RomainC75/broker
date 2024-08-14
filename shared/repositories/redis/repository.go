package redis_repo

func (repo *RedisRepo) Set(key string, value string) error {
	err := repo.client.Set(repo.ctx, key, value, 0).Err()
	if err != nil {
		return err
	}
	return nil
}

func (repo *RedisRepo) Get(key string) (string, error) {
	value, err := repo.client.Get(repo.ctx, key).Result()
	if err != nil {
		return "", err
	}
	return value, nil
}
