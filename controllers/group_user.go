package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"go-chat/storage"
)

func ListGroups(w http.ResponseWriter, r *http.Request) {
	groups, err := storage.GetAllGroups()
	if err != nil {
		http.Error(w, "Failed to fetch groups", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(groups)
}