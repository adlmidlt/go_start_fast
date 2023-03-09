package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{2, 3, 4, 7, 11}
	fmt.Print(findKthPositive2(arr, 5))

}

// twoSum  Task 1.
// Способ 1.
// Учитывая массив целых чисел nums и целочисленное целевое значение, верните индексы двух чисел так, чтобы в сумме они составляли целевое значение.
// Вы можете предположить, что каждый вход будет иметь ровно одно решение, и вы не можете использовать один и тот же элемент дважды.
// Вы можете вернуть ответ в любом порядке.
func twoSum1(nums []int, target int) []int {
	arr := make([]int, 0)

	for i := range nums {
		for j := i + 1; j < len(nums); j++ {
			sum := nums[i] + nums[j]
			if sum == target {
				arr = append(arr, i, j)
			}
		}
	}

	return arr
}

func twoSum2(nums []int, target int) []int {
	numMap := make(map[int]int)

	for i, item := range nums {
		if j, ok := numMap[target-item]; ok {
			return []int{j, i}
		}

		numMap[item] = i
	}

	return []int{}
}

// findKthPositive Task 2.
// Дан массив arr целых положительных чисел, отсортированных в строго возрастающем порядке, и целое число k.
// Возвращает k-е положительное целое число, отсутствующее в этом массиве.
func findKthPositive1(arr []int, k int) int {
	arr2 := make([]int, 0)

	for i := 1; i <= 10000; i++ {
		arr2 = append(arr2, i)
	}

	numMap := make(map[int]int)

	for i, item := range arr2 {
		numMap[i] = item
	}

	for _, item := range arr {
		for k, v := range numMap {
			if item == v {
				delete(numMap, k)
				break
			}
		}
	}

	arr3 := make([]int, 0)

	for _, v := range numMap {
		arr3 = append(arr3, v)
	}
	sort.Ints(arr3)

	result := 0
	for i := range arr3 {
		if i == k {
			result = arr3[i-1]
		}
	}

	return result
}

/*
	findKthPositive Task 2.

*
Алгоритм работает путем перебора arr и подсчета количества отсутствующих целых чисел между соседними элементами.
Переменная missingCount отслеживает количество отсутствующих целых чисел до сих пор. Если missingCount больше или равно
k, то мы знаем, что k-ое отсутствующее целое число должно быть между предыдущим числом и текущим числом, поэтому
мы возвращаем текущее число минус количество отсутствующих чисел, которые идут после k-ого отсутствующего числа.
Если мы достигли конца arr и еще не нашли k-ое отсутствующее число, мы возвращаем последний элемент arr плюс k
минус количество отсутствующих чисел, которые мы уже нашли.
*/
func findKthPositive2(arr []int, k int) int {
	missingCount := 0 //
	prevNum := 0      //

	for _, num := range arr {
		missingCount += num - prevNum - 1

		if missingCount >= k {
			return num - (missingCount - k + 1)
		}
		prevNum = num
	}

	return arr[len(arr)-1] + k - missingCount
}
