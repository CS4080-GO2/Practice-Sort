package main

import (
    "fmt"   // We would need this to print
    "math/rand" // Required for random generated list
    "time"  // Required for random generated list
    "sync"   // Required for WaitGroup if we are using recursion
)

func main() {
    // Will use to generate random integers in list
    rand.Seed(time.Now().UnixNano())

    fmt.Println("+----------------------------------------+")
    fmt.Println("|               Merge Sort               |")
    fmt.Println("+----------------------------------------+")

    size := 100  // Size of list

    // Test for list size of 10  10,000  10,000,000
    for i := 0; i < 3; i++ {
        list := generateRandomList(size)   // Randomly generate list

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
        fmt.Println("Duration:  ", duration)

        fmt.Println()

        fmt.Println("--- Merge Sort w/ Goroutine ---")
        // fmt.Println("Before:  ", temp)
        start = time.Now()
        MergeSortGo(&temp)
        duration = time.Since(start)
        // fmt.Println("After:  ", temp)
        fmt.Println("Duration:  ", duration)
        fmt.Print("\n\n")

        size *= 100
    }

    /*
        Explanation on why goroutine made merge sort slower for smaller list
        but faster for larger lists

        Source:
            https://tinyurl.com/y3kcj2uz
    */


    // Time comparison to Quick sort
    fmt.Println("+----------------------------------------+")
    fmt.Println("|               Quick Sort               |")
    fmt.Println("+----------------------------------------+")

    size = 100  // Size of list

    // Test for list size of 10  10,000  10,000,000
    for i := 0; i < 3; i++ {
        list := generateRandomList(size)   // Randomly generate list

        temp := make([]int, size)
        for i := 0; i < size; i++ {
            temp[i] = list[i]
        }

        // Want to test if goroutine will increase the speed

        fmt.Println("Testing with a list size of ", size)
        fmt.Println("--- Regular Quick Sort ---")
        // fmt.Println("Before:  ", list)
        start := time.Now()
        QuickSort(&list)
        duration := time.Since(start)
        // fmt.Println("After:  ", list)
        fmt.Println("Duration:  ", duration)

        fmt.Println()

        fmt.Println("--- Quick Sort w/ Goroutine ---")
        // fmt.Println("Before:  ", temp)
        start = time.Now()
        QuickSortGo(&temp)
        duration = time.Since(start)
        // fmt.Println("After:  ", temp)
        fmt.Println("Duration:  ", duration)
        fmt.Print("\n\n")

        size *= 100
    }

    /*
        Merge sort is better than quick sort when it comes down to larger lists

        Sources:
            https://www.geeksforgeeks.org/quicksort-better-mergesort/
            https://tinyurl.com/y3pjob4z
    */
}

func generateRandomList(size int) (randList []int) {
    /*
        Used to generate the unsorted list
    */

    randList = make([]int, size)

    // Reassigning each element in list to a random integer
    for i := 0; i < size; i++ {
        // We subtract to give the probability for negative values
        // Integers from -99 to 99
        randList[i] = rand.Intn(100) - rand.Intn(100)
    }

    return
}

func MergeSort(list *[]int) {
    /*
        This is a regular merge sort without the goroutine
    */

    size := len(*list)   // Length of the unSortedList

    // If the size of the unSortedList is 1 than the list is already sorted
    if size > 1 {
        middle := int(size / 2) // Used to split the list into right and left side

        /*
            Assigned the left side of list to the left
            Assigned the right side of list to the right
        */
        left := (*list)[:middle]
        right := (*list)[middle:]

        // Recursively applying the MergeSort to the left and right sides
        MergeSort(&left)
        MergeSort(&right)

        // Merge the two list (left & right)
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

        /*
            Assigned the left side of list to the left
            Assigned the right side of list to the right
        */
        left := (*list)[:middle]
        right := (*list)[middle:]

        /*
            Since we are calling the goroutine recursively, we need to ensure
            that the previous goroutine is done before recursively calling the
            goroutine again.
        */
        var wg sync.WaitGroup
        wg.Add(1)

        // Anonymous goroutine function that complete the mergeSort left list
        go func() {
            defer wg.Done() // Is the previous goroutine finish?
            MergeSort(&left)
        }()

        // Recursively perform MergeSort on right side
        MergeSort(&right)

        // Wait for the goroutine to complete
        wg.Wait()

        // Once completed, merge the left list and right list together
        // This will result in the a sorted list
        (*list) = Merge(left, right)
    }
}

