package services

import (
	"context"
	"time"

	"github.com/vimalkuriensam/go-employees-test/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type EmployeeService interface {
	AddEmployee(models.Employee, *mongo.Collection) (*mongo.InsertOneResult, error)
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

func (e *employeeService) AddEmployee(employee models.Employee, collection *mongo.Collection) (*mongo.InsertOneResult, error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancelFunc()
	entry := bson.D{
		{Key: "name", Value: employee.Name},
		{Key: "age", Value: employee.Age},
		{Key: "email", Value: employee.Email},
		{Key: "created_at", Value: time.Now()},
		{Key: "updated_at", Value: time.Now()},
	}
	return collection.InsertOne(ctx, entry)
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
