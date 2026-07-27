# ⚡ Functions in Go

## What is a Function?
A **function** is a reusable block of code that performs a specific task. Instead of repeating code, you write it once as a function and call it whenever needed.

**Key Points:**
- Reusable block of code
- Can accept inputs (parameters)
- Can return outputs (return values)
- Helps organize code into logical pieces
- Every Go program starts with the `main()` function

---

```go
func functionName(parameters) returnType {
    // code
    return value
}
```

**Simple Example:**
 
```go
package main
 
import "fmt"
 
func greet() {
    fmt.Println("Hello, World!")
}
 
func main() {
    greet()  // Calling the function
}
```

**Output:**
 
```
Hello, World!
```
 
---

## Functions with Parameters
Parameters let you pass data into a function.
 
```go
package main
 
import "fmt"
 
func greetUser(name string) {
    fmt.Println("Hello,", name)
}
 
func main() {
    greetUser("Manoj")
    greetUser("Priya")
}
```
 
**Output:**
 
```
Hello, Manoj
Hello, Priya
```

### Multiple Parameters
 
```go
package main
 
import "fmt"
 
func add(a int, b int) {
    fmt.Println(a + b)
}
 
func main() {
    add(5, 10)
}
```
 
**Output:**
 
```
15
```

### Shorthand for Same-Type Parameters
If multiple parameters share the same type, you can simplify:
 
```go
func add(a, b int) {
    fmt.Println(a + b)
}
```
 
This is identical to `func add(a int, b int)`.
 
---

## Functions with Return Values
Functions can send data back using `return`.
 
```go
package main
 
import "fmt"
 
func add(a int, b int) int {
    return a + b
}
 
func main() {
    result := add(5, 10)
    fmt.Println(result)  // 15
}
```
 
**Output:**
 
```
15
```
 
---

## Multiple Return Values
Go allows functions to return **more than one value** - very commonly used!
 
```go
package main
 
import "fmt"
 
func divide(a int, b int) (int, int) {
    quotient := a / b
    remainder := a % b
    return quotient, remainder
}
 
func main() {
    q, r := divide(17, 5)
    fmt.Println("Quotient:", q, "Remainder:", r)
}
```
 
**Output:**
 
```
Quotient: 3 Remainder: 2
```

### Common Pattern: Returning Value and Error
This is used **everywhere** in Go (file reading, parsing, etc.):
 
```go
package main
 
import (
    "errors"
    "fmt"
)
 
func divide(a, b int) (int, error) {
    if b == 0 {
        return 0, errors.New("cannot divide by zero")
    }
    return a / b, nil
}
 
func main() {
    result, err := divide(10, 2)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Println("Result:", result)
    
    result, err = divide(10, 0)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Println("Result:", result)
}
```

**Output:**
 
```
Result: 5
Error: cannot divide by zero
```
 
---

## Named Return Values
You can name your return values - Go automatically returns them.
 
```go
package main
 
import "fmt"
 
func rectangleArea(length, width int) (area int) {
    area = length * width
    return  // "naked" return - returns "area" automatically
}
 
func main() {
    fmt.Println(rectangleArea(5, 3))  // 15
}
```
 
**Output:**
 
```
15
```

---

## Variadic Functions (Variable Number of Arguments)
A function that accepts **any number of arguments** using `...`.
 
```go
package main
 
import "fmt"
 
func sum(numbers ...int) int {
    total := 0
    for _, num := range numbers {
        total += num
    }
    return total
}
 
func main() {
    fmt.Println(sum(1, 2, 3))        // 6
    fmt.Println(sum(1, 2, 3, 4, 5))  // 15
    fmt.Println(sum())               // 0
}
```
 
**Output:**
 
```
6
15
0
```

### Passing a Slice to Variadic Function
 
```go
package main
 
import "fmt"
 
func sum(numbers ...int) int {
    total := 0
    for _, num := range numbers {
        total += num
    }
    return total
}
 
func main() {
    nums := []int{10, 20, 30}
    fmt.Println(sum(nums...))  // 60 (spread using ...)
}
```
 
**Output:**
 
```
60
```
 
---

## Anonymous Functions
Functions without a name, often used for one-time tasks.
 
```go
package main
 
import "fmt"
 
func main() {
    // Declare and call immediately
    func() {
        fmt.Println("I'm an anonymous function!")
    }()
    
    // Assign to a variable
    greet := func(name string) {
        fmt.Println("Hello,", name)
    }
    greet("Manoj")
}
```
 
**Output:**
 
```
I'm an anonymous function!
Hello, Manoj
```
 
---