func Merge(left, right []int) (sortedList []int) {
    /*
        Used to merge the left side with the right side
        Used for merge sort
    */

    // The sorted list should be the size of the left + right side
    size := len(left) + len(right)
    sortedList = make([]int, size)    // return a sorted list

    /*
        i = for left list index
        j = for right list index
        k = for sortedList index
    */
    i, j, k := 0, 0, 0

    // Traverse through both left and right list
    for i < len(left) && j < len(right) {
        /*
            Look for the smaller value and add that into sortedList first
            Both left and right list should already be sorted in acending order
        */
        if left[i] < right[j] {
            /*
                If the value at current index of the left list is smaller than
                the value at current index of the right list, add the value of
                the current index of the left list into the sorted list
            */
            sortedList[k] = left[i]
            i++ // Move on to the next element in the left list
        } else {
            /*
                If the value at current index of the right list is smaller than
                the value at current index of the left list, add the value of
                the current index of the right list into the sorted list
            */
            sortedList[k] = right[j]
            j++ // Move on to the next element in the right list
        }
        k++ // Increment the index of the sortedList to add next smallest value
    }

    // For the remaining values in right list that are not in sortedList yet
    for j < len(right) {
        sortedList[k] = right[j]
        j++
        k++
    }

    // For the remaining values in left list that are not in sortedList yet
    for i < len(left) {
        sortedList[k] = left[i]
        i++
        k++
    }

    return  // Return the newly sorted list
}

func QuickSort(list *[]int) {
    /*
        Quick Sort implementation
    */

    size := len(*list)  // The size of the array

    if size > 1 {
        left := 0    // Index at the very left
        right := size - 1    // Index of the last element
        pivot := rand.Int() % len(*list) // Pivot idx is randomly selected

        // Swap the pivot value with the last element
        // Somewhat ensuring that the pivot is somewhere in middle of list
        (*list)[pivot], (*list)[right] = (*list)[right], (*list)[pivot]

        // Loop that iterate through list
        // If element is smaller than pivot, swap
        // Else continue to the next element
        for i, _ := range (*list) {
            if (*list)[i] < (*list)[right] {
                (*list)[left], (*list)[i] = (*list)[i], (*list)[left]
                left++
            }
        }

        (*list)[left], (*list)[right] = (*list)[right], (*list)[left]

        leftSubArr := make([]int, left)
        rightSubArr := make([]int, len(*list) - left)

        leftSubArr, rightSubArr = (*list)[:left], (*list)[left + 1:]

        QuickSort(&leftSubArr)
        QuickSort(&rightSubArr)
    }
}

func QuickSortGo(list *[]int) {
    /*
        Quick Sort with goroutine
    */

    size := len(*list)  // The size of the list

    if len(*list) > 1 {
        left := 0    // Index of the left most element
        right := size - 1    // Index of the list element in list
        pivot := rand.Intn(999) % len(*list) // Pivot idx is randomly selected

        // Swap the pivot value with the last element
        // Somewhat ensuring that the pivot is somewhere in middle of list
        (*list)[pivot], (*list)[right] = (*list)[right], (*list)[pivot]

        // Loop that iterate through list
        // If element is smaller than pivot, swap
        // Else continue to the next element
        for i, _ := range (*list) {
            if (*list)[i] < (*list)[right] {
                (*list)[left], (*list)[i] = (*list)[i], (*list)[left]
                left++
            }
        }

        (*list)[left], (*list)[right] = (*list)[right], (*list)[left]

        leftSubArr := make([]int, left)
        rightSubArr := make([]int, len(*list) - left)

        leftSubArr, rightSubArr = (*list)[:left], (*list)[left + 1:]

        /*
            Since we are calling the goroutine recursively, we need to ensure
            that the previous goroutine is done before recursively calling the
            goroutine again.
        */
        var wg sync.WaitGroup
        wg.Add(1)

        // Anonymous goroutine function that complete the quickSort left list
        go func() {
            defer wg.Done() // is the previous goroutine finish?
            QuickSort(&leftSubArr)
        }()

        QuickSort(&rightSubArr)
        wg.Wait()
    }
}
