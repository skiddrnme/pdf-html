package utils

import "pdf-golang/repository"

func Mul(price float64, quantity int) float64 {
	return price * float64(quantity)
}

func SumItems(items []repository.Item) float64 {
	var total float64
	for _, item := range items {
		total += item.Price * float64(item.Quantity)
	}
	return total
}