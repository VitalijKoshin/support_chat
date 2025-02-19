package database

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"supportchat/internal/apperrors"
	"supportchat/internal/config"

	_ "github.com/go-sql-driver/mysql"
)

type MySQLDB struct {
	*sql.DB
}

func NewMySQLDB(cfg config.MySQLConfig) (*MySQLDB, error) {
	dbDSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&multiStatements=true", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Database)
	db, err := sql.Open("mysql", dbDSN)
	if err != nil {
		return nil, apperrors.SqlOpenError.AppendMessage(err)
	}
	ctx, cancelfunc := context.WithTimeout(context.Background(), time.Second)
	defer cancelfunc()
	err = db.PingContext(ctx)
	if err != nil {
		return nil, apperrors.PingDBError.AppendMessage(err)
	}

	return &MySQLDB{db}, nil
}
