package user_handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	user "github.com/farkuy/gogo/cmd/internal/models"
)

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("id")
	if userID == "" {
		http.Error(w, "Missing id parameter", http.StatusBadRequest)
		return
	}

	testUser := user.User{Id: 1, Email: "first@mail.ru"}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(testUser)
	fmt.Fprintf(w, "Welcome to user!")
}
