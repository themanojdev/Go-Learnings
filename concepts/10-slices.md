# 📋 Slices in Go

## What is a Slice?
A **slice** is a **dynamic, flexible collection** of elements of the same type that can grow and shrink as needed.

Unlike arrays which have **fixed size**, slices can **add and remove elements** during runtime.

**Simple Comparison:**
- **Array** = Box with fixed compartments (size cannot change)
- **Slice** = Flexible bag (can add/remove items)

**Key Points:**
- Variable size (can grow and shrink)
- All elements must be the same type
- Access elements by index (0, 1, 2...)
- Based on an underlying array
- More flexible than arrays
- Pass by reference to functions
- Most commonly used collection in Go

---

## Slice Structure (3 Components)
Every slice internally has 3 parts:
 
1. **Pointer** - Points to the underlying array
2. **Length (len)** - Number of elements currently in the slice
3. **Capacity (cap)** - Total space available without reallocating


```
Slice = [Pointer to Array] [Length] [Capacity]
```

---
 
## Syntax
 
### Without Size = Slice (Dynamic)
 
```go
var sliceName []Type
```

### With Size = Array (Fixed)
 
```go
var arrayName [size]Type
```

**Key Difference:**
 
```go
arr := [3]int{1, 2, 3}     // Array - size 3 (fixed)
slice := []int{1, 2, 3}    // Slice - no size (dynamic)
```
 
---

## Declaring Slices
 
### Method 1: Declare Empty Slice
 
```go
package main
 
import "fmt"
 
func main() {
    var numbers []int  // Empty slice
    
    fmt.Println(numbers)        // []
    fmt.Println(len(numbers))   // 0
    fmt.Println(cap(numbers))   // 0
}
```

**Output:**
 
```
[]
0
0
```

### Method 2: Initialize with Values
 
```go
package main
 
import "fmt"
 
func main() {
    fruits := []string{"Apple", "Banana", "Orange"}
    
    fmt.Println(fruits)         // [Apple Banana Orange]
    fmt.Println(len(fruits))    // 3
    fmt.Println(cap(fruits))    // 3
}
```

**Output:**
 
```
[Apple Banana Orange]
3
3
```

### Method 3: Using make() Function
 
`make()` creates a slice with specific length and capacity.
 
```go
package main
 
import "fmt"
 
func main() {
    // make([]Type, length, capacity)
    slice1 := make([]int, 3)        // Length 3, capacity 3
    slice2 := make([]int, 3, 5)     // Length 3, capacity 5
    
    fmt.Println("slice1:", slice1)
    fmt.Println("len:", len(slice1), "cap:", cap(slice1))
    
    fmt.Println("slice2:", slice2)
    fmt.Println("len:", len(slice2), "cap:", cap(slice2))
}
```
 
**Output:**
 
```
slice1: [0 0 0]
len: 3 cap: 3
slice2: [0 0 0]
len: 3 cap: 5
```

### Method 4: Slicing from Array or Slice
 
You can create a slice from a portion of an array or another slice using slice notation.
 
```go
package main
 
import "fmt"
 
func main() {
    arr := [5]int{10, 20, 30, 40, 50}
    
    slice := arr[1:4]  // From index 1 to 3 (4 is exclusive)
    fmt.Println(slice)  // [20 30 40]
}
```

**Output:**
 
```
[20 30 40]
```
 
---

## Length vs Capacity
 
**Length (len)** = Number of elements currently in the slice
 
**Capacity (cap)** = Maximum elements the slice can hold without reallocating memory
 
```go
package main
 
import "fmt"
 
func main() {
    slice := make([]int, 3, 5)
    
    fmt.Println("Length:", len(slice))     // 3
    fmt.Println("Capacity:", cap(slice))   // 5
    fmt.Println("Elements:", slice)        // [0 0 0]
    
    // Add 2 more elements
    slice = append(slice, 40, 50)
    
    fmt.Println("After append:")
    fmt.Println("Length:", len(slice))     // 5
    fmt.Println("Capacity:", cap(slice))   // 5
    fmt.Println("Elements:", slice)        // [0 0 0 40 50]
}
```
 
 **Output:**
 
