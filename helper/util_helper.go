package helper

func Contains(list []string, value any) bool {
	for _, data := range list {
		if data == value {
			return true
		}
	}
	return false
}
