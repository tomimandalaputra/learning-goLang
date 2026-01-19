package main

import (
	"errors"
	"fmt"
)

type Person struct {
	Name  string
	Email string
	Age   int
}

type Address struct {
	Street  string
	Country string
}

// define method FullAddress untuk retrun string dari Address
func (a Address) FullAddress() string {
	return fmt.Sprintf("%s, %s", a.Street, a.Country)
}

// penerapan compostion Person dan Address
type Employee struct {
	EmployeeID int
	Person
	Address
	Position string
	Salary   int
	IsActive bool
}

// penerapan interface Stringer
func (e *Employee) String() string {
	return fmt.Sprintf("Employee ID: %d, Name: %s, Position: %s, Salary: %d, Address: %s", e.EmployeeID, e.Name, e.Position, e.Salary, e.Address.FullAddress())
}

// penerapan method GetInfo dengan return string dan error
func (e *Employee) GetInfo() (string, error) {
	if e.Name == "" {
		return "", errors.New("name is empty")
	}

	if e.IsActive == false {
		return "", errors.New("employee is not active")
	}

	info := e.String()

	return info, nil
}

func main() {
	employee := Employee{
		EmployeeID: 1,
		Person: Person{
			Name:  "John Doe",
			Email: "john.doe@example.com",
			Age:   30,
		},
		Address: Address{
			Street:  "123 Main St",
			Country: "USA",
		},
		Position: "Software Engineer",
		Salary:   5000,
		IsActive: false,
	}

	info, err := employee.GetInfo()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(info)

}
