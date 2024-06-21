package main

import (
	"github.com/JP-Cano/sports-management-app/src/adapters/controllers"
	"github.com/JP-Cano/sports-management-app/src/adapters/repositories"
	routes2 "github.com/JP-Cano/sports-management-app/src/adapters/routes"
	"github.com/JP-Cano/sports-management-app/src/application/config"
	"github.com/JP-Cano/sports-management-app/src/application/services"
	"github.com/JP-Cano/sports-management-app/src/infrastructure/database"
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
	routes2.SetUpHealthCheck(r, db)
	routes2.SetUpUser(r, userController)
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