```
Length: 3
Capacity: 5
Elements: [0 0 0]
After append:
Length: 5
Capacity: 5
Elements: [0 0 0 40 50]
```
 
---

## The Append Function
`append()` adds one or more elements to a slice and returns a new slice. **Important: You must reassign the result!**

```go
package main
 
import "fmt"
 
func main() {
    numbers := []int{1, 2, 3}
    
    fmt.Println("Before:", numbers)
    fmt.Println("Cap before:", cap(numbers))
    
    numbers = append(numbers, 4)  // Must reassign!
    
    fmt.Println("After:", numbers)
    fmt.Println("Cap after:", cap(numbers))
}
```
 
**Output:**
 
```
Before: [1 2 3]
Cap before: 3
After: [1 2 3 4]
Cap after: 6
```

### Method 2: Append Multiple Elements
 
```go
package main
 
import "fmt"
 
func main() {
    numbers := []int{1, 2, 3}
    
    numbers = append(numbers, 4, 5, 6)
    
    fmt.Println(numbers)  // [1 2 3 4 5 6]
}
```
 
**Output:**
 
```
[1 2 3 4 5 6]
```

### Method 3: Append One Slice to Another
 
```go
package main
 
import "fmt"
 
func main() {
    slice1 := []int{1, 2, 3}
    slice2 := []int{4, 5, 6}
    
    slice1 = append(slice1, slice2...)  // ... spreads slice2
    
    fmt.Println(slice1)  // [1 2 3 4 5 6]
}
```
 
**Output:**
 
```
[1 2 3 4 5 6]
```
 
---

## Understanding Capacity Growth
When `append()` exceeds the current capacity, Go creates a **new underlying array** with **doubled capacity** (usually). This is automatic!
 
```go
package main
 
import "fmt"
 
func main() {
    slice := make([]int, 0, 2)  // Length 0, capacity 2
    
    fmt.Println("Initial - len:", len(slice), "cap:", cap(slice))
    
    slice = append(slice, 1)
    fmt.Println("Add 1 - len:", len(slice), "cap:", cap(slice))
    
    slice = append(slice, 2)
    fmt.Println("Add 2 - len:", len(slice), "cap:", cap(slice))
    
    slice = append(slice, 3)  // Exceeds capacity!
    fmt.Println("Add 3 - len:", len(slice), "cap:", cap(slice))
    
    slice = append(slice, 4)
    fmt.Println("Add 4 - len:", len(slice), "cap:", cap(slice))
}
```

**Output:**
 
```
Initial - len: 0 cap: 2
Add 1 - len: 1 cap: 2
Add 2 - len: 2 cap: 2
Add 3 - len: 3 cap: 4
Add 4 - len: 4 cap: 4
```
 
---

## Important: Slices Share Underlying Array
Multiple slices can **point to the same underlying array**. Modifying elements affects all slices pointing to that array!
 
```go
package main
 
import "fmt"
 
func main() {
    arr := [5]int{10, 20, 30, 40, 50}
    
    slice1 := arr[0:3]  // [10 20 30]
    slice2 := arr[2:5]  // [30 40 50]
    
    fmt.Println("Before modification:")
    fmt.Println("slice1:", slice1)  // [10 20 30]
    fmt.Println("slice2:", slice2)  // [30 40 50]
    
    // Modify through slice1
    slice1[1] = 99
    
    fmt.Println("After modifying slice1[1] = 99:")
    fmt.Println("slice1:", slice1)  // [10 99 30]
    fmt.Println("slice2:", slice2)  // [30 99 50] - Also changed!
    fmt.Println("arr:", arr)        // [10 99 30 40 50]
}
```
 
**Output:**
 
```
Before modification:
slice1: [10 20 30]
slice2: [30 40 50]
After modifying slice1[1] = 99:
slice1: [10 99 30]
slice2: [30 99 50]
arr: [10 99 30 40 50]
```

