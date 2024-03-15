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

func InsertionSort(s []int) {
	if len(s) <= 1 {
		return
	}
	for i := 1; i < len(s); i++ {
		for j := i - 1; j >= 0; j-- {
			if s[j+1] < s[j] {
				s[j], s[j+1] = s[j+1], s[j]
			} else {
				break
			}
		}
	}
}

func QuickSort(s []int) {
	return
}
