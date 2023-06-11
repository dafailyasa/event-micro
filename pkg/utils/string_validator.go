package util

func IsValidStringWithLength(s string, length int) bool {
	var (
		isValidCount = false
		isEmpty      = false
	)

	if len(s) >= length {
		isValidCount = true
	}

	if s != "" {
		isEmpty = true
	}

	return isValidCount && isEmpty
}
