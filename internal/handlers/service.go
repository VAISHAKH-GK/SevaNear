package handlers

import (
	"github.com/VAISHAKH-GK/SevaNear/internal/db/generated"
	"github.com/gofiber/fiber/v3"
	"strconv"
)

type ServiceHandler struct {
	queries *generated.Queries
}

func NewServiceHandler(queries *generated.Queries) *ServiceHandler {
	return &ServiceHandler{
		queries: queries,
	}
}

func (s *ServiceHandler) PostHospital(c fiber.Ctx) error {
	type request struct {
		Name      string  `json:"name"`
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
		Address   string  `json:"address"`
		Contact   string  `json:"contact"`
	}

	var body request

	if err := c.Bind().Body(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	hospital, err := s.queries.CreateHospital(
		c.Context(),
		generated.CreateHospitalParams{
			Name:      body.Name,
			Latitude:  body.Latitude,
			Longitude: body.Longitude,
			Address:   &body.Address,
			Contact:   &body.Contact,
		},
	)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to create hospital",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(hospital)
}

func (s *ServiceHandler) GetHospitals(c fiber.Ctx) error {
	hospitals, err := s.queries.GetAllHospitals(c.Context())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "failed to fetch hospitals",
		})
	}

	return c.JSON(hospitals)
}

func (s *ServiceHandler) PostServiceType(c fiber.Ctx) error {
	type request struct {
		Name string `json:"name"`
	}

	var body request
	if err := c.Bind().Body(&body); err != nil || body.Name == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid service type",
		})
	}

	serviceType, err := s.queries.CreateServiceType(
		c.Context(),
		body.Name,
	)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "failed to create service type",
		})
	}

	return c.Status(201).JSON(serviceType)
}

func (s *ServiceHandler) GetServiceTypeByID(c fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid id"})
	}

	serviceType, err := s.queries.GetServiceTypeByID(
		c.Context(),
		int32(id),
	)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "service type not found",
		})
	}

	return c.JSON(serviceType)
}

func (s *ServiceHandler) PostService(c fiber.Ctx) error {
	type request struct {
		HospitalID    int32   `json:"hospital_id"`
		ServiceTypeID int32   `json:"service_type_id"`
		Name          string  `json:"name"`
		Provider      string  `json:"provider"`
		Description   string  `json:"description"`
		Timings       string  `json:"timings"`
		Eligibility   string  `json:"eligibility"`
		RequiredDocs  string  `json:"required_docs"`
		Contact       string  `json:"contact"`
		Latitude      float64 `json:"latitude"`
		Longitude     float64 `json:"longitude"`
	}

	var body request
	if err := c.Bind().Body(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	service, err := s.queries.CreateService(
		c.Context(),
		generated.CreateServiceParams{
			HospitalID:    body.HospitalID,
			ServiceTypeID: body.ServiceTypeID,
			Name:          body.Name,
			Provider:      &body.Provider,
			Description:   &body.Description,
			Timings:       &body.Timings,
			Eligibility:   &body.Eligibility,
			Contact:       &body.Contact,
			Latitude:      body.Latitude,
			Longitude:     body.Longitude,
		},
	)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "failed to create service",
		})
	}

	return c.Status(201).JSON(service)
}

func (s *ServiceHandler) GetServices(c fiber.Ctx) error {
	services, err := s.queries.GetAllServices(c.Context())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "failed to fetch services",
		})
	}

	return c.JSON(services)
}

func (s *ServiceHandler) GetServiceByID(c fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid id"})
	}

	service, err := s.queries.GetServiceByID(
		c.Context(),
		int32(id),
	)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "service not found",
		})
	}

	return c.JSON(service)
}

func (s *ServiceHandler) GetServicesByHospitalAndType(c fiber.Ctx) error {
	hospitalID, err := strconv.Atoi(c.Query("hospital_id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid hospital_id"})
	}

	serviceTypeID, err := strconv.Atoi(c.Query("service_type_id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid service_type_id"})
	}

	services, err := s.queries.GetServicesByHospitalID(
		c.Context(),
		int32(hospitalID),
	)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "failed to fetch services",
		})
	}

	filtered := make([]generated.Service, 0)
	for _, svc := range services {
		if svc.ServiceTypeID == int32(serviceTypeID) {
			filtered = append(filtered, svc)
		}
	}

	return c.JSON(filtered)
}

func (s *ServiceHandler) GetAllServiceTypes(c fiber.Ctx) error {
	serviceTypes, err := s.queries.GetAllServiceTypes(c.Context())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "failed to fetch service types",
		})
	}

	return c.JSON(serviceTypes)
}
