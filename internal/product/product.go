package product

type Product struct {
    ID     int
    Name   string
    Cost   float64
    Price  float64
    Seller employee.Employee
}