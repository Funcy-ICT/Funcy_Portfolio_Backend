package handler

import (
	"backend/app/interfaces/response"
	"backend/app/usecase"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

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

	fmt.Println("test", workID)
	// do
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
