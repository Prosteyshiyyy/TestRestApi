package DataBase

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func DBConnection(ctx context.Context) *pgx.Conn {
	conection, err := pgx.Connect(ctx, "postgres://postgres:200804@localhost:5432/postgres")
	if err != nil {
		panic(err)
	}
	return conection

}
