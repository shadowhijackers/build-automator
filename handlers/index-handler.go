package handlers

import (
	"fmt"
	"net/http"
)

func Index(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "%s", "YOU HAVE BEING HACKED")
}
