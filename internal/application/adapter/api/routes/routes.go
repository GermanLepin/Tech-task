package routes

import (
	"database/sql"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

type (
	CreateUserHandler interface {
		CretaeUser(w http.ResponseWriter, r *http.Request)
	}

	DeleteUserHandler interface {
		DeleteUser(w http.ResponseWriter, r *http.Request)
	}
)

func (s *service) NewRoutes() http.Handler {
	router := chi.NewRouter()

	// specify who is allowed to connect
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	router.Route("/user", func(r chi.Router) {
		r.Post("/create", s.createUserHandler.CretaeUser)
		r.Delete("/delete", s.deleteUserHandler.DeleteUser)
	})

	router.Route("/balance", func(r chi.Router) {
		// r.Get("/info", s.balanceService.BalanceInfo)
		// r.Post("/replenishment", s.balanceService.BalanceReplenishment)
		// r.Patch("/debit", s.balanceService.BalanceDebit)
		// r.Patch("/user-to-user", s.balanceService.UserToUser)
	})

	router.Route("/descriptions", func(r chi.Router) {
		// r.Post("/add", s.descriptionService.AddDescription)
		// r.Get("/get", s.descriptionService.GetDescriptions)
	})

	return router
}

func New(
	connection *sql.DB,

	createUserHandler CreateUserHandler,
	deleteUserHandler DeleteUserHandler,
) *service {
	return &service{
		connection: connection,

		createUserHandler: createUserHandler,
		deleteUserHandler: deleteUserHandler,
	}
}

type service struct {
	connection *sql.DB

	createUserHandler CreateUserHandler
	deleteUserHandler DeleteUserHandler
}
