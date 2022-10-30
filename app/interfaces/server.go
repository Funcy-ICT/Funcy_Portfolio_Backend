package interfaces

import (
	"backend/app/configs"
	"backend/app/infrastructure"
	"backend/app/interfaces/handler"
	"backend/app/usecase"
	"fmt"
	"github.com/jmoiron/sqlx"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
)

type Server struct {
	Router *chi.Mux
	db     *sqlx.DB
}

func NewServer() *Server {
	return &Server{
		Router: chi.NewRouter(),
	}
}

func (s *Server) Init() error {
	conn, err := configs.Init()
	if err != nil {
		return fmt.Errorf("failed db init. %s", err)
	}
	s.db = conn
	s.Route()
	return nil
}

func (s *Server) Route() {

	s.Router.Use(cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "HEAD", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
	}).Handler)

	//TODO　いずれちゃんとしたものに置き換えます
	s.Router.Use(middleware.Logger)
	//接続確認
	s.Router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})

	//DI
	authRepository := infrastructure.NewUserRepository(s.db)
	authUseCase := usecase.NewAuthUseCase(authRepository)
	authHandler := handler.NewAuthHandler(authUseCase)

	//認証
	s.Router.Post("/sign/up", authHandler.SignUp)
	s.Router.Post("/login", authHandler.SignIn)

	//Server.POST("/sign/up", handler.SignUp())
	//Server.POST("/sign/in", handler.SignIn())
	//
	////ユーザ関連
	//
	////作品関連
	//work := Server.Group("/work")
	//{
	//	work.POST("", middleware.Authenticate(handler.CreateWork()))
	//	work.GET("/:id", middleware.Authenticate(handler.ReadWork()))
	//}
	//works := Server.Group("/works")
	//{
	//	works.GET("/:number", middleware.Authenticate(handler.ReadWorksList()))
	//}

	//グループ関連

	//検索関連

}
