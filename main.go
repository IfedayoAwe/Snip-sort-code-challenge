package main

import (
	"fmt"
	"slices"
)

func main() {
	products := []*Product{
		{ID: 1, Name: "Alabaster Table", Price: 12.99, Created: "2019-01-04", SalesCount: 32, ViewsCount: 730},
		{ID: 2, Name: "Zebra Table", Price: 44.49, Created: "2012-01-04", SalesCount: 301, ViewsCount: 3279},
		{ID: 3, Name: "Coffee Table", Price: 10.00, Created: "2014-05-28", SalesCount: 1048, ViewsCount: 20123},
		{ID: 4, Name: "No Views", Price: 100.00, Created: "2020-01-01", SalesCount: 10, ViewsCount: 0},
	}

	manager := NewProductSorterManager()
	manager.RegisterSorter(PriceSorter{})
	manager.RegisterSorter(SalesViewRatioSorter{})
	manager.RegisterSorter(CreationDateSorter{})

	fmt.Println("Original Products:")
	for _, p := range products {
		fmt.Printf("%+v\n", *p)
	}

	fmt.Println("\nSorted by Price:")
	sortedByPrice := slices.Clone(products)
	manager.SortProducts(sortedByPrice, "Price")
	for _, p := range sortedByPrice {
		fmt.Printf("%+v\n", *p)
	}

	fmt.Println("\nSorted by Sales/View Ratio:")
	sortedByRatio := slices.Clone(products)
	manager.SortProducts(sortedByRatio, "SalesViewRatio")
	for _, p := range sortedByRatio {
		fmt.Printf("%+v\n", *p)
	}

	fmt.Println("\nSorted by Creation Date:")
	sortedByCreation := slices.Clone(products)
	manager.SortProducts(sortedByCreation, "CreationDate")
	for _, p := range sortedByCreation {
		fmt.Printf("%+v\n", *p)
	}
}
