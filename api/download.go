package api

import (
	"fmt"
	"net/http"
	"strings"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	version := r.URL.Query().Get("version")
	if version == "" {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("Missing version parameter"))
		return
	}

	platform := r.URL.Query().Get("platform")
	if platform == "" {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("Missing platform parameter"))
		return
	}

	suffix := strings.ReplaceAll(platform, "/", "_")

	w.Header().Set("Location", fmt.Sprintf("https://github.com/itzg/easy-add/releases/download/%s/easy-add_%s", version, suffix))
	w.WriteHeader(http.StatusTemporaryRedirect)
}
