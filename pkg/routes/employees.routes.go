package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/vimalkuriensam/go-employees-test/pkg/controllers"
	"github.com/vimalkuriensam/go-employees-test/pkg/services"
)

var (
	employeeService    services.EmployeeService = services.New()
	employeeController controllers.Employee     = controllers.New(employeeService)
)

func employeesRoutes(r chi.Router) {
	r.Post("/addEmployee", employeeController.CreateEmployees)
	r.Get("/getEmployee/{id}", employeeController.GetEmployee)
	r.Patch("/updateEmployee/{id}", employeeController.UpdateEmployee)
	r.Delete("/deleteEmployee/{id}", employeeController.DeleteEmployee)
}
