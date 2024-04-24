package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Employee struct {
	ID               int    `json:"id"`
	EmployeeName     string `json:"employee name"`
	EmployeeTitle    string `json:"employee title"`
	EmployeeSalary   int    `json:"employee salary"`
	EmployeeVehicle  string `json:"employee vehicle"`
	EmployeePhoneIMEI string `json:"employee phone imei"`
	EmployeeCreditCard string `json:"employee credit card"`
	EmployeeGender   string `json:"employee gender"`
}

var employees = []Employee{
	{ID: 1, EmployeeName: "John Doe", EmployeeTitle: "Software Engineer", EmployeeSalary: 100000, EmployeeVehicle: "Car", EmployeePhoneIMEI: "123456789", EmployeeCreditCard: "1234-5678-9012-3456", EmployeeGender: "Male"},
	{ID: 2, EmployeeName: "Jane Smith", EmployeeTitle: "HR Manager", EmployeeSalary: 80000, EmployeeVehicle: "None", EmployeePhoneIMEI: "987654321", EmployeeCreditCard: "9876-5432-1098-7654", EmployeeGender: "Female"},
}

func getEmployees(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, employees)
}

func createEmployee(c *gin.Context) {
	var newEmployee Employee
	if err := c.BindJSON(&newEmployee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	employees = append(employees, newEmployee)
	c.JSON(http.StatusCreated, newEmployee)
}

func getEmployeeByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	for _, emp := range employees {
		if emp.ID == id {
			c.JSON(http.StatusOK, emp)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
}

func handleNotImplemented(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"message": "Not Implemented"})
}

func main() {
	router := gin.Default()

	router.GET("/employees", getEmployees)
	// router.GET("/employee_details", getEmployees)
	router.POST("/employees", createEmployee)
	router.GET("/employees/:id", getEmployeeByID)

	router.NoRoute(handleNotImplemented)

	router.Run(":8080")
}
