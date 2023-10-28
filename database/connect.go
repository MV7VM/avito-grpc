package database

import (
	"avito/config"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"os"
)

func ConnectDB(cfg config.Config) (pool *pgxpool.Pool, err error) {
	//var cfg config.Config
	// urlExample := "postgres://username:password@localhost:5432/database_name"
	DB_url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", cfg.DB.Username, cfg.DB.Password, cfg.DB.Host, cfg.DB.Port, cfg.DB.Name)
	pool, err = pgxpool.New(context.Background(), DB_url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		//os.Exit(1)
		return nil, err
	}
	//defer conn.Close(context.Background())
	return pool, nil
}
