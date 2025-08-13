package transport

import (
	"net/http"

	user_handler "github.com/farkuy/gogo/cmd/internal/transport/user_handler"
)

func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/user/get", user_handler.GetUserHandler)
}
