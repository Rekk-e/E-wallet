package utils

// Функция для проверки нахождения строки в массиве строк
func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
