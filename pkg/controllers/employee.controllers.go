package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/vimalkuriensam/go-employees-test/pkg/config"
	"github.com/vimalkuriensam/go-employees-test/pkg/models"
	"github.com/vimalkuriensam/go-employees-test/pkg/services"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Employee interface {
	CreateEmployees(w http.ResponseWriter, req *http.Request)
	GetEmployee(w http.ResponseWriter, req *http.Request)
	UpdateEmployee(w http.ResponseWriter, req *http.Request)
	DeleteEmployee(w http.ResponseWriter, req *http.Request)
}

type employee struct {
	service services.EmployeeService
}

func New(service services.EmployeeService) Employee {
	return &employee{
		service: service,
	}
}

func (e *employee) CreateEmployees(w http.ResponseWriter, req *http.Request) {
	cfg := config.GetConfig()
	go cfg.ReadJSON(req)
	data := (<-cfg.DataChan).(config.ReadValue)
	err := services.AcceptableFields(data.D.(map[string]interface{}), services.AvailableFields["create"])
	if err != nil {
		cfg.ErrorJSON(w, req.URL.Path, err.Error(), http.StatusBadRequest)
		return
	}
	var employee *models.Employee = &models.Employee{}
	if err = json.Unmarshal(data.B, &employee); err != nil {
		cfg.ErrorJSON(w, req.URL.Path, err.Error(), http.StatusBadRequest)
		return
	}
	id, err := e.service.AddEmployee(*employee)
	if err != nil {
		cfg.ErrorJSON(w, req.URL.Path, err.Error(), http.StatusInternalServerError)
		return
	}
	employee.ID = id.InsertedID.(primitive.ObjectID).Hex()
	cfg.WriteJSON(w, http.StatusCreated, employee, fmt.Sprintf("Log with id %s created", employee.ID))
}

func (e *employee) GetEmployee(w http.ResponseWriter, req *http.Request) {}

func (e *employee) UpdateEmployee(w http.ResponseWriter, req *http.Request) {}

func (e *employee) DeleteEmployee(w http.ResponseWriter, req *http.Request) {}
