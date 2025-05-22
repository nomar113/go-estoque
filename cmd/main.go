package main

import (
	"estoque/internal/models"
	"estoque/internal/services"
	"fmt"
)

func main() {
	fmt.Println("Sistema de Estoque")
	estoque := services.NewEstoque()
	itens := []models.Item{
		{ID: 1, Name: "Fone", Quantity: 10, Price: 100},
		{ID: 2, Name: "Camiseta", Quantity: 2, Price: 55.99},
		{ID: 3, Name: "Mouse", Quantity: 1, Price: 12.99},
	}
	for _, item := range itens {
		err := estoque.AddItem(item, "Ramon")
		if err != nil {
			fmt.Println(err)
			continue
		}
	}
	for _, item := range estoque.ListItems() {
		fmt.Printf("ID: %d | Item: %s | Quantidade: %d | Preço: %.2f\n", item.ID, item.Name, item.Quantity, item.Price)
	}
	fmt.Println()
	for _, log := range estoque.ViewAuditLogs() {
		fmt.Printf("[%s] Ação: %s - Usuário: %s - Item ID: %d - Quantidade: %d - Motivo: %s\n", log.Timestamp.Format("02/01 15:04:05"), log.Action, log.User, log.ItemID, log.Quantity, log.Reason)
	}
	fmt.Println()
	fmt.Println("Valor total do estoque: R$ ", estoque.CalculateTotalCost())
	err := estoque.RemoveItem(2, 2, "Ramon")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println()
	for _, item := range estoque.ListItems() {
		fmt.Printf("ID: %d | Item: %s | Quantidade: %d | Preço: %.2f\n", item.ID, item.Name, item.Quantity, item.Price)
	}
	itemParaBuscar, err := services.FindBy(itens, func(item models.Item) bool {
		return item.Name == "Fone"
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Item encontrado: ", itemParaBuscar)
}
