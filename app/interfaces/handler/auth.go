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

	userID, err := h.authUseCase.CreateAccount(req)
	if err != nil {
		_ = response.ReturnErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	res := response.UserID{
		UserID: userID,
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

func (h *AuthHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	var req request.SignInRequest
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

	user, token, err := h.authUseCase.Login(req)
	if err != nil {
		_ = response.ReturnErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	cookie := &http.Cookie{
		Name:     "token",
		Value:    token,
		HttpOnly: true,
		Secure:   true,
	}
	http.SetCookie(w, cookie)

	res := response.UserID{
		UserID: user.UserID,
	}

	resBody, err := json.Marshal(res)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
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

	user, token, err := h.authUseCase.LoginMobile(req)
	if err != nil {
		_ = response.ReturnErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	res := response.SignInResponse{
		UserID: user.UserID,
		Token:  token,
	}
	resBody, err := json.Marshal(res)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
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

	err = h.authUseCase.CheckMail(req)
	if err != nil {
		_ = response.ReturnErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	jwt, err := auth.IssueUserToken(req.UserID)
	if err != nil {
		_ = response.ReturnErrorResponse(w, http.StatusInternalServerError, "bad request")
		return
	}

	cookie := &http.Cookie{
		Name:     "token",
		Value:    jwt,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
	}
	http.SetCookie(w, cookie)

	res := response.Token{
		Token: jwt,
	}
	resBody, err := json.Marshal(res)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", strconv.Itoa(len(resBody)))
	w.WriteHeader(http.StatusOK)
	w.Write(resBody)
}

func (h *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
	// Cookieを削除
	cookie := &http.Cookie{
		Name:     "token",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		MaxAge:   -1,
	}
	http.SetCookie(w, cookie)

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
