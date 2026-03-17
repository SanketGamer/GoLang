

func checkRedis(ctx context.Context) error {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	defer rdb.Close()
	return rdb.Ping(ctx).Err()
}