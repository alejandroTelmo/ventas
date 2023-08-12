// cmd/main.go
package main

import (
    "fmt"

    "github.com/ventas/internal/client"
    "github.com/ventas/internal/employee"
    "github.com/ventas/internal/product"
)

func main() {
    // Crear una instancia de Person
    person := person.Person{
        FirstName: "John",
        LastName:  "Doe",
        Address:   "123 Main St",
    }

    // Crear instancias de Client y Employee
    client := client.Client{Person: person}
    employee := employee.Employee{Person: person}

    // Crear un producto y asignar un empleado vendedor
    product := product.Product{
        ID:     1,
        Name:   "Product A",
        Cost:   10.0,
        Price:  20.0,
        Seller: employee,
    }

    // Imprimir información
    fmt.Println("Client:", client)
    fmt.Println("Employee:", employee)
    fmt.Println("Product:", product)
}
