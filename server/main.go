package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/febzey/port/pkg/configs"
	"github.com/febzey/port/pkg/routes"
	"github.com/febzey/port/pkg/utils"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func GetCurrentDirectory() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	return dir, nil
}

func FileServer(router *mux.Router) {
	dir, _ := GetCurrentDirectory()
	webPath := filepath.Join(dir, "../web/dist")
	router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir(webPath))))
}

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	LoadEnv()
	router := mux.NewRouter()
	routes.Public(router)
	router.Use(mux.CORSMethodMiddleware(router))
	FileServer(router)
	server := configs.ServerConfig(router)
	utils.StartServer(server)
}
