package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"go-chat/storage"
)

func GetMessages(w http.ResponseWriter, r *http.Request) {
	// Parse query parameters
	groupIDStr := r.URL.Query().Get("group_id")
	limitStr := r.URL.Query().Get("limit")
	offsetStr := r.URL.Query().Get("offset")

	// Convert string to int
	groupID, err := strconv.Atoi(groupIDStr)
	if err != nil || groupID <= 0 {
		http.Error(w, "Invalid or missing group_id", http.StatusBadRequest)
		return
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		limit = 50 // default
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil || offset < 0 {
		offset = 0
	}

	// Query messages from storage
	messages, err := storage.GetMessagesByGroup(groupID, limit, offset)
	if err != nil {
		http.Error(w, "Failed to fetch messages", http.StatusInternalServerError)
		return
	}

	// Respond as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(messages)
}
