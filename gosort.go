package gosort

func BubbleSort(s []int) {
	if len(s) <= 1 {
		return
	}
	var swap bool
	for !swap {
		swap = true
		for i := 0; i < len(s)-1; i++ {
			if s[i] > s[i+1] {
				s[i], s[i+1] = s[i+1], s[i]
				swap = false
			}
		}
	}
}
