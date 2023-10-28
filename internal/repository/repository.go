package repository

import (
	"avito/common"
	"avito/config"
	"avito/database"
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"log/slog"
	"strings"
	"sync"
)

type Repository struct {
	pool  *pgxpool.Pool
	mu    sync.RWMutex
	cache map[string]string
}

var countid uint64 = 1

const (
	alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ012345678"
)

func New(cfg config.Config) *Repository {
	pool, err := database.ConnectDB(cfg)
	if err != nil {
		slog.Error("error in create pool connects: ", err)
	}
	cache := make(map[string]string)
	return &Repository{pool: pool, cache: cache}
}

func (r *Repository) CacheRecovery() error {
	r.mu.Lock()
	defer r.mu.Unlock()
	rows, err := r.pool.Query(context.Background(), "SELECT id,sortlink, longlink FROM links")
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var id, b string
		err := rows.Scan(&countid, &id, &b)
		if err != nil {
			continue
		}
		r.cache[id] = b
	}
	if err = rows.Err(); err != nil {
		countid += 1
		return err
	}
	countid += 1
	return nil
}

func base62Encode(number uint64) string {
	length := len(alphabet)
	var encodedBuilder strings.Builder
	encodedBuilder.Grow(10)
	for ; number > 0; number = number / uint64(length) {
		encodedBuilder.WriteByte(alphabet[(number % uint64(length))])
	}

	return encodedBuilder.String()
}

func newShortLink(msg string) (string, error) {
	return base62Encode(countid), nil
}

func (r *Repository) CreateLink(msg string) (string, error) {
	shortLink, err := newShortLink(msg)
	fmt.Println("CreateLink - ", shortLink, err)
	if err != nil {
		return "", err
	}
	_, err = r.pool.Exec(context.Background(), "INSERT INTO links VALUES ($1,$2,$3)", countid, shortLink, msg)
	fmt.Println("CreateLink - exec - ", countid, shortLink, err)
	if err != nil {
		return "", err
	} else {
		countid += 1
		r.cache[shortLink] = msg
	}
	return shortLink, err
}

func (r *Repository) GetLink(shortLink string) (string, error) {
	msg, ok := r.cache[shortLink]
	fmt.Println("GetLink - ", msg, ok, r.cache[shortLink])
	fmt.Println(r.cache)
	var err error
	if ok == false {
		err = r.pool.QueryRow(context.Background(), "SELECT longlink FROM links WHERE sortlink=($1)", shortLink).Scan(&msg)
		fmt.Println("GetLink - ", msg, err)
		if errors.Is(err, pgx.ErrNoRows) {
			return "", common.ErrNotFound
		}
		if err != nil {
			return "", err
		}
		r.cache[shortLink] = msg
	}
	return msg, err
}
