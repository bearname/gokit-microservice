package postgres

import (
	"github.com/jackc/pgx"
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
	"net"
	"time"
)

type connConfig struct {
	DatabaseURI    string `envconfig:"DB_URI"`
	MaxConnections int    `envconfig:"DB_MAX_CONN" default:"5"`
	AcquireTimeout int64  `envconfig:"DB_TIMEOUT" default:"0"` // max wait time when all connections are busy (0 means no timeout)
}

func ParseEnvConfig(prefix string) (pgx.ConnPoolConfig, error) {
	var (
		err    error
		config connConfig
	)
	if err = envconfig.Process(prefix, &config); err != nil {
		return pgx.ConnPoolConfig{}, errors.Wrap(err, "failed to parse database connection environment config values")
	}
	pgxConnConfig, err := pgx.ParseURI("postgres://postgres:postgres@localhost:5432/catalog?sslMode=disable")
	if err != nil {
		return pgx.ConnPoolConfig{}, errors.Wrap(err, "failed to parse database URI from environment variable")
	}
	pgxConnConfig.Dial = (&net.Dialer{Timeout: 10 * time.Second, KeepAlive: 5 * time.Minute}).Dial
	pgxConnConfig.RuntimeParams = map[string]string{
		"standard_conforming_strings": "on",
	}
	pgxConnConfig.PreferSimpleProtocol = true

	return pgx.ConnPoolConfig{
		ConnConfig:     pgxConnConfig,
		MaxConnections: config.MaxConnections,
		AcquireTimeout: time.Duration(config.AcquireTimeout) * time.Second,
	}, nil
}

func NewConnectionPool(config pgx.ConnPoolConfig) (*pgx.ConnPool, error) {
	return pgx.NewConnPool(config)
}
