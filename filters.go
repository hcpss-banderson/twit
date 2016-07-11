package main

// ToBool is a template filter that exports a boolean value as a string.
func ToBool(value bool) string {
	s := "false"

	if value {
		return "true"
	}

	return s
}
