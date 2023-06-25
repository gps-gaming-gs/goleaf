Name: ${APP_NAME}
ListenOn: ${APP_HOST}

Log:
  Mode: ${LOG_MODE}

Telemetry:
  Name: ${APP_NAME}
  Endpoint: ${TRACE_ENDPOINT}
  Sampler: 1.0
  Batcher: jaeger

Postgres:
  Host: ${DB_HOST}
  Port: ${DB_PORT}
  SlavePort: ${DB_SLAVE_PORT}
  UserName: ${DB_USERNAME}
  Password: ${DB_PASSWORD}
  DBName: ${DB_DATABASE}
  DBPoolSize: ${DB_POOL_SIZE}
  DBTimezone: ${DB_TIMEZONE}
  DBConnMaxLifetime: ${DB_CONN_MAX_LIFETIME}
  DBDebugLevel: ${DB_DEBUG_LEVEL}

RedisCache:
  RedisSentinelNode: ${REDIS_SENTINEL_NODE}
  RedisMasterName: ${REDIS_MASTER_NAME}
  RedisDB: ${REDIS_DB}

Consul:
  Host: ${CONSUL_HOST}
  Key: ${APP_NAME}
  Meta:
    Protocol: grpc
  Tag: ${APP_TAG}
  ttl: ${CONSUL_TTL}