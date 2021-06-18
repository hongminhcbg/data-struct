package sort

func partition(a []int, low int, high int) int {
	i := low -1
	p := a[high]
	temp := 0

	for j := low; j < high; j++ {
		if a[j] < p {
			i++
			temp = a[i]
			a[i] = a[j]
			a[j] = temp
		}
	}

	temp = a[i+1]
	a[i+1] = a[high]
	a[high] = temp

	return i+1
}

func quickSort(a []int, low int, high int) {
	if high <= low {
		return
	}
	p := partition(a, low, high)

	quickSort(a, low, p-1)
	quickSort(a, p+1, high)
}

func QuickSort(a []int)  {
	quickSort(a, 0, len(a) - 1)
}