package main

import "fmt"

// BubbleSort 冒泡排序
func BubbleSort(list []int) {
	n := len(list)

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if list[i] > list[j] {
				list[i], list[j] = list[j], list[i]
			}
		}
	}
	return
}

// SelectSort 选择排序
func SelectSort(list []int) {
	n := len(list)

	for i := 0; i < n-1; i++ {
		minIndex := i

		for j := i + 1; j < n; j++ {
			if list[minIndex] > list[j] {
				minIndex = j
			}
		}

		if minIndex != i {
			list[minIndex], list[i] = list[i], list[minIndex]
		}
	}
}

// InsertSort 插入排序
func InsertSort(list []int) {
	n := len(list)

	for i := 1; i < n; i++ {
		deal := list[i]
		j := i - 1

		if deal < list[j] {
			for ; j >= 0 && deal < list[j]; j-- {
				list[j+1] = list[j]
			}
			list[j+1] = deal
		}
	}
}

// ShellSort 希尔排序
func ShellSort(list []int) {
	n := len(list)

	// 每次减半，直到步长为 1
	for step := n / 2; step >= 1; step /= 2 {
		// 开始插入排序，每一轮的步长为 step
		for i := step; i < n; i += step {
			for j := i - step; j >= 0; j -= step {
				// 满足插入那么交换元素
				if list[j+step] < list[j] {
					list[j], list[j+step] = list[j+step], list[j]
					continue
				}
				break
			}
		}
	}
}

// MergeSort 归并排序
func MergeSort(array []int, begin int, end int) {
	// 元素数量大于1时才进入递归
	if end-begin > 1 {

		// 将数组一分为二，分为 array[begin,mid) 和 array[mid,high)
		mid := begin + (end-begin+1)/2

		// 先将左边排序好
		MergeSort(array, begin, mid)

		// 再将右边排序好
		MergeSort(array, mid, end)

		// 两个有序数组进行合并
		merge(array, begin, mid, end)
	}
}

// 归并操作
func merge(array []int, begin int, mid int, end int) {
	// 申请额外的空间来合并两个有序数组，这两个数组是 array[begin,mid),array[mid,end)
	leftSize := mid - begin         // 左边数组的长度
	rightSize := end - mid          // 右边数组的长度
	newSize := leftSize + rightSize // 辅助数组的长度
	result := make([]int, 0, newSize)

	l, r := 0, 0
	for l < leftSize && r < rightSize {
		lValue := array[begin+l] // 左边数组的元素
		rValue := array[mid+r]   // 右边数组的元素
		// 小的元素先放进辅助数组里
		if lValue < rValue {
			result = append(result, lValue)
			l++
		} else {
			result = append(result, rValue)
			r++
		}
	}

	// 将剩下的元素追加到辅助数组后面
	result = append(result, array[begin+l:mid]...)
	result = append(result, array[mid+r:end]...)

	// 将辅助数组的元素复制回原数组，这样该辅助空间就可以被释放掉
	for i := 0; i < newSize; i++ {
		array[begin+i] = result[i]
	}
	return
}

// QuickSort 快速排序
func QuickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	splitData := arr[0]          //第一个数据
	low := make([]int, 0, 0)     //比我小的数据
	high := make([]int, 0, 0)    //比我大的数据
	mid := make([]int, 0, 0)     //与我一样大的数据
	mid = append(mid, splitData) //加入一个
	for i := 1; i < len(arr); i++ {
		if arr[i] < splitData {
			low = append(low, arr[i])
		} else if arr[i] > splitData {
			high = append(high, arr[i])
		} else {
			mid = append(mid, arr[i])
		}
	}
	low, high = QuickSort(low), QuickSort(high)
	newArr := append(append(low, mid...), high...)
	return newArr
}

// 冒泡排序、选择排序、插入排序、希尔排序、归并排序、堆排序、快速排序

func main() {
	list := []int{2, 1, 5, 4, 9, 3, 8, 6, 7}
	//BubbleSort(list)
	//SelectSort(list)
	//InsertSort(list)
	//ShellSort(list)
	//MergeSort(list, 0, len(list))
	list = QuickSort(list)
	fmt.Println(list)
}
