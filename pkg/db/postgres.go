// Package db содержит инфраструктурные функции для работы с PostgreSQL.
package db

import (
	"context"
	"fmt"
	"log"

	"github.com/AIdenTale/go-hack.git/internal/app/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Postgres инкапсулирует пул соединений к PostgreSQL.
type Postgres struct {
	Pool *pgxpool.Pool
}

// New создает новый пул соединений к PostgreSQL по переданной конфигурации.
func New(globalcfg *config.Config) (*Postgres, error) {
	cfg := globalcfg.Postgres

	log.Printf("%+v", cfg)

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBName, cfg.SSLMode)
	pool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		return nil, err
	}
	return &Postgres{Pool: pool}, nil
}

// Close закрывает пул соединений PostgreSQL.
func (p *Postgres) Close() {
	p.Pool.Close()
}
