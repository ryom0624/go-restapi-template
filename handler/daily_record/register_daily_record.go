package dailyRecord

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"log"
	"net/http"
	"webapp/entity"
	"webapp/handler"
)

type RegisterDailyRecordHandler struct {
	h         handler.DailyRecordRegister
	Validator *validator.Validate
}

func NewRegisterDailyRecordHandler(us handler.DailyRecordRegister, v *validator.Validate) *RegisterDailyRecordHandler {
	return &RegisterDailyRecordHandler{h: us, Validator: v}
}

func (h *RegisterDailyRecordHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req RegisterDailyRecordRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("decode error registering daily record: %s", err.Error())
		handler.RespondJSON(ctx, w, handler.ErrResponse{Error: err.Error()}, http.StatusInternalServerError)
		return
	}

	if err := h.Validator.Struct(req); err != nil {
		log.Printf("validate error daily record: %s", err.Error())
		handler.RespondJSON(ctx, w, handler.ErrResponse{Error: err.Error()}, http.StatusBadRequest)
		return
	}

	optionalRecords := make([]entity.OptionalRecord, len(req.OptionalRecords))
	for i, optionalRecord := range req.OptionalRecords {
		optionalRecords[i] = entity.OptionalRecord{
			UserDefinedId: optionalRecord.UserDefinedId,
			Value:         optionalRecord.Value,
		}
	}
	dailyRecord := &entity.DailyRecord{
		Weather:         req.Weather,
		Memo:            req.Memo,
		OptionalRecords: optionalRecords,
	}

	createdDailyRecord, err := h.h.RegisterDailyRecord(ctx, dailyRecord)
	if err != nil {
		log.Printf("Error registering user: %s", err.Error())
		handler.RespondJSON(ctx, w, handler.ErrResponse{Error: err.Error()}, http.StatusInternalServerError)
		return
	}

	handler.RespondJSON(ctx, w, createdDailyRecord, http.StatusOK)
}
