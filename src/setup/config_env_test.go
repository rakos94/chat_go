package setup

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"raks.com/kamal/practice_crud/src/utils"
)

func TestVar(t *testing.T) {
	LoadEnvTest()
	v := utils.GetDotEnvVariable(utils.OptionalValue{Value: "SERVE_PORT", DefaultValue: "8080"})
	assert.Equal(t, "9990", v)
}
