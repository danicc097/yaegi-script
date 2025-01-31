package internal

import (
	"reflect"
)

//go:generate yaegi extract github.com/bitfield/script
var Symbols map[string]map[string]reflect.Value = make(map[string]map[string]reflect.Value)
