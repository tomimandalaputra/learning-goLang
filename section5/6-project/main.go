package main

import "fmt"

type Payable interface {
	CalculatePay() float64
	String() string
}

type SalaryEmployee struct {
	Name         string
	AnnualSalary float64
}

func (se SalaryEmployee) CalculatePay() float64 {
	return se.AnnualSalary / 12.0
}

func (se SalaryEmployee) String() string {
	return fmt.Sprintf("Salaried: %s (Annual: $%.2f)", se.Name, se.AnnualSalary)
}

type HourlyEmployee struct {
	Name        string
	HourlyRate  float64
	HoursWorked float64
}

func (he HourlyEmployee) CalculatePay() float64 {
	return he.HourlyRate * he.HoursWorked
}

func (he HourlyEmployee) String() string {
	return fmt.Sprintf("Hourly: %s (Rate: $%.2f/hr, Hours: %.2f)", he.Name, he.HourlyRate, he.HoursWorked)
}

type CommissionEmployee struct {
	Name           string
	BaseSalary     float64 // monthly base salary
	CommissionRate float64 // commission rate (e.g. 0.1 for 10%)
	SalesAmount    float64 // total sales amount
}

func (ce CommissionEmployee) CalculatePay() float64 {
	return ce.BaseSalary + (ce.CommissionRate * ce.SalesAmount)
}

func (ce CommissionEmployee) String() string {
	return fmt.Sprintf("Commissioned: %s (Base: $%.2f, Rate: %.2f, Sales: $%.2f)", ce.Name, ce.BaseSalary, ce.CommissionRate, ce.SalesAmount)
}

func PrintEmployeeSummary[P fmt.Stringer](employee P) {
	fmt.Printf(" - Proccessing: %s\n", employee)
}

func ProcessPayroll(employees []Payable) {
	fmt.Println("Processing payroll...")
	totalPayroll := 0.0
	for _, employee := range employees {
		PrintEmployeeSummary(employee)
		pay := employee.CalculatePay()
		fmt.Printf("   Monthly pay: $%.2f\n", pay)
		totalPayroll += pay
	}
	fmt.Printf("\nTotal Monthly Payroll: $%.2f\n", totalPayroll)
	fmt.Println("----------------------------")
}

func main() {
	// se := SalaryEmployee{"John Doe", 50000}
	// fmt.Println(se)
	// fmt.Println(se.CalculatePay())

	// he := HourlyEmployee{"Jojo", 20.0, 40.0}
	// fmt.Println(he)
	// fmt.Println(he.CalculatePay())

	// ce := CommissionEmployee{"Donny", 1000.0, 0.1, 5000.0}
	// fmt.Println(ce)
	// fmt.Println(ce.CalculatePay())

	fmt.Println("Welcome to the payroll processor!")
	fmt.Println("----------------------------")

	ProcessPayroll([]Payable{
		SalaryEmployee{"John Doe", 72000.00},
		HourlyEmployee{"Jojo", 25.00, 160.00},
		CommissionEmployee{"Donny", 2000.00, 0.10, 15000.00},
		HourlyEmployee{"Diana", 30.00, 150.00},
	})
}
