package utils

func UseString(s *string) (value string) {
	if s == nil {
		value = ""
		return
	}
	value = *s
	return
}
