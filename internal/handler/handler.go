package handler

import "sellease-ai/internal/usecase"

type Handler struct {
	HealthHandler *healthHandler
	// UserHandler   *userHandler
	ProductHandler *productHandler
}

func Init(uc *usecase.Usecase) *Handler {
	return &Handler{
		HealthHandler: InitHealthHandler(),
		// UserHandler:   InitUserHandler(uc.User),
		ProductHandler: InitProductHandler(uc.Product),
	}
}
