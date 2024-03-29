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

	"github.com/go-chi/chi"
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

	raw, user, err := h.workUseCase.ReadWork(workID)
	if err != nil {
		_ = response.ReturnErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	images := make([]response.Image, len(raw.ImageURLs))
	for i, img := range raw.ImageURLs {
		images[i] = response.Image{
			Image: img,
		}
	}

	tags := make([]response.Tag, len(raw.Tags))
	for i, tag := range raw.Tags {
		tags[i] = response.Tag{
			Tag: tag,
		}
	}

	res := &response.ReadWork{
		Title:       raw.Title,
		Description: raw.Description,
		Thumbnail:   raw.Thumbnail,
		UserIcon:    user.Icon,
		UserName:    user.DisplayName,
		WorkUserID:  raw.UserId,
		Images:      images,
		WorkUrl:     raw.WorkUrl,
		MovieUrl:    raw.MovieUrl,
		GroupID:     raw.GroupID,
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

func (h *WorkHandler) ReadWorks(w http.ResponseWriter, r *http.Request) {
	var numberOfWorks uint
	if n, err := strconv.ParseUint(chi.URLParam(r, "number"), 10, 32); err != nil {
		_ = response.ReturnErrorResponse(w, http.StatusBadRequest, "the number out of range")
		return
	} else if n == 0 {
		_ = response.ReturnErrorResponse(w, http.StatusBadRequest, "the number out of range")
		return
	} else {
		numberOfWorks = uint(n)
	}

	works, err := h.workUseCase.ReadWorks(numberOfWorks)
	if err != nil {
		_ = response.ReturnErrorResponse(w, http.StatusInternalServerError, err.Error())
	}

	worksRes := []response.ReadWorks{}
	for _, work := range *works {
		newWorkRes := response.ReadWorks{WorkID: work.WorkID, WorkUserID: work.UserId, UserName: work.DisplayName, Title: work.Title, Thumbnail: work.Thumbnail, Description: work.Description, Icon: work.Icon}
		worksRes = append(worksRes, newWorkRes)
	}

	res := response.ReadWorksList{
		Works: worksRes,
	}

	resBody, err := json.Marshal(res)
	if err != nil {
		_ = response.ReturnErrorResponse(w, http.StatusInternalServerError, err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", strconv.Itoa(len(resBody)))
	w.WriteHeader(http.StatusOK)
	w.Write(resBody)
}

func (h *WorkHandler) DeleteWork(w http.ResponseWriter, r *http.Request) {
	workID := chi.URLParam(r, "workID")

	err := h.workUseCase.DeleteWork(workID)
	if err != nil {
		e := response.UnwrapError(err)
		_ = response.ReturnErrorResponse(w, e.Code, e.Message)
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

func (h *WorkHandler) UpdateWork(w http.ResponseWriter, r *http.Request) {
	var req request.UpdateWorkRequest
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

	workID := chi.URLParam(r, "workID")

	err = h.workUseCase.UpdateWork(req, workID)
	if err != nil {
		e := response.UnwrapError(err)
		_ = response.ReturnErrorResponse(w, e.Code, e.Message)
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
