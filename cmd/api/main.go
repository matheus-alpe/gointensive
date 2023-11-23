package main

import (
	// "encoding/json"
	"fmt"
	"net/http"

	// "net/http"

	// "github.com/go-chi/chi/v5"
	// "github.com/go-chi/chi/v5/middleware"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/matheus-alpe/gointensive/internal/entity"
)

func main() {
	port := ":8888"
	fmt.Printf("API running on: http://localhost%v\n", port)

	// r := chi.NewRouter()
	// r.Use(middleware.Logger)

	// r.Get("/order", OrderHandler)

	// http.ListenAndServe(port, r)

	e := echo.New()
	e.Use(middleware.Logger())
	e.GET("order", OrderHandler)
	e.Logger.Fatal(e.Start(port))
}

func OrderHandler(c echo.Context) error {
	order, err := entity.NewOrder("1", 10, 1)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	err = order.CalculateFinalPrice()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, order)
}

// func OrderHandler(w http.ResponseWriter, r *http.Request) {
// 	order, err := entity.NewOrder("1", 10, 1)
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 	}

// 	err = order.CalculateFinalPrice()
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 	}

// 	err = json.NewEncoder(w).Encode(order)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 	}
// }
