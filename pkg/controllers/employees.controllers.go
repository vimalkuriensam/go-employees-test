package controllers

import "net/http"

type Employee interface {
	CreateEmployees(w http.ResponseWriter, req *http.Request)
	GetEmployee(w http.ResponseWriter, req *http.Request)
	UpdateEmployee(w http.ResponseWriter, req *http.Request)
	DeleteEmployee(w http.ResponseWriter, req *http.Request)
}

type employee struct{}

func New() Employee {
	return &employee{}
}

func (e *employee) CreateEmployees(w http.ResponseWriter, req *http.Request) {}

func (e *employee) GetEmployee(w http.ResponseWriter, req *http.Request) {}

func (e *employee) UpdateEmployee(w http.ResponseWriter, req *http.Request) {}

func (e *employee) DeleteEmployee(w http.ResponseWriter, req *http.Request) {}
