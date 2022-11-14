package handler

import (
	"backend/app/interfaces/request"
	"backend/app/interfaces/response"
	"backend/app/packages/utils"
	"backend/app/usecase"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

type WorkHandler struct {
	workUseCase *usecase.WorkUseCase
}

func NewWorkHandler(workUseCase *usecase.WorkUseCase) *WorkHandler {
	return &WorkHandler{
		workUseCase: workUseCase,
	}
}

func (h *WorkHandler) CreateWork(w http.ResponseWriter, r *http.Request) {
	var req request.CreateWorkRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		_ = response.ReturnErrorResponse(w, http.StatusBadRequest, "bad request")
		return
	}
	me, _ := utils.Validate(req)
	if me != nil {
		_ = response.ReturnValidationErrorResponse(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), me)
		return
	}

	userID := r.Context().Value("user_id")
	workID, err := h.workUseCase.CreateWork(req, userID.(string))

	if err != nil {
		_ = response.ReturnErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	res := response.WorkID{
		WorkID: workID,
	}

	resBody, err := json.Marshal(res)
	if err != nil {
		_ = response.ReturnErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", strconv.Itoa(len(resBody)))
	w.WriteHeader(http.StatusOK)
	w.Write(resBody)
}

func (h *WorkHandler) ReadWork(w http.ResponseWriter, r *http.Request) {

	workID := strings.TrimPrefix(r.URL.Path, "/work/")
	if workID == "" {
		_ = response.ReturnErrorResponse(w, http.StatusBadRequest, "bad request")
		return
	}

	raw, err := h.workUseCase.ReadWork(workID)
	if err != nil {
		_ = response.ReturnErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	images := make([]response.Image, len(raw.Images))
	for i, rimg := range raw.Images {
		images[i] = response.Image{
			Image: rimg.Image,
		}
	}

	tags := make([]response.Tag, len(raw.Tags))
	for i, rtag := range raw.Tags {
		tags[i] = response.Tag{
			Tag: rtag.Tag,
		}
	}

	res := &response.ReadWorkResponse{
		Title:       raw.Title,
		Description: raw.Description,
		Images:      images,
		WorkURL:     raw.WorkURL,
		MovieUrl:    raw.MovieUrl,
		Tags:        tags,
		Security:    raw.Security,
	}

	resBody, err := json.Marshal(res)
	if err != nil {
		_ = response.ReturnErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", strconv.Itoa(len(resBody)))
	w.WriteHeader(http.StatusOK)
	w.Write(resBody)

}
