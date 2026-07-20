# 🎯 Scope in Go

## What is Scope?
Scope is the **region of code where a variable or function is allowed to be used**. It defines the boundaries within which a variable exists and can be accessed.

Think of it as **invisible walls** that decide where you can and cannot use a variable.

**Key Points:**
- Determines where a variable can be accessed
- Different levels: package, function, and block
- Go uses lexical scoping (based on code structure)
- Variables can be hidden by shadowing
- Capitalization controls public/private access

---

## Three Levels of Scope in Go
 
### 1. Package Scope (Exported & Unexported)
 
Package scope determines whether a variable or function can be accessed from other packages or only within the same package.
 
**Rule: Capitalization matters!**
- **Uppercase (Exported)**: Can be accessed from other packages
- **Lowercase (Unexported)**: Can only be accessed within the same package


**Method 1: Exported Variables**
 
```go
package main
 
import "fmt"
 
// Exported - can be used from other packages
var GlobalName = "Manoj"
 
// Exported function
func HelloWorld() {
    fmt.Println("Hello")
}
 
func main() {
    fmt.Println(GlobalName) // Works
    HelloWorld()            // Works
}
```

**Method 2: Unexported Variables**
 
```go
package main
 
import "fmt"
 
// Unexported - only for this package
var secretKey = "123"
 
// Unexported function - only for this package
func privateHelper() {
    fmt.Println("Helper")
}
 
func main() {
    fmt.Println(secretKey)   // Works (same package)
    privateHelper()          // Works (same package)
}
```

### 2. Function Scope (Local Scope)
 
Variables declared inside a function can only be used within that function.
 
**Example:**
 
```go
package main
 
import "fmt"
 
func myFunction() {
    name := "Manoj"  // Local to this function
    fmt.Println(name) // Works
}
 
func anotherFunction() {
    fmt.Println(name) // ERROR! name doesn't exist here
}
 
func main() {
    myFunction()
}
```

### 3. Block Scope
 
Variables declared inside a `{}` block (like inside `if`, `for`, `switch`) can only be used within that block.
 
**Example:**
 
```go
package main
 
import "fmt"
 
func main() {
    x := 10 // Function scope
    fmt.Println(x) // Works
 
    if x > 5 {
        y := 20 // Block scope - only exists in this if block
        fmt.Println(y) // Works
    }
 
    fmt.Println(y) // ERROR! y doesn't exist outside the if block
}
```
 
---

## Variable Shadowing
 
Variable shadowing occurs when you declare a new variable with the same name in an inner scope, which **hides** the outer variable with the same name.
 
**Method 1: Simple Shadowing**
 
```go
package main
 
import "fmt"
 
func main() {
    name := "Manoj"  // Outer scope
    fmt.Println(name) // Prints: Manoj
 
    if true {
        name := "Dev"  // Inner scope - NEW variable (shadows outer)
        fmt.Println(name) // Prints: Dev
    }
 
    fmt.Println(name) // Prints: Manoj (outer variable, not affected)
}
```
 
**Output:**
 
Manoj
 
Dev
 
Manoj

**Method 2: Why Shadowing is Dangerous**
 
When you shadow a variable, you create a new variable instead of modifying the outer one:
 
```go
package main
 
import "fmt"
 
func main() {
    count := 10
    fmt.Println(count) // Prints: 10
 
    if count > 5 {
        count := 20  // Shadow! Created a NEW variable, didn't modify outer
        fmt.Println(count) // Prints: 20
    }
 
    fmt.Println(count) // Still prints: 10 (outer unchanged!)
}
```
 
**Output:**
 
10
 
20
 
10

**Method 3: How to Avoid Shadowing**
 
Use assignment (`=`) instead of declaration (`:=`) if you want to modify the outer variable:
 
```go
package main
 
import "fmt"
 
func main() {
    count := 10
 
    if count > 5 {
        count = 20  // Use = instead of := to modify outer variable
        fmt.Println(count) // Prints: 20
    }
 
    fmt.Println(count) // Prints: 20 (outer variable actually modified)
}
```
 
**Output:**
 
20
 
20

---

## Lexical Scoping
 
Lexical scoping means a variable's scope is determined by **WHERE it's declared in the code**, not WHERE it's called from. Go uses lexical scoping based on code structure (the `{}` braces).
 

 ### The Key Rule
 
**Inner scopes can access outer scopes, but outer scopes CANNOT access inner scopes.**
 
```
Global Scope (package level)
    ↓ Can access ↓
Function Scope
    ↓ Can access ↓
Block Scope (if, for, switch)
```

---

## When to Use Different Scopes
 
Use **Package Scope** when:
- You want to share data across functions in the same package
- You want to export data to other packages (use Uppercase)
- You want private data within a package (use lowercase)
Use **Function Scope** when:
- You need variables specific to one function
- You want to prevent other functions from accessing the data
- You want local calculations
Use **Block Scope** when:
- You need temporary variables in if/for/switch blocks
- You want to limit variable visibility to a small section
- You want to avoid conflicts with outer variables

---

## 💡 Memory Points
 
1. **Scope** = Region where a variable can be accessed
2. **Package Scope** = Uppercase (exported), Lowercase (unexported)
3. **Function Scope** = Variables inside function only
4. **Block Scope** = Variables inside {} blocks only
5. **Lexical Scoping** = Determined by code structure, not execution order
6. **Key Rule** = Inner can access outer, outer cannot access inner
7. **Shadowing** = Declaring same name in inner scope hides outer
8. **Avoid Shadowing** = Use `=` instead of `:=` to modify outer variable
9. **Capitalization** = Controls public/private at package level
10. **Why It Matters** = Prevents bugs, organizes code, controls access