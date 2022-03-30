package configs

import (
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/febzey/port/pkg/utils"
	"github.com/gorilla/mux"
)

func ServerConfig(router *mux.Router) *http.Server {
	serverConnURL, _ := utils.UrlBuilder("server")
	readTimeoutSeconds, _ := strconv.Atoi(os.Getenv("SERVER_READ_TIMEOUT"))

	return &http.Server{
		Handler:     router,
		Addr:        serverConnURL,
		ReadTimeout: time.Second * time.Duration(readTimeoutSeconds),
	}
}
