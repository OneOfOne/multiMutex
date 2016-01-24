package multiMutex

func modDjb2(s string) (h int) {
	h = 5381
	for i := 0; i < len(s); i++ {
		h = h*33 + int(s[i])
	}
	if h < 0 {
		h *= -1
	}
	return
}
