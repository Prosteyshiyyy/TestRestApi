package main

import (
	"RestApi/DataBase"
	"context"
	"fmt"

	"github.com/k0kubun/pp"
)

func main() {
	ctx := context.Background()
	bd := DataBase.CreateConnection(ctx)
	if err := bd.CreateTable(ctx); err != nil {
		fmt.Println(err)
	}

	_, pps := bd.GetRows(ctx)
	pp.Println(pps)
}
