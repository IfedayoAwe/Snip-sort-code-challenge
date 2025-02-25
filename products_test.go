package main

import (
	"testing"
	"time"
)

func equalSlices(a, b []float64) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestPriceSorter(t *testing.T) {
	t.Run("Sort by Price (Descending)", func(t *testing.T) {
		products := []*Product{
			{ID: 1, Price: 20.00},
			{ID: 2, Price: 10.00},
			{ID: 3, Price: 15.00},
		}

		expected := []float64{20.00, 15.00, 10.00}

		s := PriceSorter{}
		s.Sort(products)

		var sortedPrices []float64
		for _, p := range products {
			sortedPrices = append(sortedPrices, p.Price)
		}

		if !equalSlices(sortedPrices, expected) {
			t.Errorf("PriceSorter failed, expected %v, got %v", expected, sortedPrices)
		}
	})
}
func TestSalesViewRatioSorter(t *testing.T) {
	products := []*Product{
		{ID: 1, SalesCount: 100, ViewsCount: 200}, // Ratio 0.5
		{ID: 2, SalesCount: 90, ViewsCount: 100},  // Ratio 0.9
		{ID: 3, SalesCount: 50, ViewsCount: 100},  // Ratio 0.5
		{ID: 4, SalesCount: 10, ViewsCount: 0},    // Division by zero case
	}

	s := SalesViewRatioSorter{}
	s.Sort(products)

	if products[0].ID != 2 || products[1].ID != 1 && products[1].ID != 3 {
		t.Errorf("SalesViewRatioSorter failed, got %+v", products)
	}
}

func TestCreationDateSorter(t *testing.T) {
	products := []*Product{
		{ID: 1, Created: "2023-01-01"},
		{ID: 2, Created: "2021-01-01"},
		{ID: 3, Created: "2022-01-01"},
	}

	s := CreationDateSorter{}
	s.Sort(products)

	t1, _ := time.Parse("2006-01-02", products[0].Created)
	t2, _ := time.Parse("2006-01-02", products[1].Created)
	t3, _ := time.Parse("2006-01-02", products[2].Created)

	if !(t1.After(t2) && t2.After(t3)) {
		t.Errorf("CreationDateSorter failed, got %+v", products)
	}
}

func TestProductSorterManager(t *testing.T) {
	manager := NewProductSorterManager()
	manager.RegisterSorter(PriceSorter{})
	manager.RegisterSorter(SalesViewRatioSorter{})
	manager.RegisterSorter(CreationDateSorter{})

	if len(manager.sorters) != 3 {
		t.Errorf("Expected 3 sorters, got %d", len(manager.sorters))
	}
}
