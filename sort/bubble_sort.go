package sort

func BubbleSort(a []int)  {
	for j := len(a) - 1; j > 0; j-- {
		for i := 0; i <j; i++{
			if a[i] > a[i+1] {
				temp := a[i]
				a[i] = a[i+1]
				a[i+1] = temp
			}
		}
	}
}
