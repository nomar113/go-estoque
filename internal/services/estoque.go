package services

import (
	"estoque/internal/models"
	"fmt"
	"strconv"
	"time"
)

type Estoque struct {
	items map[string]models.Item
	logs  []models.Log
}

func NewEstoque() *Estoque {
	return &Estoque{
		items: make(map[string]models.Item),
		logs:  []models.Log{},
	}
}

func (e *Estoque) AddItem(item models.Item, user string) error {
	if item.Quantity <= 0 {
		return fmt.Errorf("erro ao adicionar item: [ID:%d] possui uma quantidade inválida (zero ou negativa)", item.ID)
	}
	existingItem, exists := e.items[strconv.Itoa(item.ID)]
	if exists {
		item.Quantity += existingItem.Quantity
	}
	e.items[strconv.Itoa(item.ID)] = item
	e.logs = append(e.logs, models.Log{
		Timestamp: time.Now(),
		Action:    "Entrada de estoque",
		User:      user,
		ItemID:    item.ID,
		Quantity:  item.Quantity,
		Reason:    "Adicionando novos itens no estoque",
	})
	return nil
}

func (e *Estoque) ListItems() []models.Item {
	var itemList []models.Item
	for _, item := range e.items {
		itemList = append(itemList, item)
	}
	return itemList
}

func (e *Estoque) ViewAuditLogs() []models.Log {
	return e.logs
}

func (e *Estoque) CalculateTotalCost() float64 {
	totalCost := 0.0
	for _, item := range e.items {
		totalCost += float64(item.Quantity) * item.Price
	}
	return totalCost
}

func (e *Estoque) RemoveItem(ItemID int, quantity int, user string) error {
	existingItem, exists := e.items[strconv.Itoa(ItemID)]
	if !exists {
		return fmt.Errorf("erro ao remover item: [ID:%d] não existe", ItemID)
	}
	if quantity <= 0 {
		return fmt.Errorf("erro ao remover item: [ID:%d] quantidade informada inválida (zero ou negativa)", ItemID)
	}
	if existingItem.Quantity < quantity {
		return fmt.Errorf("erro ao remover item: [ID:%d] quantidade insuficiente no estoque", ItemID)
	}
	existingItem.Quantity -= quantity
	if existingItem.Quantity == 0 {
		delete(e.items, strconv.Itoa(ItemID))
	} else {
		e.items[strconv.Itoa(ItemID)] = existingItem
	}
	e.logs = append(e.logs, models.Log{
		Timestamp: time.Now(),
		Action:    "Saída de estoque",
		User:      user,
		ItemID:    ItemID,
		Quantity:  quantity,
		Reason:    "Removendo itens no estoque",
	})
	return nil
}

func FindBy[T any](data []T, comparator func(T) bool) ([]T, error) {
	var result []T
	for _, v := range data {
		if comparator(v) {
			result = append(result, v)
		}
	}
	if len(result) == 0 {
		return nil, fmt.Errorf("nenhum item com o foi encontrado")
	}
	return result, nil
}
