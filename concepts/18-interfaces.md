# 🔌 Interfaces in Go

## What is it
An interface defines a **set of behaviors** (methods) without saying how they're implemented. It's a contract: "if you can do these things, you qualify" — no matter what concrete type you are.

**Key Points:**
- An interface lists method signatures only, no implementation
- Any type with those methods **automatically** satisfies the interface — no `implements` keyword needed
- This is called **implicit implementation**, and it's what makes Go interfaces different from Java/C#
- Interfaces let you write code that depends on **behavior**, not on a specific type

---

## Syntax
 
```go
type InterfaceName interface {
    MethodName1(params) returnType
    MethodName2(params) returnType
}
```

**Example:**
```go
type Shape interface {
    Area() float64
}
```
 
---

## Method 1: Implicit Implementation
Any type that has the required methods automatically satisfies the interface — nothing needs to be declared.
 
```go
type Circle struct {
    Radius float64
}
 
func (c Circle) Area() float64 {
    return 3.14 * c.Radius * c.Radius
}
 
type Rectangle struct {
    Width, Height float64
}
 
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}
```
 
Both `Circle` and `Rectangle` satisfy `Shape` simply because they have an `Area()` method.
 
---

## Method 3: Empty Interface — Accepting Anything
An empty interface has no required methods, so **every type** satisfies it.
 
```go
func describe(val interface{}) {
    fmt.Println(val)
}
```

In modern Go, `any` is an alias for `interface{}`:
```go
func describe(val any) { ... }
```
 
---

## Practical Code Example
 
```go
package main
 
import "fmt"
 
type Shape interface {
    Area() float64
}
 
type Circle struct {
    Radius float64
}
 
func (c Circle) Area() float64 {
    return 3.14 * c.Radius * c.Radius
}
 
type Rectangle struct {
    Width, Height float64
}
 
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}
 
func printArea(s Shape) {
    fmt.Println("Area:", s.Area())
}
 
func main() {
    c := Circle{Radius: 5}
    r := Rectangle{Width: 4, Height: 6}
 
    printArea(c)
    printArea(r)
}
```
 
**Output:**
```
Area: 78.5
Area: 24
```

`printArea` doesn't care whether it received a `Circle` or `Rectangle` — it only cares that the value has an `Area()` method.

---

## Type Assertion — Checking the Concrete Type Inside
 
```go
var val interface{} = "hello"
 
str, ok := val.(string)
if ok {
    fmt.Println("It's a string:", str)
} else {
    fmt.Println("Not a string")
}
```
 
**Output:**
```
It's a string: hello
```

`ok` works the same way as the map lookup pattern — `true` if the assertion worked, `false` if not, avoiding a panic.
 
---

## Type Switch — Checking Multiple Possible Types
 
```go
func describe(val interface{}) {
    switch v := val.(type) {
    case int:
        fmt.Println("It's an int:", v)
    case string:
        fmt.Println("It's a string:", v)
    default:
        fmt.Println("Unknown type")
    }
}
 
describe(42)
describe("hello")
describe(3.14)
```

**Output:**
```
It's an int: 42
It's a string: hello
Unknown type
```
 
---

## The Nil Interface Trap (Classic Interview Question)
 
```go
type MyError struct{}
 
func (e *MyError) Error() string {
    return "something went wrong"
}
 
func doSomething() error {
    var err *MyError = nil
    return err  // nil pointer, but wrapped in a non-nil interface
}
 
func main() {
    err := doSomething()
    if err != nil {
        fmt.Println("Error is NOT nil!")
    }
}
```
 
**Output:**
```
Error is NOT nil!
```

**Why:** An interface value is really two parts — a type and a value. Here the type is `*MyError` (not nil), even though the value is `nil`. So the interface itself is not nil, even though it looks empty.
 
**The fix:** return a plain `nil`, not a nil pointer wrapped in a concrete type:
```go
func doSomething() error {
    return nil
}
```
 
---

 
## Method Sets — Pointer vs Value Receiver and Interface Satisfaction
Whether a type satisfies an interface depends on which receiver its methods use.
 
```go
type Shape interface {
    Area() float64
}
 
type Circle struct {
    Radius float64
}
 
// pointer receiver
func (c *Circle) Area() float64 {
    return 3.14 * c.Radius * c.Radius
}
 
func main() {
    var s Shape = Circle{Radius: 5}  // ❌ ERROR: Circle does not implement Shape
}
```
This fails because `Area()` uses a **pointer receiver**, but a plain value (`Circle{}`) was assigned to the interface. Fix by using a pointer:
 
 ```go
var s Shape = &Circle{Radius: 5}  // ✅ works
```

