# 🏷️ Custom Types in Go (type keyword)

## What is it
The `type` keyword lets you create a **new named type** based on an existing underlying type. This new type is distinct from its underlying type — Go treats them differently even though the underlying data is the same.

**Key Points:**
- `type` declares a new type or gives a name to an existing type structure
- A custom type has its own identity, separate from its underlying type
- Custom types allow you to attach **methods** — something you cannot do directly to built-in types like `float64` or `string`
- There is a difference between a **custom type** and a **type alias**

---

## Syntax
 
```go
// Custom type - creates a new distinct type
type TypeName UnderlyingType
 
// Type alias - creates another name for the SAME type
type TypeName = UnderlyingType
```
 
---

## Method 1: Basic Custom Type
 
```go
type Celsius float64
```
 
This creates a new type `Celsius` whose underlying type is `float64`. `Celsius` is now its own type — not interchangeable with plain `float64` without conversion.

```go
var temp Celsius = 25
var normal float64 = 25
 
// temp = normal        // ❌ ERROR: mismatched types
temp = Celsius(normal)  // ✅ works with explicit conversion
```
 
---

## Method 2: Type Alias (using `=`)
 
```go
type Celsius = float64
```
 
With `=`, `Celsius` is just another **name** for `float64` — they are fully interchangeable, no conversion needed.
 
```go
var temp Celsius = 25
var normal float64 = 25
 
temp = normal  // ✅ works fine, they're the same type
```
 
---

## Method 3: Custom Type on Struct
 
`type` is also used to name struct definitions — this is the most common use in real projects.
 
```go
type User struct {
    Name string
    Age  int
}
```
 
Here, `User` is a custom type built from an anonymous struct structure.
 
---

```go
package main
 
import "fmt"
 
type Celsius float64
type Fahrenheit float64
 
func (c Celsius) ToFahrenheit() Fahrenheit {
    return Fahrenheit(c*9/5 + 32)
}
 
func main() {
    var temp Celsius = 25
    result := temp.ToFahrenheit()
 
    fmt.Println("Celsius:", temp)
    fmt.Println("Fahrenheit:", result)
}
```
 
**Output:**
```
Celsius: 25
Fahrenheit: 77
```
 
---

## Custom Type vs Type Alias Comparison
 
| Aspect | Custom Type (`type X Y`) | Type Alias (`type X = Y`) |
|---|---|---|
| Relationship | New, distinct type | Same type, different name |
| Conversion needed | Yes, when mixing with underlying type | No, fully interchangeable |
| Can attach methods | Yes | Yes (methods go to underlying type) |
| Common use | Domain-specific types (Celsius, UserID) | Gradual code refactors, renaming |
 
---

## When to Use
 
- **Custom Type:** When you want type safety and meaning — e.g., preventing a `UserID` from being mixed up with a plain `int`, or attaching methods to a value
- **Type Alias:** When refactoring large codebases and you need a temporary or permanent alternate name without changing behavior

---

## Real-World / Project Usage
 
1. **Domain Modeling** — `type UserID int`, `type OrderStatus string` make code self-documenting and prevent accidental misuse of raw types
2. **Enums (Go has no native enum)** — custom types + constants simulate enums:
```go
   type Status string
 
   const (
       Active   Status = "active"
       Inactive Status = "inactive"
   )
```
3. **Attaching Behavior** — custom types let you add validation or formatting methods directly to simple values (e.g., `Celsius.ToFahrenheit()`)
4. **Interface Satisfaction** — custom types are often created specifically so methods can be attached, allowing the type to implement an interface
5. **Struct Naming** — every struct definition in Go uses `type` to give it a name, making it usable throughout the codebase

---

## Interview Questions This Covers
 
- What does the `type` keyword do in Go?
- What is the difference between a custom type and a type alias?
- Why would you create a custom type instead of using the underlying type directly?
- How does Go simulate enums without native enum support?
- Can you attach methods to primitive types like `int` or `string` directly? (No — only through a custom type)

---

## 💡 Memory Points
 
1. `type` creates a new named type based on an underlying type
2. `type X Y` → distinct type, requires conversion to mix with `Y`
3. `type X = Y` → alias, fully interchangeable with `Y`, no conversion needed
4. Custom types allow attaching methods; built-in types alone cannot have methods
5. Struct definitions always use `type` to give the struct a usable name
6. Custom types are the backbone of Go's simulated enums (type + const block)
7. Domain-specific types (`UserID`, `Celsius`, `OrderStatus`) improve readability and prevent bugs
8. Type aliases are mainly used during refactors, not for everyday type safety
9. Explicit conversion syntax: `TypeName(value)`
10. Custom types are essential for satisfying interfaces when you need to attach behavior to a value
11. The underlying type determines the custom type's storage and default operations (math, comparisons, etc.)
12. Even though `Celsius` is based on `float64`, `fmt.Println` still prints it like a normal number unless a `String()` method is defined