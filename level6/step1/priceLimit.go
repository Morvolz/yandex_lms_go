package main

const (
	minPrice = 99.0
	maxPrice = 20000.0
)

func ApplyPriceLimits(price float64) float64 {
	return min(max(price, minPrice), maxPrice)
}

// func main() {

// }
