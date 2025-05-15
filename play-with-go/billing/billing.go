package billing

import (
	"play-with-go/pricematrix"
)

//-------------------------------------------------------------
// Billing
//-------------------------------------------------------------

/*

   design issues
   ---------------
    - tight coupling between dependent (Biling) and dependency (PriceMatrix_v1)
       => can't extent new features easily
    - unit-test is difficult
       => dev/bug fix is difficult

   performance issues
    -----------------
    - too many price matrix instances created & destroyed
        => resource waste | responsive issues


     why these issues happen?

     -> dependent object managing its own dependencies

     solution
     ---------------

     -> dont't create dependency in dependent's home , get from factory ( lookup )

     any issue with this solution?
     -----------------------------

     -> on each call to getPriceMatrix, new instance is created

     best solution:
     ----------------

     -> Don't create & lookup dependency in dependent's home , inject dependency from outside
        ( Dependency Inversion Principle )

        How to implement dependency injection in Go?

        -> Constructor Injection ( required dependency is passed in constructor )
        -> Setter Injection ( optional dependency is passed in setter )

     ---------------------------------------------------------------


*/

// team: Billing
// version: 1.0
type Billing struct {
	priceMatrix pricematrix.PriceMatrix // HAS-A relationship
}

func NewBilling(priceMatrix pricematrix.PriceMatrix) *Billing {
	return &Billing{
		priceMatrix: priceMatrix,
	}
}
func (b *Billing) GetTotalPrice(cart []string) float32 {
	// Simulate calculating total price
	var total float32 = 0.0
	//priceMatrix := PriceMatrix_v1{} // Don't create dependency in dependent's home
	//PriceMatrix := getPriceMatrix("v1") // Dont's lookup dependency in dependent's home
	for _, item := range cart {
		total += b.priceMatrix.GetPrice(item)
	}
	return total
}

//-------------------------------------------------------------
