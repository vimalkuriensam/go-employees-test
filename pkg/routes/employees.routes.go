package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/vimalkuriensam/go-employees-test/pkg/controllers"
)

var (
	employeeController controllers.Employee = controllers.New()
)

func employeesRoutes(r chi.Router) {
	r.Post("/addEmployee", employeeController.CreateEmployees)
	r.Get("/getEmployee/:id", employeeController.GetEmployee)
	r.Patch("/updateEmployee/:id", employeeController.UpdateEmployee)
	r.Delete("/deleteEmployee/:id", employeeController.DeleteEmployee)
}
