package utils

import (
	"fmt"
	"os"
)

func UrlBuilder(n string) (string, error) {
	var url string

	switch n {
	case "server":
		url = fmt.Sprintf("%s:%s", os.Getenv("SERVER_HOST"), os.Getenv("SERVER_PORT"))
	default:
		return "", fmt.Errorf("urlBuilder: %s is not a valid url", n)
	}

	return url, nil

}
