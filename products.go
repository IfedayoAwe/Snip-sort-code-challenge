package main

import (
	"fmt"
	"sort"
	"time"
)

type Product struct {
	ID         int
	Name       string
	Price      float64
	Created    string
	SalesCount int
	ViewsCount int
}

type Sorter interface {
	Sort(products []*Product)
	Name() string
}

type PriceSorter struct{}

func (s PriceSorter) Sort(products []*Product) {
	sort.Slice(products, func(i, j int) bool {
		return products[i].Price > products[j].Price
	})
}

func (s PriceSorter) Name() string {
	return "Price"
}

type SalesViewRatioSorter struct{}

func (s SalesViewRatioSorter) Sort(products []*Product) {
	sort.Slice(products, func(i, j int) bool {
		if products[i].ViewsCount == 0 {
			return false // Treat as lowest ratio
		}
		if products[j].ViewsCount == 0 {
			return true // Treat as lowest ratio
		}
		ratioI := float64(products[i].SalesCount) / float64(products[i].ViewsCount)
		ratioJ := float64(products[j].SalesCount) / float64(products[j].ViewsCount)
		return ratioI > ratioJ
	})
}

func (s SalesViewRatioSorter) Name() string {
	return "SalesViewRatio"
}

type CreationDateSorter struct{}

func (s CreationDateSorter) Sort(products []*Product) {
	sort.Slice(products, func(i, j int) bool {
		timeI, errI := time.Parse("2006-01-02", products[i].Created)
		timeJ, errJ := time.Parse("2006-01-02", products[j].Created)
		if errI != nil || errJ != nil {
			return false
		}
		return timeI.After(timeJ) // Newest first
	})
}

func (s CreationDateSorter) Name() string {
	return "CreationDate"
}

type ProductSorterManager struct {
	sorters map[string]Sorter
}

func NewProductSorterManager() *ProductSorterManager {
	return &ProductSorterManager{
		sorters: make(map[string]Sorter),
	}
}

func (m *ProductSorterManager) RegisterSorter(sorter Sorter) {
	m.sorters[sorter.Name()] = sorter
}

func (m *ProductSorterManager) SortProducts(products []*Product, sorterName string) {
	sorter, ok := m.sorters[sorterName]
	if !ok {
		fmt.Println("Sorter not found:", sorterName)
		return
	}
	sorter.Sort(products)
}
