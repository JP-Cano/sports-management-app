package main

import (
	"github.com/JP-Cano/sports-management-app/src/adapters/controllers"
	"github.com/JP-Cano/sports-management-app/src/adapters/repositories"
	"github.com/JP-Cano/sports-management-app/src/adapters/routes"
	"github.com/JP-Cano/sports-management-app/src/application/config"
	"github.com/JP-Cano/sports-management-app/src/application/services"
	"github.com/JP-Cano/sports-management-app/src/application/services/file"
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
	userRepository, playerRepository := initializeRepositories(db)
	userService, _, fileService := initializeServices(userRepository, playerRepository)
	userController, fileController := initializeControllers(userService, fileService)

	routes.SetUpHealthCheck(r, db)
	routes.SetUpUser(r, userController)
	routes.SetUpFile(r, fileController)
}

func initializeRepositories(db *gorm.DB) (userRepository *repositories.User, playerRepository *repositories.Player) {
	userRepository = repositories.NewUserRepository(db)
	playerRepository = repositories.NewPlayerRepository(db)
	return
}

func initializeServices(userRepository *repositories.User, playerRepository *repositories.Player) (userService services.UserService, playerService services.PlayerService, fileService file.ExcelService) {
	userService = services.NewUserService(userRepository)
	playerService = services.NewPlayerService(playerRepository)
	fileService = file.NewPlayerService(playerService)
	return
}

func initializeControllers(userService services.UserService, fileService file.ExcelService) (userController *controllers.UserController, fileController *controllers.FileController) {
	userController = controllers.NewUserController(userService)
	fileController = controllers.NewFileController(fileService)
	return
}
