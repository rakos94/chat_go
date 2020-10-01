package setup

import "github.com/gin-gonic/gin"

// IGinMode ...
type IGinMode interface {
	SetNoMode()
	SetReleaseMode()
	SetTestMode()
}

// GinMode ...
type ginMode struct{}

// NewGinMode ...
func NewGinMode() IGinMode {
	return &ginMode{}
}

// SetNoMode ...
func (a ginMode) SetNoMode() {}

// SetReleaseMode ...
func (a ginMode) SetReleaseMode() {
	gin.SetMode(gin.ReleaseMode)
}

// SetTestMode ...
func (a ginMode) SetTestMode() {
	gin.SetMode(gin.TestMode)
}
