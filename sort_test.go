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

func TestMany(t *testing.T) {
    /*
        Testing the our merge sort with the standard sorting algorithm
        Our MergeSort function should give us the same result as the standard
        sort function
    */

    // Default of 500 test cases of randomly generated list
    numTest := 500
    list := make([]TestCase, numTest)

    for i := 0; i < numTest; i++ {
        size := 15  // The length of the list
        rList := generateList(size) // A randomly generated list

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
