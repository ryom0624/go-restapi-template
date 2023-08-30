package userDefinedRecord

import (
	"log"
	"net/http"
	"strconv"
	"webapp/handler"
)

const (
	defaultLimit  = 20
	defaultOffset = 0
)

type ListUserDefinedRecordHandler struct {
	h handler.UserDefinedRecordGetter
}

func NewListUserDefinedRecordHandler(h handler.UserDefinedRecordGetter) *ListUserDefinedRecordHandler {
	return &ListUserDefinedRecordHandler{h: h}
}

func (h *ListUserDefinedRecordHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	limitStr := r.URL.Query().Get("limit")
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		limit = defaultLimit
	}

	offsetStr := r.URL.Query().Get("offset")
	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		offset = defaultOffset
	}

	userDefinedRecords, err := h.h.GetUserDefinedRecordList(ctx, limit, offset)
	if err != nil {
		log.Printf("Error registering user: %s", err.Error())
		handler.RespondJSON(ctx, w, handler.ErrResponse{Error: err.Error()}, http.StatusInternalServerError)
		return
	}

	handler.RespondJSON(ctx, w, userDefinedRecords, http.StatusOK)
}

type GetUserDefinedRecordHandler struct {
	us handler.UserDefinedRecordGetter
}

func NewGetUserDefinedRecordHandler(us handler.UserDefinedRecordGetter) *GetUserDefinedRecordHandler {
	return &GetUserDefinedRecordHandler{us: us}
}

func (h *GetUserDefinedRecordHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Printf("Error parsing id: %s", err.Error())
		handler.RespondJSON(ctx, w, handler.ErrResponse{Error: err.Error()}, http.StatusBadRequest)
		return
	}

	userDefinedRecord, err := h.us.GetUserDefinedRecordById(ctx, int64(id))
	if err != nil {
		log.Printf("Error registering user: %s", err.Error())
		handler.RespondJSON(ctx, w, handler.ErrResponse{Error: err.Error()}, http.StatusInternalServerError)
		return
	}

	handler.RespondJSON(ctx, w, userDefinedRecord, http.StatusOK)
}
