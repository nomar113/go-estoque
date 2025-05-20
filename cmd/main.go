package main

import (
	"estoque/internal/models"
	"fmt"
)

func main() {
	fmt.Println("Sistema de Estoque")

	item1 := models.Item{
		ID:       1,
		Name:     "product 1",
		Quantity: 12,
		Price:    89.12,
	}

	item2 := models.Item{
		ID:       2,
		Name:     "product 2",
		Quantity: 5,
		Price:    47.79,
	}

	fmt.Println(item1.Info())
	fmt.Println(item2.Info())
}
