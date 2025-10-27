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

type UserinfoHandler struct {
	userinfoUseCase *usecase.UserinfoUseCase
}

func NewUserinfoHandler(userinfoUseCase *usecase.UserinfoUseCase) *UserinfoHandler {
	return &UserinfoHandler{
		userinfoUseCase: userinfoUseCase,
	}
}

func (h *UserinfoHandler) GetUserinfo(w http.ResponseWriter, r *http.Request) {
	// get params
	userID := chi.URLParam(r, "userID")

	// do
	userinfo, works, err := h.userinfoUseCase.GetUserinfo(userID)
	if err != nil {
		log.Printf("GetUserinfo failed: %v", err)
		_ = response.ReturnErrorResponse(w, http.StatusBadRequest, "An unexpected error occurred. Please try again later.")
		return
	}

	// create response
	sns := new([]string)
	for _, item := range *userinfo.SNS {
		*sns = append(*sns, item.SnsURL)
	}

	group := new([]string)
	for _, item := range *userinfo.JoinedGroups {
		*group = append(*group, item.GroupName)
	}

	skills := new([]string)
	for _, item := range *userinfo.Skills {
		*skills = append(*skills, item.SkillName)
	}

	worksRes := new([]response.ReadWorks)
	for i := range *works {
		n := (*works)[i]
		newWorkRes := response.ReadWorks{WorkID: n.WorkID, Title: n.Title, Thumbnail: n.Thumbnail, Description: n.Description, Icon: n.Icon}
		*worksRes = append(*worksRes, newWorkRes)
	}

	res := &response.UserInfo{
		Icon:            userinfo.Profile.Icon,
		HeaderImagePath: userinfo.Profile.HeaderImagePath,
		UserDescription: userinfo.Profile.Biography,
		SNS:             *sns,
		Group:           *group,
		Skills:          *skills,
		DisplayName:     userinfo.Profile.DisplayName,
		Course:          userinfo.Course,
		Mail:            userinfo.Profile.Mail,
		Works:           *worksRes,
	}

	resBody, err := json.Marshal(res)
	if err != nil {
		log.Printf("GetUserinfo failed: %v", err)
		_ = response.ReturnErrorResponse(w, http.StatusInternalServerError, "An unexpected error occurred. Please try again later.")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", strconv.Itoa(len(resBody)))
	w.WriteHeader(http.StatusOK)
	w.Write(resBody)
}

func (h *UserinfoHandler) PutUserinfo(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "userID")

	var req request.UpdateUserInfo
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Printf("PutUserinfo failed: %v", err)
		_ = response.ReturnErrorResponse(w, http.StatusBadRequest, "An unexpected error occurred. Please try again later.")
		return
	}

	ve, _ := utils.Validate(req)
	if ve != nil {
		log.Printf("PutUserinfo failed: validation error: %v", ve)
		_ = response.ReturnValidationErrorResponse(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), ve)
		return
	}

	if err := h.userinfoUseCase.UpdateUserinfo(userID, &req); err != nil {
		log.Printf("PutUserinfo failed: %v", err)
		_ = response.ReturnErrorResponse(w, http.StatusInternalServerError, "An unexpected error occurred. Please try again later.")
		return
	}

	userinfo, works, err := h.userinfoUseCase.GetUserinfo(userID)
	if err != nil {
		log.Printf("PutUserinfo failed: %v", err)
		_ = response.ReturnErrorResponse(w, http.StatusBadRequest, "An unexpected error occurred. Please try again later.")
		return
	}

	// create response
	sns := new([]string)
	for _, item := range *userinfo.SNS {
		*sns = append(*sns, item.SnsURL)
	}

	group := new([]string)
	for _, item := range *userinfo.JoinedGroups {
		*group = append(*group, item.GroupName)
	}

	skills := new([]string)
	for _, item := range *userinfo.Skills {
		*skills = append(*skills, item.SkillName)
	}

	worksRes := new([]response.ReadWorks)
	for i := range *works {
		n := (*works)[i]
		newWorkRes := response.ReadWorks{WorkID: n.WorkID, Title: n.Title, Thumbnail: n.Thumbnail, Description: n.Description, Icon: n.Icon}
		*worksRes = append(*worksRes, newWorkRes)
	}

	res := &response.UserInfo{
		Icon:            userinfo.Profile.Icon,
		HeaderImagePath: userinfo.Profile.HeaderImagePath,
		UserDescription: userinfo.Profile.Biography,
		SNS:             *sns,
		Group:           *group,
		Skills:          *skills,
		DisplayName:     userinfo.Profile.DisplayName,
		Course:          userinfo.Course,
		Mail:            userinfo.Profile.Mail,
		Works:           *worksRes,
	}

	resBody, err := json.Marshal(res)
	if err != nil {
		log.Printf("PutUserinfo failed: %v", err)
		_ = response.ReturnErrorResponse(w, http.StatusInternalServerError, "An unexpected error occurred. Please try again later.")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", strconv.Itoa(len(resBody)))
	w.WriteHeader(http.StatusOK)
	w.Write(resBody)
}
