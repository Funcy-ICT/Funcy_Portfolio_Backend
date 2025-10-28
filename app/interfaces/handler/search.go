package handler

import (
	"backend/app/interfaces/response"
	"backend/app/usecase"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type SearchHandler struct {
	searchUseCase *usecase.SearchUseCase
}

func NewSearchHandler(searchUseCase *usecase.SearchUseCase) *SearchHandler {
	return &SearchHandler{
		searchUseCase: searchUseCase,
	}
}

func (h *SearchHandler) SearchWorks(w http.ResponseWriter, r *http.Request) {
	// クエリパラメータを取得
	keyword := r.URL.Query().Get("q")
	if keyword == "" {
		_ = response.ReturnErrorResponse(w, http.StatusBadRequest, "Search keyword is required")
		return
	}

	limitStr := r.URL.Query().Get("limit")
	limit := uint(100) // デフォルト値
	if limitStr != "" {
		parsedLimit, err := strconv.ParseUint(limitStr, 10, 32)
		if err != nil {
			_ = response.ReturnErrorResponse(w, http.StatusBadRequest, "Invalid limit parameter")
			return
		}
		limit = uint(parsedLimit)
	}

	scope := r.URL.Query().Get("scope")
	if scope == "" {
		scope = "all" // デフォルト値
	}

	// 作品検索を実行
	works, err := h.searchUseCase.SearchWorks(keyword, limit, scope)
	if err != nil {
		log.Printf("SearchWorks failed: %v", err)
		_ = response.ReturnErrorResponse(w, http.StatusInternalServerError, "An unexpected error occurred. Please try again later.")
		return
	}

	// レスポンスを作成
	readWorks := make([]response.ReadWorks, len(*works))
	for i, work := range *works {
		readWorks[i] = response.ReadWorks{
			WorkID:      work.WorkID,
			Title:       work.Title,
			Thumbnail:   work.Thumbnail,
			Description: work.Description,
			Icon:        work.Icon,
			UserID:      work.UserID,
			Security:    work.Security,
		}
	}
	res := &response.ReadWorksList{Works: readWorks}
	resBody, err := json.Marshal(res)
	if err != nil {
		log.Printf("SearchWorks failed: %v", err)
		_ = response.ReturnErrorResponse(w, http.StatusInternalServerError, "An unexpected error occurred. Please try again later.")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", strconv.Itoa(len(resBody)))
	w.WriteHeader(http.StatusOK)
	w.Write(resBody)
}

func (h *SearchHandler) SearchUsers(w http.ResponseWriter, r *http.Request) {
	// クエリパラメータを取得
	keyword := r.URL.Query().Get("q")
	if keyword == "" {
		_ = response.ReturnErrorResponse(w, http.StatusBadRequest, "Search keyword is required")
		return
	}

	limitStr := r.URL.Query().Get("limit")
	limit := uint(100)
	if limitStr != "" {
		parsedLimit, err := strconv.ParseUint(limitStr, 10, 32)
		if err != nil {
			_ = response.ReturnErrorResponse(w, http.StatusBadRequest, "Invalid limit parameter")
			return
		}
		limit = uint(parsedLimit)
	}

	// ユーザー検索を実行
	users, err := h.searchUseCase.SearchUsers(keyword, limit)
	if err != nil {
		log.Printf("SearchUsers failed: %v", err)
		_ = response.ReturnErrorResponse(w, http.StatusInternalServerError, "An unexpected error occurred. Please try again later.")
		return
	}

	// レスポンスを作成
	type SearchUsersResponse struct {
		Users []interface{} `json:"users"`
	}

	// entity.UserSearchResultをinterface{}に変換
	userInterfaces := make([]interface{}, len(*users))
	for i, user := range *users {
		userInterfaces[i] = map[string]interface{}{
			"userID":      user.UserID,
			"displayName": user.DisplayName,
			"icon":        user.Icon,
			"course":      user.Course,
			"skills":      user.Skills,
		}
	}

	res := SearchUsersResponse{Users: userInterfaces}
	resBody, err := json.Marshal(res)
	if err != nil {
		log.Printf("SearchUsers failed: %v", err)
		_ = response.ReturnErrorResponse(w, http.StatusInternalServerError, "An unexpected error occurred. Please try again later.")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", strconv.Itoa(len(resBody)))
	w.WriteHeader(http.StatusOK)
	w.Write(resBody)
}
