package server

import (
	"testing"

	"github.com/gin-gonic/gin"
)

// route handlers are tested separately in handlers_test.go
// and therefore, this file is meant for route registration testing.

func TestRoutesExist(t *testing.T) {

	testApp := Config{}

	// http.Handler and gin.Engine interfaces are compatible
	// basically, type returns by testApp.routes() is global/default interface which is http.Handler
	// it needs to be more specific to Gin in this case, so we do type assertion.
	testRoutes := testApp.routes()

	ginRouter := testRoutes.(*gin.Engine)

	routes := []string{"/login", "/register"}

	for _, route := range routes {
		routeExists(t, ginRouter, route)
	}

}

func routeExists(t *testing.T, router *gin.Engine, route string) {
	found := false

	for _, registeredRoute := range router.Routes() {
		if registeredRoute.Path == route {
			found = true
			break
		}
	}

	if !found {
		t.Errorf("did not find %s in registered routes", route)
	}
}
