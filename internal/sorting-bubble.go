package internal

func BubbleSort(a []int) []int {

	// get the first index element and compare with next element
	// compare them; if first greater then second swap elements

	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a)-1; j++ {
			if a[j] > a[j+1] {
				temp := a[j]
				a[j] = a[j+1]
				a[j+1] = temp
			}
		}
	}

	return a
}
