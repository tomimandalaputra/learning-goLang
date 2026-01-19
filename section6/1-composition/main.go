package main

import "fmt"

/*
Composition adalah cara menggabungkan beberapa struct kecil
menjadi satu struct yang lebih besar tanpa memakai inheritance.
*/

type Address struct {
	Street  string
	City    string
	State   string
	ZipCode string
}

func (a Address) FullAddress() string {
	if a.Street == "" && a.City == "" {
		return "No adress provided"
	}

	return fmt.Sprintf("%s, %s, %s, %s", a.Street, a.City, a.State, a.ZipCode)
}

type Customer struct {
	CustomerID      int
	Name            string
	Email           string
	BillingAddress  Address
	ShippingAddress Address
}

func (c Customer) PrintDetails() {
	fmt.Println("Customer ID:", c.CustomerID)
	fmt.Println("Name:", c.Name)
	fmt.Println("Email:", c.Email)
	fmt.Println("Billing Address:", c.BillingAddress.FullAddress())
	fmt.Println("Shipping Address:", c.ShippingAddress.FullAddress())
}

func main() {
	customer := Customer{
		CustomerID: 1,
		Name:       "John Doe",
		Email:      "example@mail.com",
		BillingAddress: Address{
			Street:  "123 Main St",
			City:    "New York",
			State:   "NY",
			ZipCode: "10001",
		},
		ShippingAddress: Address{
			Street:  "456 Elm St",
			City:    "Los Angeles",
			State:   "CA",
			ZipCode: "90001",
		},
	}

	customer.PrintDetails()

	fmt.Println("------ customer with same address ------")
	addr := Address{
		Street:  "123 Main St",
		City:    "New York",
		State:   "NY",
		ZipCode: "10001",
	}

	customer2 := Customer{
		CustomerID:      2,
		Name:            "Jane Doe",
		Email:           "janedoe@mail.com",
		BillingAddress:  addr,
		ShippingAddress: addr,
	}

	customer2.PrintDetails()
}
