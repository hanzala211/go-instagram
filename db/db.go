package db

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"time"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"github.com/hanzala211/instagram/internal/api/models"
	"github.com/hanzala211/instagram/utils"
)

func ConnectPGDB() *pg.DB {
	db := pg.Connect(&pg.Options{
		Addr:     utils.GetEnv("DB_HOST", "localhost:5432"),
		User:     utils.GetEnv("DB_USER", "postgres"),
		Password: utils.GetEnv("DB_PASS", ""),
		Database: utils.GetEnv("DB_NAME", "postgres"),
		TLSConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	})
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	if err := db.Ping(ctx); err != nil {
		log.Fatal("Failed to connect to database", err)
	}
	return db
}

func Migrations(db *pg.DB) {
	modelSlice := []interface{}{
		(*models.User)(nil),
	}

	for _, model := range modelSlice {
		err := db.Model(model).CreateTable(&orm.CreateTableOptions{
			IfNotExists:   true,
			Temp:          false,
			FKConstraints: true,
		})
		if err != nil {
			log.Fatal("Failed to create table", err)
		}
	}

	_, err := db.Exec(`
		CREATE EXTENSION IF NOT EXISTS pgcrypto;
		ALTER TABLE users
			ALTER COLUMN "id" SET DEFAULT gen_random_uuid(),
			ALTER COLUMN "createdAt" SET DEFAULT CURRENT_TIMESTAMP,
			ALTER COLUMN "updatedAt" SET DEFAULT CURRENT_TIMESTAMP;

		CREATE OR REPLACE FUNCTION update_timestamp()
		RETURNS TRIGGER AS $$
		BEGIN
			NEW."updatedAt" := CURRENT_TIMESTAMP;
			RETURN NEW;
		END;
		$$ LANGUAGE plpgsql;

		DROP TRIGGER IF EXISTS update_users_timestamp ON users;
		CREATE TRIGGER update_users_timestamp
			BEFORE UPDATE ON users
			FOR EACH ROW
			EXECUTE FUNCTION update_timestamp();
	`)

	if err != nil {
		fmt.Printf("migration SQL execution error: %v\n", err)
	}
}
