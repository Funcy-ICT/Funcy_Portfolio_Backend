package middleware

import (
	"context"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	"backend/app/domain/entity"

	jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/jwks"
	"github.com/auth0/go-jwt-middleware/v2/validator"
)

// CustomClaims
type CustomClaims struct {
	Scope   string `json:"scope"`
	Email   string `json:"email"`
	Name    string `json:"name"`
	Picture string `json:"picture"`
}

func (c *CustomClaims) Validate(ctx context.Context) error {
	return nil
}

// EnsureValidToken
func EnsureValidToken() func(next http.Handler) http.Handler {
	issuerURL, err := url.Parse("https://" + os.Getenv("AUTH0_DOMAIN") + "/")
	if err != nil {
		log.Fatalf("Failed to parse the issuer url: %v", err)
	}

	provider := jwks.NewCachingProvider(issuerURL, 5*time.Minute)

	jwtValidator, err := validator.New(
		provider.KeyFunc,
		validator.RS256,
		issuerURL.String(),
		[]string{
			os.Getenv("AUTH0_AUDIENCE"),
			issuerURL.String() + "userinfo",
		},
		validator.WithCustomClaims(func() validator.CustomClaims {
			return &CustomClaims{}
		}),
		validator.WithAllowedClockSkew(time.Minute),
	)
	if err != nil {
		log.Fatalf("Failed to set up the jwt validator: %v", err)
	}

	errorHandler := func(w http.ResponseWriter, r *http.Request, err error) {
		log.Printf("Encountered error while validating JWT: %v", err)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`{"message":"Failed to validate JWT."}`))
	}

	middleware := jwtmiddleware.New(
		jwtValidator.ValidateToken,
		jwtmiddleware.WithErrorHandler(errorHandler),
	)

	return func(next http.Handler) http.Handler {
		return middleware.CheckJWT(next)
	}
}

func GetUserSubFromContext(ctx context.Context) string {
	claims, ok := ctx.Value(jwtmiddleware.ContextKey{}).(*validator.ValidatedClaims)
	if !ok {
		return ""
	}
	return claims.RegisteredClaims.Subject
}

func SetUserID(authRepo interface {
	GetByAuth0Sub(string) (*entity.User, error)
}) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			sub := GetUserSubFromContext(ctx)
			if sub != "" {
				user, err := authRepo.GetByAuth0Sub(sub)
				if err == nil && user != nil {
					ctx = context.WithValue(ctx, "user_id", user.UserID)
				}
			}
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
