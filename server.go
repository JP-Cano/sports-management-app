package main

import (
	"github.com/JP-Cano/sports-management-app/src/adapters/controllers"
	"github.com/JP-Cano/sports-management-app/src/adapters/repositories"
	"github.com/JP-Cano/sports-management-app/src/config"
	"github.com/JP-Cano/sports-management-app/src/infrastructure/database"
	"github.com/JP-Cano/sports-management-app/src/routes"
	"github.com/JP-Cano/sports-management-app/src/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
)

type Server struct {
	DB     *gorm.DB
	Router *gin.Engine
}

func NewServer() (*Server, error) {
	environment := config.Env().AppEnv
	if environment == "" {
		environment = "local"
	}

	config.Load(environment)
	log.Printf("Environment: %s", environment)

	dsn := config.GetDSN()
	db, err := database.New(dsn)
	if err != nil {
		return nil, err
	}

	router := gin.Default()
	router.Use(gin.Recovery())

	registerRoutes(router, db)
	return &Server{
		DB:     db,
		Router: router,
	}, nil
}

func (s *Server) Run(addr string) {
	err := s.Router.Run(addr)
	if err != nil {
		log.Fatal("Error starting server: ", err.Error())
	}
}

func registerRoutes(r *gin.Engine, db *gorm.DB) {
	userRepository := initializeRepositories(db)
	userService := initializeServices(userRepository)
	userController := initializeControllers(userService)
	routes.SetUpHealthCheck(r, db)
	routes.SetUpUser(r, userController)
}

func initializeRepositories(db *gorm.DB) (userRepository *repositories.User) {
	userRepository = repositories.NewUserRepository(db)
	return
}

func initializeServices(userRepository *repositories.User) (userService services.UserService) {
	userService = services.NewUserService(userRepository)
	return
}

func initializeControllers(userService services.UserService) (userController *controllers.UserController) {
	userController = controllers.NewUserController(userService)
	return
}
