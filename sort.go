// Source for sorting algorithm: https://tinyurl.com/yxezgv5k

package main

import (
    "fmt"   // We would need this to print
    "math/rand" // Required for random generated list
    "time"  // Required for random generated list
    "sync"   // Required for WaitGroup if we are using recursion
)

func main() {
    size := 10  // Size of list

    // Test for list size of 10  10,000  10,000,000
    for i := 0; i < 3; i++ {
        list := generateList(size)   // Randomly generate list

        temp := make([]int, size)
        for i := 0; i < size; i++ {
            temp[i] = list[i]
        }

        // Want to test if goroutine will increase the speed

        fmt.Println("Testing with a list size of ", size)
        fmt.Println("--- Regular Merge Sort ---")
        // fmt.Println("Before:  ", list)
        start := time.Now()
        MergeSort(&list)
        duration := time.Since(start)
        // fmt.Println("After:  ", list)
        fmt.Println("Duration:  ", duration, "(microseconds)")

        fmt.Println()

        fmt.Println("--- Merge Sort w/ Goroutine ---")
        // fmt.Println("Before:  ", temp)
        start = time.Now()
        MergeSortGo(&temp)
        duration = time.Since(start)
        // fmt.Println("After:  ", temp)
        fmt.Println("Duration:  ", duration, "(microseconds)")
        fmt.Println("\n")

        size *= 1000
    }

    /*
        Explanation on why goroutine made merge sort slower for smaller list
        but faster for larger lists
        
        https://tinyurl.com/y3kcj2uz
    */
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
    /*
        This is a regular merge sort without the goroutine
    */

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
        left = (*list)[:middle]
        right = (*list)[middle:]

        // Recursively applying the MergeSort to the left and right sides
        MergeSort(&left)
        MergeSort(&right)

        (*list) = Merge(left, right)
    }
}

func MergeSortGo(list *[]int) {
    /*
        This is a merge sort with goroutine
    */

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
        left = (*list)[:middle]
        right = (*list)[middle:]

        var wg sync.WaitGroup
        wg.Add(1)

        // Recursively applying the MergeSort to the left and right sides
        go func() {
            defer wg.Done() // is the previous goroutine finish?
            MergeSort(&left)
        }()

        MergeSort(&right)

        wg.Wait()
        (*list) = Merge(left, right)
    }
}

func Merge(left, right []int) (sortedList []int) {
    /*
        Used to merge the left side with the right side
    */

    sortedList = make([]int, len(left) + len(right))    // return a sorted list

    i := 0
    for len(left) > 0 && len(right) > 0 {
        // Taking turns adding the smaller value into the sorted list
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

    return  // Return the newly sorted list
}
