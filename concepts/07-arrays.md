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


**How to use this section:**
1. Try solving each question yourself first
2. If you get stuck, check the solution in `daily-progress/arrays/` folder
3. Compare your code with the solution
4. Understand the differences and learn from them

