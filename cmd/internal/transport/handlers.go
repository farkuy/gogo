package transport

import (
	"encoding/json"
	"net/http" 

	"github.com/farkuy/gogo/cmd/internal/models/user"
)

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("id")
	if userID == "" {
		http.Error(w, "Missing id parameter", http.StatusBadRequest)
		return
	}

	//TODO убрать ублюдство
	testUser := User{
		Id 1
		Email "first@gmail.com"
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
