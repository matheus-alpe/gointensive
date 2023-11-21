package main

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/matheus-alpe/gointensive/internal/infra/database"
	"github.com/matheus-alpe/gointensive/internal/usecase"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "db.sqlite3")
	if err != nil {
		panic(err)
	}

	uc := usecase.NewCalculateFinalPrice(database.NewOrderRepository(db))

	input := usecase.OrderInput{
		Id: fmt.Sprintf("%v", time.Now()),
		Price: 10.0,
		Tax: 5.0,
	}

	output, err := uc.Execute(input)
	if err != nil {
		panic(err)
	}

	fmt.Println(output)
	fmt.Println(uc.OrderRepository.GetTotalTransactions())
}
