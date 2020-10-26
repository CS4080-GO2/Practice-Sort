// Source for sorting algorithm: https://tinyurl.com/yxezgv5k

package main

import (
    "fmt"   // We would need this to print
    "math/rand" // Required for random generated list
    "time"  // Required for random generated list
    // "sync"   // Required for WaitGroup if we are using recursion
)

func main() {
    list := generateList(10)   // Randomly generate list
    fmt.Println("Before:  ", list)
    MergeSort(&list)
    fmt.Println("After:  ", list)
}

func generateList(size int) []int {
    /*
        Used to generate the unsorted list
    */

    list := make([]int, size)   // Creating a list of all 0's
    rand.Seed(time.Now().UnixNano())

    // Reassigning each element in list to a random integer
    for i := 0; i < size; i++ {
        list[i] = rand.Intn(999) - rand.Intn(999)
    }

    return list
}

func MergeSort(list *[]int) {
    size := len(*list)   // Length of the unSortedList

    // If the size of the unSortedList is 1 than the list is already sorted
    if size > 1 {
        middle := int(size / 2) // Used to split the list into right and left side

        // Creating list for the right and left side
        var (
                left = make([]int, middle)
                right = make([]int, size - middle)
            )

        /*
            Assigned the left side of list to the left
            Assigned the right side of list to the right
        */
        for i := 0; i < size; i++ {
            if i < middle {
                left[i] = (*list)[i]
            } else {
                right[i - middle] = (*list)[i]
            }
        }

        MergeSort(&left)
        MergeSort(&right)

        (*list) = Merge(left, right)
    }
}

func Merge(left, right []int) (sortedList []int) {
    sortedList = make([]int, len(left) + len(right))    // return a sorted list

    i := 0
    for len(left) > 0 && len(right) > 0 {
        if left[0] < right[0] {
            sortedList[i] = left[0]
            left = left[1:]
        } else {
            sortedList[i] = right[0]
            right = right[1:]
        }
        i++
    }

    for j := 0; j < len(left); j++ {
        sortedList[i] = left[j]
        i++
    }

    for j := 0; j < len(right); j++ {
        sortedList[i] = right[j]
        i++
    }

    return
}
