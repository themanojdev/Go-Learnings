# 📍 Pointers in Go

## What is a Pointer?
A **pointer** is a variable that **stores the memory address** of another variable instead of storing the value itself.

Think of it as **a sign that points to where your data lives in memory**, rather than the data itself.

**Key Points:**
- Stores memory address of another variable
- Declared with `*` symbol
- Access address using `&` (address-of operator)
- Access value using `*` (dereference operator)
- Allows you to modify original variables in functions
- Can be `nil` (no address assigned)

---

## Syntax
 
```go
var pointerName *Type
```
 
**The `*` means "pointer to Type"**

**Example:**
 
```go
var ptr *int        // Pointer to int
var namePtr *string // Pointer to string
var boolPtr *bool   // Pointer to bool
```

---

## Two Important Operators

### The `&` Operator (Address-of)
 
The `&` operator gives you the **memory address** of a variable.

```go
package main
 
import "fmt"
 
func main() {
    x := 10
    
    fmt.Println(x)   // 10 (the value)
    fmt.Println(&x)  // 0xc0000140b8 (the address)
}
```

**Output:**
 
```
10
0xc0000140b8
```

### The `*` Operator (Dereference)
 
The `*` operator **accesses the value** that a pointer points to.
 
```go
package main
 
import "fmt"
 
func main() {
    x := 10
    ptr := &x
    
    fmt.Println(*ptr)  // 10 (the value that ptr points to)
}
```
 
**Output:**
 
```
10
```

---

## Declaring and Using Pointers
 
### Method 1: Declare Pointer and Assign Address
 
```go
package main
 
import "fmt"
 
func main() {
    age := 25
    
    var ptr *int
    ptr = &age
    
    fmt.Println("Variable:", age)       // 25
    fmt.Println("Pointer:", ptr)        // 0xc000010080
    fmt.Println("Value via pointer:", *ptr) // 25
}
```

**Output:**
 
```
Variable: 25
Pointer: 0xc000010080
Value via pointer: 25
```

### Method 3: Modifying Variables Through Pointers
 
When you modify a value through a pointer using `*ptr`, you change the **original variable**.


```go
package main
 
import "fmt"
 
func main() {
    x := 10
    ptr := &x
    
    fmt.Println("Before:", x)  // 10
    
    *ptr = 20  // Change value through pointer
    
    fmt.Println("After:", x)   // 20
}
```
 
**Output:**
```
Before: 10
After: 20
```

---

## Pointers in Functions
 
Pointers are powerful for **modifying the original variable** inside a function.

### Without Pointers (Pass by Value)
 
When you pass a variable to a function, Go creates a **copy**. Changes don't affect the original.
 
```go
package main
 
import "fmt"
 
func increaseValue(num int) {
    num = num + 10  // Modifies only the copy
}
 
func main() {
    x := 5
    fmt.Println("Before:", x)  // 5
    
    increaseValue(x)
    
    fmt.Println("After:", x)   // Still 5 (not changed!)
}
```

**Output:**
 
```
Before: 5
After: 5
```
 
 ### With Pointers (Pass by Reference)
 
When you pass a **pointer** to a function, the function can **modify the original variable**.
 
```go
package main
 
import "fmt"
 
func increasePointer(ptr *int) {
    *ptr = *ptr + 10  // Modifies the original variable
}
 
func main() {
    x := 5
    fmt.Println("Before:", x)  // 5
    
    increasePointer(&x)
    
    fmt.Println("After:", x)   // 15 (changed!)
}
```

**Output:**
 
```
Before: 5
After: 15
```

---

## Real-World Example: Swap Function
 
```go
package main
 
import "fmt"
 
func swap(a *int, b *int) {
    // Swap the values using pointers
    temp := *a
    *a = *b
    *b = temp
}
 
func main() {
    x := 10
    y := 20
    
    fmt.Println("Before - x:", x, "y:", y)  // x: 10 y: 20
    
    swap(&x, &y)
    
    fmt.Println("After - x:", x, "y:", y)   // x: 20 y: 10
}
```
 
**Output:**
 
```
Before - x: 10 y: 20
After - x: 20 y: 10
```
 
---

## Nil Pointers
 
A pointer that hasn't been assigned an address is `nil`.
 
```go
package main
 
import "fmt"
 
func main() {
    var ptr *int
    
    fmt.Println(ptr)  // <nil>
    
    // Check if pointer is nil
    if ptr == nil {
        fmt.Println("Pointer is nil")
    }
}
```

**Output:**
 
```
<nil>
Pointer is nil
```
 
**Warning:** Never dereference a nil pointer!
 
```go
var ptr *int
fmt.Println(*ptr)  // ERROR! Panic - nil pointer dereference
```
 
---

## Pointers to Pointers
 
A pointer can point to another pointer.
 
```go
package main
 
import "fmt"
 
func main() {
    x := 10
    
    ptr1 := &x      // Pointer to x
    ptr2 := &ptr1   // Pointer to ptr1
    
    fmt.Println("x:", x)           // 10
    fmt.Println("*ptr1:", *ptr1)   // 10
    fmt.Println("**ptr2:", **ptr2) // 10 (double dereference)
}
```
 
**Output:**
 
```
x: 10
*ptr1: 10
**ptr2: 10
```
 
---

## Common Mistakes
 
### Mistake 1: Forgetting `&` When Assigning to Pointer
 
```go
package main
 
func main() {
    x := 10
    var ptr *int
    
    ptr = x  // ERROR! Can't assign int to *int
    ptr = &x // Correct! Assign address of x
}
```
 
### Mistake 2: Dereferencing Nil Pointer
 
```go
package main
 
func main() {
    var ptr *int
    
    fmt.Println(*ptr)  // ERROR! Nil pointer dereference panic
}
```

### Mistake 3: Confusing Declaration `*` with Dereference `*`
 
```go
var ptr *int    // * means "pointer to int type" (declaration)
value := *ptr   // * means "get the value" (dereference)
```
 
---

## Predict the output

```go
  a := 10
  b := &a
  *b = *b + 5
  c := b
  *c = 100
  fmt.Println(a, *b, *c)
```
```
Q1(interview classic): We have written in our notes "pass by reference" for pointers — but here's the twist interviewers love: Go is always pass-by-value, even for pointers. When you call increasePointer(&x), what exactly gets copied, and why can the function still change x even though it received a copy?
```

## When to Use Pointers
 
Use **pointers** when:
- You need to **modify the original variable** in a function
- You want to **avoid copying large data** (like big structs or slices)
- You need to work with **nil values** (indicate absence of value)
- You need **dynamic memory management**
- You're building **linked lists, trees, or graphs**
Avoid **unnecessary pointers** when:
- The function only needs to **read** the value
- The data type is **small** (int, string, bool)
- You want **simple, readable code**
---

## 💡 Memory Points
 
1. **Pointer** = Variable storing memory address of another variable
2. **Declare Pointer** = `var ptr *Type`
3. **Address-of `&`** = Get memory address of a variable
4. **Dereference `*`** = Access value that pointer points to
5. **Pass by Value** = Function gets a copy, can't modify original
6. **Pass by Reference** = Function gets pointer, can modify original
7. **Modify via Pointer** = Use `*ptr = newValue` to change original
8. **Nil Pointer** = Pointer with no address assigned, equals `nil`
9. **Never Dereference Nil** = Always check `if ptr == nil` before using
10. **Use Sparingly** = Only use pointers when necessary for clarity and efficiency

