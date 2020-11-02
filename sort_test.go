package main

import (
	"fmt"       // For Sprintf
	"math/rand" // Required for random generated list
	"reflect"   // List comparison function
	"sort"      // Will be using to test if our sort give same result as sort
	"testing"   // Required for Testing
	"time"      // Required for random generated list
)

type TestCase struct {
	Unsorted []int // The initial list
	Sorted   []int // The expected result
}

// Global variables
var (
	numTest int = 5                         // Number of test cases to run for each
	size    int = 20                        // Set the size of the list to generate
	list        = make([]TestCase, numTest) // List to hold all test cases
)

func TestMergeSort(t *testing.T) {
	/*
	   Testing the our merge sort with the standard sorting algorithm

	   Our MergeSort function should give us the same result as the standard
	   sort function.

	   Example:
	       if MergeSort(myList) == sort.Ints(myList)
	*/

	rand.Seed(time.Now().UnixNano()) // Seed used to generate random vals

	for i := 0; i < numTest; i++ {
		rList := generateRandomList(size) // A randomly generated list

		/*
		   temp is a copy of rList
		   rList will be sorted using our sorting function
		   temp will be sorted using standard implemented go sort function
		*/
		temp := make([]int, size)
		for i := 0; i < size; i++ {
			temp[i] = rList[i]
		}

		list[i].Unsorted = rList // Value is the unsorted list
		sort.Ints(temp)          // Use built-in sorting algorithm to sort
		list[i].Sorted = temp    // Assign Sorted to temp
	}

	// Iterate through the test cases
	for _, tc := range list {
		t.Run(fmt.Sprintf("%v", tc.Unsorted), func(t *testing.T) {
			// Use our MergeSort function we implemented
			MergeSort(&tc.Unsorted)

			/*
			   tc.Unsorted = actual sorted result
			   tc.Sorted = expected sorted result
			   Comparing if actual = expected, if so continue, else error
			*/
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

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < numTest; i++ {
		rList := generateRandomList(size) // A randomly generated list

		// temp is the sorted rList
		temp := make([]int, size)
		for i := 0; i < size; i++ {
			temp[i] = rList[i]
		}

		list[i].Unsorted = rList // Value is the unsorted list
		sort.Ints(temp)          // Use built-in sorting algorithm to sort
		list[i].Sorted = temp    // Assign Sorted to temp
	}

	// Iterate through the test cases
	for _, tc := range list {
		t.Run(fmt.Sprintf("%v", tc.Unsorted), func(t *testing.T) {
			// Use our MergeSortGo function we implemented
			MergeSortGo(&tc.Unsorted)

			/*
			   tc.Unsorted = actual sorted result
			   tc.Sorted = expected sorted result
			   Comparing if actual = expected, if so continue, else error
			*/
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

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < numTest; i++ {
		rList := generateRandomList(size) // A randomly generated list

		// temp is the sorted rList
		temp := make([]int, size)
		for i := 0; i < size; i++ {
			temp[i] = rList[i]
		}

		list[i].Unsorted = rList // Value is the unsorted list
		sort.Ints(temp)          // Use built-in sorting algorithm to sort
		list[i].Sorted = temp    // Assign Sorted to temp
	}

	// Iterate through the test cases
	for _, tc := range list {
		t.Run(fmt.Sprintf("%v", tc.Unsorted), func(t *testing.T) {
			// Use our QuickSort function we implemented
			QuickSort(&tc.Unsorted)

			/*
			   tc.Unsorted = actual sorted result
			   tc.Sorted = expected sorted result
			   Comparing if actual = expected, if so continue, else error
			*/
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

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < numTest; i++ {
		rList := generateRandomList(size) // A randomly generated list

		// temp is the sorted rList
		temp := make([]int, size)
		for i := 0; i < size; i++ {
			temp[i] = rList[i]
		}

		list[i].Unsorted = rList // Value is the unsorted list
		sort.Ints(temp)          // Use built-in sorting algorithm to sort
		list[i].Sorted = temp    // Assign Sorted to temp
	}

	// Iterate through the test cases
	for _, tc := range list {
		t.Run(fmt.Sprintf("%v", tc.Unsorted), func(t *testing.T) {
			// Use our QuickSortGo function we implemented
			QuickSortGo(&tc.Unsorted)

			/*
			   tc.Unsorted = actual sorted result
			   tc.Sorted = expected sorted result
			   Comparing if actual = expected, if so continue, else error
			*/
			if !reflect.DeepEqual(tc.Unsorted, tc.Sorted) {
				t.Fatalf("Not Equal - %v != %v", tc.Unsorted, tc.Sorted)
			}
		})
	}
}
func TestInsertionSort(t *testing.T) {
	/*
	   Testing our insertion sort with the standard sorting
	   algorithm.

	   Our insertion sort function should give us the same result
	   as the standard sort function.

	   Example:
	       if InsertionSort(myList) == sort.Ints(myList)
	*/

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < numTest; i++ {
		rList := generateRandomList(size) // A randomly generated list

		// temp is the sorted rList
		temp := make([]int, size)
		for i := 0; i < size; i++ {
			temp[i] = rList[i]
		}

		list[i].Unsorted = rList // Value is the unsorted list
		sort.Ints(temp)          // Use built-in sorting algorithm to sort
		list[i].Sorted = temp    // Assign Sorted to temp
	}

	// Iterate through the test cases
	for _, tc := range list {
		t.Run(fmt.Sprintf("%v", tc.Unsorted), func(t *testing.T) {
			// Use our QuickSortGo function we implemented
			InsertionSort(&tc.Unsorted)

			/*
			   tc.Unsorted = actual sorted result
			   tc.Sorted = expected sorted result
			   Comparing if actual = expected, if so continue, else error
			*/
			if !reflect.DeepEqual(tc.Unsorted, tc.Sorted) {
				t.Fatalf("Not Equal - %v != %v", tc.Unsorted, tc.Sorted)
			}
		})
	}
}
