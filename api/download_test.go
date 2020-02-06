package api

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDownloadHandler_noVariant(t *testing.T) {
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/api/dwonload?version=0.5.3&platform=linux/amd64", nil)

	DownloadHandler(w, req)

	resp := w.Result()
	assert.Equal(t, http.StatusTemporaryRedirect, resp.StatusCode)
	assert.Equal(t, "https://github.com/itzg/easy-add/releases/download/0.5.3/easy-add_linux_amd64", resp.Header.Get("Location"))
}

func TestDownloadHandler_withVariant(t *testing.T) {
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/api/dwonload?version=0.5.3&platform=linux/arm/v7", nil)

	DownloadHandler(w, req)

	resp := w.Result()
	assert.Equal(t, http.StatusTemporaryRedirect, resp.StatusCode)
	assert.Equal(t, "https://github.com/itzg/easy-add/releases/download/0.5.3/easy-add_linux_armv7", resp.Header.Get("Location"))
}
