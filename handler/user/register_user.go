package user

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"log"
	"net/http"
	"webapp/entity"
	"webapp/handler"
)

type RegisterUserHandler struct {
	us        handler.UserRegister
	Validator *validator.Validate
}

func NewRegisterUserHandler(us handler.UserRegister, v *validator.Validate) *RegisterUserHandler {
	return &RegisterUserHandler{us: us, Validator: v}
}

func (h *RegisterUserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req RegisterUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("Error registering user: %s", err.Error())
		handler.RespondJSON(ctx, w, handler.ErrResponse{Error: err.Error()}, http.StatusInternalServerError)
		return
	}

	if err := h.Validator.Struct(req); err != nil {
		log.Printf("Error registering user: %s", err.Error())
		handler.RespondJSON(ctx, w, handler.ErrResponse{Error: err.Error()}, http.StatusBadRequest)
		return
	}

	user := &entity.User{
		Birthday:      req.Birthday,
		LastName:      req.LastName,
		FirstName:     req.FirstName,
		LastNameKana:  req.LastNameKana,
		FirstNameKana: req.FirstNameKana,
		Sex:           req.Sex,
		Prefecture:    req.Prefecture,
	}

	createdUser, err := h.us.RegisterUser(ctx, user)
	if err != nil {
		log.Printf("Error registering user: %s", err.Error())
		handler.RespondJSON(ctx, w, handler.ErrResponse{Error: err.Error()}, http.StatusInternalServerError)
		return
	}

	handler.RespondJSON(ctx, w, createdUser, http.StatusOK)
}
