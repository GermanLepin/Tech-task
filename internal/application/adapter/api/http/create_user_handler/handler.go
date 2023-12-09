package create_user_handler

import (
	"balance-service/internal/application/dto"
	"errors"

	"context"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

type (
	CretaeUserService interface {
		CreateUser(ctx context.Context, user dto.User) error
	}

	JsonService interface {
		ErrorJSON(w http.ResponseWriter, err error, statusCode int) error
	}
)

func (h *handler) CretaeUser(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	var user dto.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		h.jsonService.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	if err := validateBalance(user.Balance); err != nil {
		h.jsonService.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	user.ID = uuid.New()
	if err := h.cretaeUserService.CreateUser(ctx, user); err != nil {
		h.jsonService.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	encoder := json.NewEncoder(w)
	err := encoder.Encode(&user)
	if err != nil {
		h.jsonService.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}
}

func validateBalance(balance float64) error {
	if balance < 0 {
		return errors.New("balance cannot be negative")
	}

	return nil
}

func New(
	cretaeUserService CretaeUserService,
	jsonService JsonService,
) *handler {
	return &handler{
		cretaeUserService: cretaeUserService,
		jsonService:       jsonService,
	}
}

type handler struct {
	cretaeUserService CretaeUserService
	jsonService       JsonService
}