--
 
## Copy Function
`copy()` creates a **true independent copy** of a slice that doesn't share the underlying array.

```go
package main
 
import "fmt"
 
func main() {
    slice1 := []int{1, 2, 3}
    slice2 := make([]int, len(slice1))
    
    copy(slice2, slice1)
    
    fmt.Println("Before modification:")
    fmt.Println("slice1:", slice1)  // [1 2 3]
    fmt.Println("slice2:", slice2)  // [1 2 3]
    
    // Modify slice1
    slice1[0] = 99
    
    fmt.Println("After modifying slice1[0] = 99:")
    fmt.Println("slice1:", slice1)  // [99 2 3]
    fmt.Println("slice2:", slice2)  // [1 2 3] - Not affected!
}
```

**Output:**
 
```
Before modification:
slice1: [1 2 3]
slice2: [1 2 3]
After modifying slice1[0] = 99:
slice1: [99 2 3]
slice2: [1 2 3]
```
 
---

## Deleting Elements from Slices
Go doesn't have a built-in delete function for slices. You need to use slicing to remove elements.

### Method 1: Delete by Index
Remove a single element at a specific index by using reslicing.
 
```go
package main
 
import "fmt"
 
func main() {
    slice := []int{10, 20, 30, 40, 50}
    
    index := 2  // Remove element at index 2 (value 30)
    
    // Reslice: combine everything before and after the index
    slice = append(slice[:index], slice[index+1:]...)
    
    fmt.Println(slice)  // [10 20 40 50]
}
```

**Output:**
 
```
[10 20 40 50]
```
 
**How it works:**
- `slice[:index]` = Elements before index (0 to 1)
- `slice[index+1:]` = Elements after index (3 to 4)
- `append()` combines them


### Method 2: Delete First Element
Remove the first element from a slice.
 
```go
package main
 
import "fmt"
 
func main() {
    slice := []int{10, 20, 30, 40, 50}
    
    fmt.Println("Before:", slice)
    
    slice = slice[1:]  // Remove first element
    
    fmt.Println("After:", slice)  // [20 30 40 50]
}
```
 
**Output:**
 
```
Before: [10 20 30 40 50]
After: [20 30 40 50]
```

### Method 3: Delete Last Element
Remove the last element from a slice.
 
```go
package main
 
import "fmt"
 
func main() {
    slice := []int{10, 20, 30, 40, 50}
    
    fmt.Println("Before:", slice)
    
    slice = slice[:len(slice)-1]  // Remove last element
    
    fmt.Println("After:", slice)  // [10 20 30 40]
}
```
 
**Output:**
 
```
Before: [10 20 30 40 50]
After: [10 20 30 40]
```

### Important: Delete Doesn't Free Memory
When you delete elements from a slice, Go **does not free the memory** of the underlying array immediately. The slice just points to a different range.

```go
package main
 
import "fmt"
 
func main() {
    slice := make([]int, 0, 5)
    slice = append(slice, 1, 2, 3)
    
    fmt.Println("Before delete:")
    fmt.Println("len:", len(slice), "cap:", cap(slice))
    
    slice = slice[:len(slice)-1]  // Delete last element
    
    fmt.Println("After delete:")
    fmt.Println("len:", len(slice), "cap:", cap(slice))
    // Capacity remains 5! Memory not freed.
}
```
 
**Output:**
 
```
Before delete:
len: 3 cap: 5
After delete:
len: 2 cap: 5
Capacity remains 5! Memory not freed.
```
 
---

## Looping Through Slices
 
### Method 1: Using For Loop with Index
 
```go
package main
 
import "fmt"
 
func main() {
    fruits := []string{"Apple", "Banana", "Orange"}
    
    for i := 0; i < len(fruits); i++ {
        fmt.Println(i, fruits[i])
    }
}
```

**Output:**
 
```
0 Apple
1 Banana
2 Orange
```

### Method 2: Using Range
 
