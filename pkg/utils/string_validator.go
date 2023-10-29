package util

import "reflect"

func IsValidStringWithLength(s string, length int) bool {
	var (
		isValidCount = false
		isNotEmpty   = false
	)

	if len(s) >= length {
		isValidCount = true
	}

	if s != "" {
		isNotEmpty = true
	}

	return isValidCount && isNotEmpty
}

func IsValidArrayWithLength(s []string, length int) bool {
	var (
		isValidCount = false
		isNotEmpty   = false
		isValidArray = false
	)

	if t := reflect.TypeOf(s); t.Kind() != reflect.Array {
		isValidArray = true
	}

	if len(s) >= length {
		isValidCount = true
	}

	if len(s) < 1 {
		isNotEmpty = true
	}

	return isValidCount && isNotEmpty && isValidArray
}
