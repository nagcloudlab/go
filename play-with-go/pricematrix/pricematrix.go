package pricematrix

type PriceMatrix interface {
	GetPrice(itemCode string) float32
}
