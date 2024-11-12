package config

type Config struct {
	PostgresQL PostgresQLConfig
	Redis      RedisConfig
	Server     ServerConfig
}

type ServerConfig struct {
	Port int    `env:"SERVER_PORT" env-default:"8080"`
	Mode string `env:"SERVER_MODE" env-default:"debug"`
}

type PostgresQLConfig struct {
	Host            string `env:"POSTGRESQL_HOST" env-default:"localhost"`
	Port            int    `env:"POSTGRESQL_PORT" env-default:"5432"`
	User            string `env:"POSTGRESQL_USER" env-required:"true"`
	Password        string `env:"POSTGRESQL_PASSWORD" env-required:"true"`
	DBName          string `env:"POSTGRESQL_DBNAME" env-required:"true"`
	Charset         string `env:"POSTGRESQL_CHARSET" env-default:"utf8mb4"`
	ParseTime       bool   `env:"POSTGRESQL_PARSETIME" env-default:"true"`
	Loc             string `env:"POSTGRESQL_LOC" env-default:"Local"`
	MaxIdleConns    int    `env:"POSTGRESQL_MAX_IDLE_CONNS" env-default:"10"`
	MaxOpenConns    int    `env:"POSTGRESQL_MAX_OPEN_CONNS" env-default:"100"`
	ConnMaxLifetime int    `env:"POSTGRESQL_CONN_MAX_LIFETIME" env-default:"3600"`
}

type RedisConfig struct {
	Host     string `env:"REDIS_HOST" env-default:"localhost"`
	Port     int    `env:"REDIS_PORT" env-default:"6379"`
	Username string `env:"REDIS_USERNAME" env-default:"redis_user"`
	Password string `env:"REDIS_PASSWORD" env-required:"true"`
	DB       int    `env:"REDIS_DB" env-default:"0"`
	PoolSize int    `env:"REDIS_POOL_SIZE" env-default:"10"`
}
