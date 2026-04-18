package main

import (
	"RestApi/DataBase"
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	connection := DataBase.DBConnection(ctx)
	if err := DataBase.TouchTable(ctx, connection); err != nil {
		fmt.Println(err)
	}
	fmt.Println("success")

}
