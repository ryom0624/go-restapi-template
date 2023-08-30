package dailyRecord

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"log"
	"net/http"
	"webapp/entity"
	"webapp/handler"
)

type UpdateDailyRecordHandler struct {
	h         handler.DailyRecordUpdater
	Validator *validator.Validate
}

func NewUpdateDailyRecordHandler(us handler.DailyRecordUpdater, v *validator.Validate) *UpdateDailyRecordHandler {
	return &UpdateDailyRecordHandler{h: us, Validator: v}
}

func (h *UpdateDailyRecordHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req UpdateDailyRecordRequest
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

	optionalRecords := make([]entity.OptionalRecord, len(req.OptionalRecords))
	for i, optionalRecord := range req.OptionalRecords {
		optionalRecords[i] = entity.OptionalRecord{
			Id:            optionalRecord.Id,
			DailyRecordId: req.Id,
			UserDefinedId: optionalRecord.UserDefinedId,
			Value:         optionalRecord.Value,
		}
	}
	dailyRecord := &entity.DailyRecord{
		Id:              req.Id,
		Weather:         req.Weather,
		Memo:            req.Memo,
		OptionalRecords: optionalRecords,
	}

	updatedDailyRecord, err := h.h.UpdateDailyRecord(ctx, dailyRecord)
	if err != nil {
		log.Printf("Error updating user daily record: error: %s, value: %+v", err.Error(), dailyRecord)
		handler.RespondJSON(ctx, w, handler.ErrResponse{Error: err.Error()}, http.StatusInternalServerError)
		return
	}

	handler.RespondJSON(ctx, w, updatedDailyRecord, http.StatusOK)
}
