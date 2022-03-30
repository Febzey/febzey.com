package configs

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/febzey/port/pkg/utils"
	"github.com/gorilla/mux"
)

func ServerConfig(router *mux.Router) *http.Server {
	serverConnURL, _ := utils.UrlBuilder("server")
	readTimeoutSeconds, _ := strconv.Atoi(os.Getenv("SERVER_READ_TIMEOUT"))

	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}

	//create a file server using the net/http package
	fileServer := http.FileServer(http.Dir(filepath.Join(dir, "../../web/dist")))

	fmt.Println(fileServer)

	//register the fileServer using the same serverConnURL
	router.PathPrefix("/").Handler(http.StripPrefix(serverConnURL, fileServer))

	return &http.Server{
		Handler:     router,
		Addr:        serverConnURL,
		ReadTimeout: time.Second * time.Duration(readTimeoutSeconds),
	}
}
