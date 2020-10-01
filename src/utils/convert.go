package utils

import (
	"encoding/json"
	"strconv"
)

// IMConvert interface
type IMConvert interface {
	Object(source interface{}, destination interface{}) error
	Uint(value string) (uint, error)
	// DatabaseError(err error) error
}

type imConvert struct{}

// NewIMConvert func
func NewIMConvert() IMConvert {
	return &imConvert{}
}

// Object func
func (c *imConvert) Object(source interface{}, destination interface{}) error {
	byteConvert, err := json.Marshal(source)
	if err != nil {
		return err
	}

	err = json.Unmarshal(byteConvert, &destination)
	return err
}

func (c *imConvert) Uint(value string) (uint, error) {
	pars, err := strconv.ParseUint(value, 10, 32)
	if err != nil {
		println(err.Error())
		return 0, err
	}
	return uint(pars), nil
}
