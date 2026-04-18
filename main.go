package main

import (
	"RestApi/DataBase"
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	bd := DataBase.CreateConnection(ctx)
	if err := bd.CreateTable(ctx); err != nil {
		fmt.Println(err)
	}

	if err := bd.CreateTarget(ctx, 2111, "Iphone12", "http://iphone12x", "lolz"); err != nil {
	}
	fmt.Println("Success")
}
