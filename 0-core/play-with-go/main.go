package main

import (
	"play-with-go/billing"
	"play-with-go/pricematrix"
)

func main() {
	PriceMatrixV1 := pricematrix.GetPriceMatrix("v1")
	billing := billing.NewBilling(PriceMatrixV1) // Constructor Injection

	cart := []string{"item1", "item2", "item3"}
	totalPrice := billing.GetTotalPrice(cart)
	println("Total Price: $", totalPrice)

	cart = []string{"item4", "item5", "item6"}
	totalPrice = billing.GetTotalPrice(cart)
	println("Total Price: $", totalPrice)

}

// -------------------------------------------------------------
// SOLID Principles
// -------------------------------------------------------------

// S - Single Responsibility Principle (SRP)
// A class should have only one reason to change. In other words, a class should have only one job or responsibility.
// In the above code, the PriceMatrix and Billing classes have separate responsibilities. PriceMatrix is responsible for fetching prices,
// while Billing is responsible for calculating the total price. This separation of concerns makes the code easier to maintain and extend.

// O - Open for extension/Closed  for modification Principle (OCP)
// A class should be open for extension but closed for modification. This means that you should be able to add new functionality to a class
// without changing its existing code. In the above code, we can easily add new versions of PriceMatrix (e.g., PriceMatrix_v3) without modifying
// the existing code. We can simply create a new class that implements the PriceMatrix interface and use it in the Billing class.

// L - Liskov Substitution Principle (LSP)
// Subtypes must be substitutable for their base types without altering the correctness of the program.
// In the above code, we can substitute any implementation of the PriceMatrix interface (e.g., PriceMatrix_v1 or PriceMatrix_v2) in the Billing class
// without affecting the correctness of the program. This is because the Billing class only depends on the PriceMatrix interface, not on any specific implementation.

// I - Interface Segregation Principle (ISP)
// Clients should not be forced to depend on interfaces they do not use.
// In the above code, the PriceMatrix interface is simple and focused on a single responsibility (fetching prices).
// This means that clients (e.g., Billing) are not forced to depend on methods they do not use. If we had a more complex interface with multiple methods,
// clients would be forced to implement methods they do not need, which would violate the ISP.

// D - Dependency Inversion Principle (DIP)
// High-level modules should not depend on low-level modules. Both should depend on abstractions.
// Abstractions should not depend on details. Details should depend on abstractions.
// In the above code, the Billing class (high-level module) depends on the PriceMatrix interface (abstraction) instead of a specific implementation (low-level module).
// This means that we can easily swap out different implementations of PriceMatrix without affecting the Billing class.
// This makes the code more flexible and easier to maintain. Additionally, we can inject dependencies into the Billing class (e.g., using constructor injection),
// which further decouples the classes and adheres to the DIP.
// -------------------------------------------------------------
// design patterns
// -------------------------------------------------------------

// - creational patterns
// - behavioral patterns
// - structural patterns

//--------------------------------------------------------------
