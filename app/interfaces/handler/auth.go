package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"

	"backend/app/interfaces/middleware"
	"backend/app/interfaces/request"
	"backend/app/interfaces/response"
	"backend/app/packages/utils"
	"backend/app/packages/utils/auth"
	"backend/app/usecase"

	jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/validator"
)

func isProduction() bool {
	env := os.Getenv("ENVIRONMENT")
	return env == "production"
}

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

	user, token, err := h.authUseCase.Login(req)
	if err != nil {
		log.Printf("SignIn failed: %v", err)
		_ = response.ReturnErrorResponse(w, http.StatusBadRequest, "An unexpected error occurred. Please try again later.")
		return
	}

	cookie := &http.Cookie{
		Name:     "token",
		Value:    token,
		Path:     "/",
		HttpOnly: true,
		Secure:   isProduction(),
		SameSite: http.SameSiteLaxMode,
	}
	if isProduction() {
		cookie.SameSite = http.SameSiteNoneMode
	}
	http.SetCookie(w, cookie)

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
		Name:     "token",
		Value:    jwt,
		Path:     "/",
		HttpOnly: true,
		Secure:   isProduction(),
		SameSite: http.SameSiteLaxMode,
	}
	if isProduction() {
		cookie.SameSite = http.SameSiteNoneMode
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

func (h *AuthHandler) SyncUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	sub := middleware.GetUserSubFromContext(ctx)
	if sub == "" {
		_ = response.ReturnErrorResponse(w, http.StatusUnauthorized, "No valid token found")
		return
	}

	type SyncUserRequest struct {
		Email   string `json:"email"`
		Name    string `json:"name"`
		Picture string `json:"picture"`
	}
	var req SyncUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("Failed to decode request body: %v", err)
		claims, ok := ctx.Value(jwtmiddleware.ContextKey{}).(*validator.ValidatedClaims)
		if ok {
			customClaims, ok := claims.CustomClaims.(*middleware.CustomClaims)
			if ok {
				req.Email = customClaims.Email
				req.Name = customClaims.Name
				req.Picture = customClaims.Picture
			}
		}
	}

	userID, err := h.authUseCase.SyncUser(sub, req.Email, req.Name, req.Picture)
	if err != nil {
		log.Printf("SyncUser failed: %v", err)
		_ = response.ReturnErrorResponse(w, http.StatusInternalServerError, "Failed to sync user")
		return
	}

	res := response.UserID{
		UserID: userID,
	}
	resBody, err := json.Marshal(res)
	if err != nil {
		log.Printf("SyncUser failed: %v", err)
		_ = response.ReturnErrorResponse(w, http.StatusInternalServerError, "Failed to marshal response")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", strconv.Itoa(len(resBody)))
	w.WriteHeader(http.StatusOK)
	w.Write(resBody)
}
