package setup

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnectDB(t *testing.T) {
	asserts := assert.New(t)
	LoadEnvTest()
	_, err := InitDatabase()
	asserts.Empty(err, "Db should exist")
}
