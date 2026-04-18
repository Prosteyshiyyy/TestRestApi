package DataBase

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func CreateConnection(ctx context.Context) Base {
	conection, err := pgx.Connect(ctx, "postgres://postgres:200804@localhost:5432/postgres")
	if err != nil {
		panic(err)
	}
	return Base{
		connection: conection,
	}
}

type Base struct {
	connection *pgx.Conn
}

func (b *Base) GetConnection() *pgx.Conn {
	return b.connection
}

func (b *Base) CreateTable(ctx context.Context) error {
	sqlStr := `
	CREATE TABLE IF NOT EXISTS targetsss (
	    id serial PRIMARY KEY,
	    idtar INTEGER NOT NULL,
	    url VARCHAR(1000) NOT NULL,
	    title VARCHAR(255) NOT NULL,
	    description VARCHAR(1000) NOT NULL,
	    
	    UNIQUE(idtar)
	    );

`
	_, err := b.connection.Exec(ctx, sqlStr)
	return err
}

func (b *Base) CreateTarget(ctx context.Context, idtar int, title, url, description string) error {
	sqlStr := `
	INSERT INTO targetsss (idtar,title,url,description)
	VALUES ($1,$2,$3,$4);

`

	_, err := b.connection.Exec(ctx, sqlStr, idtar, title, url, description)
	return err
}

func (b *Base) ChangeStatus(ctx context.Context, idtar int, title, url, description string) error {}
