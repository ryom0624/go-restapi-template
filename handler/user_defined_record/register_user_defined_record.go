package userDefinedRecord

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"log"
	"net/http"
	"webapp/entity"
	"webapp/handler"
)

type RegisterUserDefinedRecordHandler struct {
	us        handler.UserDefinedRecordRegister
	Validator *validator.Validate
}

func NewRegisterUserDefinedRecordHandler(us handler.UserDefinedRecordRegister, v *validator.Validate) *RegisterUserDefinedRecordHandler {
	return &RegisterUserDefinedRecordHandler{us: us, Validator: v}
}

func (h *RegisterUserDefinedRecordHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req RegisterUserDefinedRecordRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("Error registering user: %s", err.Error())
		handler.RespondJSON(ctx, w, handler.ErrResponse{Error: err.Error()}, http.StatusInternalServerError)
		return
	}

	h.Validator.RegisterValidation(entity.ValidateUserDefinedRecordTypeName, entity.ValidateUserDefinedRecordTypeFunc)

	if err := h.Validator.Struct(req); err != nil {
		log.Printf("Error registering user: %s", err.Error())
		handler.RespondJSON(ctx, w, handler.ErrResponse{Error: err.Error()}, http.StatusBadRequest)
		return
	}

	udr := &entity.UserDefinedRecord{
		ItemName: req.ItemName,
		UnitType: req.UnitType,
	}

	createdUserDefinedRecord, err := h.us.RegisterUserDefinedRecord(ctx, udr)
	if err != nil {
		log.Printf("Error registering user: %s", err.Error())
		handler.RespondJSON(ctx, w, handler.ErrResponse{Error: err.Error()}, http.StatusInternalServerError)
		return
	}

	handler.RespondJSON(ctx, w, createdUserDefinedRecord, http.StatusOK)
}
