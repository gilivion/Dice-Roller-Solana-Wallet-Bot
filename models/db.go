package models

import (
    "database/sql"
    "fmt"
)

var db *sql.DB

func InitDB(conn *sql.DB) { db = conn }

func DB() *sql.DB { return db }

func Migrate() error {
    if db == nil {
        return fmt.Errorf("db is not initialized")
    }
    stmts := []string{
        `CREATE TABLE IF NOT EXISTS wallets (
            user_id BIGINT PRIMARY KEY,
            pub_key BYTEA NOT NULL,
            priv_key BYTEA NOT NULL,
            created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
        );`,
    }
    for _, s := range stmts {
        if _, err := db.Exec(s); err != nil {
            return err
        }
    }
    return nil
}
