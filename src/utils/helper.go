package utils

import (
	"net/http/httptest"
	"os"

	"github.com/gin-gonic/gin"
)

// OptionalValue is used when optional needed
type OptionalValue struct {
	Value        string
	DefaultValue string
}

// GetDotEnvVariable ...
func GetDotEnvVariable(key OptionalValue) string {
	v := os.Getenv(key.Value)
	if v == "" {
		v = key.DefaultValue
	}
	return v
}

// Cors function cors
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.Next()
	}
}

// DoURLTest ...
func DoURLTest(r *gin.Engine, method string, route string) *httptest.ResponseRecorder {
	req := httptest.NewRequest(method, route, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}