## Functions as Values (First-Class Functions)
In Go, functions are **first-class citizens** - they can be assigned to variables, passed as arguments, and returned from other functions.
 
```go
package main
 
import "fmt"
 
func square(x int) int {
    return x * x
}
 
func main() {
    var operation func(int) int = square
    fmt.Println(operation(5))  // 25
}
```
 
**Output:**
 
```
25
```
 
---

## Functions as Parameters (Higher-Order Functions)
You can pass a function as an argument to another function.
 
```go
package main
 
import "fmt"
 
func applyOperation(a, b int, operation func(int, int) int) int {
    return operation(a, b)
}
 
func add(a, b int) int {
    return a + b
}
 
func multiply(a, b int) int {
    return a * b
}
 
func main() {
    fmt.Println(applyOperation(5, 3, add))       // 8
    fmt.Println(applyOperation(5, 3, multiply))  // 15
}
```
 
**Output:**
 
```
8
15
```

---

## Closures - Step by Step
A **closure** is a function that "remembers" variables from the place it was created, even after that outer function has finished running. Let's build up to this idea step by step.
 
### Step 1: Start with a Normal Function
A normal function has its own local variables that disappear once the function finishes.
 
```go
package main
 
import "fmt"
 
func normalFunction() {
    count := 0
    count++
    fmt.Println(count)
}
 
func main() {
    normalFunction()  // 1
    normalFunction()  // 1 (resets every time!)
    normalFunction()  // 1 (resets every time!)
}
```
 
**Output:**
 
```
1
1
1
```
Notice `count` always starts fresh at `0` every time we call `normalFunction()`. Once the function ends, `count` is gone from memory.

### Step 2: Make a Function Return Another Function
 
```go
package main
 
import "fmt"
 
func counter() func() int {
    count := 0  // This variable lives INSIDE counter()
    
    return func() int {
        count++
        return count
    }
}
 
func main() {
    fmt.Println("Just defining, not calling yet...")
}
```

**Output:**
 
```
Just defining, not calling yet...
```

Here, `counter()` doesn't print anything - it just **returns a function**. That inner function still refers to `count`, which belongs to `counter()`.

### Step 3: Call counter() to Get the Inner Function
 
```go
package main
 
import "fmt"
 
func counter() func() int {
    count := 0
    
    return func() int {
        count++
        return count
    }
}
 
func main() {
    increment := counter()  // increment now holds the inner function
    fmt.Println(increment)  // prints a memory address (it's a function)
}
```

At this point, `increment` is a variable holding the **inner function**. This inner function is still "connected" to the `count` variable from `counter()`, even though `counter()` has already finished running!

### Step 4: Call the Inner Function
 
```go
package main
 
import "fmt"
 
func counter() func() int {
    count := 0
    
    return func() int {
        count++
        return count
    }
}
 
func main() {
    increment := counter()
    
    fmt.Println(increment())  // 1
}
```
 
**Output:**
 
```
1
```

`count` was `0`, we called `increment()`, which ran `count++` (now `1`) and returned it.

### Step 5: Call It Again - The Magic Happens
 
```go
package main
 
import "fmt"
 
func counter() func() int {
    count := 0
    
    return func() int {
        count++
        return count
    }
}
 
func main() {
    increment := counter()
    
    fmt.Println(increment())  // 1
    fmt.Println(increment())  // 2
    fmt.Println(increment())  // 3
}
```
 
**Output:**
 
```
1
2
3
```

**This is the key insight:** Even though `counter()` already finished executing, the inner function **remembers** the `count` variable and keeps updating the *same* one every time you call `increment()`. Normally, local variables disappear once a function ends - but because the inner function references `count`, Go keeps `count` alive in memory for as long as the inner function still exists. This is called a **closure** - the inner function "closes over" the variable `count`.

### Step 6: Each Closure is Independent
If you call `counter()` again, you get a **brand new** `count` - it doesn't share state with the first one.
 
```go
package main
 
import "fmt"
 
func counter() func() int {
    count := 0
    
    return func() int {
        count++
        return count
    }
}
 
func main() {
    increment1 := counter()  // Gets its own count
    increment2 := counter()  // Gets a DIFFERENT count
    
    fmt.Println(increment1())  // 1
    fmt.Println(increment1())  // 2
    fmt.Println(increment2())  // 1 (independent from increment1!)
    fmt.Println(increment1())  // 3 (increment1 keeps its own count)
}
```
 
**Output:**
 
```
1
2
1
3
```
 
`increment1` and `increment2` each have their **own private copy** of `count`. This proves that every time `counter()` is called, a fresh, separate closure environment is created.

### Summary of the Steps
 
