package interfaces

import (
	"backend/app/configs"
	"backend/app/domain/entity"
	"backend/app/infrastructure"
	"backend/app/interfaces/handler"
	middleware2 "backend/app/interfaces/middleware"
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
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "HEAD", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
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
	s.Router.Post("/auth/code", authHandler.AuthCode)

	// auth
	s.Router.Group(func(mux chi.Router) {
		mux.Use(middleware2.Authentication)
		mux.Get("/health/jwt", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("ok"))
		})
		mux.Post("/work", workHandler.CreateWork)
		mux.Delete("/work/{workID}", workHandler.DeleteWork)

		mux.Get("/userinfo/{userID}", userinfoHandler.GetUserinfo)
	})

	s.Router.Post("/test", func(w http.ResponseWriter, r *http.Request) {
		i := infrastructure.NewUserInfoRepository(s.db)
		err := i.CreateNewUserinfo(&entity.Userinfo{
			Profile: &entity.Profile{
				UserID:          "7a27c640-071a-4703-a7db-cab765b3f2f1",
				HeaderImagePath: "https://example.com/hoge.png",
				Biography:       "git push origin master",
			},
			JoinedGroups: &[]*entity.GroupMember{
				{
					GroupID: "7516011c-ac6c-7889-bfd3-656cbae2f4be",
					UserID:  "7a27c640-071a-4703-a7db-cab765b3f2f1",
					Role:    "member",
					Status:  true,
				},
			},
			Skills: &[]*entity.Skill{
				{
					SkillName: "絶起",
					UserID:    "7a27c640-071a-4703-a7db-cab765b3f2f1",
				},
			},
			SNS: &[]*entity.SNS{
				{
					SnsURL: "https://twitter.com/jugesuke",
					UserID: "7a27c640-071a-4703-a7db-cab765b3f2f1",
				},
			},
		})
		if err != nil {
			w.Write([]byte(err.Error()))
		}
	})

	// no auth
	s.Router.Get("/work/{workID}", workHandler.ReadWork)
	s.Router.Get("/works/{number}", workHandler.ReadWorks)

}
