package handler

import (
	"backend/app/interfaces/request"
	"backend/app/interfaces/response"
	"backend/app/packages/utils"
	"backend/app/usecase"
	"encoding/json"
	"net/http"
	"strconv"
)

type GroupHandler struct {
	groupUseCase *usecase.GroupUseCase
}

func NewGroupHandler(groupUseCase *usecase.GroupUseCase) *GroupHandler {
	return &GroupHandler{
		groupUseCase: groupUseCase,
	}
}

func (h *GroupHandler) CreateGroup(w http.ResponseWriter, r *http.Request) {

	var req request.CreateGroupRequest
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

	groupID, err := h.groupUseCase.CreateGroup(req.Name, req.Description, req.LeaderEmail, req.Icon, req.GroupSkills)
	if err != nil {
		_ = response.ReturnErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	res := response.CreateGroupResponse{
		GroupID: groupID,
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
