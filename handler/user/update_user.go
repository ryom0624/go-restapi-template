package user

import (
	"encoding/json"
	"log"
	"net/http"
	"webapp/entity"
	"webapp/handler"
)

type UpdateUserHandler struct {
	us handler.UserUpdater
}

func NewUpdateUserHandler(us handler.UserUpdater) *UpdateUserHandler {
	return &UpdateUserHandler{us: us}
}

func (h *UpdateUserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req UpdateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("Error updating user: %s", err.Error())
		handler.RespondJSON(ctx, w, handler.ErrResponse{Error: err.Error()}, http.StatusInternalServerError)
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

	putUser, err := h.us.UpdateUser(ctx, user)
	if err != nil {
		log.Printf("Error updating user: %s", err.Error())
		handler.RespondJSON(ctx, w, handler.ErrResponse{Error: err.Error()}, http.StatusInternalServerError)
		return
	}

	handler.RespondJSON(ctx, w, putUser, http.StatusOK)
}
