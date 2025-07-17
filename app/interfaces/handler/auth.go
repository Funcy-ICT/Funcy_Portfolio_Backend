package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"backend/app/interfaces/request"
	"backend/app/interfaces/response"
	"backend/app/packages/utils"
	"backend/app/packages/utils/auth"
	"backend/app/usecase"
)

const (
	CookieAccessToken  = "token"
	CookieRefreshToken = "refresh_token"
)

type AuthHandler struct {
	authUseCase *usecase.AuthUseCase
}

func NewAuthHandler(authUseCase *usecase.AuthUseCase) *AuthHandler {
	return &AuthHandler{
		authUseCase: authUseCase,
	}
}

func (h *AuthHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	var req request.SignUpRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Printf("SignUp failed: %v", err)
		_ = response.ReturnErrorResponse(w, http.StatusBadRequest, "An unexpected error occurred. Please try again later.")
		return
	}
	me, _ := utils.Validate(req)
	if me != nil {
		log.Printf("SignUp failed: %v", me)
		_ = response.ReturnErrorResponse(w, http.StatusBadRequest, "An unexpected error occurred. Please try again later.")
		return
	}

	userID, err := h.authUseCase.CreateAccount(req)
	if err != nil {
		log.Printf("SignUp failed: %v", err)
		_ = response.ReturnErrorResponse(w, http.StatusBadRequest, "An unexpected error occurred. Please try again later.")
		return
	}

	res := response.UserID{
		UserID: userID,
	}
	resBody, err := json.Marshal(res)
	if err != nil {
		log.Printf("SignUp failed: %v", err)
		_ = response.ReturnErrorResponse(w, http.StatusBadRequest, "An unexpected error occurred. Please try again later.")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", strconv.Itoa(len(resBody)))
	w.WriteHeader(http.StatusOK)
	w.Write(resBody)
}

func (h *AuthHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	var req request.SignInRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Printf("SignIn failed: %v", err)
		_ = response.ReturnErrorResponse(w, http.StatusBadRequest, "An unexpected error occurred. Please try again later.")
		return
	}
	me, _ := utils.Validate(req)
	if me != nil {
		log.Printf("SignIn failed: %v", me)
		_ = response.ReturnErrorResponse(w, http.StatusBadRequest, "An unexpected error occurred. Please try again later.")
		return
	}

	user, token, refreshToken, err := h.authUseCase.Login(req)
	if err != nil {
		log.Printf("SignIn failed: %v", err)
		_ = response.ReturnErrorResponse(w, http.StatusBadRequest, "An unexpected error occurred. Please try again later.")
		return
	}

	// アクセストークンをクッキーに設定
	cookie := &http.Cookie{
		Name:     CookieAccessToken,
		Value:    token,
		Path:     "/",
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
		//Secure: true, // HTTPS環境で有効化
	}
	http.SetCookie(w, cookie)

	// リフレッシュトークンを別のクッキーに設定
	refreshCookie := &http.Cookie{
		Name:     CookieRefreshToken,
		Value:    refreshToken,
		Path:     "/",
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
		//Secure: true, // HTTPS環境で有効化
	}
	http.SetCookie(w, refreshCookie)

	res := response.UserID{
		UserID: user.UserID,
	}

	resBody, err := json.Marshal(res)
	if err != nil {
		log.Printf("SignIn failed: %v", err)
		_ = response.ReturnErrorResponse(w, http.StatusBadRequest, "An unexpected error occurred. Please try again later.")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", strconv.Itoa(len(resBody)))
	w.WriteHeader(http.StatusOK)
	w.Write(resBody)
}

func (h *AuthHandler) SignInMobile(w http.ResponseWriter, r *http.Request) {
	var req request.SignInRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Printf("SignInMobile failed: %v", err)
		_ = response.ReturnErrorResponse(w, http.StatusBadRequest, "An unexpected error occurred. Please try again later.")
		return
	}
	me, _ := utils.Validate(req)
	if me != nil {
		log.Printf("SignInMobile failed: %v", me)
		_ = response.ReturnErrorResponse(w, http.StatusBadRequest, "An unexpected error occurred. Please try again later.")
		return
	}

	user, token, err := h.authUseCase.LoginMobile(req)
	if err != nil {
		log.Printf("SignInMobile failed: %v", err)
		_ = response.ReturnErrorResponse(w, http.StatusBadRequest, "An unexpected error occurred. Please try again later.")
		return
	}

	res := response.SignInResponse{
		UserID: user.UserID,
		Token:  token,
	}
	resBody, err := json.Marshal(res)
	if err != nil {
		log.Printf("SignInMobile failed: %v", err)
		_ = response.ReturnErrorResponse(w, http.StatusBadRequest, "An unexpected error occurred. Please try again later.")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", strconv.Itoa(len(resBody)))
	w.WriteHeader(http.StatusOK)
	w.Write(resBody)
}

