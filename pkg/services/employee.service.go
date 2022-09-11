package services

import (
	"github.com/vimalkuriensam/go-employees-test/pkg/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type EmployeeService interface {
	AddEmployee(models.Employee) (*mongo.InsertOneResult, error)
	GetEmployee(string) *mongo.SingleResult
	UpdateEmployee(string, models.Employee) (*mongo.UpdateResult, error)
	DeleteEmployee(string) (*mongo.DeleteResult, error)
}

type employeeService struct {
	employees []models.Employee
}

func New() EmployeeService {
	return &employeeService{
		employees: []models.Employee{},
	}
}

func (e *employeeService) AddEmployee(employee models.Employee) (*mongo.InsertOneResult, error) {
	return nil, nil
}

func (e *employeeService) GetEmployee(id string) *mongo.SingleResult {
	return nil
}

func (e *employeeService) UpdateEmployee(id string, employee models.Employee) (*mongo.UpdateResult, error) {
	return nil, nil
}

func (e *employeeService) DeleteEmployee(id string) (*mongo.DeleteResult, error) {
	return nil, nil
}
