package server

import (
	"log"

	"gihub.com/Emanoel01/book-rest-api/src/server/routes"
	"github.com/gin-gonic/gin"
)

type Server struct {
	port   string
	server *gin.Engine
}

func NewServer() Server {
	return Server{
		port: "1337",
		server: gin.Default(),
	}
}


func (s *Server) Run(){
	router := routes.ConfigRoutes(s.server)

	log.Print("Server is running at port: ", s.port)
	log.Fatal(router.Run(":" + s.port))
}