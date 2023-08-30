package dailyRecord

import (
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"strconv"
	"webapp/handler"
)

const (
	defaultLimit  = 20
	defaultOffset = 0
)

type GetDailyRecordListHandler struct {
	h handler.DailyRecordGetter
}

func NewListDailyRecordHandler(h handler.DailyRecordGetter) *GetDailyRecordListHandler {
	return &GetDailyRecordListHandler{h: h}
}

func (h *GetDailyRecordListHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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

	createdDailyRecord, err := h.h.GetDailyRecordList(ctx, limit, offset)
	if err != nil {
		log.Printf("Error registering user: %s", err.Error())
		handler.RespondJSON(ctx, w, handler.ErrResponse{Error: err.Error()}, http.StatusInternalServerError)
		return
	}

	handler.RespondJSON(ctx, w, createdDailyRecord, http.StatusOK)
}

type GetDailyRecordHandler struct {
	us handler.DailyRecordGetter
}

func NewGetDailyRecordHandler(us handler.DailyRecordGetter) *GetDailyRecordHandler {
	return &GetDailyRecordHandler{us: us}
}

func (h *GetDailyRecordHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		handler.RespondJSON(ctx, w, handler.ErrResponse{Error: "required parameter is missing"}, http.StatusBadRequest)
		return
	}

	dailyRecord, err := h.us.GetDailyRecordById(ctx, int64(id))
	if err != nil {
		log.Printf("Error get daily record: %s", err.Error())
		handler.RespondJSON(ctx, w, handler.ErrResponse{Error: err.Error()}, http.StatusInternalServerError)
		return
	}

	handler.RespondJSON(ctx, w, dailyRecord, http.StatusOK)
}
