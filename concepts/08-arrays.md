# 📦 Arrays in Go

## What is an Array?

An array is a **fixed-size collection** of elements of the **same type**.

Think of it as a **box with numbered compartments** - each compartment holds one value.

**Key Points:**
- Fixed size (cannot change after creation)
- All elements must be same type
- Access elements by index (0, 1, 2...)
- Size is part of the type definition

---

## Syntax

```go
var arrayName [size]Type
```

**Example:**

```go
var numbers [5]int      // Array of 5 integers
var names [3]string     // Array of 3 strings
```

---

## Declaring and Initializing Arrays

### Method 1: Declare then Assign

```go
package main

import "fmt"

func main() {
    var numbers [3]int
    numbers[0] = 10
    numbers[1] = 20
    numbers[2] = 30
    
    fmt.Println(numbers)  // [10 20 30]
}
```

### Method 2: Declare with Initial Values

```go
package main

import "fmt"

func main() {
    scores := [3]int{85, 90, 95}
    fmt.Println(scores)  // [85 90 95]
}
```

### Method 3: Auto-size (size inferred from values)

```go
package main

import "fmt"

func main() {
    fruits := [...]string{"Apple", "Banana", "Orange"}
    fmt.Println(len(fruits))  // 3
}
```

---

## Accessing Array Elements

Use **index** to access elements (starts from 0).

```go
package main

import "fmt"

func main() {
    colors := [4]string{"Red", "Green", "Blue", "Yellow"}
    
    fmt.Println(colors[0])  // Red
    fmt.Println(colors[1])  // Green
    fmt.Println(colors[3])  // Yellow
}
```

---

## Array Length

Use `len()` to get array length.

```go
package main

import "fmt"

func main() {
    numbers := [5]int{10, 20, 30, 40, 50}
    fmt.Println(len(numbers))  // 5
}
```

---

## Looping Through Arrays

### Method 1: Using For Loop with Index

```go
package main

import "fmt"

func main() {
    fruits := [3]string{"Apple", "Banana", "Orange"}
    
    for i := 0; i < len(fruits); i++ {
        fmt.Println(i, fruits[i])
    }
}
```

**Output:**

0 Apple

1 Banana

2 Orange

### Method 2: Using Range

```go
package main

import "fmt"

func main() {
    fruits := [3]string{"Apple", "Banana", "Orange"}
    
    for index, value := range fruits {
        fmt.Println(index, value)
    }
}
```

**Output:**

0 Apple

1 Banana

2 Orange

---

## Modifying Array Elements

You can **change existing elements** but **cannot add new elements**.

```go
package main

import "fmt"

func main() {
    numbers := [3]int{1, 2, 3}
    
    numbers[0] = 10  // OK - change existing
    numbers[1] = 20  // OK - change existing
    
    fmt.Println(numbers)  // [10 20 3]
}
```

---

## Array Size is Fixed

Arrays have **fixed size** - you cannot add more elements.

```go
arr := [3]int{1, 2, 3}
arr[3] = 4  // ERROR! Index out of range
```

This causes a runtime error.

---

## Default Values

Array elements have **default values** if not initialized.

```go
package main

import "fmt"

func main() {
    var numbers [3]int
    var names [2]string
    
    fmt.Println(numbers)  // [0 0 0]
    fmt.Println(names)    // [ ]
}
```

**Default values:**
- `int` → 0
- `string` → ""
- `bool` → false
- `float64` → 0.0

---

## Multi-dimensional Arrays

Arrays can have multiple dimensions.

```go
package main

import "fmt"

func main() {
    matrix := [2][3]int{
        {1, 2, 3},
        {4, 5, 6},
    }
    
    fmt.Println(matrix[0][0])  // 1
    fmt.Println(matrix[1][2])  // 6
}
```

---

## Comparing Arrays

```go
package main

import "fmt"

func main() {
    arr1 := [3]int{1, 2, 3}
    arr2 := [3]int{1, 2, 3}
    arr3 := [3]int{1, 2, 4}
    
    fmt.Println(arr1 == arr2)  // true
    fmt.Println(arr1 == arr3)  // false
}
```

---

---

## When to Use Arrays

Use arrays when:
- You know the exact size needed
- Size won't change
- You need fixed-size data structures

