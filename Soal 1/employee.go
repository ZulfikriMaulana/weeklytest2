package main

import (
	"math/rand"
	"time"
)

type Employee struct {
	ID          int
	FullName    string
	Salary      float64
	Status      string
	Insurance   float64
	Overtime    float64
	Allowance   float64
	TotalSalary float64
}

// Fungsi untuk membuat 100 karyawan
func CreateEmployees() []Employee {
	rand.Seed(time.Now().UnixNano())
	statuses := []string{"Permanent", "Contract", "Trainee"}
	names := []string{"a1", "b1", "a2", "c4", "a5"}

	var employees []Employee
	for i := 0; i < 100; i++ {
		status := statuses[rand.Intn(len(statuses))]
		name := names[rand.Intn(len(names))]
		salary := float64(rand.Intn(10000) + 5000)
		insurance := 500000.0
		overtime := 0.0
		allowance := 0.0

		if status == "Contract" {
			overtime = 55000.0
		} else if status == "Trainee" {
			allowance = 100000.0
		}

		employee := Employee{
			ID:          101 + i,
			FullName:    name,
			Salary:      salary,
			Status:      status,
			Insurance:   insurance,
			Overtime:    overtime,
			Allowance:   allowance,
			TotalSalary: 0.0,
		}

		employees = append(employees, employee)
	}

	return employees
}

// Fungsi untuk menghitung total gaji
func CalculateTotalSalary(employee *Employee) {
	employee.TotalSalary = employee.Salary + employee.Insurance + employee.Overtime + employee.Allowance
}

func CalculateSalaries(employees []Employee) {
	for i := range employees {
		CalculateTotalSalary(&employees[i])
	}
}

func CalculateSalariesWithChannel(employees []Employee, ch chan<- bool) {
	for i := range employees {
		CalculateTotalSalary(&employees[i])
	}
	ch <- true
}
