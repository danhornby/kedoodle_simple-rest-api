package main

import (
	"fmt"
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

func TestGetStatusRoute(t *testing.T) {
	w := httptest.NewRecorder()
	_, r := gin.CreateTestContext(w)
	setupRouter(r)

	req, _ := http.NewRequest(http.MethodGet, "/status", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Header().Get("Content-Type"), "application/json")

	expectedVersion, expectedDescription, expectedCommit := "No version provided", "No description provided", "No commit provided"
	expected := fmt.Sprintf(`{
		"my-application": [
			{
				"version": "%s",
				"description": "%s",
				"sha": "%s"
			}
		]
	}`, expectedVersion, expectedDescription, expectedCommit)
	assert.JSONEq(t, expected, w.Body.String())
}

func TestMetadataEmbed(t *testing.T) {
	assert.NotEmpty(t, metadataBytes)
}

func TestSetMetadata(t *testing.T) {
	assert.Equal(t, "No version provided", Version)
	assert.Equal(t, "No description provided", Description)

	expectedVersion := "test version"
	expectedDescription := "test description"
	m := []byte(fmt.Sprintf(`{
		"version": "%s",
		"description": "%s"
	}`, expectedVersion, expectedDescription))
	setMetadata(m)

	assert.Equal(t, expectedVersion, Version)
	assert.Equal(t, expectedDescription, Description)
}
