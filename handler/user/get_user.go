package user

import (
	"log"
	"net/http"
	"webapp/handler"
)

type GetUserHandler struct {
	us handler.UserGetter
}

func NewGetUserHandler(us handler.UserGetter) *GetUserHandler {
	return &GetUserHandler{us: us}
}

func (h *GetUserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user, err := h.us.GetUserById(ctx)
	if err != nil {
		log.Printf("Error get user by id: %s", err.Error())
		handler.RespondJSON(ctx, w, handler.ErrResponse{Error: err.Error()}, http.StatusInternalServerError)
		return
	}

	handler.RespondJSON(ctx, w, user, http.StatusOK)
}