func (h *AuthHandler) AuthCode(w http.ResponseWriter, r *http.Request) {
	var req request.AuthCodeRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Printf("AuthCode failed: %v", err)
		_ = response.ReturnErrorResponse(w, http.StatusBadRequest, "An unexpected error occurred. Please try again later.")
		return
	}

	me, _ := utils.Validate(req)
	if me != nil {
		log.Printf("AuthCode failed: %v", me)
		_ = response.ReturnErrorResponse(w, http.StatusBadRequest, "An unexpected error occurred. Please try again later.")
		return
	}

	err = h.authUseCase.CheckMail(req)
	if err != nil {
		log.Printf("AuthCode failed: %v", err)
		_ = response.ReturnErrorResponse(w, http.StatusBadRequest, "An unexpected error occurred. Please try again later.")
		return
	}

	jwt, _ := auth.IssueUserToken(req.UserID)

	cookie := &http.Cookie{
		Name:     CookieAccessToken,
		Value:    jwt,
		Path:     "/",
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
		//Secure: true, // HTTPS環境で有効化
	}
	http.SetCookie(w, cookie)

	res := response.Token{
		Token: jwt,
	}
	resBody, err := json.Marshal(res)
	if err != nil {
		log.Printf("AuthCode failed: %v", err)
		_ = response.ReturnErrorResponse(w, http.StatusBadRequest, "An unexpected error occurred. Please try again later.")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", strconv.Itoa(len(resBody)))
	w.WriteHeader(http.StatusOK)
	w.Write(resBody)
}

func (h *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
	// アクセストークンクッキーを削除
	cookie := &http.Cookie{
		Name:     CookieAccessToken,
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
		//Secure: true, // HTTPS環境で有効化
		MaxAge:   -1,
	}
	http.SetCookie(w, cookie)

	// リフレッシュトークンクッキーを削除
	refreshCookie := &http.Cookie{
		Name:     CookieRefreshToken,
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
		//Secure: true, // HTTPS環境で有効化
		MaxAge:   -1,
	}
	http.SetCookie(w, refreshCookie)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "logout successful"}`))
}

func (h *AuthHandler) CheckAuth(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user_id")
	if userID == nil {
		_ = response.ReturnErrorResponse(w, http.StatusUnauthorized, "authentication required")
		return
	}

	res := response.UserID{
		UserID: userID.(string),
	}
	resBody, err := json.Marshal(res)
	if err != nil {
		log.Printf("CheckAuth failed: %v", err)
		_ = response.ReturnErrorResponse(w, http.StatusInternalServerError, "An unexpected error occurred. Please try again later.")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", strconv.Itoa(len(resBody)))
	w.WriteHeader(http.StatusOK)
	w.Write(resBody)
}

func (h *AuthHandler) RefreshToken(w http.ResponseWriter, r *http.Request) {
	refreshCookie, err := r.Cookie(CookieRefreshToken)
	if err != nil {
		log.Printf("RefreshToken failed: %v", err)
		_ = response.ReturnErrorResponse(w, http.StatusBadRequest, "refresh token not found")
		return
	}

	newAccessToken, err := h.authUseCase.RefreshToken(refreshCookie.Value)
	if err != nil {
		log.Printf("RefreshToken failed: %v", err)
		_ = response.ReturnErrorResponse(w, http.StatusBadRequest, "invalid refresh token")
		return
	}

	// 新しいアクセストークンをクッキーに設定
	cookie := &http.Cookie{
		Name:     CookieAccessToken,
		Value:    newAccessToken,
		Path:     "/",
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
		//Secure: true, // HTTPS環境で有効化
	}
	http.SetCookie(w, cookie)

	res := response.Token{
		Token: newAccessToken,
	}
	resBody, err := json.Marshal(res)
	if err != nil {
		log.Printf("RefreshToken failed: %v", err)
		_ = response.ReturnErrorResponse(w, http.StatusInternalServerError, "An unexpected error occurred. Please try again later.")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", strconv.Itoa(len(resBody)))
	w.WriteHeader(http.StatusOK)
	w.Write(resBody)
}
