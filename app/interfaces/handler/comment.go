package handler

import (
	"backend/app/interfaces/request"
	"backend/app/interfaces/response"
	"backend/app/packages/utils"
	"backend/app/usecase"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

type CommentHandler struct {
	commentUseCase *usecase.CommentUseCase
}

func NewCommentHandler(commentUseCase *usecase.CommentUseCase) *CommentHandler {
	return &CommentHandler{
		commentUseCase: commentUseCase,
	}
}

func (h *CommentHandler) GetComment(w http.ResponseWriter, r *http.Request) {
	// get params
	workID := chi.URLParam(r, "worksID")

	comments, err := h.commentUseCase.GetComment(workID)
	if err != nil {
		log.Println(err)
		_ = response.ReturnErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	// create response
	resBody, err := json.Marshal(comments)
	if err != nil {
		_ = response.ReturnErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resBody)
}

func (h *CommentHandler) CreateComment(w http.ResponseWriter, r *http.Request) {
	var req request.CreateCommentRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		_ = response.ReturnErrorResponse(w, http.StatusBadRequest, "bad request")
		return
	}

	me, err := utils.Validate(req)
	if err != nil {
		_ = response.ReturnErrorResponse(w, http.StatusInternalServerError, "bad request")
		return
	}
	if me != nil {
		_ = response.ReturnValidationErrorResponse(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), me)
		return
	}

	commentID, err := h.commentUseCase.CreateComment(req.UserID, req.WorksID, req.Content)
	if err != nil {
		_ = response.ReturnErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	res := response.CreateCommentResponse{
		StatusCode: http.StatusOK,
		Message:    "Comment created successfully",
		ID:         commentID,
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
