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
		log.Printf("GetComment failed: %v", err)
		_ = response.ReturnErrorResponse(w, http.StatusInternalServerError, "An unexpected error occurred. Please try again later.")
		return
	}

	// create response
	resBody, err := json.Marshal(comments)
	if err != nil {
		log.Printf("GetComment failed: %v", err)
		_ = response.ReturnErrorResponse(w, http.StatusInternalServerError, "An unexpected error occurred. Please try again later.")
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
		log.Printf("CreateComment failed: %v", err)
		_ = response.ReturnErrorResponse(w, http.StatusBadRequest, "An unexpected error occurred. Please try again later.")
		return
	}

	me, _ := utils.Validate(req)
	if me != nil {
		log.Printf("CreateComment failed: validation error: %v", me)
		_ = response.ReturnValidationErrorResponse(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), me)
		return
	}

	commentID, err := h.commentUseCase.CreateComment(req.UserID, req.WorksID, req.Content)
	if err != nil {
		log.Printf("CreateComment failed: %v", err)
		_ = response.ReturnErrorResponse(w, http.StatusInternalServerError, "An unexpected error occurred. Please try again later.")
		return
	}

	res := response.CreateCommentResponse{
		StatusCode: http.StatusOK,
		Message:    "Comment created successfully",
		ID:         commentID,
	}

	resBody, err := json.Marshal(res)
	if err != nil {
		log.Printf("CreateComment failed: %v", err)
		_ = response.ReturnErrorResponse(w, http.StatusInternalServerError, "An unexpected error occurred. Please try again later.")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", strconv.Itoa(len(resBody)))
	w.WriteHeader(http.StatusOK)
	w.Write(resBody)
}
