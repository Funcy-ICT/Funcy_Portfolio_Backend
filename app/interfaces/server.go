package interfaces

import (
	"backend/app/configs"
	"backend/app/infrastructure"
	"backend/app/interfaces/handler"
	middleware2 "backend/app/interfaces/middleware"
	"backend/app/packages/utils/auth"
	"backend/app/usecase"

	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/jmoiron/sqlx"
)

type processor interface {
	doSth() (error, string)
}

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
		AllowedOrigins:   []string{"http://*", "https://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "HEAD", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
	}).Handler)

	//DI
	authRepository := infrastructure.NewUserRepository(s.db)
	authUseCase := usecase.NewAuthUseCase(authRepository)
	authHandler := handler.NewAuthHandler(authUseCase)

	workRepository := infrastructure.NewWorkRepository(s.db)
	workUseCase := usecase.NewWorkUseCase(workRepository)
	workHandler := handler.NewWorkHandler(workUseCase)

	userinfoRepository := infrastructure.NewUserInfoRepository(s.db)
	userinfoUseCase := usecase.NewUserinfoUsecace(userinfoRepository, workRepository)
	userinfoHandler := handler.NewUserinfoHandler(userinfoUseCase)

	commentRepository := infrastructure.NewCommentRepository(s.db)
	commentUseCase := usecase.NewCommentUsecace(commentRepository)
	commentHandler := handler.NewCommentHandler(commentUseCase)

	groupRepository := infrastructure.NewGroupRepository(s.db)
	groupUseCase := usecase.NewGroupUseCase(groupRepository)
	groupHandler := handler.NewGroupHandler(groupUseCase)

	s.Router.Use(middleware.Logger)
	//接続確認
	s.Router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})
	s.Router.Post("/sign/up", authHandler.SignUp)
	s.Router.Post("/login", authHandler.SignIn)
	//応急処置
	s.Router.Post("/mlogin", authHandler.SignInMobile)
	//アカウント認証
	s.Router.Post("/authcode", authHandler.AuthCode)

	// auth
	s.Router.Group(func(mux chi.Router) {
		mux.Use(middleware2.Authentication)
		mux.Get("/health/jwt", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("ok"))
		})

		// Work関連のエンドポイント
		mux.Route("/work", func(r chi.Router) {
			r.Post("/", workHandler.CreateWork)
			r.Delete("/{workID}", workHandler.DeleteWork)
			r.Put("/{workID}", workHandler.UpdateWork)
		})

		// ユーザー別作品取得エンドポイント
		mux.Get("/user/works", workHandler.ReadWorksByUser)

		// ユーザー情報関連のエンドポイント
		mux.Route("/userinfo", func(r chi.Router) {
			r.Get("/{userID}", userinfoHandler.GetUserinfo)
			r.Put("/{userID}", userinfoHandler.PutUserinfo)
		})

		// グループ関連のエンドポイント
		mux.Post("/group", groupHandler.CreateGroup)

		// コメント関連のエンドポイント
		mux.Post("/comment", commentHandler.CreateComment)
	})

	// コメント関連のエンドポイント
	s.Router.Get("/comment/{worksID}", commentHandler.GetComment)

	// no auth
	s.Router.Get("/work/{workID}", workHandler.ReadWork)
	s.Router.Get("/works/{number}", workHandler.ReadWorks)
	s.Router.Get("/user/{userID}/works", workHandler.ReadWorksByUserID)

}

func issueSuperAccountToken(userID string) (string, error) {
	return auth.IssueSuperUserToken(userID)
}
