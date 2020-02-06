package api

import (
	"fmt"
	"net/http"
	"strings"
)

func DownloadHandler(w http.ResponseWriter, r *http.Request) {
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

	parts := strings.SplitN(platform, "/", 3)
	os := parts[0]
	arch := parts[1]
	variant := ""
	if len(parts) == 3 {
		variant = parts[2]
	}

	w.Header().Set("Location", fmt.Sprintf("https://github.com/itzg/easy-add/releases/download/%s/easy-add_%s_%s%s", version, os, arch, variant))
	w.WriteHeader(http.StatusTemporaryRedirect)
}
