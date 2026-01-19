package main

import "fmt"

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

type ContactInfo struct {
	Email string
	Phone string
}

func (ci ContactInfo) DisplayContact() string {
	return fmt.Sprintf("Email: %s, Phone: %s", ci.Email, ci.Phone)
}

type Company struct {
	Name string
	Address
	ContactInfo
	BusinessType string
}

func (c Company) GetProfile() {
	fmt.Println("Company Name:", c.Name)
	fmt.Println("Address:", c.Address.FullAddress())
	fmt.Println("Contact Info:", c.ContactInfo.DisplayContact())
	fmt.Println("Business Type:", c.BusinessType)
}

func main() {
	company := Company{
		Name: "Tuku Code",
		Address: Address{
			Street:  "123 Main St",
			City:    "Jakarta",
			State:   "DKI Jakarta",
			ZipCode: "12345",
		},
		ContactInfo: ContactInfo{
			Email: "example@company.com",
			Phone: "123456789",
		},
		BusinessType: "Technology",
	}

	company.GetProfile()
}
