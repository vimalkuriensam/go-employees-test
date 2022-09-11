package services

import (
	"context"
	"time"

	"github.com/vimalkuriensam/go-employees-test/pkg/config"
	"github.com/vimalkuriensam/go-employees-test/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	collection := config.GetConfig().DataBase.Collections["employees"]
	ctx, cancelFunc := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancelFunc()
	entry := bson.D{
		{Key: "name", Value: employee.Name},
		{Key: "age", Value: employee.Age},
		{Key: "email", Value: employee.Email},
		{Key: "created_at", Value: time.Now().Local()},
		{Key: "updated_at", Value: time.Now().Local()},
	}
	return collection.InsertOne(ctx, entry)
}

func (e *employeeService) GetEmployee(id string) *mongo.SingleResult {
	collection := config.GetConfig().DataBase.Collections["employees"]
	docId, _ := primitive.ObjectIDFromHex(id)
	ctx, cancelFunc := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancelFunc()
	return collection.FindOne(ctx, bson.M{"_id": docId})
}

func (e *employeeService) UpdateEmployee(id string, employee models.Employee) (*mongo.UpdateResult, error) {
	collection := config.GetConfig().DataBase.Collections["employees"]
	docId, _ := primitive.ObjectIDFromHex(id)
	ctx, cancelFunc := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancelFunc()
	return collection.UpdateByID(ctx, docId, bson.D{
		primitive.E{Key: "$set", Value: bson.D{
			primitive.E{Key: "name", Value: employee.Name},
			primitive.E{Key: "age", Value: employee.Age},
			primitive.E{Key: "email", Value: employee.Email},
			primitive.E{Key: "update_at", Value: time.Now()},
		}},
	})
}

func (e *employeeService) DeleteEmployee(id string) (*mongo.DeleteResult, error) {
	collection := config.GetConfig().DataBase.Collections["employees"]
	docId, _ := primitive.ObjectIDFromHex(id)
	ctx, cancelFunc := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancelFunc()
	return collection.DeleteOne(ctx, bson.M{"_id": docId})
}
