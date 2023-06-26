package handler

import (
	"backend/app/interfaces/response"
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
		log.Println(err)
		_ = response.ReturnErrorResponse(w, http.StatusBadRequest, "bad request")
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
		Works:           *worksRes,
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
