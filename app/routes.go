package app

import "github.com/teukuhirzi/altastore/app/controllers"

func (server *Server) initializeRoutes() {
	server.Router.HandleFunc("/", controllers.Home).Methods("GET")
	//server.Router.HandleFunc()
}
