package dailyRecord

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"log"
	"net/http"
	"webapp/handler"
)

type DeleteDailyRecordHandler struct {
	h         handler.DailyRecordDeleter
	Validator *validator.Validate
}

func NewDeleteDailyRecordHandler(drh handler.DailyRecordDeleter, v *validator.Validate) *DeleteDailyRecordHandler {
	return &DeleteDailyRecordHandler{h: drh, Validator: v}
}

func (h *DeleteDailyRecordHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req DeleteDailyRecordRequest
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

	if err := h.h.DeleteDailyRecord(ctx, req.Id); err != nil {
		log.Printf("Error registering user: %s", err.Error())
		handler.RespondJSON(ctx, w, handler.ErrResponse{Error: err.Error()}, http.StatusInternalServerError)
		return
	}

	handler.RespondJSON(ctx, w, `{"message": "success"}`, http.StatusOK)
}
