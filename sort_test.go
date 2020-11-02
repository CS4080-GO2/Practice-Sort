package main

import (
    "testing"
    "sort"
    "fmt"
    "reflect"
)

type TestCase struct {
    Unsorted []int  // The initial list
    Sorted []int    // The expected result
}

func TestMergeSort(t *testing.T) {
    /*
        Testing the our merge sort with the standard sorting algorithm

        Our MergeSort function should give us the same result as the standard
        sort function.

        Example:
            if MergeSort(myList) == sort.Ints(myList)
    */

    // Default of 500 test cases of randomly generated list
    numTest := 500
    list := make([]TestCase, numTest)

    for i := 0; i < numTest; i++ {
        size := 15  // The length of the list
        rList := generateRandomList(size) // A randomly generated list

        // temp is the sorted rList
        temp := make([]int, size)
        for i := 0; i < size; i++ {
            temp[i] = rList[i]
        }

        list[i].Unsorted = rList    // Value is the unsorted list
        sort.Ints(temp)             // Use built-in sorting algorithm to sort
        list[i].Sorted = temp       // Assign Sorted to temp
    }

    // Iterate through the test cases
    for _, tc := range list {
        t.Run(fmt.Sprintf("%v", tc.Unsorted), func(t *testing.T) {
            MergeSort(&tc.Unsorted) // Use our MergeSort function we implemented

            if !reflect.DeepEqual(tc.Unsorted, tc.Sorted) {
                t.Fatalf("Not Equal - %v != %v", tc.Unsorted, tc.Sorted)
            }
        })
    }
}

func TestMergeSortGo(t *testing.T) {
    /*
        Testing the our merge sort (w/ goroutine) with the standard sorting
        algorithm.

        Our MergeSort function (w/ goroutine) should give us the same result
        as the standard sort function.

        Example:
            if MergeSortGo(myList) == sort.Ints(myList)
    */

    // Default of 500 test cases of randomly generated list
    numTest := 500
    list := make([]TestCase, numTest)

    for i := 0; i < numTest; i++ {
        size := 15  // The length of the list
        rList := generateRandomList(size) // A randomly generated list

        // temp is the sorted rList
        temp := make([]int, size)
        for i := 0; i < size; i++ {
            temp[i] = rList[i]
        }

        list[i].Unsorted = rList    // Value is the unsorted list
        sort.Ints(temp)             // Use built-in sorting algorithm to sort
        list[i].Sorted = temp       // Assign Sorted to temp
    }

    // Iterate through the test cases
    for _, tc := range list {
        t.Run(fmt.Sprintf("%v", tc.Unsorted), func(t *testing.T) {
            MergeSortGo(&tc.Unsorted) // Use our MergeSort function we implemented

            if !reflect.DeepEqual(tc.Unsorted, tc.Sorted) {
                t.Fatalf("Not Equal - %v != %v", tc.Unsorted, tc.Sorted)
            }
        })
    }
}

func TestQuickSort(t *testing.T) {
    /*
        Testing the our quick sort with the standard sorting algorithm

        Our QuickSort function should give us the same result as the standard
        sort function.

        Example:
            if QuickSort(myList) == sort.Ints(myList)
    */

    // Default of 500 test cases of randomly generated list
    numTest := 500
    list := make([]TestCase, numTest)

    for i := 0; i < numTest; i++ {
        size := 15  // The length of the list
        rList := generateRandomList(size) // A randomly generated list

        // temp is the sorted rList
        temp := make([]int, size)
        for i := 0; i < size; i++ {
            temp[i] = rList[i]
        }

        list[i].Unsorted = rList    // Value is the unsorted list
        sort.Ints(temp)             // Use built-in sorting algorithm to sort
        list[i].Sorted = temp       // Assign Sorted to temp
    }

    // Iterate through the test cases
    for _, tc := range list {
        t.Run(fmt.Sprintf("%v", tc.Unsorted), func(t *testing.T) {
            QuickSort(&tc.Unsorted) // Use our MergeSort function we implemented

            if !reflect.DeepEqual(tc.Unsorted, tc.Sorted) {
                t.Fatalf("Not Equal - %v != %v", tc.Unsorted, tc.Sorted)
            }
        })
    }
}

func TestQuickSortGo(t *testing.T) {
    /*
        Testing the our quick sort (w/ goroutine) with the standard sorting
        algorithm.

        Our QuickSort function (w/ goroutine) should give us the same result
        as the standard sort function.

        Example:
            if QuickSortGo(myList) == sort.Ints(myList)
    */

    // Default of 500 test cases of randomly generated list
    numTest := 500
    list := make([]TestCase, numTest)

    for i := 0; i < numTest; i++ {
        size := 15  // The length of the list
        rList := generateRandomList(size) // A randomly generated list

        // temp is the sorted rList
        temp := make([]int, size)
        for i := 0; i < size; i++ {
            temp[i] = rList[i]
        }

        list[i].Unsorted = rList    // Value is the unsorted list
        sort.Ints(temp)             // Use built-in sorting algorithm to sort
        list[i].Sorted = temp       // Assign Sorted to temp
    }

    // Iterate through the test cases
    for _, tc := range list {
        t.Run(fmt.Sprintf("%v", tc.Unsorted), func(t *testing.T) {
            QuickSortGo(&tc.Unsorted) // Use our MergeSort function we implemented

            if !reflect.DeepEqual(tc.Unsorted, tc.Sorted) {
                t.Fatalf("Not Equal - %v != %v", tc.Unsorted, tc.Sorted)
            }
        })
    }
}
