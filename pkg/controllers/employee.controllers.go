package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"

	"github.com/go-chi/chi/v5"
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
	cfg.WriteJSON(w, http.StatusCreated, employee, fmt.Sprintf("Employee with id %s created", employee.ID))
}

func (e *employee) GetEmployee(w http.ResponseWriter, req *http.Request) {
	cfg := config.GetConfig()
	id := chi.URLParam(req, "id")
	var employee models.Employee = models.Employee{}
	err := e.service.GetEmployee(id).Decode(&employee)
	if err != nil {
		cfg.ErrorJSON(w, req.URL.Path, err.Error(), http.StatusBadRequest)
		return
	}
	cfg.WriteJSON(w, http.StatusOK, employee, fmt.Sprintf("Employee with id %v fetched successfully", id))
}

func (e *employee) UpdateEmployee(w http.ResponseWriter, req *http.Request) {
	id := chi.URLParam(req, "id")
	cfg := config.GetConfig()
	var employee *models.Employee = &models.Employee{}
	err := e.service.GetEmployee(id).Decode(employee)
	var priorData models.Employee = *employee
	if err != nil {
		cfg.ErrorJSON(w, req.URL.Path, err.Error(), http.StatusBadRequest)
		return
	}
	go cfg.ReadJSON(req)
	data := (<-cfg.DataChan).(config.ReadValue)
	err = services.AcceptableFields(data.D.(map[string]interface{}), services.AvailableFields["update"])
	if err != nil {
		cfg.ErrorJSON(w, req.URL.Path, err.Error(), http.StatusBadRequest)
		return
	}
	for key, value := range data.D.(map[string]interface{}) {
		field, err := services.GetStructFieldByTag(key, *employee)
		if err != nil {
			cfg.ErrorJSON(w, req.URL.Path, err.Error(), http.StatusBadRequest)
			return
		}
		reflect.ValueOf(employee).Elem().FieldByName(field).Set(reflect.ValueOf(value))
	}
	result, err := e.service.UpdateEmployee(id, *employee)
	if err != nil {
		cfg.ErrorJSON(w, req.URL.Path, err.Error(), http.StatusInternalServerError)
		return
	}
	updateResult := models.EmployeeUpdateData{
		ID:           id,
		UpdateCount:  int(result.ModifiedCount),
		PreviousData: priorData,
		UpdatedData:  *employee,
	}
	message := fmt.Sprintf("Employee with id %v updated successfully", id)
	cfg.WriteJSON(w, http.StatusOK, updateResult, message)
}

func (e *employee) DeleteEmployee(w http.ResponseWriter, req *http.Request) {
	id := chi.URLParam(req, "id")
	cfg := config.GetConfig()
	var employee models.Employee = models.Employee{}
	if err := e.service.GetEmployee(id).Decode(&employee); err != nil {
		cfg.ErrorJSON(w, req.URL.Path, err.Error(), http.StatusBadRequest)
		return
	}
	result, err := e.service.DeleteEmployee(id)
	if err != nil {
		cfg.ErrorJSON(w, req.URL.Path, err.Error(), http.StatusInternalServerError)
		return
	}
	deleteResult := models.EmployeeDeleteData{
		ID:          id,
		DeleteCount: int(result.DeletedCount),
		DeleteData:  employee,
	}
	message := fmt.Sprintf("Employee with id %v deleted successfully", id)
	cfg.WriteJSON(w, http.StatusOK, deleteResult, message)
}
