package svc

import (
    {{.imports}}
    {{.configImport}}
)

type ServiceContext struct {
	Config config.Config
    RedisClient *redis.Client
    BoDB        *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	// Redis
	redisClient := redis.NewFailoverClient(&redis.FailoverOptions{
		MasterName:    c.RedisCache.RedisMasterName,
		SentinelAddrs: strings.Split(c.RedisCache.RedisSentinelNode, ";"),
		DB:            c.RedisCache.RedisDB,
	})

	// DB
	db, err := postgrez.New(c.Postgres.Host, fmt.Sprintf("%d", c.Postgres.Port), c.Postgres.UserName,
		c.Postgres.Password, c.Postgres.DBName).
		SetTimeZone(c.Postgres.DBTimezone).
		SetLogger(logrusz.New().SetLevel(c.Postgres.DBDebugLevel).Writer()).
		Connect(postgrez.Pool(c.Postgres.DBPoolSize, c.Postgres.DBPoolSize, c.Postgres.DBConnMaxLifetime))

	if err != nil {
		panic(err)
	}

	return &ServiceContext{
		Config:c,
        RedisClient: redisClient,
        BoDB: db,
	}
}
