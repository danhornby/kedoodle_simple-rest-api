package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetRootRoute(t *testing.T) {
	w := httptest.NewRecorder()
	_, r := gin.CreateTestContext(w)
	setupRouter(r)

	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "Hello World", w.Body.String())
}