**Rule to remember:**
> If any method on a type uses a pointer receiver, only a pointer to that type satisfies the interface — not the value itself.
 
---

## Interface Embedding
Like structs, interfaces can embed other interfaces to build bigger ones from smaller ones.
 
```go
type Reader interface {
    Read(p []byte) (n int, err error)
}
 
type Writer interface {
    Write(p []byte) (n int, err error)
}
 
type ReadWriter interface {
    Reader
    Writer
}
```

Any type that has both `Read()` and `Write()` automatically satisfies `ReadWriter` — no extra code needed.
 
---

## Small Interfaces from the Standard Library
Go favors small, focused interfaces — often just one method.
 
```go
type Stringer interface {
    String() string
}
```
 
`fmt.Stringer` controls how a type prints itself:
 
```go
type Person struct {
    Name string
}
 
func (p Person) String() string {
    return "Person: " + p.Name
}
 
func main() {
    p := Person{Name: "Alice"}
    fmt.Println(p)
}
```
 
**Output:**
```
Person: Alice
```

`io.Reader` and `io.Writer` follow the same idea — small, one-purpose interfaces that many types can satisfy.
 
---

## "Accept Interfaces, Return Structs" (Go Design Principle)
 
```go
// Good: accepts an interface - works with any Shape
func printArea(s Shape) { ... }
 
// Good: returns a concrete struct - clear to the caller what they get
func NewCircle(radius float64) Circle {
    return Circle{Radius: radius}
}
```

**Why:** Accepting interfaces as parameters keeps functions flexible (they work with any implementation, including test mocks). Returning concrete structs keeps the caller's code clear and predictable about what they're getting back.
 
---

## When to Use
 
- **Interfaces:** when a function should work with multiple types that share behavior, or when you need to swap real vs test implementations
- **Empty interface (`any`):** only when you genuinely need to accept absolutely any type (used sparingly)
- **Type assertion:** when you have one specific type in mind to check for
- **Type switch:** when a value could be one of several possible types
- **Interface embedding:** when building a larger interface out of smaller, focused ones
- **Pointer receivers:** required whenever interface satisfaction needs to work with a mutable pointer, not a copy

---

## Real-World / Project Usage
 
1. **Repository pattern** — define an interface for data access (`UserRepository`) so you can swap a real database for a fake one in tests
2. **HTTP handlers** — `http.Handler` is just an interface with one method (`ServeHTTP`), letting any type act as a handler
3. **Testing/mocking** — interfaces let you inject a fake implementation during unit tests without touching real dependencies
4. **io.Reader / io.Writer** — used everywhere for files, network connections, and buffers, all sharing the same small interface
5. **Dependency injection** — services depend on interfaces, not concrete structs, making the codebase easier to extend and test

---

## Interview Questions This Covers
 
- How do interfaces work internally in Go (type + value)?
- What's the difference between Go interfaces and Java/C# interfaces?
- What is the empty interface, and when would you use it?
- What's the difference between type assertion and type switch?
- Explain the nil interface vs nil pointer trap
- Why does a value receiver method behave differently from a pointer receiver method when satisfying an interface?
- What is interface embedding?
- What does "accept interfaces, return structs" mean, and why is it good practice?
- What are `io.Reader` / `io.Writer`, and why does Go prefer small interfaces?
---
 
## 💡 Memory Points
 
1. An interface is a contract listing method signatures only, no implementation
2. Go uses implicit implementation — no `implements` keyword needed
3. Any type with the required methods automatically satisfies the interface
4. Functions accepting an interface can work with any type that satisfies it
5. The empty interface `interface{}` (or `any`) accepts absolutely any type
6. Type assertion checks for one specific type: `val.(Type)`
7. Type switch checks for multiple possible types using `switch v := val.(type)`
8. Nil interface trap: a nil pointer wrapped in an interface is NOT a nil interface
9. Always return a plain `nil`, not a nil pointer of a concrete type, to avoid the nil interface trap
10. If any method uses a pointer receiver, only a pointer satisfies the interface, not the value
11. Interfaces can embed other interfaces to build larger ones from smaller ones
12. Go favors small, focused interfaces (often one method) like `io.Reader`, `io.Writer`, `fmt.Stringer`
13. Best practice: accept interfaces as parameters, return concrete structs
14. Interfaces are essential for testing — swap real implementations for mocks without changing calling code
 