Use **slices** when:
- Size may change
- You need to add/remove elements
- (We'll learn slices next!)

---

## 💡 Memory Points

1. **Array** = Fixed-size collection of same type
2. **Syntax** = `[size]Type` or `[...]Type`
3. **Index** = Starts from 0
4. **Fixed size** = Cannot add more elements
5. **Length** = Use `len()` function
6. **Default values** = 0 for int, "" for string, false for bool
7. **Looping** = Use `for` with index or `range`
8. **Modification** = Can change existing elements
9. **Comparison** = Arrays can be compared with `==`
10. **Multi-dimensional** = Arrays can have multiple dimensions

---


# Arrays Practice Questions (Easy to Hard)

## Easy Problems (Basics)

### Question 1: Array Basics
Create an array of 5 numbers and print all elements using both for loop and range.

### Question 2: Sum of Array
Calculate the sum of all elements in an array.

### Question 3: Find Maximum
Find the maximum element in an array.

### Question 4: Find Minimum
Find the minimum element in an array.


### Question 5: Array Reversal
Reverse the order of elements in an array without using built-in functions.

---

## Medium Problems (Array Manipulation)

### Question 6: Search in Array
Search for a specific element in an array and return its index. Return -1 if not found.

### Question 7: Count Occurrences
Count how many times a specific element appears in an array.

### Question 9: Rotate Array
Rotate an array right and left by k positions.

### Question 10: Merge Two Arrays
Merge two sorted arrays into one sorted array.

## Hard Problems (Array techinique)

### Question 11: Two Sum 
Given an unsorted slice of integers nums and an integer target, return the indices of the two numbers such that they add up to the target.

### Question 12: Contains duplicates
Given an integer array nums, return true if any value appears more than once in the array, otherwise return false.

### Question 13: Move Zeros at end
You are given an integer array nums, move all 0's to the end of it while maintaining the relative order of the non-zero elements.

Note that you must do this in-place without making a copy of the array.

### Question 14: Best Time to Buy and Sell Stock
You are given an integer array prices where prices[i] is the price of NeetCoin on the ith day.

You may choose a single day to buy one NeetCoin and choose a different day in the future to sell it.

Return the maximum profit you can achieve. You may choose to not make any transactions, in which case the profit would be 0.

### Question 15: Maximum Subarray
Given an integer array nums, find the subarray with the largest sum, and return its sum.

### Question 16: Maximum Subarray with Average
Given an array of integers, find the maximum average of a contiguous subarray of size K.

### Question 17: Maximum Count of Target Elements
You are sitting in your garden looking at a long wooden fence. Birds are sitting on the fence posts. Some are beautiful Blue Jays (represented by the number 1), and some are other birds (represented by the number 0).
You have a special camera lens that can only see exactly three fence posts at a time (K=3). You want to slide your camera along the fence to find the spot where you can see the maximum number of Blue Jays at once!

### Question 18: The Gold Star Telescope (Counting Specific Targets of Size K)
You are looking at a row of slots in the night sky through a telescope. Each slot has something different:

7 = A Rare Gold Star (Our target!)

1 = A normal star

2 = A moon

0 = A dark cloud

Your telescope lens can only see exactly four slots at a time (K=4). You want to slide the telescope along the sky to find the view that has the maximum number of Rare Gold Stars (7s) at once.

### Question 19: Products of Array Except Self
Given an integer array nums, return an array output where output[i] is the product of all the elements of nums except nums[i].

Each product is guaranteed to fit in a 32-bit integer.

Follow-up: Could you solve it in 
O(n)
O(n) time without using the division operation?

Example 1:

Input: nums = [1,2,4,6]

Output: [48,24,12,8]

### Question 20: Remove Duplicates From Sorted Array
You are given an integer array nums sorted in non-decreasing order. Your task is to remove duplicates from nums in-place so that each element appears only once.

After removing the duplicates, return the number of unique elements, denoted as k, such that the first k elements of nums contain the unique elements.

Note:

- The order of the unique elements should remain the same as in the original array.
- It is not necessary to consider elements beyond the first k positions of the array.
- To be accepted, the first k elements of nums must contain all the unique elements.

Return k as the final result.


## How to use this section:
1. Try solving each question yourself first
2. If you get stuck, check the solution in `daily-progress/arrays/` folder
3. Compare your code with the solution
4. Understand the differences and learn from them

