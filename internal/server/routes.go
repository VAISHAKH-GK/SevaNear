package server

import (
	"github.com/VAISHAKH-GK/SevaNear/internal/handlers"
	"github.com/gofiber/fiber/v3"
)

func (s *WebServer) RegisterRoutes() {
	serviceHandler := handlers.NewServiceHandler(s.DB.Queries)

	s.App.Get("/", s.HandleIndexRotue)

	s.App.Post("/hospitals", serviceHandler.PostHospital)
	s.App.Get("/hospitals", serviceHandler.GetHospitals)

	s.App.Post("/service-types", serviceHandler.PostServiceType)
	s.App.Get("/service-types", serviceHandler.GetAllServiceTypes)
	s.App.Get("/service-types/:id", serviceHandler.GetServiceTypeByID)

	s.App.Post("/services", serviceHandler.PostService)
	s.App.Get("/services", serviceHandler.GetServices)
	s.App.Get("/services/filter", serviceHandler.GetServicesByHospitalAndType)
	s.App.Get("/services/:id", serviceHandler.GetServiceByID)
}

func (s *WebServer) HandleIndexRotue(c fiber.Ctx) error {
	return c.SendString("Index Router")
}
