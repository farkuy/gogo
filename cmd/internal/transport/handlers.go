package transport

import (
	"encoding/json"
	"net/http"

	user "github.com/farkuy/gogo/cmd/internal/models"
)

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("id")
	if userID == "" {
		http.Error(w, "Missing id parameter", http.StatusBadRequest)
		return
	}

	//TODO убрать ублюдство
	testUser := user.User{Id: 1, Email: "first@mail.ru"}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(testUser)
}
