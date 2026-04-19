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
	    title VARCHAR(255) NOT NULL,
	    price Integer NOT NULL,
	    url VARCHAR(1000) NOT NULL,	    
	    description VARCHAR(1000) NOT NULL,
	    UNIQUE(idtar)
	    );

`
	_, err := b.connection.Exec(ctx, sqlStr)
	return err
}

func (b *Base) CreateTarget(ctx context.Context, idtar, price int, title, url, description string) error {
	sqlStr := `
	INSERT INTO targetsss (idtar,price,url,description, title)
	VALUES ($1,$2,$3,$4,$5);

`

	_, err := b.connection.Exec(ctx, sqlStr, idtar, price, url, description, title)
	return err
}

func (b *Base) ChangePrice(ctx context.Context, id, price int) error {
	sqlStr := `
	UPDATE targetsss 
	SET price = $2
	WHERE id = $1;
`
	_, err := b.connection.Exec(ctx, sqlStr, id, price)
	return err
}

func (b *Base) GetRows(ctx context.Context) (error, []TargetModel) {
	sqlStr := `
	SELECT id, idtar, price, title, url
	FROM targetsss
	WHERE price < 50000
	ORDER BY price ASC LIMIT 10;
`
	targets := make([]TargetModel, 0)
	rows, err := b.connection.Query(ctx, sqlStr)
	if err != nil {
		return err, []TargetModel{}
	}
	defer rows.Close()
	for rows.Next() {
		var target = TargetModel{}
		err := rows.Scan(
			&target.Id,
			&target.Idtar,
			&target.Price,
			&target.Title,
			&target.URl,
		)
		if err != nil {
			return err, targets
		}

		targets = append(targets, target)

	}
	return nil, targets
}
