package main

import (
	"fmt"

	"github.com/matheus-alpe/gointensive/internal/entity"
)

func main() {
    order := entity.Order{
        Id: "123",
        Price: 59.95,
        Tax: 5.0,
    }

    order.CalculateFinalPrice()
    fmt.Println(order.FinalPrice)
}