1. Normal functions lose their local variables when they end
2. A function can **return another function**
3. That returned function can reference variables from the outer function
4. Even after the outer function ends, those variables stay alive because the inner function still needs them
5. Every call to the outer function creates a **new, independent** set of variables
6. This "remembering" behavior is called a **closure**

### Another Closure Example: Multiplier
 
```go
package main
 
import "fmt"
 
func makeMultiplier(factor int) func(int) int {
    return func(x int) int {
        return x * factor
    }
}
 
func main() {
    double := makeMultiplier(2)
    triple := makeMultiplier(3)
    
    fmt.Println(double(5))  // 10
    fmt.Println(triple(5))  // 15
}
```
 
**Output:**
 
```
10
15
```
 
---

## Passing Slices, Maps, and Pointers to Functions

### Passing a Slice (Reference Type)
Since slices are reference types, changes inside a function affect the original.
 
```go
package main
 
import "fmt"
 
func modifySlice(s []int) {
    s[0] = 100  // Modifies original slice
}
 
func main() {
    slice := []int{1, 2, 3}
    modifySlice(slice)
    fmt.Println(slice)  // [100 2 3] - changed!
}
```
 
**Output:**
 
```
[100 2 3]
```

### Passing a Value (Pass by Value)
 
```go
package main
 
import "fmt"
 
func modifyValue(x int) {
    x = 100  // Only changes the local copy
}
 
func main() {
    num := 5
    modifyValue(num)
    fmt.Println(num)  // 5 - unchanged!
}
```
 
**Output:**
 
```
5
```
 
 ### Passing a Pointer (Pass by Reference)
To let a function modify a plain value type like `int`, pass a **pointer** to it instead.
 
```go
package main
 
import "fmt"
 
func modifyPointer(x *int) {
    *x = 100  // Dereference and change the actual value
}
 
func main() {
    num := 5
    modifyPointer(&num)  // Pass the address of num
    fmt.Println(num)     // 100 - changed!
}
```
 
**Output:**
 
```
100
```

### Side-by-Side Comparison: Value vs Pointer
 
```go
package main
 
import "fmt"
 
func modifyValue(x int) {
    x = 999
}
 
func modifyPointer(x *int) {
    *x = 999
}
 
func main() {
    a := 5
    b := 5
    
    modifyValue(a)
    modifyPointer(&b)
    
    fmt.Println("a (pass by value):", a)     // 5 - unchanged
    fmt.Println("b (pass by pointer):", b)   // 999 - changed
}
```
 
**Output:**
 
```
a (pass by value): 5
b (pass by pointer): 999
```
 
### Why Use Pointers in Functions?
 
1. **Modify the original value** - Necessary when you want changes to persist outside the function
2. **Avoid copying large data** - Passing a pointer to a large struct is cheaper than copying the whole struct
3. **Signal optionality** - A pointer can be `nil`, representing "no value," which a plain value type cannot

---

## Defer with Functions (Quick Recap)
 
`defer` delays execution until the function returns.
 
```go
package main
 
import "fmt"
 
func process() {
    defer fmt.Println("Cleanup done")
    fmt.Println("Processing...")
}
 
func main() {
    process()
}
```
 
**Output:**
 
```
Processing...
Cleanup done
```
 
---
 
## When to Use What
 
Use **regular functions** when:
- You have reusable logic used in multiple places
Use **multiple return values** when:
- You need to return a result and an error (very common pattern)
Use **variadic functions** when:
- You don't know how many arguments will be passed (like `fmt.Println`)
Use **anonymous functions** when:
- You need a quick, one-time function (often with goroutines or defer)
Use **closures** when:
- You need a function that maintains state between calls
- Building things like counters, generators, or middleware
Use **pointers as parameters** when:
- You need to modify the original value inside a function
- You want to avoid copying large data structures

---

## 💡 Memory Points
 
1. **Function** = Reusable block of code, declared with `func`
2. **Parameters** = Inputs passed to a function
3. **Return values** = Outputs sent back from a function
4. **Multiple returns** = Go supports returning more than one value, commonly `(result, error)`
5. **Named returns** = Declared in the signature, returned automatically with a naked `return`
6. **Variadic functions** = Accept unlimited arguments using `...Type`
7. **Anonymous functions** = Functions without a name, often used inline
8. **First-class functions** = Functions can be stored in variables, passed, and returned
9. **Closure** = Inner function that "remembers" and keeps alive variables from its outer function
10. **Independent closures** = Each call to the outer function creates a separate, private variable set
11. **Pass by value** = Function gets a copy, original is unchanged (plain types)
12. **Pass by reference** = Slices/maps share underlying data; pointers let you modify plain values directly
 