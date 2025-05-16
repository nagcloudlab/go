package pricematrix

// Factory Method
func GetPriceMatrix(version string) PriceMatrix {
	if version == "v1" {
		return &PriceMatrix_v1{}
	} else if version == "v2" {
		return &PriceMatrix_v2{}
	}
	return nil
}
