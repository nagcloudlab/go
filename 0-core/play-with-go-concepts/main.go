package main

import "sort"

type Apple struct {
	Color  string
	Weight int
}

type Product struct {
	Name  string
	Price float64
}

func main() {

	appleInventory := []Apple{
		{Color: "red", Weight: 100},
		{Color: "green", Weight: 200},
		{Color: "yellow", Weight: 150},
		{Color: "red", Weight: 120},
		{Color: "green", Weight: 180},
	}

	// productInventory := []Product{
	// 	{Name: "Laptop", Price: 999.99},
	// 	{Name: "Smartphone", Price: 499.99},
	// 	{Name: "Tablet", Price: 299.99},
	// 	{Name: "Smartwatch", Price: 199.99},
	// 	{Name: "Headphones", Price: 149.99},
	// }

	// Filtering apples based on color=green
	// greenApples := Filter(appleInventory, func(a Apple) bool {
	// 	return a.Color == "green"
	// })
	// displaying filtered apples
	// for _, apple := range greenApples {
	// 	println("Filtered Apple - Color:", apple.Color, "Weight:", apple.Weight)
	// }
	// Filtering products based on price < 500
	// cheapProducts := Filter(productInventory, func(p Product) bool {
	// 	return p.Price < 500
	// })

	// displaying filtered products
	// for _, product := range cheapProducts {
	// 	println("Filtered Product - Name:", product.Name, "Price:", product.Price)
	// }

	// Sorting apples based on weight

	// appleInventory = Sort(appleInventory, func(a, b Apple) bool {
	// 	return a.Weight < b.Weight
	// })
	sort.Slice(appleInventory, func(i, j int) bool {
		return appleInventory[i].Weight < appleInventory[j].Weight
	})

	// displaying sorted apples
	for _, apple := range appleInventory {
		println("Sorted Apple - Color:", apple.Color, "Weight:", apple.Weight)
	}
}

// Implementation of a generic function to sort items based on a comparator
func Sort[T any](items []T, comparator func(T, T) bool) []T {
	for i := 0; i < len(items)-1; i++ {
		for j := 0; j < len(items)-i-1; j++ {
			if comparator(items[j+1], items[j]) {
				items[j], items[j+1] = items[j+1], items[j]
			}
		}
	}
	return items
}

// Implementation of a generic function to filter items based on a predicate
func Filter[T any](items []T, predicate func(T) bool) []T {
	var result []T
	for _, item := range items {
		if predicate(item) {
			result = append(result, item)
		}
	}
	return result
}
