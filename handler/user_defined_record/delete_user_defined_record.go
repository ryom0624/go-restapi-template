package userDefinedRecord

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"log"
	"net/http"
	"webapp/handler"
)

type DeleteUserDefinedRecordHandler struct {
	h         handler.UserDefinedRecordDeleter
	Validator *validator.Validate
}

func NewDeleteUserDefinedRecordHandler(drh handler.UserDefinedRecordDeleter, v *validator.Validate) *DeleteUserDefinedRecordHandler {
	return &DeleteUserDefinedRecordHandler{h: drh, Validator: v}
}

func (h *DeleteUserDefinedRecordHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req DeleteUserDefinedRecordRequest
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

	if err := h.h.DeleteUserDefinedRecord(ctx, req.Id); err != nil {
		log.Printf("Error registering user: %s", err.Error())
		handler.RespondJSON(ctx, w, handler.ErrResponse{Error: err.Error()}, http.StatusInternalServerError)
		return
	}

	handler.RespondJSON(ctx, w, `{"message": "success"}`, http.StatusOK)
}
