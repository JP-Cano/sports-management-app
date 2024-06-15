package main

import (
	"github.com/JP-Cano/sports-management-app/src/config"
	"github.com/JP-Cano/sports-management-app/src/infrastructure/database"
	"github.com/JP-Cano/sports-management-app/src/routes"
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

	return &Server{
		DB:     db,
		Router: router,
	}, nil
}

func (s *Server) Run(addr string) {
	registerRoutes(s.Router, s.DB)
	err := s.Router.Run(addr)
	if err != nil {
		log.Fatal("Error starting server: ", err.Error())
	}
}

func registerRoutes(r *gin.Engine, db *gorm.DB) {
	routes.SetUpHealthCheck(r, db)
}