```go
package main
 
import "fmt"
 
func main() {
    fruits := []string{"Apple", "Banana", "Orange"}
    
    for index, value := range fruits {
        fmt.Println(index, value)
    }
}
```
 
**Output:**
 
```
0 Apple
1 Banana
2 Orange
```
 
---

## Passing Slices to Functions
Slices are **passed by reference**, so functions can modify the slice's elements.
 
```go
package main
 
import "fmt"
 
func modifySlice(slice []int) {
    // Modify elements
    for i := 0; i < len(slice); i++ {
        slice[i] = slice[i] * 2
    }
}
 
func main() {
    numbers := []int{1, 2, 3}
    
    fmt.Println("Before:", numbers)  // [1 2 3]
    
    modifySlice(numbers)
    
    fmt.Println("After:", numbers)   // [2 4 6]
}
```
 
**Output:**
 
```
Before: [1 2 3]
After: [2 4 6]
```
 
---

## Slicing Notation
 
Extract a portion of a slice or array using `[start:end]` notation.
 
```go
package main
 
import "fmt"
 
func main() {
    slice := []int{10, 20, 30, 40, 50}
    
    fmt.Println(slice[1:3])   // [20 30] - from index 1 to 2
    fmt.Println(slice[:3])    // [10 20 30] - from start to 2
    fmt.Println(slice[2:])    // [30 40 50] - from index 2 to end
    fmt.Println(slice[:])     // [10 20 30 40 50] - entire slice
}
```
 
**Output:**
 
```
[20 30]
[10 20 30]
[30 40 50]
[10 20 30 40 50]
```
 
---

## Common Slice Operations
 
### Check if Slice is Empty
 
```go
slice := []int{}
 
if len(slice) == 0 {
    fmt.Println("Slice is empty")
}
```
 
### Find Element in Slice
 
```go
package main
 
import "fmt"
 
func contains(slice []int, target int) bool {
    for _, value := range slice {
        if value == target {
            return true
        }
    }
    return false
}
 
func main() {
    numbers := []int{1, 2, 3, 4, 5}
    
    fmt.Println(contains(numbers, 3))  // true
    fmt.Println(contains(numbers, 10)) // false
}
```
 
**Output:**
 
```
true
false
```
 
---

## Key Differences: Array vs Slice
 
| Feature | Array | Slice |
|---------|-------|-------|
| **Size** | Fixed at creation | Dynamic |
| **Syntax** | `[size]Type` | `[]Type` |
| **Can grow** | No | Yes (with append) |
| **Can shrink** | No | Yes |
| **Pass to function** | Copies entire array | Passes reference |
| **Capacity** | Always equals length | Can be >= length |
| **Underlying array** | Owns the array | May share array |
 
---

## When to Use Slices
 
Use **slices** when:
- Size is unknown or will change
- You need to add/remove elements dynamically
- You want efficient function arguments (no copying)
- Working with collections (most common in Go)

Use **arrays** when:
- Size is fixed and known at compile time
- You want guaranteed size (like RGB triplet [3]byte)
- You need the data on the stack for performance

---

## 💡 Memory Points
 
1. **Slice** = Dynamic collection, variable size
2. **Three parts** = Pointer, Length, Capacity
3. **Syntax** = `[]Type` (no size = slice, `[size]Type` = array)
4. **make()** = Creates slice with `make([]Type, length, capacity)`
5. **Length** = Current elements in slice (len())
6. **Capacity** = Maximum elements before reallocation (cap())
7. **append()** = Adds elements, doubles capacity when needed
8. **Must reassign** = `slice = append(slice, ...)` after add/delete
9. **Delete by index** = `slice = append(slice[:i], slice[i+1:]...)`
10. **Delete first** = `slice = slice[1:]`
11. **Delete last** = `slice = slice[:len(slice)-1]`
12. **Shared array** = Slices of same array share underlying data
13. **copy()** = Creates independent copy, breaks shared reference
14. **Delete doesn't free** = Memory remains, only length changes
 
 

