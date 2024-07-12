package main

import (
	"fmt"
	"time"
)

func main() {
	employees := CreateEmployees()

	// Benchmark tanpa channel
	start := time.Now()
	CalculateSalaries(employees)
	elapsed := time.Since(start)
	fmt.Printf("Time taken without channels: %s\n", elapsed)

	// Benchmark dengan channel
	ch := make(chan bool)
	start = time.Now()
	go CalculateSalariesWithChannel(employees, ch)
	<-ch
	elapsed = time.Since(start)
	fmt.Printf("Time taken with channels: %s\n", elapsed)

	// Print data sampel karyawan
	for i, employee := range employees {
		if i < 5 {
			fmt.Printf("ID: %d, Name: %s, Total Salary: %.2f\n", employee.ID, employee.FullName, employee.TotalSalary)
		}
	}
}
