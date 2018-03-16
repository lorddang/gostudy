package main

// 冒泡排序法
import (
	"fmt"
)

func main() {
	arr := [5]int{6, 1, 4, 5, 2}
	// selectSort(arr[:])
	// bubbleSort(arr[:])
	insertSort(arr[:])
	fmt.Println(arr)
}

func bubbleSort(slice []int) {
	fmt.Println(slice)
	for i := 0; i < len(slice)-1; i++ {
		for j := 0; j < len(slice)-1-i; j++ {
			if slice[j] > slice[j+1] {
				fmt.Printf("%v exchange with %v ", slice[j], slice[j+1])
				tmp := slice[j]
				slice[j] = slice[j+1]
				slice[j+1] = tmp

			}
			fmt.Printf("\t %v\n", slice)
		}
		fmt.Printf("%v\n", slice)
	}
	fmt.Println(slice)
}

func selectSort(slice []int) {
	fmt.Println(slice)
	for i := 0; i < len(slice)-1; i++ {

		for j := i + 1; j < len(slice); j++ {
			if slice[i] < slice[j] {
				tmp := slice[i]
				slice[i] = slice[j]
				slice[j] = tmp
			}
		}

	}

}


func insertSort(slice []int) {

	for i := 1; i < len(slice); i++ {
		temp := slice[i]
		var j int
		for j = i;j > 0 && temp <= slice[j - 1]; j -- {
					slice[j] = slice[j - 1]
		}
		slice[j] = temp
	}
	fmt.Println(slice)

}

func quickSort(slice []int)  {



}