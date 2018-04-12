package jiup

func includes(arr []string, val string) bool {
	for i := range arr {
		if arr[i] == val {
			return true
		}
	}
	return false
}
