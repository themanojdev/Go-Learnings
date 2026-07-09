# 📦 Packages and Imports

## What is a Package?

A package is a **folder containing Go files** that work together.

Think of it as a group of related functions and code organized in one place.

**Key Point:** Every Go file must start with `package packagename`

---

## What is go.mod?

`go.mod` is a **module file** that tells Go about your project.

**Why do you need it?**
- Enables custom imports in your project
- Tells Go what the project name is
- Without it, Go can't find your custom packages

**How to create `go.mod`:**

In your project root folder (`go-learning/`), run:
```bash
go mod init go-learning
```

This creates a `go.mod` file with:
module go-learning

**Important:** The module name in `go.mod` must match your import paths.

---

## Import Alias (Custom Import Names)

An import alias allows you to **give a custom name** to a package when importing it.

**Why use alias?**
- Simplify long import paths
- Make code more readable
- Avoid naming conflicts

**Syntax:**
```go
import (
    customname "actual/package/path"
)
```

**Example 1: Long Path with Alias**
```go
import (
    pkg "go-learning/concepts/example-code/01-packages-imports"
)

func main() {
    pkg.PrintHelloWorld()  // Using alias instead of full path
}
```

**Example 2: Renaming for Clarity**
```go
import (
    math1 "package/math"
    math2 "package/advanced-math"
)

func main() {
    result1 := math1.Add(10, 5)
    result2 := math2.ComplexAdd(10, 5)
}
```

**When to Use Alias:**
- Long folder paths (easier to type)
- Confusing package names (clearer intent)
- Avoiding naming conflicts
- Making code more readable

---

## Types of Packages

### **1. Main Package**

```go
package main
```

- Special package used to create **executable programs**
- Must have a `main()` function
- This is the program that runs directly
- Example: Your first Go program

### **2. Built-in Packages**

Packages created by Google:
- `fmt` - Print output, format strings
- `math` - Math operations
- `strings` - Work with text
- `time` - Date and time
- `os` - Operating system operations

### **3. Custom Packages**

Packages you create yourself to organize your code.

---

## What is Import?

Import is **bringing code from other packages** into your file so you can use it.

**Syntax:**
```go
import "packagename"
```

**Example:**
```go
import "fmt"
```

This brings the `fmt` package into your file.

---

## How to Use a Package

Once imported, call functions using:

packagename.FunctionName()

**Examples**

***Single Import:***
```go
package main

import "fmt"

func main() {
    // Calling a function from the imported 'fmt' package
    fmt.Println("Hello World") 
}
```

***Multiple Imports:***
When importing more than one package, use a factored `import` block enclosed in parentheses:

```go
package main

import (
    "fmt"
    "math"
)

func main() {
    fmt.Println("Calculating square root:")
    result := math.Sqrt(16)
    fmt.Println(result) // Output: 4
}
```

---

## Visibility: Capital vs Lowercase

Go uses **capitalization** to control what can be used outside a package:

| Capitalization | Visibility | Example |
|---|---|---|
| **Uppercase** | Public (can use from other packages) | `Add()`, `Hello()` |
| **lowercase** | Private (only inside package) | `add()`, `hello()` |

**Example:**
```go
package math

func Add(a, b int) int {      // Public - others can use this
    return a + b
}

func subtract(a, b int) int { // Private - only this package can use
    return a - b
}
```

---

## Custom Package Creation

Creating your own package to organize code.

**Steps:**

1. **Create a folder** for your package
2. **Create a Go file** with `package name` declaration
3. **Define public functions** (capital letters)
4. **Import the package** in your main file
5. **Use the functions** with `packagename.FunctionName()`

---

## 💡 Memory Points

1. **Package** = Folder with Go files that work together
2. **go.mod** = Module file needed for custom imports
3. **Main package** = Special package for executable programs
4. **Import** = Bringing code from other packages into your file
5. **Syntax:** `import "packagename"` or `import "modulename/packagename"`
6. **Alias** = Custom name for a package: `import alias "path"`
7. **Call function:** `packagename.FunctionName()`
8. **Uppercase** = Public, **lowercase** = Private
9. Every Go file must have `package` declaration at the top
10. Create `go.mod` with: `go mod init modulename`

---