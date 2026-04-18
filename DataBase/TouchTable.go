package DataBase

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func TouchTable(ctx context.Context, conn *pgx.Conn) error {
	sqlStr := `
	CREATE TABLE IF NOT EXISTS targets (
	    id SERIAL PRIMARY KEY,
	    idtar INTEGER NOT NULL,
	    nametar VARCHAR(255) NOT NULL,
	    url VARCHAR(1000) NOT NULL,
	    active BOOLEAN NOT NULL,
	    UNIQUE(idtar)
	);
`
	_, err := conn.Exec(ctx, sqlStr)
	return err
}
