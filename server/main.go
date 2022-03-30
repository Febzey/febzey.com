package main

import (
	"fmt"

	"github.com/febzey/port/pkg/routes"
	"github.com/febzey/port/pkg/utils"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	"github.com/febzey/port/pkg/configs"
)

//create main function
func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}

	router := mux.NewRouter()
	routes.Public(router)

	router.Use(mux.CORSMethodMiddleware(router))

	server := configs.ServerConfig(router)

	utils.StartServer(server)

}
