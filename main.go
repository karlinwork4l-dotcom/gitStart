package main

import (
	"encoding/json"
	"net/http"
	"time"

	"modernc.org/libc/uuid"
)

func main() {
	http.HandleFunc("/links", linksHandler)
	http.HandleFunc("/links/", linksHandler)

	http.ListenAndServe(":8080", nil)
}

func linksHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		var input Link

		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		newLink := Link{
			ID:          uuid.New().String(),
			URL:         input.URL,
			Title:       input.Title,
			Description: input.Description,
			CreatedAt:   time.Now(),
		}

		links = append(links, newLink)

		w.Header().Set("Context-Type", "application/json")
		json.NewEncoder(w).Encode(newLink)

	case http.MethodGet:
		w.Header().Set("Context-Type", "application/json")
		json.NewEncoder(w).Encode(links)

	default:
		http.Error(w, "method not allowed", http.StatusBadRequest)
	}

}
