package server

import (
	// "github.com/VAISHAKH-GK/SevaNear/internal/middleware"
	"github.com/gofiber/fiber/v3"
)

func (s *WebServer) RegisterRoutes() {
	// repos := repository.NewRepositories(s.DB.Queries)

	// authService := services.NewAuthService(
	// 	repos.User,
	// 	repos.Session,
	// 	repos.RefreshToken,
	// 	s.Config.JWTSecret,
	// )
	// authHandler := handlers.NewAuthHandler(authService)

	s.App.Get("/", s.HandleIndexRotue)

}

func (s *WebServer) HandleIndexRotue(c fiber.Ctx) error {
	return c.SendString("Index Router")
}
