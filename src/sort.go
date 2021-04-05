package leetcode

//bubble sort

func BubbleSort(arr []int) []int {
	length := len(arr)
	if length == 0 {
		return arr
	}
	for i := 0; i < length-1; i++ {
		for j := 1; j < length-i-1; j++ {
			if arr[j-1] > arr[j] {
				temp := arr[j]
				arr[j] = arr[j-1]
				arr[j-1] = temp
			}
		}
	}
	return arr
}

func QuickSort(arr []int) []int {
	length := len(arr)
	if length < 2 {
		return arr
	}
	head, end := 0, length-1
	pio := arr[head]
	for head < end {
		if arr[head+1] > pio {
			arr[head+1], arr[end] = arr[end], arr[head+1]
			end--
		} else {
			arr[head], arr[head+1] = arr[head+1], arr[head]
			head--
		}
	}
	QuickSort(arr[:head])
	QuickSort(arr[head+1:])
}
