package pricematrix

type PriceMatrix_v2 struct{}

func (p *PriceMatrix_v2) GetPrice(itemCode string) float32 {
	// Simulate fetching price from a api
	return 200.0
}
